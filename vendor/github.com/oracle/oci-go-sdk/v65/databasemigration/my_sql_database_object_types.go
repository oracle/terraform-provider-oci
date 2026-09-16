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

// MySqlDatabaseObjectTypesEnum Enum with underlying type: string
type MySqlDatabaseObjectTypesEnum string

// Set of constants representing the allowable values for MySqlDatabaseObjectTypesEnum
const (
	MySqlDatabaseObjectTypesTable        MySqlDatabaseObjectTypesEnum = "TABLE"
	MySqlDatabaseObjectTypesView         MySqlDatabaseObjectTypesEnum = "VIEW"
	MySqlDatabaseObjectTypesIndex        MySqlDatabaseObjectTypesEnum = "INDEX"
	MySqlDatabaseObjectTypesTrigger      MySqlDatabaseObjectTypesEnum = "TRIGGER"
	MySqlDatabaseObjectTypesProcedure    MySqlDatabaseObjectTypesEnum = "PROCEDURE"
	MySqlDatabaseObjectTypesFunction     MySqlDatabaseObjectTypesEnum = "FUNCTION"
	MySqlDatabaseObjectTypesEvent        MySqlDatabaseObjectTypesEnum = "EVENT"
	MySqlDatabaseObjectTypesDatabase     MySqlDatabaseObjectTypesEnum = "DATABASE"
	MySqlDatabaseObjectTypesSchema       MySqlDatabaseObjectTypesEnum = "SCHEMA"
	MySqlDatabaseObjectTypesTablespace   MySqlDatabaseObjectTypesEnum = "TABLESPACE"
	MySqlDatabaseObjectTypesServer       MySqlDatabaseObjectTypesEnum = "SERVER"
	MySqlDatabaseObjectTypesLogfileGroup MySqlDatabaseObjectTypesEnum = "LOGFILE_GROUP"
	MySqlDatabaseObjectTypesUser         MySqlDatabaseObjectTypesEnum = "USER"
	MySqlDatabaseObjectTypesRole         MySqlDatabaseObjectTypesEnum = "ROLE"
)

var mappingMySqlDatabaseObjectTypesEnum = map[string]MySqlDatabaseObjectTypesEnum{
	"TABLE":         MySqlDatabaseObjectTypesTable,
	"VIEW":          MySqlDatabaseObjectTypesView,
	"INDEX":         MySqlDatabaseObjectTypesIndex,
	"TRIGGER":       MySqlDatabaseObjectTypesTrigger,
	"PROCEDURE":     MySqlDatabaseObjectTypesProcedure,
	"FUNCTION":      MySqlDatabaseObjectTypesFunction,
	"EVENT":         MySqlDatabaseObjectTypesEvent,
	"DATABASE":      MySqlDatabaseObjectTypesDatabase,
	"SCHEMA":        MySqlDatabaseObjectTypesSchema,
	"TABLESPACE":    MySqlDatabaseObjectTypesTablespace,
	"SERVER":        MySqlDatabaseObjectTypesServer,
	"LOGFILE_GROUP": MySqlDatabaseObjectTypesLogfileGroup,
	"USER":          MySqlDatabaseObjectTypesUser,
	"ROLE":          MySqlDatabaseObjectTypesRole,
}

var mappingMySqlDatabaseObjectTypesEnumLowerCase = map[string]MySqlDatabaseObjectTypesEnum{
	"table":         MySqlDatabaseObjectTypesTable,
	"view":          MySqlDatabaseObjectTypesView,
	"index":         MySqlDatabaseObjectTypesIndex,
	"trigger":       MySqlDatabaseObjectTypesTrigger,
	"procedure":     MySqlDatabaseObjectTypesProcedure,
	"function":      MySqlDatabaseObjectTypesFunction,
	"event":         MySqlDatabaseObjectTypesEvent,
	"database":      MySqlDatabaseObjectTypesDatabase,
	"schema":        MySqlDatabaseObjectTypesSchema,
	"tablespace":    MySqlDatabaseObjectTypesTablespace,
	"server":        MySqlDatabaseObjectTypesServer,
	"logfile_group": MySqlDatabaseObjectTypesLogfileGroup,
	"user":          MySqlDatabaseObjectTypesUser,
	"role":          MySqlDatabaseObjectTypesRole,
}

// GetMySqlDatabaseObjectTypesEnumValues Enumerates the set of values for MySqlDatabaseObjectTypesEnum
func GetMySqlDatabaseObjectTypesEnumValues() []MySqlDatabaseObjectTypesEnum {
	values := make([]MySqlDatabaseObjectTypesEnum, 0)
	for _, v := range mappingMySqlDatabaseObjectTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMySqlDatabaseObjectTypesEnumStringValues Enumerates the set of values in String for MySqlDatabaseObjectTypesEnum
func GetMySqlDatabaseObjectTypesEnumStringValues() []string {
	return []string{
		"TABLE",
		"VIEW",
		"INDEX",
		"TRIGGER",
		"PROCEDURE",
		"FUNCTION",
		"EVENT",
		"DATABASE",
		"SCHEMA",
		"TABLESPACE",
		"SERVER",
		"LOGFILE_GROUP",
		"USER",
		"ROLE",
	}
}

// GetMappingMySqlDatabaseObjectTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMySqlDatabaseObjectTypesEnum(val string) (MySqlDatabaseObjectTypesEnum, bool) {
	enum, ok := mappingMySqlDatabaseObjectTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
