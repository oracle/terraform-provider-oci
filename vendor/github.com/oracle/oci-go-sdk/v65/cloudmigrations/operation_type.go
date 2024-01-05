// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"strings"
)

// OperationTypeEnum Enum with underlying type: string
type OperationTypeEnum string

// Set of constants representing the allowable values for OperationTypeEnum
const (
	OperationTypeCreateMigration           OperationTypeEnum = "CREATE_MIGRATION"
	OperationTypeUpdateMigration           OperationTypeEnum = "UPDATE_MIGRATION"
	OperationTypeRefreshMigration          OperationTypeEnum = "REFRESH_MIGRATION"
	OperationTypeDeleteMigration           OperationTypeEnum = "DELETE_MIGRATION"
	OperationTypeMoveMigration             OperationTypeEnum = "MOVE_MIGRATION"
	OperationTypeStartAssetReplication     OperationTypeEnum = "START_ASSET_REPLICATION"
	OperationTypeStartMigrationReplication OperationTypeEnum = "START_MIGRATION_REPLICATION"
	OperationTypeCreateReplicationSchedule OperationTypeEnum = "CREATE_REPLICATION_SCHEDULE"
	OperationTypeUpdateReplicationSchedule OperationTypeEnum = "UPDATE_REPLICATION_SCHEDULE"
	OperationTypeDeleteReplicationSchedule OperationTypeEnum = "DELETE_REPLICATION_SCHEDULE"
	OperationTypeMoveReplicationSchedule   OperationTypeEnum = "MOVE_REPLICATION_SCHEDULE"
	OperationTypeCreateMigrationPlan       OperationTypeEnum = "CREATE_MIGRATION_PLAN"
	OperationTypeUpdateMigrationPlan       OperationTypeEnum = "UPDATE_MIGRATION_PLAN"
	OperationTypeDeleteMigrationPlan       OperationTypeEnum = "DELETE_MIGRATION_PLAN"
	OperationTypeMoveMigrationPlan         OperationTypeEnum = "MOVE_MIGRATION_PLAN"
	OperationTypeRefreshMigrationPlan      OperationTypeEnum = "REFRESH_MIGRATION_PLAN"
	OperationTypeExecuteMigrationPlan      OperationTypeEnum = "EXECUTE_MIGRATION_PLAN"
	OperationTypeRefreshMigrationAsset     OperationTypeEnum = "REFRESH_MIGRATION_ASSET"
	OperationTypeCreateMigrationAsset      OperationTypeEnum = "CREATE_MIGRATION_ASSET"
	OperationTypeDeleteMigrationAsset      OperationTypeEnum = "DELETE_MIGRATION_ASSET"
	OperationTypeCreateTargetAsset         OperationTypeEnum = "CREATE_TARGET_ASSET"
	OperationTypeUpdateTargetAsset         OperationTypeEnum = "UPDATE_TARGET_ASSET"
	OperationTypeDeleteTargetAsset         OperationTypeEnum = "DELETE_TARGET_ASSET"
)

var mappingOperationTypeEnum = map[string]OperationTypeEnum{
	"CREATE_MIGRATION":            OperationTypeCreateMigration,
	"UPDATE_MIGRATION":            OperationTypeUpdateMigration,
	"REFRESH_MIGRATION":           OperationTypeRefreshMigration,
	"DELETE_MIGRATION":            OperationTypeDeleteMigration,
	"MOVE_MIGRATION":              OperationTypeMoveMigration,
	"START_ASSET_REPLICATION":     OperationTypeStartAssetReplication,
	"START_MIGRATION_REPLICATION": OperationTypeStartMigrationReplication,
	"CREATE_REPLICATION_SCHEDULE": OperationTypeCreateReplicationSchedule,
	"UPDATE_REPLICATION_SCHEDULE": OperationTypeUpdateReplicationSchedule,
	"DELETE_REPLICATION_SCHEDULE": OperationTypeDeleteReplicationSchedule,
	"MOVE_REPLICATION_SCHEDULE":   OperationTypeMoveReplicationSchedule,
	"CREATE_MIGRATION_PLAN":       OperationTypeCreateMigrationPlan,
	"UPDATE_MIGRATION_PLAN":       OperationTypeUpdateMigrationPlan,
	"DELETE_MIGRATION_PLAN":       OperationTypeDeleteMigrationPlan,
	"MOVE_MIGRATION_PLAN":         OperationTypeMoveMigrationPlan,
	"REFRESH_MIGRATION_PLAN":      OperationTypeRefreshMigrationPlan,
	"EXECUTE_MIGRATION_PLAN":      OperationTypeExecuteMigrationPlan,
	"REFRESH_MIGRATION_ASSET":     OperationTypeRefreshMigrationAsset,
	"CREATE_MIGRATION_ASSET":      OperationTypeCreateMigrationAsset,
	"DELETE_MIGRATION_ASSET":      OperationTypeDeleteMigrationAsset,
	"CREATE_TARGET_ASSET":         OperationTypeCreateTargetAsset,
	"UPDATE_TARGET_ASSET":         OperationTypeUpdateTargetAsset,
	"DELETE_TARGET_ASSET":         OperationTypeDeleteTargetAsset,
}

