/*
 *
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package client

import (
	"context"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/internal/buffer"
	"google.golang.org/grpc/internal/grpclog"
)

// ErrResourceTypeUnsupported is an error used to indicate an unsupported xDS
// resource type. The wrapped ErrStr contains the details.
type ErrResourceTypeUnsupported struct {
	ErrStr string
}

// Error helps implements the error interface.
func (e ErrResourceTypeUnsupported) Error() string {
	return e.ErrStr
}

// VersionedClient is the interface to be provided by the transport protocol
// specific client implementations. This mainly deals with the actual sending
// and receiving of messages.
type VersionedClient interface {
	// NewStream returns a new grpc.ClientStream specific to the underlying
	// transport protocol version.
	NewStream(ctx context.Context) (grpc.ClientStream, error)

	// params: resources, typeURL, version, nonce

	// SendRequest constructs and sends out a DiscoveryRequest message specific
	// to the underlying transport protocol version.
	SendRequest(s grpc.ClientStream, resourceNames []string, typeURL string, version string, nonce string) error

	// RecvResponse uses the provided stream to receive a response specific to
	// the underlying transport protocol version.
	RecvResponse(stream grpc.ClientStream) (proto.Message, error)

	// HandleResponse parses and validates the received response and notifies
	// the top-level client which in turn notifies the registered watchers.
	//
	// Return values are: typeURL, version, nonce, error.
	// If the provided protobuf message contains a resource type which is not
	// supported, implementations must return an error of type
	// ErrResourceTypeUnsupported.
	HandleResponse(proto.Message) (string, string, string, error)
}

// TransportHelper contains all xDS transport protocol related functionality
// which is common across different versioned client implementations.
//
// TransportHelper takes care of sending and receiving xDS requests and
// responses on an ADS stream. It also takes care of ACK/NACK handling. It
// delegates to the actual versioned client implementations wherever
// appropriate.
//
// Implements the APIClient interface which makes it possible for versioned
// client implementations to embed this type, and thereby satisfy the interface
// requirements.
type TransportHelper struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	vClient  VersionedClient
	logger   *grpclog.PrefixLogger
	backoff  func(int) time.Duration
	streamCh chan grpc.ClientStream
	sendCh   *buffer.Unbounded

	mu sync.Mutex
	// Message specific watch infos, protected by the above mutex. These are
	// written to, after successfully reading from the update channel, and are
	// read from when recovering from a broken stream to resend the xDS
	// messages. When the user of this client object cancels a watch call,
	// these are set to nil. All accesses to the map protected and any value
	// inside the map should be protected with the above mutex.
	watchMap map[string]map[string]bool
	// versionMap contains the version that was acked (the version in the ack
	// request that was sent on wire). The key is typeURL, the value is the
	// version string, becaues the versions for different resource types should
	// be independent.
	versionMap map[string]string
	// nonceMap contains the nonce from the most recent received response.
	nonceMap map[string]string
}

// NewTransportHelper creates a new transport helper to be used by versioned
// client implementations.
func NewTransportHelper(vc VersionedClient, logger *grpclog.PrefixLogger, backoff func(int) time.Duration) *TransportHelper {
	ctx, cancelCtx := context.WithCancel(context.Background())
	t := &TransportHelper{
		ctx:       ctx,
		cancelCtx: cancelCtx,
		vClient:   vc,
		logger:    logger,
		backoff:   backoff,

		streamCh:   make(chan grpc.ClientStream, 1),
		sendCh:     buffer.NewUnbounded(),
		watchMap:   make(map[string]map[string]bool),
		versionMap: make(map[string]string),
		nonceMap:   make(map[string]string),
	}

	go t.run()
	return t
}

// AddWatch adds a watch for an xDS resource given its type and name.
func (t *TransportHelper) AddWatch(resourceType, resourceName string) {
	t.sendCh.Put(&watchAction{
		typeURL:  resourceType,
		remove:   false,
		resource: resourceName,
	})
}

// RemoveWatch cancels an already registered watch for an xDS resource
// given its type and name.
func (t *TransportHelper) RemoveWatch(resourceType, resourceName string) {
	t.sendCh.Put(&watchAction{
		typeURL:  resourceType,
		remove:   true,
		resource: resourceName,
	})
}

// Close closes the transport helper.
func (t *TransportHelper) Close() {
	t.cancelCtx()
}

// run starts an ADS stream (and backs off exponentially, if the previous
// stream failed without receiving a single reply) and runs the sender and
// receiver routines to send and receive data from the stream respectively.
func (t *TransportHelper) run() {
	go t.send()
	// TODO: start a goroutine monitoring ClientConn's connectivity state, and
	// report error (and log) when stats is transient failure.

	retries := 0
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
		}

		if retries != 0 {
			timer := time.NewTimer(t.backoff(retries))
			select {
			case <-timer.C:
			case <-t.ctx.Done():
				if !timer.Stop() {
					<-timer.C
				}
				return
			}
		}

		retries++
		stream, err := t.vClient.NewStream(t.ctx)
		if err != nil {
			t.logger.Warningf("xds: ADS stream creation failed: %v", err)
			continue
		}
		t.logger.Infof("ADS stream created")

		select {
		case <-t.streamCh:
		default:
		}
		t.streamCh <- stream
		if t.recv(stream) {
			retries = 0
		}
	}
}

// send is a separate goroutine for sending watch requests on the xds stream.
//
// It watches the stream channel for new streams, and the request channel for
// new requests to send on the stream.
//
// For each new request (watchAction), it's
//  - processed and added to the watch map
//    - so resend will pick them up when there are new streams
//  - sent on the current stream if there's one
//    - the current stream is cleared when any send on it fails
//
// For each new stream, all the existing requests will be resent.
//
// Note that this goroutine doesn't do anything to the old stream when there's a
// new one. In fact, there should be only one stream in progress, and new one
// should only be created when the old one fails (recv returns an error).
func (t *TransportHelper) send() {
	var stream grpc.ClientStream
	for {
		select {
		case <-t.ctx.Done():
			return
		case stream = <-t.streamCh:
			if !t.sendExisting(stream) {
				// send failed, clear the current stream.
				stream = nil
			}
		case u := <-t.sendCh.Get():
			t.sendCh.Load()

			var (
				target                  []string
				typeURL, version, nonce string
				send                    bool
			)
			switch update := u.(type) {
			case *watchAction:
				target, typeURL, version, nonce = t.processWatchInfo(update)
			case *ackAction:
				target, typeURL, version, nonce, send = t.processAckInfo(update, stream)
				if !send {
					continue
				}
			}
			if stream == nil {
				// There's no stream yet. Skip the request. This request
				// will be resent to the new streams. If no stream is
				// created, the watcher will timeout (same as server not
				// sending response back).
				continue
			}
			if err := t.vClient.SendRequest(stream, target, typeURL, version, nonce); err != nil {
				t.logger.Errorf("ADS request failed: %v", err)
				// send failed, clear the current stream.
				stream = nil
			}
		}
	}
}

// sendExisting sends out xDS requests for registered watchers when recovering
// from a broken stream.
//
// We call stream.Send() here with the lock being held. It should be OK to do
// that here because the stream has just started and Send() usually returns
// quickly (once it pushes the message onto the transport layer) and is only
// ever blocked if we don't have enough flow control quota.
func (t *TransportHelper) sendExisting(stream grpc.ClientStream) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Reset the ack versions when the stream restarts.
	t.versionMap = make(map[string]string)
	t.nonceMap = make(map[string]string)

	for typeURL, s := range t.watchMap {
		if err := t.vClient.SendRequest(stream, mapToSlice(s), typeURL, "", ""); err != nil {
			t.logger.Errorf("ADS request failed: %v", err)
			return false
		}
	}

	return true
}

// recv receives xDS responses on the provided ADS stream and branches out to
// message specific handlers.
func (t *TransportHelper) recv(stream grpc.ClientStream) bool {
	success := false
	for {
		resp, err := t.vClient.RecvResponse(stream)
		if err != nil {
			t.logger.Warningf("ADS stream is closed with error: %v", err)
			return success
		}
		typeURL, version, nonce, err := t.vClient.HandleResponse(resp)
		if e, ok := err.(ErrResourceTypeUnsupported); ok {
			t.logger.Warningf("%s", e.ErrStr)
			continue
		}
		if err != nil {
			t.sendCh.Put(&ackAction{
				typeURL: typeURL,
				version: "",
				nonce:   nonce,
				stream:  stream,
			})
			t.logger.Warningf("Sending NACK for response type: %v, version: %v, nonce: %v, reason: %v", typeURL, version, nonce, err)
			continue
		}
		t.sendCh.Put(&ackAction{
			typeURL: typeURL,
			version: version,
			nonce:   nonce,
			stream:  stream,
		})
		t.logger.Infof("Sending ACK for response type: %v, version: %v, nonce: %v", typeURL, version, nonce)
		success = true
	}
}

func mapToSlice(m map[string]bool) (ret []string) {
	for i := range m {
		ret = append(ret, i)
	}
	return
}

type watchAction struct {
	typeURL  string
	remove   bool // Whether this is to remove watch for the resource.
	resource string
}

// processWatchInfo pulls the fields needed by the request from a watchAction.
//
// It also updates the watch map in v2c.
func (t *TransportHelper) processWatchInfo(w *watchAction) (target []string, typeURL, ver, nonce string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	var current map[string]bool
	current, ok := t.watchMap[w.typeURL]
	if !ok {
		current = make(map[string]bool)
		t.watchMap[w.typeURL] = current
	}

	if w.remove {
		delete(current, w.resource)
		if len(current) == 0 {
			delete(t.watchMap, w.typeURL)
		}
	} else {
		current[w.resource] = true
	}

	typeURL = w.typeURL
	target = mapToSlice(current)
	// We don't reset version or nonce when a new watch is started. The version
	// and nonce from previous response are carried by the request unless the
	// stream is recreated.
	ver = t.versionMap[typeURL]
	nonce = t.nonceMap[typeURL]
	return target, typeURL, ver, nonce
}

type ackAction struct {
	typeURL string
	version string // NACK if version is an empty string.
	nonce   string
	// ACK/NACK are tagged with the stream it's for. When the stream is down,
	// all the ACK/NACK for this stream will be dropped, and the version/nonce
	// won't be updated.
	stream grpc.ClientStream
}

// processAckInfo pulls the fields needed by the ack request from a ackAction.
//
// If no active watch is found for this ack, it returns false for send.
func (t *TransportHelper) processAckInfo(ack *ackAction, stream grpc.ClientStream) (target []string, typeURL, version, nonce string, send bool) {
	if ack.stream != stream {
		// If ACK's stream isn't the current sending stream, this means the ACK
		// was pushed to queue before the old stream broke, and a new stream has
		// been started since. Return immediately here so we don't update the
		// nonce for the new stream.
		return nil, "", "", "", false
	}
	typeURL = ack.typeURL

	t.mu.Lock()
	defer t.mu.Unlock()

	// Update the nonce no matter if we are going to send the ACK request on
	// wire. We may not send the request if the watch is canceled. But the nonce
	// needs to be updated so the next request will have the right nonce.
	nonce = ack.nonce
	t.nonceMap[typeURL] = nonce

	s, ok := t.watchMap[typeURL]
	if !ok || len(s) == 0 {
		// We don't send the request ack if there's no active watch (this can be
		// either the server sends responses before any request, or the watch is
		// canceled while the ackAction is in queue), because there's no resource
		// name. And if we send a request with empty resource name list, the
		// server may treat it as a wild card and send us everything.
		return nil, "", "", "", false
	}
	send = true
	target = mapToSlice(s)

	version = ack.version
	if version == "" {
		// This is a nack, get the previous acked version.
		version = t.versionMap[typeURL]
		// version will still be an empty string if typeURL isn't
		// found in versionMap, this can happen if there wasn't any ack
		// before.
	} else {
		t.versionMap[typeURL] = version
	}
	return target, typeURL, version, nonce, send
}
