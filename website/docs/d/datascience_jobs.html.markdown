---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_jobs"
sidebar_current: "docs-oci-datasource-datascience-jobs"
description: |-
  Provides the list of Jobs in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_jobs
This data source provides the list of Jobs in Oracle Cloud Infrastructure Data Science service.

List jobs in the specified compartment.

## Example Usage

```hcl
data "oci_datascience_jobs" "test_jobs" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	created_by = var.job_created_by
	display_name = var.job_display_name
	id = var.job_id
	project_id = oci_datascience_project.test_project.id
	state = var.job_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `created_by` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the resource.
* `display_name` - (Optional) <b>Filter</b> results by its user-friendly name.
* `id` - (Optional) <b>Filter</b> results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type. 
* `project_id` - (Optional) <b>Filter</b> results by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project.
* `state` - (Optional) <b>Filter</b> results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `jobs` - The list of jobs.

### Job Reference

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
	* `shape_name` - The shape used to launch the job run instances.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_details` - Logging configuration for resource. 
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
* `lifecycle_details` - The state of the job.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `state` - The state of the job.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

