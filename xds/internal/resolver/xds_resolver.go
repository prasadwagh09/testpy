/*
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

// Package resolver implements the xds resolver, that does LDS and RDS to find
// the cluster to use.
package resolver

import (
	"fmt"

	"google.golang.org/grpc/internal/grpclog"
	"google.golang.org/grpc/internal/grpcsync"
	"google.golang.org/grpc/resolver"

	iresolver "google.golang.org/grpc/internal/resolver"
	xdsclient "google.golang.org/grpc/xds/internal/client"
)

const xdsScheme = "xds"

// For overriding in unittests.
var newXDSClient = func() (xdsClientInterface, error) { return xdsclient.New() }

func init() {
	resolver.Register(&xdsResolverBuilder{})
}

type xdsResolverBuilder struct{}

// Build helps implement the resolver.Builder interface.
//
// The xds bootstrap process is performed (and a new xds client is built) every
// time an xds resolver is built.
func (b *xdsResolverBuilder) Build(t resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	r := &xdsResolver{
		target:         t,
		cc:             cc,
		closed:         grpcsync.NewEvent(),
		updateCh:       make(chan suWithError, 1),
		activeClusters: make(map[string]*clusterInfo),
	}
	r.logger = prefixLogger((r))
	r.logger.Infof("Creating resolver for target: %+v", t)

	client, err := newXDSClient()
	if err != nil {
		return nil, fmt.Errorf("xds: failed to create xds-client: %v", err)
	}
	r.client = client

	// Register a watch on the xdsClient for the user's dial target.
	cancelWatch := watchService(r.client, r.target.Endpoint, r.handleServiceUpdate, r.logger)
	r.logger.Infof("Watch started on resource name %v with xds-client %p", r.target.Endpoint, r.client)
	r.cancelWatch = func() {
		cancelWatch()
		r.logger.Infof("Watch cancel on resource name %v with xds-client %p", r.target.Endpoint, r.client)
	}

	go r.run()
	return r, nil
}

// Name helps implement the resolver.Builder interface.
func (*xdsResolverBuilder) Scheme() string {
	return xdsScheme
}

// xdsClientInterface contains methods from xdsClient.Client which are used by
// the resolver. This will be faked out in unittests.
type xdsClientInterface interface {
	WatchListener(serviceName string, cb func(xdsclient.ListenerUpdate, error)) func()
	WatchRouteConfig(routeName string, cb func(xdsclient.RouteConfigUpdate, error)) func()
	Close()
}

// suWithError wraps the ServiceUpdate and error received through a watch API
// callback, so that it can pushed onto the update channel as a single entity.
type suWithError struct {
	su          serviceUpdate
	emptyUpdate bool
	err         error
}

// xdsResolver implements the resolver.Resolver interface.
//
// It registers a watcher for ServiceConfig updates with the xdsClient object
// (which performs LDS/RDS queries for the same), and passes the received
// updates to the ClientConn.
type xdsResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	closed *grpcsync.Event

	logger *grpclog.PrefixLogger

	// The underlying xdsClient which performs all xDS requests and responses.
	client xdsClientInterface
	// A channel for the watch API callback to write service updates on to. The
	// updates are read by the run goroutine and passed on to the ClientConn.
	updateCh chan suWithError
	// cancelWatch is the function to cancel the watcher.
	cancelWatch func()

	// activeClusters is a map from cluster name to a ref count.  Only read or
	// written during a service update (synchronous).
	activeClusters map[string]*clusterInfo

	curConfigSelector *configSelector
}

// run is a long running goroutine which blocks on receiving service updates
// and passes it on the ClientConn.
func (r *xdsResolver) run() {
	for {
		select {
		case <-r.closed.Done():
			return
		case update := <-r.updateCh:
			if update.err != nil {
				r.logger.Warningf("Watch error on resource %v from xds-client %p, %v", r.target.Endpoint, r.client, update.err)
				if xdsclient.ErrType(update.err) == xdsclient.ErrorTypeResourceNotFound {
					// If error is resource-not-found, it means the LDS resource
					// was removed. Send an empty service config, which picks
					// pick-first, with no address, and puts the ClientConn into
					// transient failure..
					r.cc.UpdateState(resolver.State{
						ServiceConfig: r.cc.ParseServiceConfig("{}"),
					})
					r.curConfigSelector.decRefs()
					r.curConfigSelector = nil
					continue
				}
				// Send error to ClientConn, and balancers, if error is not
				// resource not found.
				r.cc.ReportError(update.err)
				continue
			}
			var cs *configSelector
			if !update.emptyUpdate {
				var err error
				if cs, err = r.serviceUpdateToCS(update.su); err != nil {
					r.logger.Warningf("%v", err)
					r.cc.ReportError(err)
					continue
				}
			} else {
				cs = r.curConfigSelector
			}
			cs.incRefs()
			sc, err := r.serviceConfigJSON()
			if err != nil {
				r.logger.Warningf("%v", err)
				r.cc.ReportError(err)
				cs.decRefs()
				continue
			}
			r.logger.Infof("Received update on resource %v from xds-client %p, generated service config: %v", r.target.Endpoint, r.client, sc)
			state := iresolver.SetConfigSelector(resolver.State{
				ServiceConfig: r.cc.ParseServiceConfig(sc),
			}, cs)
			r.cc.UpdateState(state)
			r.curConfigSelector.decRefs()
			r.curConfigSelector = cs
		}
	}
}

// handleServiceUpdate is the callback which handles service updates. It writes
// the received update to the update channel, which is picked by the run
// goroutine.
func (r *xdsResolver) handleServiceUpdate(su serviceUpdate, err error) {
	if r.closed.HasFired() {
		// Do not pass updates to the ClientConn once the resolver is closed.
		return
	}
	r.updateCh <- suWithError{su: su, err: err}
}

// ResolveNow is a no-op at this point.
func (*xdsResolver) ResolveNow(o resolver.ResolveNowOptions) {}

// Close closes the resolver, and also closes the underlying xdsClient.
func (r *xdsResolver) Close() {
	r.cancelWatch()
	r.client.Close()
	r.closed.Fire()
	r.logger.Infof("Shutdown")
}
