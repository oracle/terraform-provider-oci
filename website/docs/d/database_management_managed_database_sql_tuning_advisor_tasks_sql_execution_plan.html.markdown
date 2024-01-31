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

Retrieves a SQL execution plan for the SQL being tuned.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan" {
	#Required
	attribute = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_attribute
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_object_id = oci_objectstorage_object.test_object.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

	#Optional
	opc_named_credential_id = var.managed_database_sql_tuning_advisor_tasks_sql_execution_plan_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `attribute` - (Required) The attribute of the SQL execution plan.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `sql_object_id` - (Required) The SQL object ID for the SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `plan` - A SQL execution plan as a list of steps.
	* `access_predicates` - The predicates used to locate rows in an access structure. For example, start or stop predicates for an index range scan. 
	* `attribute` - The text string identifying the type of execution plan.
	* `bytes` - The number of bytes returned by the current operation.
	* `cardinality` - The number of rows returned by the current operation (estimated by the CBO).
	* `cost` - The cost of the current operation estimated by the cost-based optimizer (CBO).
	* `cpu_cost` - The CPU cost of the current operation.
	* `filter_predicates` - The predicates used to filter rows before producing them.
	* `io_cost` - The I/O cost of the current operation.
	* `number_of_search_column` - Number of index columns with start and stop keys (that is, the number of columns with matching predicates). 
	* `object` - The name of the object.
	* `object_node` - The name of the database link used to reference the object.
	* `object_owner` - The owner of the object.
	* `object_position` - The numbered position of the object name in the original SQL statement.
	* `object_type` - The descriptive modifier that further describes the type of object.
	* `operation` - The name of the operation performed at this step.
	* `optimizer_mode` - The current mode of the optimizer, such as all_rows, first_rows_n (where n = 1, 10, 100, 1000, and so on).
	* `options` - The options used for the operation performed at this step.
	* `other` - Information about parallel execution servers and parallel queries
	* `other_tag` - Describes the function of the SQL text in the OTHER column.
	* `parent_step_id` - The ID of the next step that operates on the results of this step. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `partition_id` - The ID of the step in the execution plan that has computed the pair of values of partitionStart and partitionStop. 
	* `partition_start` - A step may get data from a range of partitions of a partitioned object, such as table or index, based on predicates and sorting order. The partionStart is the starting partition of the range. The partitionStop is the ending partition of the range. 
	* `partition_stop` - A step may get data from a range of partitions of a partitioned object, such as table or index, based on predicates and sorting order. The partionStart is the starting partition of the range. The partitionStop is the ending partition of the range. 
	* `plan_hash_value` - The numerical representation of the SQL execution plan.
	* `position` - The order of processing for steps with the same parent ID.
	* `remarks` - The place for comments that can be added to the steps of the execution plan.
	* `step_id` - The identification number of a step in the SQL execution plan. This is unique within the SQL execution plan. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
	* `temp_space` - The temporary space usage (in bytes) of the operation (sort or hash-join) as estimated by the CBO.
	* `time` - The elapsed time (in seconds) of the operation as estimated by the CBO.

