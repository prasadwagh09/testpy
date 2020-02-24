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

package bootstrap

import "google.golang.org/grpc/grpclog"

const (
	prefix = "[xds-bootstrap] "
)

func debugf(format string, args ...interface{}) {
	if grpclog.V(2) {
		grpclog.Infof(prefix+format, args...)
	}
}

func infof(format string, args ...interface{}) {
	grpclog.Infof(prefix+format, args...)
}

func warningf(format string, args ...interface{}) {
	grpclog.Warningf(prefix+format, args...)
}
