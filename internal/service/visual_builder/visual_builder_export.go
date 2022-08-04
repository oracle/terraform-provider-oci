package visual_builder

import (
	oci_visual_builder "github.com/oracle/oci-go-sdk/v65/visualbuilder"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("visual_builder", visualBuilderResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportVisualBuilderVbInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_visual_builder_vb_instance",
	DatasourceClass:        "oci_visual_builder_vb_instances",
	DatasourceItemsAttr:    "vb_instance_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "vb_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_visual_builder.VbInstanceLifecycleStateActive),
	},
}

var visualBuilderResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVisualBuilderVbInstanceHints},
	},
}
