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
* `job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job run.
* `job_infrastructure_configuration_details` - The job infrastructure configuration details (shape, block storage, etc.) 
	* `block_storage_size_in_gbs` - The size of the block storage volume to attach to the instance running the job 
	* `job_infrastructure_type` - The infrastructure type used for job run.
	* `shape_name` - The shape used to launch the job run instances.
	* `subnet_id` - The subnet to create a secondary vnic in to attach to the instance running the job 
* `job_log_configuration_override_details` - Logging configuration for resource. 
	* `enable_auto_log_creation` - If automatic on-behalf-of log object creation is enabled for job runs. 
	* `enable_logging` - If customer logging is enabled for job runs.
	* `log_group_id` - The log group id for where log objects are for job runs. 
	* `log_id` - The log id the job run will push logs too. 
* `lifecycle_details` - Details of the state of the job run.
* `log_details` - Customer logging details for job run. 
	* `log_group_id` - The log group id for where log objects will be for job runs. 
	* `log_id` - The log id of the log object the job run logs will be shipped to. 
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project to associate the job with.
* `state` - The state of the job run.
* `time_accepted` - The date and time the job run was accepted in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_finished` - The date and time the job run request was finished in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_started` - The date and time the job run request was started in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

