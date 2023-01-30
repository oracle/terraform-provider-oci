package database_tools

import (
	oci_database_tools "github.com/oracle/oci-go-sdk/v65/databasetools"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("database_tools", databaseToolsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDatabaseToolsDatabaseToolsPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_private_endpoint",
	DatasourceClass:        "oci_database_tools_database_tools_private_endpoints",
	DatasourceItemsAttr:    "database_tools_private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.LifecycleStateActive),
	},
}

var exportDatabaseToolsDatabaseToolsConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_connection",
	DatasourceClass:        "oci_database_tools_database_tools_connections",
	DatasourceItemsAttr:    "database_tools_connection_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_connection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.LifecycleStateActive),
	},
}

var databaseToolsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsPrivateEndpointHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsConnectionHints},
	},
}
