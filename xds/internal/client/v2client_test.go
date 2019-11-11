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
	"errors"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/xds/internal/client/fakexds"

	discoverypb "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	ldspb "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	rdspb "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	xdspb "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	basepb "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	routepb "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	httppb "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	listenerpb "github.com/envoyproxy/go-control-plane/envoy/config/listener/v2"
	anypb "github.com/golang/protobuf/ptypes/any"
	structpb "github.com/golang/protobuf/ptypes/struct"
)

const (
	defaultTestTimeout       = 2 * time.Second
	goodLDSTarget1           = "lds.target.good:1111"
	goodLDSTarget2           = "lds.target.good:2222"
	uninterestingLDSTarget   = "lds.target.uninteresting"
	goodRouteName1           = "GoodRouteConfig1"
	goodRouteName2           = "GoodRouteConfig2"
	uninterestingRouteName   = "UninterestingRouteName"
	goodMatchingDomain       = "lds.target.good"
	uninterestingDomain      = "uninteresting.domain"
	goodClusterName1         = "GoodClusterName1"
	goodClusterName2         = "GoodClusterName2"
	uninterestingClusterName = "UninterestingClusterName"
	httpConnManagerURL       = "type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager"
)

var (
	goodNodeProto = &basepb.Node{
		Id: "ENVOY_NODE_ID",
		Metadata: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"TRAFFICDIRECTOR_GRPC_HOSTNAME": {
					Kind: &structpb.Value_StringValue{StringValue: "trafficdirector"},
				},
			},
		},
	}
	goodLDSRequest = &discoverypb.DiscoveryRequest{
		Node:          goodNodeProto,
		TypeUrl:       listenerURL,
		ResourceNames: []string{goodLDSTarget1},
	}
	goodHTTPConnManager1 = &httppb.HttpConnectionManager{
		RouteSpecifier: &httppb.HttpConnectionManager_Rds{
			Rds: &httppb.Rds{
				RouteConfigName: goodRouteName1,
			},
		},
	}
	marshaledConnMgr1, _ = proto.Marshal(goodHTTPConnManager1)
	goodHTTPConnManager2 = &httppb.HttpConnectionManager{
		RouteSpecifier: &httppb.HttpConnectionManager_Rds{
			Rds: &httppb.Rds{
				RouteConfigName: goodRouteName2,
			},
		},
	}
	marshaledConnMgr2, _ = proto.Marshal(goodHTTPConnManager2)
	emptyHTTPConnManager = &httppb.HttpConnectionManager{
		RouteSpecifier: &httppb.HttpConnectionManager_Rds{
			Rds: &httppb.Rds{},
		},
	}
	emptyMarshaledConnMgr, _     = proto.Marshal(emptyHTTPConnManager)
	connMgrWithInlineRouteConfig = &httppb.HttpConnectionManager{
		RouteSpecifier: &httppb.HttpConnectionManager_RouteConfig{
			RouteConfig: &rdspb.RouteConfiguration{
				Name: goodRouteName1,
			},
		},
	}
	marshaledConnMgrWithInlineRouteConfig, _ = proto.Marshal(connMgrWithInlineRouteConfig)
	connMgrWithScopedRoutes                  = &httppb.HttpConnectionManager{
		RouteSpecifier: &httppb.HttpConnectionManager_ScopedRoutes{},
	}
	marshaledConnMgrWithScopedRoutes, _ = proto.Marshal(connMgrWithScopedRoutes)
	goodListener1                       = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgr1,
			},
		},
	}
	marshaledListener1, _ = proto.Marshal(goodListener1)
	goodListener2         = &ldspb.Listener{
		Name: goodLDSTarget2,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgr1,
			},
		},
	}
	marshaledListener2, _ = proto.Marshal(goodListener2)
	otherGoodListener2    = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgr2,
			},
		},
	}
	otherMarshaledListener2, _ = proto.Marshal(otherGoodListener2)
	uninterestingListener      = &ldspb.Listener{
		Name: uninterestingLDSTarget,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgr1,
			},
		},
	}
	uninterestingMarshaledListener, _ = proto.Marshal(uninterestingListener)
	noAPIListener                     = &ldspb.Listener{Name: goodLDSTarget1}
	marshaledNoAPIListener, _         = proto.Marshal(noAPIListener)
	badAPIListener1                   = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   []byte{1, 2, 3, 4},
			},
		},
	}
	badlyMarshaledAPIListener1, _ = proto.Marshal(badAPIListener1)
	badAPIListener2               = &ldspb.Listener{
		Name: goodLDSTarget2,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   []byte{1, 2, 3, 4},
			},
		},
	}
	badlyMarshaledAPIListener2, _ = proto.Marshal(badAPIListener2)
	badResourceListener           = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: listenerURL,
				Value:   marshaledListener1,
			},
		},
	}
	marshaledBadResourceListener, _ = proto.Marshal(badResourceListener)
	listenerWithEmptyHTTPConnMgr    = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   emptyMarshaledConnMgr,
			},
		},
	}
	marshaledListenerWithEmptyHTTPConnMgr, _ = proto.Marshal(listenerWithEmptyHTTPConnMgr)
	listenerWithInlineRouteConfig            = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgrWithInlineRouteConfig,
			},
		},
	}
	marshaledListenerWithInlineRouteConfig, _ = proto.Marshal(listenerWithInlineRouteConfig)
	listenerWithScopedRoutesRouteConfig       = &ldspb.Listener{
		Name: goodLDSTarget1,
		ApiListener: &listenerpb.ApiListener{
			ApiListener: &anypb.Any{
				TypeUrl: httpConnManagerURL,
				Value:   marshaledConnMgrWithScopedRoutes,
			},
		},
	}
	goodLDSResponse1 = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener1,
			},
		},
		TypeUrl: listenerURL,
	}
	goodLDSResponse2 = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener2,
			},
		},
		TypeUrl: listenerURL,
	}
	emptyLDSResponse          = &discoverypb.DiscoveryResponse{TypeUrl: listenerURL}
	badlyMarshaledLDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   []byte{1, 2, 3, 4},
			},
		},
		TypeUrl: listenerURL,
	}
	badResourceTypeInLDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledConnMgr1,
			},
		},
		TypeUrl: listenerURL,
	}
	ldsResponseWithMultipleResources = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener2,
			},
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener1,
			},
		},
		TypeUrl: listenerURL,
	}
	noAPIListenerLDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledNoAPIListener,
			},
		},
		TypeUrl: listenerURL,
	}
	goodBadUglyLDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener2,
			},
			{
				TypeUrl: listenerURL,
				Value:   marshaledListener1,
			},
			{
				TypeUrl: listenerURL,
				Value:   badlyMarshaledAPIListener2,
			},
		},
		TypeUrl: listenerURL,
	}
	badlyMarshaledRDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   []byte{1, 2, 3, 4},
			},
		},
		TypeUrl: routeURL,
	}
	badResourceTypeInRDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   marshaledConnMgr1,
			},
		},
		TypeUrl: routeURL,
	}
	emptyRouteConfig             = &xdspb.RouteConfiguration{}
	marshaledEmptyRouteConfig, _ = proto.Marshal(emptyRouteConfig)
	noDomainsInRouteConfig       = &xdspb.RouteConfiguration{
		VirtualHosts: []*routepb.VirtualHost{{}},
	}
	noVirtualHostsInRDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   marshaledEmptyRouteConfig,
			},
		},
		TypeUrl: routeURL,
	}
	goodRouteConfig1 = &xdspb.RouteConfiguration{
		Name: goodRouteName1,
		VirtualHosts: []*routepb.VirtualHost{
			{
				Domains: []string{uninterestingDomain},
				Routes: []*routepb.Route{
					{
						Action: &routepb.Route_Route{
							Route: &routepb.RouteAction{
								ClusterSpecifier: &routepb.RouteAction_Cluster{Cluster: uninterestingClusterName},
							},
						},
					},
				},
			},
			{
				Domains: []string{goodMatchingDomain},
				Routes: []*routepb.Route{
					{
						Action: &routepb.Route_Route{
							Route: &routepb.RouteAction{
								ClusterSpecifier: &routepb.RouteAction_Cluster{Cluster: goodClusterName1},
							},
						},
					},
				},
			},
		},
	}
	marshaledGoodRouteConfig1, _ = proto.Marshal(goodRouteConfig1)
	goodRouteConfig2             = &xdspb.RouteConfiguration{
		Name: goodRouteName2,
		VirtualHosts: []*routepb.VirtualHost{
			{
				Domains: []string{uninterestingDomain},
				Routes: []*routepb.Route{
					{
						Action: &routepb.Route_Route{
							Route: &routepb.RouteAction{
								ClusterSpecifier: &routepb.RouteAction_Cluster{Cluster: uninterestingClusterName},
							},
						},
					},
				},
			},
			{
				Domains: []string{goodMatchingDomain},
				Routes: []*routepb.Route{
					{
						Action: &routepb.Route_Route{
							Route: &routepb.RouteAction{
								ClusterSpecifier: &routepb.RouteAction_Cluster{Cluster: goodClusterName2},
							},
						},
					},
				},
			},
		},
	}
	marshaledGoodRouteConfig2, _ = proto.Marshal(goodRouteConfig2)
	uninterestingRouteConfig     = &xdspb.RouteConfiguration{
		Name: uninterestingRouteName,
		VirtualHosts: []*routepb.VirtualHost{
			{
				Domains: []string{uninterestingDomain},
				Routes: []*routepb.Route{
					{
						Action: &routepb.Route_Route{
							Route: &routepb.RouteAction{
								ClusterSpecifier: &routepb.RouteAction_Cluster{Cluster: uninterestingClusterName},
							},
						},
					},
				},
			},
		},
	}
	marshaledUninterestingRouteConfig, _ = proto.Marshal(uninterestingRouteConfig)
	goodRDSResponse1                     = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   marshaledGoodRouteConfig1,
			},
		},
		TypeUrl: routeURL,
	}
	goodRDSResponse2 = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   marshaledGoodRouteConfig2,
			},
		},
		TypeUrl: routeURL,
	}
	uninterestingRDSResponse = &discoverypb.DiscoveryResponse{
		Resources: []*anypb.Any{
			{
				TypeUrl: routeURL,
				Value:   marshaledUninterestingRouteConfig,
			},
		},
		TypeUrl: routeURL,
	}
)

