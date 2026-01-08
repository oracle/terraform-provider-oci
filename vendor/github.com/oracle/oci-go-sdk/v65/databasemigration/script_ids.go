// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ScriptIdsEnum Enum with underlying type: string
type ScriptIdsEnum string

// Set of constants representing the allowable values for ScriptIdsEnum
const (
	ScriptIdsUserCreationSqlScript ScriptIdsEnum = "USER_CREATION_SQL_SCRIPT"
)

var mappingScriptIdsEnum = map[string]ScriptIdsEnum{
	"USER_CREATION_SQL_SCRIPT": ScriptIdsUserCreationSqlScript,
}

var mappingScriptIdsEnumLowerCase = map[string]ScriptIdsEnum{
	"user_creation_sql_script": ScriptIdsUserCreationSqlScript,
}

// GetScriptIdsEnumValues Enumerates the set of values for ScriptIdsEnum
func GetScriptIdsEnumValues() []ScriptIdsEnum {
	values := make([]ScriptIdsEnum, 0)
	for _, v := range mappingScriptIdsEnum {
		values = append(values, v)
	}
	return values
}

// GetScriptIdsEnumStringValues Enumerates the set of values in String for ScriptIdsEnum
func GetScriptIdsEnumStringValues() []string {
	return []string{
		"USER_CREATION_SQL_SCRIPT",
	}
}

// GetMappingScriptIdsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScriptIdsEnum(val string) (ScriptIdsEnum, bool) {
	enum, ok := mappingScriptIdsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
