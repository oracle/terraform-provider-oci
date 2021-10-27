---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipeline_stages"
sidebar_current: "docs-oci-datasource-devops-build_pipeline_stages"
description: |-
  Provides the list of Build Pipeline Stages in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_pipeline_stages
This data source provides the list of Build Pipeline Stages in Oracle Cloud Infrastructure Devops service.

Returns summary of list of all Stages in a compartment or buildPipeline


## Example Usage

```hcl
data "oci_devops_build_pipeline_stages" "test_build_pipeline_stages" {

	#Optional
	build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
	compartment_id = var.compartment_id
	display_name = var.build_pipeline_stage_display_name
	id = var.build_pipeline_stage_id
	state = var.build_pipeline_stage_state
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_id` - (Optional) The ID of the parent build pipeline.
* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `state` - (Optional) A filter to return the stages that match with the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `build_pipeline_stage_collection` - The list of build_pipeline_stage_collection.

### BuildPipelineStage Reference

The following attributes are exported:

* `build_pipeline_id` - Build Pipeline Identifier
* `build_pipeline_stage_predecessor_collection` - The containing collection for the predecessors of a Stage.
	* `items` - A list of BuildPipelineStagePredecessors for a stage.
		* `id` - The id of the predecessor stage. If a stages is the first stage in the pipeline, then the id is the pipeline's id.
* `build_pipeline_stage_type` - List of stage types. It includes 'Wait stage', 'Build Stage', 'Deliver Artifact Stage' and 'Trigger Deployment Stage'. 
* `build_source_collection` - Collection of Build Sources.
	* `items` - Collection of Build sources. In case of UPDATE operation, replaces existing Build sources list. Merging with existing Build Sources is not supported.
		* `branch` - branch name
		* `connection_id` - Connection identifier pertinent to GITHUB source provider
		* `connection_type` - The type of Source Provider.
		* `name` - Name of the Build source. This must be unique within a BuildSourceCollection. The name can be used by customers to locate the working directory pertinent to this repository.
		* `repository_id` - The Devops Code Repository Id
		* `repository_url` - Url for repository
* `build_spec_file` - The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
* `compartment_id` - Compartment Identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deliver_artifact_collection` - Specifies an array of Artifacts that need to be pushed to the artifactory stores.
	* `items` - Collection of Artifacts that were generated in the Build Stage and need to be pushed to the artifactory stores. In case of UPDATE operation, replaces existing artifacts list. Merging with existing artifacts is not supported.
		* `artifact_id` - Artifact Identifier which contains the Artifact Definition.
		* `artifact_name` - Name of the artifact specified in the build_spec.yml file.
* `deploy_pipeline_id` - A target Pipeline ocid that will be run in this stage.
* `description` - Optional description about the BuildStage
* `display_name` - Stage identifier which can be renamed and is not necessarily unique
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation
* `image` - Image name for the Build Environment
* `is_pass_all_parameters_enabled` - A boolean flag specifies whether the parameters should be passed during the deployment trigger.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `primary_build_source` - Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
* `project_id` - Project Identifier
* `stage_execution_timeout_in_seconds` - Timeout for the Build Stage Execution. Value in seconds.
* `state` - The current state of the Stage. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time at which the Stage was created. An RFC3339 formatted datetime string
* `time_updated` - The time at which the Stage was updated. An RFC3339 formatted datetime string
* `wait_criteria` - Specifies wait criteria for the Wait stage.
	* `wait_duration` - The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - Wait criteria type.

