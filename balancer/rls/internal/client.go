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

package rls

import (
	"context"
	"time"

	"google.golang.org/grpc"
	rlspb "google.golang.org/grpc/balancer/rls/internal/proto/grpc_lookup_v1"
)

// For gRPC services using RLS, the value of target_type in the
// RouteLookupServiceRequest will be set to this.
const grpcTargetType = "grpc"

// rlsClient is a simple wrapper around a RouteLookupService client which
// provides non-blocking semantics on top of a blocking unary RPC call.
//
// The RLS LB policy creates a new rlsClient object with the following values:
// * a grpc.ClientConn to the RLS server using appropriate credentials from the
//   parent channel
// * dialTarget corresponding to the original user dial target, e.g.
//   "firestore.googleapis.com".
//
// The RLS LB policy uses an adaptive throttler to perform client side
// throttling and asks this client to make an RPC call only after checking with
// the throttler.
type rlsClient struct {
	cc         *grpc.ClientConn
	stub       rlspb.RouteLookupServiceClient
	dialTarget string
	rpcTimeout time.Duration
}

func newRLSClient(cc *grpc.ClientConn, dialTarget string, rpcTimeout time.Duration) *rlsClient {
	return &rlsClient{
		cc:         cc,
		stub:       rlspb.NewRouteLookupServiceClient(cc),
		dialTarget: dialTarget,
		rpcTimeout: rpcTimeout,
	}
}

type lookupCallback func(target, headerData string, err error)

// lookup starts a RouteLookup RPC in a separate goroutine and returns the
// results (and error, if any) in the provided callback.
func (c *rlsClient) lookup(path string, keyMap map[string]string, cb lookupCallback) {
	ctx, cancel := context.WithTimeout(context.Background(), c.rpcTimeout)
	go func() {
		resp, err := c.stub.RouteLookup(ctx, &rlspb.RouteLookupRequest{
			Server:     c.dialTarget,
			Path:       path,
			TargetType: grpcTargetType,
			KeyMap:     keyMap,
		})
		cb(resp.GetTarget(), resp.GetHeaderData(), err)
		cancel()
	}()
}
