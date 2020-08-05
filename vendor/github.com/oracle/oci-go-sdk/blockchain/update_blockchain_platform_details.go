// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateBlockchainPlatformDetails Blockchain Platform details for updating a service.
type UpdateBlockchainPlatformDetails struct {

	// Platform Description
	Description *string `mandatory:"false" json:"description"`

	// Storage size in TBs
	StorageSizeInTBs *float64 `mandatory:"false" json:"storageSizeInTBs"`

	Replicas *ReplicaDetails `mandatory:"false" json:"replicas"`

	// Number of total OCPUs to allocate
	TotalOcpuCapacity *int `mandatory:"false" json:"totalOcpuCapacity"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateBlockchainPlatformDetails) String() string {
	return common.PointerString(m)
}
