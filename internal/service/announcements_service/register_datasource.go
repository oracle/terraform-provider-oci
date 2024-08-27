// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package announcements_service

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_announcements_service_announcement_subscription", AnnouncementsServiceAnnouncementSubscriptionDataSource())
	tfresource.RegisterDatasource("oci_announcements_service_announcement_subscriptions", AnnouncementsServiceAnnouncementSubscriptionsDataSource())
	tfresource.RegisterDatasource("oci_announcements_service_services", AnnouncementsServiceServicesDataSource())
}
