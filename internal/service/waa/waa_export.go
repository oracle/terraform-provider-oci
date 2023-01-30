package waa

import (
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("waa", waaResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportWaaWebAppAccelerationPolicyHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waa_web_app_acceleration_policy",
	DatasourceClass:        "oci_waa_web_app_acceleration_policies",
	DatasourceItemsAttr:    "web_app_acceleration_policy_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "web_app_acceleration_policy",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waa.WebAppAccelerationPolicyLifecycleStateActive),
	},
}

var exportWaaWebAppAccelerationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_waa_web_app_acceleration",
	DatasourceClass:        "oci_waa_web_app_accelerations",
	DatasourceItemsAttr:    "web_app_acceleration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "web_app_acceleration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_waa.WebAppAccelerationLifecycleStateActive),
	},
}

var waaResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportWaaWebAppAccelerationPolicyHints},
		{TerraformResourceHints: exportWaaWebAppAccelerationHints},
	},
}
