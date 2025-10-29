// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"strings"
)

// DatabaseToolsIdentityStatusEnum Enum with underlying type: string
type DatabaseToolsIdentityStatusEnum string

// Set of constants representing the allowable values for DatabaseToolsIdentityStatusEnum
const (
	DatabaseToolsIdentityStatusAvailable   DatabaseToolsIdentityStatusEnum = "AVAILABLE"
	DatabaseToolsIdentityStatusUnavailable DatabaseToolsIdentityStatusEnum = "UNAVAILABLE"
)

var mappingDatabaseToolsIdentityStatusEnum = map[string]DatabaseToolsIdentityStatusEnum{
	"AVAILABLE":   DatabaseToolsIdentityStatusAvailable,
	"UNAVAILABLE": DatabaseToolsIdentityStatusUnavailable,
}

var mappingDatabaseToolsIdentityStatusEnumLowerCase = map[string]DatabaseToolsIdentityStatusEnum{
	"available":   DatabaseToolsIdentityStatusAvailable,
	"unavailable": DatabaseToolsIdentityStatusUnavailable,
}

// GetDatabaseToolsIdentityStatusEnumValues Enumerates the set of values for DatabaseToolsIdentityStatusEnum
func GetDatabaseToolsIdentityStatusEnumValues() []DatabaseToolsIdentityStatusEnum {
	values := make([]DatabaseToolsIdentityStatusEnum, 0)
	for _, v := range mappingDatabaseToolsIdentityStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsIdentityStatusEnumStringValues Enumerates the set of values in String for DatabaseToolsIdentityStatusEnum
func GetDatabaseToolsIdentityStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

// GetMappingDatabaseToolsIdentityStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsIdentityStatusEnum(val string) (DatabaseToolsIdentityStatusEnum, bool) {
	enum, ok := mappingDatabaseToolsIdentityStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
