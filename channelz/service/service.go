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

//go:generate protoc -I ../service_proto --go_out=plugins=grpc:../service_proto ../service_proto/service.proto

// Package service provides an implementation for channelz service server.
package service

import (
	"net"
	"time"

	"github.com/golang/protobuf/ptypes"
	wrpb "github.com/golang/protobuf/ptypes/wrappers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz"
	pb "google.golang.org/grpc/channelz/service_proto"
	"google.golang.org/grpc/connectivity"
)

const (
	secToNano  = 1e9
	usecToNano = 1e3
)

// RegisterChannelzServiceToServer registers the channelz service to the given server.
func RegisterChannelzServiceToServer(s *grpc.Server) {
	pb.RegisterChannelzServer(s, &serverImpl{})
}

func newCZServer() pb.ChannelzServer {
	return &serverImpl{}
}

type serverImpl struct{}

func connectivityStateToProto(s connectivity.State) *pb.ChannelConnectivityState {
	switch s {
	case connectivity.Idle:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_IDLE}
	case connectivity.Connecting:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_CONNECTING}
	case connectivity.Ready:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_READY}
	case connectivity.TransientFailure:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_TRANSIENT_FAILURE}
	case connectivity.Shutdown:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_SHUTDOWN}
	default:
		return &pb.ChannelConnectivityState{State: pb.ChannelConnectivityState_UNKNOWN}
	}
}

func channelMetricToProto(cm *channelz.ChannelMetric) *pb.Channel {
	c := &pb.Channel{}
	c.Ref = &pb.ChannelRef{ChannelId: cm.ID, Name: cm.RefName}

	c.Data = &pb.ChannelData{
		State:          connectivityStateToProto(cm.ChannelData.State),
		Target:         cm.ChannelData.Target,
		CallsStarted:   cm.ChannelData.CallsStarted,
		CallsSucceeded: cm.ChannelData.CallsSucceeded,
		CallsFailed:    cm.ChannelData.CallsFailed,
	}
	if ts, err := ptypes.TimestampProto(cm.ChannelData.LastCallStartedTimestamp); err == nil {
		c.Data.LastCallStartedTimestamp = ts
	}
	nestedChans := make([]*pb.ChannelRef, 0, len(cm.NestedChans))
	for id, ref := range cm.NestedChans {
		nestedChans = append(nestedChans, &pb.ChannelRef{ChannelId: id, Name: ref})
	}
	c.ChannelRef = nestedChans

	subChans := make([]*pb.SubchannelRef, 0, len(cm.SubChans))
	for id, ref := range cm.SubChans {
		subChans = append(subChans, &pb.SubchannelRef{SubchannelId: id, Name: ref})
	}
	c.SubchannelRef = subChans

	sockets := make([]*pb.SocketRef, 0, len(cm.Sockets))
	for id, ref := range cm.Sockets {
		sockets = append(sockets, &pb.SocketRef{SocketId: id, Name: ref})
	}
	c.SocketRef = sockets
	return c
}

func subChannelMetricToProto(cm *channelz.SubChannelMetric) *pb.Subchannel {
	sc := &pb.Subchannel{}
	sc.Ref = &pb.SubchannelRef{SubchannelId: cm.ID, Name: cm.RefName}

	sc.Data = &pb.ChannelData{
		State:          connectivityStateToProto(cm.ChannelData.State),
		Target:         cm.ChannelData.Target,
		CallsStarted:   cm.ChannelData.CallsStarted,
		CallsSucceeded: cm.ChannelData.CallsSucceeded,
		CallsFailed:    cm.ChannelData.CallsFailed,
	}
	if ts, err := ptypes.TimestampProto(cm.ChannelData.LastCallStartedTimestamp); err == nil {
		sc.Data.LastCallStartedTimestamp = ts
	}
	nestedChans := make([]*pb.ChannelRef, 0, len(cm.NestedChans))
	for id, ref := range cm.NestedChans {
		nestedChans = append(nestedChans, &pb.ChannelRef{ChannelId: id, Name: ref})
	}
	sc.ChannelRef = nestedChans

	subChans := make([]*pb.SubchannelRef, 0, len(cm.SubChans))
	for id, ref := range cm.SubChans {
		subChans = append(subChans, &pb.SubchannelRef{SubchannelId: id, Name: ref})
	}
	sc.SubchannelRef = subChans

	sockets := make([]*pb.SocketRef, 0, len(cm.Sockets))
	for id, ref := range cm.Sockets {
		sockets = append(sockets, &pb.SocketRef{SocketId: id, Name: ref})
	}
	sc.SocketRef = sockets
	return sc
}

