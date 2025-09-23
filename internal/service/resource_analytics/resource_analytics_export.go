package resource_analytics

import (
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("resource_analytics", resourceAnalyticsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportResourceAnalyticsResourceAnalyticsInstanceHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_resource_analytics_resource_analytics_instance",
	DatasourceClass:        "oci_resource_analytics_resource_analytics_instances",
	DatasourceItemsAttr:    "resource_analytics_instance_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "resource_analytics_instance",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateActive),
		string(oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateNeedsAttention),
	},
}

var exportResourceAnalyticsTenancyAttachmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_resource_analytics_tenancy_attachment",
	DatasourceClass:        "oci_resource_analytics_tenancy_attachments",
	DatasourceItemsAttr:    "tenancy_attachment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "tenancy_attachment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateActive),
		string(oci_resource_analytics.TenancyAttachmentLifecycleStateNeedsAttention),
	},
}

var exportResourceAnalyticsMonitoredRegionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_resource_analytics_monitored_region",
	DatasourceClass:        "oci_resource_analytics_monitored_regions",
	DatasourceItemsAttr:    "monitored_region_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "monitored_region",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_resource_analytics.MonitoredRegionLifecycleStateActive),
	},
}

var resourceAnalyticsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportResourceAnalyticsResourceAnalyticsInstanceHints},
		{TerraformResourceHints: exportResourceAnalyticsTenancyAttachmentHints},
		{TerraformResourceHints: exportResourceAnalyticsMonitoredRegionHints},
	},
}
