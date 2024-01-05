// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_document

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_ai_document_model", AiDocumentModelResource())
	tfresource.RegisterResource("oci_ai_document_processor_job", AiDocumentProcessorJobResource())
	tfresource.RegisterResource("oci_ai_document_project", AiDocumentProjectResource())
}
