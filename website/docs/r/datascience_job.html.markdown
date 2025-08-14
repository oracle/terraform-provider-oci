---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_job"
sidebar_current: "docs-oci-resource-datascience-job"
description: |-
  Provides the Job resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_job
This resource provides the Job resource in Oracle Cloud Infrastructure Data Science service.

Creates a job.

## Example Usage

```hcl
resource "oci_datascience_job" "test_job" {
	#Required
	compartment_id = var.compartment_id
	project_id = oci_datascience_project.test_project.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.job_description
	display_name = var.job_display_name
	freeform_tags = {"Department"= "Finance"}
	job_configuration_details {
		#Required
		job_type = var.job_job_configuration_details_job_type

		#Optional
		command_line_arguments = var.job_job_configuration_details_command_line_arguments
		environment_variables = var.job_job_configuration_details_environment_variables
		maximum_runtime_in_minutes = var.job_job_configuration_details_maximum_runtime_in_minutes
		startup_probe_details {
			#Required
			command = var.job_job_configuration_details_startup_probe_details_command
			job_probe_check_type = var.job_job_configuration_details_startup_probe_details_job_probe_check_type

			#Optional
			cpu_baseline = var.job_job_infrastructure_configuration_details_job_shape_config_details_cpu_baseline
			failure_threshold = var.job_job_configuration_details_startup_probe_details_failure_threshold
			initial_delay_in_seconds = var.job_job_configuration_details_startup_probe_details_initial_delay_in_seconds
			memory_in_gbs = var.job_job_infrastructure_configuration_details_job_shape_config_details_memory_in_gbs
			ocpus = var.job_job_infrastructure_configuration_details_job_shape_config_details_ocpus
			period_in_seconds = var.job_job_configuration_details_startup_probe_details_period_in_seconds
		}
	}
	job_environment_configuration_details {
		#Required
		image = var.job_job_environment_configuration_details_image
		job_environment_type = var.job_job_environment_configuration_details_job_environment_type

		#Optional
		cmd = var.job_job_environment_configuration_details_cmd
		entrypoint = var.job_job_environment_configuration_details_entrypoint
		image_digest = var.job_job_environment_configuration_details_image_digest
		image_signature_id = oci_datascience_image_signature.test_image_signature.id
	}
	job_infrastructure_configuration_details {
		#Required
		job_infrastructure_type = var.job_job_infrastructure_configuration_details_job_infrastructure_type

		#Optional
		block_storage_size_in_gbs = var.job_job_infrastructure_configuration_details_block_storage_size_in_gbs
		job_shape_config_details {

			#Optional
			memory_in_gbs = var.job_job_infrastructure_configuration_details_job_shape_config_details_memory_in_gbs
			ocpus = var.job_job_infrastructure_configuration_details_job_shape_config_details_ocpus
		}
		shape_name = oci_core_shape.test_shape.name
		subnet_id = oci_core_subnet.test_subnet.id
	}
	job_log_configuration_details {

		#Optional
		enable_auto_log_creation = var.job_job_log_configuration_details_enable_auto_log_creation
		enable_logging = var.job_job_log_configuration_details_enable_logging
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
	job_node_configuration_details {
		#Required
		job_node_type = var.job_job_node_configuration_details_job_node_type

		#Optional
		job_network_configuration {
			#Required
			job_network_type = var.job_job_node_configuration_details_job_network_configuration_job_network_type

			#Optional
			subnet_id = oci_core_subnet.test_subnet.id
		}
		job_node_group_configuration_details_list {
			#Required
			name = var.job_job_node_configuration_details_job_node_group_configuration_details_list_name

			#Optional
			job_configuration_details {
				#Required
				job_type = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_job_type

				#Optional
				command_line_arguments = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_command_line_arguments
				environment_variables = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_environment_variables
				maximum_runtime_in_minutes = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_maximum_runtime_in_minutes
				startup_probe_details {
					#Required
					command = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_startup_probe_details_command
					job_probe_check_type = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_startup_probe_details_job_probe_check_type

					#Optional
					failure_threshold = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_startup_probe_details_failure_threshold
					initial_delay_in_seconds = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_startup_probe_details_initial_delay_in_seconds
					period_in_seconds = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_configuration_details_startup_probe_details_period_in_seconds
				}
			}
			job_environment_configuration_details {
				#Required
				image = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_environment_configuration_details_image
				job_environment_type = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_environment_configuration_details_job_environment_type

				#Optional
				cmd = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_environment_configuration_details_cmd
				entrypoint = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_environment_configuration_details_entrypoint
				image_digest = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_environment_configuration_details_image_digest
				image_signature_id = oci_datascience_image_signature.test_image_signature.id
			}
			job_infrastructure_configuration_details {
				#Required
				job_infrastructure_type = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_infrastructure_configuration_details_job_infrastructure_type

				#Optional
				block_storage_size_in_gbs = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_infrastructure_configuration_details_block_storage_size_in_gbs
				job_shape_config_details {

					#Optional
					memory_in_gbs = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_infrastructure_configuration_details_job_shape_config_details_memory_in_gbs
					ocpus = var.job_job_node_configuration_details_job_node_group_configuration_details_list_job_infrastructure_configuration_details_job_shape_config_details_ocpus
				}
				shape_name = oci_core_shape.test_shape.name
				subnet_id = oci_core_subnet.test_subnet.id
			}
			minimum_success_replicas = var.job_job_node_configuration_details_job_node_group_configuration_details_list_minimum_success_replicas
			replicas = var.job_job_node_configuration_details_job_node_group_configuration_details_list_replicas
		}
		maximum_runtime_in_minutes = var.job_job_node_configuration_details_maximum_runtime_in_minutes
		startup_order = var.job_job_node_configuration_details_startup_order
	}
	job_storage_mount_configuration_details_list {
		#Required
		destination_directory_name = var.job_job_storage_mount_configuration_details_list_destination_directory_name
		storage_type = var.job_job_storage_mount_configuration_details_list_storage_type

		#Optional
		bucket = var.job_job_storage_mount_configuration_details_list_bucket
		destination_path = var.job_job_storage_mount_configuration_details_list_destination_path
		export_id = oci_file_storage_export.test_export.id
		mount_target_id = oci_file_storage_mount_target.test_mount_target.id
		namespace = var.job_job_storage_mount_configuration_details_list_namespace
		prefix = var.job_job_storage_mount_configuration_details_list_prefix
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the job.
* `delete_related_job_runs` - (Optional) (Updatable) Delete all related JobRuns upon deletion of the Job. 
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `job_configuration_details` - (Optional) The job configuration details 
	* `command_line_arguments` - (Applicable when job_type=DEFAULT) The arguments to pass to the job. 
	* `environment_variables` - (Applicable when job_type=DEFAULT) Environment variables to set for the job. 
	* `job_type` - (Required) The type of job.
	* `maximum_runtime_in_minutes` - (Applicable when job_type=DEFAULT) A time bound for the execution of the job. Timer starts when the job becomes active. 
	* `startup_probe_details` - (Applicable when job_type=DEFAULT) The probe indicates whether the application within the job run has started.
		* `command` - (Required) The commands to run in the target job run to perform the startup probe
		* `failure_threshold` - (Optional) How many times the job will try before giving up when a probe fails.
		* `initial_delay_in_seconds` - (Optional) Number of seconds after the job run has started before a startup probe is initiated.
		* `job_probe_check_type` - (Required) The probe check type to perform the startup probe and specifies the type of health check for a job.
		* `period_in_seconds` - (Optional) Number of seconds how often the job run should perform a startup probe
* `job_environment_configuration_details` - (Optional) Environment configuration to capture job runtime dependencies.
	* `cmd` - (Optional) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
	* `entrypoint` - (Optional) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
	* `image` - (Required) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
	* `image_digest` - (Optional) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
	* `image_signature_id` - (Optional) OCID of the container image signature
	* `job_environment_type` - (Required) The environment configuration type used for job runtime.
* `job_infrastructure_configuration_details` - (Optional) (Updatable) The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - (Required when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) (Updatable) The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - (Required) (Updatable) The infrastructure type used for job run.
	* `job_shape_config_details` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) (Updatable) Details for the job run shape configuration. Specify only when a flex shape is selected.
		* `cpu_baseline` - (Applicable when job_infrastructure_type=ME_STANDALONE | STANDALONE) (Updatable) The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) (Updatable) The total amount of memory available to the job run instance, in gigabytes. 
		* `ocpus` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) (Updatable) The total number of OCPUs available to the job run instance. 
	* `shape_name` - (Required when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) (Updatable) The name that corresponds to the JobShapeSummary to use for the job node
	* `subnet_id` - (Required when job_infrastructure_type=STANDALONE) (Updatable) The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_details` - (Optional) Logging configuration for resource. 
	* `enable_auto_log_creation` - (Optional) If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - (Optional) If customer logging is enabled for job runs.
	* `log_group_id` - (Optional) The log group id for where log objects are for job runs. 
	* `log_id` - (Optional) The log id the job run will push logs too. 
