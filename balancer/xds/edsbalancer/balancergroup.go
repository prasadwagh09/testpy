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
 */

package edsbalancer

import (
	"context"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
)

type pickerState struct {
	weight uint32
	picker balancer.Picker
	state  connectivity.State
}

// balancerGroup is a group of balancers with weights.
type balancerGroup struct {
	cc balancer.ClientConn

	mu           sync.Mutex
	idToBalancer map[string]balancer.Balancer
	// A separate mutex??? No, because we usually want to read the other map.
	scToID map[balancer.SubConn]string
	// All balancer ID exists as a key in this map. IDs not in map are either
	// removed or never existed.
	idToPickerState map[string]*pickerState
}

func newBalancerGroup(cc balancer.ClientConn) *balancerGroup {
	return &balancerGroup{
		cc: cc,

		idToBalancer:    make(map[string]balancer.Balancer),
		scToID:          make(map[balancer.SubConn]string),
		idToPickerState: make(map[string]*pickerState),
	}
}

// add adds a balancer built by builder to the group, with given id and picking weight.
func (bg *balancerGroup) add(id string, weight uint32, builder balancer.Builder) {
	bg.mu.Lock()
	if _, ok := bg.idToBalancer[id]; ok {
		bg.mu.Unlock()
		grpclog.Warningf("balancer group: adding a balancer with existing ID: %s", id)
		return
	}
	bg.mu.Unlock()
	bgcc := &balancerGroupCC{
		id:    id,
		group: bg,
	}
	b := builder.Build(bgcc, balancer.BuildOptions{})
	bg.mu.Lock()
	bg.idToBalancer[id] = b
	bg.idToPickerState[id] = &pickerState{
		weight: weight,
		// Start everything in IDLE. We don't count IDLE when aggregating (as
		// opposite to e.g. READY, 1 READY results in overall READY).
		state: connectivity.Idle,
	}
	bg.mu.Unlock()
}

func (bg *balancerGroup) remove(id string) {
	bg.mu.Lock()
	defer bg.mu.Unlock()

	// Close balancer.
	if b, ok := bg.idToBalancer[id]; ok {
		b.Close()
		delete(bg.idToBalancer, id)
	}
	// Remove picker. This also results in future updates for this ID to be
	// ignored.
	delete(bg.idToPickerState, id)
	// Remove SubConns.
	for sc, bid := range bg.scToID {
		if bid == id {
			bg.cc.RemoveSubConn(sc)
			delete(bg.scToID, sc)
		}
	}

	// Update state and picker to reflect the changes.
	bg.cc.UpdateBalancerState(buildPickerAndState(bg.idToPickerState))
}

func (bg *balancerGroup) changeWeight(id string, newWeight uint32) {
	// NOTE: This probably doesn't need to update the picker. But it seems better
	// to do the update because it's still a change in the picker (which is
	// balancer's snapshot).
	bg.mu.Lock()
	defer bg.mu.Unlock()

	pState, ok := bg.idToPickerState[id]
	if !ok {
		return
	}
	if pState.weight == newWeight {
		return
	}
	pState.weight = newWeight

	// Update state and picker to reflect the changes.
	bg.cc.UpdateBalancerState(buildPickerAndState(bg.idToPickerState))
}

// Actions from ClientConn, forward to sub-balancers.

// SubConn state change: find the corresponding balancer and then forward.
func (bg *balancerGroup) handleSubConnStateChange(sc balancer.SubConn, state connectivity.State) {
	grpclog.Infof("balancer group: handle subconn state change: %p, %v", sc, state)
	bg.mu.Lock()
	var b balancer.Balancer
	if id, ok := bg.scToID[sc]; ok {
		if state == connectivity.Shutdown {
			// Only delete sc from the map when state changed to Shutdown.
			delete(bg.scToID, sc)
		}
		b = bg.idToBalancer[id]
	}
	bg.mu.Unlock()
	if b == nil {
		grpclog.Infof("balancer group: balancer not found for sc state change")
		return
	}

	b.HandleSubConnStateChange(sc, state)
}

// Address change: forward to balancer.
func (bg *balancerGroup) handleResolvedAddrs(id string, addrs []resolver.Address) {
	bg.mu.Lock()
	b, ok := bg.idToBalancer[id]
	bg.mu.Unlock()
	if !ok {
		grpclog.Infof("balancer group: balancer with id %q not found", id)
		return
	}
	b.HandleResolvedAddrs(addrs, nil)
}

// TODO: handleServiceConfig()
//
// For BNS address for slicer, comes from endpoint.Metadata. It will be sent
// from parent to sub-balancers as service config.

// Actions from sub-balancers, forward to ClientConn.

