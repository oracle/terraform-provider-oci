// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// TaskTypeEnum Enum with underlying type: string
type TaskTypeEnum string

// Set of constants representing the allowable values for TaskTypeEnum
const (
	TaskTypeSavedSearch             TaskTypeEnum = "SAVED_SEARCH"
	TaskTypeAcceleration            TaskTypeEnum = "ACCELERATION"
	TaskTypePurge                   TaskTypeEnum = "PURGE"
	TaskTypeAccelerationMaintenance TaskTypeEnum = "ACCELERATION_MAINTENANCE"
)

var mappingTaskType = map[string]TaskTypeEnum{
	"SAVED_SEARCH":             TaskTypeSavedSearch,
	"ACCELERATION":             TaskTypeAcceleration,
	"PURGE":                    TaskTypePurge,
	"ACCELERATION_MAINTENANCE": TaskTypeAccelerationMaintenance,
}

// GetTaskTypeEnumValues Enumerates the set of values for TaskTypeEnum
func GetTaskTypeEnumValues() []TaskTypeEnum {
	values := make([]TaskTypeEnum, 0)
	for _, v := range mappingTaskType {
		values = append(values, v)
	}
	return values
}
