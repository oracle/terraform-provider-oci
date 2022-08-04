package announcements_service

import (
	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("announcements_service", announcementsServiceResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportAnnouncementsServiceAnnouncementSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_announcements_service_announcement_subscription",
	DatasourceClass:        "oci_announcements_service_announcement_subscriptions",
	DatasourceItemsAttr:    "announcement_subscription_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "announcement_subscription",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_announcements_service.AnnouncementSubscriptionLifecycleStateActive),
	},
}

var exportAnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_announcements_service_announcement_subscriptions_actions_change_compartment",
	ResourceAbbreviation: "announcement_subscriptions_actions_change_compartment",
}

var exportAnnouncementsServiceAnnouncementSubscriptionsFilterGroupHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_announcements_service_announcement_subscriptions_filter_group",
	ResourceAbbreviation: "announcement_subscriptions_filter_group",
}

var announcementsServiceResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportAnnouncementsServiceAnnouncementSubscriptionHints},
	},
}
