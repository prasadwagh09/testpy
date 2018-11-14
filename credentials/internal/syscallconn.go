// +build !appengine

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

package internal

import (
	"errors"
	"net"
	"syscall"
)

// syscallConn keeps reference of rawConn to support syscall.Conn for channelz.
// SyscallConn() (the method in interface syscall.Conn) is explicitly
// implemented on this type,
//
// Interface syscall.Conn is implemented by most net.Conn implementations (e.g.
// TCPConn, UnixConn), but is not part of net.Conn interface. So wrapper conns
// that embed net.Conn don't implement syscall.Conn. (Side note: tls.Conn
// doesn't embed net.Conn, so even if syscall.Conn is part of net.Conn, it won't
// help here).
type syscallConn struct {
	net.Conn
	rawConn net.Conn
}

// WrapSyscallConn tries to wrapper rawConn and newConn into a net.Conn that
// implements syscall.Conn. rawConn will be used to support syscall, and newConn
// will be used for read/write.
//
// This function returns newConn if rawConn doesn't implement syscall.Conn.
func WrapSyscallConn(rawConn, newConn net.Conn) net.Conn {
	if _, ok := rawConn.(syscall.Conn); !ok {
		return newConn
	}
	return &syscallConn{
		Conn:    newConn,
		rawConn: rawConn,
	}
}

// implements the syscall.Conn interface
func (c *syscallConn) SyscallConn() (syscall.RawConn, error) {
	conn, ok := c.rawConn.(syscall.Conn)
	if !ok {
		// This should never happen because we already checked rawConn in
		// newConnSyscall(). It is kept to avoid panic.
		return nil, errors.New("RawConn does not implement syscall.Conn")
	}
	return conn.SyscallConn()
}
