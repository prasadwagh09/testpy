// +build grpcgoid

/*
 *
 * Copyright 2019 gRPC authors.
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

package profiling

import (
	"runtime"
)

// This stubbed function usually returns zero (see goid_regular.go); however,
// if grpc is built with `-tags 'grpcgoid'`, a runtime.GoId function, which
// does not exist in the Go standard library, is expected. While not necessary,
// sometimes, visualising grpc profiling data in trace-viewer is much nicer
// with goroutines separated from each other.
//
// Several other approaches were considered before arriving at this:
//
// 1. Using a CGO module: CGO usually has access to some things that regular
//    Go does not. Till go1.4, CGO used to have access to the goroutine struct
//    because the Go runtime was written in C. However, 1.5+ uses a native Go
//    runtime; as a result, CGO does not have access to the goroutine structure
//    anymore in modern Go. Besides, CGO interop wasn't fast enough (estimated
//    to be ~170ns/op). This would also make building grpc require a C
//    compiler, which isn't a requirement currently, breaking a lot of stuff.
//
// 2. Using runtime.Stack stacktrace: While this would remove the need for a
//    modified Go runtime, this is ridiculously slow, thanks to the all the
//    string processing shenanigans required to extract the goroutine ID (about
//    ~2000ns/op).
//
// 3. Using Go version-specific build tags: For any given Go version, the
//    goroutine struct has a fixed structure. As a result, the goroutine ID
//    could be extracted if we know the offset using some assembly. This would
//    be faster then #1 and #2, but is harder to maintain. This would require
//    special Go code that's both architecture-specific and go version-specific
//    (a quadratic number of variants to maintain).
//
// 4. This approach, which requires a simple two-line modification to the Go
//    runtime to expose the current goroutine's ID. This is the chosen approach
//    and it takes about ~2 ns/op, which is negligible in the face of the tens
//    of microseconds that grpc takes to complete a RPC request.
//
// A personal note (adtac@): I'm fully aware of the Go team's position on why
// goroutine ID isn't exposed and I agree with most of it. However, this is
// used exclusively for profiling data visualisation, which qualifies as
// debugging; given that runtime.Stack is fine with exposing the goroutine ID
// (in an inefficient way), I'm sure the Go team would be comfortable with this
// conditional build tag.
func goId() int64 {
	return runtime.GoId()
}