* `job_node_configuration_details` - (Optional) The job node configuration details
	* `job_network_configuration` - (Optional) The job network configuration details 
		* `job_network_type` - (Required) job network type
		* `subnet_id` - (Required when job_network_type=CUSTOM_NETWORK) The custom subnet id
	* `job_node_group_configuration_details_list` - (Optional) List of JobNodeGroupConfigurationDetails
		* `job_configuration_details` - (Optional) The job configuration details 
			* `command_line_arguments` - (Applicable when job_type=DEFAULT) The arguments to pass to the job. 
			* `environment_variables` - (Applicable when job_type=DEFAULT) Environment variables to set for the job. 
			* `job_type` - (Required) The type of job.
			* `maximum_runtime_in_minutes` - (Applicable when job_type=DEFAULT) A time bound for the execution of the job. Timer starts when the job becomes active. 
			* `startup_probe_details` - (Applicable when job_type=DEFAULT) The probe indicates whether the application within the job run has started.
				* `command` - (Required) The commands to run in the target job run to perform the startup probe
				* `failure_threshold` - (Optional) How many times the job will try before giving up when a probe fails.
				* `initial_delay_in_seconds` - (Optional) Number of seconds after the job run has started before a startup probe is initiated.
				* `job_probe_check_type` - (Required) The probe check type to perform the startup probe and specifies the type of health check for a job.
				* `period_in_seconds` - (Optional) Number of seconds how often the job run should perform a startup probe
		* `job_environment_configuration_details` - (Optional) Environment configuration to capture job runtime dependencies.
			* `cmd` - (Optional) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
			* `entrypoint` - (Optional) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
			* `image` - (Required) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
			* `image_digest` - (Optional) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
			* `image_signature_id` - (Optional) OCID of the container image signature
			* `job_environment_type` - (Required) The environment configuration type used for job runtime.
		* `job_infrastructure_configuration_details` - (Optional) The job infrastructure configuration details (shape, block storage, etc.) 
			* `block_storage_size_in_gbs` - (Required when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) The size of the block storage volume to attach to the instance running the job 
			* `job_infrastructure_type` - (Required) The infrastructure type used for job run.
			* `job_shape_config_details` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) Details for the job run shape configuration. Specify only when a flex shape is selected.
				* `memory_in_gbs` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) The total amount of memory available to the job run instance, in gigabytes. 
				* `ocpus` - (Applicable when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) The total number of OCPUs available to the job run instance. 
			* `shape_name` - (Required when job_infrastructure_type=ME_STANDALONE | MULTI_NODE | STANDALONE) The name that corresponds to the JobShapeSummary to use for the job node
			* `subnet_id` - (Required when job_infrastructure_type=STANDALONE) The subnet to create a secondary vnic in to attach to the instance running the job 
		* `minimum_success_replicas` - (Optional) The minimum threshold of successful replicas for node group to be successful. All replicas need to succeed if this is not specified.
		* `name` - (Required) node group name.
		* `replicas` - (Optional) The number of nodes.
	* `job_node_type` - (Required) The node type used for job run.
	* `maximum_runtime_in_minutes` - (Optional) A time bound for the execution of the job run. Timer starts when the job run is in progress. 
	* `startup_order` - (Optional) The execution order of node groups
