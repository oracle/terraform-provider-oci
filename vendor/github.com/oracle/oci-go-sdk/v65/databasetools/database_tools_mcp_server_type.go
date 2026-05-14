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

// DatabaseToolsMcpServerTypeEnum Enum with underlying type: string
type DatabaseToolsMcpServerTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpServerTypeEnum
const (
	DatabaseToolsMcpServerTypeDefault DatabaseToolsMcpServerTypeEnum = "DEFAULT"
)

var mappingDatabaseToolsMcpServerTypeEnum = map[string]DatabaseToolsMcpServerTypeEnum{
	"DEFAULT": DatabaseToolsMcpServerTypeDefault,
}

var mappingDatabaseToolsMcpServerTypeEnumLowerCase = map[string]DatabaseToolsMcpServerTypeEnum{
	"default": DatabaseToolsMcpServerTypeDefault,
}

// GetDatabaseToolsMcpServerTypeEnumValues Enumerates the set of values for DatabaseToolsMcpServerTypeEnum
func GetDatabaseToolsMcpServerTypeEnumValues() []DatabaseToolsMcpServerTypeEnum {
	values := make([]DatabaseToolsMcpServerTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpServerTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpServerTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpServerTypeEnum
func GetDatabaseToolsMcpServerTypeEnumStringValues() []string {
	return []string{
		"DEFAULT",
	}
}

// GetMappingDatabaseToolsMcpServerTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpServerTypeEnum(val string) (DatabaseToolsMcpServerTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpServerTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
