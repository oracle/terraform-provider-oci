package cloud_migrations

import (
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("cloud_migrations", cloudMigrationsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportCloudMigrationsMigrationAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_migrations_migration_asset",
	DatasourceClass:        "oci_cloud_migrations_migration_assets",
	DatasourceItemsAttr:    "migration_asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "migration_asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_migrations.MigrationAssetLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationAssetLifecycleStateActive),
	},
}

var exportCloudMigrationsMigrationPlanHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_migrations_migration_plan",
	DatasourceClass:        "oci_cloud_migrations_migration_plans",
	DatasourceItemsAttr:    "migration_plan_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "migration_plan",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_migrations.MigrationPlanLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationPlanLifecycleStateActive),
	},
}

var exportCloudMigrationsTargetAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_migrations_target_asset",
	DatasourceClass:        "oci_cloud_migrations_target_assets",
	DatasourceItemsAttr:    "target_asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "target_asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_migrations.TargetAssetLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.TargetAssetLifecycleStateActive),
	},
}

var exportCloudMigrationsMigrationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_migrations_migration",
	DatasourceClass:        "oci_cloud_migrations_migrations",
	DatasourceItemsAttr:    "migration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "migration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_migrations.MigrationLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationLifecycleStateActive),
	},
}

var exportCloudMigrationsReplicationScheduleHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_cloud_migrations_replication_schedule",
	DatasourceClass:        "oci_cloud_migrations_replication_schedules",
	DatasourceItemsAttr:    "replication_schedule_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "replication_schedule",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_cloud_migrations.ReplicationScheduleLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.ReplicationScheduleLifecycleStateActive),
	},
}

var cloudMigrationsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportCloudMigrationsMigrationHints},
		{TerraformResourceHints: exportCloudMigrationsReplicationScheduleHints},
		{TerraformResourceHints: exportCloudMigrationsMigrationPlanHints},
	},
	"oci_cloud_migrations_migration_plan": {
		{
			TerraformResourceHints: exportCloudMigrationsTargetAssetHints,
			DatasourceQueryParams: map[string]string{
				"migration_plan_id": "id",
			},
		},
	},
	"oci_cloud_migrations_migration": {
		{
			TerraformResourceHints: exportCloudMigrationsMigrationAssetHints,
			DatasourceQueryParams: map[string]string{
				"migration_id": "id",
			},
		},
	},
}
