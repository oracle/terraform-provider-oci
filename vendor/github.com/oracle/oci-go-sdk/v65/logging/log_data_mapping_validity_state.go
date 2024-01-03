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

// LogDataMappingValidityStateEnum Enum with underlying type: string
type LogDataMappingValidityStateEnum string

// Set of constants representing the allowable values for LogDataMappingValidityStateEnum
const (
	LogDataMappingValidityStateValid     LogDataMappingValidityStateEnum = "VALID"
	LogDataMappingValidityStateInvalid   LogDataMappingValidityStateEnum = "INVALID"
	LogDataMappingValidityStateUndefined LogDataMappingValidityStateEnum = "UNDEFINED"
)

var mappingLogDataMappingValidityStateEnum = map[string]LogDataMappingValidityStateEnum{
	"VALID":     LogDataMappingValidityStateValid,
	"INVALID":   LogDataMappingValidityStateInvalid,
	"UNDEFINED": LogDataMappingValidityStateUndefined,
}

var mappingLogDataMappingValidityStateEnumLowerCase = map[string]LogDataMappingValidityStateEnum{
	"valid":     LogDataMappingValidityStateValid,
	"invalid":   LogDataMappingValidityStateInvalid,
	"undefined": LogDataMappingValidityStateUndefined,
}

// GetLogDataMappingValidityStateEnumValues Enumerates the set of values for LogDataMappingValidityStateEnum
func GetLogDataMappingValidityStateEnumValues() []LogDataMappingValidityStateEnum {
	values := make([]LogDataMappingValidityStateEnum, 0)
	for _, v := range mappingLogDataMappingValidityStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogDataMappingValidityStateEnumStringValues Enumerates the set of values in String for LogDataMappingValidityStateEnum
func GetLogDataMappingValidityStateEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
		"UNDEFINED",
	}
}

// GetMappingLogDataMappingValidityStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogDataMappingValidityStateEnum(val string) (LogDataMappingValidityStateEnum, bool) {
	enum, ok := mappingLogDataMappingValidityStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
