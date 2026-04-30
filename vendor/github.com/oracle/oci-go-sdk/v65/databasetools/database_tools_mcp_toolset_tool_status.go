// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools API
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsMcpToolsetToolStatusEnum Enum with underlying type: string
type DatabaseToolsMcpToolsetToolStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpToolsetToolStatusEnum
const (
	DatabaseToolsMcpToolsetToolStatusEnabled  DatabaseToolsMcpToolsetToolStatusEnum = "ENABLED"
	DatabaseToolsMcpToolsetToolStatusDisabled DatabaseToolsMcpToolsetToolStatusEnum = "DISABLED"
)

var mappingDatabaseToolsMcpToolsetToolStatusEnum = map[string]DatabaseToolsMcpToolsetToolStatusEnum{
	"ENABLED":  DatabaseToolsMcpToolsetToolStatusEnabled,
	"DISABLED": DatabaseToolsMcpToolsetToolStatusDisabled,
}

var mappingDatabaseToolsMcpToolsetToolStatusEnumLowerCase = map[string]DatabaseToolsMcpToolsetToolStatusEnum{
	"enabled":  DatabaseToolsMcpToolsetToolStatusEnabled,
	"disabled": DatabaseToolsMcpToolsetToolStatusDisabled,
}

// GetDatabaseToolsMcpToolsetToolStatusEnumValues Enumerates the set of values for DatabaseToolsMcpToolsetToolStatusEnum
func GetDatabaseToolsMcpToolsetToolStatusEnumValues() []DatabaseToolsMcpToolsetToolStatusEnum {
	values := make([]DatabaseToolsMcpToolsetToolStatusEnum, 0)
	for _, v := range mappingDatabaseToolsMcpToolsetToolStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpToolsetToolStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpToolsetToolStatusEnum
func GetDatabaseToolsMcpToolsetToolStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseToolsMcpToolsetToolStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpToolsetToolStatusEnum(val string) (DatabaseToolsMcpToolsetToolStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpToolsetToolStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
