// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_vision_model", AiVisionModelDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_models", AiVisionModelsDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_project", AiVisionProjectDataSource())
	tfresource.RegisterDatasource("oci_ai_vision_projects", AiVisionProjectsDataSource())
}
