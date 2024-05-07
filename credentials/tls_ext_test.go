/*
 *
 * Copyright 2023 gRPC authors.
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

package credentials_test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/envconfig"
	"google.golang.org/grpc/internal/grpctest"
	"google.golang.org/grpc/internal/stubserver"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/testdata"

	testgrpc "google.golang.org/grpc/interop/grpc_testing"
	testpb "google.golang.org/grpc/interop/grpc_testing"
)

const defaultTestTimeout = 10 * time.Second

type s struct {
	grpctest.Tester
}

func Test(t *testing.T) {
	grpctest.RunSubTests(t, s{})
}

var serverCert tls.Certificate
var certPool *x509.CertPool
var serverName = "x.test.example.com"

func init() {
	var err error
	serverCert, err = tls.LoadX509KeyPair(testdata.Path("x509/server1_cert.pem"), testdata.Path("x509/server1_key.pem"))
	if err != nil {
		panic(fmt.Sprintf("tls.LoadX509KeyPair(server1.pem, server1.key) failed: %v", err))
	}

	b, err := os.ReadFile(testdata.Path("x509/server_ca_cert.pem"))
	if err != nil {
		panic(fmt.Sprintf("Error reading CA cert file: %v", err))
	}
	certPool = x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(b) {
		panic("Error appending cert from PEM")
	}
}

// Tests that the MinVersion of tls.Config is set to 1.2 if it is not already
// set by the user.
func (s) TestTLS_MinVersion12(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	// Create server creds without a minimum version.
	serverCreds := credentials.NewTLS(&tls.Config{
		// MinVersion should be set to 1.2 by gRPC by default.
		Certificates: []tls.Certificate{serverCert},
	})
	ss := stubserver.StubServer{
		EmptyCallF: func(context.Context, *testpb.Empty) (*testpb.Empty, error) {
			return &testpb.Empty{}, nil
		},
	}

	// Create client creds that supports V1.0-V1.1.
	clientCreds := credentials.NewTLS(&tls.Config{
		ServerName: serverName,
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS10,
		MaxVersion: tls.VersionTLS11,
	})

	// Start server and client separately, because Start() blocks on a
	// successful connection, which we will not get.
	if err := ss.StartServer(grpc.Creds(serverCreds)); err != nil {
		t.Fatalf("Error starting server: %v", err)
	}
	defer ss.Stop()

	cc, err := grpc.NewClient(ss.Address, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		t.Fatalf("grpc.NewClient error: %v", err)
	}
	defer cc.Close()

	client := testgrpc.NewTestServiceClient(cc)

	const wantStr = "authentication handshake failed"
	if _, err = client.EmptyCall(ctx, &testpb.Empty{}); status.Code(err) != codes.Unavailable || !strings.Contains(status.Convert(err).Message(), wantStr) {
		t.Fatalf("EmptyCall err = %v; want code=%v, message contains %q", err, codes.Unavailable, wantStr)
	}
}

// Tests that the MinVersion of tls.Config is not changed if it is set by the
// user.
func (s) TestTLS_MinVersionOverridable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	var allCipherSuites []uint16
	for _, cs := range tls.CipherSuites() {
		allCipherSuites = append(allCipherSuites, cs.ID)
	}

	// Create server creds that allow v1.0.
	serverCreds := credentials.NewTLS(&tls.Config{
		MinVersion:   tls.VersionTLS10,
		Certificates: []tls.Certificate{serverCert},
		CipherSuites: allCipherSuites,
	})
	ss := stubserver.StubServer{
		EmptyCallF: func(context.Context, *testpb.Empty) (*testpb.Empty, error) {
			return &testpb.Empty{}, nil
		},
	}

	// Create client creds that supports V1.0-V1.1.
	clientCreds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		RootCAs:      certPool,
		CipherSuites: allCipherSuites,
		MinVersion:   tls.VersionTLS10,
		MaxVersion:   tls.VersionTLS11,
	})

	if err := ss.Start([]grpc.ServerOption{grpc.Creds(serverCreds)}, grpc.WithTransportCredentials(clientCreds)); err != nil {
		t.Fatalf("Error starting stub server: %v", err)
	}
	defer ss.Stop()

	if _, err := ss.Client.EmptyCall(ctx, &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall err = %v; want <nil>", err)
	}
}

// Tests that CipherSuites is set to exclude HTTP/2 forbidden suites by default.
func (s) TestTLS_CipherSuites(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	// Create server creds without cipher suites.
	serverCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
	})
	ss := stubserver.StubServer{
		EmptyCallF: func(context.Context, *testpb.Empty) (*testpb.Empty, error) {
			return &testpb.Empty{}, nil
		},
	}

	// Create client creds that use a forbidden suite only.
	clientCreds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		RootCAs:      certPool,
		CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_128_CBC_SHA},
		MaxVersion:   tls.VersionTLS12, // TLS1.3 cipher suites are not configurable, so limit to 1.2.
	})

	// Start server and client separately, because Start() blocks on a
	// successful connection, which we will not get.
	if err := ss.StartServer(grpc.Creds(serverCreds)); err != nil {
		t.Fatalf("Error starting server: %v", err)
	}
	defer ss.Stop()

	cc, err := grpc.NewClient("dns:"+ss.Address, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		t.Fatalf("grpc.NewClient error: %v", err)
	}
	defer cc.Close()

	client := testgrpc.NewTestServiceClient(cc)

	const wantStr = "authentication handshake failed"
	if _, err = client.EmptyCall(ctx, &testpb.Empty{}); status.Code(err) != codes.Unavailable || !strings.Contains(status.Convert(err).Message(), wantStr) {
		t.Fatalf("EmptyCall err = %v; want code=%v, message contains %q", err, codes.Unavailable, wantStr)
	}
}

// Tests that CipherSuites is not overridden when it is set.
func (s) TestTLS_CipherSuitesOverridable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	// Create server that allows only a forbidden cipher suite.
	serverCreds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_128_CBC_SHA},
	})
	ss := stubserver.StubServer{
		EmptyCallF: func(context.Context, *testpb.Empty) (*testpb.Empty, error) {
			return &testpb.Empty{}, nil
		},
	}

	// Create server that allows only a forbidden cipher suite.
	clientCreds := credentials.NewTLS(&tls.Config{
		ServerName:   serverName,
		RootCAs:      certPool,
		CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_128_CBC_SHA},
		MaxVersion:   tls.VersionTLS12, // TLS1.3 cipher suites are not configurable, so limit to 1.2.
	})

	if err := ss.Start([]grpc.ServerOption{grpc.Creds(serverCreds)}, grpc.WithTransportCredentials(clientCreds)); err != nil {
		t.Fatalf("Error starting stub server: %v", err)
	}
	defer ss.Stop()

	if _, err := ss.Client.EmptyCall(ctx, &testpb.Empty{}); err != nil {
		t.Fatalf("EmptyCall err = %v; want <nil>", err)
	}
}

// TestTLS_DisabledALPN tests the behaviour of a gRPC client when connecting to
// a server that doesn't support ALPN.
func (s) TestTLS_DisabledALPN(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTestTimeout)
	defer cancel()

	initialVal := envconfig.EnforceALPNEnabled
	defer func() {
		envconfig.EnforceALPNEnabled = initialVal
	}()

	// Start a non gRPC TLS server.
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		NextProtos:   []string{}, // Empty list indicates ALPN is disabled.
	}
	listner, err := tls.Listen("tcp", ":0", config)
	if err != nil {
		t.Fatalf("Error starting TLS server: %v", err)
	}
	defer listner.Close()

	tests := []struct {
		name                   string
		alpnEnforced           bool
		wantErrMatchPattern    string
		wantErrNonMatchPattern string
	}{
		{
			name:                "enforced",
			alpnEnforced:        true,
			wantErrMatchPattern: "transport: .*missing selected ALPN property",
		},
		{
			name:                   "not_enforced",
			wantErrNonMatchPattern: "transport:",
		},
		{
			name:                   "default_value",
			wantErrNonMatchPattern: "transport:",
			alpnEnforced:           initialVal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			envconfig.EnforceALPNEnabled = tc.alpnEnforced
			// Listen to one TCP connection request.
			go func() {
				conn, err := listner.Accept()
				if err != nil {
					t.Errorf("tls.Accept failed err = %v", err)
				} else {
					_, _ = conn.Write([]byte("Hello, World!"))
					_ = conn.Close()
				}
			}()

			clientCreds := credentials.NewTLS(&tls.Config{
				ServerName: serverName,
				RootCAs:    certPool,
			})

			cc, err := grpc.NewClient("dns:"+listner.Addr().String(), grpc.WithTransportCredentials(clientCreds))
			if err != nil {
				t.Fatalf("grpc.NewClient error: %v", err)
			}
			defer cc.Close()
			client := testgrpc.NewTestServiceClient(cc)
			_, rpcErr := client.EmptyCall(ctx, &testpb.Empty{})

			if gotCode := status.Code(rpcErr); gotCode != codes.Unavailable {
				t.Errorf("EmptyCall returned unexpected code: got=%v, want=%v", gotCode, codes.Unavailable)
			}

			matchPat, err := regexp.Compile(tc.wantErrMatchPattern)
			if err != nil {
				t.Fatalf("Error message match pattern %q is invalid due to error: %v", tc.wantErrMatchPattern, err)
			}

			if tc.wantErrMatchPattern != "" && !matchPat.MatchString(status.Convert(rpcErr).Message()) {
				t.Errorf("EmptyCall err = %v; want pattern match %q", rpcErr, matchPat)
			}

			nonMatchPat, err := regexp.Compile(tc.wantErrNonMatchPattern)
			if err != nil {
				t.Fatalf("Error message non-match pattern %q is invalid due to error: %v", tc.wantErrNonMatchPattern, err)
			}

			if tc.wantErrNonMatchPattern != "" && nonMatchPat.MatchString(status.Convert(rpcErr).Message()) {
				t.Errorf("EmptyCall err = %v; want pattern missing %q", rpcErr, nonMatchPat)
			}
		})
	}
}
