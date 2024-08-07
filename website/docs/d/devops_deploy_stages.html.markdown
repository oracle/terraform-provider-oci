---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_stages"
sidebar_current: "docs-oci-datasource-devops-deploy_stages"
description: |-
  Provides the list of Deploy Stages in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_stages
This data source provides the list of Deploy Stages in Oracle Cloud Infrastructure Devops service.

Retrieves a list of deployment stages.

## Example Usage

```hcl
data "oci_devops_deploy_stages" "test_deploy_stages" {

	#Optional
	compartment_id = var.compartment_id
	deploy_pipeline_id = oci_devops_deploy_pipeline.test_deploy_pipeline.id
	display_name = var.deploy_stage_display_name
	id = var.deploy_stage_id
	state = var.deploy_stage_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `deploy_pipeline_id` - (Optional) The ID of the parent pipeline.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `state` - (Optional) A filter to return only deployment stages that matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `deploy_stage_collection` - The list of deploy_stage_collection.

### DeployStage Reference

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
* `helm_command_artifact_ids` - List of Helm command artifact OCIDs.
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
* `timeout_in_seconds` -Time to wait for execution of a Shell/Helm stage. Defaults to 36000 seconds for Shell and 300 seconds for Helm Stage
* `traffic_shift_target` - Specifies the target or destination backend set.
* `values_artifact_ids` - List of values.yaml file artifact OCIDs.
* `wait_criteria` - Specifies wait criteria for the Wait stage.
	* `wait_duration` - The absolute wait duration. An ISO 8601 formatted duration string. Minimum waitDuration should be 5 seconds. Maximum waitDuration can be up to 2 days.
	* `wait_type` - Wait criteria type.
