// Copyright 2024 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: test_client_streaming.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClientStreamingService_ClientMethod_FullMethodName = "/main.ClientStreamingService/clientMethod"
)

// ClientStreamingServiceClient is the client API for ClientStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamingServiceClient interface {
	ClientMethod(ctx context.Context, opts ...grpc.CallOption) (ClientStreamingService_ClientMethodClient, error)
}

type clientStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamingServiceClient(cc grpc.ClientConnInterface) ClientStreamingServiceClient {
	return &clientStreamingServiceClient{cc}
}

func (c *clientStreamingServiceClient) ClientMethod(ctx context.Context, opts ...grpc.CallOption) (ClientStreamingService_ClientMethodClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStreamingService_ServiceDesc.Streams[0], ClientStreamingService_ClientMethod_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamingServiceClientMethodClient{stream}
	return x, nil
}

type ClientStreamingService_ClientMethodClient interface {
	Send(*EventRequest) error
	CloseAndRecv() (*EventResponse, error)
	grpc.ClientStream
}

type clientStreamingServiceClientMethodClient struct {
	grpc.ClientStream
}

func (x *clientStreamingServiceClientMethodClient) Send(m *EventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamingServiceClientMethodClient) CloseAndRecv() (*EventResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EventResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamingServiceServer is the server API for ClientStreamingService service.
// All implementations must embed UnimplementedClientStreamingServiceServer
// for forward compatibility
type ClientStreamingServiceServer interface {
	ClientMethod(ClientStreamingService_ClientMethodServer) error
	mustEmbedUnimplementedClientStreamingServiceServer()
}

// UnimplementedClientStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamingServiceServer struct {
}

func (UnimplementedClientStreamingServiceServer) ClientMethod(ClientStreamingService_ClientMethodServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientMethod not implemented")
}
func (UnimplementedClientStreamingServiceServer) mustEmbedUnimplementedClientStreamingServiceServer() {
}

// UnsafeClientStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamingServiceServer will
// result in compilation errors.
type UnsafeClientStreamingServiceServer interface {
	mustEmbedUnimplementedClientStreamingServiceServer()
}

func RegisterClientStreamingServiceServer(s grpc.ServiceRegistrar, srv ClientStreamingServiceServer) {
	s.RegisterService(&ClientStreamingService_ServiceDesc, srv)
}

func _ClientStreamingService_ClientMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamingServiceServer).ClientMethod(&clientStreamingServiceClientMethodServer{stream})
}

type ClientStreamingService_ClientMethodServer interface {
	SendAndClose(*EventResponse) error
	Recv() (*EventRequest, error)
	grpc.ServerStream
}

type clientStreamingServiceClientMethodServer struct {
	grpc.ServerStream
}

func (x *clientStreamingServiceClientMethodServer) SendAndClose(m *EventResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamingServiceClientMethodServer) Recv() (*EventRequest, error) {
	m := new(EventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamingService_ServiceDesc is the grpc.ServiceDesc for ClientStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.ClientStreamingService",
	HandlerType: (*ClientStreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "clientMethod",
			Handler:       _ClientStreamingService_ClientMethod_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test_client_streaming.proto",
}
