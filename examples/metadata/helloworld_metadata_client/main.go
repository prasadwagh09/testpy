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

package main

import (
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/metadata/helloworld"
	"google.golang.org/grpc/metadata"
)

const (
	address         = "localhost:9527"
	timestampFormat = time.StampNano // "Jan _2 15:04:05.000"
)

func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	log.Printf("------------ unary ------------")
	// create metadata and context
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// call RPC
	var header, trailer metadata.MD
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:")
		for i, e := range l {
			log.Printf(" %d. %s", i, e)
		}
	}
	log.Printf("message:")
	log.Printf(" - %s", r.Message)
	if t, ok := trailer["timestamp"]; ok {
		log.Printf("timestamp from trailer:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
}

func serverStreamingWithMetadata(c pb.GreeterClient, names []string) {
	log.Printf("------------ server streaming ------------")
	// create metadata and context
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// call RPC
	stream, err := c.ServerStreamingSayHello(ctx, &pb.StreamingHelloRequest{Names: names})
	if err != nil {
		log.Fatalf("failed to call ServerStreamingSayHello: %v", err)
	}

	// read header
	header, err := stream.Header()
	if err != nil {
		log.Fatalf("failed to get header from stream: %v", err)
	}
	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:")
		for i, e := range l {
			log.Printf(" %d. %s", i, e)
		}
	}

	// read response
	var rpcStatus error
	log.Printf("message:")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		log.Printf(" - %s", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}

	// read trailer
	trailer := stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		log.Printf("timestamp from trailer:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
}

func clientStreamWithMetadata(c pb.GreeterClient, names []string) {
	log.Printf("------------ client streaming ------------")
	// create metadata and context
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// call RPC
	stream, err := c.ClientStreamingSayHello(ctx)
	if err != nil {
		log.Fatalf("failed to call ClientStreamingSayHello: %v\n", err)
	}

	// read header
	header, err := stream.Header()
	if err != nil {
		log.Fatalf("failed to get header from stream: %v", err)
	}
	if t, ok := header["timestamp"]; ok {
		log.Printf("timestamp from header:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
	if l, ok := header["location"]; ok {
		log.Printf("location from header:")
		for i, e := range l {
			log.Printf(" %d. %s", i, e)
		}
	}

	// send request to stream
	for _, name := range names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("failed to send streaming: %v\n", err)
		}
	}

	// read response
	r, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to CloseAndRecv: %v\n", err)
	}
	log.Printf("message:")
	for _, m := range r.Messages {
		log.Printf(" - %s\n", m)
	}

	// read trailer
	trailer := stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		log.Printf("timestamp from trailer:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}
}

func bidirectionalWithMetadata(c pb.GreeterClient, names []string) {
	log.Printf("------------ bidirectional ------------")
	// create metadata and context
	md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// call RPC
	stream, err := c.BidirectionalStreamingSayHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidirectionalStreamingSayHello: %v\n", err)
	}

	go func() {
		// read header
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}
		if t, ok := header["timestamp"]; ok {
			log.Printf("timestamp from header:")
			for i, e := range t {
				log.Printf(" %d. %s", i, e)
			}
		}
		if l, ok := header["location"]; ok {
			log.Printf("location from header:")
			for i, e := range l {
				log.Printf(" %d. %s", i, e)
			}
		}

		// send request
		for _, name := range names {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// read response
	var rpcStatus error
	log.Printf("message:")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		log.Printf(" - %s", r.Message)
	}
	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server streaming: %v", rpcStatus)
	}

	// read trailer
	trailer := stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		log.Printf("timestamp from trailer:")
		for i, e := range t {
			log.Printf(" %d. %s", i, e)
		}
	}

}

var names = []string{
	"Anne",
	"Hope",
	"Margeret",
	"Jamar",
	"Judson",
	"Carrol",
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	unaryCallWithMetadata(c, names[0])
	time.Sleep(1 * time.Second)

	serverStreamingWithMetadata(c, names)
	time.Sleep(1 * time.Second)

	clientStreamWithMetadata(c, names)
	time.Sleep(1 * time.Second)

	bidirectionalWithMetadata(c, names)
}
