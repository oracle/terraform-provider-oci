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

// MigrationScopeOracleEnum Enum with underlying type: string
type MigrationScopeOracleEnum string

// Set of constants representing the allowable values for MigrationScopeOracleEnum
const (
	MigrationScopeOracleSchema MigrationScopeOracleEnum = "SCHEMA"
	MigrationScopeOracleFull   MigrationScopeOracleEnum = "FULL"
)

var mappingMigrationScopeOracleEnum = map[string]MigrationScopeOracleEnum{
	"SCHEMA": MigrationScopeOracleSchema,
	"FULL":   MigrationScopeOracleFull,
}

var mappingMigrationScopeOracleEnumLowerCase = map[string]MigrationScopeOracleEnum{
	"schema": MigrationScopeOracleSchema,
	"full":   MigrationScopeOracleFull,
}

// GetMigrationScopeOracleEnumValues Enumerates the set of values for MigrationScopeOracleEnum
func GetMigrationScopeOracleEnumValues() []MigrationScopeOracleEnum {
	values := make([]MigrationScopeOracleEnum, 0)
	for _, v := range mappingMigrationScopeOracleEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationScopeOracleEnumStringValues Enumerates the set of values in String for MigrationScopeOracleEnum
func GetMigrationScopeOracleEnumStringValues() []string {
	return []string{
		"SCHEMA",
		"FULL",
	}
}

// GetMappingMigrationScopeOracleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationScopeOracleEnum(val string) (MigrationScopeOracleEnum, bool) {
	enum, ok := mappingMigrationScopeOracleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
