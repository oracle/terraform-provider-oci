// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TaskExecutionTypeEnum Enum with underlying type: string
type TaskExecutionTypeEnum string

// Set of constants representing the allowable values for TaskExecutionTypeEnum
const (
	TaskExecutionTypeScript TaskExecutionTypeEnum = "SCRIPT"
	TaskExecutionTypeApi    TaskExecutionTypeEnum = "API"
)

var mappingTaskExecutionTypeEnum = map[string]TaskExecutionTypeEnum{
	"SCRIPT": TaskExecutionTypeScript,
	"API":    TaskExecutionTypeApi,
}

var mappingTaskExecutionTypeEnumLowerCase = map[string]TaskExecutionTypeEnum{
	"script": TaskExecutionTypeScript,
	"api":    TaskExecutionTypeApi,
}

// GetTaskExecutionTypeEnumValues Enumerates the set of values for TaskExecutionTypeEnum
func GetTaskExecutionTypeEnumValues() []TaskExecutionTypeEnum {
	values := make([]TaskExecutionTypeEnum, 0)
	for _, v := range mappingTaskExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskExecutionTypeEnumStringValues Enumerates the set of values in String for TaskExecutionTypeEnum
func GetTaskExecutionTypeEnumStringValues() []string {
	return []string{
		"SCRIPT",
		"API",
	}
}

// GetMappingTaskExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskExecutionTypeEnum(val string) (TaskExecutionTypeEnum, bool) {
	enum, ok := mappingTaskExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
