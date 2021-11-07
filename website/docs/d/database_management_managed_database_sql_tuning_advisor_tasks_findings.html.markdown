---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_findings"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_findings"
description: |-
  Provides the list of Managed Database Sql Tuning Advisor Tasks Findings in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_findings
This data source provides the list of Managed Database Sql Tuning Advisor Tasks Findings in Oracle Cloud Infrastructure Database Management service.

Takes in a task id, and a finding/object type filter and applies some SQLs to find return the output.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_findings" "test_managed_database_sql_tuning_advisor_tasks_findings" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

	#Optional
	begin_exec_id = oci_database_management_begin_exec.test_begin_exec.id
	end_exec_id = oci_database_management_end_exec.test_end_exec.id
	finding_filter = var.managed_database_sql_tuning_advisor_tasks_finding_finding_filter
	index_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_index_hash_filter
	search_period = var.managed_database_sql_tuning_advisor_tasks_finding_search_period
	stats_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter
}
```

## Argument Reference

The following arguments are supported:

* `begin_exec_id` - (Optional) The optional greater than or equal to filter on the execution ID related to a specific SQL Tuning Advisor task.
* `end_exec_id` - (Optional) The optional less than or equal to query parameter to filter on the execution ID related to a specific SQL Tuning Advisor task.
* `finding_filter` - (Optional) Filters which findings get shown in the report
* `index_hash_filter` - (Optional) The hash value of the index table name.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `search_period` - (Optional) How far back the API will search for begin and end exec id, if not supplied. Unused if beginExecId and endExecId optional query params are both supplied.
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `stats_hash_filter` - (Optional) The hash value of the object for the statistic finding search.


## Attributes Reference

The following attributes are exported:

* `sql_tuning_advisor_task_finding_collection` - The list of sql_tuning_advisor_task_finding_collection.

### ManagedDatabaseSqlTuningAdvisorTasksFinding Reference

The following attributes are exported:

* `items` - An array of the findings for a tuning task.
	* `db_time_benefit` - Time benefit in seconds for the highest-rated finding for this object.
	* `is_alternative_plan_finding_present` - Whether an alternative execution plan was found for this SQL statement.
	* `is_error_finding_present` - Whether there is an error in this SQL statement.
	* `is_index_finding_present` - Whether an index recommendation was found for this SQL statement.
	* `is_miscellaneous_finding_present` - Whether a miscellaneous finding was found for this SQL statement.
	* `is_restructure_sql_finding_present` - Whether a restructure SQL recommendation was found for this SQL statement.
	* `is_sql_profile_finding_implemented` - Whether a SQL Profile recommendation has been implemented for this SQL statement.
	* `is_sql_profile_finding_present` - Whether a SQL Profile recommendation was found for this SQL statement.
	* `is_stats_finding_present` - Whether a statistics recommendation was found for this SQL statement.
	* `is_timeout_finding_present` - Whether the task timed out.
	* `parsing_schema` - Parsing schema of the object.
	* `per_execution_percentage` - The per-execution percentage benefit.
	* `sql_key` - Unique key of this SQL statement
	* `sql_text` - Text of the SQL statement.
	* `sql_tuning_advisor_task_id` - Unique identifier of the task. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_execution_id` - Execution id of the analyzed SQL object. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_id` - Key of the object to which these recommendations apply. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

