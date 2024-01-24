// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_artifacts_container_configuration", ArtifactsContainerConfigurationDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_image", ArtifactsContainerImageDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_image_signature", ArtifactsContainerImageSignatureDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_image_signatures", ArtifactsContainerImageSignaturesDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_images", ArtifactsContainerImagesDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_repositories", ArtifactsContainerRepositoriesDataSource())
	tfresource.RegisterDatasource("oci_artifacts_container_repository", ArtifactsContainerRepositoryDataSource())
	tfresource.RegisterDatasource("oci_artifacts_generic_artifact", ArtifactsGenericArtifactDataSource())
	tfresource.RegisterDatasource("oci_artifacts_generic_artifacts", ArtifactsGenericArtifactsDataSource())
	tfresource.RegisterDatasource("oci_artifacts_repositories", ArtifactsRepositoriesDataSource())
	tfresource.RegisterDatasource("oci_artifacts_repository", ArtifactsRepositoryDataSource())
}
