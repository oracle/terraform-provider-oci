---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_task"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_task"
description: |-
  Provides details about a specific Managed Database Sql Tuning Advisor Task in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_task
This data source provides details about a specific Managed Database Sql Tuning Advisor Task resource in Oracle Cloud Infrastructure Database Management service.

Lists the SQL Tuning Advisor tasks for the specified Managed Database.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_task" "test_managed_database_sql_tuning_advisor_task" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id

	#Optional
	name = var.managed_database_sql_tuning_advisor_task_name
	status = var.managed_database_sql_tuning_advisor_task_status
	time_greater_than_or_equal_to = var.managed_database_sql_tuning_advisor_task_time_greater_than_or_equal_to
	time_less_than_or_equal_to = var.managed_database_sql_tuning_advisor_task_time_less_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) The optional query parameter to filter the SQL Tuning Advisor task list by name.
* `status` - (Optional) The optional query parameter to filter the SQL Tuning Advisor task list by status.
* `time_greater_than_or_equal_to` - (Optional) The optional greater than or equal to query parameter to filter the timestamp.
* `time_less_than_or_equal_to` - (Optional) The optional less than or equal to query parameter to filter the timestamp.


## Attributes Reference

The following attributes are exported:

* `items` - A list of SQL Tuning Advisor tasks.
	* `days_to_expire` - The number of days left before the task expires. If the value equals -1, then the task has no expiration time (UNLIMITED).
	* `description` - The description of the SQL Tuning Advisor task.
	* `instance_id` - The instance ID of the SQL Tuning Advisor task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `name` - The name of the SQL Tuning Advisor task.
	* `owner` - The owner of the SQL Tuning Advisor task.
	* `recommendation_count` - The number of recommendations provided for the SQL Tuning Advisor task.
	* `sql_tuning_advisor_task_id` - The unique identifier of the SQL Tuning Advisor task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `task_status` - The status of the SQL Tuning Advisor task.
	* `time_created` - The Creation date of the SQL Tuning Advisor task.
	* `time_execution_ended` - The end time of the task execution.
	* `time_execution_started` - The start time of the task execution.
	* `total_sql_statements` - The total number of SQL statements related to the SQL Tuning Advisor task.

