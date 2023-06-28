package dataflow

import (
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dataflow", dataflowResourceGraph)
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDataflowApplicationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataflow_application",
	DatasourceClass:        "oci_dataflow_applications",
	DatasourceItemsAttr:    "applications",
	ResourceAbbreviation:   "application",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataflow.ApplicationLifecycleStateActive),
	},
}

var exportDataflowPrivateEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataflow_private_endpoint",
	DatasourceClass:        "oci_dataflow_private_endpoints",
	DatasourceItemsAttr:    "private_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "private_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataflow.PrivateEndpointLifecycleStateActive),
		string(oci_dataflow.PrivateEndpointLifecycleStateInactive),
	},
}

var exportDataflowRunStatementHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataflow_run_statement",
	DatasourceClass:        "oci_dataflow_run_statements",
	DatasourceItemsAttr:    "statement_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "run_statement",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataflow.StatementLifecycleStateSucceeded),
	},
}

var exportDataflowPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataflow_pool",
	DatasourceClass:        "oci_dataflow_pools",
	DatasourceItemsAttr:    "pool_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "pool",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataflow.PoolLifecycleStateActive),
	},
}
var exportDataflowSqlEndpointHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dataflow_sql_endpoint",
	DatasourceClass:        "oci_dataflow_sql_endpoints",
	DatasourceItemsAttr:    "sql_endpoint_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "sql_endpoint",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dataflow.SqlEndpointLifecycleStateActive),
	},
}

var dataflowResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataflowApplicationHints},
		{TerraformResourceHints: exportDataflowPrivateEndpointHints},
		{TerraformResourceHints: exportDataflowPoolHints},
		{TerraformResourceHints: exportDataflowSqlEndpointHints},
	},
}
