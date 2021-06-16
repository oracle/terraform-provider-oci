---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_stage"
sidebar_current: "docs-oci-resource-devops-deploy_stage"
description: |-
  Provides the Deploy Stage resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_deploy_stage
This resource provides the Deploy Stage resource in Oracle Cloud Infrastructure Devops service.

Creates a new deployment stage.

## Example Usage

```hcl
resource "oci_devops_deploy_stage" "test_deploy_stage" {
	#Required
	deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
	deploy_stage_predecessor_collection {
		#Required
		items {
			#Required
			id = var.deploy_stage_deploy_stage_predecessor_collection_items_id
		}
	}
	deploy_stage_type = var.deploy_stage_deploy_stage_type

	#Optional
	approval_policy {
		#Required
		approval_policy_type = var.deploy_stage_approval_policy_approval_policy_type
		number_of_approvals_required = var.deploy_stage_approval_policy_number_of_approvals_required
	}
	blue_backend_ips {

		#Optional
		items = var.deploy_stage_blue_backend_ips_items
	}
	compute_instance_group_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	config = var.deploy_stage_config
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	deploy_artifact_ids = var.deploy_stage_deploy_artifact_ids
	deployment_spec_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	description = var.deploy_stage_description
	display_name = var.deploy_stage_display_name
	docker_image_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	failure_policy {
		#Required
		policy_type = var.deploy_stage_failure_policy_policy_type

		#Optional
		failure_count = var.deploy_stage_failure_policy_failure_count
		failure_percentage = var.deploy_stage_failure_policy_failure_percentage
	}
	freeform_tags = {"bar-key"= "value"}
	function_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	function_timeout_in_seconds = var.deploy_stage_function_timeout_in_seconds
	green_backend_ips {

		#Optional
		items = var.deploy_stage_green_backend_ips_items
	}
	is_async = var.deploy_stage_is_async
	is_validation_enabled = var.deploy_stage_is_validation_enabled
	kubernetes_manifest_deploy_artifact_ids = var.deploy_stage_kubernetes_manifest_deploy_artifact_ids
	load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	max_memory_in_mbs = var.deploy_stage_max_memory_in_mbs
	namespace = var.deploy_stage_namespace
	oke_cluster_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	rollback_policy {

		#Optional
		policy_type = var.deploy_stage_rollback_policy_policy_type
	}
	rollout_policy {
		#Required
		policy_type = var.deploy_stage_rollout_policy_policy_type

		#Optional
		batch_count = var.deploy_stage_rollout_policy_batch_count
		batch_delay_in_seconds = var.deploy_stage_rollout_policy_batch_delay_in_seconds
		batch_percentage = var.deploy_stage_rollout_policy_batch_percentage
		ramp_limit_percent = var.deploy_stage_rollout_policy_ramp_limit_percent
	}
	traffic_shift_target = var.deploy_stage_traffic_shift_target
	wait_criteria {
		#Required
		wait_duration = var.deploy_stage_wait_criteria_wait_duration
		wait_type = var.deploy_stage_wait_criteria_wait_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `approval_policy` - (Required when deploy_stage_type=MANUAL_APPROVAL) (Updatable) Specifies the approval policy.
	* `approval_policy_type` - (Required) (Updatable) Approval policy type.
	* `number_of_approvals_required` - (Required) (Updatable) A minimum number of approvals required for stage to proceed.
* `blue_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `compute_instance_group_deploy_environment_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) A compute instance group environment OCID for rolling deployment.
* `config` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - (Applicable when deploy_stage_type=INVOKE_FUNCTION) (Updatable) Optional binary artifact OCID user may provide to this stage.
* `deploy_artifact_ids` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) Additional file artifact OCIDs.
* `deploy_pipeline_id` - (Required) The OCID of a pipeline.
* `deploy_stage_predecessor_collection` - (Required) (Updatable) Collection containing the predecessors of a stage.
	* `items` - (Required) (Updatable) A list of stage predecessors for a stage.
		* `id` - (Required) (Updatable) The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
* `deploy_stage_type` - (Required) (Updatable) Deployment stage type.
* `deployment_spec_deploy_artifact_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) The OCID of the artifact that contains the deployment specification.
* `description` - (Optional) (Updatable) Optional description about the deployment stage.
* `display_name` - (Optional) (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `docker_image_deploy_artifact_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) A Docker image artifact OCID.
* `failure_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	* `failure_count` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT) (Updatable) The threshold count of failed instances in the group, which when reached or exceeded sets the stage as FAILED.
	* `failure_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE) (Updatable) The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	* `policy_type` - (Required) (Updatable) Specifies if the failure instance size is given by absolute number or by percentage.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_deploy_environment_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION | INVOKE_FUNCTION) (Updatable) Function environment OCID.
* `function_timeout_in_seconds` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Timeout for execution of the Function. Value in seconds.
* `green_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `is_async` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether this stage executes asynchronously.
* `is_validation_enabled` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether the invoked function should be validated.
* `kubernetes_manifest_deploy_artifact_ids` - (Required when deploy_stage_type=OKE_DEPLOYMENT) (Updatable) List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
* `load_balancer_config` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies config for load balancer traffic shift stages.
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The OCID of the load balancer.
* `max_memory_in_mbs` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Maximum usable memory for the Function (in MB).
* `namespace` - (Applicable when deploy_stage_type=OKE_DEPLOYMENT) (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_cluster_deploy_environment_id` - (Required when deploy_stage_type=OKE_DEPLOYMENT) (Updatable) Kubernetes cluster environment OCID for deployment.
* `rollback_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT) (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT) (Updatable) Specifies type of the deployment stage rollback policy.
* `rollout_policy` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE) (Updatable) The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - (Required) (Updatable) The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Indicates the criteria to stop.
* `traffic_shift_target` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies the target or destination backend set.
* `wait_criteria` - (Required when deploy_stage_type=WAIT) (Updatable) Specifies wait criteria for the Wait stage.
	* `wait_duration` - (Required) (Updatable) The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - (Required) (Updatable) Wait criteria type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `approval_policy` - Specifies the approval policy.
	* `approval_policy_type` - Approval policy type.
	* `number_of_approvals_required` - A minimum number of approvals required for stage to proceed.
