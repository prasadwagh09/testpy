/*
 *
 * Copyright 2014 gRPC authors.
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

package transport

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

func (s) TestTimeoutDecode(t *testing.T) {
	for _, test := range []struct {
		// input
		s string
		// output
		d   time.Duration
		err error
	}{
		{"1234S", time.Second * 1234, nil},
		{"1234x", 0, fmt.Errorf("transport: timeout unit is not recognized: %q", "1234x")},
		{"1", 0, fmt.Errorf("transport: timeout string is too short: %q", "1")},
		{"", 0, fmt.Errorf("transport: timeout string is too short: %q", "")},
	} {
		d, err := decodeTimeout(test.s)
		if d != test.d || fmt.Sprint(err) != fmt.Sprint(test.err) {
			t.Fatalf("timeoutDecode(%q) = %d, %v, want %d, %v", test.s, int64(d), err, int64(test.d), test.err)
		}
	}
}

func (s) TestEncodeGrpcMessage(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Hello", "Hello"},
		{"\u0000", "%00"},
		{"%", "%25"},
		{"系统", "%E7%B3%BB%E7%BB%9F"},
		{string([]byte{0xff, 0xfe, 0xfd}), "%EF%BF%BD%EF%BF%BD%EF%BF%BD"},
	} {
		actual := encodeGrpcMessage(tt.input)
		if tt.expected != actual {
			t.Errorf("encodeGrpcMessage(%q) = %q, want %q", tt.input, actual, tt.expected)
		}
	}

	// make sure that all the visible ASCII chars except '%' are not percent encoded.
	for i := ' '; i <= '~' && i != '%'; i++ {
		output := encodeGrpcMessage(string(i))
		if output != string(i) {
			t.Errorf("encodeGrpcMessage(%v) = %v, want %v", string(i), output, string(i))
		}
	}

	// make sure that all the invisible ASCII chars and '%' are percent encoded.
	for i := rune(0); i == '%' || (i >= rune(0) && i < ' ') || (i > '~' && i <= rune(127)); i++ {
		output := encodeGrpcMessage(string(i))
		expected := fmt.Sprintf("%%%02X", i)
		if output != expected {
			t.Errorf("encodeGrpcMessage(%v) = %v, want %v", string(i), output, expected)
		}
	}
}

func (s) TestDecodeGrpcMessage(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Hello", "Hello"},
		{"H%61o", "Hao"},
		{"H%6", "H%6"},
		{"%G0", "%G0"},
		{"%E7%B3%BB%E7%BB%9F", "系统"},
		{"%EF%BF%BD", "�"},
	} {
		actual := decodeGrpcMessage(tt.input)
		if tt.expected != actual {
			t.Errorf("decodeGrpcMessage(%q) = %q, want %q", tt.input, actual, tt.expected)
		}
	}

	// make sure that all the visible ASCII chars except '%' are not percent decoded.
	for i := ' '; i <= '~' && i != '%'; i++ {
		output := decodeGrpcMessage(string(i))
		if output != string(i) {
			t.Errorf("decodeGrpcMessage(%v) = %v, want %v", string(i), output, string(i))
		}
	}

	// make sure that all the invisible ASCII chars and '%' are percent decoded.
	for i := rune(0); i == '%' || (i >= rune(0) && i < ' ') || (i > '~' && i <= rune(127)); i++ {
		output := decodeGrpcMessage(fmt.Sprintf("%%%02X", i))
		if output != string(i) {
			t.Errorf("decodeGrpcMessage(%v) = %v, want %v", fmt.Sprintf("%%%02X", i), output, string(i))
		}
	}
}

// Decode an encoded string should get the same thing back, except for invalid
// utf8 chars.
func (s) TestDecodeEncodeGrpcMessage(t *testing.T) {
	testCases := []struct {
		orig string
		want string
	}{
		{"", ""},
		{"hello", "hello"},
		{"h%6", "h%6"},
		{"%G0", "%G0"},
		{"系统", "系统"},
		{"Hello, 世界", "Hello, 世界"},

		{string([]byte{0xff, 0xfe, 0xfd}), "���"},
		{string([]byte{0xff}) + "Hello" + string([]byte{0xfe}) + "世界" + string([]byte{0xfd}), "�Hello�世界�"},
	}
	for _, tC := range testCases {
		got := decodeGrpcMessage(encodeGrpcMessage(tC.orig))
		if got != tC.want {
			t.Errorf("decodeGrpcMessage(encodeGrpcMessage(%q)) = %q, want %q", tC.orig, got, tC.want)
		}
	}
}

const binaryValue = "\u0080"

func (s) TestEncodeMetadataHeader(t *testing.T) {
	for _, test := range []struct {
		// input
		kin string
		vin string
		// output
		vout string
	}{
		{"key", "abc", "abc"},
		{"KEY", "abc", "abc"},
		{"key-bin", "abc", "YWJj"},
		{"key-bin", binaryValue, "woA"},
	} {
		v := encodeMetadataHeader(test.kin, test.vin)
		if !reflect.DeepEqual(v, test.vout) {
			t.Fatalf("encodeMetadataHeader(%q, %q) = %q, want %q", test.kin, test.vin, v, test.vout)
		}
	}
}

func (s) TestDecodeMetadataHeader(t *testing.T) {
	for _, test := range []struct {
		// input
		kin string
		vin string
		// output
		vout string
		err  error
	}{
		{"a", "abc", "abc", nil},
		{"key-bin", "Zm9vAGJhcg==", "foo\x00bar", nil},
		{"key-bin", "Zm9vAGJhcg", "foo\x00bar", nil},
		{"key-bin", "woA=", binaryValue, nil},
		{"a", "abc,efg", "abc,efg", nil},
	} {
		v, err := decodeMetadataHeader(test.kin, test.vin)
		if !reflect.DeepEqual(v, test.vout) || !reflect.DeepEqual(err, test.err) {
			t.Fatalf("decodeMetadataHeader(%q, %q) = %q, %v, want %q, %v", test.kin, test.vin, v, err, test.vout, test.err)
		}
	}
}

func (s) TestDecodeHeaderH2ErrCode(t *testing.T) {
	for _, test := range []struct {
		name string
		// input
		metaHeaderFrame *http2.MetaHeadersFrame
		serverSide      bool
		// output
		wantCode http2.ErrCode
	}{
		{
			name: "valid header",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: []hpack.HeaderField{
				{Name: "content-type", Value: "application/grpc"},
			}},
			wantCode: http2.ErrCodeNo,
		},
		{
			name: "valid header serverSide",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: []hpack.HeaderField{
				{Name: "content-type", Value: "application/grpc"},
			}},
			serverSide: true,
			wantCode:   http2.ErrCodeNo,
		},
		{
			name: "invalid grpc status header field",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: []hpack.HeaderField{
				{Name: "content-type", Value: "application/grpc"},
				{Name: "grpc-status", Value: "xxxx"},
			}},
			wantCode: http2.ErrCodeProtocol,
		},
		{
			name: "invalid http content type",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: []hpack.HeaderField{
				{Name: "content-type", Value: "application/json"},
			}},
			wantCode: http2.ErrCodeProtocol,
		},
		{
			name: "http fallback and invalid http status",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: []hpack.HeaderField{
				// No content type provided then fallback into handling http error.
				{Name: ":status", Value: "xxxx"},
			}},
			wantCode: http2.ErrCodeProtocol,
		},
		{
			name:            "http2 frame size exceeds",
			metaHeaderFrame: &http2.MetaHeadersFrame{Fields: nil, Truncated: true},
			wantCode:        http2.ErrCodeFrameSize,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			state := &decodeState{serverSide: test.serverSide}
			if h2code, _ := state.decodeHeader(test.metaHeaderFrame); h2code != test.wantCode {
				t.Fatalf("decodeState.decodeHeader(%v) = %v, want %v", test.metaHeaderFrame, h2code, test.wantCode)
			}
		})
	}
}

func (s) TestParseDialTarget(t *testing.T) {
	for _, test := range []struct {
		target, wantNet, wantAddr string
	}{
		{"unix:a", "unix", "a"},
		{"unix:a/b/c", "unix", "a/b/c"},
		{"unix:/a", "unix", "/a"},
		{"unix:/a/b/c", "unix", "/a/b/c"},
		{"unix://a", "unix", "a"},
		{"unix://a/b/c", "unix", "/b/c"},
		{"unix:///a", "unix", "/a"},
		{"unix:///a/b/c", "unix", "/a/b/c"},
		{"unix:etcd:0", "unix", "etcd:0"},
		{"unix:///tmp/unix-3", "unix", "/tmp/unix-3"},
		{"unix://domain", "unix", "domain"},
		{"unix://etcd:0", "unix", "etcd:0"},
		{"unix:///etcd:0", "unix", "/etcd:0"},
		{"unix-abstract:abc", "unix", "\x00abc"},
		{"unix-abstract:abc edf", "unix", "\x00abc edf"},
		{"unix-abstract:///abc", "unix", "\x00///abc"},
		{"unix-abstract:unix:abc", "unix", "\x00unix:abc"},
		{"passthrough://unix://domain", "tcp", "passthrough://unix://domain"},
		{"https://google.com:443", "tcp", "https://google.com:443"},
		{"dns:///google.com", "tcp", "dns:///google.com"},
		{"/unix/socket/address", "tcp", "/unix/socket/address"},
	} {
		gotNet, gotAddr := parseDialTarget(test.target)
		if gotNet != test.wantNet || gotAddr != test.wantAddr {
			t.Errorf("parseDialTarget(%q) = %s, %s want %s, %s", test.target, gotNet, gotAddr, test.wantNet, test.wantAddr)
		}
	}
}
