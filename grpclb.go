/*
 *
 * Copyright 2016 gRPC authors.
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

package grpc

import (
	"strings"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
	lbpb "google.golang.org/grpc/grpclb/grpc_lb_v1/messages"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const (
	lbTokeyKey             = "lb-token"
	defaultFallbackTimeout = 10 * time.Second
)

func convertDuration(d *lbpb.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return time.Duration(d.Seconds)*time.Second + time.Duration(d.Nanos)*time.Nanosecond
}

// Client API for LoadBalancer service.
// Mostly copied from generated pb.go file.
// To avoid circular dependency.
type loadBalancerClient struct {
	cc *ClientConn
}

func (c *loadBalancerClient) BalanceLoad(ctx context.Context, opts ...CallOption) (*balanceLoadClientStream, error) {
	desc := &StreamDesc{
		StreamName:    "BalanceLoad",
		ServerStreams: true,
		ClientStreams: true,
	}
	stream, err := NewClientStream(ctx, desc, c.cc, "/grpc.lb.v1.LoadBalancer/BalanceLoad", opts...)
	if err != nil {
		return nil, err
	}
	x := &balanceLoadClientStream{stream}
	return x, nil
}

type balanceLoadClientStream struct {
	ClientStream
}

func (x *balanceLoadClientStream) Send(m *lbpb.LoadBalanceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *balanceLoadClientStream) Recv() (*lbpb.LoadBalanceResponse, error) {
	m := new(lbpb.LoadBalanceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NewLBBuilder creates a builder for grpclb.
func NewLBBuilder() balancer.Builder {
	// TODO(bar grpclb) this function is exported for testing only, remove it when resolver supports selecting grpclb.
	return NewLBBuilderWithFallbackTimeout(defaultFallbackTimeout)
}

// NewLBBuilderWithFallbackTimeout creates a grpclb builder with the given
// fallbackTimeout.
//
// Only call this function when a non-default fallback timeout is needed.
func NewLBBuilderWithFallbackTimeout(fallbackTimeout time.Duration) balancer.Builder {
	return &lbBuilder{
		fallbackTimeout: fallbackTimeout,
	}
}

type lbBuilder struct {
	fallbackTimeout time.Duration
}

func (b *lbBuilder) Name() string {
	return "grpclb"
}

func (b *lbBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
	// This generates a manual resolver builder with a random scheme. This
	// scheme will be used to dial to remote LB, so we can send filtered address
	// updates to remote LB ClientConn using this manual resolver.
	r, cleanup := manual.GenerateAndRegisterManualResolver()

	var target string
	targetSplitted := strings.Split(cc.Target(), ":///")
	if len(targetSplitted) < 2 {
		target = cc.Target()
	} else {
		target = targetSplitted[1]
	}

	lb := &lbBalancer{
		cc:              cc,
		target:          target,
		opt:             opt,
		fallbackTimeout: b.fallbackTimeout,
		doneCh:          make(chan struct{}),

		manualResolver:     r,
		unregisterResolver: cleanup,
		csEvltr:            &connectivityStateEvaluator{},
		subConns:           make(map[resolver.Address]balancer.SubConn),
		scStates:           make(map[balancer.SubConn]connectivity.State),
		picker:             newPicker(nil, nil, nil, nil),
		clientStats:        &rpcStats{},
	}

	return lb
}

type lbBalancer struct {
	cc              balancer.ClientConn
	target          string
	opt             balancer.BuildOptions
	fallbackTimeout time.Duration
	doneCh          chan struct{}

	// manualResolver is used in the remote LB ClientConn inside grpclb. When
	// resolved address updates are received by grpclb, filtered updates will be
	// send to remote LB ClientConn through this resolver.
	manualResolver *manual.Resolver
	// Call this function in close to unregister the manual resolver.
	unregisterResolver func()
	// The ClientConn to talk to the remote balancer.
	ccRemoteLB *ClientConn

	mu sync.Mutex
	// The full server list including drops, used to check if the newly received
	// serverList contains anything new. Each generate picker will also have
	// reference to this list to do the first layer pick.
	fullServerList []*lbpb.Server
	// All backends addresses, with metadata set to nil. This list contains all
	// backend addresses in the same order and with the same duplicates as in
	// serverlist. When generating picker, a SubConn slice with the same order
	// but with only READY SCs will be gerenated.
	backendAddrs []resolver.Address
	// Roundrobin functionalities.
	csEvltr  *connectivityStateEvaluator
	state    connectivity.State
	subConns map[resolver.Address]balancer.SubConn   // Used to new/remove SubConn.
	scStates map[balancer.SubConn]connectivity.State // Used to filter READY SubConns.
	picker   *lbPicker
	// Support fallback to resolved backend addresses if there's no response
	// from remote balancer within fallbackTimeout.
	fallbackTimerExpired bool
	serverListReceived   bool
	// resolvedBackendAddrs is resolvedAddrs minus remote balancers. It's set
	// when resolved address updates are received, and read in the goroutine
	// handling fallback.
	resolvedBackendAddrs []resolver.Address

	// Support client side load reporting. Each picker gets a reference to this,
	// and will update its content.
	clientStats *rpcStats
}

// regeneratePicker takes a snapshot of the balancer, and generate a picker from
// it. The picker
//  - always return ErrTransientFailure if the balancer is in TransientFailure,
//  - do two layer roundrobin pick otherwise.
// Caller must hold lb.mu.
func (lb *lbBalancer) regeneratePicker() {
	if lb.state == connectivity.TransientFailure {
		lb.picker = newPicker(balancer.ErrTransientFailure, nil, nil, nil)
		return
	}
	var readySCs []balancer.SubConn
	for _, a := range lb.backendAddrs {
		if sc, ok := lb.subConns[a]; ok {
			if st, ok := lb.scStates[sc]; ok && st == connectivity.Ready {
				readySCs = append(readySCs, sc)
			}
		}
	}
	lb.picker = newPicker(nil, lb.fullServerList, readySCs, lb.clientStats)
	return
}

func (lb *lbBalancer) HandleSubConnStateChange(sc balancer.SubConn, s connectivity.State) {
	grpclog.Infof("lbBalancer: handle SubConn state change: %p, %v", sc, s)
	lb.mu.Lock()
	defer lb.mu.Unlock()

	oldS, ok := lb.scStates[sc]
	if !ok {
		grpclog.Infof("lbBalancer: got state changes for an unknown SubConn: %p, %v", sc, s)
		return
	}
	lb.scStates[sc] = s
	switch s {
	case connectivity.Idle:
		sc.Connect()
	case connectivity.Shutdown:
		// When an address was removed by resolver, b called RemoveSubConn but
		// kept the sc's state in scStates. Remove state for this sc here.
		delete(lb.scStates, sc)
	}

	oldAggrState := lb.state
	lb.state = lb.csEvltr.recordTransition(oldS, s)

	// Regenerate picker when one of the following happens:
	//  - this sc became ready from not-ready
	//  - this sc became not-ready from ready
	//  - the aggregated state of balancer became TransientFailure from non-TransientFailure
	//  - the aggregated state of balancer became non-TransientFailure from TransientFailure
	if (oldS == connectivity.Ready) != (s == connectivity.Ready) ||
		(lb.state == connectivity.TransientFailure) != (oldAggrState == connectivity.TransientFailure) {
		lb.regeneratePicker()
	}

	lb.cc.UpdateBalancerState(lb.state, lb.picker)
	return
}

// fallbackToBackendsAfter blocks for fallbackTimeout and falls back to use
// resolved backends (backends received from resolver, not from remote balancer)
// if no connection to remote balancers was successful.
func (lb *lbBalancer) fallbackToBackendsAfter(fallbackTimeout time.Duration) {
	timer := time.NewTimer(fallbackTimeout)
	defer func() {
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
	}()
	select {
	case <-timer.C:
	case <-lb.doneCh:
		return
	}
	lb.mu.Lock()
	if lb.serverListReceived {
		lb.mu.Unlock()
		return
	}
	lb.fallbackTimerExpired = true
	lb.refreshSubConns(lb.resolvedBackendAddrs)
	lb.mu.Unlock()
}

// HandleResolvedAddrs sends the updated remoteLB addresses to remoteLB
// clientConn. The remoteLB clientConn will handle creating/removing remoteLB
// connections.
func (lb *lbBalancer) HandleResolvedAddrs(addrs []resolver.Address, err error) {
	grpclog.Infof("lbBalancer: handleResolvedResult: %+v", addrs)
	if len(addrs) <= 0 {
		return
	}

	var (
		remoteBalancerAddrs []resolver.Address
		backendAddrs        []resolver.Address
	)
	for _, a := range addrs {
		if a.Type == resolver.GRPCLB {
			remoteBalancerAddrs = append(remoteBalancerAddrs, a)
		} else {
			backendAddrs = append(backendAddrs, a)
		}
	}

	if lb.ccRemoteLB == nil {
		if len(remoteBalancerAddrs) <= 0 {
			grpclog.Fatalf("grpclb: no remote balancer address is available, should never happen")
		}
		// First time receiving resolved addresses, create a cc to remote
		// balancers.
		lb.dialRemoteLB(remoteBalancerAddrs[0].ServerName)
		// Start the fallback goroutine.
		go lb.fallbackToBackendsAfter(lb.fallbackTimeout)
	}

	// cc to remote balancers uses lb.manualResolver. Send the updated remote
	// balancer addresses to it through manualResolver.
	lb.manualResolver.NewAddress(remoteBalancerAddrs)

	lb.mu.Lock()
	lb.resolvedBackendAddrs = backendAddrs
	// If serverListReceived is true, connection to remote balancer was
	// successful and there's no need to do fallback anymore.
	// If fallbackTimerExpired is false, fallback hasn't happened yet.
	if !lb.serverListReceived && lb.fallbackTimerExpired {
		// This means we received a new list of resolved backends, and we are
		// still in fallback mode. Need to update the list of backends we are
		// using to the new list of backends.
		lb.refreshSubConns(lb.resolvedBackendAddrs)
	}
	lb.mu.Unlock()
}

func (lb *lbBalancer) Close() {
	select {
	case <-lb.doneCh:
		return
	default:
	}
	close(lb.doneCh)
	if lb.ccRemoteLB != nil {
		lb.ccRemoteLB.Close()
	}
}
