// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, agent configurations, log data models,
// continuous queries, and managed continuous queries.
// For more information, see https://docs.oracle.com/en-us/iaas/Content/Logging/Concepts/loggingoverview.htm.
//

package logging

import (
	"strings"
)

// LogDataModelValidityStateEnum Enum with underlying type: string
type LogDataModelValidityStateEnum string

// Set of constants representing the allowable values for LogDataModelValidityStateEnum
const (
	LogDataModelValidityStateValid     LogDataModelValidityStateEnum = "VALID"
	LogDataModelValidityStateInvalid   LogDataModelValidityStateEnum = "INVALID"
	LogDataModelValidityStateUndefined LogDataModelValidityStateEnum = "UNDEFINED"
)

var mappingLogDataModelValidityStateEnum = map[string]LogDataModelValidityStateEnum{
	"VALID":     LogDataModelValidityStateValid,
	"INVALID":   LogDataModelValidityStateInvalid,
	"UNDEFINED": LogDataModelValidityStateUndefined,
}

var mappingLogDataModelValidityStateEnumLowerCase = map[string]LogDataModelValidityStateEnum{
	"valid":     LogDataModelValidityStateValid,
	"invalid":   LogDataModelValidityStateInvalid,
	"undefined": LogDataModelValidityStateUndefined,
}

// GetLogDataModelValidityStateEnumValues Enumerates the set of values for LogDataModelValidityStateEnum
func GetLogDataModelValidityStateEnumValues() []LogDataModelValidityStateEnum {
	values := make([]LogDataModelValidityStateEnum, 0)
	for _, v := range mappingLogDataModelValidityStateEnum {
		values = append(values, v)
	}
	return values
}

// GetLogDataModelValidityStateEnumStringValues Enumerates the set of values in String for LogDataModelValidityStateEnum
func GetLogDataModelValidityStateEnumStringValues() []string {
	return []string{
		"VALID",
		"INVALID",
		"UNDEFINED",
	}
}

// GetMappingLogDataModelValidityStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogDataModelValidityStateEnum(val string) (LogDataModelValidityStateEnum, bool) {
	enum, ok := mappingLogDataModelValidityStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
