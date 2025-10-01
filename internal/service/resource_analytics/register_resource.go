// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_resource_analytics_monitored_region", ResourceAnalyticsMonitoredRegionResource())
	tfresource.RegisterResource("oci_resource_analytics_resource_analytics_instance", ResourceAnalyticsResourceAnalyticsInstanceResource())
	tfresource.RegisterResource("oci_resource_analytics_resource_analytics_instance_oac_management", ResourceAnalyticsResourceAnalyticsInstanceOacManagementResource())
	tfresource.RegisterResource("oci_resource_analytics_tenancy_attachment", ResourceAnalyticsTenancyAttachmentResource())
}
