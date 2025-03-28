// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// TaskScopeEnum Enum with underlying type: string
type TaskScopeEnum string

// Set of constants representing the allowable values for TaskScopeEnum
const (
	TaskScopeLocal  TaskScopeEnum = "LOCAL"
	TaskScopeShared TaskScopeEnum = "SHARED"
)

var mappingTaskScopeEnum = map[string]TaskScopeEnum{
	"LOCAL":  TaskScopeLocal,
	"SHARED": TaskScopeShared,
}

var mappingTaskScopeEnumLowerCase = map[string]TaskScopeEnum{
	"local":  TaskScopeLocal,
	"shared": TaskScopeShared,
}

// GetTaskScopeEnumValues Enumerates the set of values for TaskScopeEnum
func GetTaskScopeEnumValues() []TaskScopeEnum {
	values := make([]TaskScopeEnum, 0)
	for _, v := range mappingTaskScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskScopeEnumStringValues Enumerates the set of values in String for TaskScopeEnum
func GetTaskScopeEnumStringValues() []string {
	return []string{
		"LOCAL",
		"SHARED",
	}
}

// GetMappingTaskScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskScopeEnum(val string) (TaskScopeEnum, bool) {
	enum, ok := mappingTaskScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
