// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

// LogSavedSearchLifecycleStateEnum Enum with underlying type: string
type LogSavedSearchLifecycleStateEnum string

// Set of constants representing the allowable values for LogSavedSearchLifecycleStateEnum
const (
	LogSavedSearchLifecycleStateCreating LogSavedSearchLifecycleStateEnum = "CREATING"
	LogSavedSearchLifecycleStateActive   LogSavedSearchLifecycleStateEnum = "ACTIVE"
	LogSavedSearchLifecycleStateUpdating LogSavedSearchLifecycleStateEnum = "UPDATING"
	LogSavedSearchLifecycleStateInactive LogSavedSearchLifecycleStateEnum = "INACTIVE"
	LogSavedSearchLifecycleStateDeleting LogSavedSearchLifecycleStateEnum = "DELETING"
	LogSavedSearchLifecycleStateFailed   LogSavedSearchLifecycleStateEnum = "FAILED"
)

var mappingLogSavedSearchLifecycleState = map[string]LogSavedSearchLifecycleStateEnum{
	"CREATING": LogSavedSearchLifecycleStateCreating,
	"ACTIVE":   LogSavedSearchLifecycleStateActive,
	"UPDATING": LogSavedSearchLifecycleStateUpdating,
	"INACTIVE": LogSavedSearchLifecycleStateInactive,
	"DELETING": LogSavedSearchLifecycleStateDeleting,
	"FAILED":   LogSavedSearchLifecycleStateFailed,
}

// GetLogSavedSearchLifecycleStateEnumValues Enumerates the set of values for LogSavedSearchLifecycleStateEnum
func GetLogSavedSearchLifecycleStateEnumValues() []LogSavedSearchLifecycleStateEnum {
	values := make([]LogSavedSearchLifecycleStateEnum, 0)
	for _, v := range mappingLogSavedSearchLifecycleState {
		values = append(values, v)
	}
	return values
}
