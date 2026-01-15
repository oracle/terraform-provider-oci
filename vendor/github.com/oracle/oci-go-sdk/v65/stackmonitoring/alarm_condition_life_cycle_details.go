// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// AlarmConditionLifeCycleDetailsEnum Enum with underlying type: string
type AlarmConditionLifeCycleDetailsEnum string

// Set of constants representing the allowable values for AlarmConditionLifeCycleDetailsEnum
const (
	AlarmConditionLifeCycleDetailsNotApplied     AlarmConditionLifeCycleDetailsEnum = "NOT_APPLIED"
	AlarmConditionLifeCycleDetailsApplied        AlarmConditionLifeCycleDetailsEnum = "APPLIED"
	AlarmConditionLifeCycleDetailsPartialApplied AlarmConditionLifeCycleDetailsEnum = "PARTIAL_APPLIED"
	AlarmConditionLifeCycleDetailsError          AlarmConditionLifeCycleDetailsEnum = "ERROR"
)

var mappingAlarmConditionLifeCycleDetailsEnum = map[string]AlarmConditionLifeCycleDetailsEnum{
	"NOT_APPLIED":     AlarmConditionLifeCycleDetailsNotApplied,
	"APPLIED":         AlarmConditionLifeCycleDetailsApplied,
	"PARTIAL_APPLIED": AlarmConditionLifeCycleDetailsPartialApplied,
	"ERROR":           AlarmConditionLifeCycleDetailsError,
}

var mappingAlarmConditionLifeCycleDetailsEnumLowerCase = map[string]AlarmConditionLifeCycleDetailsEnum{
	"not_applied":     AlarmConditionLifeCycleDetailsNotApplied,
	"applied":         AlarmConditionLifeCycleDetailsApplied,
	"partial_applied": AlarmConditionLifeCycleDetailsPartialApplied,
	"error":           AlarmConditionLifeCycleDetailsError,
}

// GetAlarmConditionLifeCycleDetailsEnumValues Enumerates the set of values for AlarmConditionLifeCycleDetailsEnum
func GetAlarmConditionLifeCycleDetailsEnumValues() []AlarmConditionLifeCycleDetailsEnum {
	values := make([]AlarmConditionLifeCycleDetailsEnum, 0)
	for _, v := range mappingAlarmConditionLifeCycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmConditionLifeCycleDetailsEnumStringValues Enumerates the set of values in String for AlarmConditionLifeCycleDetailsEnum
func GetAlarmConditionLifeCycleDetailsEnumStringValues() []string {
	return []string{
		"NOT_APPLIED",
		"APPLIED",
		"PARTIAL_APPLIED",
		"ERROR",
	}
}

// GetMappingAlarmConditionLifeCycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmConditionLifeCycleDetailsEnum(val string) (AlarmConditionLifeCycleDetailsEnum, bool) {
	enum, ok := mappingAlarmConditionLifeCycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
