/*
 *
 * Copyright 2015 gRPC authors.
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

// Package glogger defines glog-based logging for grpc.
package glogger

import (
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc/grpclog"
)

func init() {
	grpclog.SetLogger(&glogger{})
}

type glogger struct{}

func (g *glogger) Print(l grpclog.Severity, args ...interface{}) {
	switch l {
	case grpclog.InfoLog:
		glog.InfoDepth(2, args...)
	case grpclog.WarningLog:
		glog.WarningDepth(2, args...)
	case grpclog.ErrorLog:
		glog.ErrorDepth(2, args...)
	case grpclog.FatalLog:
		glog.FatalDepth(2, args...)
	}
}

func (g *glogger) Println(l grpclog.Severity, args ...interface{}) {
	switch l {
	case grpclog.InfoLog:
		glog.InfoDepth(2, fmt.Sprintln(args...))
	case grpclog.WarningLog:
		glog.WarningDepth(2, fmt.Sprintln(args...))
	case grpclog.ErrorLog:
		glog.ErrorDepth(2, fmt.Sprintln(args...))
	case grpclog.FatalLog:
		glog.FatalDepth(2, fmt.Sprintln(args...))
	}
}

func (g *glogger) Printf(l grpclog.Severity, format string, args ...interface{}) {
	switch l {
	case grpclog.InfoLog:
		glog.InfoDepth(2, fmt.Sprintf(format, args...))
	case grpclog.WarningLog:
		glog.WarningDepth(2, fmt.Sprintf(format, args...))
	case grpclog.ErrorLog:
		glog.ErrorDepth(2, fmt.Sprintf(format, args...))
	case grpclog.FatalLog:
		glog.FatalDepth(2, fmt.Sprintf(format, args...))
	}
}

func (g *glogger) V(l grpclog.VerboseLevel) bool {
	return bool(glog.V(glog.Level(l)))
}
