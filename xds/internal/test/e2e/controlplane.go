/*
 *
 * Copyright 2021 gRPC authors.
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
 */

package e2e

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	xdsinternal "google.golang.org/grpc/internal/xds"
	"google.golang.org/grpc/xds/internal/testutils/e2e"
)

const configJSONPath = "./configs/bootstrap-%s.json"

type controlPlane struct {
	server *e2e.ManagementServer
	nodeID string
}

func newControlPlane(testName string) (*controlPlane, error) {
	// Spin up an xDS management server on a local port.
	server, err := e2e.StartManagementServer()
	if err != nil {
		return nil, fmt.Errorf("failed to spin up the xDS management server: %v", err)
	}

	// Create a bootstrap file in a temporary directory.
	nodeID := uuid.New().String()
	bootstrapContents, err := xdsinternal.BootstrapContents(xdsinternal.BootstrapOptions{
		Version:                            xdsinternal.TransportV3,
		NodeID:                             nodeID,
		ServerURI:                          server.Address,
		ServerListenerResourceNameTemplate: e2e.ServerListenerResourceNameTemplate,
	})
	if err != nil {
		server.Stop()
		return nil, fmt.Errorf("failed to create bootstrap file: %v", err)
	}

	jsonPath := fmt.Sprintf(configJSONPath, testName)
	dir := filepath.Dir(jsonPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}

	if err := ioutil.WriteFile(jsonPath, bootstrapContents, os.ModePerm); err != nil {
		return nil, fmt.Errorf("ioutil.WriteFile(bootstrap.json) failed: %v", err)
	}

	return &controlPlane{
		server: server,
		nodeID: nodeID,
	}, nil
}

func (cp *controlPlane) stop() {
	cp.server.Stop()

	// TODO: maybe? os.RemoveAll(cp.tempDir)
}
