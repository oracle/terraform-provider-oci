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

// BlockchainPlatform Blockchain Platform Instance Description.
type BlockchainPlatform struct {

	// unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Platform Instance Display name, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Role of platform - FOUNDER or PARTICIPANT
	PlatformRole BlockchainPlatformPlatformRoleEnum `mandatory:"true" json:"platformRole"`

	// Compute shape - STANDARD or ENTERPRISE_SMALL or ENTERPRISE_MEDIUM or ENTERPRISE_LARGE or ENTERPRISE_EXTRA_LARGE or ENTERPRISE_CUSTOM
	ComputeShape BlockchainPlatformComputeShapeEnum `mandatory:"true" json:"computeShape"`

	// Platform Instance Description
	Description *string `mandatory:"false" json:"description"`

	// Bring your own license
	IsByol *bool `mandatory:"false" json:"isByol"`

	// The time the the Platform Instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the Platform Instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Platform Version
	PlatformVersion *string `mandatory:"false" json:"platformVersion"`

	// The version of the Platform Instance.
	ServiceVersion *string `mandatory:"false" json:"serviceVersion"`

	// Type of Platform shape - DEFAULT or CUSTOM
	PlatformShapeType BlockchainPlatformPlatformShapeTypeEnum `mandatory:"false" json:"platformShapeType,omitempty"`

	// Type of Load Balancer shape - LB_100_MBPS or LB_400_MBPS. Default is LB_100_MBPS.
	LoadBalancerShape BlockchainPlatformLoadBalancerShapeEnum `mandatory:"false" json:"loadBalancerShape,omitempty"`

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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BlockchainPlatform) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBlockchainPlatformPlatformRoleEnum(string(m.PlatformRole)); !ok && m.PlatformRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformRole: %s. Supported values are: %s.", m.PlatformRole, strings.Join(GetBlockchainPlatformPlatformRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBlockchainPlatformComputeShapeEnum(string(m.ComputeShape)); !ok && m.ComputeShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComputeShape: %s. Supported values are: %s.", m.ComputeShape, strings.Join(GetBlockchainPlatformComputeShapeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBlockchainPlatformPlatformShapeTypeEnum(string(m.PlatformShapeType)); !ok && m.PlatformShapeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformShapeType: %s. Supported values are: %s.", m.PlatformShapeType, strings.Join(GetBlockchainPlatformPlatformShapeTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBlockchainPlatformLoadBalancerShapeEnum(string(m.LoadBalancerShape)); !ok && m.LoadBalancerShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LoadBalancerShape: %s. Supported values are: %s.", m.LoadBalancerShape, strings.Join(GetBlockchainPlatformLoadBalancerShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBlockchainPlatformLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBlockchainPlatformLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BlockchainPlatformPlatformRoleEnum Enum with underlying type: string
type BlockchainPlatformPlatformRoleEnum string

// Set of constants representing the allowable values for BlockchainPlatformPlatformRoleEnum
const (
	BlockchainPlatformPlatformRoleFounder     BlockchainPlatformPlatformRoleEnum = "FOUNDER"
	BlockchainPlatformPlatformRoleParticipant BlockchainPlatformPlatformRoleEnum = "PARTICIPANT"
)

var mappingBlockchainPlatformPlatformRoleEnum = map[string]BlockchainPlatformPlatformRoleEnum{
	"FOUNDER":     BlockchainPlatformPlatformRoleFounder,
	"PARTICIPANT": BlockchainPlatformPlatformRoleParticipant,
}

// GetBlockchainPlatformPlatformRoleEnumValues Enumerates the set of values for BlockchainPlatformPlatformRoleEnum
func GetBlockchainPlatformPlatformRoleEnumValues() []BlockchainPlatformPlatformRoleEnum {
	values := make([]BlockchainPlatformPlatformRoleEnum, 0)
	for _, v := range mappingBlockchainPlatformPlatformRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockchainPlatformPlatformRoleEnumStringValues Enumerates the set of values in String for BlockchainPlatformPlatformRoleEnum
func GetBlockchainPlatformPlatformRoleEnumStringValues() []string {
	return []string{
		"FOUNDER",
		"PARTICIPANT",
	}
}

// GetMappingBlockchainPlatformPlatformRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockchainPlatformPlatformRoleEnum(val string) (BlockchainPlatformPlatformRoleEnum, bool) {
	mappingBlockchainPlatformPlatformRoleEnumIgnoreCase := make(map[string]BlockchainPlatformPlatformRoleEnum)
	for k, v := range mappingBlockchainPlatformPlatformRoleEnum {
		mappingBlockchainPlatformPlatformRoleEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockchainPlatformPlatformRoleEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingBlockchainPlatformComputeShapeEnum = map[string]BlockchainPlatformComputeShapeEnum{
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
	for _, v := range mappingBlockchainPlatformComputeShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockchainPlatformComputeShapeEnumStringValues Enumerates the set of values in String for BlockchainPlatformComputeShapeEnum
func GetBlockchainPlatformComputeShapeEnumStringValues() []string {
	return []string{
		"STANDARD",
		"ENTERPRISE_SMALL",
		"ENTERPRISE_MEDIUM",
		"ENTERPRISE_LARGE",
		"ENTERPRISE_EXTRA_LARGE",
		"ENTERPRISE_CUSTOM",
	}
}

// GetMappingBlockchainPlatformComputeShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockchainPlatformComputeShapeEnum(val string) (BlockchainPlatformComputeShapeEnum, bool) {
	mappingBlockchainPlatformComputeShapeEnumIgnoreCase := make(map[string]BlockchainPlatformComputeShapeEnum)
	for k, v := range mappingBlockchainPlatformComputeShapeEnum {
		mappingBlockchainPlatformComputeShapeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockchainPlatformComputeShapeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// BlockchainPlatformPlatformShapeTypeEnum Enum with underlying type: string
type BlockchainPlatformPlatformShapeTypeEnum string

// Set of constants representing the allowable values for BlockchainPlatformPlatformShapeTypeEnum
const (
	BlockchainPlatformPlatformShapeTypeDefault BlockchainPlatformPlatformShapeTypeEnum = "DEFAULT"
	BlockchainPlatformPlatformShapeTypeCustom  BlockchainPlatformPlatformShapeTypeEnum = "CUSTOM"
)

var mappingBlockchainPlatformPlatformShapeTypeEnum = map[string]BlockchainPlatformPlatformShapeTypeEnum{
	"DEFAULT": BlockchainPlatformPlatformShapeTypeDefault,
	"CUSTOM":  BlockchainPlatformPlatformShapeTypeCustom,
}

// GetBlockchainPlatformPlatformShapeTypeEnumValues Enumerates the set of values for BlockchainPlatformPlatformShapeTypeEnum
func GetBlockchainPlatformPlatformShapeTypeEnumValues() []BlockchainPlatformPlatformShapeTypeEnum {
	values := make([]BlockchainPlatformPlatformShapeTypeEnum, 0)
	for _, v := range mappingBlockchainPlatformPlatformShapeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockchainPlatformPlatformShapeTypeEnumStringValues Enumerates the set of values in String for BlockchainPlatformPlatformShapeTypeEnum
func GetBlockchainPlatformPlatformShapeTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"CUSTOM",
	}
}

// GetMappingBlockchainPlatformPlatformShapeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockchainPlatformPlatformShapeTypeEnum(val string) (BlockchainPlatformPlatformShapeTypeEnum, bool) {
	mappingBlockchainPlatformPlatformShapeTypeEnumIgnoreCase := make(map[string]BlockchainPlatformPlatformShapeTypeEnum)
	for k, v := range mappingBlockchainPlatformPlatformShapeTypeEnum {
		mappingBlockchainPlatformPlatformShapeTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockchainPlatformPlatformShapeTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// BlockchainPlatformLoadBalancerShapeEnum Enum with underlying type: string
type BlockchainPlatformLoadBalancerShapeEnum string

// Set of constants representing the allowable values for BlockchainPlatformLoadBalancerShapeEnum
const (
	BlockchainPlatformLoadBalancerShape100Mbps BlockchainPlatformLoadBalancerShapeEnum = "LB_100_MBPS"
	BlockchainPlatformLoadBalancerShape400Mbps BlockchainPlatformLoadBalancerShapeEnum = "LB_400_MBPS"
)

var mappingBlockchainPlatformLoadBalancerShapeEnum = map[string]BlockchainPlatformLoadBalancerShapeEnum{
	"LB_100_MBPS": BlockchainPlatformLoadBalancerShape100Mbps,
	"LB_400_MBPS": BlockchainPlatformLoadBalancerShape400Mbps,
}

// GetBlockchainPlatformLoadBalancerShapeEnumValues Enumerates the set of values for BlockchainPlatformLoadBalancerShapeEnum
func GetBlockchainPlatformLoadBalancerShapeEnumValues() []BlockchainPlatformLoadBalancerShapeEnum {
	values := make([]BlockchainPlatformLoadBalancerShapeEnum, 0)
	for _, v := range mappingBlockchainPlatformLoadBalancerShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockchainPlatformLoadBalancerShapeEnumStringValues Enumerates the set of values in String for BlockchainPlatformLoadBalancerShapeEnum
func GetBlockchainPlatformLoadBalancerShapeEnumStringValues() []string {
	return []string{
		"LB_100_MBPS",
		"LB_400_MBPS",
	}
}

// GetMappingBlockchainPlatformLoadBalancerShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockchainPlatformLoadBalancerShapeEnum(val string) (BlockchainPlatformLoadBalancerShapeEnum, bool) {
	mappingBlockchainPlatformLoadBalancerShapeEnumIgnoreCase := make(map[string]BlockchainPlatformLoadBalancerShapeEnum)
	for k, v := range mappingBlockchainPlatformLoadBalancerShapeEnum {
		mappingBlockchainPlatformLoadBalancerShapeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockchainPlatformLoadBalancerShapeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingBlockchainPlatformLifecycleStateEnum = map[string]BlockchainPlatformLifecycleStateEnum{
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
	for _, v := range mappingBlockchainPlatformLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBlockchainPlatformLifecycleStateEnumStringValues Enumerates the set of values in String for BlockchainPlatformLifecycleStateEnum
func GetBlockchainPlatformLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"SCALING",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingBlockchainPlatformLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlockchainPlatformLifecycleStateEnum(val string) (BlockchainPlatformLifecycleStateEnum, bool) {
	mappingBlockchainPlatformLifecycleStateEnumIgnoreCase := make(map[string]BlockchainPlatformLifecycleStateEnum)
	for k, v := range mappingBlockchainPlatformLifecycleStateEnum {
		mappingBlockchainPlatformLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingBlockchainPlatformLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
