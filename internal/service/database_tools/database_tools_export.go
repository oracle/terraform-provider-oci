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

var exportDatabaseToolsDatabaseToolsIdentityHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_identity",
	DatasourceClass:        "oci_database_tools_database_tools_identities",
	DatasourceItemsAttr:    "database_tools_identity_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_identity",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateActive),
		string(oci_database_tools.DatabaseToolsIdentityLifecycleStateNeedsAttention),
	},
}

var exportDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_database_api_gateway_config",
	DatasourceClass:        "oci_database_tools_database_tools_database_api_gateway_configs",
	DatasourceItemsAttr:    "database_tools_database_api_gateway_config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_database_api_gateway_config",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.DatabaseToolsDatabaseApiGatewayConfigLifecycleStateActive),
	},
}

var exportDatabaseToolsDatabaseToolsMcpToolsetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_mcp_toolset",
	DatasourceClass:        "oci_database_tools_database_tools_mcp_toolsets",
	DatasourceItemsAttr:    "database_tools_mcp_toolset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_mcp_toolset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.DatabaseToolsMcpToolsetLifecycleStateActive),
	},
}

var exportDatabaseToolsDatabaseToolsSqlReportHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_sql_report",
	DatasourceClass:        "oci_database_tools_database_tools_sql_reports",
	DatasourceItemsAttr:    "database_tools_sql_report_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_sql_report",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.DatabaseToolsSqlReportLifecycleStateActive),
	},
}

var exportDatabaseToolsDatabaseToolsMcpServerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_database_tools_mcp_server",
	DatasourceClass:        "oci_database_tools_database_tools_mcp_servers",
	DatasourceItemsAttr:    "database_tools_mcp_server_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_mcp_server",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateActive),
		string(oci_database_tools.DatabaseToolsMcpServerLifecycleStateNeedsAttention),
	},
}

var databaseToolsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsPrivateEndpointHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsConnectionHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsIdentityHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsDatabaseApiGatewayConfigHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsMcpToolsetHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsSqlReportHints},
		{TerraformResourceHints: exportDatabaseToolsDatabaseToolsMcpServerHints},
	},
}
