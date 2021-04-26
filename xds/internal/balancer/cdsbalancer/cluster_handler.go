/*
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

package cdsbalancer

import (
	"errors"
	"sync"

	xdsclient "google.golang.org/grpc/xds/internal/client"
)


// clusterHandler will be given a name representing a cluster. It will then update the CDS policy
// constantly with a list of Clusters to pass down to XdsClusterResolverLoadBalancingPolicyConfig
// in a stream like fashion.
type clusterHandler struct {
	// A mutex to protect entire tree of clusters.
	clusterMutex    sync.Mutex
	root            *clusterNode
	rootClusterName string

	// A way to ping CDS Balanccer about any updates or errors to a Node in the tree.
	// This will either get called from this handler constructing an update or from a child with an error.
	// Capacity of one as the only update CDS Balancer cares about is the most recent update.
	updateChannel chan clusterHandlerUpdate

	xdsClient      xdsClientInterface
}

// He has something called WatchCluster which takes xdsClientInterface and also a callback.

func (ch *clusterHandler) updateRootCluster(rootClusterName string) {
	ch.clusterMutex.Lock()
	defer ch.clusterMutex.Unlock()
	if ch.root == nil {
		// Construct a root node on first update.
		ch.root = createClusterNode(rootClusterName, ch.xdsClient, ch)
		ch.rootClusterName = rootClusterName
		return
	}
	// Check if root cluster was changed. If it was, delete old one and start new one, if not
	// do nothing.
	if rootClusterName != ch.rootClusterName {
		ch.root.delete()
		ch.root = createClusterNode(rootClusterName, ch.xdsClient, ch)
		ch.rootClusterName = rootClusterName
	}
}


// This function tries to construct a cluster update to send to CDS.
func (ch *clusterHandler) constructClusterUpdate() {
	// If there was an error received no op, as this simply means one of the children hasn't received an update yet.
	if clusterUpdate, err := ch.root.constructClusterUpdate(); err == nil {
		// For a ClusterUpdate, the only update CDS cares about is the most recent one, so opportunistically drain the
		// update channel before sending the new update.
		select {
		case <-ch.updateChannel:
		default:
		}
		ch.updateChannel<-clusterHandlerUpdate{chu: clusterUpdate, err: nil}
	}
}

// This logically represents a cluster. This handles all the logic for starting and stopping
// a cluster watch, handling any updates, and constructing a list recursively for the ClusterHandler.
type clusterNode struct {
	// A way to cancel the watch for the cluster.
	cancelFunc func()

	// A list of children, as the Node can be an aggregate Cluster.
	children []*clusterNode

	// A ClusterUpdate in order to build a list of cluster updates for CDS to send down to child
	// XdsClusterResolverLoadBalancingPolicy.
	clusterUpdate xdsclient.ClusterUpdate

	// This boolean determines whether this Node has received an update or not. This isn't the best practice,
	// but this will protect a list of Cluster Updates from being constructed if a cluster in the tree has not received
	// an update yet.
	receivedUpdate bool

	clusterHandler *clusterHandler

}

// CreateClusterNode creates a cluster node from a given clusterName. This will also start the watch for that cluster.
func createClusterNode(clusterName string, xdsClient xdsClientInterface, topLevelHandler *clusterHandler) *clusterNode {
	print("|createClusterNode() called on ", clusterName, "|")
	c := &clusterNode{
		clusterHandler: topLevelHandler,
	}
	print("About to communicate with the xds client for:", clusterName, "|")
	// Communicate with the xds client here.
	c.cancelFunc = xdsClient.WatchCluster(clusterName, c.handleResp)
	return c
}

// This function cancels the cluster watch on the cluster and all of it's children.
func (c *clusterNode) delete() {
	c.cancelFunc()
	for _, child := range c.children {
		child.delete()
	}
}

// Construct cluster update (potentially a list of ClusterUpdates) for a node.
func (c *clusterNode) constructClusterUpdate() ([]xdsclient.ClusterUpdate, error) { // THIS LOGIC IS WRONG
	print("Tries constructing a cluster update for Cluster ", c.clusterUpdate.ServiceName, ". ")
	// If the cluster has not yet received an update, the cluster update is not yet ready.
	if !c.receivedUpdate {
		return nil, errors.New("Tried to construct a cluster update on a cluster that has not received an update.")
	}

	// Base case - LogicalDNS or EDS. Both of these cluster types will be tied to a single ClusterUpdate.
	if c.clusterUpdate.ClusterType != xdsclient.ClusterTypeAggregate {
		print("Returning the cluster update for service (base node):", c.clusterUpdate.ServiceName)
		return []xdsclient.ClusterUpdate{c.clusterUpdate}, nil
	}

	// If an aggregate construct a list by recursively calling down to all of it's children.
	var childrenUpdates []xdsclient.ClusterUpdate
	for _, child := range c.children {
		childUpdateList, err := child.constructClusterUpdate()
		if err != nil {
			return nil, err
		}
		childrenUpdates = append(childrenUpdates, childUpdateList...)
	}
	print("Returning the update list for an aggregate cluster")
	return childrenUpdates, nil
}

// handleResp handles a xds response for a particular cluster. This function also
// handles any logic with regards to any child state that may have changed.
// Problem statement: perhaps this logic is wrong
func (c *clusterNode) handleResp(clusterUpdate xdsclient.ClusterUpdate, err error) {
	c.clusterHandler.clusterMutex.Lock()
	defer c.clusterHandler.clusterMutex.Unlock()
	print("SERVICE NAME:", clusterUpdate.ServiceName)
	if err != nil { // Write this error for run() to pick up in CDS LB policy.
		c.clusterHandler.updateChannel<-clusterHandlerUpdate{chu: nil, err: err}
		return
	}
	// This variable will determine whether there was a delta with regards to this clusterupdate. If there was, at the end
	// of the response ping ClusterHandler to send CDS Policy a new list of ClusterUpdates.
	var delta bool
	// If this is the first update to the Cluster Node, ping ClusterHandler at the end of method to try and build a config.
	if !c.receivedUpdate {
		delta = true
	}

	c.receivedUpdate = true
	c.clusterUpdate = clusterUpdate

	// This map will be empty if the cluster update specifies cluster is an EDS or LogicalDNS cluster, as will have no children.
	newChildren := make(map[string]struct{})
	if clusterUpdate.ClusterType == xdsclient.ClusterTypeAggregate {
		for _, childName := range clusterUpdate.PrioritizedClusterNames {
			print("New children added: ", childName)
			newChildren[childName] = struct{}{}
		}
	}

	for _, child := range c.children {
		// If the child is still present in the update, then there is nothing to do for that child name in the update.
		if _, found := newChildren[child.clusterUpdate.ServiceName]; found {
			print("New children deleted: ", child.clusterUpdate.ServiceName)
			delete(newChildren, child.clusterUpdate.ServiceName)
		} else { // If the child is no longer present in the update, that cluster can be deleted.
			delta = true
			child.delete()
		}
	}

	// Whatever clusters are left over here from the update are all new children, so create CDS watches for those clusters.
	for child, _ := range newChildren {
		print("Creating cds watch in:", child)
		delta = true
		c.children = append(c.children, createClusterNode(child, c.clusterHandler.xdsClient, c.clusterHandler))
	}
	// If there was a change in the state of the children, ping the ClusterHandler to try and construct a new update to send back
	// to CDS.
	if delta {
		c.clusterHandler.constructClusterUpdate()
	}
}
