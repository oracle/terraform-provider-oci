---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_summary_report"
description: |-
  Provides details about a specific Managed Database Sql Tuning Advisor Tasks Summary Report in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report
This data source provides details about a specific Managed Database Sql Tuning Advisor Tasks Summary Report resource in Oracle Cloud Infrastructure Database Management service.

Gets the summary report for the specific SQL Tuning Advisor task.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report" "test_managed_database_sql_tuning_advisor_tasks_summary_report" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id

	#Optional
	begin_exec_id_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_begin_exec_id_greater_than_or_equal_to
	end_exec_id_less_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_end_exec_id_less_than_or_equal_to
	search_period = var.managed_database_sql_tuning_advisor_tasks_summary_report_search_period
	time_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_greater_than_or_equal_to
	time_less_than_or_equal_to = var.managed_database_sql_tuning_advisor_tasks_summary_report_time_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `begin_exec_id_greater_than_or_equal_to` - (Optional) The optional greater than or equal to filter on the execution ID related to a specific SQL Tuning Advisor task. This is applicable only for Auto SQL Tuning tasks.
* `end_exec_id_less_than_or_equal_to` - (Optional) The optional less than or equal to query parameter to filter on the execution ID related to a specific SQL Tuning Advisor task. This is applicable only for Auto SQL Tuning tasks.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `search_period` - (Optional) How far back the API will search for begin and end exec id. Unused if neither exec ids nor time filter query params are supplied. This is applicable only for Auto SQL Tuning tasks.
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `time_greater_than_or_equal_to` - (Optional) The optional greater than or equal to query parameter to filter the timestamp. This is applicable only for Auto SQL Tuning tasks.
* `time_less_than_or_equal_to` - (Optional) The optional less than or equal to query parameter to filter the timestamp. This is applicable only for Auto SQL Tuning tasks.


## Attributes Reference

The following attributes are exported:

* `index_findings` - The list of object findings related to indexes.
	* `index_columns` - Columns of the index.
	* `index_hash_value` - Numerical representation of the index.
	* `index_name` - Name of the index.
	* `reference_count` - The number of times the index is referenced within the SQL Tuning advisor task findings.
	* `schema` - Schema related to the index.
	* `table_name` - Table's name related to the index.
* `object_stat_findings` - The list of object findings related to statistics.
	* `object` - Name of the object.
	* `object_hash_value` - Numerical representation of the object.
	* `object_type` - Type of the object.
	* `problem_type` - Type of statistics problem related to the object.
	* `reference_count` - The number of the times the object is referenced within the SQL Tuning advisor task findings.
	* `schema` - Schema of the object.
* `statistics` - Statistics of statements and findings for the SQL Tuning Advisor summary report.
	* `finding_benefits` - The finding benefits data for the SQL Tuning Advisor summary report.
		* `db_time_after_implemented` - The count of database time benefit after SQL recommendations are implemented.
		* `db_time_after_recommended` - The count of Potential database time after SQL recommendations are implemented.
		* `db_time_before_implemented` - The count of database time benefit before SQL recommendations are implemented.
		* `db_time_before_recommended` - The count of Potential database time before SQL recommendations are implemented.
	* `finding_counts` - The finding counts data for the SQL Tuning Advisor summary report.
		* `alternate_plan` - The count of distinct SQL statements with alternative plan recommendations.
		* `implemented_sql_profile` - The count of distinct SQL statements with implemented SQL profiles.
		* `index` - The count of distinct SQL statements with index recommendations.
		* `recommended_sql_profile` - The count of distinct SQL statements with recommended SQL profiles.
		* `restructure` - The count of distinct SQL statements with restructure SQL recommendations.
		* `statistics` - The count of distinct SQL statements with stale/missing optimizer statistics recommendations.
	* `statement_counts` - The statement counts data for the SQL Tuning Advisor summary report.
		* `distinct_sql` - The count of distinct SQL statements.
		* `error_count` - The count of distinct SQL statements with errors.
		* `finding_count` - The count of distinct SQL statements with findings.
		* `total_sql` - The total count of SQL statements.
* `task_info` - SQL Tuning advisor task general info.
	* `description` - The SQL Tuning Advisor task description. Not defined on Auto SQL Tuning tasks.
	* `id` - The SQL Tuning Advisor task id. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `name` - The SQL Tuning Advisor task name.
	* `owner` - The SQL Tuning Advisor task user owner.
	* `running_time` - The total running time in seconds. Not defined on Auto SQL Tuning tasks.
	* `status` - The SQL Tuning Advisor task status. Not defined on Auto SQL Tuning tasks.
	* `time_ended` - End timestamp of task execution.
	* `time_started` - Start timestamp of task execution.

