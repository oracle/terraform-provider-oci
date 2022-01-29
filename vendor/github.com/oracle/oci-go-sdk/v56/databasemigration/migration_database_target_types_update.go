// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

// MigrationDatabaseTargetTypesUpdateEnum Enum with underlying type: string
type MigrationDatabaseTargetTypesUpdateEnum string

// Set of constants representing the allowable values for MigrationDatabaseTargetTypesUpdateEnum
const (
	MigrationDatabaseTargetTypesUpdateAdbSRemap                MigrationDatabaseTargetTypesUpdateEnum = "ADB_S_REMAP"
	MigrationDatabaseTargetTypesUpdateAdbDRemap                MigrationDatabaseTargetTypesUpdateEnum = "ADB_D_REMAP"
	MigrationDatabaseTargetTypesUpdateAdbDAutocreate           MigrationDatabaseTargetTypesUpdateEnum = "ADB_D_AUTOCREATE"
	MigrationDatabaseTargetTypesUpdateNonAdbRemap              MigrationDatabaseTargetTypesUpdateEnum = "NON_ADB_REMAP"
	MigrationDatabaseTargetTypesUpdateNonAdbAutocreate         MigrationDatabaseTargetTypesUpdateEnum = "NON_ADB_AUTOCREATE"
	MigrationDatabaseTargetTypesUpdateTargetDefaultsRemap      MigrationDatabaseTargetTypesUpdateEnum = "TARGET_DEFAULTS_REMAP"
	MigrationDatabaseTargetTypesUpdateTargetDefaultsAutocreate MigrationDatabaseTargetTypesUpdateEnum = "TARGET_DEFAULTS_AUTOCREATE"
)

var mappingMigrationDatabaseTargetTypesUpdateEnum = map[string]MigrationDatabaseTargetTypesUpdateEnum{
	"ADB_S_REMAP":                MigrationDatabaseTargetTypesUpdateAdbSRemap,
	"ADB_D_REMAP":                MigrationDatabaseTargetTypesUpdateAdbDRemap,
	"ADB_D_AUTOCREATE":           MigrationDatabaseTargetTypesUpdateAdbDAutocreate,
	"NON_ADB_REMAP":              MigrationDatabaseTargetTypesUpdateNonAdbRemap,
	"NON_ADB_AUTOCREATE":         MigrationDatabaseTargetTypesUpdateNonAdbAutocreate,
	"TARGET_DEFAULTS_REMAP":      MigrationDatabaseTargetTypesUpdateTargetDefaultsRemap,
	"TARGET_DEFAULTS_AUTOCREATE": MigrationDatabaseTargetTypesUpdateTargetDefaultsAutocreate,
}

// GetMigrationDatabaseTargetTypesUpdateEnumValues Enumerates the set of values for MigrationDatabaseTargetTypesUpdateEnum
func GetMigrationDatabaseTargetTypesUpdateEnumValues() []MigrationDatabaseTargetTypesUpdateEnum {
	values := make([]MigrationDatabaseTargetTypesUpdateEnum, 0)
	for _, v := range mappingMigrationDatabaseTargetTypesUpdateEnum {
		values = append(values, v)
	}
	return values
}

// GetMigrationDatabaseTargetTypesUpdateEnumStringValues Enumerates the set of values in String for MigrationDatabaseTargetTypesUpdateEnum
func GetMigrationDatabaseTargetTypesUpdateEnumStringValues() []string {
	return []string{
		"ADB_S_REMAP",
		"ADB_D_REMAP",
		"ADB_D_AUTOCREATE",
		"NON_ADB_REMAP",
		"NON_ADB_AUTOCREATE",
		"TARGET_DEFAULTS_REMAP",
		"TARGET_DEFAULTS_AUTOCREATE",
	}
}
