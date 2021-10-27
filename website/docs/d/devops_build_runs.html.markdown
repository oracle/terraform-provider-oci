---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_runs"
sidebar_current: "docs-oci-datasource-devops-build_runs"
description: |-
  Provides the list of Build Runs in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_runs
This data source provides the list of Build Runs in Oracle Cloud Infrastructure Devops service.

Returns a list of build runs summary.


## Example Usage

```hcl
data "oci_devops_build_runs" "test_build_runs" {

	#Optional
	build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
	compartment_id = var.compartment_id
	display_name = var.build_run_display_name
	id = var.build_run_id
	project_id = oci_devops_project.test_project.id
	state = var.build_run_state
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_id` - (Optional) Unique build pipeline identifier.
* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only Build Runs that matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `build_run_summary_collection` - The list of build_run_summary_collection.

### BuildRun Reference

The following attributes are exported:

* `build_outputs` - Outputs from the Build
	* `artifact_override_parameters` - Specifies the list of artifact override arguments at the time of deployment.
		* `items` - List of artifact override arguments at the time of deployment.
			* `deploy_artifact_id` - The OCID of the artifact to which this parameter applies.
			* `name` - Name of the parameter (case-sensitive).
			* `value` - Value of the parameter.
	* `delivered_artifacts` - Specifies the list of Artifacts delivered via DeliverArtifactStage
		* `items` - List of Artifacts delivered via DeliverArtifactStage
			* `artifact_repository_id` - The OCID of the artifact registry repository used by the DeliverArtifactStage
			* `artifact_type` - Type of Artifact Delivered
			* `delivered_artifact_hash` - The Hash of the OCIR artifact pushed by the DeliverArtifactStage
			* `delivered_artifact_id` - The OCID of the artifact pushed by the DeliverArtifactStage
			* `deploy_artifact_id` - The OCID of the deploy artifact definition
			* `image_uri` - The imageUri of the OCIR artifact pushed by the DeliverArtifactStage
			* `output_artifact_name` - Name of the output artifact defined in the build spec
			* `path` - Path of the repository where artifact was pushed
			* `version` - Version of the artifact pushed
	* `exported_variables` - Specifies list of Exported Variables. 
		* `items` - List of exported variables
			* `name` - Name of the parameter (Case-sensitive). 
			* `value` - value of the argument
* `build_pipeline_id` - Pipeline Identifier
* `build_run_arguments` - Specifies list of arguments passed along with the BuildRun. 
	* `items` - List of arguments provided at the time of BuildRun.
		* `name` - Name of the parameter (Case-sensitive). 
		* `value` - value of the argument
