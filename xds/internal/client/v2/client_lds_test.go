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

package v2

import (
	"testing"
	"time"

	v2xdspb "github.com/envoyproxy/go-control-plane/envoy/api/v2"

	xdsclient "google.golang.org/grpc/xds/internal/client"
)

// TestLDSHandleResponse starts a fake xDS server, makes a ClientConn to it,
// and creates a client using it. Then, it registers a watchLDS and tests
// different LDS responses.
func (s) TestLDSHandleResponse(t *testing.T) {
	tests := []struct {
		name          string
		ldsResponse   *v2xdspb.DiscoveryResponse
		wantErr       bool
		wantUpdate    *xdsclient.ListenerUpdate
		wantUpdateErr bool
	}{
		// No APIListener in the response. Just one test case here for a bad
		// ApiListener, since the others are covered in
		// TestGetRouteConfigNameFromListener.
		// Response contains one listener and it is good.
		// Response contains multiple good listeners, including the one we are
		// interested in.
		{
			name:          "multiple-good-listener",
			ldsResponse:   ldsResponseWithMultipleResources,
			wantErr:       false,
			wantUpdate:    &xdsclient.ListenerUpdate{RouteConfigName: goodRouteName1},
			wantUpdateErr: false,
		},
		// Response contains two good listeners (one interesting and one
		// uninteresting), and one badly marshaled listener. This will cause a
		// nack because the uninteresting listener will still be parsed.
		{
			name:          "good-bad-ugly-listeners",
			ldsResponse:   goodBadUglyLDSResponse,
			wantErr:       true,
			wantUpdate:    nil,
			wantUpdateErr: false,
		},
		// Response contains one listener, but we are not interested in it.
		{
			name:          "one-uninteresting-listener",
			ldsResponse:   goodLDSResponse2,
			wantErr:       false,
			wantUpdate:    nil,
			wantUpdateErr: false,
		},
		// Response constains no resources. This is the case where the server
		// does not know about the target we are interested in.
		{
			name:          "empty-response",
			ldsResponse:   emptyLDSResponse,
			wantErr:       false,
			wantUpdate:    nil,
			wantUpdateErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testWatchHandle(t, &watchHandleTestcase{
				rType:            xdsclient.ListenerResource,
				resourceName:     goodLDSTarget1,
				responseToHandle: test.ldsResponse,
				wantHandleErr:    test.wantErr,
				wantUpdate:       test.wantUpdate,
				wantUpdateErr:    test.wantUpdateErr,
			})
		})
	}
}

// TestLDSHandleResponseWithoutWatch tests the case where the client receives
// an LDS response without a registered watcher.
func (s) TestLDSHandleResponseWithoutWatch(t *testing.T) {
	_, cc, cleanup := startServerAndGetCC(t)
	defer cleanup()

	v2c, err := newV2Client(&testUpdateReceiver{
		f: func(xdsclient.ResourceType, map[string]interface{}) {},
	}, cc, goodNodeProto, func(int) time.Duration { return 0 }, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer v2c.Close()

	if v2c.handleLDSResponse(badResourceTypeInLDSResponse) == nil {
		t.Fatal("v2c.handleLDSResponse() succeeded, should have failed")
	}

	if v2c.handleLDSResponse(goodLDSResponse1) != nil {
		t.Fatal("v2c.handleLDSResponse() succeeded, should have failed")
	}
}
