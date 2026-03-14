package self

import (
	oci_self "github.com/oracle/oci-go-sdk/v65/self"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("self", selfResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportSelfSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_self_subscription",
	DatasourceClass:        "oci_self_subscriptions",
	DatasourceItemsAttr:    "subscription_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "subscription",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_self.LifecycleStateEnumActive),
	},
}

var selfResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportSelfSubscriptionHints},
	},
}
