// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"strings"
)

// LogPipelineDestinationTypesEnum Enum with underlying type: string
type LogPipelineDestinationTypesEnum string

// Set of constants representing the allowable values for LogPipelineDestinationTypesEnum
const (
	LogPipelineDestinationTypesLogging       LogPipelineDestinationTypesEnum = "LOGGING"
	LogPipelineDestinationTypesObjectStorage LogPipelineDestinationTypesEnum = "OBJECT_STORAGE"
	LogPipelineDestinationTypesLoggingSearch LogPipelineDestinationTypesEnum = "LOGGING_SEARCH"
)

var mappingLogPipelineDestinationTypesEnum = map[string]LogPipelineDestinationTypesEnum{
	"LOGGING":        LogPipelineDestinationTypesLogging,
	"OBJECT_STORAGE": LogPipelineDestinationTypesObjectStorage,
	"LOGGING_SEARCH": LogPipelineDestinationTypesLoggingSearch,
}

var mappingLogPipelineDestinationTypesEnumLowerCase = map[string]LogPipelineDestinationTypesEnum{
	"logging":        LogPipelineDestinationTypesLogging,
	"object_storage": LogPipelineDestinationTypesObjectStorage,
	"logging_search": LogPipelineDestinationTypesLoggingSearch,
}

// GetLogPipelineDestinationTypesEnumValues Enumerates the set of values for LogPipelineDestinationTypesEnum
func GetLogPipelineDestinationTypesEnumValues() []LogPipelineDestinationTypesEnum {
	values := make([]LogPipelineDestinationTypesEnum, 0)
	for _, v := range mappingLogPipelineDestinationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogPipelineDestinationTypesEnumStringValues Enumerates the set of values in String for LogPipelineDestinationTypesEnum
func GetLogPipelineDestinationTypesEnumStringValues() []string {
	return []string{
		"LOGGING",
		"OBJECT_STORAGE",
		"LOGGING_SEARCH",
	}
}

// GetMappingLogPipelineDestinationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogPipelineDestinationTypesEnum(val string) (LogPipelineDestinationTypesEnum, bool) {
	enum, ok := mappingLogPipelineDestinationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
