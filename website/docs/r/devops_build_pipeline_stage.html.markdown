---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipeline_stage"
sidebar_current: "docs-oci-resource-devops-build_pipeline_stage"
description: |-
  Provides the Build Pipeline Stage resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_build_pipeline_stage
This resource provides the Build Pipeline Stage resource in Oracle Cloud Infrastructure Devops service.

Creates a new Stage.


## Example Usage

```hcl
resource "oci_devops_build_pipeline_stage" "test_build_pipeline_stage" {
	#Required
	build_pipeline_id = oci_devops_build_pipeline.test_build_pipeline.id
	build_pipeline_stage_predecessor_collection {
		#Required
		items {
			#Required
			id = var.build_pipeline_stage_build_pipeline_stage_predecessor_collection_items_id
		}
	}
	build_pipeline_stage_type = var.build_pipeline_stage_build_pipeline_stage_type

	#Optional
	build_source_collection {

		#Optional
		items {
			#Required
			connection_type = var.build_pipeline_stage_build_source_collection_items_connection_type

			#Optional
			branch = var.build_pipeline_stage_build_source_collection_items_branch
			connection_id = oci_devops_connection.test_connection.id
			name = var.build_pipeline_stage_build_source_collection_items_name
			repository_id = oci_artifacts_repository.test_repository.id
			repository_url = var.build_pipeline_stage_build_source_collection_items_repository_url
		}
	}
	build_spec_file = var.build_pipeline_stage_build_spec_file
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deliver_artifact_collection {

		#Optional
		items {

			#Optional
			artifact_id = oci_devops_artifact.test_artifact.id
			artifact_name = var.build_pipeline_stage_deliver_artifact_collection_items_artifact_name
		}
	}
	deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
	description = var.build_pipeline_stage_description
	display_name = var.build_pipeline_stage_display_name
	freeform_tags = {"bar-key"= "value"}
	image = var.build_pipeline_stage_image
	is_pass_all_parameters_enabled = var.build_pipeline_stage_is_pass_all_parameters_enabled
	primary_build_source = var.build_pipeline_stage_primary_build_source
	stage_execution_timeout_in_seconds = var.build_pipeline_stage_stage_execution_timeout_in_seconds
	wait_criteria {
		#Required
		wait_duration = var.build_pipeline_stage_wait_criteria_wait_duration
		wait_type = var.build_pipeline_stage_wait_criteria_wait_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_id` - (Required) buildPipeline Identifier
* `build_pipeline_stage_predecessor_collection` - (Required) (Updatable) The containing collection for the predecessors of a Stage.
	* `items` - (Required) (Updatable) A list of BuildPipelineStagePredecessors for a stage.
		* `id` - (Required) (Updatable) The id of the predecessor stage. If a stages is the first stage in the pipeline, then the id is the pipeline's id.
* `build_pipeline_stage_type` - (Required) (Updatable) List of stage types. It includes 'Wait stage', 'Build Stage', 'Deliver Artifact Stage' and 'Trigger Deployment Stage'. 
* `build_source_collection` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Collection of Build Sources.
	* `items` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Collection of Build sources. In case of UPDATE operation, replaces existing Build sources list. Merging with existing Build Sources is not supported.
		* `branch` - (Required when build_pipeline_stage_type=BUILD) (Updatable) branch name
		* `connection_id` - (Required when connection_type=GITHUB | GITLAB) (Updatable) Connection identifier pertinent to GITHUB source provider
		* `connection_type` - (Required) (Updatable) The type of Source Provider.
		* `name` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Name of the Build source. This must be unique within a BuildSourceCollection. The name can be used by customers to locate the working directory pertinent to this repository.
		* `repository_id` - (Required when connection_type=DEVOPS_CODE_REPOSITORY) (Updatable) The Devops Code Repository Id
		* `repository_url` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Url for repository
* `build_spec_file` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) The path to the build specification file for this Environment. The default location if not specified is build_spec.yaml
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deliver_artifact_collection` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Specifies an array of Artifacts that need to be pushed to the artifactory stores.
	* `items` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Collection of Artifacts that were generated in the Build Stage and need to be pushed to the artifactory stores. In case of UPDATE operation, replaces existing artifacts list. Merging with existing artifacts is not supported.
		* `artifact_id` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Artifact Identifier which contains the Artifact Definition.
		* `artifact_name` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Name of the artifact specified in the build_spec.yml file.
* `deploy_pipeline_id` - (Required when build_pipeline_stage_type=TRIGGER_DEPLOYMENT_PIPELINE) (Updatable) A target Pipeline ocid that will be run in this stage.
* `description` - (Optional) (Updatable) Optional description about the Stage
* `display_name` - (Optional) (Updatable) Stage identifier which can be renamed and is not necessarily unique
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `image` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Image name for the Build Environment
* `is_pass_all_parameters_enabled` - (Required when build_pipeline_stage_type=TRIGGER_DEPLOYMENT_PIPELINE) (Updatable) A boolean flag specifies whether the parameters should be passed during the deployment trigger.
* `primary_build_source` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) Name of the BuildSource in which the build_spec.yml file need to be located. If not specified, the 1st entry in the BuildSource collection will be chosen as Primary.
* `stage_execution_timeout_in_seconds` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) Timeout for the Build Stage Execution. Value in seconds.
* `wait_criteria` - (Required when build_pipeline_stage_type=WAIT) (Updatable) Specifies wait criteria for wait stage.
	* `wait_duration` - (Required) (Updatable) The absolute wait duration. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days. 
	* `wait_type` - (Required) (Updatable) wait criteria sub type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Build Pipeline Stage
	* `update` - (Defaults to 20 minutes), when updating the Build Pipeline Stage
	* `delete` - (Defaults to 20 minutes), when destroying the Build Pipeline Stage


## Import

BuildPipelineStages can be imported using the `id`, e.g.

```
$ terraform import oci_devops_build_pipeline_stage.test_build_pipeline_stage "id"
```

