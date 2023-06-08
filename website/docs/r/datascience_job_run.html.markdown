---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_job_run"
sidebar_current: "docs-oci-resource-datascience-job_run"
description: |-
  Provides the Job Run resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_job_run
This resource provides the Job Run resource in Oracle Cloud Infrastructure Data Science service.

Creates a job run.

## Example Usage

```hcl
resource "oci_datascience_job_run" "test_job_run" {
	#Required
	compartment_id = var.compartment_id
	job_id = oci_datascience_job.test_job.id
	project_id = oci_datascience_project.test_project.id

	#Optional
  asynchronous = var.asynchronous
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.job_run_display_name
	freeform_tags = {"Department"= "Finance"}
	job_configuration_override_details {
		#Required
		job_type = var.job_run_job_configuration_override_details_job_type

		#Optional
		command_line_arguments = var.job_run_job_configuration_override_details_command_line_arguments
		environment_variables = var.job_run_job_configuration_override_details_environment_variables
		maximum_runtime_in_minutes = var.job_run_job_configuration_override_details_maximum_runtime_in_minutes
	}
	job_environment_configuration_override_details {
		#Required
		image = var.job_run_job_environment_configuration_override_details_image
		job_environment_type = var.job_run_job_environment_configuration_override_details_job_environment_type

		#Optional
		cmd = var.job_run_job_environment_configuration_override_details_cmd
		entrypoint = var.job_run_job_environment_configuration_override_details_entrypoint
		image_digest = var.job_run_job_environment_configuration_override_details_image_digest
		image_signature_id = oci_datascience_image_signature.test_image_signature.id
	}
	job_log_configuration_override_details {

		#Optional
		enable_auto_log_creation = var.job_run_job_log_configuration_override_details_enable_auto_log_creation
		enable_logging = var.job_run_job_log_configuration_override_details_enable_logging
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
	}
	opc_parent_rpt_url = var.job_run_opc_parent_rpt_url
}
```

## Argument Reference

The following arguments are supported:

* `asynchronous` - (Optional) If set to true, do not wait for the JobRun to reach completion prior to returning. Can be useful for JobRuns with a long duration.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `job_configuration_override_details` - (Optional) The job configuration details 
	* `command_line_arguments` - (Optional) The arguments to pass to the job. 
	* `environment_variables` - (Optional) Environment variables to set for the job. 
	* `job_type` - (Required) The type of job.
	* `maximum_runtime_in_minutes` - (Optional) A time bound for the execution of the job. Timer starts when the job becomes active. 
* `job_environment_configuration_override_details` - (Optional) Environment configuration to capture job runtime dependencies.
	* `cmd` - (Optional) The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
	* `entrypoint` - (Optional) The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
	* `image` - (Required) The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
	* `image_digest` - (Optional) The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
	* `image_signature_id` - (Optional) OCID of the container image signature
	* `job_environment_type` - (Required) The environment configuration type used for job runtime.
* `job_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job to create a run for.
* `job_log_configuration_override_details` - (Optional) Logging configuration for resource. 
	* `enable_auto_log_creation` - (Optional) If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - (Optional) If customer logging is enabled for job runs.
	* `log_group_id` - (Optional) The log group id for where log objects are for job runs. 
	* `log_id` - (Optional) The log id the job run will push logs too. 
* `opc_parent_rpt_url` - (Optional) URL to fetch the Resource Principal Token from the parent resource. 
* `project_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the job.
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
* `job_environment_configuration_override_details` - Environment configuration to capture job runtime dependencies.
	* `cmd` - The container image run [CMD](https://docs.docker.com/engine/reference/builder/#cmd) as a list of strings. Use `CMD` as arguments to the `ENTRYPOINT` or the only command to run in the absence of an `ENTRYPOINT`. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. 
	* `entrypoint` - The container image run [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint) as a list of strings. Accept the `CMD` as extra arguments. The combined size of `CMD` and `ENTRYPOINT` must be less than 2048 bytes. More information on how `CMD` and `ENTRYPOINT` interact are [here](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact). 
	* `image` - The full path to the Oracle Container Repository (OCIR) registry, image, and tag in a canonical format. Acceptable format: `<region>.ocir.io/<registry>/<image>:<tag>` `<region>.ocir.io/<registry>/<image>:<tag>@digest` 
	* `image_digest` - The digest of the container image. For example, `sha256:881303a6b2738834d795a32b4a98eb0e5e3d1cad590a712d1e04f9b2fa90a030` 
	* `image_signature_id` - OCID of the container image signature
	* `job_environment_type` - The environment configuration type used for job runtime.
* `job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.
* `job_infrastructure_configuration_details` - The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - The infrastructure type used for job run.
	* `job_shape_config_details` - Details for the job run shape configuration. Specify only when a flex shape is selected.
		* `memory_in_gbs` - The total amount of memory available to the job run instance, in gigabytes. 
		* `ocpus` - The total number of OCPUs available to the job run instance. 
	* `shape_name` - The shape used to launch the job run instances.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_override_details` - Logging configuration for resource. 
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - If customer logging is enabled for job runs.
	* `log_group_id` - The log group id for where log objects are for job runs. 
	* `log_id` - The log id the job run will push logs too. 
* `job_storage_mount_configuration_details_list` - Collection of JobStorageMountConfigurationDetails.
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
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `state` - The state of the job run.
* `time_accepted` - The date and time the job run was accepted in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_finished` - The date and time the job run request was finished in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_started` - The date and time the job run request was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Job Run
	* `update` - (Defaults to 20 minutes), when updating the Job Run
	* `delete` - (Defaults to 20 minutes), when destroying the Job Run


## Import

JobRuns can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_job_run.test_job_run "id"
```

