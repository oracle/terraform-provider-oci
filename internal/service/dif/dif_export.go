package dif

import (
	oci_dif "github.com/oracle/oci-go-sdk/v65/dif"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("dif", difResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDifStackHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_dif_stack",
	DatasourceClass:        "oci_dif_stacks",
	DatasourceItemsAttr:    "stack_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stack",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_dif.StackLifecycleStateActive),
		string(oci_dif.StackLifecycleStateNeedsAttention),
	},
}

var difResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDifStackHints},
	},
}
