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

Gets an array of the details of the findings that match specific filters.


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
	opc_named_credential_id = var.managed_database_sql_tuning_advisor_tasks_finding_opc_named_credential_id
	search_period = var.managed_database_sql_tuning_advisor_tasks_finding_search_period
	stats_hash_filter = var.managed_database_sql_tuning_advisor_tasks_finding_stats_hash_filter
}
```

## Argument Reference

The following arguments are supported:

* `begin_exec_id` - (Optional) The optional greater than or equal to filter on the execution ID related to a specific SQL Tuning Advisor task.
* `end_exec_id` - (Optional) The optional less than or equal to query parameter to filter on the execution ID related to a specific SQL Tuning Advisor task.
* `finding_filter` - (Optional) The filter used to display specific findings in the report.
* `index_hash_filter` - (Optional) The hash value of the index table name.
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `search_period` - (Optional) The search period during which the API will search for begin and end exec id, if not supplied. Unused if beginExecId and endExecId optional query params are both supplied. 
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `stats_hash_filter` - (Optional) The hash value of the object for the statistic finding search.


## Attributes Reference

The following attributes are exported:

* `sql_tuning_advisor_task_finding_collection` - The list of sql_tuning_advisor_task_finding_collection.

### ManagedDatabaseSqlTuningAdvisorTasksFinding Reference

The following attributes are exported:

* `items` - An array of the findings for a tuning task.
	* `db_time_benefit` - The time benefit (in seconds) for the highest-rated finding for this object.
	* `is_alternative_plan_finding_present` - Indicates whether an alternative execution plan was reported for this SQL statement.
	* `is_error_finding_present` - Indicates whether there is an error in this SQL statement.
	* `is_index_finding_present` - Indicates whether an index recommendation was reported for this SQL statement.
	* `is_miscellaneous_finding_present` - Indicates whether a miscellaneous finding was reported for this SQL statement.
	* `is_restructure_sql_finding_present` - Indicates whether a restructure SQL recommendation was reported for this SQL statement.
	* `is_sql_profile_finding_implemented` - Indicates whether a SQL Profile recommendation has been implemented for this SQL statement.
	* `is_sql_profile_finding_present` - Indicates whether a SQL Profile recommendation was reported for this SQL statement.
	* `is_stats_finding_present` - Indicates whether a statistics recommendation was reported for this SQL statement.
	* `is_timeout_finding_present` - Indicates whether the task timed out.
	* `parsing_schema` - The parsing schema of the object.
	* `per_execution_percentage` - The per-execution percentage benefit.
	* `sql_key` - The unique key of this SQL statement.
	* `sql_text` - The text of the SQL statement.
	* `sql_tuning_advisor_task_id` - The unique identifier of the SQL Tuning Advisor task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_execution_id` - The execution id of the analyzed SQL object. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_id` - The key of the object to which these recommendations apply. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

