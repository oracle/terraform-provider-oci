// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_artifacts_container_configuration", ArtifactsContainerConfigurationResource())
	tfresource.RegisterResource("oci_artifacts_container_image_signature", ArtifactsContainerImageSignatureResource())
	tfresource.RegisterResource("oci_artifacts_container_repository", ArtifactsContainerRepositoryResource())
	tfresource.RegisterResource("oci_artifacts_generic_artifact", ArtifactsGenericArtifactResource())
	tfresource.RegisterResource("oci_artifacts_repository", ArtifactsRepositoryResource())
}
