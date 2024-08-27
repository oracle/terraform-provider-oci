package delegate_access_control

import (
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("delegate_access_control", delegateAccessControlResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportDelegateAccessControlDelegationSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_delegate_access_control_delegation_subscription",
	DatasourceClass:        "oci_delegate_access_control_delegation_subscriptions",
	DatasourceItemsAttr:    "delegation_subscription_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "delegation_subscription",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_delegate_access_control.DelegationSubscriptionLifecycleStateActive),
	},
}

var exportDelegateAccessControlDelegationControlHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_delegate_access_control_delegation_control",
	DatasourceClass:        "oci_delegate_access_control_delegation_controls",
	DatasourceItemsAttr:    "delegation_control_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "delegation_control",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_delegate_access_control.DelegationControlLifecycleStateActive),
		string(oci_delegate_access_control.DelegationControlLifecycleStateNeedsAttention),
	},
}

var delegateAccessControlResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDelegateAccessControlDelegationSubscriptionHints},
		{TerraformResourceHints: exportDelegateAccessControlDelegationControlHints},
	},
}
