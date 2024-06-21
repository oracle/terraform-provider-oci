// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"strings"
)

// DatabaseCombinationEnum Enum with underlying type: string
type DatabaseCombinationEnum string

// Set of constants representing the allowable values for DatabaseCombinationEnum
const (
	DatabaseCombinationMysql  DatabaseCombinationEnum = "MYSQL"
	DatabaseCombinationOracle DatabaseCombinationEnum = "ORACLE"
)

var mappingDatabaseCombinationEnum = map[string]DatabaseCombinationEnum{
	"MYSQL":  DatabaseCombinationMysql,
	"ORACLE": DatabaseCombinationOracle,
}

var mappingDatabaseCombinationEnumLowerCase = map[string]DatabaseCombinationEnum{
	"mysql":  DatabaseCombinationMysql,
	"oracle": DatabaseCombinationOracle,
}

// GetDatabaseCombinationEnumValues Enumerates the set of values for DatabaseCombinationEnum
func GetDatabaseCombinationEnumValues() []DatabaseCombinationEnum {
	values := make([]DatabaseCombinationEnum, 0)
	for _, v := range mappingDatabaseCombinationEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseCombinationEnumStringValues Enumerates the set of values in String for DatabaseCombinationEnum
func GetDatabaseCombinationEnumStringValues() []string {
	return []string{
		"MYSQL",
		"ORACLE",
	}
}

// GetMappingDatabaseCombinationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseCombinationEnum(val string) (DatabaseCombinationEnum, bool) {
	enum, ok := mappingDatabaseCombinationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
