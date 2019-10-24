// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/tapds.proto

package envoy_service_tap_v2alpha

import (
	context "context"
	fmt "fmt"
	v2 "google.golang.org/grpc/xds/internal/proto/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TapResource struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config               *TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_4eef591ebf2a5317, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v2alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v2alpha/tapds.proto", fileDescriptor_4eef591ebf2a5317)
}

var fileDescriptor_4eef591ebf2a5317 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x6e, 0xe2, 0x40,
	0x10, 0xc6, 0x6f, 0x2d, 0xc4, 0xe9, 0x96, 0x82, 0x93, 0xaf, 0x38, 0xf0, 0x21, 0x0e, 0x71, 0xdc,
	0x1d, 0xba, 0x62, 0x7d, 0x72, 0x3a, 0x94, 0xca, 0x41, 0xa9, 0x91, 0x71, 0x93, 0x2a, 0x1a, 0xcc,
	0x06, 0x56, 0xb2, 0x77, 0x36, 0xde, 0xc5, 0x82, 0x36, 0x55, 0xfa, 0xbc, 0x57, 0x9a, 0xbc, 0x42,
	0x9e, 0x22, 0x55, 0xe4, 0x3f, 0xa0, 0x90, 0xc8, 0xa9, 0xd2, 0x8d, 0x3c, 0xdf, 0xf7, 0x1b, 0xcf,
	0x7e, 0x43, 0x7f, 0x73, 0x99, 0xe1, 0xce, 0xd5, 0x3c, 0xcd, 0x44, 0xc4, 0x5d, 0x03, 0xca, 0xcd,
	0x3c, 0x88, 0xd5, 0x1a, 0xf2, 0x7a, 0xa9, 0x99, 0x4a, 0xd1, 0xa0, 0xdd, 0x2d, 0x64, 0xac, 0x92,
	0x31, 0x03, 0x8a, 0x55, 0x32, 0xa7, 0x57, 0x12, 0x40, 0x09, 0x37, 0xf3, 0xdc, 0xa5, 0xd0, 0x11,
	0x66, 0x3c, 0xdd, 0x95, 0x46, 0xe7, 0x4f, 0x3d, 0x3f, 0xc2, 0x24, 0x41, 0x59, 0xe9, 0x7a, 0x2b,
	0xc4, 0x55, 0xcc, 0x0b, 0x0c, 0x48, 0x89, 0x06, 0x8c, 0x40, 0x59, 0x8d, 0x77, 0xbe, 0x67, 0x10,
	0x8b, 0x25, 0x18, 0xee, 0xee, 0x8b, 0xb2, 0x31, 0x5c, 0xd3, 0x56, 0x08, 0x2a, 0xe0, 0x1a, 0x37,
	0x69, 0xc4, 0xed, 0x1f, 0xb4, 0x21, 0x21, 0xe1, 0x1d, 0x32, 0x20, 0xe3, 0x2f, 0xfe, 0xe7, 0x27,
	0xbf, 0x91, 0x5a, 0x03, 0x12, 0x14, 0x1f, 0xed, 0x53, 0xda, 0x8c, 0x50, 0x5e, 0x89, 0x55, 0xc7,
	0x1a, 0x90, 0x71, 0xcb, 0x1b, 0xb1, 0xda, 0xa5, 0x58, 0x08, 0xea, 0xac, 0xd0, 0x06, 0x95, 0xc7,
	0xbb, 0xb7, 0xe8, 0xb7, 0x10, 0xd4, 0x74, 0xbf, 0xdf, 0xbc, 0x74, 0xd9, 0x17, 0xf4, 0xeb, 0xdc,
	0xa4, 0x1c, 0x92, 0x83, 0x45, 0xdb, 0xfd, 0x8a, 0x0c, 0x4a, 0xb0, 0xcc, 0x63, 0x07, 0x4f, 0xc0,
	0xaf, 0x37, 0x5c, 0x1b, 0xe7, 0x67, 0x6d, 0x5f, 0x2b, 0x94, 0x9a, 0x0f, 0x3f, 0x8d, 0xc9, 0x7f,
	0x62, 0x2f, 0x68, 0x7b, 0xca, 0x63, 0x03, 0x2f, 0xc8, 0xbf, 0x5e, 0x39, 0xf3, 0xf6, 0x1b, 0xfc,
	0xe8, 0x7d, 0xd1, 0xd1, 0x8c, 0x2d, 0x6d, 0x9f, 0x73, 0x13, 0xad, 0x3f, 0xf2, 0xef, 0x47, 0x37,
	0x0f, 0x8f, 0x77, 0x56, 0x7f, 0xd8, 0x3d, 0x3a, 0x88, 0x89, 0x01, 0x75, 0x59, 0x3e, 0xa6, 0x9e,
	0x90, 0x7f, 0xfe, 0x84, 0xfe, 0x15, 0x58, 0xa2, 0x54, 0x8a, 0xdb, 0x5d, 0x7d, 0x1a, 0x3e, 0x0d,
	0xf3, 0x53, 0x9c, 0xe5, 0x89, 0xcf, 0xc8, 0x2d, 0x21, 0x8b, 0x66, 0x91, 0xfe, 0xc9, 0x73, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xca, 0x3c, 0xff, 0x65, 0xbe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v2alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v2alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(srv TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(srv TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(ctx context.Context, req *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v2alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v2alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v2alpha/tapds.proto",
}
