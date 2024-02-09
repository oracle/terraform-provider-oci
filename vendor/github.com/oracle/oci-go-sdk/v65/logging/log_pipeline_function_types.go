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

// LogPipelineFunctionTypesEnum Enum with underlying type: string
type LogPipelineFunctionTypesEnum string

// Set of constants representing the allowable values for LogPipelineFunctionTypesEnum
const (
	LogPipelineFunctionTypesFilter    LogPipelineFunctionTypesEnum = "FILTER"
	LogPipelineFunctionTypesTransform LogPipelineFunctionTypesEnum = "TRANSFORM"
	LogPipelineFunctionTypesDrop      LogPipelineFunctionTypesEnum = "DROP"
	LogPipelineFunctionTypesRename    LogPipelineFunctionTypesEnum = "RENAME"
)

var mappingLogPipelineFunctionTypesEnum = map[string]LogPipelineFunctionTypesEnum{
	"FILTER":    LogPipelineFunctionTypesFilter,
	"TRANSFORM": LogPipelineFunctionTypesTransform,
	"DROP":      LogPipelineFunctionTypesDrop,
	"RENAME":    LogPipelineFunctionTypesRename,
}

var mappingLogPipelineFunctionTypesEnumLowerCase = map[string]LogPipelineFunctionTypesEnum{
	"filter":    LogPipelineFunctionTypesFilter,
	"transform": LogPipelineFunctionTypesTransform,
	"drop":      LogPipelineFunctionTypesDrop,
	"rename":    LogPipelineFunctionTypesRename,
}

// GetLogPipelineFunctionTypesEnumValues Enumerates the set of values for LogPipelineFunctionTypesEnum
func GetLogPipelineFunctionTypesEnumValues() []LogPipelineFunctionTypesEnum {
	values := make([]LogPipelineFunctionTypesEnum, 0)
	for _, v := range mappingLogPipelineFunctionTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogPipelineFunctionTypesEnumStringValues Enumerates the set of values in String for LogPipelineFunctionTypesEnum
func GetLogPipelineFunctionTypesEnumStringValues() []string {
	return []string{
		"FILTER",
		"TRANSFORM",
		"DROP",
		"RENAME",
	}
}

// GetMappingLogPipelineFunctionTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogPipelineFunctionTypesEnum(val string) (LogPipelineFunctionTypesEnum, bool) {
	enum, ok := mappingLogPipelineFunctionTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
