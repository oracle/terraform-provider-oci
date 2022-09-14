package oce

import (
	oci_oce "github.com/oracle/oci-go-sdk/v65/oce"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("oce", oceResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOceOceInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_oce_oce_instance",
	DatasourceClass:        "oci_oce_oce_instances",
	DatasourceItemsAttr:    "oce_instances",
	ResourceAbbreviation:   "oce_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_oce.LifecycleStateActive),
	},
}

var oceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOceOceInstanceHints},
	},
}
