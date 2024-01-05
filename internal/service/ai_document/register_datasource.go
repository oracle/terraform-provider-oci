// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_ai_document_model", AiDocumentModelDataSource())
	tfresource.RegisterDatasource("oci_ai_document_models", AiDocumentModelsDataSource())
	tfresource.RegisterDatasource("oci_ai_document_processor_job", AiDocumentProcessorJobDataSource())
	tfresource.RegisterDatasource("oci_ai_document_project", AiDocumentProjectDataSource())
	tfresource.RegisterDatasource("oci_ai_document_projects", AiDocumentProjectsDataSource())
}
