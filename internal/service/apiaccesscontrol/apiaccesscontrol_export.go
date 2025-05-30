package apiaccesscontrol

import (
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("apiaccesscontrol", apiaccesscontrolResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportApiaccesscontrolPrivilegedApiControlHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apiaccesscontrol_privileged_api_control",
	DatasourceClass:        "oci_apiaccesscontrol_privileged_api_controls",
	DatasourceItemsAttr:    "privileged_api_control_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "privileged_api_control",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateActive),
		string(oci_apiaccesscontrol.PrivilegedApiControlLifecycleStateNeedsAttention),
	},
}

var exportApiaccesscontrolPrivilegedApiRequestHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_apiaccesscontrol_privileged_api_request",
	DatasourceClass:        "oci_apiaccesscontrol_privileged_api_requests",
	DatasourceItemsAttr:    "privileged_api_request_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "privileged_api_request",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_apiaccesscontrol.PrivilegedApiRequestLifecycleStateAccepted),
	},
}

var apiaccesscontrolResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportApiaccesscontrolPrivilegedApiControlHints},
		{TerraformResourceHints: exportApiaccesscontrolPrivilegedApiRequestHints},
	},
}
