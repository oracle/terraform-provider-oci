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

// LogPipelineLifecycleStateEnum Enum with underlying type: string
type LogPipelineLifecycleStateEnum string

// Set of constants representing the allowable values for LogPipelineLifecycleStateEnum
const (
	LogPipelineLifecycleStateCreating LogPipelineLifecycleStateEnum = "CREATING"
	LogPipelineLifecycleStateActive   LogPipelineLifecycleStateEnum = "ACTIVE"
	LogPipelineLifecycleStateUpdating LogPipelineLifecycleStateEnum = "UPDATING"
	LogPipelineLifecycleStateInactive LogPipelineLifecycleStateEnum = "INACTIVE"
	LogPipelineLifecycleStateDeleting LogPipelineLifecycleStateEnum = "DELETING"
	LogPipelineLifecycleStateFailed   LogPipelineLifecycleStateEnum = "FAILED"
)

var mappingLogPipelineLifecycleStateEnum = map[string]LogPipelineLifecycleStateEnum{
	"CREATING": LogPipelineLifecycleStateCreating,
	"ACTIVE":   LogPipelineLifecycleStateActive,
	"UPDATING": LogPipelineLifecycleStateUpdating,
	"INACTIVE": LogPipelineLifecycleStateInactive,
	"DELETING": LogPipelineLifecycleStateDeleting,
	"FAILED":   LogPipelineLifecycleStateFailed,
}

var mappingLogPipelineLifecycleStateEnumLowerCase = map[string]LogPipelineLifecycleStateEnum{
	"creating": LogPipelineLifecycleStateCreating,
	"active":   LogPipelineLifecycleStateActive,
	"updating": LogPipelineLifecycleStateUpdating,
	"inactive": LogPipelineLifecycleStateInactive,
	"deleting": LogPipelineLifecycleStateDeleting,
	"failed":   LogPipelineLifecycleStateFailed,
}

// GetLogPipelineLifecycleStateEnumValues Enumerates the set of values for LogPipelineLifecycleStateEnum
func GetLogPipelineLifecycleStateEnumValues() []LogPipelineLifecycleStateEnum {
	values := make([]LogPipelineLifecycleStateEnum, 0)
	for _, v := range mappingLogPipelineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogPipelineLifecycleStateEnumStringValues Enumerates the set of values in String for LogPipelineLifecycleStateEnum
func GetLogPipelineLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingLogPipelineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogPipelineLifecycleStateEnum(val string) (LogPipelineLifecycleStateEnum, bool) {
	enum, ok := mappingLogPipelineLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