* `job_storage_mount_configuration_details_list` - (Optional) (Updatable) Collection of JobStorageMountConfigurationDetails.
	* `bucket` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage bucket
	* `destination_directory_name` - (Required) (Updatable) The local directory name to be mounted
	* `destination_path` - (Optional) (Updatable) The local path of the mounted directory, excluding directory name.
	* `export_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the export
	* `mount_target_id` - (Required when storage_type=FILE_STORAGE) (Updatable) OCID of the mount target
	* `namespace` - (Required when storage_type=OBJECT_STORAGE) (Updatable) The object storage namespace
	* `prefix` - (Applicable when storage_type=OBJECT_STORAGE) (Updatable) Prefix in the bucket to mount
	* `storage_type` - (Required) (Updatable) The type of storage.
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `job_artifact` - (Optional) The job artifact to upload. This can be done in a separate step or from cli/sdk. The Job will remain in "Creating" state until its artifact is uploaded.
* `artifact_content_disposition` - (Optional, Required if `job_artifact` is set) This header allows you to specify a filename during upload. This file name is used to dispose of the file contents while downloading the file. Example: `attachment; filename=job-artifact.py`
* `artifact_content_length` - (Optional, Required if `job_artifact` is set) The content length of the body.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the project.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the job.
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
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
		* `cpu_baseline` - The baseline OCPU utilization for a subcore burstable VM instance. If this attribute is left blank, it will default to `BASELINE_1_1`. The following values are supported: BASELINE_1_8 - baseline usage is 1/8 of an OCPU. BASELINE_1_2 - baseline usage is 1/2 of an OCPU. BASELINE_1_1 - baseline usage is an entire OCPU. This represents a non-burstable instance. 
		* `memory_in_gbs` - The total amount of memory available to the job run instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the job run instance. 
	* `shape_name` - The name that corresponds to the JobShapeSummary to use for the job node
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_details` - Logging configuration for resource. 
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - If customer logging is enabled for job runs.
	* `log_group_id` - The log group id for where log objects are for job runs. 
	* `log_id` - The log id the job run will push logs too. 
* `job_node_configuration_details` - The job node configuration details
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
* `lifecycle_details` - The state of the job.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `state` - The state of the job.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Job
	* `update` - (Defaults to 20 minutes), when updating the Job
	* `delete` - (Defaults to 20 minutes), when destroying the Job


## Import

Jobs can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_job.test_job "id"
```

