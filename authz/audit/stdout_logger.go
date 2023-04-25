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

package audit

import (
	"encoding/json"
	"fmt"
)

type StdOutLogger struct{}

func (logger *StdOutLogger) Log(event *Event) error {
	jsonBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal log data to JSON: %v", err)
	}

	fmt.Println(string(jsonBytes))
	return nil
}

func (logger *StdOutLogger) Configure(configJSON string) {
	fmt.Println("StdOutLogger doesn't support any configs")
}

func (logger *StdOutLogger) ToJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal logger to JSON: %v", err)
	}

	return jsonBytes, nil
}
