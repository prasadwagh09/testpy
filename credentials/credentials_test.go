/*
 *
 * Copyright 2016, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package credentials

import (
	"crypto/tls"
	"net"
	"testing"

	"golang.org/x/net/context"
)

func TestTLSOverrideServerName(t *testing.T) {
	expectedServerName := "server.name"
	c := NewTLS(nil)
	c.OverrideServerName(expectedServerName)
	if c.Info().ServerName != expectedServerName {
		t.Fatalf("c.Info().ServerName = %v, want %v", c.Info().ServerName, expectedServerName)
	}
}

func TestTLSClone(t *testing.T) {
	expectedServerName := "server.name"
	c := NewTLS(nil)
	c.OverrideServerName(expectedServerName)
	cc := c.Clone()
	if cc.Info().ServerName != expectedServerName {
		t.Fatalf("cc.Info().ServerName = %v, want %v", cc.Info().ServerName, expectedServerName)
	}
	cc.OverrideServerName("")
	if c.Info().ServerName != expectedServerName {
		t.Fatalf("Change in clone should not affect the original, c.Info().ServerName = %v, want %v", c.Info().ServerName, expectedServerName)
	}
}

func TestTLSClientHandshakeReturnsAuthInfo(t *testing.T) {
	localPort := ":5050"
	tlsDir := "../test/testdata/"
	lis, err := net.Listen("tcp", localPort)
	if err != nil {
		t.Fatalf("Failed to start local server. Listener error: %v", err)
	}
	serverTLS, err := NewServerTLSFromFile(tlsDir+"server1.pem", tlsDir+"server1.key")
	if err != nil {
		t.Fatalf("Failed to create server TLS. Error: %v", err)
	}
	var serverAuthInfo AuthInfo
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		serverRawConn, _ := lis.Accept()
		serverConn := tls.Server(serverRawConn, serverTLS.(*tlsCreds).config)
		serverErr := serverConn.Handshake()
		if serverErr != nil {
			t.Fatalf("Error on server while handshake. Error: %v", serverErr)
		}
		serverAuthInfo = TLSInfo{serverConn.ConnectionState()}
	}()
	defer lis.Close()
	conn, err := net.Dial("tcp", localPort)
	if err != nil {
		t.Fatalf("Client failed to connect to local server. Error: %v", err)
	}
	c := NewTLS(&tls.Config{InsecureSkipVerify: true})
	_, authInfo, err := c.ClientHandshake(context.Background(), localPort, conn)
	if err != nil {
		t.Fatalf("Error on client while handshake. Error: %v", err)
	}
	select {
	case <-done:
		// wait until server has populated the serverAuthInfo struct.
	}
	if authInfo.AuthType() != serverAuthInfo.AuthType() {
		t.Fatalf("c.ClientHandshake(_, %v, _) = %v, want %v.", localPort, authInfo, serverAuthInfo)
	}
}

func TestTLSServerHandshakeReturnsAuthInfo(t *testing.T) {
	localPort := ":5050"
	tlsDir := "../test/testdata/"
	lis, err := net.Listen("tcp", localPort)
	if err != nil {
		t.Fatalf("Failed to start local server. Listener error: %v", err)
	}
	serverTLS, err := NewServerTLSFromFile(tlsDir+"server1.pem", tlsDir+"server1.key")
	if err != nil {
		t.Fatalf("Failed to create server TLS. Error: %v", err)
	}
	var serverAuthInfo AuthInfo
	done := make(chan bool)
	go func() {
		defer func() {
			done <- true
		}()
		serverRawConn, _ := lis.Accept()
		var serverErr error
		_, serverAuthInfo, serverErr = serverTLS.ServerHandshake(serverRawConn)
		if serverErr != nil {
			t.Fatalf("Error on server while handshake. Error: %v", serverErr)
		}
	}()
	defer lis.Close()
	conn, err := net.Dial("tcp", localPort)
	if err != nil {
		t.Fatalf("Client failed to connect to local server. Error: %v", err)
	}
	c := NewTLS(&tls.Config{InsecureSkipVerify: true})
	clientConn := tls.Client(conn, c.(*tlsCreds).config)
	err = clientConn.Handshake()
	if err != nil {
		t.Fatalf("Error on client while handshake. Error: %v", err)
	}
	authInfo := TLSInfo{clientConn.ConnectionState()}
	select {
	case <-done:
		// wait until server has populated the serverAuthInfo struct.
	}
	if authInfo.AuthType() != serverAuthInfo.AuthType() {
		t.Fatalf("ServerHandshake(_) = %v, want %v.", serverAuthInfo, authInfo)
	}

}
