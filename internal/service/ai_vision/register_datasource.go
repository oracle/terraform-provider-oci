// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_vision_model", AiVisionModelDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_models", AiVisionModelsDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_project", AiVisionProjectDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_projects", AiVisionProjectsDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_group", AiVisionStreamGroupDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_groups", AiVisionStreamGroupsDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_job", AiVisionStreamJobDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_jobs", AiVisionStreamJobsDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_source", AiVisionStreamSourceDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_stream_sources", AiVisionStreamSourcesDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_vision_private_endpoint", AiVisionVisionPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_vision_private_endpoints", AiVisionVisionPrivateEndpointsDataSource())
}
