---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_sql_execution_plan"
description: |-
  Provides details about a specific Managed Database Sql Tuning Advisor Tasks Sql Execution Plan in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan
This data source provides details about a specific Managed Database Sql Tuning Advisor Tasks Sql Execution Plan resource in Oracle Cloud Infrastructure Database Management service.

Retrieve a SQL execution plan for a SQL being tuned, for original or new plan


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" {
	#Required
	attribute = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_object_id = oci_objectstorage_object.test_object.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}
```

## Argument Reference

The following arguments are supported:

* `attribute` - (Required) The attribute of the SQL execution plan.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `sql_object_id` - (Required) The SQL object id for the SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `plan` - A SQL execution plan as a list of steps
	* `access_predicates` - Predicates used to locate rows in an access structure. For example, start or stop predicates for an index range scan. 
	* `attribute` - Text string identifying the type of the execution plan.
	* `bytes` - Number of bytes returned by the current operation.
	* `cardinality` - Number of rows returned by the current operation (estimated by the CBO).
	* `cost` - Cost of the current operation estimated by the cost-based optimizer (CBO).
	* `cpu_cost` - The CPU cost of the current operation.
	* `filter_predicates` - Predicates used to filter rows before producing them.
	* `io_cost` - The I/O cost of the current operation.
	* `number_of_search_column` - Number of index columns with start and stop keys (that is, the number of columns with matching predicates) 
	* `object` - Name of the object.
	* `object_node` - Name of the database link used to reference the object.
	* `object_owner` - Owner of the object.
	* `object_position` - Numbered position of the object name in the original SQL statement.
	* `object_type` - Descriptive modifier that further describes the type of object.
	* `operation` - Name of the operation performed at this step
	* `optimizer_mode` - Current mode of the optimizer, such as all_rows, first_rows_n (where n = 1, 10, 100, 1000 etc).
	* `options` - Options used for the operation performed at this step.
	* `other` - Information about parallel execution servers and parallel queries
	* `other_tag` - Describes the function of the SQL text in the OTHER column.
	* `parent_step_id` - ID of the next step that operates on the results of this step. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `partition_id` - The id of the step in the execution plan that has computed the pair of values of the partitionStart and partitionStop 
	* `partition_start` - A step may get data from a range of partitions of a partitioned object, such table or index, based on predicates and sorting order. The partionStart is the starting partition of the range. The partitionStop is the ending partition of the range 
	* `partition_stop` - A step may get data from a range of partitions of a partitioned object, such table or index, based on predicates and sorting order. The partionStart is the starting partition of the range. The partitionStop is the ending partition of the range 
	* `plan_hash_value` - Numerical representation of the execution plan
	* `position` - Order of processing for steps with the same parent ID.
	* `remarks` - Place for comments that can be added to the steps of the execution plan.
	* `step_id` - Identification number for this step in the execution plan. It is unique within the execution plan. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `temp_space` - Temporary space usage (in bytes) of the operation (sort or hash-join) as estimated by the CBO.
	* `time` - Elapsed time (in seconds) of the operation as estimated by the CBO.

