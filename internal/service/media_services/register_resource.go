// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_media_services_media_asset", MediaServicesMediaAssetResource())
	tfresource.RegisterResource("oci_media_services_media_workflow", MediaServicesMediaWorkflowResource())
	tfresource.RegisterResource("oci_media_services_media_workflow_configuration", MediaServicesMediaWorkflowConfigurationResource())
	tfresource.RegisterResource("oci_media_services_media_workflow_job", MediaServicesMediaWorkflowJobResource())
	tfresource.RegisterResource("oci_media_services_stream_cdn_config", MediaServicesStreamCdnConfigResource())
	tfresource.RegisterResource("oci_media_services_stream_distribution_channel", MediaServicesStreamDistributionChannelResource())
	tfresource.RegisterResource("oci_media_services_stream_packaging_config", MediaServicesStreamPackagingConfigResource())
}
