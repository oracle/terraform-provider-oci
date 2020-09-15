// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

// LogLifecycleStateEnum Enum with underlying type: string
type LogLifecycleStateEnum string

// Set of constants representing the allowable values for LogLifecycleStateEnum
const (
	LogLifecycleStateCreating LogLifecycleStateEnum = "CREATING"
	LogLifecycleStateActive   LogLifecycleStateEnum = "ACTIVE"
	LogLifecycleStateUpdating LogLifecycleStateEnum = "UPDATING"
	LogLifecycleStateInactive LogLifecycleStateEnum = "INACTIVE"
	LogLifecycleStateDeleting LogLifecycleStateEnum = "DELETING"
	LogLifecycleStateFailed   LogLifecycleStateEnum = "FAILED"
)

var mappingLogLifecycleState = map[string]LogLifecycleStateEnum{
	"CREATING": LogLifecycleStateCreating,
	"ACTIVE":   LogLifecycleStateActive,
	"UPDATING": LogLifecycleStateUpdating,
	"INACTIVE": LogLifecycleStateInactive,
	"DELETING": LogLifecycleStateDeleting,
	"FAILED":   LogLifecycleStateFailed,
}

// GetLogLifecycleStateEnumValues Enumerates the set of values for LogLifecycleStateEnum
func GetLogLifecycleStateEnumValues() []LogLifecycleStateEnum {
	values := make([]LogLifecycleStateEnum, 0)
	for _, v := range mappingLogLifecycleState {
		values = append(values, v)
	}
	return values
}
