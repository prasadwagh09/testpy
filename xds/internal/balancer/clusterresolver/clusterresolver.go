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

// Package clusterresolver contains the implementation of the
// cluster_resolver_experimental LB policy which resolves endpoint addresses
// using a list of one or more discovery mechanisms.
package clusterresolver

import (
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/internal/balancer/nop"
	"google.golang.org/grpc/internal/buffer"
	"google.golang.org/grpc/internal/envconfig"
	"google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/internal/pretty"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	"google.golang.org/grpc/xds/internal/balancer/outlierdetection"
	"google.golang.org/grpc/xds/internal/balancer/priority"
	"google.golang.org/grpc/xds/internal/xdsclient"
	"google.golang.org/grpc/xds/internal/xdsclient/xdsresource"
)

// Name is the name of the cluster_resolver balancer.
const Name = "cluster_resolver_experimental"

var (
	errBalancerClosed = errors.New("cdsBalancer is closed")
	newChildBalancer  = func(bb balancer.Builder, cc balancer.ClientConn, o balancer.BuildOptions) balancer.Balancer {
		return bb.Build(cc, o)
	}
)

func init() {
	balancer.Register(bb{})
}

type bb struct{}

// Build helps implement the balancer.Builder interface.
func (bb) Build(cc balancer.ClientConn, opts balancer.BuildOptions) balancer.Balancer {
	priorityBuilder := balancer.Get(priority.Name)
	if priorityBuilder == nil {
		logger.Errorf("%q LB policy is needed but not registered", priority.Name)
		return nop.NewBalancer(cc, fmt.Errorf("%q LB policy is needed but not registered", priority.Name))
	}
	priorityConfigParser, ok := priorityBuilder.(balancer.ConfigParser)
	if !ok {
		logger.Errorf("%q LB policy does not implement a config parser", priority.Name)
		return nop.NewBalancer(cc, fmt.Errorf("%q LB policy does not implement a config parser", priority.Name))
	}

	b := &clusterResolverBalancer{
		bOpts:    opts,
		updateCh: buffer.NewUnbounded(),
		closed:   grpcsync.NewEvent(),
		done:     grpcsync.NewEvent(),

		priorityBuilder:      priorityBuilder,
		priorityConfigParser: priorityConfigParser,
	}
	b.logger = prefixLogger(b)
	b.logger.Infof("Created")

	b.resourceWatcher = newResourceResolver(b, b.logger)
	b.cc = &ccWrapper{
		ClientConn:      cc,
		b:               b,
		resourceWatcher: b.resourceWatcher,
	}

	go b.run()
	return b
}

func (bb) Name() string {
	return Name
}

func (bb) ParseConfig(j json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	odBuilder := balancer.Get(outlierdetection.Name)
	if odBuilder == nil {
		// Shouldn't happen, registered through imported Outlier Detection,
		// defensive programming.
		return nil, fmt.Errorf("%q LB policy is needed but not registered", outlierdetection.Name)
	}
	odParser, ok := odBuilder.(balancer.ConfigParser)
	if !ok {
		// Shouldn't happen, imported Outlier Detection builder has this method.
		return nil, fmt.Errorf("%q LB policy does not implement a config parser", outlierdetection.Name)
	}

	var cfg *LBConfig
	if err := json.Unmarshal(j, &cfg); err != nil {
		return nil, fmt.Errorf("unable to unmarshal balancer config %s into cluster-resolver config, error: %v", string(j), err)
	}

	if envconfig.XDSOutlierDetection {
		for i, dm := range cfg.DiscoveryMechanisms {
			lbCfg, err := odParser.ParseConfig(dm.OutlierDetection)
			if err != nil {
				return nil, fmt.Errorf("error parsing Outlier Detection config %v: %v", dm.OutlierDetection, err)
			}
			odCfg, ok := lbCfg.(*outlierdetection.LBConfig)
			if !ok {
				// Shouldn't happen, Parser built at build time with Outlier Detection
				// builder pulled from gRPC LB Registry.
				return nil, fmt.Errorf("odParser returned config with unexpected type %T: %v", lbCfg, lbCfg)
			}
			cfg.DiscoveryMechanisms[i].outlierDetection = *odCfg
		}
	}
	if err := json.Unmarshal(cfg.XDSLBPolicy, &cfg.xdsLBPolicy); err != nil {
		// This will never occur, valid configuration is emitted from the xDS
		// Client. Validity is already checked in the xDS Client, however, this
		// double validation is present because Unmarshalling and Validating are
		// coupled into one json.Unmarshal operation). We will switch this in
		// the future to two separate operations.
		return nil, fmt.Errorf("error unmarshaling xDS LB Policy: %v", err)
	}
	return cfg, nil
}

