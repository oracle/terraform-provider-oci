// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cluster Placement Groups API
//
// API for managing cluster placement groups.
//

package clusterplacementgroups

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClusterPlacementGroup A cluster placement group, which is a logical grouping of resources that offer low latency within an availability domain by being placed in close physical proximity.
type ClusterPlacementGroup struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group.
	Id *string `mandatory:"true" json:"id"`

	// The user-friendly name of the cluster placement group. The display name for a cluster placement must be unique and you cannot change it. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// A description of the cluster placement group.
	Description *string `mandatory:"true" json:"description"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the cluster placement group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The availability domain of the cluster placement group.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The type of cluster placement group.
	ClusterPlacementGroupType ClusterPlacementGroupTypeEnum `mandatory:"true" json:"clusterPlacementGroupType"`

	// The time the cluster placement group was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the ClusterPlacementGroup.
	LifecycleState ClusterPlacementGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the cluster placement group was updated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, lifecycle details for a resource in a Failed state might include information to act on.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	PlacementInstruction *PlacementInstructionDetails `mandatory:"false" json:"placementInstruction"`

	Capabilities *CapabilitiesCollection `mandatory:"false" json:"capabilities"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m ClusterPlacementGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClusterPlacementGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingClusterPlacementGroupTypeEnum(string(m.ClusterPlacementGroupType)); !ok && m.ClusterPlacementGroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterPlacementGroupType: %s. Supported values are: %s.", m.ClusterPlacementGroupType, strings.Join(GetClusterPlacementGroupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClusterPlacementGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClusterPlacementGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ClusterPlacementGroupLifecycleStateEnum Enum with underlying type: string
type ClusterPlacementGroupLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterPlacementGroupLifecycleStateEnum
const (
	ClusterPlacementGroupLifecycleStateCreating ClusterPlacementGroupLifecycleStateEnum = "CREATING"
	ClusterPlacementGroupLifecycleStateUpdating ClusterPlacementGroupLifecycleStateEnum = "UPDATING"
	ClusterPlacementGroupLifecycleStateActive   ClusterPlacementGroupLifecycleStateEnum = "ACTIVE"
	ClusterPlacementGroupLifecycleStateInactive ClusterPlacementGroupLifecycleStateEnum = "INACTIVE"
	ClusterPlacementGroupLifecycleStateDeleting ClusterPlacementGroupLifecycleStateEnum = "DELETING"
	ClusterPlacementGroupLifecycleStateDeleted  ClusterPlacementGroupLifecycleStateEnum = "DELETED"
	ClusterPlacementGroupLifecycleStateFailed   ClusterPlacementGroupLifecycleStateEnum = "FAILED"
)

var mappingClusterPlacementGroupLifecycleStateEnum = map[string]ClusterPlacementGroupLifecycleStateEnum{
	"CREATING": ClusterPlacementGroupLifecycleStateCreating,
	"UPDATING": ClusterPlacementGroupLifecycleStateUpdating,
	"ACTIVE":   ClusterPlacementGroupLifecycleStateActive,
	"INACTIVE": ClusterPlacementGroupLifecycleStateInactive,
	"DELETING": ClusterPlacementGroupLifecycleStateDeleting,
	"DELETED":  ClusterPlacementGroupLifecycleStateDeleted,
	"FAILED":   ClusterPlacementGroupLifecycleStateFailed,
}

var mappingClusterPlacementGroupLifecycleStateEnumLowerCase = map[string]ClusterPlacementGroupLifecycleStateEnum{
	"creating": ClusterPlacementGroupLifecycleStateCreating,
	"updating": ClusterPlacementGroupLifecycleStateUpdating,
	"active":   ClusterPlacementGroupLifecycleStateActive,
	"inactive": ClusterPlacementGroupLifecycleStateInactive,
	"deleting": ClusterPlacementGroupLifecycleStateDeleting,
	"deleted":  ClusterPlacementGroupLifecycleStateDeleted,
	"failed":   ClusterPlacementGroupLifecycleStateFailed,
}

// GetClusterPlacementGroupLifecycleStateEnumValues Enumerates the set of values for ClusterPlacementGroupLifecycleStateEnum
func GetClusterPlacementGroupLifecycleStateEnumValues() []ClusterPlacementGroupLifecycleStateEnum {
	values := make([]ClusterPlacementGroupLifecycleStateEnum, 0)
	for _, v := range mappingClusterPlacementGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetClusterPlacementGroupLifecycleStateEnumStringValues Enumerates the set of values in String for ClusterPlacementGroupLifecycleStateEnum
func GetClusterPlacementGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingClusterPlacementGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClusterPlacementGroupLifecycleStateEnum(val string) (ClusterPlacementGroupLifecycleStateEnum, bool) {
	enum, ok := mappingClusterPlacementGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
