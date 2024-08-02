---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_run"
sidebar_current: "docs-oci-datasource-devops-build_run"
description: |-
  Provides details about a specific Build Run in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_run
This data source provides details about a specific Build Run resource in Oracle Cloud Infrastructure Devops service.

Returns the details of a build run for a given build run ID.


## Example Usage

```hcl
data "oci_devops_build_run" "test_build_run" {
	#Required
	build_run_id = oci_devops_build_run.test_build_run.id
}
```

## Argument Reference

The following arguments are supported:

* `build_run_id` - (Required) Unique build run identifier.


## Attributes Reference

The following attributes are exported:

* `build_outputs` - Outputs from the build.
	* `artifact_override_parameters` - Specifies the list of artifact override arguments at the time of deployment.
		* `items` - List of artifact override arguments at the time of deployment.
			* `deploy_artifact_id` - The OCID of the artifact to which this parameter applies.
			* `name` - Name of the parameter (case-sensitive).
			* `value` - Value of the parameter.
	* `delivered_artifacts` - Specifies the list of artifacts delivered through the Deliver Artifacts stage.
		* `items` - List of artifacts delivered through the Deliver Artifacts stage.
			* `artifact_repository_id` - The OCID of the artifact registry repository used by the DeliverArtifactStage
			* `artifact_type` - Type of artifact delivered.
			* `delivered_artifact_hash` - The hash of the container registry artifact pushed by the Deliver Artifacts stage.
			* `delivered_artifact_id` - The OCID of the artifact pushed by the Deliver Artifacts stage.
			* `deploy_artifact_id` - The OCID of the deployment artifact definition.
			* `image_uri` - The imageUri of the OCIR artifact pushed by the DeliverArtifactStage
			* `output_artifact_name` - Name of the output artifact defined in the build specification file.
			* `path` - Path of the repository where artifact was pushed
			* `version` - Version of the artifact pushed
	* `exported_variables` - Specifies list of exported variables. 
		* `items` - List of exported variables.
			* `name` - Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$. 
			* `value` - Value of the argument.
	* `vulnerability_audit_summary_collection` - List of vulnerability audit summary.
		* `items` - List of vulnerability audit summary.
			* `build_stage_id` - Build stage OCID where scan was configured.
			* `commit_hash` - Commit hash used while retrieving the pom file for vulnerabilityAudit.
			* `vulnerability_audit_id` - The OCID of the vulnerability audit.
* `build_pipeline_id` - The OCID of the build pipeline.
* `build_run_arguments` - Specifies list of arguments passed along with the build run. 
	* `items` - List of arguments provided at the time of running the build.
		* `name` - Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$. Example: 'Build_Pipeline_param' is not same as 'build_pipeline_Param' 
		* `value` - Value of the argument.