// New SubConn: forward to ClientConn, and also create a map from sc to
// balancer, so state update will find the right balancer.
//
// One note about removing SubConn: only forward to ClientConn, but not delete
// from map. Delete sc from the map only when state changes to Shutdown. Since
// it's just forwarding the action, there's no need for a removeSubConn()
// wrapper function.
func (bg *balancerGroup) newSubConn(id string, addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	sc, err := bg.cc.NewSubConn(addrs, opts)
	if err != nil {
		return nil, err
	}

	bg.mu.Lock()
	bg.scToID[sc] = id
	bg.mu.Unlock()

	return sc, nil
}

func (bg *balancerGroup) updateBalancerState(id string, state connectivity.State, picker balancer.Picker) {
	grpclog.Infof("balancer group: update balancer state: %v, %v, %p", id, state, picker)
	bg.mu.Lock()
	defer bg.mu.Unlock()
	pickerSt, ok := bg.idToPickerState[id]
	if !ok {
		// All state starts in IDLE. It ID is not in map, it's either removed,
		// or never existed.
		grpclog.Infof("balancer group: pickerState not found when update picker/state")
		return
	}

	pickerSt.picker = picker
	pickerSt.state = state

	bg.cc.UpdateBalancerState(buildPickerAndState(bg.idToPickerState))
}

func (bg *balancerGroup) close() {
	for _, b := range bg.idToBalancer {
		b.Close()
	}
}

func buildPickerAndState(m map[string]*pickerState) (connectivity.State, balancer.Picker) {
	var readyN, connectingN int
	readyPickerWithWeights := make([]pickerState, 0, len(m))
	for _, ps := range m {
		switch ps.state {
		case connectivity.Ready:
			readyN++
			readyPickerWithWeights = append(readyPickerWithWeights, *ps)
		case connectivity.Connecting:
			connectingN++
		}
	}
	var aggregatedState connectivity.State
	switch {
	case readyN > 0:
		aggregatedState = connectivity.Ready
	case connectingN > 0:
		aggregatedState = connectivity.Connecting
	default:
		aggregatedState = connectivity.TransientFailure
	}
	if aggregatedState == connectivity.TransientFailure {
		return aggregatedState, base.NewErrPicker(balancer.ErrTransientFailure)
	}

	return aggregatedState, newPickerGroup(readyPickerWithWeights)
}

type pickerGroup struct {
	readyPickerWithWeights []pickerState
	length                 int

	mu    sync.Mutex
	idx   int    // The index of the picker that will be picked
	count uint32 // The number of times the current picker has been picked.
}

// newPickerGroup takes pickers with weights, and group them into one picker.
//
// Note it only takes ready pickers. The map shouldn't contain non-ready
// pickers.
//
// TODO: (bg) confirm this is the expected behavior: non-ready balancers should
// be ignored when picking. Only ready balancers are picked.
func newPickerGroup(readyPickerWithWeights []pickerState) *pickerGroup {
	return &pickerGroup{
		readyPickerWithWeights: readyPickerWithWeights,
		length:                 len(readyPickerWithWeights),
	}
}

func (pg *pickerGroup) Pick(ctx context.Context, opts balancer.PickOptions) (conn balancer.SubConn, done func(balancer.DoneInfo), err error) {
	if pg.length <= 0 {
		return nil, nil, balancer.ErrNoSubConnAvailable
	}

	// TODO: the WRR algorithm needs a design.
	// MAYBE: move WRR implmentation to util.go as a separate struct.
	pg.mu.Lock()
	pickerSt := pg.readyPickerWithWeights[pg.idx]
	p := pickerSt.picker
	pg.count++
	if pg.count >= pickerSt.weight {
		pg.idx = (pg.idx + 1) % pg.length
		pg.count = 0
	}
	pg.mu.Unlock()

	return p.Pick(ctx, opts)
}

// balancerGroupCC is a balancer.ClientConn implementation with the balancer ID.
type balancerGroupCC struct {
	id    string
	group *balancerGroup
}

func (bgcc *balancerGroupCC) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	return bgcc.group.newSubConn(bgcc.id, addrs, opts)
}

func (bgcc *balancerGroupCC) RemoveSubConn(sc balancer.SubConn) {
	bgcc.group.cc.RemoveSubConn(sc)
}

func (bgcc *balancerGroupCC) UpdateBalancerState(state connectivity.State, picker balancer.Picker) {
	bgcc.group.updateBalancerState(bgcc.id, state, picker)
}

func (bgcc *balancerGroupCC) ResolveNow(opt resolver.ResolveNowOption) {
	bgcc.group.cc.ResolveNow(opt)
}

func (bgcc *balancerGroupCC) Target() string {
	return bgcc.group.cc.Target()
}
