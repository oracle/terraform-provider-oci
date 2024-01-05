// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_ai_language_endpoint", AiLanguageEndpointResource())
	tfresource.RegisterResource("oci_ai_language_model", AiLanguageModelResource())
	tfresource.RegisterResource("oci_ai_language_project", AiLanguageProjectResource())
}
