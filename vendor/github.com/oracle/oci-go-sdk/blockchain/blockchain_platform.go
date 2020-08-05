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

// BlockchainPlatform Blockchain Platform Instance Description.
type BlockchainPlatform struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Platform Instance Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Role of platform - founder or participant
	PlatformRole BlockchainPlatformPlatformRoleEnum `mandatory:"true" json:"platformRole"`

	// Type of compute shape - one of Standard, (Enterprise) Small, Medium, Large or Extra Large
	ComputeShape BlockchainPlatformComputeShapeEnum `mandatory:"true" json:"computeShape"`

	// Platform Instance Description
	Description *string `mandatory:"false" json:"description"`

	// Bring your own license
	IsByol *bool `mandatory:"false" json:"isByol"`

	// The time the the Platform Instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Platform Instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The version of the Platform Instance.
	ServiceVersion *string `mandatory:"false" json:"serviceVersion"`

	// Service endpoint URL, valid post-provisioning
	ServiceEndpoint *string `mandatory:"false" json:"serviceEndpoint"`

	// The current state of the Platform Instance.
	LifecycleState BlockchainPlatformLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// An message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Storage size in TBs
	StorageSizeInTBs *float64 `mandatory:"false" json:"storageSizeInTBs"`

	// Storage used in TBs
	StorageUsedInTBs *float64 `mandatory:"false" json:"storageUsedInTBs"`

	// True for multi-AD blockchain plaforms, false for single-AD
	IsMultiAD *bool `mandatory:"false" json:"isMultiAD"`

	// Number of total OCPUs allocated to the platform cluster
	TotalOcpuCapacity *int `mandatory:"false" json:"totalOcpuCapacity"`

	ComponentDetails *BlockchainPlatformComponentDetails `mandatory:"false" json:"componentDetails"`

	Replicas *ReplicaDetails `mandatory:"false" json:"replicas"`

	// List of OcpuUtilization for all hosts
	HostOcpuUtilizationInfo []OcpuUtilizationInfo `mandatory:"false" json:"hostOcpuUtilizationInfo"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BlockchainPlatform) String() string {
	return common.PointerString(m)
}

// BlockchainPlatformPlatformRoleEnum Enum with underlying type: string
type BlockchainPlatformPlatformRoleEnum string

// Set of constants representing the allowable values for BlockchainPlatformPlatformRoleEnum
const (
	BlockchainPlatformPlatformRoleFounder     BlockchainPlatformPlatformRoleEnum = "FOUNDER"
	BlockchainPlatformPlatformRoleParticipant BlockchainPlatformPlatformRoleEnum = "PARTICIPANT"
)

var mappingBlockchainPlatformPlatformRole = map[string]BlockchainPlatformPlatformRoleEnum{
	"FOUNDER":     BlockchainPlatformPlatformRoleFounder,
	"PARTICIPANT": BlockchainPlatformPlatformRoleParticipant,
}

// GetBlockchainPlatformPlatformRoleEnumValues Enumerates the set of values for BlockchainPlatformPlatformRoleEnum
func GetBlockchainPlatformPlatformRoleEnumValues() []BlockchainPlatformPlatformRoleEnum {
	values := make([]BlockchainPlatformPlatformRoleEnum, 0)
	for _, v := range mappingBlockchainPlatformPlatformRole {
		values = append(values, v)
	}
	return values
}

// BlockchainPlatformComputeShapeEnum Enum with underlying type: string
type BlockchainPlatformComputeShapeEnum string

// Set of constants representing the allowable values for BlockchainPlatformComputeShapeEnum
const (
	BlockchainPlatformComputeShapeStandard             BlockchainPlatformComputeShapeEnum = "STANDARD"
	BlockchainPlatformComputeShapeEnterpriseSmall      BlockchainPlatformComputeShapeEnum = "ENTERPRISE_SMALL"
	BlockchainPlatformComputeShapeEnterpriseMedium     BlockchainPlatformComputeShapeEnum = "ENTERPRISE_MEDIUM"
	BlockchainPlatformComputeShapeEnterpriseLarge      BlockchainPlatformComputeShapeEnum = "ENTERPRISE_LARGE"
	BlockchainPlatformComputeShapeEnterpriseExtraLarge BlockchainPlatformComputeShapeEnum = "ENTERPRISE_EXTRA_LARGE"
	BlockchainPlatformComputeShapeEnterpriseCustom     BlockchainPlatformComputeShapeEnum = "ENTERPRISE_CUSTOM"
)

var mappingBlockchainPlatformComputeShape = map[string]BlockchainPlatformComputeShapeEnum{
	"STANDARD":               BlockchainPlatformComputeShapeStandard,
	"ENTERPRISE_SMALL":       BlockchainPlatformComputeShapeEnterpriseSmall,
	"ENTERPRISE_MEDIUM":      BlockchainPlatformComputeShapeEnterpriseMedium,
	"ENTERPRISE_LARGE":       BlockchainPlatformComputeShapeEnterpriseLarge,
	"ENTERPRISE_EXTRA_LARGE": BlockchainPlatformComputeShapeEnterpriseExtraLarge,
	"ENTERPRISE_CUSTOM":      BlockchainPlatformComputeShapeEnterpriseCustom,
}

// GetBlockchainPlatformComputeShapeEnumValues Enumerates the set of values for BlockchainPlatformComputeShapeEnum
func GetBlockchainPlatformComputeShapeEnumValues() []BlockchainPlatformComputeShapeEnum {
	values := make([]BlockchainPlatformComputeShapeEnum, 0)
	for _, v := range mappingBlockchainPlatformComputeShape {
		values = append(values, v)
	}
	return values
}

// BlockchainPlatformLifecycleStateEnum Enum with underlying type: string
type BlockchainPlatformLifecycleStateEnum string

// Set of constants representing the allowable values for BlockchainPlatformLifecycleStateEnum
const (
	BlockchainPlatformLifecycleStateCreating BlockchainPlatformLifecycleStateEnum = "CREATING"
	BlockchainPlatformLifecycleStateUpdating BlockchainPlatformLifecycleStateEnum = "UPDATING"
	BlockchainPlatformLifecycleStateActive   BlockchainPlatformLifecycleStateEnum = "ACTIVE"
	BlockchainPlatformLifecycleStateDeleting BlockchainPlatformLifecycleStateEnum = "DELETING"
	BlockchainPlatformLifecycleStateDeleted  BlockchainPlatformLifecycleStateEnum = "DELETED"
	BlockchainPlatformLifecycleStateScaling  BlockchainPlatformLifecycleStateEnum = "SCALING"
	BlockchainPlatformLifecycleStateInactive BlockchainPlatformLifecycleStateEnum = "INACTIVE"
	BlockchainPlatformLifecycleStateFailed   BlockchainPlatformLifecycleStateEnum = "FAILED"
)

var mappingBlockchainPlatformLifecycleState = map[string]BlockchainPlatformLifecycleStateEnum{
	"CREATING": BlockchainPlatformLifecycleStateCreating,
	"UPDATING": BlockchainPlatformLifecycleStateUpdating,
	"ACTIVE":   BlockchainPlatformLifecycleStateActive,
	"DELETING": BlockchainPlatformLifecycleStateDeleting,
	"DELETED":  BlockchainPlatformLifecycleStateDeleted,
	"SCALING":  BlockchainPlatformLifecycleStateScaling,
	"INACTIVE": BlockchainPlatformLifecycleStateInactive,
	"FAILED":   BlockchainPlatformLifecycleStateFailed,
}

// GetBlockchainPlatformLifecycleStateEnumValues Enumerates the set of values for BlockchainPlatformLifecycleStateEnum
func GetBlockchainPlatformLifecycleStateEnumValues() []BlockchainPlatformLifecycleStateEnum {
	values := make([]BlockchainPlatformLifecycleStateEnum, 0)
	for _, v := range mappingBlockchainPlatformLifecycleState {
		values = append(values, v)
	}
	return values
}
