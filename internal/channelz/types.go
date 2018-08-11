/*
 *
 * Copyright 2018 gRPC authors.
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

package channelz

import (
	"net"
	"time"

	"sync"

	"sync/atomic"

	"fmt"

	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// entry represents a node in the channelz database.
type entry interface {
	// addChild adds a child e, whose channelz id is id to child list
	addChild(id int64, e entry)
	// deleteChild deletes a child with channelz id to be id from child list
	deleteChild(id int64)
	// triggerDelete tries to delete self from channelz database. However, if child
	// list is not empty, then deletion from the database is on hold until the last
	// child is deleted from database.
	triggerDelete()
	// deleteSelfIfReady check whether triggerDelete() has been called before, and whether child
	// list is now empty. If both conditions are met, then delete self from database.
	deleteSelfIfReady()
}

// dummyEntry is a fake entry to handle entry not found case.
type dummyEntry struct {
	idNotFound int64
}

func (d *dummyEntry) addChild(id int64, e entry) {
	// Note: It is possible for a normal program to reach here under race condition.
	// For example, there could be a race between ClientConn.Close() info being propagated
	// to addrConn and http2Client. ClientConn.Close() cancel the context and result
	// in http2Client to error. The error info is then caught by transport monitor
	// and before addrConn.tearDown() is called in side ClientConn.Close(). Therefore,
	// the addrConn will create a new transport. And when registering the new transport in
	// channelz, its parent addrConn could have already been torn down and deleted
	// from channelz tracking, and thus reach the code here.
	grpclog.Infof("attempt to add child of type %T with id %d to a parent (id=%d) that doesn't currently exist", e, id, d.idNotFound)
}

func (d *dummyEntry) deleteChild(id int64) {
	// It is possible for a normal program to reach here under race condition.
	// Refer to the example described in addChild().
	grpclog.Infof("attempt to delete child with id %d from a parent (id=%d) that doesn't currently exist", id, d.idNotFound)
}

func (d *dummyEntry) triggerDelete() {
	grpclog.Warningf("attempt to delete an entry (id=%d) that doesn't currently exist", d.idNotFound)
}

func (*dummyEntry) deleteSelfIfReady() {
	// code should not reach here. deleteSelfIfReady is always called on an existing entry.
}

// ChannelMetric defines the info channelz provides for a specific Channel, which
// includes ChannelInternalMetric and channelz-specific data, such as channelz id,
// child list, etc.
type ChannelMetric struct {
	// ID is the channelz id of this channel.
	ID int64
	// RefName is the human readable reference string of this channel.
	RefName string
	// ChannelData contains channel internal metric reported by the channel through
	// ChannelzMetric().
	ChannelData *ChannelInternalMetric
	// NestedChans tracks the nested channel type children of this channel in the format of
	// a map from nested channel channelz id to corresponding reference string.
	NestedChans map[int64]string
	// SubChans tracks the subchannel type children of this channel in the format of a
	// map from subchannel channelz id to corresponding reference string.
	SubChans map[int64]string
	// Sockets tracks the socket type children of this channel in the format of a map
	// from socket channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow channel having sockets directly,
	// therefore, this is field is unused.
	Sockets map[int64]string
	// Trace contains the most recent traced events.
	Trace *ChannelTrace
}

// SubChannelMetric defines the info channelz provides for a specific SubChannel,
// which includes ChannelInternalMetric and channelz-specific data, such as
// channelz id, child list, etc.
type SubChannelMetric struct {
	// ID is the channelz id of this subchannel.
	ID int64
	// RefName is the human readable reference string of this subchannel.
	RefName string
	// ChannelData contains subchannel internal metric reported by the subchannel
	// through ChannelzMetric().
	ChannelData *ChannelInternalMetric
	// NestedChans tracks the nested channel type children of this subchannel in the format of
	// a map from nested channel channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow subchannel to have nested channels
	// as children, therefore, this field is unused.
	NestedChans map[int64]string
	// SubChans tracks the subchannel type children of this subchannel in the format of a
	// map from subchannel channelz id to corresponding reference string.
	// Note current grpc implementation doesn't allow subchannel to have subchannels
	// as children, therefore, this field is unused.
	SubChans map[int64]string
	// Sockets tracks the socket type children of this subchannel in the format of a map
	// from socket channelz id to corresponding reference string.
	Sockets map[int64]string
	// Trace contains the most recent traced events.
	Trace *ChannelTrace
}

// ChannelInternalMetric defines the struct that the implementor of Channel interface
// should return from ChannelzMetric().
type ChannelInternalMetric struct {
	// current connectivity state of the channel.
	State connectivity.State
	// The target this channel originally tried to connect to.  May be absent
	Target string
	// The number of calls started on the channel.
	CallsStarted int64
	// The number of calls that have completed with an OK status.
	CallsSucceeded int64
	// The number of calls that have a completed with a non-OK status.
	CallsFailed int64
	// The last time a call was started on the channel.
	LastCallStartedTimestamp time.Time
}

type ChannelTrace struct {
	EventNum     int64
	CreationTime time.Time
	Events       []*TraceEvent
}

type TraceEvent struct {
	Desc         string
	Severity     Severity
	Timestamp    time.Time
	ID           int64
	RefName      string
	IsRefChannel bool
}

// Channel is the interface that should be satisfied in order to be tracked by
// channelz as Channel or SubChannel.
type Channel interface {
	ChannelzMetric() *ChannelInternalMetric
}

type dummyChannel struct{}

func (d *dummyChannel) ChannelzMetric() *ChannelInternalMetric {
	return &ChannelInternalMetric{}
}

type channel struct {
	refName       string
	c             Channel
	closeCalled   bool
	nestedChans   map[int64]string
	subChans      map[int64]string
	id            int64
	pid           int64
	cm            *channelMap
	trace         *channelTrace
	traceRefCount int32
}

func (c *channel) addChild(id int64, e entry) {
	switch v := e.(type) {
	case *subChannel:
		c.subChans[id] = v.refName
	case *channel:
		c.nestedChans[id] = v.refName
	default:
		grpclog.Errorf("cannot add a child (id = %d) of type %T to a channel", id, e)
	}
}

func (c *channel) deleteChild(id int64) {
	delete(c.subChans, id)
	delete(c.nestedChans, id)
	c.deleteSelfIfReady()
}

func (c *channel) triggerDelete() {
	c.cm.traceChannelDeleted(c.pid, c.id)
	c.closeCalled = true
	c.deleteSelfIfReady()
}

func (c *channel) deleteSelfIfReady() {
	if !c.closeCalled || len(c.subChans)+len(c.nestedChans) != 0 {
		return
	}
	if c.getTraceCount() != 0 {
		// free the grpc struct (i.e. ChannelzChannel(wrapped ClientConn))
		c.c = &dummyChannel{}
	} else {
		c.cm.deleteEntry(c.id)
	}
	// not top channel
	if c.pid != 0 {
		c.cm.findEntry(c.pid).deleteChild(c.id)
	}
}

func (c *channel) getChannelTrace() *channelTrace {
	return c.trace
}

func (c *channel) incrTraceCount() {
	atomic.AddInt32(&c.traceRefCount, 1)
}

func (c *channel) decrTraceCount() {
	atomic.AddInt32(&c.traceRefCount, -1)
}

func (c *channel) getTraceCount() int {
	i := atomic.LoadInt32(&c.traceRefCount)
	return int(i)
}

func (c *channel) cleanup() {
	if c.getTraceCount() != 0 {
		// should never gets here
		return
	}
	c.cm.deleteEntry(c.id)
	c.trace.clear()
}

func (c *channel) getRefName() string {
	return c.refName
}

type subChannel struct {
	refName       string
	c             Channel
	closeCalled   bool
	sockets       map[int64]string
	id            int64
	pid           int64
	cm            *channelMap
	trace         *channelTrace
	traceRefCount int32
}

func (sc *subChannel) addChild(id int64, e entry) {
	if v, ok := e.(*normalSocket); ok {
		sc.sockets[id] = v.refName
	} else {
		grpclog.Errorf("cannot add a child (id = %d) of type %T to a subChannel", id, e)
	}
}

func (sc *subChannel) deleteChild(id int64) {
	delete(sc.sockets, id)
	sc.deleteSelfIfReady()
}

func (sc *subChannel) triggerDelete() {
	sc.cm.traceSubChannelDeleted(sc.pid, sc.id)
	sc.closeCalled = true
	sc.deleteSelfIfReady()
}

func (sc *subChannel) deleteSelfIfReady() {
	if !sc.closeCalled || len(sc.sockets) != 0 {
		return
	}
	if sc.getTraceCount() != 0 {
		// free the grpc struct (i.e. addrConn)
		sc.c = &dummyChannel{}
	} else {
		sc.cm.deleteEntry(sc.id)
	}
	sc.cm.findEntry(sc.pid).deleteChild(sc.id)
}

func (sc *subChannel) getChannelTrace() *channelTrace {
	return sc.trace
}

func (sc *subChannel) incrTraceCount() {
	atomic.AddInt32(&sc.traceRefCount, 1)
}

func (sc *subChannel) decrTraceCount() {
	atomic.AddInt32(&sc.traceRefCount, -1)
}

func (sc *subChannel) getTraceCount() int {
	i := atomic.LoadInt32(&sc.traceRefCount)
	return int(i)
}

func (sc *subChannel) cleanup() {
	if sc.getTraceCount() != 0 {
		// should never gets here
		return
	}
	sc.cm.deleteEntry(sc.id)
	sc.trace.clear()
}

func (sc *subChannel) getRefName() string {
	return sc.refName
}

// SocketMetric defines the info channelz provides for a specific Socket, which
// includes SocketInternalMetric and channelz-specific data, such as channelz id, etc.
type SocketMetric struct {
	// ID is the channelz id of this socket.
	ID int64
	// RefName is the human readable reference string of this socket.
	RefName string
	// SocketData contains socket internal metric reported by the socket through
	// ChannelzMetric().
	SocketData *SocketInternalMetric
}

// SocketInternalMetric defines the struct that the implementor of Socket interface
// should return from ChannelzMetric().
type SocketInternalMetric struct {
	// The number of streams that have been started.
	StreamsStarted int64
	// The number of streams that have ended successfully:
	// On client side, receiving frame with eos bit set.
	// On server side, sending frame with eos bit set.
	StreamsSucceeded int64
	// The number of streams that have ended unsuccessfully:
	// On client side, termination without receiving frame with eos bit set.
	// On server side, termination without sending frame with eos bit set.
	StreamsFailed int64
	// The number of messages successfully sent on this socket.
	MessagesSent     int64
	MessagesReceived int64
	// The number of keep alives sent.  This is typically implemented with HTTP/2
	// ping messages.
	KeepAlivesSent int64
	// The last time a stream was created by this endpoint.  Usually unset for
	// servers.
	LastLocalStreamCreatedTimestamp time.Time
	// The last time a stream was created by the remote endpoint.  Usually unset
	// for clients.
	LastRemoteStreamCreatedTimestamp time.Time
	// The last time a message was sent by this endpoint.
	LastMessageSentTimestamp time.Time
	// The last time a message was received by this endpoint.
	LastMessageReceivedTimestamp time.Time
	// The amount of window, granted to the local endpoint by the remote endpoint.
	// This may be slightly out of date due to network latency.  This does NOT
	// include stream level or TCP level flow control info.
	LocalFlowControlWindow int64
	// The amount of window, granted to the remote endpoint by the local endpoint.
	// This may be slightly out of date due to network latency.  This does NOT
	// include stream level or TCP level flow control info.
	RemoteFlowControlWindow int64
	// The locally bound address.
	LocalAddr net.Addr
	// The remote bound address.  May be absent.
	RemoteAddr net.Addr
	// Optional, represents the name of the remote endpoint, if different than
	// the original target name.
	RemoteName    string
	SocketOptions *SocketOptionData
	Security      credentials.ChannelzSecurityValue
}

// Socket is the interface that should be satisfied in order to be tracked by
// channelz as Socket.
type Socket interface {
	ChannelzMetric() *SocketInternalMetric
}

type listenSocket struct {
	refName string
	s       Socket
	id      int64
	pid     int64
	cm      *channelMap
}

func (ls *listenSocket) addChild(id int64, e entry) {
	grpclog.Errorf("cannot add a child (id = %d) of type %T to a listen socket", id, e)
}

func (ls *listenSocket) deleteChild(id int64) {
	grpclog.Errorf("cannot delete a child (id = %d) from a listen socket", id)
}

func (ls *listenSocket) triggerDelete() {
	ls.cm.deleteEntry(ls.id)
	ls.cm.findEntry(ls.pid).deleteChild(ls.id)
}

func (ls *listenSocket) deleteSelfIfReady() {
	grpclog.Errorf("cannot call deleteSelfIfReady on a listen socket")
}

type normalSocket struct {
	refName string
	s       Socket
	id      int64
	pid     int64
	cm      *channelMap
}

func (ns *normalSocket) addChild(id int64, e entry) {
	grpclog.Errorf("cannot add a child (id = %d) of type %T to a normal socket", id, e)
}

func (ns *normalSocket) deleteChild(id int64) {
	grpclog.Errorf("cannot delete a child (id = %d) from a normal socket", id)
}

func (ns *normalSocket) triggerDelete() {
	ns.cm.deleteEntry(ns.id)
	ns.cm.findEntry(ns.pid).deleteChild(ns.id)
}

func (ns *normalSocket) deleteSelfIfReady() {
	grpclog.Errorf("cannot call deleteSelfIfReady on a normal socket")
}

// ServerMetric defines the info channelz provides for a specific Server, which
// includes ServerInternalMetric and channelz-specific data, such as channelz id,
// child list, etc.
type ServerMetric struct {
	// ID is the channelz id of this server.
	ID int64
	// RefName is the human readable reference string of this server.
	RefName string
	// ServerData contains server internal metric reported by the server through
	// ChannelzMetric().
	ServerData *ServerInternalMetric
	// ListenSockets tracks the listener socket type children of this server in the
	// format of a map from socket channelz id to corresponding reference string.
	ListenSockets map[int64]string
}

// ServerInternalMetric defines the struct that the implementor of Server interface
// should return from ChannelzMetric().
type ServerInternalMetric struct {
	// The number of incoming calls started on the server.
	CallsStarted int64
	// The number of incoming calls that have completed with an OK status.
	CallsSucceeded int64
	// The number of incoming calls that have a completed with a non-OK status.
	CallsFailed int64
	// The last time a call was started on the server.
	LastCallStartedTimestamp time.Time
}

// Server is the interface to be satisfied in order to be tracked by channelz as
// Server.
type Server interface {
	ChannelzMetric() *ServerInternalMetric
}

type server struct {
	refName       string
	s             Server
	closeCalled   bool
	sockets       map[int64]string
	listenSockets map[int64]string
	id            int64
	cm            *channelMap
}

func (s *server) addChild(id int64, e entry) {
	switch v := e.(type) {
	case *normalSocket:
		s.sockets[id] = v.refName
	case *listenSocket:
		s.listenSockets[id] = v.refName
	default:
		grpclog.Errorf("cannot add a child (id = %d) of type %T to a server", id, e)
	}
}

func (s *server) deleteChild(id int64) {
	delete(s.sockets, id)
	delete(s.listenSockets, id)
	s.deleteSelfIfReady()
}

func (s *server) triggerDelete() {
	s.closeCalled = true
	s.deleteSelfIfReady()
}

func (s *server) deleteSelfIfReady() {
	if !s.closeCalled || len(s.sockets)+len(s.listenSockets) != 0 {
		return
	}
	s.cm.deleteEntry(s.id)
}

type tracedChannel interface {
	getChannelTrace() *channelTrace
	incrTraceCount()
	decrTraceCount()
	cleanup()
	getRefName() string
}

type eventType int

const (
	channelCreate eventType = iota
	channelDelete
	subChannelCreate
	subChannelDelete
	channelConnectivityChange
	subChannelConnectivityChange
	subChannelPickNewAddress
	addressResolutionChange
)

type event struct {
	t         eventType
	timestamp time.Time
	refId     int64
	refName   string
	//distinguish whether the referenced entity is a channel or subchannel.
	isRefChannel      bool
	connectivityState connectivity.State
	addrEventType     AddressResolutionChangeType
	addrEventDesc     string
}

// change the format string inside this function will lead to test failures.
func (e *event) getDesc() string {
	switch e.t {
	case channelCreate:
		if e.refId != 0 {
			return fmt.Sprintf("Nested Channel (id: %d[%s]) Created", e.refId, e.refName)
		}
		return "Channel Created"
	case channelDelete:
		if e.refId != 0 {
			return fmt.Sprintf("Nested Channel (id: %d[%s]) Deleted", e.refId, e.refName)
		}
		return "Channel Deleted"
	case subChannelCreate:
		if e.refId != 0 {
			return fmt.Sprintf("SubChannel (id: %d[%s]) Created", e.refId, e.refName)
		}
		return "Subchannel Created"
	case subChannelDelete:
		if e.refId != 0 {
			return fmt.Sprintf("SubChannel (id: %d[%s]) Deleted", e.refId, e.refName)
		}
		return "Subchannel Deleted"
	case channelConnectivityChange:
		return fmt.Sprintf("Channel's connectivity state changed to %s", e.connectivityState.String())
	case subChannelConnectivityChange:
		return fmt.Sprintf("Subchannel's connectivity state changed to %s", e.connectivityState.String())
	case subChannelPickNewAddress:
		return fmt.Sprintf("Subchannel picked a new address: %q", e.addrEventDesc)
	case addressResolutionChange:
		switch e.addrEventType {
		case ServiceConfigChange:
			return fmt.Sprintf("New service config resolved: \"%s\"", e.addrEventDesc)
		case NonEmptyAddressList:
			return fmt.Sprintf("Adressses resolved (from empty address state): %q", e.addrEventDesc)
		case EmptyAddressList:
			return "Addresses resolved is empty"
		case NewLBPolicy:
			return fmt.Sprintf("New LB policy from address resolution: %q", e.addrEventDesc)
		default:
			//should not get here
		}
	default:
		//should not get here
	}
	return ""
}

func (e *event) getSeverity() Severity {
	switch e.t {
	case channelCreate:
		return CtINFO
	case channelDelete:
		return CtINFO
	case subChannelCreate:
		return CtINFO
	case subChannelDelete:
		return CtINFO
	case channelConnectivityChange:
		return CtINFO
	case subChannelConnectivityChange:
		return CtINFO
	case subChannelPickNewAddress:
		return CtINFO
	case addressResolutionChange:
		switch e.addrEventType {
		case ServiceConfigChange:
			return CtINFO
		case NonEmptyAddressList:
			return CtINFO
		case EmptyAddressList:
			return CtWarning
		case NewLBPolicy:
			return CtINFO
		default:
			//should not get here
		}
	default:
		//should not get here
	}
	return CtUNKNOWN
}

type channelTrace struct {
	cm          *channelMap
	createdTime time.Time
	mu          sync.Mutex
	events      []*event
}

func (c *channelTrace) append(e *event) {
	c.mu.Lock()
	if len(c.events) == getMaxTraceEntry() {
		del := c.events[0]
		c.events = c.events[1:]
		if del.t == channelDelete || del.t == subChannelDelete {
			// start recursive cleanup in a goroutine to not block the call originated from grpc.
			go c.cm.startCleanup(del.refId)
		}
	}
	e.timestamp = time.Now()
	c.events = append(c.events, e)
	c.mu.Unlock()
}

func (c *channelTrace) clear() {
	c.mu.Lock()
	for _, e := range c.events {
		if e.t == channelDelete || e.t == subChannelDelete {
			if v, ok := c.cm.findEntry(e.refId).(tracedChannel); ok {
				v.decrTraceCount()
				v.cleanup()
			}
		}
	}
	c.mu.Unlock()
}

// nid is the nested channel id.
func (c *channelTrace) ChannelCreated(nid int64, ref string) {
	c.append(&event{t: channelCreate, refId: nid, refName: ref, isRefChannel: true})
}

func (c *channelTrace) ChannelDeleted(nid int64, ref string) {
	c.append(&event{t: channelDelete, refId: nid, refName: ref, isRefChannel: true})
}

func (c *channelTrace) SubChannelCreated(scID int64, ref string) {
	c.append(&event{t: subChannelCreate, refId: scID, refName: ref})
}

func (c *channelTrace) SubChannelDeleted(scID int64, ref string) {
	c.append(&event{t: subChannelDelete, refId: scID, refName: ref})
}

func (c *channelTrace) ChannelConnectivityChange(s connectivity.State) {
	c.append(&event{t: channelConnectivityChange, connectivityState: s})
}

func (c *channelTrace) SubChannelConnectivityChange(s connectivity.State) {
	c.append(&event{t: subChannelConnectivityChange, connectivityState: s})
}

func (c *channelTrace) SubChannelPickNewAddress(addr string) {
	c.append(&event{t: subChannelPickNewAddress, addrEventDesc: addr})
}

type AddressResolutionChangeType int

const (
	ServiceConfigChange AddressResolutionChangeType = iota
	NonEmptyAddressList
	EmptyAddressList
	NewLBPolicy
)

func (c *channelTrace) AddressResolutionChange(t AddressResolutionChangeType, desc string) {
	c.append(&event{t: addressResolutionChange, addrEventType: t, addrEventDesc: desc})
}

type Severity int

const (
	CtUNKNOWN Severity = iota
	CtINFO
	CtWarning
	CtError
)

func (c *channelTrace) dumpData() *ChannelTrace {
	c.mu.Lock()
	ct := &ChannelTrace{EventNum: int64(len(c.events)), CreationTime: c.createdTime}
	ct.Events = make([]*TraceEvent, 0, len(c.events))
	for _, e := range c.events {
		ct.Events = append(ct.Events, &TraceEvent{
			Desc:         e.getDesc(),
			Severity:     e.getSeverity(),
			Timestamp:    e.timestamp,
			ID:           e.refId,
			RefName:      e.refName,
			IsRefChannel: e.isRefChannel,
		})
	}
	c.mu.Unlock()
	return ct
}
