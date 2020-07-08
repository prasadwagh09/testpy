/*
 *
 * Copyright 2020 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package test

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	testpb "google.golang.org/grpc/test/grpc_testing"
)

// unixServer is used to test servers listening over a unix socket.
type unixServer struct {
	// Guarantees we satisfy this interface; panics if unimplemented methods are called.
	testpb.TestServiceServer

	// Customizable implementations of server handlers.
	emptyCall func(ctx context.Context, in *testpb.Empty) (*testpb.Empty, error)

	// A client connected to this service the test may use.  Created in Start().
	client testpb.TestServiceClient
	cc     *grpc.ClientConn
	s      *grpc.Server

	cleanups []func() // Lambdas executed in Stop(); populated by Start().
}

func (us *unixServer) EmptyCall(ctx context.Context, in *testpb.Empty) (*testpb.Empty, error) {
	return us.emptyCall(ctx, in)
}

func (us *unixServer) Start(network, address, target string) error {
	lis, err := net.Listen(network, address)
	if err != nil {
		return err
	}
	us.cleanups = append(us.cleanups, func() { lis.Close() })

	s := grpc.NewServer()
	testpb.RegisterTestServiceServer(s, us)
	go s.Serve(lis)
	us.cleanups = append(us.cleanups, s.Stop)
	us.s = s

	cc, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("grpc.Dial(%q) = %v", target, err)
	}
	us.cc = cc
	if err := us.waitForReady(cc); err != nil {
		return err
	}
	us.cleanups = append(us.cleanups, func() { cc.Close() })

	us.client = testpb.NewTestServiceClient(cc)

	return nil
}

func (us *unixServer) waitForReady(cc *grpc.ClientConn) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
		s := cc.GetState()
		if s == connectivity.Ready {
			return nil
		}
		if !cc.WaitForStateChange(ctx, s) {
			return ctx.Err()
		}
	}
}

func (us *unixServer) Stop() {
	for i := len(us.cleanups) - 1; i >= 0; i-- {
		us.cleanups[i]()
	}
}

func runUnixTest(t *testing.T, address, target, expectedAuthority string) {
	if err := os.RemoveAll(address); err != nil {
		t.Fatalf("Error removing socket file %v: %v\n", address, err)
	}
	us := &unixServer{
		emptyCall: func(ctx context.Context, in *testpb.Empty) (*testpb.Empty, error) {
			if md, ok := metadata.FromIncomingContext(ctx); ok {
				if auths, ok := md[":authority"]; ok {
					if len(auths) < 1 {
						return nil, status.Error(codes.Unauthenticated, "no authority header")
					}
					if auths[0] != expectedAuthority {
						return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("invalid authority header %v, expected %v", auths[0], expectedAuthority))
					}
				} else {
					return nil, status.Error(codes.Unauthenticated, "no authority header")
				}
			} else {
				return nil, status.Error(codes.Unauthenticated, "failed to parse metadata")
			}
			return &testpb.Empty{}, nil
		},
	}
	if err := us.Start("unix", address, target); err != nil {
		t.Fatalf("Error starting endpoint server: %v\n", err)
		return
	}
	defer us.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := us.client.EmptyCall(ctx, &testpb.Empty{})
	if err != nil {
		t.Errorf("us.client.EmptyCall(_, _) = _, %v; want _, nil\n", err)
	}
}

func (s) TestUnix1(t *testing.T) {
	runUnixTest(t, "sock.sock", "unix:sock.sock", "localhost")
}

func (s) TestUnix2(t *testing.T) {
	runUnixTest(t, "/tmp/sock.sock", "unix:/tmp/sock.sock", "localhost")
}

func (s) TestUnix3(t *testing.T) {
	runUnixTest(t, "/tmp/sock.sock", "unix:///tmp/sock.sock", "localhost")
}
