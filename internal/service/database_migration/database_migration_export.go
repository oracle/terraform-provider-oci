package database_migration

//
//import (
//	oci_database_migration "github.com/oracle/oci-go-sdk/v65/databasemigration"
//
//	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
//)
//
//func init() {
//	tf_export.RegisterCompartmentGraphs("database_migration", databaseMigrationResourceGraph)
//}
//
//// Custom overrides for generating composite IDs within the resource discovery framework
//
//// Hints for discovering and exporting this resource to configuration and state files
//var exportDatabaseMigrationConnectionHints = &tf_export.TerraformResourceHints{
//	ResourceClass:          "oci_database_migration_connection",
//	DatasourceClass:        "oci_database_migration_connections",
//	DatasourceItemsAttr:    "connection_collection",
//	IsDatasourceCollection: true,
//	ResourceAbbreviation:   "connection",
//	RequireResourceRefresh: true,
//	DiscoverableLifecycleStates: []string{
//		string(oci_database_migration.LifecycleStatesActive),
//	},
//}
//
//var exportDatabaseMigrationMigrationHints = &tf_export.TerraformResourceHints{
//	ResourceClass:          "oci_database_migration_migration",
//	DatasourceClass:        "oci_database_migration_migrations",
//	DatasourceItemsAttr:    "migration_collection",
//	IsDatasourceCollection: true,
//	ResourceAbbreviation:   "migration",
//	RequireResourceRefresh: true,
//	DiscoverableLifecycleStates: []string{
//		string(oci_database_migration.LifecycleStatesActive),
//	},
//}
//
//var databaseMigrationResourceGraph = tf_export.TerraformResourceGraph{
//	"oci_identity_compartment": {
//		{TerraformResourceHints: exportDatabaseMigrationMigrationHints},
//		{TerraformResourceHints: exportDatabaseMigrationConnectionHints},
//	},
//}
