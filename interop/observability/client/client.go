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

package main

import (
	"context"
	"flag"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/gcp/observability"
	"google.golang.org/grpc/interop"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
)

var (
	serverHost           = flag.String("server_host", "localhost", "The server host name")
	serverPort           = flag.Int("server_port", 10000, "The server port number")
	exporterSleepSeconds = flag.Int("exporter_sleep_seconds", 0, "Number of seconds to wait to export observability data")
	testCase             = flag.String("test_case", "large_unary", "The action to perform")
)

func main() {
	err := observability.Start(context.Background())
	if err != nil {
		log.Fatalf("observability start failed: %v", err)
	}
	defer observability.End()
	flag.Parse()
	serverAddr := *serverHost
	if *serverPort != 0 {
		serverAddr = net.JoinHostPort(*serverHost, strconv.Itoa(*serverPort))
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
	}
	defer conn.Close()
	tc := testgrpc.NewTestServiceClient(conn)
	testCases := strings.Split(*testCase, ",")
	for _, singleCase := range testCases {
		if singleCase == "ping_pong" {
			interop.DoPingPong(tc)
		} else if singleCase == "large_unary" {
			interop.DoLargeUnaryCall(tc)
		} else if singleCase == "custom_metadata" {
			interop.DoCustomMetadata(tc)
		} else {
			log.Fatalf("Invalid test case: %s", singleCase)
		}
	}
	log.Printf("Sleeping %d seconds before closing client", *exporterSleepSeconds)
	time.Sleep(time.Duration(*exporterSleepSeconds) * time.Second)
}
