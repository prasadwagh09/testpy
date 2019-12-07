// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// EnableRequest defines the fields in a /Profiling/Enable method request to
// toggle profiling on and off within a gRPC program.
type EnableRequest struct {
	// Setting this to true will enable profiling. Setting this to false will
	// disable profiling.
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnableRequest) Reset()         { *m = EnableRequest{} }
func (m *EnableRequest) String() string { return proto.CompactTextString(m) }
func (*EnableRequest) ProtoMessage()    {}
func (*EnableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *EnableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableRequest.Unmarshal(m, b)
}
func (m *EnableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableRequest.Marshal(b, m, deterministic)
}
func (m *EnableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableRequest.Merge(m, src)
}
func (m *EnableRequest) XXX_Size() int {
	return xxx_messageInfo_EnableRequest.Size(m)
}
func (m *EnableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnableRequest proto.InternalMessageInfo

func (m *EnableRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// EnableResponse defines the fields in a /Profiling/Enable method response.
type EnableResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnableResponse) Reset()         { *m = EnableResponse{} }
func (m *EnableResponse) String() string { return proto.CompactTextString(m) }
func (*EnableResponse) ProtoMessage()    {}
func (*EnableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *EnableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableResponse.Unmarshal(m, b)
}
func (m *EnableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableResponse.Marshal(b, m, deterministic)
}
func (m *EnableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableResponse.Merge(m, src)
}
func (m *EnableResponse) XXX_Size() int {
	return xxx_messageInfo_EnableResponse.Size(m)
}
func (m *EnableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnableResponse proto.InternalMessageInfo

// GetStreamStats defines the fields in a /Profiling/GetStreamStats method
// request to retrieve stream-level stats in a gRPC client/server.
type GetStreamStatsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStreamStatsRequest) Reset()         { *m = GetStreamStatsRequest{} }
func (m *GetStreamStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStreamStatsRequest) ProtoMessage()    {}
func (*GetStreamStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetStreamStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStreamStatsRequest.Unmarshal(m, b)
}
func (m *GetStreamStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStreamStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStreamStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStreamStatsRequest.Merge(m, src)
}
func (m *GetStreamStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStreamStatsRequest.Size(m)
}
func (m *GetStreamStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStreamStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStreamStatsRequest proto.InternalMessageInfo

// A TimerProto measures the start and end of execution of a component within
// gRPC that's being profiled. It includes a tag and some additional metadata
// to identify itself.
type TimerProto struct {
	// tags is a comma-separated list of strings used to tag a timer.
	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	// begin_sec and begin_nsec are the start epoch second and nanosecond,
	// respectively, of the component profiled by this timer in UTC. begin_nsec
	// must be a non-negative integer.
	BeginSec  int64 `protobuf:"varint,2,opt,name=begin_sec,json=beginSec,proto3" json:"begin_sec,omitempty"`
	BeginNsec int32 `protobuf:"varint,3,opt,name=begin_nsec,json=beginNsec,proto3" json:"begin_nsec,omitempty"`
	// end_sec and end_nsec are the end epoch second and nanosecond,
	// respectively, of the component profiled by this timer in UTC. end_nsec
	// must be a non-negative integer.
	EndSec  int64 `protobuf:"varint,4,opt,name=end_sec,json=endSec,proto3" json:"end_sec,omitempty"`
	EndNsec int32 `protobuf:"varint,5,opt,name=end_nsec,json=endNsec,proto3" json:"end_nsec,omitempty"`
	// go_id is the goroutine ID of the component being profiled.
	GoId                 int64    `protobuf:"varint,6,opt,name=go_id,json=goId,proto3" json:"go_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimerProto) Reset()         { *m = TimerProto{} }
func (m *TimerProto) String() string { return proto.CompactTextString(m) }
func (*TimerProto) ProtoMessage()    {}
func (*TimerProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *TimerProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimerProto.Unmarshal(m, b)
}
func (m *TimerProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimerProto.Marshal(b, m, deterministic)
}
func (m *TimerProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimerProto.Merge(m, src)
}
func (m *TimerProto) XXX_Size() int {
	return xxx_messageInfo_TimerProto.Size(m)
}
func (m *TimerProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TimerProto.DiscardUnknown(m)
}

var xxx_messageInfo_TimerProto proto.InternalMessageInfo

func (m *TimerProto) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *TimerProto) GetBeginSec() int64 {
	if m != nil {
		return m.BeginSec
	}
	return 0
}

func (m *TimerProto) GetBeginNsec() int32 {
	if m != nil {
		return m.BeginNsec
	}
	return 0
}

func (m *TimerProto) GetEndSec() int64 {
	if m != nil {
		return m.EndSec
	}
	return 0
}

func (m *TimerProto) GetEndNsec() int32 {
	if m != nil {
		return m.EndNsec
	}
	return 0
}

func (m *TimerProto) GetGoId() int64 {
	if m != nil {
		return m.GoId
	}
	return 0
}

// A StatProto is a collection of TimerProtos along with some additional
// metadata to tag and identify itself.
type StatProto struct {
	// tags is a comma-separated list of strings used to categorize a stat.
	Tags string `protobuf:"bytes,1,opt,name=tags,proto3" json:"tags,omitempty"`
	// timer_protos is an array of TimerProtos, each representing a different
	// (but possibly overlapping) component within this stat.
	TimerProtos []*TimerProto `protobuf:"bytes,2,rep,name=timer_protos,json=timerProtos,proto3" json:"timer_protos,omitempty"`
	// metadata is an array of bytes used to uniquely identify a stat with an
	// undefined encoding format. For example, the StatProtos returned by the
	// /Profiling/GetStreamStats service use the metadata field to encode the
	// connection ID and the stream ID of each query.
	Metadata             []byte   `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatProto) Reset()         { *m = StatProto{} }
func (m *StatProto) String() string { return proto.CompactTextString(m) }
func (*StatProto) ProtoMessage()    {}
func (*StatProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *StatProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatProto.Unmarshal(m, b)
}
func (m *StatProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatProto.Marshal(b, m, deterministic)
}
func (m *StatProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatProto.Merge(m, src)
}
func (m *StatProto) XXX_Size() int {
	return xxx_messageInfo_StatProto.Size(m)
}
func (m *StatProto) XXX_DiscardUnknown() {
	xxx_messageInfo_StatProto.DiscardUnknown(m)
}

var xxx_messageInfo_StatProto proto.InternalMessageInfo

func (m *StatProto) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *StatProto) GetTimerProtos() []*TimerProto {
	if m != nil {
		return m.TimerProtos
	}
	return nil
}

func (m *StatProto) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*EnableRequest)(nil), "grpc.go.profiling.v1alpha.EnableRequest")
	proto.RegisterType((*EnableResponse)(nil), "grpc.go.profiling.v1alpha.EnableResponse")
	proto.RegisterType((*GetStreamStatsRequest)(nil), "grpc.go.profiling.v1alpha.GetStreamStatsRequest")
	proto.RegisterType((*TimerProto)(nil), "grpc.go.profiling.v1alpha.TimerProto")
	proto.RegisterType((*StatProto)(nil), "grpc.go.profiling.v1alpha.StatProto")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdd, 0x4a, 0xeb, 0x40,
	0x10, 0xc7, 0xd9, 0x7e, 0xe4, 0x24, 0xd3, 0x0f, 0x0e, 0x7b, 0x38, 0x34, 0xed, 0xe1, 0x40, 0x08,
	0x2a, 0xe9, 0x4d, 0x5a, 0xeb, 0x1b, 0x08, 0xa2, 0xde, 0x48, 0x49, 0xbd, 0x12, 0xa4, 0x6c, 0x93,
	0x71, 0x8d, 0xa4, 0xd9, 0x98, 0x5d, 0xfb, 0x06, 0x3e, 0x8a, 0xaf, 0xe4, 0xf3, 0xc8, 0x6e, 0x6c,
	0xa4, 0x60, 0x8b, 0x57, 0xbb, 0xb3, 0x33, 0xbf, 0xd9, 0x99, 0xff, 0x1f, 0x7a, 0x12, 0xcb, 0x4d,
	0x1a, 0x63, 0x58, 0x94, 0x42, 0x09, 0x3a, 0xe4, 0x65, 0x11, 0x87, 0x5c, 0xe8, 0xf0, 0x21, 0xcd,
	0xd2, 0x9c, 0x87, 0x9b, 0x53, 0x96, 0x15, 0x8f, 0xcc, 0x1f, 0x43, 0xef, 0x22, 0x67, 0xab, 0x0c,
	0x23, 0x7c, 0x7e, 0x41, 0xa9, 0xa8, 0x0b, 0xbf, 0xd0, 0x3c, 0x24, 0x2e, 0xf1, 0x48, 0x60, 0x47,
	0xdb, 0xd0, 0xff, 0x0d, 0xfd, 0x6d, 0xa9, 0x2c, 0x44, 0x2e, 0xd1, 0x1f, 0xc0, 0xdf, 0x4b, 0x54,
	0x0b, 0x55, 0x22, 0x5b, 0x2f, 0x14, 0x53, 0xf2, 0xb3, 0x89, 0xff, 0x46, 0x00, 0x6e, 0xd3, 0x35,
	0x96, 0x73, 0xf3, 0x3f, 0x85, 0x96, 0x62, 0x5c, 0x9a, 0x86, 0x4e, 0x64, 0xee, 0xf4, 0x1f, 0x38,
	0x2b, 0xe4, 0x69, 0xbe, 0x94, 0x18, 0xbb, 0x0d, 0x8f, 0x04, 0xcd, 0xc8, 0x36, 0x0f, 0x0b, 0x8c,
	0xe9, 0x7f, 0x80, 0x2a, 0x99, 0xeb, 0x6c, 0xd3, 0x23, 0x41, 0x3b, 0xaa, 0xca, 0x6f, 0x24, 0xc6,
	0x74, 0xa0, 0x67, 0x4c, 0x0c, 0xd9, 0x32, 0xa4, 0x85, 0x79, 0xa2, 0xb9, 0x21, 0xd8, 0x3a, 0x61,
	0xa8, 0xb6, 0xa1, 0x74, 0xa1, 0x61, 0xfe, 0x40, 0x9b, 0x8b, 0x65, 0x9a, 0xb8, 0x96, 0x21, 0x5a,
	0x5c, 0x5c, 0x27, 0xfe, 0x2b, 0x01, 0x47, 0x0f, 0xbe, 0x7f, 0xcc, 0x2b, 0xe8, 0x2a, 0xbd, 0xc8,
	0xd2, 0x28, 0x29, 0xdd, 0x86, 0xd7, 0x0c, 0x3a, 0xb3, 0xe3, 0x70, 0xaf, 0xa2, 0xe1, 0xd7, 0xde,
	0x51, 0x47, 0xd5, 0x77, 0x49, 0x47, 0x60, 0xaf, 0x51, 0xb1, 0x84, 0x29, 0x66, 0x36, 0xea, 0x46,
	0x75, 0x3c, 0x7b, 0x27, 0xe0, 0xcc, 0xb7, 0x9d, 0xe8, 0x3d, 0x58, 0x95, 0xd0, 0x34, 0x38, 0xf0,
	0xcf, 0x8e, 0x6d, 0xa3, 0xf1, 0x0f, 0x2a, 0x2b, 0xd7, 0xe8, 0x13, 0xf4, 0x77, 0x5d, 0xa3, 0xd3,
	0x03, 0xf0, 0xb7, 0x06, 0x8f, 0x8e, 0x0e, 0x10, 0xb5, 0xa0, 0x53, 0x72, 0x1e, 0xdc, 0x9d, 0x70,
	0x21, 0x78, 0x86, 0x21, 0x17, 0x19, 0xcb, 0x79, 0x28, 0x4a, 0x3e, 0xd1, 0xe8, 0xa4, 0xe6, 0x26,
	0x46, 0xd9, 0x95, 0x65, 0x8e, 0xb3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x92, 0xb6, 0xa5,
	0xbb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfilingClient is the client API for Profiling service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilingClient interface {
	// Enable allows users to toggle profiling on and off remotely.
	Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (Profiling_GetStreamStatsClient, error)
}

