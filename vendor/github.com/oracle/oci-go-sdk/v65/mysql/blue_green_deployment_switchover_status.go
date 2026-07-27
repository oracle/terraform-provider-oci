// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// BlueGreenDeploymentSwitchoverStatusEnum Enum with underlying type: string
type BlueGreenDeploymentSwitchoverStatusEnum string

// Set of constants representing the allowable values for BlueGreenDeploymentSwitchoverStatusEnum
const (
	BlueGreenDeploymentSwitchoverStatusNotStarted           BlueGreenDeploymentSwitchoverStatusEnum = "NOT_STARTED"
	BlueGreenDeploymentSwitchoverStatusValidating           BlueGreenDeploymentSwitchoverStatusEnum = "VALIDATING"
	BlueGreenDeploymentSwitchoverStatusReplicatingData      BlueGreenDeploymentSwitchoverStatusEnum = "REPLICATING_DATA"
	BlueGreenDeploymentSwitchoverStatusReadyForSwitchover   BlueGreenDeploymentSwitchoverStatusEnum = "READY_FOR_SWITCHOVER"
	BlueGreenDeploymentSwitchoverStatusSwitchoverInProgress BlueGreenDeploymentSwitchoverStatusEnum = "SWITCHOVER_IN_PROGRESS"
	BlueGreenDeploymentSwitchoverStatusSwitchoverCompleted  BlueGreenDeploymentSwitchoverStatusEnum = "SWITCHOVER_COMPLETED"
	BlueGreenDeploymentSwitchoverStatusSwitchoverFailed     BlueGreenDeploymentSwitchoverStatusEnum = "SWITCHOVER_FAILED"
)

var mappingBlueGreenDeploymentSwitchoverStatusEnum = map[string]BlueGreenDeploymentSwitchoverStatusEnum{
	"NOT_STARTED":            BlueGreenDeploymentSwitchoverStatusNotStarted,
	"VALIDATING":             BlueGreenDeploymentSwitchoverStatusValidating,
	"REPLICATING_DATA":       BlueGreenDeploymentSwitchoverStatusReplicatingData,
	"READY_FOR_SWITCHOVER":   BlueGreenDeploymentSwitchoverStatusReadyForSwitchover,
	"SWITCHOVER_IN_PROGRESS": BlueGreenDeploymentSwitchoverStatusSwitchoverInProgress,
	"SWITCHOVER_COMPLETED":   BlueGreenDeploymentSwitchoverStatusSwitchoverCompleted,
	"SWITCHOVER_FAILED":      BlueGreenDeploymentSwitchoverStatusSwitchoverFailed,
}

var mappingBlueGreenDeploymentSwitchoverStatusEnumLowerCase = map[string]BlueGreenDeploymentSwitchoverStatusEnum{
	"not_started":            BlueGreenDeploymentSwitchoverStatusNotStarted,
	"validating":             BlueGreenDeploymentSwitchoverStatusValidating,
	"replicating_data":       BlueGreenDeploymentSwitchoverStatusReplicatingData,
	"ready_for_switchover":   BlueGreenDeploymentSwitchoverStatusReadyForSwitchover,
	"switchover_in_progress": BlueGreenDeploymentSwitchoverStatusSwitchoverInProgress,
	"switchover_completed":   BlueGreenDeploymentSwitchoverStatusSwitchoverCompleted,
	"switchover_failed":      BlueGreenDeploymentSwitchoverStatusSwitchoverFailed,
}

// GetBlueGreenDeploymentSwitchoverStatusEnumValues Enumerates the set of values for BlueGreenDeploymentSwitchoverStatusEnum
func GetBlueGreenDeploymentSwitchoverStatusEnumValues() []BlueGreenDeploymentSwitchoverStatusEnum {
	values := make([]BlueGreenDeploymentSwitchoverStatusEnum, 0)
	for _, v := range mappingBlueGreenDeploymentSwitchoverStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetBlueGreenDeploymentSwitchoverStatusEnumStringValues Enumerates the set of values in String for BlueGreenDeploymentSwitchoverStatusEnum
func GetBlueGreenDeploymentSwitchoverStatusEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"VALIDATING",
		"REPLICATING_DATA",
		"READY_FOR_SWITCHOVER",
		"SWITCHOVER_IN_PROGRESS",
		"SWITCHOVER_COMPLETED",
		"SWITCHOVER_FAILED",
	}
}

// GetMappingBlueGreenDeploymentSwitchoverStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBlueGreenDeploymentSwitchoverStatusEnum(val string) (BlueGreenDeploymentSwitchoverStatusEnum, bool) {
	enum, ok := mappingBlueGreenDeploymentSwitchoverStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
