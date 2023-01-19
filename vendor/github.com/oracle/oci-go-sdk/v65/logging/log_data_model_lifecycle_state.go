// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"strings"
)

// LogDataModelLifecycleStateEnum Enum with underlying type: string
type LogDataModelLifecycleStateEnum string

// Set of constants representing the allowable values for LogDataModelLifecycleStateEnum
const (
	LogDataModelLifecycleStateCreating LogDataModelLifecycleStateEnum = "CREATING"
	LogDataModelLifecycleStateUpdating LogDataModelLifecycleStateEnum = "UPDATING"
	LogDataModelLifecycleStateDeleting LogDataModelLifecycleStateEnum = "DELETING"
	LogDataModelLifecycleStateActive   LogDataModelLifecycleStateEnum = "ACTIVE"
	LogDataModelLifecycleStateFailed   LogDataModelLifecycleStateEnum = "FAILED"
)

var mappingLogDataModelLifecycleStateEnum = map[string]LogDataModelLifecycleStateEnum{
	"CREATING": LogDataModelLifecycleStateCreating,
	"UPDATING": LogDataModelLifecycleStateUpdating,
	"DELETING": LogDataModelLifecycleStateDeleting,
	"ACTIVE":   LogDataModelLifecycleStateActive,
	"FAILED":   LogDataModelLifecycleStateFailed,
}

var mappingLogDataModelLifecycleStateEnumLowerCase = map[string]LogDataModelLifecycleStateEnum{
	"creating": LogDataModelLifecycleStateCreating,
	"updating": LogDataModelLifecycleStateUpdating,
	"deleting": LogDataModelLifecycleStateDeleting,
	"active":   LogDataModelLifecycleStateActive,
	"failed":   LogDataModelLifecycleStateFailed,
}

// GetLogDataModelLifecycleStateEnumValues Enumerates the set of values for LogDataModelLifecycleStateEnum
func GetLogDataModelLifecycleStateEnumValues() []LogDataModelLifecycleStateEnum {
	values := make([]LogDataModelLifecycleStateEnum, 0)
	for _, v := range mappingLogDataModelLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogDataModelLifecycleStateEnumStringValues Enumerates the set of values in String for LogDataModelLifecycleStateEnum
func GetLogDataModelLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"DELETING",
		"ACTIVE",
		"FAILED",
	}
}

// GetMappingLogDataModelLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogDataModelLifecycleStateEnum(val string) (LogDataModelLifecycleStateEnum, bool) {
	enum, ok := mappingLogDataModelLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
