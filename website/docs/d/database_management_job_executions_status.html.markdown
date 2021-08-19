---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_job_executions_status"
sidebar_current: "docs-oci-datasource-database_management-job_executions_status"
description: |-
  Provides details about a specific Job Executions Status in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_job_executions_status
This data source provides details about a specific Job Executions Status resource in Oracle Cloud Infrastructure Database Management service.

Gets the number of job executions grouped by status for a job, Managed Database, or Database Group in a specific compartment. Only one of the parameters, jobId, managedDatabaseId, or managedDatabaseGroupId should be provided.

## Example Usage

```hcl
data "oci_database_management_job_executions_status" "test_job_executions_status" {
	#Required
	compartment_id = var.compartment_id
	end_time = var.job_executions_status_end_time
	start_time = var.job_executions_status_start_time

	#Optional
	id = var.job_executions_status_id
	managed_database_group_id = oci_database_management_managed_database_group.test_managed_database_group.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	name = var.job_executions_status_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `end_time` - (Required) The end time of the time range to retrieve the status summary of job executions in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 
* `id` - (Optional) The identifier of the resource.
* `managed_database_group_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
* `managed_database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `start_time` - (Required) The start time of the time range to retrieve the status summary of job executions in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 


## Attributes Reference

The following attributes are exported:

* `items` - A list of JobExecutionsSummary objects.
	* `count` - The number of job executions of a particular status.
	* `status` - The status of the job execution.

