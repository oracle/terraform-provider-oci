// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"strings"
)

// DatabaseToolsExternalAuthenticationStatusEnum Enum with underlying type: string
type DatabaseToolsExternalAuthenticationStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsExternalAuthenticationStatusEnum
const (
	DatabaseToolsExternalAuthenticationStatusAvailable   DatabaseToolsExternalAuthenticationStatusEnum = "AVAILABLE"
	DatabaseToolsExternalAuthenticationStatusUnavailable DatabaseToolsExternalAuthenticationStatusEnum = "UNAVAILABLE"
	DatabaseToolsExternalAuthenticationStatusEnabled     DatabaseToolsExternalAuthenticationStatusEnum = "ENABLED"
	DatabaseToolsExternalAuthenticationStatusUnknown     DatabaseToolsExternalAuthenticationStatusEnum = "UNKNOWN"
)

var mappingDatabaseToolsExternalAuthenticationStatusEnum = map[string]DatabaseToolsExternalAuthenticationStatusEnum{
	"AVAILABLE":   DatabaseToolsExternalAuthenticationStatusAvailable,
	"UNAVAILABLE": DatabaseToolsExternalAuthenticationStatusUnavailable,
	"ENABLED":     DatabaseToolsExternalAuthenticationStatusEnabled,
	"UNKNOWN":     DatabaseToolsExternalAuthenticationStatusUnknown,
}

var mappingDatabaseToolsExternalAuthenticationStatusEnumLowerCase = map[string]DatabaseToolsExternalAuthenticationStatusEnum{
	"available":   DatabaseToolsExternalAuthenticationStatusAvailable,
	"unavailable": DatabaseToolsExternalAuthenticationStatusUnavailable,
	"enabled":     DatabaseToolsExternalAuthenticationStatusEnabled,
	"unknown":     DatabaseToolsExternalAuthenticationStatusUnknown,
}

// GetDatabaseToolsExternalAuthenticationStatusEnumValues Enumerates the set of values for DatabaseToolsExternalAuthenticationStatusEnum
func GetDatabaseToolsExternalAuthenticationStatusEnumValues() []DatabaseToolsExternalAuthenticationStatusEnum {
	values := make([]DatabaseToolsExternalAuthenticationStatusEnum, 0)
	for _, v := range mappingDatabaseToolsExternalAuthenticationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsExternalAuthenticationStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsExternalAuthenticationStatusEnum
func GetDatabaseToolsExternalAuthenticationStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
		"ENABLED",
		"UNKNOWN",
	}
}

// GetMappingDatabaseToolsExternalAuthenticationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsExternalAuthenticationStatusEnum(val string) (DatabaseToolsExternalAuthenticationStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsExternalAuthenticationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
