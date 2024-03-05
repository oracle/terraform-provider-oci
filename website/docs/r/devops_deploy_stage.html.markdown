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
	are_hooks_enabled = var.deploy_stage_are_hooks_enabled
	blue_backend_ips {

		#Optional
		items = var.deploy_stage_blue_backend_ips_items
	}
	blue_green_strategy {
		#Required
		ingress_name = var.deploy_stage_blue_green_strategy_ingress_name
		namespace_a = var.deploy_stage_blue_green_strategy_namespace_a
		namespace_b = var.deploy_stage_blue_green_strategy_namespace_b
		strategy_type = var.deploy_stage_blue_green_strategy_strategy_type
	}
	canary_strategy {
		#Required
		ingress_name = var.deploy_stage_canary_strategy_ingress_name
		namespace = var.deploy_stage_canary_strategy_namespace
		strategy_type = var.deploy_stage_canary_strategy_strategy_type
	}
	command_spec_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	compute_instance_group_blue_green_deployment_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_canary_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_canary_traffic_shift_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	compute_instance_group_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	config = var.deploy_stage_config
	container_config {
		#Required
		container_config_type = var.deploy_stage_container_config_container_config_type
		network_channel {
			#Required
			network_channel_type = var.deploy_stage_container_config_network_channel_network_channel_type
			subnet_id = oci_core_subnet.test_subnet.id

			#Optional
			nsg_ids = var.deploy_stage_container_config_network_channel_nsg_ids
		}
		shape_config {
			#Required
			ocpus = var.deploy_stage_container_config_shape_config_ocpus

			#Optional
			memory_in_gbs = var.deploy_stage_container_config_shape_config_memory_in_gbs
		}
		shape_name = oci_core_shape.test_shape.name

		#Optional
		availability_domain = var.deploy_stage_container_config_availability_domain
		compartment_id = var.compartment_id
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	deploy_artifact_ids = var.deploy_stage_deploy_artifact_ids
	deploy_environment_id_a = var.deploy_stage_deploy_environment_id_a
	deploy_environment_id_b = var.deploy_stage_deploy_environment_id_b
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
	helm_chart_deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
	helm_command_artifact_ids = var.deploy_stage_helm_command_artifact_ids
	is_async = var.deploy_stage_is_async
	is_debug_enabled = var.deploy_stage_is_debug_enabled
	is_force_enabled = var.deploy_stage_is_force_enabled
	is_uninstall_on_stage_delete = var.deploy_stage_is_uninstall_on_stage_delete
	is_validation_enabled = var.deploy_stage_is_validation_enabled
	kubernetes_manifest_deploy_artifact_ids = var.deploy_stage_kubernetes_manifest_deploy_artifact_ids
	load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	max_history = var.deploy_stage_max_history
	max_memory_in_mbs = var.deploy_stage_max_memory_in_mbs
	namespace = var.deploy_stage_namespace
	oke_blue_green_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_canary_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_canary_traffic_shift_deploy_stage_id = oci_devops_deploy_stage.test_deploy_stage.id
	oke_cluster_deploy_environment_id = oci_devops_deploy_environment.test_deploy_environment.id
	production_load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_production_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	purpose = var.deploy_stage_purpose
	release_name = var.deploy_stage_release_name
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
	set_string {

		#Optional
		items {

			#Optional
			name = var.deploy_stage_set_string_items_name
			value = var.deploy_stage_set_string_items_value
		}
	}
	set_values {

		#Optional
		items {

			#Optional
			name = var.deploy_stage_set_values_items_name
			value = var.deploy_stage_set_values_items_value
		}
	}
	should_cleanup_on_fail = var.deploy_stage_should_cleanup_on_fail
	should_not_wait = var.deploy_stage_should_not_wait
	should_reset_values = var.deploy_stage_should_reset_values
	should_reuse_values = var.deploy_stage_should_reuse_values
	should_skip_crds = var.deploy_stage_should_skip_crds
	should_skip_render_subchart_notes = var.deploy_stage_should_skip_render_subchart_notes
	test_load_balancer_config {

		#Optional
		backend_port = var.deploy_stage_test_load_balancer_config_backend_port
		listener_name = oci_load_balancer_listener.test_listener.name
		load_balancer_id = oci_load_balancer_load_balancer.test_load_balancer.id
	}
	timeout_in_seconds = var.deploy_stage_timeout_in_seconds
	traffic_shift_target = var.deploy_stage_traffic_shift_target
	values_artifact_ids = var.deploy_stage_values_artifact_ids
	wait_criteria {
		#Required
		wait_duration = var.deploy_stage_wait_criteria_wait_duration
		wait_type = var.deploy_stage_wait_criteria_wait_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `approval_policy` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL | MANUAL_APPROVAL | OKE_CANARY_APPROVAL) (Updatable) Specifies the approval policy.
	* `approval_policy_type` - (Required) (Updatable) Approval policy type.
	* `number_of_approvals_required` - (Required) (Updatable) A minimum number of approvals required for stage to proceed.
* `are_hooks_enabled` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Disable pre/post upgrade hooks. Set to false by default.
* `blue_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `blue_green_strategy` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT) Specifies the required blue green release strategy for OKE deployment.
	* `ingress_name` - (Required) Name of the Ingress resource.
	* `namespace_a` - (Required) First Namespace for deployment.
	* `namespace_b` - (Required) Second Namespace for deployment.
	* `strategy_type` - (Required) Blue Green strategy type
* `canary_strategy` - (Required when deploy_stage_type=OKE_CANARY_DEPLOYMENT) Specifies the required canary release strategy for OKE deployment.
	* `ingress_name` - (Required) Name of the Ingress resource.
	* `namespace` - (Required) Canary namespace to be used for Kubernetes canary deployment.
	* `strategy_type` - (Required) Canary strategy type.
* `command_spec_deploy_artifact_id` - (Required when deploy_stage_type=SHELL) (Updatable) The OCID of the artifact that contains the command specification.
* `compute_instance_group_blue_green_deployment_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT) The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
* `compute_instance_group_canary_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT) A compute instance group canary stage OCID for load balancer.
* `compute_instance_group_canary_traffic_shift_deploy_stage_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL) (Updatable) A compute instance group canary traffic shift stage OCID for load balancer.
* `compute_instance_group_deploy_environment_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) A compute instance group environment OCID for rolling deployment.
* `config` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) User provided key and value pair configuration, which is assigned through constants or parameter.
* `container_config` - (Required when deploy_stage_type=SHELL) (Updatable) Specifies the container configuration.
	* `availability_domain` - (Optional) (Updatable) Availability domain where the ContainerInstance will be created.
	* `compartment_id` - (Optional) (Updatable) The OCID of the compartment where the ContainerInstance will be created.
	* `container_config_type` - (Required) (Updatable) Container configuration type.
	* `network_channel` - (Required) (Updatable) Specifies the configuration needed when the target Oracle Cloud Infrastructure resource, i.e., OKE cluster, resides in customer's private network. 
		* `network_channel_type` - (Required) (Updatable) Network channel type.
		* `nsg_ids` - (Optional) (Updatable) An array of network security group OCIDs.
		* `subnet_id` - (Required) (Updatable) The OCID of the subnet where VNIC resources will be created for private endpoint.
	* `shape_config` - (Required) (Updatable) Determines the size and amount of resources available to the instance.
		* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to the instance, in gigabytes.
		* `ocpus` - (Required) (Updatable) The total number of OCPUs available to the instance.
	* `shape_name` - (Required) (Updatable) The shape of the ContainerInstance. The shape determines the resources available to the ContainerInstance.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - (Applicable when deploy_stage_type=INVOKE_FUNCTION) (Updatable) Optional artifact OCID. The artifact will be included in the body for the function invocation during the stage's execution. If the DeployArtifact.argumentSubstituitionMode is set to SUBSTITUTE_PLACEHOLDERS, then the pipeline parameter values will be used to replace the placeholders in the artifact content. 
* `deploy_artifact_ids` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) The list of file artifact OCIDs to deploy.
* `deploy_environment_id_a` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT) First compute instance group environment OCID for deployment.
* `deploy_environment_id_b` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT) Second compute instance group environment OCID for deployment.
* `deploy_pipeline_id` - (Required) The OCID of a pipeline.
* `deploy_stage_predecessor_collection` - (Required) (Updatable) Collection containing the predecessors of a stage.
	* `items` - (Required) (Updatable) A list of stage predecessors for a stage.
		* `id` - (Required) (Updatable) The OCID of the predecessor stage. If a stage is the first stage in the pipeline, then the ID is the pipeline's OCID.
* `deploy_stage_type` - (Required) (Updatable) Deployment stage type.
* `deployment_spec_deploy_artifact_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) The OCID of the artifact that contains the deployment specification.
* `description` - (Optional) (Updatable) Optional description about the deployment stage.
* `display_name` - (Optional) (Updatable) Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `docker_image_deploy_artifact_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) A Docker image artifact OCID.
* `failure_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT) (Updatable) Specifies a failure policy for a compute instance group rolling deployment stage.
	* `failure_count` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT) (Updatable) The threshold count of failed instances in the group, which when reached or exceeded sets the stage as FAILED.
	* `failure_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE) (Updatable) The failure percentage threshold, which when reached or exceeded sets the stage as FAILED. Percentage is computed as the ceiling value of the number of failed instances over the total count of the instances in the group.
	* `policy_type` - (Required) (Updatable) Specifies if the failure instance size is given by absolute number or by percentage.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `function_deploy_environment_id` - (Required when deploy_stage_type=DEPLOY_FUNCTION | INVOKE_FUNCTION) (Updatable) Function environment OCID.
* `function_timeout_in_seconds` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Timeout for execution of the Function. Value in seconds.
* `green_backend_ips` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Collection of backend environment IP addresses.
	* `items` - (Applicable when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The IP address of the backend server. A server could be a compute instance or a load balancer.
* `helm_chart_deploy_artifact_id` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Helm chart artifact OCID.
* `helm_command_artifact_ids` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) List of Helm command artifact OCIDs.
* `is_async` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether this stage executes asynchronously.
* `is_debug_enabled` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Enables helm --debug option to stream output to tf stdout. Set to false by default.
* `is_force_enabled` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Force resource update through delete; or if required, recreate. Set to false by default.
* `is_uninstall_on_stage_delete` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Uninstall the Helm chart release on deleting the stage.
* `is_validation_enabled` - (Required when deploy_stage_type=INVOKE_FUNCTION) (Updatable) A boolean flag specifies whether the invoked function should be validated.
* `kubernetes_manifest_deploy_artifact_ids` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT | OKE_CANARY_DEPLOYMENT | OKE_DEPLOYMENT) (Updatable) List of Kubernetes manifest artifact OCIDs.
* `load_balancer_config` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) The OCID of the load balancer.
* `max_history` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Limit the maximum number of revisions saved per release. Use 0 for no limit. Set to 10 by default
* `max_memory_in_mbs` - (Applicable when deploy_stage_type=DEPLOY_FUNCTION) (Updatable) Maximum usable memory for the Function (in MB).
* `namespace` - (Applicable when deploy_stage_type=OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Default namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_blue_green_deploy_stage_id` - (Required when deploy_stage_type=OKE_BLUE_GREEN_TRAFFIC_SHIFT) The OCID of the upstream OKE blue-green deployment stage in this pipeline.
* `oke_canary_deploy_stage_id` - (Required when deploy_stage_type=OKE_CANARY_TRAFFIC_SHIFT) The OCID of an upstream OKE canary deployment stage in this pipeline.
* `oke_canary_traffic_shift_deploy_stage_id` - (Required when deploy_stage_type=OKE_CANARY_APPROVAL) The OCID of an upstream OKE canary deployment traffic shift stage in this pipeline.
* `oke_cluster_deploy_environment_id` - (Required when deploy_stage_type=OKE_BLUE_GREEN_DEPLOYMENT | OKE_CANARY_DEPLOYMENT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Kubernetes cluster environment OCID for deployment.
* `production_load_balancer_config` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Specifies configuration for load balancer traffic shift stages. The load balancer specified here should be an Application load balancer type. Network load balancers are not supported.
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) The OCID of the load balancer.
* `purpose` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) The purpose of running this Helm stage
* `release_name` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Default name of the chart instance. Must be unique within a Kubernetes namespace.
* `rollback_policy` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_DEPLOYMENT | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies type of the deployment stage rollback policy.
* `rollout_policy` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - (Required when policy_type=COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE) (Updatable) The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - (Required) (Updatable) The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT | LOAD_BALANCER_TRAFFIC_SHIFT | OKE_CANARY_TRAFFIC_SHIFT) (Updatable) Indicates the criteria to stop.
* `set_string` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies the name and value pairs to set helm values.
	* `items` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) List of parameters defined to set helm value.
		* `name` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Name of the parameter (case-sensitive).
		* `value` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Value of the parameter.
