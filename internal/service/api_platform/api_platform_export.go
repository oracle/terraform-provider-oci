package api_platform

import (
	oci_api_platform "github.com/oracle/oci-go-sdk/v65/apiplatform"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("api_platform", apiPlatformResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportApiPlatformApiPlatformInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_api_platform_api_platform_instance",
	DatasourceClass:        "oci_api_platform_api_platform_instances",
	DatasourceItemsAttr:    "api_platform_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "api_platform_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_api_platform.ApiPlatformInstanceLifecycleStateActive),
	},
}

var apiPlatformResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApiPlatformApiPlatformInstanceHints},
	},
}