* `build_run_progress` - The run progress details of a BuildRun.
	* `build_pipeline_stage_run_progress` - Map of stage OCIDs to BuildPipelineStageRunProgress model.
		* `actual_build_runner_shape` - Name of Build Runner shape where this Build Stage is running.
		* `actual_build_runner_shape_config` - Build Runner Shape configuration.
			* `memory_in_gbs` - The total amount of memory set for the instance in gigabytes.
			* `ocpus` - The total number of OCPUs set for the instance.
		* `artifact_override_parameters` - Specifies the list of artifact override arguments at the time of deployment.
			* `items` - List of artifact override arguments at the time of deployment.
				* `deploy_artifact_id` - The OCID of the artifact to which this parameter applies.
				* `name` - Name of the parameter (case-sensitive).
				* `value` - Value of the parameter.
		* `build_pipeline_stage_id` - Stage id
		* `build_pipeline_stage_predecessors` - The containing collection for the predecessors of a Stage.
			* `items` - A list of BuildPipelineStagePredecessors for a stage.
				* `id` - The id of the predecessor stage. If a stages is the first stage in the pipeline, then the id is the pipeline's id.
		* `build_pipeline_stage_type` - Stage sub types.
		* `build_source_collection` - Collection of Build Sources.
			* `items` - Collection of Build sources. In case of UPDATE operation, replaces existing Build sources list. Merging with existing Build Sources is not supported.
				* `branch` - branch name
				* `connection_id` - Connection identifier pertinent to GITHUB source provider
				* `connection_type` - The type of Source Provider.
				* `name` - Name of the Build source. This must be unique within a BuildSourceCollection. The name can be used by customers to locate the working directory pertinent to this repository.
				* `repository_id` - The Devops Code Repository Id
				* `repository_url` - Url for repository
		* `build_spec_file` - The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
		* `delivered_artifacts` - Specifies the list of Artifacts delivered via DeliverArtifactStage
			* `items` - List of Artifacts delivered via DeliverArtifactStage
				* `artifact_repository_id` - The OCID of the artifact registry repository used by the DeliverArtifactStage
				* `artifact_type` - Type of Artifact Delivered
				* `delivered_artifact_hash` - The Hash of the OCIR artifact pushed by the DeliverArtifactStage
				* `delivered_artifact_id` - The OCID of the artifact pushed by the DeliverArtifactStage
				* `deploy_artifact_id` - The OCID of the deploy artifact definition
				* `image_uri` - The imageUri of the OCIR artifact pushed by the DeliverArtifactStage
				* `output_artifact_name` - Name of the output artifact defined in the build spec
				* `path` - Path of the repository where artifact was pushed
				* `version` - Version of the artifact pushed
		* `deployment_id` - Identifier of the Deployment Trigerred.
		* `exported_variables` - Specifies list of Exported Variables. 
			* `items` - List of exported variables
				* `name` - Name of the parameter (Case-sensitive). 
				* `value` - value of the argument
		* `image` - Image name for the Build Environment
		* `primary_build_source` - Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
		* `stage_display_name` - BuildRun identifier which can be renamed and is not necessarily unique
		* `stage_execution_timeout_in_seconds` - Timeout for the Build Stage Execution. Value in seconds.
		* `status` - The current status of the Stage.
		* `steps` - The details about all the steps in a Build Stage
			* `name` - Name of the step.
			* `state` - State of the step.
			* `time_finished` - Time when the step finished.
			* `time_started` - Time when the step started.
		* `time_finished` - The time the Stage was finished executing. An RFC3339 formatted datetime string
		* `time_started` - The time the Stage was started executing. An RFC3339 formatted datetime string
	* `time_finished` - The time the BuildRun is finished. An RFC3339 formatted datetime string
	* `time_started` - The time the the BuildRun is started. An RFC3339 formatted datetime string
* `build_run_source` - The source from which this Build Run was triggered
	* `repository_id` - The Devops Code Repository RepoId that invoked this build run
	* `source_type` - Source from which this build run was triggered
	* `trigger_id` - The Trigger that invoked this build run
	* `trigger_info` - Trigger details that need to be used for the BuildRun
		* `actions` - The list of actions that are to be performed for this Trigger
			* `build_pipeline_id` - The id of the build pipeline to be triggered
			* `filter` - The filters for the trigger
				* `events` - The events, example PUSH, PULL_REQUEST_MERGE etc.
				* `include` - Attributes to filter Devops Code Repository events
					* `base_ref` - The target branch for pull requests; not applicable for push
					* `head_ref` - Branch for push event; source branch for pull requests
				* `trigger_source` - Source of the Trigger (allowed values are - GITHUB, GITLAB)
			* `type` - The type of action that will be taken (allowed value - TRIGGER_BUILD_PIPELINE)
		* `display_name` - Name for Trigger.
* `commit_info` - Commit details that need to be used for the BuildRun
	* `commit_hash` - Commit Hash pertinent to the repository URL and Branch specified.
	* `repository_branch` - Name of the repository branch.
	* `repository_url` - Repository URL
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `display_name` - BuildRun identifier which can be renamed and is not necessarily unique
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - Project Identifier
* `state` - The current state of the BuildRun.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the the BuildRun was created. An RFC3339 formatted datetime string
* `time_updated` - The time the BuildRun was updated. An RFC3339 formatted datetime string

