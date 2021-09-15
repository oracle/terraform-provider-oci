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
	job_configuration_details {
		#Required
		job_type = var.job_job_configuration_details_job_type

		#Optional
		command_line_arguments = var.job_job_configuration_details_command_line_arguments
		environment_variables = var.job_job_configuration_details_environment_variables
		maximum_runtime_in_minutes = var.job_job_configuration_details_maximum_runtime_in_minutes
	}
	job_infrastructure_configuration_details {
		#Required
		block_storage_size_in_gbs = var.job_job_infrastructure_configuration_details_block_storage_size_in_gbs
		job_infrastructure_type = var.job_job_infrastructure_configuration_details_job_infrastructure_type
		shape_name = oci_core_shape.test_shape.name
		subnet_id = oci_core_subnet.test_subnet.id
	}
	project_id = oci_datascience_project.test_project.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.job_description
	display_name = var.job_display_name
	freeform_tags = {"Department"= "Finance"}
	job_log_configuration_details {

		#Optional
		enable_auto_log_creation = var.job_job_log_configuration_details_enable_auto_log_creation
		enable_logging = var.job_job_log_configuration_details_enable_logging
		log_group_id = oci_logging_log_group.test_log_group.id
		log_id = oci_logging_log.test_log.id
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
* `job_configuration_details` - (Required) The job configuration details 
	* `command_line_arguments` - (Optional) The arguments to pass to the job. 
	* `environment_variables` - (Optional) Environment variables to set for the job. 
	* `job_type` - (Required) The type of job.
	* `maximum_runtime_in_minutes` - (Optional) A time bound for the execution of the job. Timer starts when the job becomes active. 
* `job_infrastructure_configuration_details` - (Required) (Updatable) The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - (Required) (Updatable) The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - (Required) (Updatable) The infrastructure type used for job run.
	* `shape_name` - (Required) (Updatable) The shape used to launch the job run instances.
	* `subnet_id` - (Required) (Updatable) The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_details` - (Optional) Logging configuration for resource. 
	* `enable_auto_log_creation` - (Optional) If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - (Optional) If customer logging is enabled for job runs.
	* `log_group_id` - (Optional) The log group id for where log objects are for job runs. 
	* `log_id` - (Optional) The log id the job run will push logs too. 
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
* `job_infrastructure_configuration_details` - The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - The infrastructure type used for job run.
	* `shape_name` - The shape used to launch the job run instances.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_details` - Logging configuration for resource. 
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - If customer logging is enabled for job runs.
	* `log_group_id` - The log group id for where log objects are for job runs. 
	* `log_id` - The log id the job run will push logs too. 
* `lifecycle_details` - The state of the job.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `state` - The state of the job.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2020-08-06T21:10:29.41Z 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Job
	* `update` - (Defaults to 20 minutes), when updating the Job
	* `delete` - (Defaults to 20 minutes), when destroying the Job


## Import

Jobs can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_job.test_job "id"
```

