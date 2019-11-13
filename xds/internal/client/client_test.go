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
	"errors"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/xds/internal/client/bootstrap"
	"google.golang.org/grpc/xds/internal/client/fakexds"

	corepb "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

const balancerName = "dummyBalancer"

var validConfig = bootstrap.Config{
	BalancerName: balancerName,
	Creds:        grpc.WithInsecure(),
	NodeProto:    &corepb.Node{},
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name            string
		opts            Options
		wantErr         bool
		wantDialOptsLen int
	}{
		{name: "empty-opts", opts: Options{}, wantErr: true},
		{
			name: "empty-balancer-name",
			opts: Options{
				Config: bootstrap.Config{
					Creds:     grpc.WithInsecure(),
					NodeProto: &corepb.Node{},
				},
			},
			wantErr: true,
		},
		{
			name: "empty-dial-creds",
			opts: Options{
				Config: bootstrap.Config{
					BalancerName: "dummy",
					NodeProto:    &corepb.Node{},
				},
			},
			wantErr: true,
		},
		{
			name: "empty-node-proto",
			opts: Options{
				Config: bootstrap.Config{
					BalancerName: "dummy",
					Creds:        grpc.WithInsecure(),
				},
			},
			wantErr: true,
		},
		{
			name:            "without-extra-dialoptions",
			opts:            Options{Config: validConfig},
			wantErr:         false,
			wantDialOptsLen: 1,
		},
		{
			name: "without-extra-dialoptions",
			opts: Options{
				Config:   validConfig,
				DialOpts: []grpc.DialOption{grpc.WithDisableRetry()},
			},
			wantErr:         false,
			wantDialOptsLen: 2,
		},
	}

	for _, test := range tests {
		func() {
			oldDialFunc := dialFunc
			dialFunc = func(ctx context.Context, target string, dopts ...grpc.DialOption) (*grpc.ClientConn, error) {
				if target != balancerName {
					t.Fatalf("%s: in dialFunc() got target: %v, want %v", test.name, target, balancerName)
				}
				if len(dopts) != test.wantDialOptsLen {
					t.Fatalf("%s: found %d dialOptions, want %d", test.name, len(dopts), test.wantDialOptsLen)
				}
				return grpc.DialContext(ctx, target, dopts...)
			}
			defer func() {
				dialFunc = oldDialFunc
			}()
			if _, err := NewClient(test.opts); (err != nil) != test.wantErr {
				t.Fatalf("%s: NewClient(%+v) = %v, wantErr: %v", test.name, test.opts, err, test.wantErr)
			}
		}()
	}
}

func TestWatchForServiceUpdate(t *testing.T) {
	fakeServer, fakeCC, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	oldDialFunc := dialFunc
	dialFunc = func(_ context.Context, _ string, _ ...grpc.DialOption) (*grpc.ClientConn, error) {
		return fakeCC, nil
	}
	defer func() {
		dialFunc = oldDialFunc
	}()

	xdsClient, err := NewClient(Options{Config: validConfig})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	defer xdsClient.Close()
	t.Log("Created an xdsClient...")

	callbackCh := make(chan error, 1)
	cancelWatch := xdsClient.WatchForServiceUpdate(goodLDSTarget1, func(su ServiceUpdate, err error) {
		if su.Cluster != goodClusterName1 {
			callbackCh <- fmt.Errorf("got clusterName: %+v, want clusterName: %+v", su.Cluster, goodClusterName1)
			return
		}
		if err != nil {
			callbackCh <- fmt.Errorf("xdsClient.WatchForServiceUpdate returned error: %v", err)
			return
		}
		callbackCh <- nil
		return
	})
	defer cancelWatch()
	t.Log("Registered a watcher for service updates...")

	// Make the fakeServer send LDS and RDS responses.
	<-fakeServer.RequestChan
	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodLDSResponse1}
	<-fakeServer.RequestChan
	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodRDSResponse1}

	timer := time.NewTimer(defaultTestTimeout)
	select {
	case <-timer.C:
		t.Fatal("Timeout when expecting a service update")
	case err := <-callbackCh:
		timer.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestWatchForServiceUpdateWithNoResponseFromServer(t *testing.T) {
	fakeServer, fakeCC, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	oldDialFunc := dialFunc
	dialFunc = func(_ context.Context, _ string, _ ...grpc.DialOption) (*grpc.ClientConn, error) {
		return fakeCC, nil
	}
	defer func() {
		dialFunc = oldDialFunc
	}()

	xdsClient, err := NewClient(Options{Config: validConfig})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	defer xdsClient.Close()
	t.Log("Created an xdsClient...")

	oldWatchExpiryTimeout := defaultWatchExpiryTimeout
	defaultWatchExpiryTimeout = 1 * time.Second
	defer func() {
		defaultWatchExpiryTimeout = oldWatchExpiryTimeout
	}()

	callbackCh := make(chan error, 1)
	cancelWatch := xdsClient.WatchForServiceUpdate(goodLDSTarget1, func(su ServiceUpdate, err error) {
		if su.Cluster != "" {
			callbackCh <- fmt.Errorf("got clusterName: %+v, want empty clusterName", su.Cluster)
			return
		}
		if err == nil {
			callbackCh <- errors.New("xdsClient.WatchForServiceUpdate returned error non-nil error")
			return
		}
		callbackCh <- nil
		return
	})
	defer cancelWatch()
	t.Log("Registered a watcher for service updates...")

	// Wait for one request from the client, but send no reponses.
	<-fakeServer.RequestChan

	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		t.Fatal("Timeout when expecting a service update")
	case err := <-callbackCh:
		timer.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestWatchForServiceUpdateEmptyRDS(t *testing.T) {
	fakeServer, fakeCC, cleanup := fakexds.StartClientAndServer(t)
	defer cleanup()

	oldDialFunc := dialFunc
	dialFunc = func(_ context.Context, _ string, _ ...grpc.DialOption) (*grpc.ClientConn, error) {
		return fakeCC, nil
	}
	defer func() {
		dialFunc = oldDialFunc
	}()

	xdsClient, err := NewClient(Options{Config: validConfig})
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	defer xdsClient.Close()
	t.Log("Created an xdsClient...")

	oldWatchExpiryTimeout := defaultWatchExpiryTimeout
	defaultWatchExpiryTimeout = 1 * time.Second
	defer func() {
		defaultWatchExpiryTimeout = oldWatchExpiryTimeout
	}()

	callbackCh := make(chan error, 1)
	cancelWatch := xdsClient.WatchForServiceUpdate(goodLDSTarget1, func(su ServiceUpdate, err error) {
		if su.Cluster != "" {
			callbackCh <- fmt.Errorf("got clusterName: %+v, want empty clusterName", su.Cluster)
			return
		}
		if err == nil {
			callbackCh <- errors.New("xdsClient.WatchForServiceUpdate returned error non-nil error")
			return
		}
		callbackCh <- nil
		return
	})
	defer cancelWatch()
	t.Log("Registered a watcher for service updates...")

	// Send a good LDS response, but send an empty RDS response.
	<-fakeServer.RequestChan
	fakeServer.ResponseChan <- &fakexds.Response{Resp: goodLDSResponse1}
	<-fakeServer.RequestChan
	fakeServer.ResponseChan <- &fakexds.Response{Resp: noVirtualHostsInRDSResponse}

	timer := time.NewTimer(2 * time.Second)
	select {
	case <-timer.C:
		t.Fatal("Timeout when expecting a service update")
	case err := <-callbackCh:
		timer.Stop()
		if err != nil {
			t.Fatal(err)
		}
	}
}
