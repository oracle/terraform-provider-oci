---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_recommendation"
description: |-
  Provides details about a specific Managed Database Sql Tuning Advisor Tasks Recommendation in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation
This data source provides details about a specific Managed Database Sql Tuning Advisor Tasks Recommendation resource in Oracle Cloud Infrastructure Database Management service.

Gets the findings and possible actions for a given object in a SQL tuning task.
The task ID and object ID are used to retrieve the findings and recommendations.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation" "test_managed_database_sql_tuning_advisor_tasks_recommendation" {
	#Required
	execution_id = oci_database_management_execution.test_execution.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_object_id = oci_objectstorage_object.test_object.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}
```

## Argument Reference

The following arguments are supported:

* `execution_id` - (Required) The execution ID for an execution of a SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `sql_object_id` - (Required) The SQL object ID for the SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `items` - A list of SQL Tuning Advisor recommendations.
	* `benefit` - The percentage benefit of this implementation.
	* `finding` - Summary of the issue found in the SQL statement.
	* `implement_action_sql` - Action sql to be implemented based on the recommendation result.
	* `rationale` - Describes the reasoning behind the recommendation and how it relates to the finding.
	* `recommendation` - The recommendation for a specific finding.
	* `recommendation_key` - The unique identifier of the recommendation in the scope of the task.
	* `recommendation_type` - Type of recommendation.
	* `sql_tuning_advisor_task_id` - The unique identifier of the task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_id` - The key of the object to which these recommendations apply. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

