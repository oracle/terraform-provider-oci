// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_announcements_service_announcement_subscription", AnnouncementsServiceAnnouncementSubscriptionResource())
	tfresource.RegisterResource("oci_announcements_service_announcement_subscriptions_actions_change_compartment", AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResource())
	tfresource.RegisterResource("oci_announcements_service_announcement_subscriptions_filter_group", AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResource())
}
