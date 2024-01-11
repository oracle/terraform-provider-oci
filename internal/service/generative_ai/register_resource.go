// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_generative_ai_dedicated_ai_cluster", GenerativeAiDedicatedAiClusterResource())
	tfresource.RegisterResource("oci_generative_ai_endpoint", GenerativeAiEndpointResource())
	tfresource.RegisterResource("oci_generative_ai_model", GenerativeAiModelResource())
}
