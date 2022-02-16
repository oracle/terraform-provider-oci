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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScaledPlatformMeteringPreview) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