var mappingOperationTypeEnumLowerCase = map[string]OperationTypeEnum{
	"create_migration":            OperationTypeCreateMigration,
	"update_migration":            OperationTypeUpdateMigration,
	"refresh_migration":           OperationTypeRefreshMigration,
	"delete_migration":            OperationTypeDeleteMigration,
	"move_migration":              OperationTypeMoveMigration,
	"start_asset_replication":     OperationTypeStartAssetReplication,
	"start_migration_replication": OperationTypeStartMigrationReplication,
	"create_replication_schedule": OperationTypeCreateReplicationSchedule,
	"update_replication_schedule": OperationTypeUpdateReplicationSchedule,
	"delete_replication_schedule": OperationTypeDeleteReplicationSchedule,
	"move_replication_schedule":   OperationTypeMoveReplicationSchedule,
	"create_migration_plan":       OperationTypeCreateMigrationPlan,
	"update_migration_plan":       OperationTypeUpdateMigrationPlan,
	"delete_migration_plan":       OperationTypeDeleteMigrationPlan,
	"move_migration_plan":         OperationTypeMoveMigrationPlan,
	"refresh_migration_plan":      OperationTypeRefreshMigrationPlan,
	"execute_migration_plan":      OperationTypeExecuteMigrationPlan,
	"refresh_migration_asset":     OperationTypeRefreshMigrationAsset,
	"create_migration_asset":      OperationTypeCreateMigrationAsset,
	"delete_migration_asset":      OperationTypeDeleteMigrationAsset,
	"create_target_asset":         OperationTypeCreateTargetAsset,
	"update_target_asset":         OperationTypeUpdateTargetAsset,
	"delete_target_asset":         OperationTypeDeleteTargetAsset,
}

// GetOperationTypeEnumValues Enumerates the set of values for OperationTypeEnum
func GetOperationTypeEnumValues() []OperationTypeEnum {
	values := make([]OperationTypeEnum, 0)
	for _, v := range mappingOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOperationTypeEnumStringValues Enumerates the set of values in String for OperationTypeEnum
func GetOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_MIGRATION",
		"UPDATE_MIGRATION",
		"REFRESH_MIGRATION",
		"DELETE_MIGRATION",
		"MOVE_MIGRATION",
		"START_ASSET_REPLICATION",
		"START_MIGRATION_REPLICATION",
		"CREATE_REPLICATION_SCHEDULE",
		"UPDATE_REPLICATION_SCHEDULE",
		"DELETE_REPLICATION_SCHEDULE",
		"MOVE_REPLICATION_SCHEDULE",
		"CREATE_MIGRATION_PLAN",
		"UPDATE_MIGRATION_PLAN",
		"DELETE_MIGRATION_PLAN",
		"MOVE_MIGRATION_PLAN",
		"REFRESH_MIGRATION_PLAN",
		"EXECUTE_MIGRATION_PLAN",
		"REFRESH_MIGRATION_ASSET",
		"CREATE_MIGRATION_ASSET",
		"DELETE_MIGRATION_ASSET",
		"CREATE_TARGET_ASSET",
		"UPDATE_TARGET_ASSET",
		"DELETE_TARGET_ASSET",
	}
}

// GetMappingOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOperationTypeEnum(val string) (OperationTypeEnum, bool) {
	enum, ok := mappingOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
