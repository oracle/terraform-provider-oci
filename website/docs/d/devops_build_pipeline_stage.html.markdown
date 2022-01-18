---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_build_pipeline_stage"
sidebar_current: "docs-oci-datasource-devops-build_pipeline_stage"
description: |-
  Provides details about a specific Build Pipeline Stage in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_build_pipeline_stage
This data source provides details about a specific Build Pipeline Stage resource in Oracle Cloud Infrastructure Devops service.

Retrieves a stage based on the stage ID provided in the request.

## Example Usage

```hcl
data "oci_devops_build_pipeline_stage" "test_build_pipeline_stage" {
	#Required
	build_pipeline_stage_id = oci_devops_build_pipeline_stage.test_build_pipeline_stage.id
}
```

## Argument Reference

The following arguments are supported:

* `build_pipeline_stage_id` - (Required) Unique stage identifier.


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

