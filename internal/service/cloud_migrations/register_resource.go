// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_cloud_migrations_migration", CloudMigrationsMigrationResource())
	tfresource.RegisterResource("oci_cloud_migrations_migration_asset", CloudMigrationsMigrationAssetResource())
	tfresource.RegisterResource("oci_cloud_migrations_migration_plan", CloudMigrationsMigrationPlanResource())
	tfresource.RegisterResource("oci_cloud_migrations_replication_schedule", CloudMigrationsReplicationScheduleResource())
	tfresource.RegisterResource("oci_cloud_migrations_target_asset", CloudMigrationsTargetAssetResource())
}
