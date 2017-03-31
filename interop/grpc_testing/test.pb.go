// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package grpc_testing is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Empty
	Payload
	EchoStatus
	SimpleRequest
	SimpleResponse
	StreamingInputCallRequest
	StreamingInputCallResponse
	ResponseParameters
	StreamingOutputCallRequest
	StreamingOutputCallResponse
*/
package grpc_testing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of payload that should be returned.
type PayloadType int32

const (
	// Compressable text format.
	PayloadType_COMPRESSABLE PayloadType = 0
	// Uncompressable binary format.
	PayloadType_UNCOMPRESSABLE PayloadType = 1
	// Randomly chosen from all other formats defined in this enum.
	PayloadType_RANDOM PayloadType = 2
)

var PayloadType_name = map[int32]string{
	0: "COMPRESSABLE",
	1: "UNCOMPRESSABLE",
	2: "RANDOM",
}
var PayloadType_value = map[string]int32{
	"COMPRESSABLE":   0,
	"UNCOMPRESSABLE": 1,
	"RANDOM":         2,
}

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}
func (x PayloadType) String() string {
	return proto.EnumName(PayloadType_name, int32(x))
}
func (x *PayloadType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PayloadType_value, data, "PayloadType")
	if err != nil {
		return err
	}
	*x = PayloadType(value)
	return nil
}
func (PayloadType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A block of data, to simply increase gRPC message size.
type Payload struct {
	// The type of data in body.
	Type *PayloadType `protobuf:"varint,1,opt,name=type,enum=grpc.testing.PayloadType" json:"type,omitempty"`
	// Primary contents of payload.
	Body             []byte `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Payload) GetType() PayloadType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PayloadType_COMPRESSABLE
}

func (m *Payload) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

// A protobuf representation for grpc status. This is used by test
// clients to specify a status that the server should attempt to return.
type EchoStatus struct {
	Code             *int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message          *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EchoStatus) Reset()                    { *m = EchoStatus{} }
func (m *EchoStatus) String() string            { return proto.CompactTextString(m) }
func (*EchoStatus) ProtoMessage()               {}
func (*EchoStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EchoStatus) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *EchoStatus) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// Unary request.
type SimpleRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, server randomly chooses one from other formats.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Desired payload size in the response from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	ResponseSize *int32 `protobuf:"varint,2,opt,name=response_size,json=responseSize" json:"response_size,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	// Whether SimpleResponse should include username.
	FillUsername *bool `protobuf:"varint,4,opt,name=fill_username,json=fillUsername" json:"fill_username,omitempty"`
	// Whether SimpleResponse should include OAuth scope.
	FillOauthScope *bool `protobuf:"varint,5,opt,name=fill_oauth_scope,json=fillOauthScope" json:"fill_oauth_scope,omitempty"`
	// Whether server should return a given status
	ResponseStatus   *EchoStatus `protobuf:"bytes,7,opt,name=response_status,json=responseStatus" json:"response_status,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SimpleRequest) Reset()                    { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string            { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()               {}
func (*SimpleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SimpleRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *SimpleRequest) GetResponseSize() int32 {
	if m != nil && m.ResponseSize != nil {
		return *m.ResponseSize
	}
	return 0
}

func (m *SimpleRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleRequest) GetFillUsername() bool {
	if m != nil && m.FillUsername != nil {
		return *m.FillUsername
	}
	return false
}

func (m *SimpleRequest) GetFillOauthScope() bool {
	if m != nil && m.FillOauthScope != nil {
		return *m.FillOauthScope
	}
	return false
}

func (m *SimpleRequest) GetResponseStatus() *EchoStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

// Unary response, as configured by the request.
type SimpleResponse struct {
	// Payload to increase message size.
	Payload *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	// The user the request came from, for verifying authentication was
	// successful when the client expected it.
	Username *string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	// OAuth scope.
	OauthScope       *string `protobuf:"bytes,3,opt,name=oauth_scope,json=oauthScope" json:"oauth_scope,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SimpleResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *SimpleResponse) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *SimpleResponse) GetOauthScope() string {
	if m != nil && m.OauthScope != nil {
		return *m.OauthScope
	}
	return ""
}

// Client-streaming request.
type StreamingInputCallRequest struct {
	// Optional input payload sent along with the request.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingInputCallRequest) Reset()                    { *m = StreamingInputCallRequest{} }
func (m *StreamingInputCallRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingInputCallRequest) ProtoMessage()               {}
func (*StreamingInputCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StreamingInputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Client-streaming response.
type StreamingInputCallResponse struct {
	// Aggregated size of payloads received from the client.
	AggregatedPayloadSize *int32 `protobuf:"varint,1,opt,name=aggregated_payload_size,json=aggregatedPayloadSize" json:"aggregated_payload_size,omitempty"`
	XXX_unrecognized      []byte `json:"-"`
}

func (m *StreamingInputCallResponse) Reset()                    { *m = StreamingInputCallResponse{} }
func (m *StreamingInputCallResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingInputCallResponse) ProtoMessage()               {}
func (*StreamingInputCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StreamingInputCallResponse) GetAggregatedPayloadSize() int32 {
	if m != nil && m.AggregatedPayloadSize != nil {
		return *m.AggregatedPayloadSize
	}
	return 0
}

// Configuration for a particular response.
type ResponseParameters struct {
	// Desired payload sizes in responses from the server.
	// If response_type is COMPRESSABLE, this denotes the size before compression.
	Size *int32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	// Desired interval between consecutive responses in the response stream in
	// microseconds.
	IntervalUs       *int32 `protobuf:"varint,2,opt,name=interval_us,json=intervalUs" json:"interval_us,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ResponseParameters) Reset()                    { *m = ResponseParameters{} }
func (m *ResponseParameters) String() string            { return proto.CompactTextString(m) }
func (*ResponseParameters) ProtoMessage()               {}
func (*ResponseParameters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResponseParameters) GetSize() int32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *ResponseParameters) GetIntervalUs() int32 {
	if m != nil && m.IntervalUs != nil {
		return *m.IntervalUs
	}
	return 0
}

// Server-streaming request.
type StreamingOutputCallRequest struct {
	// Desired payload type in the response from the server.
	// If response_type is RANDOM, the payload from each response in the stream
	// might be of different types. This is to simulate a mixed type of payload
	// stream.
	ResponseType *PayloadType `protobuf:"varint,1,opt,name=response_type,json=responseType,enum=grpc.testing.PayloadType" json:"response_type,omitempty"`
	// Configuration for each expected response message.
	ResponseParameters []*ResponseParameters `protobuf:"bytes,2,rep,name=response_parameters,json=responseParameters" json:"response_parameters,omitempty"`
	// Optional input payload sent along with the request.
	Payload *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	// Whether server should return a given status
	ResponseStatus   *EchoStatus `protobuf:"bytes,7,opt,name=response_status,json=responseStatus" json:"response_status,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *StreamingOutputCallRequest) Reset()                    { *m = StreamingOutputCallRequest{} }
func (m *StreamingOutputCallRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingOutputCallRequest) ProtoMessage()               {}
func (*StreamingOutputCallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *StreamingOutputCallRequest) GetResponseType() PayloadType {
	if m != nil && m.ResponseType != nil {
		return *m.ResponseType
	}
	return PayloadType_COMPRESSABLE
}

func (m *StreamingOutputCallRequest) GetResponseParameters() []*ResponseParameters {
	if m != nil {
		return m.ResponseParameters
	}
	return nil
}

func (m *StreamingOutputCallRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *StreamingOutputCallRequest) GetResponseStatus() *EchoStatus {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

// Server-streaming response, as configured by the request and parameters.
type StreamingOutputCallResponse struct {
	// Payload to increase response size.
	Payload          *Payload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StreamingOutputCallResponse) Reset()                    { *m = StreamingOutputCallResponse{} }
func (m *StreamingOutputCallResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingOutputCallResponse) ProtoMessage()               {}
func (*StreamingOutputCallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StreamingOutputCallResponse) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "grpc.testing.Empty")
	proto.RegisterType((*Payload)(nil), "grpc.testing.Payload")
	proto.RegisterType((*EchoStatus)(nil), "grpc.testing.EchoStatus")
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterType((*StreamingInputCallRequest)(nil), "grpc.testing.StreamingInputCallRequest")
	proto.RegisterType((*StreamingInputCallResponse)(nil), "grpc.testing.StreamingInputCallResponse")
	proto.RegisterType((*ResponseParameters)(nil), "grpc.testing.ResponseParameters")
	proto.RegisterType((*StreamingOutputCallRequest)(nil), "grpc.testing.StreamingOutputCallRequest")
	proto.RegisterType((*StreamingOutputCallResponse)(nil), "grpc.testing.StreamingOutputCallResponse")
	proto.RegisterEnum("grpc.testing.PayloadType", PayloadType_name, PayloadType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	// One empty request followed by one empty response.
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error)
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error)
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error)
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/EmptyCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/grpc.testing.TestService/UnaryCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) StreamingOutputCall(ctx context.Context, in *StreamingOutputCallRequest, opts ...grpc.CallOption) (TestService_StreamingOutputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpc.testing.TestService/StreamingOutputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingOutputCallClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_StreamingOutputCallClient interface {
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingOutputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingOutputCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) StreamingInputCall(ctx context.Context, opts ...grpc.CallOption) (TestService_StreamingInputCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/grpc.testing.TestService/StreamingInputCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceStreamingInputCallClient{stream}
	return x, nil
}

type TestService_StreamingInputCallClient interface {
	Send(*StreamingInputCallRequest) error
	CloseAndRecv() (*StreamingInputCallResponse, error)
	grpc.ClientStream
}

type testServiceStreamingInputCallClient struct {
	grpc.ClientStream
}

func (x *testServiceStreamingInputCallClient) Send(m *StreamingInputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallClient) CloseAndRecv() (*StreamingInputCallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamingInputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) FullDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_FullDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/grpc.testing.TestService/FullDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceFullDuplexCallClient{stream}
	return x, nil
}

type TestService_FullDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceFullDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceFullDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) HalfDuplexCall(ctx context.Context, opts ...grpc.CallOption) (TestService_HalfDuplexCallClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[3], c.cc, "/grpc.testing.TestService/HalfDuplexCall", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceHalfDuplexCallClient{stream}
	return x, nil
}

