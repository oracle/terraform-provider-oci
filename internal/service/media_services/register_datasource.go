// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_media_services_media_asset", MediaServicesMediaAssetDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_asset_distribution_channel_attachment", MediaServicesMediaAssetDistributionChannelAttachmentDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_assets", MediaServicesMediaAssetsDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow", MediaServicesMediaWorkflowDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_configuration", MediaServicesMediaWorkflowConfigurationDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_configurations", MediaServicesMediaWorkflowConfigurationsDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_job", MediaServicesMediaWorkflowJobDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_job_fact", MediaServicesMediaWorkflowJobFactDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_job_facts", MediaServicesMediaWorkflowJobFactsDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_jobs", MediaServicesMediaWorkflowJobsDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflow_task_declaration", MediaServicesMediaWorkflowTaskDeclarationDataSource())
	tfresource.RegisterDatasource("oci_media_services_media_workflows", MediaServicesMediaWorkflowsDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_cdn_config", MediaServicesStreamCdnConfigDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_cdn_configs", MediaServicesStreamCdnConfigsDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_distribution_channel", MediaServicesStreamDistributionChannelDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_distribution_channels", MediaServicesStreamDistributionChannelsDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_packaging_config", MediaServicesStreamPackagingConfigDataSource())
	tfresource.RegisterDatasource("oci_media_services_stream_packaging_configs", MediaServicesStreamPackagingConfigsDataSource())
	tfresource.RegisterDatasource("oci_media_services_system_media_workflow", MediaServicesSystemMediaWorkflowDataSource())
}