* `set_values` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Specifies the name and value pairs to set helm values.
	* `items` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) List of parameters defined to set helm value.
		* `name` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Name of the parameter (case-sensitive).
		* `value` - (Required when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Value of the parameter.
* `should_cleanup_on_fail` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Allow deletion of new resources created during when an upgrade fails. Set to false by default.
* `should_not_wait` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) Does not wait until all the resources are in a ready state to mark the release as successful if set to true. Set to false by default.
* `should_reset_values` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) During upgrade, reset the values to the ones built into the chart. It overrides shouldReuseValues. Set to false by default.
* `should_reuse_values` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) During upgrade, reuse the values of the last release and merge overrides from the command line. Set to false by default.
* `should_skip_crds` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) If set, no CRDs are installed. By default, CRDs are installed only if they are not present already. Set to false by default.
* `should_skip_render_subchart_notes` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) If set, renders subchart notes along with the parent. Set to false by default.
* `test_load_balancer_config` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Specifies configuration for load balancer traffic shift stages. The load balancer specified here should be an Application load balancer type. Network load balancers are not supported. 
	* `backend_port` - (Applicable when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Listen port for the backend server.
	* `listener_name` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) Name of the load balancer listener.
	* `load_balancer_id` - (Required when deploy_stage_type=COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT | COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT) (Updatable) The OCID of the load balancer.
