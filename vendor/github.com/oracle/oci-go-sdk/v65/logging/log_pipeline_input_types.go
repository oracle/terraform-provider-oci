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

// LogPipelineInputTypesEnum Enum with underlying type: string
type LogPipelineInputTypesEnum string

// Set of constants representing the allowable values for LogPipelineInputTypesEnum
const (
	LogPipelineInputTypesOciLogObject LogPipelineInputTypesEnum = "OCI_LOG_OBJECT"
)

var mappingLogPipelineInputTypesEnum = map[string]LogPipelineInputTypesEnum{
	"OCI_LOG_OBJECT": LogPipelineInputTypesOciLogObject,
}

var mappingLogPipelineInputTypesEnumLowerCase = map[string]LogPipelineInputTypesEnum{
	"oci_log_object": LogPipelineInputTypesOciLogObject,
}

// GetLogPipelineInputTypesEnumValues Enumerates the set of values for LogPipelineInputTypesEnum
func GetLogPipelineInputTypesEnumValues() []LogPipelineInputTypesEnum {
	values := make([]LogPipelineInputTypesEnum, 0)
	for _, v := range mappingLogPipelineInputTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetLogPipelineInputTypesEnumStringValues Enumerates the set of values in String for LogPipelineInputTypesEnum
func GetLogPipelineInputTypesEnumStringValues() []string {
	return []string{
		"OCI_LOG_OBJECT",
	}
}

// GetMappingLogPipelineInputTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogPipelineInputTypesEnum(val string) (LogPipelineInputTypesEnum, bool) {
	enum, ok := mappingLogPipelineInputTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
