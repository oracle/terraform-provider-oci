package sch

import (
	oci_sch "github.com/oracle/oci-go-sdk/v65/sch"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("sch", schResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportSchServiceConnectorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_sch_service_connector",
	DatasourceClass:        "oci_sch_service_connectors",
	DatasourceItemsAttr:    "service_connector_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "service_connector",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_sch.LifecycleStateActive),
	},
}

var schResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportSchServiceConnectorHints},
	},
}
