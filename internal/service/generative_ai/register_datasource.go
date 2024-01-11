// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_generative_ai_dedicated_ai_cluster", GenerativeAiDedicatedAiClusterDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_dedicated_ai_clusters", GenerativeAiDedicatedAiClustersDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_endpoint", GenerativeAiEndpointDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_endpoints", GenerativeAiEndpointsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_model", GenerativeAiModelDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_models", GenerativeAiModelsDataSource())
}
