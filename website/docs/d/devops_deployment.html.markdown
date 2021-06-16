---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deployment"
sidebar_current: "docs-oci-datasource-devops-deployment"
description: |-
  Provides details about a specific Deployment in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deployment
This data source provides details about a specific Deployment resource in Oracle Cloud Infrastructure Devops service.

Retrieves a deployment by identifier.

## Example Usage

```hcl
data "oci_devops_deployment" "test_deployment" {
	#Required
	deployment_id = oci_devops_deployment.test_deployment.id
}
```

## Argument Reference

The following arguments are supported:

* `deployment_id` - (Required) Unique deployment identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of a compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_override_arguments` - Specifies the list of artifact override arguments at the time of deployment.
	* `items` - List of artifact override arguments at the time of deployment.
		* `deploy_artifact_id` - The OCID of the artifact to which this parameter applies.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - Value of the parameter.
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
* `deploy_pipeline_id` - The OCID of a pipeline.
* `deploy_stage_id` - Specifies the OCID of the stage to be deployed.
* `deployment_arguments` - Specifies list of arguments passed along with the deployment.
	* `items` - List of arguments provided at the time of deployment.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - value of the argument.
* `deployment_execution_progress` - The execution progress details of a deployment.
	* `deploy_stage_execution_progress` - Map of stage OCIDs to deploy stage execution progress model.
		* `approval_actions` - 
			* `action` - The action of the user on the DevOps deployment stage.
			* `subject_id` - The subject ID of the user who approves or disapproves a DevOps deployment stage.
		* `deploy_stage_display_name` - Stage display name. Avoid entering confidential information.
		* `deploy_stage_execution_progress_details` - Details about stage execution for all the target environments.
			* `rollback_steps` - Details about all the rollback steps for one target environment.
				* `name` - Name of the step.
				* `state` - State of the step.
				* `time_finished` - Time when the step finished.
				* `time_started` - Time when the step started.
			* `steps` - Details about all the steps for one target environment.
				* `name` - Name of the step.
				* `state` - State of the step.
				* `time_finished` - Time when the step finished.
				* `time_started` - Time when the step started.
			* `target_group` - Group for the target environment for example, the batch number for an Instance Group deployment.
			* `target_id` - The function ID, instance ID or the cluster ID. For Wait stage it will be the stage ID.
		* `deploy_stage_id` - The OCID of the stage.
		* `deploy_stage_predecessors` - Collection containing the predecessors of a stage.
			* `items` - A list of stage predecessors for a stage.
				* `id` - The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
		* `deploy_stage_type` - Deployment stage type.
		* `status` - The current state of the stage.
		* `time_finished` - Time the stage finished executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
		* `time_started` - Time the stage started executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	* `time_finished` - Time the deployment is finished. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
	* `time_started` - Time the deployment is started. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `deployment_type` - Specifies type of Deployment
* `display_name` - Deployment identifier which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `previous_deployment_id` - Specifies the OCID of the previous deployment to be redeployed.
* `project_id` - The OCID of a project.
* `state` - The current state of the deployment.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the deployment was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

