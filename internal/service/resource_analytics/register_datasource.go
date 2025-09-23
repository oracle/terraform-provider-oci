// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_resource_analytics_monitored_region", ResourceAnalyticsMonitoredRegionDataSource())
	tfresource.RegisterDatasource("oci_resource_analytics_monitored_regions", ResourceAnalyticsMonitoredRegionsDataSource())
	tfresource.RegisterDatasource("oci_resource_analytics_resource_analytics_instance", ResourceAnalyticsResourceAnalyticsInstanceDataSource())
	tfresource.RegisterDatasource("oci_resource_analytics_resource_analytics_instances", ResourceAnalyticsResourceAnalyticsInstancesDataSource())
	tfresource.RegisterDatasource("oci_resource_analytics_tenancy_attachment", ResourceAnalyticsTenancyAttachmentDataSource())
	tfresource.RegisterDatasource("oci_resource_analytics_tenancy_attachments", ResourceAnalyticsTenancyAttachmentsDataSource())
}