func securityToProto(se channelz.SecurityValue) *pb.Security {
	switch v := se.(type) {
	case *channelz.TLSSecurityValue:
		return &pb.Security{
			&pb.Security_Tls_{
				&pb.Security_Tls{
					CipherSuite:       &pb.Security_Tls_StandardName{v.StandardName},
					LocalCertificate:  v.LocalCertificate,
					RemoteCertificate: v.RemoteCertificate,
				},
			},
		}
	case *channelz.OtherSecurityValue:
		anyval, err := ptypes.MarshalAny(v.Value)
		if err != nil {
			return &pb.Security{
				&pb.Security_Other{
					&pb.Security_OtherSecurity{
						Name: v.Name,
					},
				},
			}
		}
		return &pb.Security{
			&pb.Security_Other{
				&pb.Security_OtherSecurity{
					Name:  v.Name,
					Value: anyval,
				},
			},
		}
	}
	return nil
}

func sockoptToProto(skopts *channelz.SocketOptionData) []*pb.SocketOption {
	var opts []*pb.SocketOption
	if skopts.Linger != nil {
		additional, err := ptypes.MarshalAny(&pb.SocketOptionLinger{Active: skopts.Linger.Onoff != 0, Duration: ptypes.DurationProto(time.Duration(int64(skopts.Linger.Linger) * secToNano))})
		if err == nil {
			opts = append(opts, &pb.SocketOption{Name: "SO_LINGER", Additional: additional})
		}
	}
	if skopts.RecvTimeout != nil {
		additional, err := ptypes.MarshalAny(&pb.SocketOptionTimeout{ptypes.DurationProto(time.Duration(skopts.RecvTimeout.Sec*secToNano + skopts.RecvTimeout.Usec*usecToNano))})
		if err == nil {
			opts = append(opts, &pb.SocketOption{Name: "SO_RCVTIMEO", Additional: additional})
		}
	}
	if skopts.SendTimeout != nil {
		additional, err := ptypes.MarshalAny(&pb.SocketOptionTimeout{ptypes.DurationProto(time.Duration(skopts.SendTimeout.Sec*secToNano + skopts.SendTimeout.Usec*usecToNano))})
		if err == nil {
			opts = append(opts, &pb.SocketOption{Name: "SO_SNDTIMEO", Additional: additional})
		}
	}
	if skopts.TCPInfo != nil {
		additional, err := ptypes.MarshalAny(&pb.SocketOptionTcpInfo{
			TcpiState:       uint32(skopts.TCPInfo.State),
			TcpiCaState:     uint32(skopts.TCPInfo.Ca_state),
			TcpiRetransmits: uint32(skopts.TCPInfo.Retransmits),
			TcpiProbes:      uint32(skopts.TCPInfo.Probes),
			TcpiBackoff:     uint32(skopts.TCPInfo.Backoff),
			TcpiOptions:     uint32(skopts.TCPInfo.Options),
			// https://golang.org/pkg/syscall/#TCPInfo
			// TCPInfo struct does not contain info about TcpiSndWscale and TcpiRcvWscale.
			TcpiRto:          skopts.TCPInfo.Rto,
			TcpiAto:          skopts.TCPInfo.Ato,
			TcpiSndMss:       skopts.TCPInfo.Snd_mss,
			TcpiRcvMss:       skopts.TCPInfo.Rcv_mss,
			TcpiUnacked:      skopts.TCPInfo.Unacked,
			TcpiSacked:       skopts.TCPInfo.Sacked,
			TcpiLost:         skopts.TCPInfo.Lost,
			TcpiRetrans:      skopts.TCPInfo.Retrans,
			TcpiFackets:      skopts.TCPInfo.Fackets,
			TcpiLastDataSent: skopts.TCPInfo.Last_data_sent,
			TcpiLastAckSent:  skopts.TCPInfo.Last_ack_sent,
			TcpiLastDataRecv: skopts.TCPInfo.Last_data_recv,
			TcpiLastAckRecv:  skopts.TCPInfo.Last_ack_recv,
			TcpiPmtu:         skopts.TCPInfo.Pmtu,
			TcpiRcvSsthresh:  skopts.TCPInfo.Rcv_ssthresh,
			TcpiRtt:          skopts.TCPInfo.Rtt,
			TcpiRttvar:       skopts.TCPInfo.Rttvar,
			TcpiSndSsthresh:  skopts.TCPInfo.Snd_ssthresh,
			TcpiSndCwnd:      skopts.TCPInfo.Snd_cwnd,
			TcpiAdvmss:       skopts.TCPInfo.Advmss,
			TcpiReordering:   skopts.TCPInfo.Reordering,
		})
		if err == nil {
			opts = append(opts, &pb.SocketOption{Name: "TCP_INFO", Additional: additional})
		}
	}
	return opts
}

