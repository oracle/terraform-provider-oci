// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"strings"
)

// DatabaseConsoleStatusEnum Enum with underlying type: string
type DatabaseConsoleStatusEnum string

// Set of constants representing the allowable values for DatabaseConsoleStatusEnum
const (
	DatabaseConsoleStatusEnabled  DatabaseConsoleStatusEnum = "ENABLED"
	DatabaseConsoleStatusDisabled DatabaseConsoleStatusEnum = "DISABLED"
)

var mappingDatabaseConsoleStatusEnum = map[string]DatabaseConsoleStatusEnum{
	"ENABLED":  DatabaseConsoleStatusEnabled,
	"DISABLED": DatabaseConsoleStatusDisabled,
}

var mappingDatabaseConsoleStatusEnumLowerCase = map[string]DatabaseConsoleStatusEnum{
	"enabled":  DatabaseConsoleStatusEnabled,
	"disabled": DatabaseConsoleStatusDisabled,
}

// GetDatabaseConsoleStatusEnumValues Enumerates the set of values for DatabaseConsoleStatusEnum
func GetDatabaseConsoleStatusEnumValues() []DatabaseConsoleStatusEnum {
	values := make([]DatabaseConsoleStatusEnum, 0)
	for _, v := range mappingDatabaseConsoleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseConsoleStatusEnumStringValues Enumerates the set of values in String for DatabaseConsoleStatusEnum
func GetDatabaseConsoleStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingDatabaseConsoleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseConsoleStatusEnum(val string) (DatabaseConsoleStatusEnum, bool) {
	enum, ok := mappingDatabaseConsoleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
