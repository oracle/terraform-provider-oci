// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_generative_ai_dedicated_ai_cluster", GenerativeAiDedicatedAiClusterDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_dedicated_ai_clusters", GenerativeAiDedicatedAiClustersDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_endpoint", GenerativeAiEndpointDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_endpoints", GenerativeAiEndpointsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_generative_ai_private_endpoint", GenerativeAiGenerativeAiPrivateEndpointDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_generative_ai_private_endpoints", GenerativeAiGenerativeAiPrivateEndpointsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_application", GenerativeAiHostedApplicationDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_application_storage", GenerativeAiHostedApplicationStorageDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_application_storages", GenerativeAiHostedApplicationStoragesDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_applications", GenerativeAiHostedApplicationsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_deployment", GenerativeAiHostedDeploymentDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_hosted_deployments", GenerativeAiHostedDeploymentsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_imported_model", GenerativeAiImportedModelDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_imported_models", GenerativeAiImportedModelsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_model", GenerativeAiModelDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_models", GenerativeAiModelsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_project", GenerativeAiProjectDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_projects", GenerativeAiProjectsDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_semantic_store", GenerativeAiSemanticStoreDataSource())
	tfresource.RegisterDatasource("oci_generative_ai_semantic_stores", GenerativeAiSemanticStoresDataSource())
}
