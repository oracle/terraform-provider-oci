package gdp

import (
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("gdp", gdpResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportGdpGdpPipelineHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_gdp_gdp_pipeline",
	DatasourceClass:        "oci_gdp_gdp_pipelines",
	DatasourceItemsAttr:    "gdp_pipeline_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "gdp_pipeline",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_gdp.GdpPipelineLifecycleStateActive),
		string(oci_gdp.GdpPipelineLifecycleStateNeedsAttention),
	},
}

var gdpResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportGdpGdpPipelineHints},
	},
}
