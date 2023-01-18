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

// HeatWaveNode A HeatWave node is a compute host that is part of a HeatWave cluster.
type HeatWaveNode struct {

	// The ID of the node within MySQL HeatWave cluster.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// The current state of the MySQL HeatWave node.
	LifecycleState HeatWaveNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MySQL HeatWave node was created,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the MySQL HeatWave node was updated,
	// as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m HeatWaveNode) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HeatWaveNode) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHeatWaveNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetHeatWaveNodeLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HeatWaveNodeLifecycleStateEnum Enum with underlying type: string
type HeatWaveNodeLifecycleStateEnum string

// Set of constants representing the allowable values for HeatWaveNodeLifecycleStateEnum
const (
	HeatWaveNodeLifecycleStateCreating HeatWaveNodeLifecycleStateEnum = "CREATING"
	HeatWaveNodeLifecycleStateActive   HeatWaveNodeLifecycleStateEnum = "ACTIVE"
	HeatWaveNodeLifecycleStateInactive HeatWaveNodeLifecycleStateEnum = "INACTIVE"
	HeatWaveNodeLifecycleStateUpdating HeatWaveNodeLifecycleStateEnum = "UPDATING"
	HeatWaveNodeLifecycleStateDeleting HeatWaveNodeLifecycleStateEnum = "DELETING"
	HeatWaveNodeLifecycleStateDeleted  HeatWaveNodeLifecycleStateEnum = "DELETED"
	HeatWaveNodeLifecycleStateFailed   HeatWaveNodeLifecycleStateEnum = "FAILED"
)

var mappingHeatWaveNodeLifecycleStateEnum = map[string]HeatWaveNodeLifecycleStateEnum{
	"CREATING": HeatWaveNodeLifecycleStateCreating,
	"ACTIVE":   HeatWaveNodeLifecycleStateActive,
	"INACTIVE": HeatWaveNodeLifecycleStateInactive,
	"UPDATING": HeatWaveNodeLifecycleStateUpdating,
	"DELETING": HeatWaveNodeLifecycleStateDeleting,
	"DELETED":  HeatWaveNodeLifecycleStateDeleted,
	"FAILED":   HeatWaveNodeLifecycleStateFailed,
}

var mappingHeatWaveNodeLifecycleStateEnumLowerCase = map[string]HeatWaveNodeLifecycleStateEnum{
	"creating": HeatWaveNodeLifecycleStateCreating,
	"active":   HeatWaveNodeLifecycleStateActive,
	"inactive": HeatWaveNodeLifecycleStateInactive,
	"updating": HeatWaveNodeLifecycleStateUpdating,
	"deleting": HeatWaveNodeLifecycleStateDeleting,
	"deleted":  HeatWaveNodeLifecycleStateDeleted,
	"failed":   HeatWaveNodeLifecycleStateFailed,
}

// GetHeatWaveNodeLifecycleStateEnumValues Enumerates the set of values for HeatWaveNodeLifecycleStateEnum
func GetHeatWaveNodeLifecycleStateEnumValues() []HeatWaveNodeLifecycleStateEnum {
	values := make([]HeatWaveNodeLifecycleStateEnum, 0)
	for _, v := range mappingHeatWaveNodeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHeatWaveNodeLifecycleStateEnumStringValues Enumerates the set of values in String for HeatWaveNodeLifecycleStateEnum
func GetHeatWaveNodeLifecycleStateEnumStringValues() []string {
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

// GetMappingHeatWaveNodeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeatWaveNodeLifecycleStateEnum(val string) (HeatWaveNodeLifecycleStateEnum, bool) {
	enum, ok := mappingHeatWaveNodeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
