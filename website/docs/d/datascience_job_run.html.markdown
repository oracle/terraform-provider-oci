---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_job_run"
sidebar_current: "docs-oci-datasource-datascience-job_run"
description: |-
  Provides details about a specific Job Run in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_job_run
This data source provides details about a specific Job Run resource in Oracle Cloud Infrastructure Data Science service.

Gets a job run.

## Example Usage

```hcl
data "oci_datascience_job_run" "test_job_run" {
	#Required
	job_run_id = oci_datascience_job_run.test_job_run.id
}
```

## Argument Reference

The following arguments are supported:

* `job_run_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job run.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the job run.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.
* `job_configuration_override_details` - The job configuration details 
	* `command_line_arguments` - The arguments to pass to the job. 
	* `environment_variables` - Environment variables to set for the job. 
	* `job_type` - The type of job.
	* `maximum_runtime_in_minutes` - A time bound for the execution of the job. Timer starts when the job becomes active. 
	* `startup_probe_details` - The probe indicates whether the application within the job run has started.
		* `command` - The commands to run in the target job run to perform the startup probe
		* `failure_threshold` - How many times the job will try before giving up when a probe fails.
		* `initial_delay_in_seconds` - Number of seconds after the job run has started before a startup probe is initiated.
		* `job_probe_check_type` - The probe check type to perform the startup probe and specifies the type of health check for a job.
		* `period_in_seconds` - Number of seconds how often the job run should perform a startup probe
* `job_environment_configuration_override_details` - Environment configuration to capture job runtime dependencies.
	* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
	* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
	* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
	* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
	* `image_signature_id` - OCID of the container image signature
	* `job_environment_type` - The environment configuration type used for job runtime.
* `job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
* `job_infrastructure_configuration_details` - The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - The infrastructure type used for job run.
	* `job_shape_config_details` - Details for the job run shape configuration. Specify only when a flex shape is selected.
		* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory available to the job run instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the job run instance. 
	* `shape_name` - The name that corresponds to the JobShapeSummary to use for the job node
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_infrastructure_configuration_override_details` - The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - The infrastructure type used for job run.
	* `job_shape_config_details` - Details for the job run shape configuration. Specify only when a flex shape is selected.
		* `memory_in_gbs` - The total amount of memory available to the job run instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the job run instance. 
	* `shape_name` - The name that corresponds to the JobShapeSummary to use for the job node
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_override_details` - Logging configuration for resource. 
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - If customer logging is enabled for job runs.
	* `log_group_id` - The log group id for where log objects are for job runs. 
	* `log_id` - The log id the job run will push logs too. 
* `job_node_configuration_override_details` - The job node configuration details
	* `job_network_configuration` - The job network configuration details 
		* `job_network_type` - job network type
		* `subnet_id` - The custom subnet id
	* `job_node_group_configuration_details_list` - List of JobNodeGroupConfigurationDetails
		* `job_configuration_details` - The job configuration details 
			* `command_line_arguments` - The arguments to pass to the job. 
			* `environment_variables` - Environment variables to set for the job. 
			* `job_type` - The type of job.
			* `maximum_runtime_in_minutes` - A time bound for the execution of the job. Timer starts when the job becomes active. 
			* `startup_probe_details` - The probe indicates whether the application within the job run has started.
				* `command` - The commands to run in the target job run to perform the startup probe
				* `failure_threshold` - How many times the job will try before giving up when a probe fails.
				* `initial_delay_in_seconds` - Number of seconds after the job run has started before a startup probe is initiated.
				* `job_probe_check_type` - The probe check type to perform the startup probe and specifies the type of health check for a job.
				* `period_in_seconds` - Number of seconds how often the job run should perform a startup probe
		* `job_environment_configuration_details` - Environment configuration to capture job runtime dependencies.
			* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
			* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
			* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
			* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
			* `image_signature_id` - OCID of the container image signature
			* `job_environment_type` - The environment configuration type used for job runtime.
		* `job_infrastructure_configuration_details` - The job infrastructure configuration details (shape, block storage, etc.) 
			* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
			* `job_infrastructure_type` - The infrastructure type used for job run.
			* `job_shape_config_details` - Details for the job run shape configuration. Specify only when a flex shape is selected.
				* `memory_in_gbs` - The total amount of memory available to the job run instance, in gigabytes. 
				* `ocpus` - The total number of OCPUs available to the job run instance. 
			* `shape_name` - The name that corresponds to the JobShapeSummary to use for the job node
			* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
		* `minimum_success_replicas` - The minimum threshold of successful replicas for node group to be successful. All replicas need to succeed if this is not specified.
		* `name` - node group name.
		* `replicas` - The number of nodes.
	* `job_node_type` - The node type used for job run.
	* `maximum_runtime_in_minutes` - A time bound for the execution of the job run. Timer starts when the job run is in progress. 
	* `startup_order` - The execution order of node groups
* `job_storage_mount_configuration_details_list` - Collection of JobStorageMountConfigurationDetails.
	* `bucket` - The object storage bucket
	* `destination_directory_name` - The local directory name to be mounted
	* `destination_path` - The local path of the mounted directory, excluding directory name.
	* `export_id` - OCID of the export
	* `mount_target_id` - OCID of the mount target
	* `namespace` - The object storage namespace
	* `prefix` - Prefix in the bucket to mount
	* `storage_type` - The type of storage.
* `job_storage_mount_configuration_override_details_list` - Collection of JobStorageMountConfigurationDetails.
	* `bucket` - The object storage bucket
	* `destination_directory_name` - The local directory name to be mounted
	* `destination_path` - The local path of the mounted directory, excluding directory name.
	* `export_id` - OCID of the export
	* `mount_target_id` - OCID of the mount target
	* `namespace` - The object storage namespace
	* `prefix` - Prefix in the bucket to mount
	* `storage_type` - The type of storage.
* `lifecycle_details` - Details of the state of the job run.
* `log_details` - Customer logging details for job run. 
	* `log_group_id` - The log group id for where log objects will be for job runs. 
	* `log_id` - The log id of the log object the job run logs will be shipped to. 
* `node_group_details_list` - Collection of NodeGroupDetails
	* `lifecycle_details` - The state details of the node group.
	* `name` - node group name.
	* `state` - The state of the node group.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job run with.
* `state` - The state of the job run.
* `time_accepted` - The date and time the job run was accepted in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_finished` - The date and time the job run request was finished in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_started` - The date and time the job run request was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

