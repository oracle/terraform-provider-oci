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

// ScaledBlockchainPlatformPreview Blockchain Platform Instance Description Preview after Scaling.
type ScaledBlockchainPlatformPreview struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Platform Instance Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of compute shape - one of Standard, (Enterprise) Small, Medium, Large or Extra Large
	ComputeShape *string `mandatory:"true" json:"computeShape"`

	// Platform Instance Description
	Description *string `mandatory:"false" json:"description"`

	// Storage size in TBs
	StorageSizeInTBs *float64 `mandatory:"false" json:"storageSizeInTBs"`

	// Storage size in TBs post scaling
	StorageSizeInTBsPostScaling *float64 `mandatory:"false" json:"storageSizeInTBsPostScaling"`

	ComponentDetails *BlockchainPlatformComponentDetails `mandatory:"false" json:"componentDetails"`

	Replicas *ReplicaDetails `mandatory:"false" json:"replicas"`

	ComponentDetailsPostScaling *BlockchainPlatformComponentDetails `mandatory:"false" json:"componentDetailsPostScaling"`

	ReplicasPostScaling *ReplicaDetails `mandatory:"false" json:"replicasPostScaling"`

	// List of OcpuUtilization for all hosts
	HostOcpuUtilizationInfo []OcpuUtilizationInfo `mandatory:"false" json:"hostOcpuUtilizationInfo"`

	// List of OcpuUtilization for all hosts after scaling
	HostOcpuUtilizationInfoPostScaling []OcpuUtilizationInfo `mandatory:"false" json:"hostOcpuUtilizationInfoPostScaling"`

	// Number of new VMs that would be created
	NewVmCount *int `mandatory:"false" json:"newVmCount"`

	MeteringPreview *ScaledPlatformMeteringPreview `mandatory:"false" json:"meteringPreview"`

	ScalePayload *ScaleBlockchainPlatformDetails `mandatory:"false" json:"scalePayload"`
}

func (m ScaledBlockchainPlatformPreview) String() string {
	return common.PointerString(m)
}
