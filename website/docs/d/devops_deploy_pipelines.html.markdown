---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_pipelines"
sidebar_current: "docs-oci-datasource-devops-deploy_pipelines"
description: |-
  Provides the list of Deploy Pipelines in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_pipelines
This data source provides the list of Deploy Pipelines in Oracle Cloud Infrastructure Devops service.

Returns a list of deployment pipelines.

## Example Usage

```hcl
data "oci_devops_deploy_pipelines" "test_deploy_pipelines" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.deploy_pipeline_display_name
	id = var.deploy_pipeline_id
	project_id = oci_devops_project.test_project.id
	state = var.deploy_pipeline_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only DeployPipelines that matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `deploy_pipeline_collection` - The list of deploy_pipeline_collection.

### DeployPipeline Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment where the pipeline is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_pipeline_artifacts` - List of all artifacts used in the pipeline.
	* `items` - List of all artifacts used in the pipeline.
		* `deploy_artifact_id` - The OCID of an artifact
		* `deploy_pipeline_stages` - List of stages.
			* `items` - List of stages.
				* `deploy_stage_id` - The OCID of a stage
				* `display_name` - Display name of the stage. Avoid entering confidential information.
		* `display_name` - Display name of the artifact. Avoid entering confidential information.
* `deploy_pipeline_environments` - List of all environments used in the pipeline.
	* `items` - List of all environments used in the pipeline.
		* `deploy_environment_id` - The OCID of an Environment
		* `deploy_pipeline_stages` - List of stages.
			* `items` - List of stages.
				* `deploy_stage_id` - The OCID of a stage
				* `display_name` - Display name of the stage. Avoid entering confidential information.
		* `display_name` - Display name of the environment. Avoid entering confidential information.
* `deploy_pipeline_parameters` - Specifies list of parameters present in the deployment pipeline. In case of Update operation, replaces existing parameters list. Merging with existing parameters is not supported.
	* `items` - List of parameters defined for a deployment pipeline.
		* `default_value` - Default value of the parameter.
		* `description` - Description of the parameter.
		* `name` - Name of the parameter (case-sensitive). Parameter name must be ^[a-zA-Z][a-zA-Z_0-9]*$.
* `description` - Optional description about the deployment pipeline.
* `display_name` - Deployment pipeline display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of a project.
* `state` - The current state of the deployment pipeline.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the deployment pipeline was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment pipeline was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