type TestService_HalfDuplexCallClient interface {
	Send(*StreamingOutputCallRequest) error
	Recv() (*StreamingOutputCallResponse, error)
	grpc.ClientStream
}

type testServiceHalfDuplexCallClient struct {
	grpc.ClientStream
}

func (x *testServiceHalfDuplexCallClient) Send(m *StreamingOutputCallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallClient) Recv() (*StreamingOutputCallResponse, error) {
	m := new(StreamingOutputCallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	// One empty request followed by one empty response.
	EmptyCall(context.Context, *Empty) (*Empty, error)
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// One request followed by a sequence of responses (streamed download).
	// The server returns the payload with client desired type and sizes.
	StreamingOutputCall(*StreamingOutputCallRequest, TestService_StreamingOutputCallServer) error
	// A sequence of requests followed by one response (streamed upload).
	// The server returns the aggregated size of client payload as the result.
	StreamingInputCall(TestService_StreamingInputCallServer) error
	// A sequence of requests with each request served by the server immediately.
	// As one request could lead to multiple responses, this interface
	// demonstrates the idea of full duplexing.
	FullDuplexCall(TestService_FullDuplexCallServer) error
	// A sequence of requests followed by a sequence of responses.
	// The server buffers all the client requests and then serves them in order. A
	// stream of responses are returned to the client when the server starts with
	// first request.
	HalfDuplexCall(TestService_HalfDuplexCallServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/EmptyCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.TestService/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_StreamingOutputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamingOutputCallRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).StreamingOutputCall(m, &testServiceStreamingOutputCallServer{stream})
}

type TestService_StreamingOutputCallServer interface {
	Send(*StreamingOutputCallResponse) error
	grpc.ServerStream
}

type testServiceStreamingOutputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingOutputCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_StreamingInputCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).StreamingInputCall(&testServiceStreamingInputCallServer{stream})
}