// testOp contains all data related to one particular test operation. Not all
// fields make sense for all tests.
type testOp struct {
	// target is the resource name to watch for.
	target string
	// responseToSend is the xDS response sent to the client
	responseToSend *fakexds.Response
	// wantOpData is the operation specific output that we expect.
	wantOpData interface{}
	// wantOpErr specfies whether the main operation should return an error.
	wantOpErr bool
	// wantRetry specifies whether or not the client is expected to kill the
	// stream because of an error, and expected to backoff and retry.
	wantRetry bool
	// wantRequest is the LDS request expected to be sent by the client.
	wantRequest *fakexds.Request
	// wantRDSCache is the expected rdsCache at the end of an operation.
	wantRDSCache map[string]string
	// wantWatchCallback specifies if the watch callback should be invoked.
	wantWatchCallback bool
}

// TestV2ClientBackoffAfterRecvError verifies if the v2Client backoffs when it
// encounters a Recv error while receiving an LDS response.
func TestV2ClientBackoffAfterRecvError(t *testing.T) {
	fakeServer, client, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	// Override the v2Client backoff function with this, so that we can verify
	// that a backoff actually was triggerred.
	boCh := make(chan int, 1)
	clientBackoff := func(v int) time.Duration {
		boCh <- v
		return 0
	}

	v2c := newV2Client(client, goodNodeProto, clientBackoff)
	defer v2c.close()
	t.Log("Started xds v2Client...")

	v2c.watchLDS(goodLDSTarget1, func(u ldsUpdate, err error) {
		t.Fatalf("Received unexpected LDS callback with ldsUpdate {%+v} and error {%v}", u, err)
	})
	<-fakeServer.RequestChan
	t.Log("FakeServer received request...")

	fakeServer.ResponseChan <- &fakexds.Response{Err: errors.New("RPC error")}
	t.Log("Bad LDS response pushed to fakeServer...")

	timer := time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
		t.Fatal("time out when expecting LDS update")
	case <-boCh:
		t.Log("v2Client backed off before retrying...")
	}
}

