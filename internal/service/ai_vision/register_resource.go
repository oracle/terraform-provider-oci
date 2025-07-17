// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_vision

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_ai_vision_model", AiVisionModelResource())
	tfresource.RegisterResource("oci_ai_vision_project", AiVisionProjectResource())
	tfresource.RegisterResource("oci_ai_vision_stream_group", AiVisionStreamGroupResource())
	tfresource.RegisterResource("oci_ai_vision_stream_job", AiVisionStreamJobResource())
	tfresource.RegisterResource("oci_ai_vision_stream_source", AiVisionStreamSourceResource())
	tfresource.RegisterResource("oci_ai_vision_vision_private_endpoint", AiVisionVisionPrivateEndpointResource())
}
