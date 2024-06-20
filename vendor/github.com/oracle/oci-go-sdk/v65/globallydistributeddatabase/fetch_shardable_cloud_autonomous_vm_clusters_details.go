// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FetchShardableCloudAutonomousVmClustersDetails Details required for fetch sharded cloud autonomous vm clusters.
type FetchShardableCloudAutonomousVmClustersDetails struct {

	// Compartment id of cloud autonomous vm clusters.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Lifecycle states for shardable Cloud autonomous vm cluster.
	LifecycleState FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Detailed message for the lifecycle state.
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// List only clusters for which atleast given minimum CDB count is available.
	MinimumAvailableCdbCount *int `mandatory:"false" json:"minimumAvailableCdbCount"`

	// Flag to indicate of response shall also include clusters for which no more capacity is left to create new resources.
	AreDepletedClustersIncluded *bool `mandatory:"false" json:"areDepletedClustersIncluded"`

	// Region code of regions for which sharded cloud autonomous vm clusters need to be fetched.
	Regions []string `mandatory:"false" json:"regions"`
}

func (m FetchShardableCloudAutonomousVmClustersDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FetchShardableCloudAutonomousVmClustersDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum Enum with underlying type: string
type FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum string

// Set of constants representing the allowable values for FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum
const (
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateActive         FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "ACTIVE"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateFailed         FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "FAILED"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateNeedsAttention FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "NEEDS_ATTENTION"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateInactive       FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "INACTIVE"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleting       FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "DELETING"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleted        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "DELETED"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUpdating       FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "UPDATING"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateCreating       FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "CREATING"
	FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUnavailable    FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = "UNAVAILABLE"
)

var mappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum = map[string]FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum{
	"ACTIVE":          FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateActive,
	"FAILED":          FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateFailed,
	"NEEDS_ATTENTION": FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateNeedsAttention,
	"INACTIVE":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateInactive,
	"DELETING":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleting,
	"DELETED":         FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleted,
	"UPDATING":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUpdating,
	"CREATING":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateCreating,
	"UNAVAILABLE":     FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUnavailable,
}

var mappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumLowerCase = map[string]FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum{
	"active":          FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateActive,
	"failed":          FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateFailed,
	"needs_attention": FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateNeedsAttention,
	"inactive":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateInactive,
	"deleting":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleting,
	"deleted":         FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateDeleted,
	"updating":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUpdating,
	"creating":        FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateCreating,
	"unavailable":     FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateUnavailable,
}

// GetFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumValues Enumerates the set of values for FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum
func GetFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumValues() []FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum {
	values := make([]FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum, 0)
	for _, v := range mappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumStringValues Enumerates the set of values in String for FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum
func GetFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"UNAVAILABLE",
	}
}

// GetMappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum(val string) (FetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnum, bool) {
	enum, ok := mappingFetchShardableCloudAutonomousVmClustersDetailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
