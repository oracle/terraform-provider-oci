package media_services

import (
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("media_services", mediaServicesResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportMediaServicesStreamPackagingConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_stream_packaging_config",
	DatasourceClass:        "oci_media_services_stream_packaging_configs",
	DatasourceItemsAttr:    "stream_packaging_config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_packaging_config",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.StreamPackagingConfigLifecycleStateActive),
		string(oci_media_services.StreamPackagingConfigLifecycleStateNeedsAttention),
	},
}

var exportMediaServicesMediaWorkflowHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_media_workflow",
	DatasourceClass:        "oci_media_services_media_workflows",
	DatasourceItemsAttr:    "media_workflow_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "media_workflow",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.MediaWorkflowLifecycleStateActive),
		string(oci_media_services.MediaWorkflowLifecycleStateNeedsAttention),
	},
}

var exportMediaServicesStreamDistributionChannelHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_stream_distribution_channel",
	DatasourceClass:        "oci_media_services_stream_distribution_channels",
	DatasourceItemsAttr:    "stream_distribution_channel_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_distribution_channel",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.StreamDistributionChannelLifecycleStateActive),
		string(oci_media_services.StreamDistributionChannelLifecycleStateNeedsAttention),
	},
}

var exportMediaServicesMediaWorkflowJobHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_media_workflow_job",
	DatasourceClass:        "oci_media_services_media_workflow_jobs",
	DatasourceItemsAttr:    "media_workflow_job_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "media_workflow_job",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.MediaWorkflowJobLifecycleStateSucceeded),
	},
}

var exportMediaServicesStreamCdnConfigHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_stream_cdn_config",
	DatasourceClass:        "oci_media_services_stream_cdn_configs",
	DatasourceItemsAttr:    "stream_cdn_config_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "stream_cdn_config",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.StreamCdnConfigLifecycleStateActive),
		string(oci_media_services.StreamCdnConfigLifecycleStateNeedsAttention),
	},
}

var exportMediaServicesMediaAssetHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_media_asset",
	DatasourceClass:        "oci_media_services_media_assets",
	DatasourceItemsAttr:    "media_asset_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "media_asset",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.LifecycleStateActive),
	},
}

var exportMediaServicesMediaWorkflowConfigurationHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_media_services_media_workflow_configuration",
	DatasourceClass:        "oci_media_services_media_workflow_configurations",
	DatasourceItemsAttr:    "media_workflow_configuration_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "media_workflow_configuration",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_media_services.MediaWorkflowConfigurationLifecycleStateActive),
		string(oci_media_services.MediaWorkflowLifecycleStateNeedsAttention),
	},
}

var mediaServicesResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportMediaServicesMediaWorkflowHints},
		{TerraformResourceHints: exportMediaServicesStreamDistributionChannelHints},
		{TerraformResourceHints: exportMediaServicesMediaWorkflowJobHints},
		{TerraformResourceHints: exportMediaServicesMediaAssetHints},
		{TerraformResourceHints: exportMediaServicesMediaWorkflowConfigurationHints},
	},
}
