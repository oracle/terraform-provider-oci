// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnalyticsCluster DEPRECATED -- please use HeatWave API instead.
// An Analytics Cluster is a database accelerator for a DB System.
type AnalyticsCluster struct {

	// The OCID of the parent DB System this Analytics Cluster is attached to.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The shape determines resources to allocate to the Analytics
	// Cluster nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing compute instances, of the
	// specified shape, in the Analytics Cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// An Analytics Cluster Node is a compute host that is part of an Analytics Cluster.
	ClusterNodes []AnalyticsClusterNode `mandatory:"true" json:"clusterNodes"`

	// The current state of the Analytics Cluster.
	LifecycleState AnalyticsClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the Analytics Cluster was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the Analytics Cluster was last updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AnalyticsCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnalyticsCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAnalyticsClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAnalyticsClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AnalyticsClusterLifecycleStateEnum Enum with underlying type: string
type AnalyticsClusterLifecycleStateEnum string

// Set of constants representing the allowable values for AnalyticsClusterLifecycleStateEnum
const (
	AnalyticsClusterLifecycleStateCreating AnalyticsClusterLifecycleStateEnum = "CREATING"
	AnalyticsClusterLifecycleStateActive   AnalyticsClusterLifecycleStateEnum = "ACTIVE"
	AnalyticsClusterLifecycleStateInactive AnalyticsClusterLifecycleStateEnum = "INACTIVE"
	AnalyticsClusterLifecycleStateUpdating AnalyticsClusterLifecycleStateEnum = "UPDATING"
	AnalyticsClusterLifecycleStateDeleting AnalyticsClusterLifecycleStateEnum = "DELETING"
	AnalyticsClusterLifecycleStateDeleted  AnalyticsClusterLifecycleStateEnum = "DELETED"
	AnalyticsClusterLifecycleStateFailed   AnalyticsClusterLifecycleStateEnum = "FAILED"
)

var mappingAnalyticsClusterLifecycleStateEnum = map[string]AnalyticsClusterLifecycleStateEnum{
	"CREATING": AnalyticsClusterLifecycleStateCreating,
	"ACTIVE":   AnalyticsClusterLifecycleStateActive,
	"INACTIVE": AnalyticsClusterLifecycleStateInactive,
	"UPDATING": AnalyticsClusterLifecycleStateUpdating,
	"DELETING": AnalyticsClusterLifecycleStateDeleting,
	"DELETED":  AnalyticsClusterLifecycleStateDeleted,
	"FAILED":   AnalyticsClusterLifecycleStateFailed,
}

var mappingAnalyticsClusterLifecycleStateEnumLowerCase = map[string]AnalyticsClusterLifecycleStateEnum{
	"creating": AnalyticsClusterLifecycleStateCreating,
	"active":   AnalyticsClusterLifecycleStateActive,
	"inactive": AnalyticsClusterLifecycleStateInactive,
	"updating": AnalyticsClusterLifecycleStateUpdating,
	"deleting": AnalyticsClusterLifecycleStateDeleting,
	"deleted":  AnalyticsClusterLifecycleStateDeleted,
	"failed":   AnalyticsClusterLifecycleStateFailed,
}

// GetAnalyticsClusterLifecycleStateEnumValues Enumerates the set of values for AnalyticsClusterLifecycleStateEnum
func GetAnalyticsClusterLifecycleStateEnumValues() []AnalyticsClusterLifecycleStateEnum {
	values := make([]AnalyticsClusterLifecycleStateEnum, 0)
	for _, v := range mappingAnalyticsClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAnalyticsClusterLifecycleStateEnumStringValues Enumerates the set of values in String for AnalyticsClusterLifecycleStateEnum
func GetAnalyticsClusterLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingAnalyticsClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAnalyticsClusterLifecycleStateEnum(val string) (AnalyticsClusterLifecycleStateEnum, bool) {
	enum, ok := mappingAnalyticsClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
