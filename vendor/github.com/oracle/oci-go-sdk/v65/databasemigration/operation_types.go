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

// OperationTypesEnum Enum with underlying type: string
type OperationTypesEnum string

// Set of constants representing the allowable values for OperationTypesEnum
const (
	OperationTypesCreateMigration            OperationTypesEnum = "CREATE_MIGRATION"
	OperationTypesCloneMigration             OperationTypesEnum = "CLONE_MIGRATION"
	OperationTypesDeleteMigration            OperationTypesEnum = "DELETE_MIGRATION"
	OperationTypesUpdateMigration            OperationTypesEnum = "UPDATE_MIGRATION"
	OperationTypesStartMigration             OperationTypesEnum = "START_MIGRATION"
	OperationTypesValidateMigration          OperationTypesEnum = "VALIDATE_MIGRATION"
	OperationTypesCreateConnection           OperationTypesEnum = "CREATE_CONNECTION"
	OperationTypesDeleteConnection           OperationTypesEnum = "DELETE_CONNECTION"
	OperationTypesUpdateConnection           OperationTypesEnum = "UPDATE_CONNECTION"
	OperationTypesCreateAssessment           OperationTypesEnum = "CREATE_ASSESSMENT"
	OperationTypesDeleteAssessment           OperationTypesEnum = "DELETE_ASSESSMENT"
	OperationTypesUpdateAssessment           OperationTypesEnum = "UPDATE_ASSESSMENT"
	OperationTypesRunAssessorAction          OperationTypesEnum = "RUN_ASSESSOR_ACTION"
	OperationTypesRunMigrateTableAssessor    OperationTypesEnum = "RUN_MIGRATE_TABLE_ASSESSOR"
	OperationTypesRunCompatibilityAssessor   OperationTypesEnum = "RUN_COMPATIBILITY_ASSESSOR"
	OperationTypesRunViabilityAssessor       OperationTypesEnum = "RUN_VIABILITY_ASSESSOR"
	OperationTypesConfigureMigrationAssessor OperationTypesEnum = "CONFIGURE_MIGRATION_ASSESSOR"
	OperationTypesCreateMigrationAssessor    OperationTypesEnum = "CREATE_MIGRATION_ASSESSOR"
	OperationTypesPrepareSourceAssessor      OperationTypesEnum = "PREPARE_SOURCE_ASSESSOR"
	OperationTypesPrepareTargetAssessor      OperationTypesEnum = "PREPARE_TARGET_ASSESSOR"
	OperationTypesConfirmViabilityAssessor   OperationTypesEnum = "CONFIRM_VIABILITY_ASSESSOR"
)

var mappingOperationTypesEnum = map[string]OperationTypesEnum{
	"CREATE_MIGRATION":             OperationTypesCreateMigration,
	"CLONE_MIGRATION":              OperationTypesCloneMigration,
	"DELETE_MIGRATION":             OperationTypesDeleteMigration,
	"UPDATE_MIGRATION":             OperationTypesUpdateMigration,
	"START_MIGRATION":              OperationTypesStartMigration,
	"VALIDATE_MIGRATION":           OperationTypesValidateMigration,
	"CREATE_CONNECTION":            OperationTypesCreateConnection,
	"DELETE_CONNECTION":            OperationTypesDeleteConnection,
	"UPDATE_CONNECTION":            OperationTypesUpdateConnection,
	"CREATE_ASSESSMENT":            OperationTypesCreateAssessment,
	"DELETE_ASSESSMENT":            OperationTypesDeleteAssessment,
	"UPDATE_ASSESSMENT":            OperationTypesUpdateAssessment,
	"RUN_ASSESSOR_ACTION":          OperationTypesRunAssessorAction,
	"RUN_MIGRATE_TABLE_ASSESSOR":   OperationTypesRunMigrateTableAssessor,
	"RUN_COMPATIBILITY_ASSESSOR":   OperationTypesRunCompatibilityAssessor,
	"RUN_VIABILITY_ASSESSOR":       OperationTypesRunViabilityAssessor,
	"CONFIGURE_MIGRATION_ASSESSOR": OperationTypesConfigureMigrationAssessor,
	"CREATE_MIGRATION_ASSESSOR":    OperationTypesCreateMigrationAssessor,
	"PREPARE_SOURCE_ASSESSOR":      OperationTypesPrepareSourceAssessor,
	"PREPARE_TARGET_ASSESSOR":      OperationTypesPrepareTargetAssessor,
	"CONFIRM_VIABILITY_ASSESSOR":   OperationTypesConfirmViabilityAssessor,
}

var mappingOperationTypesEnumLowerCase = map[string]OperationTypesEnum{
	"create_migration":             OperationTypesCreateMigration,
	"clone_migration":              OperationTypesCloneMigration,
	"delete_migration":             OperationTypesDeleteMigration,
	"update_migration":             OperationTypesUpdateMigration,
	"start_migration":              OperationTypesStartMigration,
	"validate_migration":           OperationTypesValidateMigration,
	"create_connection":            OperationTypesCreateConnection,
	"delete_connection":            OperationTypesDeleteConnection,
	"update_connection":            OperationTypesUpdateConnection,
	"create_assessment":            OperationTypesCreateAssessment,
	"delete_assessment":            OperationTypesDeleteAssessment,
	"update_assessment":            OperationTypesUpdateAssessment,
	"run_assessor_action":          OperationTypesRunAssessorAction,
	"run_migrate_table_assessor":   OperationTypesRunMigrateTableAssessor,
	"run_compatibility_assessor":   OperationTypesRunCompatibilityAssessor,
	"run_viability_assessor":       OperationTypesRunViabilityAssessor,
	"configure_migration_assessor": OperationTypesConfigureMigrationAssessor,
	"create_migration_assessor":    OperationTypesCreateMigrationAssessor,
	"prepare_source_assessor":      OperationTypesPrepareSourceAssessor,
	"prepare_target_assessor":      OperationTypesPrepareTargetAssessor,
	"confirm_viability_assessor":   OperationTypesConfirmViabilityAssessor,
}

// GetOperationTypesEnumValues Enumerates the set of values for OperationTypesEnum
func GetOperationTypesEnumValues() []OperationTypesEnum {
	values := make([]OperationTypesEnum, 0)
	for _, v := range mappingOperationTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypesEnumStringValues Enumerates the set of values in String for OperationTypesEnum
func GetOperationTypesEnumStringValues() []string {
	return []string{
		"CREATE_MIGRATION",
		"CLONE_MIGRATION",
		"DELETE_MIGRATION",
		"UPDATE_MIGRATION",
		"START_MIGRATION",
		"VALIDATE_MIGRATION",
		"CREATE_CONNECTION",
		"DELETE_CONNECTION",
		"UPDATE_CONNECTION",
		"CREATE_ASSESSMENT",
		"DELETE_ASSESSMENT",
		"UPDATE_ASSESSMENT",
		"RUN_ASSESSOR_ACTION",
		"RUN_MIGRATE_TABLE_ASSESSOR",
		"RUN_COMPATIBILITY_ASSESSOR",
		"RUN_VIABILITY_ASSESSOR",
		"CONFIGURE_MIGRATION_ASSESSOR",
		"CREATE_MIGRATION_ASSESSOR",
		"PREPARE_SOURCE_ASSESSOR",
		"PREPARE_TARGET_ASSESSOR",
		"CONFIRM_VIABILITY_ASSESSOR",
	}
}

// GetMappingOperationTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypesEnum(val string) (OperationTypesEnum, bool) {
	enum, ok := mappingOperationTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
