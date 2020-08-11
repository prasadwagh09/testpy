// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package helloworld

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterService is the service API for Greeter service.
// Fields should be assigned to their respective handler implementations only before
// RegisterGreeterService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type GreeterService struct {
	// Sends a greeting
	SayHello func(context.Context, *HelloRequest) (*HelloReply, error)
}

func (s *GreeterService) sayHello(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.SayHello == nil {
		return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
	}
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterGreeterService registers a service implementation with a gRPC server.
func RegisterGreeterService(s grpc.ServiceRegistrar, srv *GreeterService) {
	sd := grpc.ServiceDesc{
		ServiceName: "helloworld.Greeter",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "SayHello",
				Handler:    srv.sayHello,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "examples/helloworld/helloworld/helloworld.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewGreeterService creates a new GreeterService containing the
// implemented methods of the Greeter service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care.
func NewGreeterService(s interface{}) *GreeterService {
	ns := &GreeterService{}
	if h, ok := s.(interface {
		SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	}); ok {
		ns.SayHello = h.SayHello
	}
	return ns
}

// UnstableGreeterService is the service API for Greeter service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableGreeterService interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}
