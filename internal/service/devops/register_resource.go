// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_devops_build_pipeline", DevopsBuildPipelineResource())
	tfresource.RegisterResource("oci_devops_build_pipeline_stage", DevopsBuildPipelineStageResource())
	tfresource.RegisterResource("oci_devops_build_run", DevopsBuildRunResource())
	tfresource.RegisterResource("oci_devops_connection", DevopsConnectionResource())
	tfresource.RegisterResource("oci_devops_deploy_artifact", DevopsDeployArtifactResource())
	tfresource.RegisterResource("oci_devops_deploy_environment", DevopsDeployEnvironmentResource())
	tfresource.RegisterResource("oci_devops_deploy_pipeline", DevopsDeployPipelineResource())
	tfresource.RegisterResource("oci_devops_deploy_stage", DevopsDeployStageResource())
	tfresource.RegisterResource("oci_devops_deployment", DevopsDeploymentResource())
	tfresource.RegisterResource("oci_devops_project", DevopsProjectResource())
	tfresource.RegisterResource("oci_devops_repository", DevopsRepositoryResource())
	tfresource.RegisterResource("oci_devops_repository_mirror", DevopsRepositoryMirrorResource())
	tfresource.RegisterResource("oci_devops_repository_ref", DevopsRepositoryRefResource())
	tfresource.RegisterResource("oci_devops_trigger", DevopsTriggerResource())
}
