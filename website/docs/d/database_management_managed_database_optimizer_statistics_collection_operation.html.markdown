---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_optimizer_statistics_collection_operation"
sidebar_current: "docs-oci-datasource-database_management-managed_database_optimizer_statistics_collection_operation"
description: |-
  Provides details about a specific Managed Database Optimizer Statistics Collection Operation in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_optimizer_statistics_collection_operation
This data source provides details about a specific Managed Database Optimizer Statistics Collection Operation resource in Oracle Cloud Infrastructure Database Management service.

Gets a detailed report of the Optimizer Statistics Collection operation for the specified Managed Database.

## Example Usage

```hcl
data "oci_database_management_managed_database_optimizer_statistics_collection_operation" "test_managed_database_optimizer_statistics_collection_operation" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	optimizer_statistics_collection_operation_id = oci_database_management_optimizer_statistics_collection_operation.test_optimizer_statistics_collection_operation.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `optimizer_statistics_collection_operation_id` - (Required) The ID of the Optimizer Statistics Collection operation.


## Attributes Reference

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

