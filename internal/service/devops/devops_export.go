package devops

import (
	"fmt"

	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDevopsRepositoryRefHints.GetIdFn = getDevopsRepositoryRefId
	exportDevopsRepositorySettingHints.GetIdFn = getDevopsRepositorySettingId
	exportDevopsProjectRepositorySettingHints.GetIdFn = getDevopsProjectRepositorySettingId
	tf_export.RegisterCompartmentGraphs("devops", devopsResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDevopsRepositoryRefId(resource *tf_export.OCIResource) (string, error) {

	refName, ok := resource.SourceAttributes["ref_name"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find refName for Devops RepositoryRef")
	}
	repositoryId := resource.Parent.Id
	return GetRepositoryRefCompositeId(refName, repositoryId), nil
}

func getDevopsRepositorySettingId(resource *tf_export.OCIResource) (string, error) {

	repositoryId := resource.Parent.Id
	return GetRepositorySettingCompositeId(repositoryId), nil
}

func getDevopsProjectRepositorySettingId(resource *tf_export.OCIResource) (string, error) {

	projectId := resource.Parent.Id
	return GetProjectRepositorySettingCompositeId(projectId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDevopsProjectHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_project",
	DatasourceClass:        "oci_devops_projects",
	DatasourceItemsAttr:    "project_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "project",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.ProjectLifecycleStateActive),
	},
}

var exportDevopsDeployEnvironmentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_deploy_environment",
	DatasourceClass:        "oci_devops_deploy_environments",
	DatasourceItemsAttr:    "deploy_environment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deploy_environment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.DeployEnvironmentLifecycleStateActive),
		string(oci_devops.DeployEnvironmentLifecycleStateNeedsAttention),
	},
}

var exportDevopsDeployArtifactHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_deploy_artifact",
	DatasourceClass:        "oci_devops_deploy_artifacts",
	DatasourceItemsAttr:    "deploy_artifact_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deploy_artifact",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.DeployArtifactLifecycleStateActive),
	},
}

var exportDevopsDeployPipelineHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_deploy_pipeline",
	DatasourceClass:        "oci_devops_deploy_pipelines",
	DatasourceItemsAttr:    "deploy_pipeline_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deploy_pipeline",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.DeployPipelineLifecycleStateActive),
	},
}

var exportDevopsDeployStageHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_deploy_stage",
	DatasourceClass:        "oci_devops_deploy_stages",
	DatasourceItemsAttr:    "deploy_stage_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deploy_stage",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.DeployStageLifecycleStateActive),
	},
}

var exportDevopsDeploymentHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_deployment",
	DatasourceClass:        "oci_devops_deployments",
	DatasourceItemsAttr:    "deployment_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "deployment",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.DeploymentLifecycleStateSucceeded),
	},
}

var exportDevopsRepositoryHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_repository",
	DatasourceClass:        "oci_devops_repositories",
	DatasourceItemsAttr:    "repository_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "repository",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.RepositoryLifecycleStateActive),
	},
}

var exportDevopsRepositoryRefHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_repository_ref",
	DatasourceClass:        "oci_devops_repository_refs",
	DatasourceItemsAttr:    "repository_ref_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "repository_ref",
	RequireResourceRefresh: true,
}

var exportDevopsBuildPipelineHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_build_pipeline",
	DatasourceClass:        "oci_devops_build_pipelines",
	DatasourceItemsAttr:    "build_pipeline_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "build_pipeline",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.BuildPipelineLifecycleStateActive),
	},
}

var exportDevopsBuildRunHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_build_run",
	DatasourceClass:        "oci_devops_build_runs",
	DatasourceItemsAttr:    "build_run_summary_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "build_run",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.BuildRunLifecycleStateSucceeded),
	},
}

var exportDevopsConnectionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_connection",
	DatasourceClass:        "oci_devops_connections",
	DatasourceItemsAttr:    "connection_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "connection",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.ConnectionLifecycleStateActive),
	},
}

var exportDevopsBuildPipelineStageHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_build_pipeline_stage",
	DatasourceClass:        "oci_devops_build_pipeline_stages",
	DatasourceItemsAttr:    "build_pipeline_stage_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "build_pipeline_stage",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.BuildPipelineStageLifecycleStateActive),
	},
}

var exportDevopsTriggerHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_devops_trigger",
	DatasourceClass:        "oci_devops_triggers",
	DatasourceItemsAttr:    "trigger_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "trigger",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_devops.TriggerLifecycleStateActive),
	},
}

var exportDevopsRepositoryMirrorHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_devops_repository_mirror",
	ResourceAbbreviation: "repository_mirror",
}

var exportDevopsRepositorySettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_devops_repository_setting",
	DatasourceClass:      "oci_devops_repository_setting",
	ResourceAbbreviation: "repository_setting",
}

var exportDevopsProjectRepositorySettingHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_devops_project_repository_setting",
	DatasourceClass:      "oci_devops_project_repository_setting",
	ResourceAbbreviation: "project_repository_setting",
}

var devopsResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportDevopsProjectHints},
		{TerraformResourceHints: exportDevopsDeployEnvironmentHints},
		{TerraformResourceHints: exportDevopsDeployArtifactHints},
		{TerraformResourceHints: exportDevopsDeployPipelineHints},
		{TerraformResourceHints: exportDevopsDeployStageHints},
		{TerraformResourceHints: exportDevopsDeploymentHints},
		{TerraformResourceHints: exportDevopsRepositoryHints},
		{TerraformResourceHints: exportDevopsBuildPipelineHints},
		{TerraformResourceHints: exportDevopsBuildRunHints},
		{TerraformResourceHints: exportDevopsConnectionHints},
		{TerraformResourceHints: exportDevopsBuildPipelineStageHints},
		{TerraformResourceHints: exportDevopsTriggerHints},
	},
	"oci_devops_project": {
		{
			TerraformResourceHints: exportDevopsProjectRepositorySettingHints,
			DatasourceQueryParams: map[string]string{
				"project_id": "id",
			},
		},
	},
	"oci_devops_repository": {
		{
			TerraformResourceHints: exportDevopsRepositoryRefHints,
			DatasourceQueryParams: map[string]string{
				"repository_id": "id",
			},
		},
		{
			TerraformResourceHints: exportDevopsRepositorySettingHints,
			DatasourceQueryParams: map[string]string{
				"repository_id": "id",
			},
		},
	},
}