// TestV2ClientRetriesAfterBrokenStream verifies the case where a stream
// encountered a Recv() error, and is expected to send out xDS requests for
// registered watchers once it comes back up again.
func TestV2ClientRetriesAfterBrokenStream(t *testing.T) {
	fakeServer, client, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	v2c := newV2Client(client, goodNodeProto, func(int) time.Duration { return 0 })
	defer v2c.close()
	t.Log("Started xds v2Client...")

	callbackCh := make(chan struct{}, 1)
	v2c.watchLDS(goodLDSTarget1, func(u ldsUpdate, err error) {
		t.Logf("Received LDS callback with ldsUpdate {%+v} and error {%v}", u, err)
		callbackCh <- struct{}{}
	})
	<-fakeServer.RequestChan
	t.Log("FakeServer received request...")

	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodLDSResponse1}
	t.Log("Good LDS response pushed to fakeServer...")

	timer := time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
		t.Fatal("time out when expecting LDS update")
	case <-callbackCh:
	}

	fakeServer.ResponseChan <- &fakexds.Response{Err: errors.New("RPC error")}
	t.Log("Bad LDS response pushed to fakeServer...")

	timer = time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
		t.Fatal("time out when expecting LDS update")
	case gotRequest := <-fakeServer.RequestChan:
		t.Log("received LDS request after stream re-creation")
		if !proto.Equal(gotRequest.Req, goodLDSRequest) {
			t.Fatalf("gotRequest: %+v, wantRequest: %+v", gotRequest.Req, goodLDSRequest)
		}
	}
}

// TestV2ClientCancelWatch verifies that the registered watch callback is not
// invoked if a response is received after the watcher is cancelled.
func TestV2ClientCancelWatch(t *testing.T) {
	fakeServer, client, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	v2c := newV2Client(client, goodNodeProto, func(int) time.Duration { return 0 })
	defer v2c.close()
	t.Log("Started xds v2Client...")

	callbackCh := make(chan struct{}, 1)
	cancelFunc := v2c.watchLDS(goodLDSTarget1, func(u ldsUpdate, err error) {
		t.Logf("Received LDS callback with ldsUpdate {%+v} and error {%v}", u, err)
		callbackCh <- struct{}{}
	})
	<-fakeServer.RequestChan
	t.Log("FakeServer received request...")

	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodLDSResponse1}
	t.Log("Good LDS response pushed to fakeServer...")

	timer := time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
		t.Fatal("time out when expecting LDS update")
	case <-callbackCh:
	}

	cancelFunc()

	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodLDSResponse1}
	t.Log("Another good LDS response pushed to fakeServer...")

	timer = time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
	case <-callbackCh:
		t.Fatalf("Watch callback invoked after the watcher was cancelled")
	}
}
