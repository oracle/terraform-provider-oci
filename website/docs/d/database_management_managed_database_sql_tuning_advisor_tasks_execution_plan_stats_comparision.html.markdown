---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision"
description: |-
  Provides details about a specific Managed Database Sql Tuning Advisor Tasks Execution Plan Stats Comparision in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision
This data source provides details about a specific Managed Database Sql Tuning Advisor Tasks Execution Plan Stats Comparision resource in Oracle Cloud Infrastructure Database Management service.

Retrieves a comparison of the existing SQL execution plan and a new plan.
A SQL tuning task may suggest a new execution plan for a SQL,
and this API retrieves the comparison report of the statistics of the two plans.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" "test_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision" {
	#Required
	execution_id = oci_database_management_execution.test_execution.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_object_id = oci_objectstorage_object.test_object.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

	#Optional
	opc_named_credential_id = var.managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `execution_id` - (Required) The execution ID for an execution of a SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `sql_object_id` - (Required) The SQL object ID for the SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `modified` - The statistics of a SQL execution plan. 
	* `plan_stats` - A map contains the statistics for the SQL execution using the plan. The key of the map is the metric's name. The value of the map is the metric's value. 
	* `plan_status` - The status of the execution using the plan. 
	* `plan_type` - The type of the original or modified plan with profile, index, and so on.
* `original` - The statistics of a SQL execution plan. 
	* `plan_stats` - A map contains the statistics for the SQL execution using the plan. The key of the map is the metric's name. The value of the map is the metric's value. 
	* `plan_status` - The status of the execution using the plan. 
	* `plan_type` - The type of the original or modified plan with profile, index, and so on.

