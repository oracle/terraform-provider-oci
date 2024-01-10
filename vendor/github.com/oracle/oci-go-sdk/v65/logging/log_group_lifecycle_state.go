// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"strings"
)

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

var mappingLogGroupLifecycleStateEnum = map[string]LogGroupLifecycleStateEnum{
	"CREATING": LogGroupLifecycleStateCreating,
	"ACTIVE":   LogGroupLifecycleStateActive,
	"UPDATING": LogGroupLifecycleStateUpdating,
	"INACTIVE": LogGroupLifecycleStateInactive,
	"DELETING": LogGroupLifecycleStateDeleting,
	"FAILED":   LogGroupLifecycleStateFailed,
}

var mappingLogGroupLifecycleStateEnumLowerCase = map[string]LogGroupLifecycleStateEnum{
	"creating": LogGroupLifecycleStateCreating,
	"active":   LogGroupLifecycleStateActive,
	"updating": LogGroupLifecycleStateUpdating,
	"inactive": LogGroupLifecycleStateInactive,
	"deleting": LogGroupLifecycleStateDeleting,
	"failed":   LogGroupLifecycleStateFailed,
}

// GetLogGroupLifecycleStateEnumValues Enumerates the set of values for LogGroupLifecycleStateEnum
func GetLogGroupLifecycleStateEnumValues() []LogGroupLifecycleStateEnum {
	values := make([]LogGroupLifecycleStateEnum, 0)
	for _, v := range mappingLogGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogGroupLifecycleStateEnumStringValues Enumerates the set of values in String for LogGroupLifecycleStateEnum
func GetLogGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingLogGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogGroupLifecycleStateEnum(val string) (LogGroupLifecycleStateEnum, bool) {
	enum, ok := mappingLogGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