func addrToProto(a net.Addr) *pb.Address {
	switch a.Network() {
	case "udp":
		// TODO: Address_OtherAddress{}. Need proto def for Value.
	case "ip":
		// Note zone info is discarded through the conversion.
		return &pb.Address{Address: &pb.Address_TcpipAddress{TcpipAddress: &pb.Address_TcpIpAddress{IpAddress: a.(*net.IPAddr).IP}}}
	case "ip+net":
		// Note mask info is discarded through the conversion.
		return &pb.Address{Address: &pb.Address_TcpipAddress{TcpipAddress: &pb.Address_TcpIpAddress{IpAddress: a.(*net.IPNet).IP}}}
	case "tcp":
		// Note zone info is discarded through the conversion.
		return &pb.Address{Address: &pb.Address_TcpipAddress{TcpipAddress: &pb.Address_TcpIpAddress{IpAddress: a.(*net.TCPAddr).IP, Port: int32(a.(*net.TCPAddr).Port)}}}
	case "unix", "unixgram", "unixpacket":
		return &pb.Address{Address: &pb.Address_UdsAddress_{UdsAddress: &pb.Address_UdsAddress{Filename: a.String()}}}
	default:
	}
	return &pb.Address{}
}

func socketMetricToProto(sm *channelz.SocketMetric) *pb.Socket {
	s := &pb.Socket{}
	s.Ref = &pb.SocketRef{SocketId: sm.ID, Name: sm.RefName}

	s.Data = &pb.SocketData{
		StreamsStarted:   sm.SocketData.StreamsStarted,
		StreamsSucceeded: sm.SocketData.StreamsSucceeded,
		StreamsFailed:    sm.SocketData.StreamsFailed,
		MessagesSent:     sm.SocketData.MessagesSent,
		MessagesReceived: sm.SocketData.MessagesReceived,
		KeepAlivesSent:   sm.SocketData.KeepAlivesSent,
	}
	if ts, err := ptypes.TimestampProto(sm.SocketData.LastLocalStreamCreatedTimestamp); err == nil {
		s.Data.LastLocalStreamCreatedTimestamp = ts
	}
	if ts, err := ptypes.TimestampProto(sm.SocketData.LastRemoteStreamCreatedTimestamp); err == nil {
		s.Data.LastRemoteStreamCreatedTimestamp = ts
	}
	if ts, err := ptypes.TimestampProto(sm.SocketData.LastMessageSentTimestamp); err == nil {
		s.Data.LastMessageSentTimestamp = ts
	}
	if ts, err := ptypes.TimestampProto(sm.SocketData.LastMessageReceivedTimestamp); err == nil {
		s.Data.LastMessageReceivedTimestamp = ts
	}
	s.Data.LocalFlowControlWindow = &wrpb.Int64Value{Value: sm.SocketData.LocalFlowControlWindow}
	s.Data.RemoteFlowControlWindow = &wrpb.Int64Value{Value: sm.SocketData.RemoteFlowControlWindow}

	if sm.SocketData.SocketOptions != nil {
		s.Data.Option = sockoptToProto(sm.SocketData.SocketOptions)
	}
	if sm.SocketData.Security != nil {
		s.Security = securityToProto(sm.SocketData.Security)
	}

	if sm.SocketData.LocalAddr != nil {
		s.Local = addrToProto(sm.SocketData.LocalAddr)
	}
	if sm.SocketData.RemoteAddr != nil {
		s.Remote = addrToProto(sm.SocketData.RemoteAddr)
	}
	s.RemoteName = sm.SocketData.RemoteName
	return s
}