type profilingClient struct {
	cc *grpc.ClientConn
}

func NewProfilingClient(cc *grpc.ClientConn) ProfilingClient {
	return &profilingClient{cc}
}

func (c *profilingClient) Enable(ctx context.Context, in *EnableRequest, opts ...grpc.CallOption) (*EnableResponse, error) {
	out := new(EnableResponse)
	err := c.cc.Invoke(ctx, "/grpc.go.profiling.v1alpha.Profiling/Enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilingClient) GetStreamStats(ctx context.Context, in *GetStreamStatsRequest, opts ...grpc.CallOption) (Profiling_GetStreamStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Profiling_serviceDesc.Streams[0], "/grpc.go.profiling.v1alpha.Profiling/GetStreamStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &profilingGetStreamStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Profiling_GetStreamStatsClient interface {
	Recv() (*StatProto, error)
	grpc.ClientStream
}

type profilingGetStreamStatsClient struct {
	grpc.ClientStream
}

func (x *profilingGetStreamStatsClient) Recv() (*StatProto, error) {
	m := new(StatProto)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProfilingServer is the server API for Profiling service.
type ProfilingServer interface {
	// Enable allows users to toggle profiling on and off remotely.
	Enable(context.Context, *EnableRequest) (*EnableResponse, error)
	// GetStreamStats is used to retrieve an array of stream-level stats from a
	// gRPC client/server.
	GetStreamStats(*GetStreamStatsRequest, Profiling_GetStreamStatsServer) error
}

// UnimplementedProfilingServer can be embedded to have forward compatible implementations.
type UnimplementedProfilingServer struct {
}

func (*UnimplementedProfilingServer) Enable(ctx context.Context, req *EnableRequest) (*EnableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (*UnimplementedProfilingServer) GetStreamStats(req *GetStreamStatsRequest, srv Profiling_GetStreamStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStreamStats not implemented")
}

func RegisterProfilingServer(s *grpc.Server, srv ProfilingServer) {
	s.RegisterService(&_Profiling_serviceDesc, srv)
}

func _Profiling_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilingServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.go.profiling.v1alpha.Profiling/Enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilingServer).Enable(ctx, req.(*EnableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiling_GetStreamStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStreamStatsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProfilingServer).GetStreamStats(m, &profilingGetStreamStatsServer{stream})
}

type Profiling_GetStreamStatsServer interface {
	Send(*StatProto) error
	grpc.ServerStream
}

type profilingGetStreamStatsServer struct {
	grpc.ServerStream
}

func (x *profilingGetStreamStatsServer) Send(m *StatProto) error {
	return x.ServerStream.SendMsg(m)
}

var _Profiling_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.go.profiling.v1alpha.Profiling",
	HandlerType: (*ProfilingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enable",
			Handler:    _Profiling_Enable_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStreamStats",
			Handler:       _Profiling_GetStreamStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
