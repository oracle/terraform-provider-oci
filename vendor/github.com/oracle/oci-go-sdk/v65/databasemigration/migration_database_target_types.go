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

// MigrationDatabaseTargetTypesEnum Enum with underlying type: string
type MigrationDatabaseTargetTypesEnum string

// Set of constants representing the allowable values for MigrationDatabaseTargetTypesEnum
const (
	MigrationDatabaseTargetTypesAdbSRemap        MigrationDatabaseTargetTypesEnum = "ADB_S_REMAP"
	MigrationDatabaseTargetTypesAdbDRemap        MigrationDatabaseTargetTypesEnum = "ADB_D_REMAP"
	MigrationDatabaseTargetTypesAdbDAutocreate   MigrationDatabaseTargetTypesEnum = "ADB_D_AUTOCREATE"
	MigrationDatabaseTargetTypesNonAdbRemap      MigrationDatabaseTargetTypesEnum = "NON_ADB_REMAP"
	MigrationDatabaseTargetTypesNonAdbAutocreate MigrationDatabaseTargetTypesEnum = "NON_ADB_AUTOCREATE"
)

var mappingMigrationDatabaseTargetTypesEnum = map[string]MigrationDatabaseTargetTypesEnum{
	"ADB_S_REMAP":        MigrationDatabaseTargetTypesAdbSRemap,
	"ADB_D_REMAP":        MigrationDatabaseTargetTypesAdbDRemap,
	"ADB_D_AUTOCREATE":   MigrationDatabaseTargetTypesAdbDAutocreate,
	"NON_ADB_REMAP":      MigrationDatabaseTargetTypesNonAdbRemap,
	"NON_ADB_AUTOCREATE": MigrationDatabaseTargetTypesNonAdbAutocreate,
}

var mappingMigrationDatabaseTargetTypesEnumLowerCase = map[string]MigrationDatabaseTargetTypesEnum{
	"adb_s_remap":        MigrationDatabaseTargetTypesAdbSRemap,
	"adb_d_remap":        MigrationDatabaseTargetTypesAdbDRemap,
	"adb_d_autocreate":   MigrationDatabaseTargetTypesAdbDAutocreate,
	"non_adb_remap":      MigrationDatabaseTargetTypesNonAdbRemap,
	"non_adb_autocreate": MigrationDatabaseTargetTypesNonAdbAutocreate,
}

// GetMigrationDatabaseTargetTypesEnumValues Enumerates the set of values for MigrationDatabaseTargetTypesEnum
func GetMigrationDatabaseTargetTypesEnumValues() []MigrationDatabaseTargetTypesEnum {
	values := make([]MigrationDatabaseTargetTypesEnum, 0)
	for _, v := range mappingMigrationDatabaseTargetTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationDatabaseTargetTypesEnumStringValues Enumerates the set of values in String for MigrationDatabaseTargetTypesEnum
func GetMigrationDatabaseTargetTypesEnumStringValues() []string {
	return []string{
		"ADB_S_REMAP",
		"ADB_D_REMAP",
		"ADB_D_AUTOCREATE",
		"NON_ADB_REMAP",
		"NON_ADB_AUTOCREATE",
	}
}

// GetMappingMigrationDatabaseTargetTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMigrationDatabaseTargetTypesEnum(val string) (MigrationDatabaseTargetTypesEnum, bool) {
	enum, ok := mappingMigrationDatabaseTargetTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