// ccUpdate wraps a clientConn update received from gRPC.
type ccUpdate struct {
	state balancer.ClientConnState
	err   error
}

type exitIdle struct{}

// clusterResolverBalancer resolves endpoint addresses using a list of one or
// more discovery mechanisms.
type clusterResolverBalancer struct {
	cc              balancer.ClientConn
	bOpts           balancer.BuildOptions
	updateCh        *buffer.Unbounded // Channel for updates from gRPC.
	resourceWatcher *resourceResolver
	logger          *grpclog.PrefixLogger
	closed          *grpcsync.Event
	done            *grpcsync.Event

	priorityBuilder      balancer.Builder
	priorityConfigParser balancer.ConfigParser

	config          *LBConfig
	configRaw       *serviceconfig.ParseResult
	xdsClient       xdsclient.XDSClient    // xDS client to watch EDS resource.
	attrsWithClient *attributes.Attributes // Attributes with xdsClient attached to be passed to the child policies.

	child               balancer.Balancer
	priorities          []priorityConfig
	watchUpdateReceived bool
}

// handleClientConnUpdate handles a ClientConnUpdate received from gRPC.
//
// A good update results in creation of endpoint resolvers for the configured
// discovery mechanisms. An update with an error results in cancellation of any
// existing endpoint resolution and propagation of the same to the child policy.
func (b *clusterResolverBalancer) handleClientConnUpdate(update *ccUpdate) {
	if err := update.err; err != nil {
		b.handleErrorFromUpdate(err, true)
		return
	}

	b.logger.Infof("Received new balancer config: %v", pretty.ToJSON(update.state.BalancerConfig))
	cfg, _ := update.state.BalancerConfig.(*LBConfig)
	if cfg == nil {
		b.logger.Warningf("Ignoring unsupported balancer configuration of type: %T", update.state.BalancerConfig)
		return
	}

	b.config = cfg
	b.configRaw = update.state.ResolverState.ServiceConfig
	b.resourceWatcher.updateMechanisms(cfg.DiscoveryMechanisms)

	// The child policy is created only after all configured discovery
	// mechanisms have been successfully returned endpoints. If that is not the
	// case, we return early.
	if !b.watchUpdateReceived {
		return
	}
	b.updateChildConfig()
}

// handleResourceUpdate handles a resource update or error from the resource
// resolver by propagating the same to the child LB policy.
func (b *clusterResolverBalancer) handleResourceUpdate(update *resourceUpdate) {
	if err := update.err; err != nil {
		b.handleErrorFromUpdate(err, false)
		return
	}

	b.watchUpdateReceived = true
	b.priorities = update.priorities

	// An update from the resource resolver contains resolved endpoint addresses
	// for all configured discovery mechanisms ordered by priority. This is used
	// to generate configuration for the priority LB policy.
	b.updateChildConfig()
}

// updateChildConfig builds child policy configuration using endpoint addresses
// returned by the resource resolver and child policy configuration provided by
// parent LB policy.
//
// A child policy is created if one doesn't already exist. The newly built
// configuration is then pushed to the child policy.
func (b *clusterResolverBalancer) updateChildConfig() {
	if b.child == nil {
		b.child = newChildBalancer(b.priorityBuilder, b.cc, b.bOpts)
	}

	childCfgBytes, addrs, err := buildPriorityConfigJSON(b.priorities, &b.config.xdsLBPolicy)
	if err != nil {
		b.logger.Warningf("Failed to build child policy config: %v", err)
		return
	}
	childCfg, err := b.priorityConfigParser.ParseConfig(childCfgBytes)
	if err != nil {
		b.logger.Warningf("Failed to parse child policy config. This should never happen because the config was generated: %v", err)
		return
	}
	b.logger.Infof("Built child policy config: %v", pretty.ToJSON(childCfg))

	endpoints := make([]resolver.Endpoint, len(addrs))
	for i, a := range addrs {
		endpoints[i].Attributes = a.BalancerAttributes
		endpoints[i].Addresses = []resolver.Address{a}
		endpoints[i].Addresses[0].BalancerAttributes = nil
	}
	if err := b.child.UpdateClientConnState(balancer.ClientConnState{
		ResolverState: resolver.State{
			Endpoints:     endpoints,
			Addresses:     addrs,
			ServiceConfig: b.configRaw,
			Attributes:    b.attrsWithClient,
		},
		BalancerConfig: childCfg,
	}); err != nil {
		b.logger.Warningf("Failed to push config to child policy: %v", err)
	}
}

