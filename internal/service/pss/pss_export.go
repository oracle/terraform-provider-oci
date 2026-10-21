package pss

import (
	oci_pss "github.com/oracle/oci-go-sdk/v65/pss"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("pss", pssResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportPssInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_pss_instance",
	DatasourceClass:        "oci_pss_instances",
	DatasourceItemsAttr:    "instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_pss.InstanceLifecycleStateActive),
	},
}

var pssResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportPssInstanceHints},
	},
}