* `timeout_in_seconds` - (Applicable when deploy_stage_type=SHELL | OKE_HELM_CHART_DEPLOYMENT) (Updatable) Time to wait for execution of a Shell/Helm stage. Defaults to 36000 seconds for Shell and 300 seconds for Helm Stage
* `traffic_shift_target` - (Required when deploy_stage_type=LOAD_BALANCER_TRAFFIC_SHIFT) (Updatable) Specifies the target or destination backend set.
* `values_artifact_ids` - (Applicable when deploy_stage_type=OKE_HELM_CHART_DEPLOYMENT) (Updatable) List of values.yaml file artifact OCIDs.
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
* `are_hooks_enabled` - Disable pre/post upgrade hooks. Set to false by default.
* `blue_backend_ips` - Collection of backend environment IP addresses.
	* `items` - The IP address of the backend server. A server could be a compute instance or a load balancer.
* `blue_green_strategy` - Specifies the required blue green release strategy for OKE deployment.
	* `ingress_name` - Name of the Ingress resource.
	* `namespace_a` - First Namespace for deployment.
	* `namespace_b` - Second Namespace for deployment.
	* `strategy_type` - Blue Green strategy type
* `canary_strategy` - Specifies the required canary release strategy for OKE deployment.
	* `ingress_name` - Name of the Ingress resource.
	* `namespace` - Canary namespace to be used for Kubernetes canary deployment.
	* `strategy_type` - Canary strategy type.
