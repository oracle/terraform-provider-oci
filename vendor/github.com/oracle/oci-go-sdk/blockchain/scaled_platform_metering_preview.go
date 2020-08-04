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

// ScaledPlatformMeteringPreview Blockchain Platform Metering Preview after Scaling
type ScaledPlatformMeteringPreview struct {

	// Number of total OCPU allocation for the blockchain platform
	TotalOcpuAllocation *float32 `mandatory:"false" json:"totalOcpuAllocation"`

	// Number of total OCPU allocation for the blockchain platform after Scaling
	TotalOcpuAllocationPostScaling *float32 `mandatory:"false" json:"totalOcpuAllocationPostScaling"`

	// Current Storage metered units in TBs
	StorageMeteredUnits *float64 `mandatory:"false" json:"storageMeteredUnits"`

	// Extra Storage units required in TBs
	ExtraStorageMeteredUnits *float64 `mandatory:"false" json:"extraStorageMeteredUnits"`

	// Total Post Scaling Storage metered units in TBs
	StorageMeteredUnitsPostScaling *float64 `mandatory:"false" json:"storageMeteredUnitsPostScaling"`
}

func (m ScaledPlatformMeteringPreview) String() string {
	return common.PointerString(m)
}
