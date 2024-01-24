// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_language_endpoint", AiLanguageEndpointDataSource())
	tfresource.RegisterDatasource("oci_ai_language_endpoints", AiLanguageEndpointsDataSource())
	tfresource.RegisterDatasource("oci_ai_language_model", AiLanguageModelDataSource())
	tfresource.RegisterDatasource("oci_ai_language_model_evaluation_results", AiLanguageModelEvaluationResultsDataSource())
	tfresource.RegisterDatasource("oci_ai_language_model_type", AiLanguageModelTypeDataSource())
	tfresource.RegisterDatasource("oci_ai_language_models", AiLanguageModelsDataSource())
	tfresource.RegisterDatasource("oci_ai_language_project", AiLanguageProjectDataSource())
	tfresource.RegisterDatasource("oci_ai_language_projects", AiLanguageProjectsDataSource())
}