type TestService_StreamingInputCallServer interface {
	SendAndClose(*StreamingInputCallResponse) error
	Recv() (*StreamingInputCallRequest, error)
	grpc.ServerStream
}

type testServiceStreamingInputCallServer struct {
	grpc.ServerStream
}

func (x *testServiceStreamingInputCallServer) SendAndClose(m *StreamingInputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceStreamingInputCallServer) Recv() (*StreamingInputCallRequest, error) {
	m := new(StreamingInputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_FullDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).FullDuplexCall(&testServiceFullDuplexCallServer{stream})
}

type TestService_FullDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceFullDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceFullDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceFullDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_HalfDuplexCall_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).HalfDuplexCall(&testServiceHalfDuplexCallServer{stream})
}

type TestService_HalfDuplexCallServer interface {
	Send(*StreamingOutputCallResponse) error
	Recv() (*StreamingOutputCallRequest, error)
	grpc.ServerStream
}

type testServiceHalfDuplexCallServer struct {
	grpc.ServerStream
}

func (x *testServiceHalfDuplexCallServer) Send(m *StreamingOutputCallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceHalfDuplexCallServer) Recv() (*StreamingOutputCallRequest, error) {
	m := new(StreamingOutputCallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
		{
			MethodName: "UnaryCall",
			Handler:    _TestService_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingOutputCall",
			Handler:       _TestService_StreamingOutputCall_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamingInputCall",
			Handler:       _TestService_StreamingInputCall_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FullDuplexCall",
			Handler:       _TestService_FullDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "HalfDuplexCall",
			Handler:       _TestService_HalfDuplexCall_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

// Client API for UnimplementedService service

type UnimplementedServiceClient interface {
	// A call that no server should implement
	UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type unimplementedServiceClient struct {
	cc *grpc.ClientConn
}

func NewUnimplementedServiceClient(cc *grpc.ClientConn) UnimplementedServiceClient {
	return &unimplementedServiceClient{cc}
}

func (c *unimplementedServiceClient) UnimplementedCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/grpc.testing.UnimplementedService/UnimplementedCall", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UnimplementedService service

type UnimplementedServiceServer interface {
	// A call that no server should implement
	UnimplementedCall(context.Context, *Empty) (*Empty, error)
}

func RegisterUnimplementedServiceServer(s *grpc.Server, srv UnimplementedServiceServer) {
	s.RegisterService(&_UnimplementedService_serviceDesc, srv)
}

func _UnimplementedService_UnimplementedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnimplementedServiceServer).UnimplementedCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.UnimplementedService/UnimplementedCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnimplementedServiceServer).UnimplementedCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _UnimplementedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.UnimplementedService",
	HandlerType: (*UnimplementedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnimplementedCall",
			Handler:    _UnimplementedService_UnimplementedCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x69, 0x42, 0xda, 0x49, 0x6a, 0xc2, 0x94, 0x0a, 0x37, 0x45, 0x22, 0x32, 0x07, 0x0c,
	0x12, 0x01, 0x45, 0x82, 0x03, 0x12, 0xa0, 0xd2, 0xa6, 0xa2, 0x52, 0xdb, 0x14, 0xbb, 0x39, 0x47,
	0x4b, 0x32, 0x75, 0x2d, 0xf9, 0x0b, 0x7b, 0x5d, 0x91, 0x1e, 0xf8, 0x33, 0xfc, 0x08, 0x0e, 0xfc,
	0x39, 0xb4, 0x6b, 0x3b, 0x71, 0xd2, 0x54, 0x34, 0x7c, 0xdd, 0x76, 0xdf, 0xbe, 0xf9, 0x78, 0x33,
	0xcf, 0x06, 0xe0, 0x14, 0xf3, 0x76, 0x18, 0x05, 0x3c, 0xc0, 0xba, 0x1d, 0x85, 0xc3, 0xb6, 0x00,
	0x1c, 0xdf, 0xd6, 0xab, 0x50, 0xe9, 0x7a, 0x21, 0x1f, 0xeb, 0x87, 0x50, 0x3d, 0x61, 0x63, 0x37,
	0x60, 0x23, 0x7c, 0x06, 0x65, 0x3e, 0x0e, 0x49, 0x53, 0x5a, 0x8a, 0xa1, 0x76, 0xb6, 0xda, 0xc5,
	0x80, 0x76, 0x46, 0x3a, 0x1d, 0x87, 0x64, 0x4a, 0x1a, 0x22, 0x94, 0x3f, 0x05, 0xa3, 0xb1, 0x56,
	0x6a, 0x29, 0x46, 0xdd, 0x94, 0x67, 0xfd, 0x35, 0x40, 0x77, 0x78, 0x1e, 0x58, 0x9c, 0xf1, 0x24,
	0x16, 0x8c, 0x61, 0x30, 0x4a, 0x13, 0x56, 0x4c, 0x79, 0x46, 0x0d, 0xaa, 0x1e, 0xc5, 0x31, 0xb3,
	0x49, 0x06, 0xae, 0x99, 0xf9, 0x55, 0xff, 0x5e, 0x82, 0x75, 0xcb, 0xf1, 0x42, 0x97, 0x4c, 0xfa,
	0x9c, 0x50, 0xcc, 0xf1, 0x2d, 0xac, 0x47, 0x14, 0x87, 0x81, 0x1f, 0xd3, 0xe0, 0x66, 0x9d, 0xd5,
	0x73, 0xbe, 0xb8, 0xe1, 0xa3, 0x42, 0x7c, 0xec, 0x5c, 0xa6, 0x15, 0x2b, 0x53, 0x92, 0xe5, 0x5c,
	0x12, 0x3e, 0x87, 0x6a, 0x98, 0x66, 0xd0, 0x56, 0x5a, 0x8a, 0x51, 0xeb, 0x6c, 0x2e, 0x4c, 0x6f,
	0xe6, 0x2c, 0x91, 0xf5, 0xcc, 0x71, 0xdd, 0x41, 0x12, 0x53, 0xe4, 0x33, 0x8f, 0xb4, 0x72, 0x4b,
	0x31, 0x56, 0xcd, 0xba, 0x00, 0xfb, 0x19, 0x86, 0x06, 0x34, 0x24, 0x29, 0x60, 0x09, 0x3f, 0x1f,
	0xc4, 0xc3, 0x20, 0x24, 0xad, 0x22, 0x79, 0xaa, 0xc0, 0x7b, 0x02, 0xb6, 0x04, 0x8a, 0x3b, 0x70,
	0x67, 0xda, 0xa4, 0x9c, 0x9b, 0x56, 0x95, 0x7d, 0x68, 0xb3, 0x7d, 0x4c, 0xe7, 0x6a, 0xaa, 0x13,
	0x01, 0xf2, 0xae, 0x7f, 0x05, 0x35, 0x1f, 0x5c, 0x8a, 0x17, 0x45, 0x29, 0x37, 0x12, 0xd5, 0x84,
	0xd5, 0x89, 0x9e, 0x74, 0x2f, 0x93, 0x3b, 0x3e, 0x84, 0x5a, 0x51, 0xc6, 0x8a, 0x7c, 0x86, 0x60,
	0x22, 0x41, 0x3f, 0x84, 0x2d, 0x8b, 0x47, 0xc4, 0x3c, 0xc7, 0xb7, 0x0f, 0xfc, 0x30, 0xe1, 0xbb,
	0xcc, 0x75, 0xf3, 0x25, 0x2e, 0xdb, 0x8a, 0x7e, 0x0a, 0xcd, 0x45, 0xd9, 0x32, 0x65, 0xaf, 0xe0,
	0x3e, 0xb3, 0xed, 0x88, 0x6c, 0xc6, 0x69, 0x34, 0xc8, 0x62, 0xd2, 0xed, 0xa6, 0x36, 0xdb, 0x9c,
	0x3e, 0x67, 0xa9, 0xc5, 0x9a, 0xf5, 0x03, 0xc0, 0x3c, 0xc7, 0x09, 0x8b, 0x98, 0x47, 0x9c, 0x22,
	0xe9, 0xd0, 0x42, 0xa8, 0x3c, 0x0b, 0xb9, 0x8e, 0xcf, 0x29, 0xba, 0x60, 0x62, 0xc7, 0x99, 0x67,
	0x20, 0x87, 0xfa, 0xb1, 0xfe, 0xad, 0x54, 0xe8, 0xb0, 0x97, 0xf0, 0x39, 0xc1, 0x7f, 0xea, 0xda,
	0x8f, 0xb0, 0x31, 0x89, 0x0f, 0x27, 0xad, 0x6a, 0xa5, 0xd6, 0x8a, 0x51, 0xeb, 0xb4, 0x66, 0xb3,
	0x5c, 0x95, 0x64, 0x62, 0x74, 0x55, 0xe6, 0xd2, 0x1e, 0xff, 0x0b, 0xa6, 0x3c, 0x86, 0xed, 0x85,
	0x43, 0xfa, 0x4d, 0x87, 0x3e, 0x7d, 0x07, 0xb5, 0xc2, 0xcc, 0xb0, 0x01, 0xf5, 0xdd, 0xde, 0xd1,
	0x89, 0xd9, 0xb5, 0xac, 0x9d, 0xf7, 0x87, 0xdd, 0xc6, 0x2d, 0x44, 0x50, 0xfb, 0xc7, 0x33, 0x98,
	0x82, 0x00, 0xb7, 0xcd, 0x9d, 0xe3, 0xbd, 0xde, 0x51, 0xa3, 0xd4, 0xf9, 0x51, 0x86, 0xda, 0x29,
	0xc5, 0xdc, 0xa2, 0xe8, 0xc2, 0x19, 0x12, 0xbe, 0x84, 0x35, 0xf9, 0x0b, 0x14, 0x6d, 0xe1, 0xc6,
	0x9c, 0x2e, 0xf1, 0xd0, 0x5c, 0x04, 0xe2, 0x3e, 0xac, 0xf5, 0x7d, 0x16, 0xa5, 0x61, 0xdb, 0xb3,
	0x8c, 0x99, 0xdf, 0x57, 0xf3, 0xc1, 0xe2, 0xc7, 0x6c, 0x00, 0x2e, 0x6c, 0x2c, 0x98, 0x0f, 0x1a,
	0x73, 0x41, 0xd7, 0xfa, 0xac, 0xf9, 0xe4, 0x06, 0xcc, 0xb4, 0xd6, 0x0b, 0x05, 0x1d, 0xc0, 0xab,
	0x1f, 0x15, 0x3e, 0xbe, 0x26, 0xc5, 0xfc, 0x47, 0xdc, 0x34, 0x7e, 0x4d, 0x4c, 0x4b, 0x19, 0xa2,
	0x94, 0xba, 0x9f, 0xb8, 0xee, 0x5e, 0x12, 0xba, 0xf4, 0xe5, 0x9f, 0x69, 0x32, 0x14, 0xa9, 0x4a,
	0xfd, 0xc0, 0xdc, 0xb3, 0xff, 0x50, 0xaa, 0xd3, 0x87, 0x7b, 0x7d, 0x5f, 0x6e, 0xd0, 0x23, 0x9f,
	0xd3, 0x28, 0x77, 0xd1, 0x1b, 0xb8, 0x3b, 0x83, 0x2f, 0xe7, 0xa6, 0x9f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xdd, 0xb5, 0x50, 0x6f, 0xa2, 0x07, 0x00, 0x00,
}
