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

// Package defaults is the API to set global default StatsHandler, which will be
// applied to clients and servers created after. This package is for monitoring
// purpose only. All fields are read-only. All APIs are experimental.
package defaults // import "google.golang.org/grpc/stats/defaults"

import (
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/stats"
)

var (
	defaultClientStatsHandler stats.Handler
	defaultServerStatsHandler stats.Handler
	logger                    = grpclog.Component("default-statshandler")
)

// SetDefaultClientStatsHandler sets the default StatsHandler for all clients
// that created after this call.
func SetDefaultClientStatsHandler(handler stats.Handler) {
	defaultClientStatsHandler = handler
}

// GetDefaultClientStatsHandler gets the pointer of the default client
// StatsHandler. Should only be used internally.
func GetDefaultClientStatsHandler() stats.Handler {
	if defaultClientStatsHandler != nil {
		logger.Infof("default client StatsHandler injected")
	}
	return defaultClientStatsHandler
}

// SetDefaultServerStatsHandler sets the default StatsHandler for all servers
// that created after this call.
func SetDefaultServerStatsHandler(handler stats.Handler) {
	defaultServerStatsHandler = handler
}

// GetDefaultServerStatsHandler gets the pointer of the default server
// StatsHandler. Should only be used internally.
func GetDefaultServerStatsHandler() stats.Handler {
	if defaultServerStatsHandler != nil {
		logger.Infof("default server StatsHandler injected")
	}
	return defaultServerStatsHandler
}
