package oda

import (
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("oda", odaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportOdaOdaInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_oda_oda_instance",
	DatasourceClass:      "oci_oda_oda_instances",
	DatasourceItemsAttr:  "oda_instances",
	ResourceAbbreviation: "oda_instance",
	DiscoverableLifecycleStates: []string{
		string(oci_oda.OdaInstanceLifecycleStateActive),
	},
}

var odaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOdaOdaInstanceHints},
	},
}
