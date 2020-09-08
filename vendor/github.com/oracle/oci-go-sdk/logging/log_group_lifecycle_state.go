// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

// LogGroupLifecycleStateEnum Enum with underlying type: string
type LogGroupLifecycleStateEnum string

// Set of constants representing the allowable values for LogGroupLifecycleStateEnum
const (
	LogGroupLifecycleStateCreating LogGroupLifecycleStateEnum = "CREATING"
	LogGroupLifecycleStateActive   LogGroupLifecycleStateEnum = "ACTIVE"
	LogGroupLifecycleStateUpdating LogGroupLifecycleStateEnum = "UPDATING"
	LogGroupLifecycleStateInactive LogGroupLifecycleStateEnum = "INACTIVE"
	LogGroupLifecycleStateDeleting LogGroupLifecycleStateEnum = "DELETING"
	LogGroupLifecycleStateFailed   LogGroupLifecycleStateEnum = "FAILED"
)

var mappingLogGroupLifecycleState = map[string]LogGroupLifecycleStateEnum{
	"CREATING": LogGroupLifecycleStateCreating,
	"ACTIVE":   LogGroupLifecycleStateActive,
	"UPDATING": LogGroupLifecycleStateUpdating,
	"INACTIVE": LogGroupLifecycleStateInactive,
	"DELETING": LogGroupLifecycleStateDeleting,
	"FAILED":   LogGroupLifecycleStateFailed,
}

// GetLogGroupLifecycleStateEnumValues Enumerates the set of values for LogGroupLifecycleStateEnum
func GetLogGroupLifecycleStateEnumValues() []LogGroupLifecycleStateEnum {
	values := make([]LogGroupLifecycleStateEnum, 0)
	for _, v := range mappingLogGroupLifecycleState {
		values = append(values, v)
	}
	return values
}
