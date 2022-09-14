package autoscaling

import (
	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("auto_scaling", autoScalingResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAutoScalingAutoScalingConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_autoscaling_auto_scaling_configuration",
	DatasourceClass:        "oci_autoscaling_auto_scaling_configurations",
	DatasourceItemsAttr:    "auto_scaling_configurations",
	ResourceAbbreviation:   "auto_scaling_configuration",
	RequireResourceRefresh: true,
}

var autoScalingResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAutoScalingAutoScalingConfigurationHints},
	},
}
