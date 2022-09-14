package ons

import (
	"fmt"

	oci_ons "github.com/oracle/oci-go-sdk/v65/ons"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportOnsNotificationTopicHints.GetIdFn = getOnsNotificationTopicId
	tf_export.RegisterCompartmentGraphs("ons", onsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getOnsNotificationTopicId(resource *tf_export.OCIResource) (string, error) {
	id, ok := resource.SourceAttributes["topic_id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find topic id for ons notification topic")
	}
	return id, nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportOnsNotificationTopicHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_ons_notification_topic",
	DatasourceClass:      "oci_ons_notification_topics",
	DatasourceItemsAttr:  "notification_topics",
	ResourceAbbreviation: "notification_topic",
	DiscoverableLifecycleStates: []string{
		string(oci_ons.NotificationTopicLifecycleStateActive),
	},
}

var exportOnsSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_ons_subscription",
	DatasourceClass:      "oci_ons_subscriptions",
	DatasourceItemsAttr:  "subscriptions",
	ResourceAbbreviation: "subscription",
	DiscoverableLifecycleStates: []string{
		string(oci_ons.SubscriptionLifecycleStatePending),
		string(oci_ons.SubscriptionLifecycleStateActive),
	},
}

var onsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportOnsNotificationTopicHints},
		{TerraformResourceHints: exportOnsSubscriptionHints},
	},
}
