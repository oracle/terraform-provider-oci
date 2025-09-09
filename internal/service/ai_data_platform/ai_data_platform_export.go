package ai_data_platform

import (
	oci_ai_data_platform "github.com/oracle/oci-go-sdk/v65/aidataplatform"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("ai_data_platform", aiDataPlatformResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAiDataPlatformAiDataPlatformHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_ai_data_platform_ai_data_platform",
	DatasourceClass:        "oci_ai_data_platform_ai_data_platforms",
	DatasourceItemsAttr:    "ai_data_platform_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "ai_data_platform",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_ai_data_platform.AiDataPlatformLifecycleStateActive),
	},
}

var aiDataPlatformResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAiDataPlatformAiDataPlatformHints},
	},
}
