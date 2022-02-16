// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"strings"
)

// TaskTypeEnum Enum with underlying type: string
type TaskTypeEnum string

// Set of constants representing the allowable values for TaskTypeEnum
const (
	TaskTypeSavedSearch             TaskTypeEnum = "SAVED_SEARCH"
	TaskTypeAcceleration            TaskTypeEnum = "ACCELERATION"
	TaskTypePurge                   TaskTypeEnum = "PURGE"
	TaskTypeAccelerationMaintenance TaskTypeEnum = "ACCELERATION_MAINTENANCE"
)

var mappingTaskTypeEnum = map[string]TaskTypeEnum{
	"SAVED_SEARCH":             TaskTypeSavedSearch,
	"ACCELERATION":             TaskTypeAcceleration,
	"PURGE":                    TaskTypePurge,
	"ACCELERATION_MAINTENANCE": TaskTypeAccelerationMaintenance,
}

// GetTaskTypeEnumValues Enumerates the set of values for TaskTypeEnum
func GetTaskTypeEnumValues() []TaskTypeEnum {
	values := make([]TaskTypeEnum, 0)
	for _, v := range mappingTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskTypeEnumStringValues Enumerates the set of values in String for TaskTypeEnum
func GetTaskTypeEnumStringValues() []string {
	return []string{
		"SAVED_SEARCH",
		"ACCELERATION",
		"PURGE",
		"ACCELERATION_MAINTENANCE",
	}
}

// GetMappingTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskTypeEnum(val string) (TaskTypeEnum, bool) {
	mappingTaskTypeEnumIgnoreCase := make(map[string]TaskTypeEnum)
	for k, v := range mappingTaskTypeEnum {
		mappingTaskTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
