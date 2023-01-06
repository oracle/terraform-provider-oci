// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// HeatWaveCluster A HeatWave cluster is a database accelerator for a DB System.
type HeatWaveCluster struct {

	// The OCID of the parent DB System this HeatWave cluster is attached to.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The shape determines resources to allocate to the HeatWave
	// nodes - CPU cores, memory.
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// The number of analytics-processing compute instances, of the
	// specified shape, in the HeatWave cluster.
	ClusterSize *int `mandatory:"true" json:"clusterSize"`

	// A HeatWave node is a compute host that is part of a HeatWave cluster.
	ClusterNodes []HeatWaveNode `mandatory:"true" json:"clusterNodes"`

	// The current state of the HeatWave cluster.
	LifecycleState HeatWaveClusterLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the HeatWave cluster was created,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the HeatWave cluster was last updated,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Additional information about the current lifecycleState.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m HeatWaveCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHeatWaveClusterLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHeatWaveClusterLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HeatWaveClusterLifecycleStateEnum Enum with underlying type: string
type HeatWaveClusterLifecycleStateEnum string

// Set of constants representing the allowable values for HeatWaveClusterLifecycleStateEnum
const (
	HeatWaveClusterLifecycleStateCreating HeatWaveClusterLifecycleStateEnum = "CREATING"
	HeatWaveClusterLifecycleStateActive   HeatWaveClusterLifecycleStateEnum = "ACTIVE"
	HeatWaveClusterLifecycleStateInactive HeatWaveClusterLifecycleStateEnum = "INACTIVE"
	HeatWaveClusterLifecycleStateUpdating HeatWaveClusterLifecycleStateEnum = "UPDATING"
	HeatWaveClusterLifecycleStateDeleting HeatWaveClusterLifecycleStateEnum = "DELETING"
	HeatWaveClusterLifecycleStateDeleted  HeatWaveClusterLifecycleStateEnum = "DELETED"
	HeatWaveClusterLifecycleStateFailed   HeatWaveClusterLifecycleStateEnum = "FAILED"
)

var mappingHeatWaveClusterLifecycleStateEnum = map[string]HeatWaveClusterLifecycleStateEnum{
	"CREATING": HeatWaveClusterLifecycleStateCreating,
	"ACTIVE":   HeatWaveClusterLifecycleStateActive,
	"INACTIVE": HeatWaveClusterLifecycleStateInactive,
	"UPDATING": HeatWaveClusterLifecycleStateUpdating,
	"DELETING": HeatWaveClusterLifecycleStateDeleting,
	"DELETED":  HeatWaveClusterLifecycleStateDeleted,
	"FAILED":   HeatWaveClusterLifecycleStateFailed,
}

var mappingHeatWaveClusterLifecycleStateEnumLowerCase = map[string]HeatWaveClusterLifecycleStateEnum{
	"creating": HeatWaveClusterLifecycleStateCreating,
	"active":   HeatWaveClusterLifecycleStateActive,
	"inactive": HeatWaveClusterLifecycleStateInactive,
	"updating": HeatWaveClusterLifecycleStateUpdating,
	"deleting": HeatWaveClusterLifecycleStateDeleting,
	"deleted":  HeatWaveClusterLifecycleStateDeleted,
	"failed":   HeatWaveClusterLifecycleStateFailed,
}

// GetHeatWaveClusterLifecycleStateEnumValues Enumerates the set of values for HeatWaveClusterLifecycleStateEnum
func GetHeatWaveClusterLifecycleStateEnumValues() []HeatWaveClusterLifecycleStateEnum {
	values := make([]HeatWaveClusterLifecycleStateEnum, 0)
	for _, v := range mappingHeatWaveClusterLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHeatWaveClusterLifecycleStateEnumStringValues Enumerates the set of values in String for HeatWaveClusterLifecycleStateEnum
func GetHeatWaveClusterLifecycleStateEnumStringValues() []string {
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

// GetMappingHeatWaveClusterLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeatWaveClusterLifecycleStateEnum(val string) (HeatWaveClusterLifecycleStateEnum, bool) {
	enum, ok := mappingHeatWaveClusterLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
