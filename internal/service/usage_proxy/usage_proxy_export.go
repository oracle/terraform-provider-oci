package usage_proxy

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportUsageProxySubscriptionRedeemableUserHints.GetIdFn = getUsageProxySubscriptionRedeemableUserId
	tf_export.RegisterTenancyGraphs("usage_proxy", usageProxyResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getUsageProxySubscriptionRedeemableUserId(resource *tf_export.OCIResource) (string, error) {

	subscriptionId := resource.Parent.Id
	tenancyId, ok := resource.Parent.SourceAttributes["tenancy_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find bucket for ObjectStorage ReplicationPolicy")
	}
	return GetSubscriptionRedeemableUserCompositeId(subscriptionId, tenancyId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportUsageProxySubscriptionRedeemableUserHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_usage_proxy_subscription_redeemable_user",
	DatasourceClass:        "oci_usage_proxy_subscription_redeemable_users",
	DatasourceItemsAttr:    "redeemable_user_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "subscription_redeemable_user",
}

var usageProxyResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
