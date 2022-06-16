/*
 *
 * Copyright 2022 gRPC authors.
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

package xds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/testutils"
	"google.golang.org/grpc/internal/testutils/xds/e2e"

	v3discoverypb "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	testgrpc "google.golang.org/grpc/test/grpc_testing"
	testpb "google.golang.org/grpc/test/grpc_testing"
)

// TestClientResourceVersionAfterStreamRestart tests the scenario where the
// xdsClient's ADS stream to the management server gets broken. This test
// verifies that the version number on the initial request on the new stream
// indicates the most recent version seen by the client on the previous stream.
func (s) TestClientResourceVersionAfterStreamRestart(t *testing.T) {
	// Create a restartable listener which can close existing connections.
	l, err := testutils.LocalTCPListener()
	if err != nil {
		t.Fatalf("testutils.LocalTCPListener() failed: %v", err)
	}
	lis := testutils.NewRestartableListener(l)

	// Maps to store ACK versions before and after stream restart.
	ackVersionsBeforeRestart := make(map[string]string)
	ackVersionsAfterRestart := make(map[string]string)
	// Channels to notify that all expected resources have been requested by the
	// xdsClient before and after stream restart.
	resourcesRequestedBeforeStreamClose := make(chan struct{}, 1)
	resourcesRequestedAfterStreamClose := make(chan struct{}, 1)
	// Event to notify stream closure.
	streamClosed := grpcsync.NewEvent()
	const wantResources = 4

	managementServer, nodeID, _, resolver, cleanup1 := e2e.SetupManagementServer(t, &e2e.ManagementServerOptions{
		Listener: lis,
		OnStreamRequest: func(id int64, request *v3discoverypb.DiscoveryRequest) error {
			// Populate the versions in the appropriate map based on whether the
			// stream has closed.
			if !streamClosed.HasFired() {
				// Prior to stream closure, record only non-empty version numbers. The
				// client first requests for a resource with version set to empty
				// string. After receipt of the response, it sends another request for
				// the same resource, this time with a non-empty version string. This
				// corresponds to ACKs, and this is what we want to capture.
				if len(request.GetResourceNames()) != 0 && request.GetVersionInfo() != "" {
					ackVersionsBeforeRestart[request.GetTypeUrl()] = request.GetVersionInfo()
					if len(ackVersionsBeforeRestart) == wantResources {
						select {
						case resourcesRequestedBeforeStreamClose <- struct{}{}:
						default:
						}
					}
				}
				return nil
			}
			// After stream closure, capture the first request for every resource.
			// This should not be set to an empty version string, but instead should
			// be set to the version last ACKed before stream closure.
			if len(request.GetResourceNames()) != 0 {
				if ackVersionsAfterRestart[request.GetTypeUrl()] == "" {
					ackVersionsAfterRestart[request.GetTypeUrl()] = request.GetVersionInfo()
				}
				if len(ackVersionsAfterRestart) == wantResources {
					select {
					case resourcesRequestedAfterStreamClose <- struct{}{}:
					default:
					}
				}
			}
			return nil
		},
		OnStreamClosed: func(int64) {
			streamClosed.Fire()
		},
	})
	defer cleanup1()

	port, cleanup2 := startTestService(t, nil)
	defer cleanup2()

	const serviceName = "my-service-client-side-xds"
	resources := e2e.DefaultClientResources(e2e.ResourceParams{
		DialTarget: serviceName,
		NodeID:     nodeID,
		Host:       "localhost",
		Port:       port,
		SecLevel:   e2e.SecurityLevelNone,
	})
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()
	if err := managementServer.Update(ctx, resources); err != nil {
		t.Fatal(err)
	}

	// Create a ClientConn and make a successful RPC.
	cc, err := grpc.Dial(fmt.Sprintf("xds:///%s", serviceName), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithResolvers(resolver))
	if err != nil {
		t.Fatalf("failed to dial local test server: %v", err)
	}
	defer cc.Close()

	client := testgrpc.NewTestServiceClient(cc)
	if _, err := client.EmptyCall(ctx, &testpb.Empty{}); err != nil {
		t.Fatalf("rpc EmptyCall() failed: %v", err)
	}

	// Wait for all the resources to be ACKed.
	select {
	case <-resourcesRequestedBeforeStreamClose:
	case <-ctx.Done():
		t.Fatal("timeout waiting for resource to be ACKed before stream restart")
	}

	// Stop the listener on the management server. This will cause the client to
	// backoff and recreate the stream.
	lis.Stop()

	// Wait for the stream to be closed on the server.
	<-streamClosed.Done()

	// Restart the listener on the management server to be able to accept
	// reconnect attempts from the client.
	lis.Restart()

	// Wait for all the previously sent resources to be re-requested.
	select {
	case <-resourcesRequestedAfterStreamClose:
	case <-ctx.Done():
		t.Fatal("timeout when waiting for all resources to be re-requested after stream restart")
	}

	if !cmp.Equal(ackVersionsBeforeRestart, ackVersionsAfterRestart) {
		t.Fatalf("ackVersionsBeforeRestart: %v and ackVersionsAfterRestart: %v don't match", ackVersionsBeforeRestart, ackVersionsAfterRestart)
	}
	if _, err := client.EmptyCall(ctx, &testpb.Empty{}, grpc.WaitForReady(true)); err != nil {
		t.Fatalf("rpc EmptyCall() failed: %v", err)
	}
}
