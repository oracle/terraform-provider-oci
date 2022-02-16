// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ScaleBlockchainPlatformDetails Scale operation details for a blockchain platform. The scale operation payload has multiple options
// - Add one or more Ordering Service Node (addOsns)
// - Add one or more Peers (addPeers)
// - Add more replicas of CA, Console and Rest Proxy (addReplicas)
// - Add more storage to the platform (addStorage)
// - Modify the CPU allocation for Peer Nodes (modifyPeers)
// - Remove one or more replicas of CA, Console and Rest Proxy (removeReplicas)
// - Remove one or more Ordering Service Node (removeOsns)
// - Remove one or more Peers (removePeers).
// The scale operation payload must have at least one of the above options.
type ScaleBlockchainPlatformDetails struct {

	// new OSNs to add
	AddOsns []CreateOsnDetails `mandatory:"false" json:"addOsns"`

	AddReplicas *ReplicaDetails `mandatory:"false" json:"addReplicas"`

	// new Peers to add
	AddPeers []CreatePeerDetails `mandatory:"false" json:"addPeers"`

	AddStorage *ScaleStorageDetails `mandatory:"false" json:"addStorage"`

	// modify ocpu allocation to existing Peers
	ModifyPeers []ModifyPeerDetails `mandatory:"false" json:"modifyPeers"`

	RemoveReplicas *ReplicaDetails `mandatory:"false" json:"removeReplicas"`

	// OSN id list to remove
	RemoveOsns []string `mandatory:"false" json:"removeOsns"`

	// Peer id list to remove
	RemovePeers []string `mandatory:"false" json:"removePeers"`
}

func (m ScaleBlockchainPlatformDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScaleBlockchainPlatformDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