* `command_spec_deploy_artifact_id` - The OCID of the artifact that contains the command specification.
* `compartment_id` - The OCID of a compartment.
* `compute_instance_group_blue_green_deployment_deploy_stage_id` - The OCID of the upstream compute instance group blue-green deployment stage in this pipeline.
* `compute_instance_group_canary_deploy_stage_id` - The OCID of an upstream compute instance group canary deployment stage ID in this pipeline.
* `compute_instance_group_canary_traffic_shift_deploy_stage_id` - A compute instance group canary traffic shift stage OCID for load balancer.
* `compute_instance_group_deploy_environment_id` - A compute instance group environment OCID for rolling deployment.
* `config` - User provided key and value pair configuration, which is assigned through constants or parameter.
* `container_config` - Specifies the container configuration.
	* `availability_domain` - Availability domain where the ContainerInstance will be created.
	* `compartment_id` - The OCID of the compartment where the ContainerInstance will be created.
	* `container_config_type` - Container configuration type.
	* `network_channel` - Specifies the configuration needed when the target Oracle Cloud Infrastructure resource, i.e., OKE cluster, resides in customer's private network. 
		* `network_channel_type` - Network channel type.
		* `nsg_ids` - An array of network security group OCIDs.
		* `subnet_id` - The OCID of the subnet where VNIC resources will be created for private endpoint.
	* `shape_config` - Determines the size and amount of resources available to the instance.
		* `memory_in_gbs` - The total amount of memory available to the instance, in gigabytes.
		* `ocpus` - The total number of OCPUs available to the instance.
	* `shape_name` - The shape of the ContainerInstance. The shape determines the resources available to the ContainerInstance.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_id` - Optional artifact OCID. The artifact will be included in the body for the function invocation during the stage's execution. If the DeployArtifact.argumentSubstituitionMode is set to SUBSTITUTE_PLACEHOLDERS, then the pipeline parameter values will be used to replace the placeholders in the artifact content. 
* `deploy_artifact_ids` - The list of file artifact OCIDs to deploy.
* `deploy_environment_id_a` - First compute instance group environment OCID for deployment.
* `deploy_environment_id_b` - Second compute instance group environment OCID for deployment.
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
* `helm_chart_deploy_artifact_id` - Helm chart artifact OCID.
* `helm_command_artifact_ids` - List of Helm command artifact OCIDs.
* `id` - Unique identifier that is immutable on creation.
* `is_async` - A boolean flag specifies whether this stage executes asynchronously.
* `is_debug_enabled` - Enables helm --debug option to stream output to tf stdout. Set to false by default.
* `is_force_enabled` - Force resource update through delete; or if required, recreate. Set to false by default.
* `is_uninstall_on_stage_delete` - Uninstall the Helm chart release on deleting the stage.
* `is_validation_enabled` - A boolean flag specifies whether the invoked function must be validated.
* `kubernetes_manifest_deploy_artifact_ids` - List of Kubernetes manifest artifact OCIDs.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `max_history` - Limit the maximum number of revisions saved per release. Use 0 for no limit. Set to 10 by default
* `max_memory_in_mbs` - Maximum usable memory for the Function (in MB).
* `namespace` - Default Namespace to be used for Kubernetes deployment when not specified in the manifest.
* `oke_blue_green_deploy_stage_id` - The OCID of the upstream OKE blue-green deployment stage in this pipeline.
* `oke_canary_deploy_stage_id` - The OCID of an upstream OKE canary deployment stage in this pipeline.
* `oke_canary_traffic_shift_deploy_stage_id` - The OCID of an upstream OKE canary deployment traffic shift stage in this pipeline.
* `oke_cluster_deploy_environment_id` - Kubernetes cluster environment OCID for deployment.
* `production_load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `project_id` - The OCID of a project.
* `purpose` - The purpose of running this Helm stage
* `release_name` - Release name of the Helm chart.
* `rollback_policy` - Specifies the rollback policy. This is initiated on the failure of certain stage types.
	* `policy_type` - Specifies type of the deployment stage rollback policy.
