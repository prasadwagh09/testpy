// Copyright 2015 gRPC authors.
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

// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: grpc/testing/report_qps_scenario_service.proto

package grpc_testing

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ReportQpsScenarioService_ReportScenario_FullMethodName = "/grpc.testing.ReportQpsScenarioService/ReportScenario"
)

// ReportQpsScenarioServiceClient is the client API for ReportQpsScenarioService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportQpsScenarioServiceClient interface {
	// Report results of a QPS test benchmark scenario.
	ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error)
}

type reportQpsScenarioServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportQpsScenarioServiceClient(cc grpc.ClientConnInterface) ReportQpsScenarioServiceClient {
	return &reportQpsScenarioServiceClient{cc}
}

func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, in *ScenarioResult, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, ReportQpsScenarioService_ReportScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportQpsScenarioServiceServer is the server API for ReportQpsScenarioService service.
// All implementations must embed UnimplementedReportQpsScenarioServiceServer
// for forward compatibility
type ReportQpsScenarioServiceServer interface {
	// Report results of a QPS test benchmark scenario.
	ReportScenario(context.Context, *ScenarioResult) (*Void, error)
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

// UnimplementedReportQpsScenarioServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportQpsScenarioServiceServer struct {
}

func (UnimplementedReportQpsScenarioServiceServer) ReportScenario(context.Context, *ScenarioResult) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportScenario not implemented")
}
func (UnimplementedReportQpsScenarioServiceServer) mustEmbedUnimplementedReportQpsScenarioServiceServer() {
}

// UnsafeReportQpsScenarioServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportQpsScenarioServiceServer will
// result in compilation errors.
type UnsafeReportQpsScenarioServiceServer interface {
	mustEmbedUnimplementedReportQpsScenarioServiceServer()
}

func RegisterReportQpsScenarioServiceServer(s grpc.ServiceRegistrar, srv ReportQpsScenarioServiceServer) {
	s.RegisterService(&ReportQpsScenarioService_ServiceDesc, srv)
}

func _ReportQpsScenarioService_ReportScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScenarioResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportQpsScenarioService_ReportScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportQpsScenarioServiceServer).ReportScenario(ctx, req.(*ScenarioResult))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportQpsScenarioService_ServiceDesc is the grpc.ServiceDesc for ReportQpsScenarioService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportQpsScenarioService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.ReportQpsScenarioService",
	HandlerType: (*ReportQpsScenarioServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportScenario",
			Handler:    _ReportQpsScenarioService_ReportScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/testing/report_qps_scenario_service.proto",
}
