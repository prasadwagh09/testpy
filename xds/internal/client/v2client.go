/*
 *
 * Copyright 2019 gRPC authors.
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
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/internal/buffer"

	corepb "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	adsgrpc "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
)

// The value chosen here is based on the default value of the
// initial_fetch_timeout field in corepb.ConfigSource proto.
const defaultWatchTimer = 15 * time.Second

// v2Client performs the actual xDS RPCs using the xDS v2 API. It creates a
// single ADS stream on which the different types of xDS requests and responses
// are multiplexed.
// The reason for splitting this out from the top level xdsClient object is
// because there is already an xDS v3Aplha API in development. If and when we
// want to switch to that, this seperation will ease that process.
type v2Client struct {
	ctx       context.Context
	cancelCtx context.CancelFunc

	// ClientConn to the xDS gRPC server. Owned by the parent xdsClient.
	cc        *grpc.ClientConn
	nodeProto *corepb.Node
	backoff   func(int) time.Duration

	// watchCh in the channel onto which watchInfo objects are pushed by the
	// watch API, and it is read and acted upon by the send() goroutine.
	watchCh *buffer.Unbounded

	mu sync.Mutex
	// Message specific watch infos, protected by the above mutex. These are
	// written to, after successfully reading from the update channel, and are
	// read from when recovering from a broken stream to resend the xDS
	// messages. When the user of this client object cancels a watch call,
	// these are set to nil. All accesses to the map protected and any value
	// inside the map should be protected with the above mutex.
	watchMap map[resourceType]*watchInfo
	// rdsCache maintains a mapping of {routeConfigName --> clusterName} from
	// validated route configurations received in RDS responses. We cache all
	// valid route configurations, whether or not we are interested in them
	// when we received them (because we could become interested in them in the
	// future and the server wont send us those resources again).
	// Protected by the above mutex.
	rdsCache map[string]string
}

// newV2Client creates a new v2Client initialized with the passed arguments.
func newV2Client(cc *grpc.ClientConn, nodeProto *corepb.Node, backoff func(int) time.Duration) *v2Client {
	v2c := &v2Client{
		cc:        cc,
		nodeProto: nodeProto,
		backoff:   backoff,
		watchCh:   buffer.NewUnbounded(),
		watchMap:  make(map[resourceType]*watchInfo),
		rdsCache:  make(map[string]string),
	}
	v2c.ctx, v2c.cancelCtx = context.WithCancel(context.Background())

	go v2c.run()
	return v2c
}

// close cleans up resources and goroutines allocated by this client.
func (v2c *v2Client) close() {
	v2c.cancelCtx()
}

// run starts an ADS stream (and backs off exponentially, if the previous
// stream failed without receiving a single reply) and runs the sender and
// receiver routines to send and receive data from the stream respectively.
func (v2c *v2Client) run() {
	retries := 0
	for {
		select {
		case <-v2c.ctx.Done():
			return
		default:
		}

		if retries != 0 {
			t := time.NewTimer(v2c.backoff(retries))
			select {
			case <-t.C:
			case <-v2c.ctx.Done():
				if !t.Stop() {
					<-t.C
				}
				return
			}
		}

		retries++
		cli := adsgrpc.NewAggregatedDiscoveryServiceClient(v2c.cc)
		stream, err := cli.StreamAggregatedResources(v2c.ctx, grpc.WaitForReady(true))
		if err != nil {
			grpclog.Infof("xds: ADS stream creation failed: %v", err)
			continue
		}

		// send() could be blocked on reading updates from the different update
		// channels when it is not actually sending out messages. So, we need a
		// way to break out of send() when recv() returns. This done channel is
		// used to for that purpose.
		done := make(chan struct{})
		go v2c.send(stream, done)
		if v2c.recv(stream) {
			retries = 0
		}
		close(done)
	}
}

// sendExisting sends out xDS requests for registered watchers when recovering
// from a broken stream.
//
// We call stream.Send() here with the lock being held. It should be OK to do
// that here because the stream has just started and Send() usually returns
// quickly (once it pushes the message onto the transport layer) and is only
// ever blocked if we don't have enough flow control quota.
func (v2c *v2Client) sendExisting(stream adsStream) bool {
	v2c.mu.Lock()
	defer v2c.mu.Unlock()

	for wType, wi := range v2c.watchMap {
		switch wType {
		case ldsResource:
			if !v2c.sendLDS(stream, wi.target) {
				return false
			}
		case rdsResource:
			if !v2c.sendRDS(stream, wi.target) {
				return false
			}
		}
	}

	return true
}

// send reads from message specific update channels and sends out actual xDS
// requests on the provided ADS stream.
func (v2c *v2Client) send(stream adsStream, done chan struct{}) {
	if !v2c.sendExisting(stream) {
		return
	}

	for {
		select {
		case <-v2c.ctx.Done():
			return
		case u := <-v2c.watchCh.Get():
			v2c.watchCh.Load()
			wi := u.(*watchInfo)
			v2c.mu.Lock()
			if wi.state == watchCancelled {
				v2c.mu.Unlock()
				continue
			}
			wi.state = watchStarted
			target := wi.target
			v2c.checkWatchTargetInCache(wi)
			v2c.updateWatchMap(wi)
			v2c.mu.Unlock()

			switch wi.wType {
			case ldsResource:
				if !v2c.sendLDS(stream, target) {
					return
				}
			case rdsResource:
				if !v2c.sendRDS(stream, target) {
					return
				}
			}
		case <-done:
			return
		}
	}
}

// recv receives xDS responses on the provided ADS stream and branches out to
// message specific handlers.
func (v2c *v2Client) recv(stream adsStream) bool {
	success := false
	for {
		resp, err := stream.Recv()
		if err != nil {
			grpclog.Warningf("xds: ADS stream recv failed: %v", err)
			return success
		}
		switch resp.GetTypeUrl() {
		case listenerURL:
			if err := v2c.handleLDSResponse(resp); err != nil {
				grpclog.Warningf("xds: LDS response handler failed: %v", err)
				return success
			}
		case routeURL:
			if err := v2c.handleRDSResponse(resp); err != nil {
				grpclog.Warningf("xds: RDS response handler failed: %v", err)
				return success
			}
		default:
			grpclog.Warningf("xds: unknown response URL type: %v", resp.GetTypeUrl())
		}
		success = true
	}
}

// watchLDS registers an LDS watcher for the provided target. Updates
// corresponding to received LDS responses will be pushed to the provided
// callback. The caller can cancel the watch by invoking the returned cancel
// function.
// The provided callback should not block or perform any expensive operations
// or call other methods of the v2Client object.
func (v2c *v2Client) watchLDS(target string, ldsCb ldsCallback) (cancel func()) {
	wi := &watchInfo{wType: ldsResource, target: []string{target}, callback: ldsCb}
	v2c.watchCh.Put(wi)
	return func() {
		v2c.mu.Lock()
		defer v2c.mu.Unlock()
		if wi.state == watchEnqueued {
			wi.state = watchCancelled
			return
		}
		v2c.watchMap[ldsResource] = nil
	}
}

// watchRDS registers an RDS watcher for the provided routeName. Updates
// corresponding to received RDS responses will be pushed to the provided
// callback. The caller can cancel the watch by invoking the returned cancel
// function.
// The provided callback should not block or perform any expensive operations
// or call other methods of the v2Client object.
func (v2c *v2Client) watchRDS(routeName string, rdsCb rdsCallback) (cancel func()) {
	wi := &watchInfo{wType: rdsResource, target: []string{routeName}, callback: rdsCb}
	v2c.watchCh.Put(wi)
	return func() {
		v2c.mu.Lock()
		defer v2c.mu.Unlock()
		if wi.state == watchEnqueued {
			wi.state = watchCancelled
			return
		}
		v2c.watchMap[rdsResource].cancel()
		v2c.watchMap[rdsResource] = nil
	}
}

// updateWatchMap takes care of updating watchInfo state in the map in a clean
// way. Mainly it takes care of closing a watch timer if one exists, and sets
// up the timer for the new watcher.
//
// Caller should hold v2c.mu
func (v2c *v2Client) updateWatchMap(wi *watchInfo) {
	if existing := v2c.watchMap[wi.wType]; existing != nil {
		existing.cancel()
	}

	v2c.watchMap[wi.wType] = wi
	switch wi.wType {
	case rdsResource:
		wi.timer = time.AfterFunc(defaultWatchTimer, func() {
			wi.callback.(rdsCallback)(rdsUpdate{cluster: ""}, fmt.Errorf("xds: RDS target %s not found", wi.target))
		})
	}
}

// checkWatchTargetInCache is called when a new watch call is handled in
// send(). This method checks if the resource is found in the cache, and if so,
// returns it to the watcher.
// This is required only for RDS and EDS.
//
// Caller should hold v2c.mu
func (v2c *v2Client) checkWatchTargetInCache(wi *watchInfo) {
	switch wi.wType {
	case rdsResource:
		routeName := wi.target[0]
		if cluster := v2c.rdsCache[routeName]; cluster != "" {
			if v2c.watchMap[ldsResource] == nil {
				grpclog.Warningf("xds: no LDS watcher found when handling RDS watch for route {%v} from cache", routeName)
				return
			}
			wi.timer.Stop()
			wi.callback.(rdsCallback)(rdsUpdate{cluster: cluster}, nil)
		}
	}
}
