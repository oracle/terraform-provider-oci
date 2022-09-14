package osp_gateway

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportOspGatewaySubscriptionHints.GetIdFn = getOspGatewaySubscriptionId
	tf_export.RegisterCompartmentGraphs("osp_gateway", ospGatewayResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getOspGatewaySubscriptionId(resource *tf_export.OCIResource) (string, error) {

	subscriptionId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find subscriptionId for OspGateway Subscription")
	}
	return GetSubscriptionCompositeId(subscriptionId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportOspGatewaySubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_osp_gateway_subscription",
	DatasourceClass:        "oci_osp_gateway_subscriptions",
	DatasourceItemsAttr:    "subscription_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "subscription",
}

var ospGatewayResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