* `blue_backend_ips` - Collection of backend environment IP addresses.
	* `items` - The IP address of the backend server. A server could be a compute instance or a load balancer.
* `compartment_id` - The OCID of a compartment.
* `compute_instance_group_deploy_environment_id` - A compute instance group environment OCID for rolling deployment.
* `config` - User provided key and value pair configuration, which is assigned through constants or parameter.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - Optional binary artifact OCID user may provide to this stage.
* `deploy_artifact_ids` - Additional file artifact OCIDs.
* `deploy_pipeline_id` - The OCID of a pipeline.
* `deploy_stage_predecessor_collection` - Collection containing the predecessors of a stage.
	* `items` - A list of stage predecessors for a stage.
		* `id` - The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
* `deploy_stage_type` - Deployment stage type.
* `deployment_spec_deploy_artifact_id` - The OCID of the artifact that contains the deployment specification.
* `description` - Optional description about the deployment stage.
* `display_name` - Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `docker_image_deploy_artifact_id` - A Docker image artifact OCID.
* `failure_policy` - Specifies a failure policy for a compute instance group rolling deployment stage.
	* `failure_count` - The threshold count of failed instances in the group, which when reached or exceeded sets the stage as FAILED.
	* `failure_percentage` - The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	* `policy_type` - Specifies if the failure instance size is given by absolute number or by percentage.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_deploy_environment_id` - Function environment OCID.
* `function_timeout_in_seconds` - Timeout for execution of the Function. Value in seconds.
* `green_backend_ips` - Collection of backend environment IP addresses.
	* `items` - The IP address of the backend server. A server could be a compute instance or a load balancer.
* `id` - Unique identifier that is immutable on creation.
* `is_async` - A boolean flag specifies whether this stage executes asynchronously.
* `is_validation_enabled` - A boolean flag specifies whether the invoked function must be validated.
* `kubernetes_manifest_deploy_artifact_ids` - List of Kubernetes manifest artifact OCIDs, the manifests should not include any job resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `load_balancer_config` - Specifies config for load balancer traffic shift stages.
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `max_memory_in_mbs` - Maximum usable memory for the Function (in MB).
* `namespace` - Default Namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_cluster_deploy_environment_id` - Kubernetes cluster environment OCID for deployment.
* `project_id` - The OCID of a project.
* `rollback_policy` - Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - Specifies type of the deployment stage rollback policy.
* `rollout_policy` - Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - Indicates the criteria to stop.
* `state` - The current state of the deployment stage.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `traffic_shift_target` - Specifies the target or destination backend set.
* `wait_criteria` - Specifies wait criteria for the Wait stage.
	* `wait_duration` - The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - Wait criteria type.

## Import

DeployStages can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deploy_stage.test_deploy_stage "id"
```

