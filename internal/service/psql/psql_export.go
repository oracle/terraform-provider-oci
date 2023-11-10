package psql

import (
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("psql", psqlResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportPsqlConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_psql_configuration",
	DatasourceClass:        "oci_psql_configurations",
	DatasourceItemsAttr:    "configuration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_psql.ConfigurationLifecycleStateActive),
	},
}

var exportPsqlDbSystemHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_psql_db_system",
	DatasourceClass:        "oci_psql_db_systems",
	DatasourceItemsAttr:    "db_system_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "db_system",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_psql.DbSystemLifecycleStateActive),
		string(oci_psql.DbSystemLifecycleStateNeedsAttention),
	},
}

var exportPsqlBackupHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_psql_backup",
	DatasourceClass:        "oci_psql_backups",
	DatasourceItemsAttr:    "backup_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "backup",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_psql.BackupLifecycleStateActive),
	},
}

var psqlResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportPsqlConfigurationHints},
		{TerraformResourceHints: exportPsqlDbSystemHints},
		{TerraformResourceHints: exportPsqlBackupHints},
	},
}
