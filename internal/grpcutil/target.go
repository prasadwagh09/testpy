/*
 *
 * Copyright 2020 gRPC authors.
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

// Package grpcutil provides a bunch of utility functions to be used across the
// gRPC codebase.
package grpcutil

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

// split2 returns the values from strings.SplitN(s, sep, 2).
// If sep is not found, it returns ("", "", false) instead.
func split2(s, sep string) (string, string, bool) {
	spl := strings.SplitN(s, sep, 2)
	if len(spl) < 2 {
		return "", "", false
	}
	return spl[0], spl[1], true
}

// ParseTarget splits target into a resolver.Target struct containing scheme,
// authority and endpoint. skipUnixColonParsing indicates that the parse should
// not parse "unix:[path]" cases. This should be true in cases where a custom
// dialer is present, to prevent a behavior change.
//
// If target is not a valid scheme://authority/endpoint, it returns {Endpoint:
// target}.
func ParseTarget(target string, skipUnixColonParsing bool) (ret resolver.Target) {
	var ok bool
	ret.Scheme, ret.Endpoint, ok = split2(target, "://")
	if !ok {
		if strings.HasPrefix(target, "unix:") && !skipUnixColonParsing {
			// Handle the "unix:[path]" case, because splitting on :// only
			// handles the "unix://[/absolute/path]" case. Only handle if the
			// dialer is nil, to avoid a behavior change with custom dialers.
			return resolver.Target{Scheme: "unix", Endpoint: target[len("unix:"):]}
		}
		return resolver.Target{Endpoint: target}
	}
	ret.Authority, ret.Endpoint, ok = split2(ret.Endpoint, "/")
	if !ok {
		return resolver.Target{Endpoint: target}
	}
	if ret.Scheme == "unix" {
		// Prevents behavior change in "unix:///[...]" case.
		if skipUnixColonParsing && ret.Authority == "" {
			return resolver.Target{Endpoint: target}
		}
		// Add the "/" back in the unix case, so the unix resolver receives the
		// actual endpoint.
		ret.Endpoint = "/" + ret.Endpoint
	}
	// Prevents behavior change in "passthrough:///unix:///a/b/c" case.
	if !skipUnixColonParsing && ret.Scheme == "passthrough" && strings.HasPrefix(ret.Endpoint, "unix:") {
		return ParseTarget(ret.Endpoint, false)
	}
	return ret
}