// handleErrorFromUpdate handles errors from the parent LB policy and endpoint
// resolvers. fromParent is true if error is from the parent LB policy. In both
// cases, the error is propagated to the child policy, if one exists.
func (b *clusterResolverBalancer) handleErrorFromUpdate(err error, fromParent bool) {
	b.logger.Warningf("Received error: %v", err)

	// A resource-not-found error from the parent LB policy means that the LDS
	// or CDS resource was removed. This should result in endpoint resolvers
	// being stopped here.
	//
	// A resource-not-found error from the EDS endpoint resolver means that the
	// EDS resource was removed. No action needs to be taken for this, and we
	// should continue watching the same EDS resource.
	if fromParent && xdsresource.ErrType(err) == xdsresource.ErrorTypeResourceNotFound {
		b.resourceWatcher.stop()
	}

	if b.child != nil {
		b.child.ResolverError(err)
		return
	}
	b.cc.UpdateState(balancer.State{
		ConnectivityState: connectivity.TransientFailure,
		Picker:            base.NewErrPicker(err),
	})
}

// run is a long-running goroutine that handles updates from gRPC and endpoint
// resolvers. The methods handling the individual updates simply push them onto
// a channel which is read and acted upon from here.
func (b *clusterResolverBalancer) run() {
	for {
		select {
		case u, ok := <-b.updateCh.Get():
			if !ok {
				return
			}
			b.updateCh.Load()
			switch update := u.(type) {
			case *ccUpdate:
				b.handleClientConnUpdate(update)
			case exitIdle:
				if b.child == nil {
					b.logger.Errorf("xds: received ExitIdle with no child balancer")
					break
				}
				// This implementation assumes the child balancer supports
				// ExitIdle (but still checks for the interface's existence to
				// avoid a panic if not).  If the child does not, no subconns
				// will be connected.
				if ei, ok := b.child.(balancer.ExitIdler); ok {
					ei.ExitIdle()
				}
			}
		case u := <-b.resourceWatcher.updateChannel:
			b.handleResourceUpdate(u)

		// Close results in stopping the endpoint resolvers and closing the
		// underlying child policy and is the only way to exit this goroutine.
		case <-b.closed.Done():
			b.resourceWatcher.stop()

			if b.child != nil {
				b.child.Close()
				b.child = nil
			}
			b.updateCh.Close()
			// This is the *ONLY* point of return from this function.
			b.logger.Infof("Shutdown")
			b.done.Fire()
			return
		}
	}
}

// Following are methods to implement the balancer interface.

func (b *clusterResolverBalancer) UpdateClientConnState(state balancer.ClientConnState) error {
	if b.closed.HasFired() {
		b.logger.Warningf("Received update from gRPC {%+v} after close", state)
		return errBalancerClosed
	}

	if b.xdsClient == nil {
		c := xdsclient.FromResolverState(state.ResolverState)
		if c == nil {
			return balancer.ErrBadResolverState
		}
		b.xdsClient = c
		b.attrsWithClient = state.ResolverState.Attributes
	}

	b.updateCh.Put(&ccUpdate{state: state})
	return nil
}

// ResolverError handles errors reported by the xdsResolver.
func (b *clusterResolverBalancer) ResolverError(err error) {
	if b.closed.HasFired() {
		b.logger.Warningf("Received resolver error {%v} after close", err)
		return
	}
	b.updateCh.Put(&ccUpdate{err: err})
}

// UpdateSubConnState handles subConn updates from gRPC.
func (b *clusterResolverBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
	b.logger.Errorf("UpdateSubConnState(%v, %+v) called unexpectedly", sc, state)
}

// Close closes the cdsBalancer and the underlying child balancer.
func (b *clusterResolverBalancer) Close() {
	b.closed.Fire()
	<-b.done.Done()
}

func (b *clusterResolverBalancer) ExitIdle() {
	b.updateCh.Put(exitIdle{})
}

// ccWrapper overrides ResolveNow(), so that re-resolution from the child
// policies will trigger the DNS resolver in cluster_resolver balancer.  It
// also intercepts NewSubConn calls in case children don't set the
// StateListener, to allow redirection to happen via this cluster_resolver
// balancer.
type ccWrapper struct {
	balancer.ClientConn
	b               *clusterResolverBalancer
	resourceWatcher *resourceResolver
}

func (c *ccWrapper) ResolveNow(resolver.ResolveNowOptions) {
	c.resourceWatcher.resolveNow()
}

func (c *ccWrapper) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (sc balancer.SubConn, err error) {
	// No need to override opts.StateListener; just forward all calls to the
	// child that created the SubConn.
	return c.ClientConn.NewSubConn(addrs, opts)
}
