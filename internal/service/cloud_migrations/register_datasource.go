// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_cloud_migrations_migration", CloudMigrationsMigrationDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_asset", CloudMigrationsMigrationAssetDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_assets", CloudMigrationsMigrationAssetsDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_plan", CloudMigrationsMigrationPlanDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_plan_available_shape", CloudMigrationsMigrationPlanAvailableShapeDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_plan_available_shapes", CloudMigrationsMigrationPlanAvailableShapesDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migration_plans", CloudMigrationsMigrationPlansDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_migrations", CloudMigrationsMigrationsDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_replication_schedule", CloudMigrationsReplicationScheduleDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_replication_schedules", CloudMigrationsReplicationSchedulesDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_target_asset", CloudMigrationsTargetAssetDataSource())
	tfresource.RegisterDatasource("oci_cloud_migrations_target_assets", CloudMigrationsTargetAssetsDataSource())
}
