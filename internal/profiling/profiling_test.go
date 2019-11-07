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
	"fmt"
	"sync"
	"testing"
)

func BenchmarkTimer(b *testing.B) {
	for routines := 1; routines <= 1<<8; routines <<= 1 {
		b.Run(fmt.Sprintf("goroutines:%d", routines), func(b *testing.B) {
			stat := NewStat("foo")
			perRoutine := b.N / routines
			var wg sync.WaitGroup
			for r := 0; r < routines; r++ {
				wg.Add(1)
				go func() {
					for i := 0; i < perRoutine; i++ {
						timer := stat.NewTimer("bar")
						timer.Egress()
					}
					wg.Done()
				}()
			}
			wg.Wait()
		})
	}
}
