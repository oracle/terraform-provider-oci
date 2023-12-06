---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_optimizer_statistics_collection_aggregations"
sidebar_current: "docs-oci-datasource-database_management-managed_database_optimizer_statistics_collection_aggregations"
description: |-
  Provides the list of Managed Database Optimizer Statistics Collection Aggregations in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_optimizer_statistics_collection_aggregations
This data source provides the list of Managed Database Optimizer Statistics Collection Aggregations in Oracle Cloud Infrastructure Database Management service.

Gets a list of the optimizer statistics collection operations per hour, grouped by task or object status for the specified Managed Database.
You must specify a value for the GroupByQueryParam to determine whether the data should be grouped by task status or task object status.
Optionally, you can specify a date-time range (of seven days) to obtain collection aggregations within the specified time range.
If the date-time range is not specified, then the operations in the last seven days are listed.
You can further filter the results by providing the optional type of TaskTypeQueryParam.
If the task type if not provided, then both Auto and Manual tasks are considered for aggregation.


## Example Usage

```hcl
data "oci_database_management_managed_database_optimizer_statistics_collection_aggregations" "test_managed_database_optimizer_statistics_collection_aggregations" {
	#Required
	group_type = var.managed_database_optimizer_statistics_collection_aggregation_group_type
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	end_time_less_than_or_equal_to = var.managed_database_optimizer_statistics_collection_aggregation_end_time_less_than_or_equal_to
	start_time_greater_than_or_equal_to = var.managed_database_optimizer_statistics_collection_aggregation_start_time_greater_than_or_equal_to
	task_type = var.managed_database_optimizer_statistics_collection_aggregation_task_type
}
```

## Argument Reference

The following arguments are supported:

* `end_time_less_than_or_equal_to` - (Optional) The end time of the time range to retrieve the optimizer statistics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 
* `group_type` - (Required) The optimizer statistics tasks grouped by type.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `start_time_greater_than_or_equal_to` - (Optional) The start time of the time range to retrieve the optimizer statistics of a Managed Database in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'". 
* `task_type` - (Optional) The filter types of the optimizer statistics tasks.


## Attributes Reference

The following attributes are exported:

* `optimizer_statistics_collection_aggregations_collection` - The list of optimizer_statistics_collection_aggregations_collection.

### ManagedDatabaseOptimizerStatisticsCollectionAggregation Reference

The following attributes are exported:

* `items` - The list of Optimizer Statistics Collection details.
	* `completed` - The number of tasks or objects for which statistics gathering is completed.
	* `failed` - The number of tasks or objects for which statistics gathering failed.
	* `group_by` - The optimizer statistics tasks grouped by type.
	* `in_progress` - The number of tasks or objects for which statistics gathering is in progress.
	* `pending` - The number of tasks or objects for which statistics are yet to be gathered.
	* `skipped` - The number of tasks or objects for which statistics gathering was skipped.
	* `time_end` - Indicates the end of the hour as the statistics are aggregated per hour.
	* `time_start` - Indicates the start of the hour as the statistics are aggregated per hour.
	* `timed_out` - The number of tasks or objects for which statistics gathering timed out.
	* `total` - The total number of tasks or objects for which statistics collection is finished. This number is the sum of all the tasks or objects with various statuses: pending, inProgress, completed, failed, skipped, timedOut, and unknown. 
	* `unknown` - The number of tasks or objects for which the status of statistics gathering is unknown.

