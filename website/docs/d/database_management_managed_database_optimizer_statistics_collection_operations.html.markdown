---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_optimizer_statistics_collection_operations"
sidebar_current: "docs-oci-datasource-database_management-managed_database_optimizer_statistics_collection_operations"
description: |-
  Provides the list of Managed Database Optimizer Statistics Collection Operations in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_optimizer_statistics_collection_operations
This data source provides the list of Managed Database Optimizer Statistics Collection Operations in Oracle Cloud Infrastructure Database Management service.

Lists the optimizer statistics (Auto and Manual) task operation summary for the specified Managed Database.
The summary includes the details of each operation and the number of tasks grouped by status: Completed, In Progress, Failed, and so on.
Optionally, you can specify a date-time range (of seven days) to obtain the list of operations that fall within the specified time range.
If the date-time range is not specified, then the operations in the last seven days are listed.
This API also enables the pagination of results and the opc-next-page response header indicates whether there is a next page.
If you use the same header value in a consecutive request, the next page records are returned.
To obtain the required results, you can apply the different types of filters supported by this API.


## Example Usage

```hcl
data "oci_database_management_managed_database_optimizer_statistics_collection_operations" "test_managed_database_optimizer_statistics_collection_operations" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	end_time_less_than_or_equal_to = var.managed_database_optimizer_statistics_collection_operation_end_time_less_than_or_equal_to
	filter_by = var.managed_database_optimizer_statistics_collection_operation_filter_by
	start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_operation_start_time_greater_than_or_equal_to
	task_type = var.managed_database_optimizer_statistics_collection_operation_task_type
}
```

## Argument Reference

The following arguments are supported:

* `end_time_less_than_or_equal_to` - (Optional) The end time of the time range to retrieve the optimizer statistics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 
* `filter_by` - (Optional) The parameter used to filter the optimizer statistics operations. Any property of the OptimizerStatisticsCollectionOperationSummary can be used to define the filter condition. The allowed conditional operators are AND or OR, and the allowed binary operators are are >, < and =. Any other operator is regarded invalid. Example: jobName=<replace with job name> AND status=<replace with status> 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `start_time_greater_than_or_equal_to` - (Optional) The start time of the time range to retrieve the optimizer statistics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 
* `task_type` - (Optional) The filter types of the optimizer statistics tasks.


## Attributes Reference

The following attributes are exported:

* `optimizer_statistics_collection_operations_collection` - The list of optimizer_statistics_collection_operations_collection.

### ManagedDatabaseOptimizerStatisticsCollectionOperation Reference

The following attributes are exported:

* `completed_count` - The number of objects for which statistics collection is completed.
* `database` - The summary of the Managed Database resource.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database resides.
	* `db_deployment_type` - The infrastructure used to deploy the Oracle Database.
	* `db_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, Non-container Database, Autonomous Database, or Autonomous Container Database. 
	* `db_type` - The type of Oracle Database installation.
	* `db_version` - The version of the Oracle Database.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
	* `name` - The name of the Managed Database.
* `duration_in_seconds` - The time it takes to complete the operation (in seconds).
* `end_time` - The end time of the operation.
* `failed_count` - The number of objects for which statistics collection failed.
* `id` - The ID of the operation.
* `in_progress_count` - The number of objects for which statistics collection is in progress.
* `job_name` - The name of the job.
* `operation_name` - The name of the operation.
* `start_time` - The start time of the operation.
* `status` - The status of the operation such as Completed, and Failed.
* `target` - The target object type such as Table, Index, and Partition.
* `tasks` - An array of Optimizer Statistics Collection task details.
	* `status` - The status of the Optimizer Statistics Collection task.
	* `target` - The name of the target object for which statistics are gathered.
	* `target_type` - The type of target object.
	* `time_end` - The end time of the Optimizer Statistics Collection task.
	* `time_start` - The start time of the Optimizer Statistics Collection task.
* `timed_out_count` - The number of objects for which statistics collection timed out.
* `total_objects_count` - The total number of objects for which statistics is collected. This number is the sum of all the objects with various statuses: completed, inProgress, failed, and timedOut. 

