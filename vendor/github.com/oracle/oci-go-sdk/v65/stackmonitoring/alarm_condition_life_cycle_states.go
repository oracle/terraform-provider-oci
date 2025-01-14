// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"strings"
)

// AlarmConditionLifeCycleStatesEnum Enum with underlying type: string
type AlarmConditionLifeCycleStatesEnum string

// Set of constants representing the allowable values for AlarmConditionLifeCycleStatesEnum
const (
	AlarmConditionLifeCycleStatesCreating AlarmConditionLifeCycleStatesEnum = "CREATING"
	AlarmConditionLifeCycleStatesActive   AlarmConditionLifeCycleStatesEnum = "ACTIVE"
	AlarmConditionLifeCycleStatesInactive AlarmConditionLifeCycleStatesEnum = "INACTIVE"
	AlarmConditionLifeCycleStatesUpdating AlarmConditionLifeCycleStatesEnum = "UPDATING"
	AlarmConditionLifeCycleStatesDeleted  AlarmConditionLifeCycleStatesEnum = "DELETED"
)

var mappingAlarmConditionLifeCycleStatesEnum = map[string]AlarmConditionLifeCycleStatesEnum{
	"CREATING": AlarmConditionLifeCycleStatesCreating,
	"ACTIVE":   AlarmConditionLifeCycleStatesActive,
	"INACTIVE": AlarmConditionLifeCycleStatesInactive,
	"UPDATING": AlarmConditionLifeCycleStatesUpdating,
	"DELETED":  AlarmConditionLifeCycleStatesDeleted,
}

var mappingAlarmConditionLifeCycleStatesEnumLowerCase = map[string]AlarmConditionLifeCycleStatesEnum{
	"creating": AlarmConditionLifeCycleStatesCreating,
	"active":   AlarmConditionLifeCycleStatesActive,
	"inactive": AlarmConditionLifeCycleStatesInactive,
	"updating": AlarmConditionLifeCycleStatesUpdating,
	"deleted":  AlarmConditionLifeCycleStatesDeleted,
}

// GetAlarmConditionLifeCycleStatesEnumValues Enumerates the set of values for AlarmConditionLifeCycleStatesEnum
func GetAlarmConditionLifeCycleStatesEnumValues() []AlarmConditionLifeCycleStatesEnum {
	values := make([]AlarmConditionLifeCycleStatesEnum, 0)
	for _, v := range mappingAlarmConditionLifeCycleStatesEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmConditionLifeCycleStatesEnumStringValues Enumerates the set of values in String for AlarmConditionLifeCycleStatesEnum
func GetAlarmConditionLifeCycleStatesEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingAlarmConditionLifeCycleStatesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmConditionLifeCycleStatesEnum(val string) (AlarmConditionLifeCycleStatesEnum, bool) {
	enum, ok := mappingAlarmConditionLifeCycleStatesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
