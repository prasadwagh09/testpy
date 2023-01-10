/*
 *
 * Copyright 2022 gRPC authors.
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

// Binary server is an example server.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/authz"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "google.golang.org/grpc/examples/features/proto/echo"
)

const (
	unaryEchoWriterRole      = "UNARY_ECHO:W"
	streamEchoReaderRole     = "STREAM_ECHO:R"
	streamEchoWriterRole     = "STREAM_ECHO:W"
	streamEchoReadWriterRole = "STREAM_ECHO:RW"
)

var (
	port = flag.Int("port", 50051, "the port to serve on")

	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")

	authzPolicy = newPolicy("authz",
		[]string{
			newHeaderRule("allow_UnaryEcho", "/grpc.examples.echo.Echo/UnaryEcho", unaryEchoWriterRole),
			newHeaderRule("allow_ClientStreamingEcho", "/grpc.examples.echo.Echo/ClientStreamingEcho", streamEchoWriterRole),
			newHeaderRule("allow_ServerStreamingEcho", "/grpc.examples.echo.Echo/ServerStreamingEcho", streamEchoReaderRole),
			newHeaderRule("allow_BidirectionalStreamingEcho", "/grpc.examples.echo.Echo/BidirectionalStreamingEcho", streamEchoReadWriterRole),
		},
		nil,
	)
	mockedMetadata = newMockedMetadata()
)

func newMockedMetadata() metadata.MD {
	md := metadata.MD{}
	roles := []string{
		unaryEchoWriterRole,
		streamEchoReadWriterRole,
	}
	for _, role := range roles {
		md.Set(role, "true")
	}
	return md
}

func newPolicy(name string, allowRules, denyRules []string) string {
	return fmt.Sprintf(`{
		"name": "%s",
		"allow_rules":[%s],
		"deny_rules":[%s]
	}`, name,
		strings.Join(allowRules, ","),
		strings.Join(denyRules, ","),
	)
}

func newHeaderRule(name, path, role string) string {
	return fmt.Sprintf(`{
		"name": "%s",
		"request": {
			"paths":["%s"],
			"headers": [
				{
					"key": "%s",
					"values": ["true"]
				}
			]
		}
	}`, name, path, role)
}

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
func logger(format string, a ...interface{}) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("unary echoing message %q\n", in.Message)
	return &pb.EchoResponse{Message: in.Message}, nil
}

func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			fmt.Printf("server: error receiving from stream: %v\n", err)
			return err
		}
		fmt.Printf("bidi echoing message %q\n", in.Message)
		stream.Send(&pb.EchoResponse{Message: in.Message})
	}
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}

func authUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	m, err := handler(ctx, req)
	if err != nil {
		logger("RPC failed with error %v", err)
	}
	return m, err
}

// wrappedStream wraps around the embedded grpc.ServerStream, and intercepts the RecvMsg and
// SendMsg method call.
type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	logger("Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	logger("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(ctx context.Context, s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s, ctx}
}

func authStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// authentication (token verification)
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return errMissingMetadata
	}
	if !valid(md["authorization"]) {
		return errInvalidToken
	}

	err := handler(srv, newWrappedStream(ss.Context(), ss))
	if err != nil {
		logger("RPC failed with error %v", err)
	}
	return err
}

func injectResourceAccessUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// Assuming authentication was successful in the middleware chain, fetch the subject access, mocked here
	mdCtx := metadata.NewIncomingContext(ctx, mockedMetadata)
	return handler(mdCtx, req)
}

func injectResourceAccessStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Assuming authentication was successful in the middleware chain, fetch the subject access, mocked here
	mdCtx := metadata.NewIncomingContext(ss.Context(), mockedMetadata)
	return handler(srv, newWrappedStream(mdCtx, ss))
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create tls based credential.
	creds, err := credentials.NewServerTLSFromFile(data.Path("x509/server_cert.pem"), data.Path("x509/server_key.pem"))
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	staticInteceptor, err := authz.NewStatic(authzPolicy)
	if err != nil {
		log.Fatalf("failed to create static authz interceptor: %v", err)
	}
	s := grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			authUnaryInterceptor,
			injectResourceAccessUnaryInterceptor,
			staticInteceptor.UnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			authStreamInterceptor,
			injectResourceAccessStreamInterceptor,
			staticInteceptor.StreamInterceptor,
		),
	)

	// Register EchoServer on the server.
	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
