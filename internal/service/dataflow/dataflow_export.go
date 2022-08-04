package dataflow

import (
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dataflow", dataflowResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

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

var dataflowResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDataflowApplicationHints},
		{TerraformResourceHints: exportDataflowPrivateEndpointHints},
	},
}