* `rollout_policy` - Description of rollout policy for load balancer traffic shift stage.
	* `batch_count` - The number that will be used to determine how many instances will be deployed concurrently.
	* `batch_delay_in_seconds` - The duration of delay between batch rollout. The default delay is 1 minute.
	* `batch_percentage` - The percentage that will be used to determine how many instances will be deployed concurrently.
	* `policy_type` - The type of policy used for rolling out a deployment stage.
	* `ramp_limit_percent` - Indicates the criteria to stop.
* `set_string` - Specifies the name and value pairs to set helm values.
	* `items` - List of parameters defined to set helm value.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - Value of the parameter.
* `set_values` - Specifies the name and value pairs to set helm values.
	* `items` - List of parameters defined to set helm value.
		* `name` - Name of the parameter (case-sensitive).
		* `value` - Value of the parameter.
* `should_cleanup_on_fail` - Allow deletion of new resources created during when an upgrade fails. Set to false by default.
* `should_not_wait` - Waits until all the resources are in a ready state to mark the release as successful. Set to false by default.
* `should_reset_values` - During upgrade, reset the values to the ones built into the chart. It overrides shouldReuseValues. Set to false by default.
* `should_reuse_values` - During upgrade, reuse the values of the last release and merge overrides from the command line. Set to false by default.
* `should_skip_crds` - If set, no CRDs are installed. By default, CRDs are installed only if they are not present already. Set to false by default.
* `should_skip_render_subchart_notes` - If set, renders subchart notes along with the parent. Set to false by default.
* `state` - The current state of the deployment stage.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `test_load_balancer_config` - Specifies config for load balancer traffic shift stages. The Load Balancer specified here should be an Application Load Balancer type. Network Load Balancers are not supported. 
	* `backend_port` - Listen port for the backend server.
	* `listener_name` - Name of the load balancer listener.
	* `load_balancer_id` - The OCID of the load balancer.
* `time_created` - Time the deployment stage was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment stage was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `timeout_in_seconds` - Time to wait for execution of a Shell/Helm stage. Defaults to 36000 seconds for Shell and 300 seconds for Helm Stage.
* `traffic_shift_target` - Specifies the target or destination backend set.
* `values_artifact_ids` - List of values.yaml file artifact OCIDs.
* `wait_criteria` - Specifies wait criteria for the Wait stage.
	* `wait_duration` - The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - Wait criteria type.
	
## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Deploy Stage
	* `update` - (Defaults to 20 minutes), when updating the Deploy Stage
	* `delete` - (Defaults to 20 minutes), when destroying the Deploy Stage


## Import

DeployStages can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deploy_stage.test_deploy_stage "id"
```