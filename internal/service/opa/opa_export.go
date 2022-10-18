package opa

import (
	oci_opa "github.com/oracle/oci-go-sdk/v65/opa"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("opa", opaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOpaOpaInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_opa_opa_instance",
	DatasourceClass:        "oci_opa_opa_instances",
	DatasourceItemsAttr:    "opa_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "opa_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_opa.OpaInstanceLifecycleStateActive),
	},
}

var opaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOpaOpaInstanceHints},
	},
}
