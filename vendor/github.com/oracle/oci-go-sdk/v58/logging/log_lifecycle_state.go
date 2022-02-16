// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, and delete log groups, log objects, and agent configurations.
//

package logging

import (
	"strings"
)

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

var mappingLogLifecycleStateEnum = map[string]LogLifecycleStateEnum{
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
	for _, v := range mappingLogLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogLifecycleStateEnumStringValues Enumerates the set of values in String for LogLifecycleStateEnum
func GetLogLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingLogLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogLifecycleStateEnum(val string) (LogLifecycleStateEnum, bool) {
	mappingLogLifecycleStateEnumIgnoreCase := make(map[string]LogLifecycleStateEnum)
	for k, v := range mappingLogLifecycleStateEnum {
		mappingLogLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingLogLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
