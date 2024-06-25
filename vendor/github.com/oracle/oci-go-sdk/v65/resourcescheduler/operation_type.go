// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeStartResource             OperationTypeEnum = "START_RESOURCE"
	OperationTypeStopResource              OperationTypeEnum = "STOP_RESOURCE"
	OperationTypeChangeScheduleCompartment OperationTypeEnum = "CHANGE_SCHEDULE_COMPARTMENT"
	OperationTypeCreateSchedule            OperationTypeEnum = "CREATE_SCHEDULE"
	OperationTypeUpdateSchedule            OperationTypeEnum = "UPDATE_SCHEDULE"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"START_RESOURCE":              OperationTypeStartResource,
	"STOP_RESOURCE":               OperationTypeStopResource,
	"CHANGE_SCHEDULE_COMPARTMENT": OperationTypeChangeScheduleCompartment,
	"CREATE_SCHEDULE":             OperationTypeCreateSchedule,
	"UPDATE_SCHEDULE":             OperationTypeUpdateSchedule,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"start_resource":              OperationTypeStartResource,
	"stop_resource":               OperationTypeStopResource,
	"change_schedule_compartment": OperationTypeChangeScheduleCompartment,
	"create_schedule":             OperationTypeCreateSchedule,
	"update_schedule":             OperationTypeUpdateSchedule,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"START_RESOURCE",
		"STOP_RESOURCE",
		"CHANGE_SCHEDULE_COMPARTMENT",
		"CREATE_SCHEDULE",
		"UPDATE_SCHEDULE",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
