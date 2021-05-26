// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
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

var mappingHeatWaveNodeLifecycleState = map[string]HeatWaveNodeLifecycleStateEnum{
	"CREATING": HeatWaveNodeLifecycleStateCreating,
	"ACTIVE":   HeatWaveNodeLifecycleStateActive,
	"INACTIVE": HeatWaveNodeLifecycleStateInactive,
	"UPDATING": HeatWaveNodeLifecycleStateUpdating,
	"DELETING": HeatWaveNodeLifecycleStateDeleting,
	"DELETED":  HeatWaveNodeLifecycleStateDeleted,
	"FAILED":   HeatWaveNodeLifecycleStateFailed,
}

// GetHeatWaveNodeLifecycleStateEnumValues Enumerates the set of values for HeatWaveNodeLifecycleStateEnum
func GetHeatWaveNodeLifecycleStateEnumValues() []HeatWaveNodeLifecycleStateEnum {
	values := make([]HeatWaveNodeLifecycleStateEnum, 0)
	for _, v := range mappingHeatWaveNodeLifecycleState {
		values = append(values, v)
	}
	return values
}