func (s *serverImpl) GetTopChannels(ctx context.Context, req *pb.GetTopChannelsRequest) (*pb.GetTopChannelsResponse, error) {
	metrics, end := channelz.GetTopChannels(req.GetStartChannelId())
	resp := &pb.GetTopChannelsResponse{}
	for _, m := range metrics {
		resp.Channel = append(resp.Channel, channelMetricToProto(m))
	}
	resp.End = end
	return resp, nil
}

func serverMetricToProto(sm *channelz.ServerMetric) *pb.Server {
	s := &pb.Server{}
	s.Ref = &pb.ServerRef{ServerId: sm.ID, Name: sm.RefName}

	s.Data = &pb.ServerData{
		CallsStarted:   sm.ServerData.CallsStarted,
		CallsSucceeded: sm.ServerData.CallsSucceeded,
		CallsFailed:    sm.ServerData.CallsFailed,
	}

	if ts, err := ptypes.TimestampProto(sm.ServerData.LastCallStartedTimestamp); err == nil {
		s.Data.LastCallStartedTimestamp = ts
	}
	sockets := make([]*pb.SocketRef, 0, len(sm.ListenSockets))
	for id, ref := range sm.ListenSockets {
		sockets = append(sockets, &pb.SocketRef{SocketId: id, Name: ref})
	}
	s.ListenSocket = sockets
	return s
}

func (s *serverImpl) GetServers(ctx context.Context, req *pb.GetServersRequest) (*pb.GetServersResponse, error) {
	metrics, end := channelz.GetServers(req.GetStartServerId())
	resp := &pb.GetServersResponse{}
	for _, m := range metrics {
		resp.Server = append(resp.Server, serverMetricToProto(m))
	}
	resp.End = end
	return resp, nil
}

func (s *serverImpl) GetServerSockets(ctx context.Context, req *pb.GetServerSocketsRequest) (*pb.GetServerSocketsResponse, error) {
	metrics, end := channelz.GetServerSockets(req.GetServerId(), req.GetStartSocketId())
	resp := &pb.GetServerSocketsResponse{}
	for _, m := range metrics {
		resp.SocketRef = append(resp.SocketRef, &pb.SocketRef{SocketId: m.ID, Name: m.RefName})
	}
	resp.End = end
	return resp, nil
}

func (s *serverImpl) GetChannel(ctx context.Context, req *pb.GetChannelRequest) (*pb.GetChannelResponse, error) {
	var metric *channelz.ChannelMetric
	if metric = channelz.GetChannel(req.GetChannelId()); metric == nil {
		return &pb.GetChannelResponse{}, nil
	}
	resp := &pb.GetChannelResponse{Channel: channelMetricToProto(metric)}
	return resp, nil
}

func (s *serverImpl) GetSubchannel(ctx context.Context, req *pb.GetSubchannelRequest) (*pb.GetSubchannelResponse, error) {
	var metric *channelz.SubChannelMetric
	if metric = channelz.GetSubChannel(req.GetSubchannelId()); metric == nil {
		return &pb.GetSubchannelResponse{}, nil
	}
	resp := &pb.GetSubchannelResponse{Subchannel: subChannelMetricToProto(metric)}
	return resp, nil
}

func (s *serverImpl) GetSocket(ctx context.Context, req *pb.GetSocketRequest) (*pb.GetSocketResponse, error) {
	var metric *channelz.SocketMetric
	if metric = channelz.GetSocket(req.GetSocketId()); metric == nil {
		return &pb.GetSocketResponse{}, nil
	}
	resp := &pb.GetSocketResponse{Socket: socketMetricToProto(metric)}
	return resp, nil
}