* `build_run_progress` - The run progress details of a build run.
	* `build_pipeline_stage_run_progress` - Map of stage OCIDs to build pipeline stage run progress model.
		* `actual_build_runner_shape` - Name of Build Runner shape where this Build Stage is running.
		* `actual_build_runner_shape_config` - Build Runner Shape configuration.
			* `memory_in_gbs` - The total amount of memory set for the instance in gigabytes.
			* `ocpus` - The total number of OCPUs set for the instance.
		* `artifact_override_parameters` - Specifies the list of artifact override arguments at the time of deployment.
			* `items` - List of artifact override arguments at the time of deployment.
				* `deploy_artifact_id` - The OCID of the artifact to which this parameter applies.
				* `name` - Name of the parameter (case-sensitive).
				* `value` - Value of the parameter.
		* `build_pipeline_stage_id` - The stage OCID.
		* `build_pipeline_stage_predecessors` - The collection containing the predecessors of a stage.
			* `items` - A list of build pipeline stage predecessors for a stage.
				* `id` - The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID. 
		* `build_pipeline_stage_type` - Stage types.
		* `build_source_collection` - Collection of build sources.
			* `items` - Collection of build sources. In case of UPDATE operation, replaces existing build sources list. Merging with existing build sources is not supported.
				* `branch` - Branch name.
				* `connection_id` - Connection identifier pertinent to Bitbucket Server source provider
				* `connection_type` - The type of source provider.
				* `name` - Name of the build source. This must be unique within a build source collection. The name can be used by customers to locate the working directory pertinent to this repository.
				* `repository_id` - The DevOps code repository ID.
				* `repository_url` - URL for the repository.
		* `build_spec_file` - The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
		* `delivered_artifacts` - Specifies the list of artifacts delivered through the Deliver Artifacts stage.
			* `items` - List of artifacts delivered through the Deliver Artifacts stage.
				* `artifact_repository_id` - The OCID of the artifact registry repository used by the DeliverArtifactStage
				* `artifact_type` - Type of artifact delivered.
				* `delivered_artifact_hash` - The hash of the container registry artifact pushed by the Deliver Artifacts stage.
				* `delivered_artifact_id` - The OCID of the artifact pushed by the Deliver Artifacts stage.
				* `deploy_artifact_id` - The OCID of the deployment artifact definition.
				* `image_uri` - The imageUri of the OCIR artifact pushed by the DeliverArtifactStage
				* `output_artifact_name` - Name of the output artifact defined in the build specification file.
				* `path` - Path of the repository where artifact was pushed
				* `version` - Version of the artifact pushed
		* `deployment_id` - Identifier of the deployment triggered.
		* `exported_variables` - Specifies list of exported variables. 
			* `items` - List of exported variables.
				* `name` - Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$. 
				* `value` - Value of the argument.
		* `image` - Image name for the Build Environment
		* `primary_build_source` - Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
		* `private_access_config` - Specifies the configuration needed when the target Oracle Cloud Infrastructure resource, i.e., OKE cluster, resides in customer's private network. 
			* `network_channel_type` - Network channel type.
			* `nsg_ids` - An array of network security group OCIDs.
			* `subnet_id` - The OCID of the subnet where VNIC resources will be created for private endpoint.
		* `stage_display_name` - Build Run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
		* `stage_execution_timeout_in_seconds` - Timeout for the Build Stage Execution. Value in seconds.
		* `status` - The current status of the stage.
		* `steps` - The details about all the steps in a Build stage
			* `name` - Name of the step.
			* `state` - State of the step.
			* `time_finished` - Time when the step finished.
			* `time_started` - Time when the step started.
		* `time_finished` - The time the stage finished executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
		* `time_started` - The time the stage started executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	* `time_finished` - The time the build run finished. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	* `time_started` - The time the build run started. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `build_run_source` - The source from which the build run is triggered.
	* `repository_id` - The DevOps code repository identifier that invoked the build run.
	* `source_type` - The source from which the build run is triggered.
	* `trigger_id` - The trigger that invoked the build run.
	* `trigger_info` - Trigger details that need to be used for the BuildRun
		* `actions` - The list of actions that are to be performed for this Trigger
			* `build_pipeline_id` - The OCID of the build pipeline to be triggered.
			* `build_run_filter` - The filters for the trigger.
				* `events` - The events, for example, PUSH, PULL_REQUEST_CREATED, PULL_REQUEST_UPDATED.
				* `exclude` - Attributes to filter GitLab self-hosted server events. File filter criteria - Changes only affecting excluded files will not invoke a build. if both include and exclude filter are used then exclusion filter will be applied on the result set of inclusion filter.
					* `file_filter` - Attributes to support include/exclude files for triggering build runs.
						* `file_paths` - The file paths/glob pattern for files.
				* `include` - Attributes to filter GitLab self-hosted server events.
					* `base_ref` - The target branch for pull requests; not applicable for push requests.
					* `head_ref` - Branch for push event; source branch for pull requests.
					* `repository_name` - The repository name for trigger events.
				* `trigger_source` - Source of the trigger. Allowed values are, GITHUB and GITLAB.
			* `type` - The type of action that will be taken. Allowed value is TRIGGER_BUILD_PIPELINE.
		* `display_name` - Name for Trigger.
* `commit_info` - Commit details that need to be used for the build run.
	* `commit_hash` - Commit hash pertinent to the repository URL and the specified branch.
	* `repository_branch` - Name of the repository branch.
	* `repository_url` - Repository URL.
* `compartment_id` - The OCID of the compartment where the build is running.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - Build run display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of the DevOps project.
* `state` - The current state of the build run.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the build run was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the build run was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

