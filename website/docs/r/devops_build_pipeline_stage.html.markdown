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

Creates a new stage.


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

* `build_pipeline_id` - (Required) The OCID of the build pipeline.
* `build_pipeline_stage_predecessor_collection` - (Required) (Updatable) The collection containing the predecessors of a stage.
	* `items` - (Required) (Updatable) A list of build pipeline stage predecessors for a stage.
		* `id` - (Required) (Updatable) The ID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's ID.
* `build_pipeline_stage_type` - (Required) (Updatable) Defines the stage type, which is one of the following: Build, Deliver Artifacts, Wait, and Trigger Deployment. 
* `build_source_collection` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Collection of build sources.
	* `items` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Collection of build sources. In case of UPDATE operation, replaces existing build sources list. Merging with existing build sources is not supported.
		* `branch` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Branch name.
		* `connection_id` - (Required when connection_type=GITHUB | GITLAB) (Updatable) Connection identifier pertinent to GitHub source provider.
		* `connection_type` - (Required) (Updatable) The type of source provider.
		* `name` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Name of the build source. This must be unique within a build source collection. The name can be used by customers to locate the working directory pertinent to this repository.
		* `repository_id` - (Required when connection_type=DEVOPS_CODE_REPOSITORY) (Updatable) The DevOps code repository ID.
		* `repository_url` - (Required when build_pipeline_stage_type=BUILD) (Updatable) URL for the repository.
* `build_spec_file` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deliver_artifact_collection` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Specifies an array of artifacts that need to be pushed to the artifactory stores.
	* `items` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Collection of artifacts that were generated in the Build stage and need to be pushed to the artifactory stores. In case of UPDATE operation, replaces existing artifacts list. Merging with existing artifacts is not supported.
		* `artifact_id` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Artifact identifier that contains the artifact definition.
		* `artifact_name` - (Required when build_pipeline_stage_type=DELIVER_ARTIFACT) (Updatable) Name of the artifact specified in the build_spec.yaml file.
* `deploy_pipeline_id` - (Required when build_pipeline_stage_type=TRIGGER_DEPLOYMENT_PIPELINE) (Updatable) A target deployment pipeline OCID that will run in this stage.
* `description` - (Optional) (Updatable) Optional description about the stage.
* `display_name` - (Optional) (Updatable) Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `image` - (Required when build_pipeline_stage_type=BUILD) (Updatable) Image name for the build environment
* `is_pass_all_parameters_enabled` - (Required when build_pipeline_stage_type=TRIGGER_DEPLOYMENT_PIPELINE) (Updatable) A boolean flag that specifies whether all the parameters must be passed when the deployment is triggered.
* `primary_build_source` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) Name of the build source where the build_spec.yml file is located. If not specified, the first entry in the build source collection is chosen as primary build source.
* `stage_execution_timeout_in_seconds` - (Applicable when build_pipeline_stage_type=BUILD) (Updatable) Timeout for the build stage execution. Specify value in seconds.
* `wait_criteria` - (Required when build_pipeline_stage_type=WAIT) (Updatable) Specifies wait criteria for the Wait stage.
	* `wait_duration` - (Required) (Updatable) The absolute wait duration. Minimum wait duration must be 5 seconds. Maximum wait duration can be up to 2 days. 
	* `wait_type` - (Required) (Updatable) Wait criteria type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `build_pipeline_id` - The OCID of the build pipeline.
* `build_pipeline_stage_predecessor_collection` - The collection containing the predecessors of a stage.
	* `items` - A list of build pipeline stage predecessors for a stage.
		* `id` - The ID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's ID.
* `build_pipeline_stage_type` - Defines the stage type, which is one of the following: Build, Deliver Artifacts, Wait, and Trigger Deployment. 
* `build_source_collection` - Collection of build sources.
	* `items` - Collection of build sources. In case of UPDATE operation, replaces existing build sources list. Merging with existing build sources is not supported.
		* `branch` - Branch name.
		* `connection_id` - Connection identifier pertinent to GitHub source provider.
		* `connection_type` - The type of source provider.
		* `name` - Name of the build source. This must be unique within a build source collection. The name can be used by customers to locate the working directory pertinent to this repository.
		* `repository_id` - The DevOps code repository ID.
		* `repository_url` - URL for the repository.
* `build_spec_file` - The path to the build specification file for this environment. The default location of the file if not specified is build_spec.yaml.
* `compartment_id` - The OCID of the compartment where the pipeline is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deliver_artifact_collection` - Specifies an array of artifacts that need to be pushed to the artifactory stores.
	* `items` - Collection of artifacts that were generated in the Build stage and need to be pushed to the artifactory stores. In case of UPDATE operation, replaces existing artifacts list. Merging with existing artifacts is not supported.
		* `artifact_id` - Artifact identifier that contains the artifact definition.
		* `artifact_name` - Name of the artifact specified in the build_spec.yaml file.
* `deploy_pipeline_id` - A target deployment pipeline OCID that will run in this stage.
* `description` - Optional description about the build stage.
* `display_name` - Stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `image` - Image name for the build environment.
* `is_pass_all_parameters_enabled` - A boolean flag that specifies whether all the parameters must be passed when the deployment is triggered.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `primary_build_source` - Name of the build source where the build_spec.yml file is located. If not specified, then the first entry in the build source collection is chosen as primary build source.
* `project_id` - The OCID of the DevOps project.
* `stage_execution_timeout_in_seconds` - Timeout for the build stage execution. Specify value in seconds.
* `state` - The current state of the stage. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - The time the stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
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

