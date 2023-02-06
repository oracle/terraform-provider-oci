---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deployment"
sidebar_current: "docs-oci-resource-devops-deployment"
description: |-
  Provides the Deployment resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_deployment
This resource provides the Deployment resource in Oracle Cloud Infrastructure Devops service.

Creates a new deployment.

## Example Usage

```hcl
resource "oci_devops_deployment" "test_deployment" {
	#Required
	deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
	deployment_type = var.deployment_deployment_type

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deploy_artifact_override_arguments {

		#Optional
		items {

			#Optional
			deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
			name = var.deployment_deploy_artifact_override_arguments_items_name
			value = var.deployment_deploy_artifact_override_arguments_items_value
		}
	}
	deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	deploy_stage_override_arguments {

		#Optional
		items {

			#Optional
			deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
			name = var.deployment_deploy_stage_override_arguments_items_name
			value = var.deployment_deploy_stage_override_arguments_items_value
		}
	}
	deployment_arguments {

		#Optional
		items {

			#Optional
			name = var.deployment_deployment_arguments_items_name
			value = var.deployment_deployment_arguments_items_value
		}
	}
	display_name = var.deployment_display_name
	freeform_tags = {"bar-key"= "value"}
	previous_deployment_id = oci_devops_deployment.test_deployment.id
	trigger_new_devops_deployment = var.deployment_trigger_new_devops_deployment_bool
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_override_arguments` - (Applicable when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Specifies the list of artifact override arguments at the time of deployment.
	* `items` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) List of artifact override arguments at the time of deployment.
		* `deploy_artifact_id` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) The OCID of the artifact to which this parameter applies.
		* `name` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Name of the parameter (case-sensitive).
		* `value` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Value of the parameter.
* `deploy_pipeline_id` - (Required) The OCID of a pipeline.
* `deploy_stage_id` - (Required when deployment_type=SINGLE_STAGE_DEPLOYMENT | SINGLE_STAGE_REDEPLOYMENT) Specifies the OCID of the stage to be redeployed.
* `deploy_stage_override_arguments` - (Applicable when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Specifies the list of arguments to be overriden per Stage at the time of deployment.
	* `items` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) List of stage override arguments at the time of deployment.
		* `deploy_stage_id` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) The OCID of the stage.
		* `name` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Name of the parameter (case-sensitive).
		* `value` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Value of the parameter.
* `deployment_arguments` - (Applicable when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Specifies list of arguments passed along with the deployment.
	* `items` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) List of arguments provided at the time of deployment.
        * `name` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) Name of the parameter (case-sensitive).
        * `value` - (Required when deployment_type=PIPELINE_DEPLOYMENT | SINGLE_STAGE_DEPLOYMENT) value of the argument.
        *  To retrieve Helm Diff for Helm stages in the pipeline add deployment_arguments with name=PLAN_DRY_RUN and value=true
* `deployment_type` - (Required) (Updatable) Specifies type for this deployment.
* `display_name` - (Optional) (Updatable) Deployment display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `previous_deployment_id` - (Required when deployment_type=PIPELINE_REDEPLOYMENT | SINGLE_STAGE_REDEPLOYMENT) Specifies the OCID of the previous deployment to be redeployed.
* `trigger_new_devops_deployment` - (Optional) A boolean specifying if a new deployment should be created on every apply. As long as this value is set to true in the config, every apply will trigger a new deployment to be created. The existing deployment resource will be replaced with the new one in the state file (deployment resources are never deleted, they persist as a store of records, but your state file will only track the latest one created with this resource block). 

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values. Additionally, while `trigger_new_devops_deployment`
is set true each apply will force the destruction and recreation of the resource with the new property values.

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
* `deploy_stage_id` - Specifies the OCID of the stage to be redeployed.
* `deploy_stage_override_arguments` - Specifies the list of arguments to be overriden per Stage at the time of deployment.
	* `items` - List of stage override arguments at the time of deployment.
		* `deploy_stage_id` - The OCID of the stage.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - Value of the parameter.
* `deployment_arguments` - Specifies list of arguments passed along with the deployment.
	* `items` - List of arguments provided at the time of deployment.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - value of the argument.
* `deployment_execution_progress` - The execution progress details of a deployment.
	* `deploy_stage_execution_progress` - Map of stage OCIDs to deploy stage execution progress model.
        * `approval_actions` - 
            * `action` - The action of the user on the DevOps deployment stage.
            * `reason` - The reason for approving or rejecting the deployment.
            * `subject_id` - The subject ID of the user who approves or disapproves a DevOps deployment stage.
        * `chart_url` - The URL of an OCIR repository. 
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
        * `environment_id` - The OCID of the environment where the artifacts were deployed.
        * `helm_diff` - Helm Diff output Example: greeting: Version 1.0 + greeting: Version 1.1
          - Helm Diff output shows a diff explaining what a helm upgrade would change during plan, if deployment_arguments PLAN_DRY_RUN with value true is added
          - Helm Diff output is only supported for PIPELINE_DEPLOYMENT and SINGLE_STAGE_DEPLOYMENT deployment type
          - This diff is rendered only for OKE_HELM_CHART_DEPLOYMENT Stage
        * `namespace` - Default namespace to be used for Kubernetes deployment when not specified in the manifest.
        * `release_name` - Release name of the Helm chart.
        * `status` - The current state of the stage.
        * `time_finished` - Time the stage finished executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
        * `time_started` - Time the stage started executing. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
        * `version` - The version of the helm chart stored in OCIR repository.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Deployment
	* `update` - (Defaults to 20 minutes), when updating the Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Deployment


## Import

Deployments can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deployment.test_deployment "id"
```

