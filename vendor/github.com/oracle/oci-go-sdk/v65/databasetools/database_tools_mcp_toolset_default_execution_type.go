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

// DatabaseToolsMcpToolsetDefaultExecutionTypeEnum Enum with underlying type: string
type DatabaseToolsMcpToolsetDefaultExecutionTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsMcpToolsetDefaultExecutionTypeEnum
const (
	DatabaseToolsMcpToolsetDefaultExecutionTypeSynchronous  DatabaseToolsMcpToolsetDefaultExecutionTypeEnum = "SYNCHRONOUS"
	DatabaseToolsMcpToolsetDefaultExecutionTypeAsynchronous DatabaseToolsMcpToolsetDefaultExecutionTypeEnum = "ASYNCHRONOUS"
)

var mappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum = map[string]DatabaseToolsMcpToolsetDefaultExecutionTypeEnum{
	"SYNCHRONOUS":  DatabaseToolsMcpToolsetDefaultExecutionTypeSynchronous,
	"ASYNCHRONOUS": DatabaseToolsMcpToolsetDefaultExecutionTypeAsynchronous,
}

var mappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnumLowerCase = map[string]DatabaseToolsMcpToolsetDefaultExecutionTypeEnum{
	"synchronous":  DatabaseToolsMcpToolsetDefaultExecutionTypeSynchronous,
	"asynchronous": DatabaseToolsMcpToolsetDefaultExecutionTypeAsynchronous,
}

// GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumValues Enumerates the set of values for DatabaseToolsMcpToolsetDefaultExecutionTypeEnum
func GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumValues() []DatabaseToolsMcpToolsetDefaultExecutionTypeEnum {
	values := make([]DatabaseToolsMcpToolsetDefaultExecutionTypeEnum, 0)
	for _, v := range mappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsMcpToolsetDefaultExecutionTypeEnum
func GetDatabaseToolsMcpToolsetDefaultExecutionTypeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNCHRONOUS",
	}
}

// GetMappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnum(val string) (DatabaseToolsMcpToolsetDefaultExecutionTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsMcpToolsetDefaultExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
