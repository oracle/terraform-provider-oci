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

// AlarmConditionSeverityEnum Enum with underlying type: string
type AlarmConditionSeverityEnum string

// Set of constants representing the allowable values for AlarmConditionSeverityEnum
const (
	AlarmConditionSeverityCritical AlarmConditionSeverityEnum = "CRITICAL"
	AlarmConditionSeverityWarning  AlarmConditionSeverityEnum = "WARNING"
)

var mappingAlarmConditionSeverityEnum = map[string]AlarmConditionSeverityEnum{
	"CRITICAL": AlarmConditionSeverityCritical,
	"WARNING":  AlarmConditionSeverityWarning,
}

var mappingAlarmConditionSeverityEnumLowerCase = map[string]AlarmConditionSeverityEnum{
	"critical": AlarmConditionSeverityCritical,
	"warning":  AlarmConditionSeverityWarning,
}

// GetAlarmConditionSeverityEnumValues Enumerates the set of values for AlarmConditionSeverityEnum
func GetAlarmConditionSeverityEnumValues() []AlarmConditionSeverityEnum {
	values := make([]AlarmConditionSeverityEnum, 0)
	for _, v := range mappingAlarmConditionSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmConditionSeverityEnumStringValues Enumerates the set of values in String for AlarmConditionSeverityEnum
func GetAlarmConditionSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"WARNING",
	}
}

// GetMappingAlarmConditionSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmConditionSeverityEnum(val string) (AlarmConditionSeverityEnum, bool) {
	enum, ok := mappingAlarmConditionSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
