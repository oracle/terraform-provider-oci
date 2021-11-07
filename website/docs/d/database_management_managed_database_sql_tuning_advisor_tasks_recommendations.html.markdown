---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations"
sidebar_current: "docs-oci-datasource-database_management-managed_database_sql_tuning_advisor_tasks_recommendations"
description: |-
  Provides the list of Managed Database Sql Tuning Advisor Tasks Recommendations in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations
This data source provides the list of Managed Database Sql Tuning Advisor Tasks Recommendations in Oracle Cloud Infrastructure Database Management service.

Takes in a task id and object id and returns the recommendations/findings.


## Example Usage

```hcl
data "oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations" "test_managed_database_sql_tuning_advisor_tasks_recommendations" {
	#Required
	execution_id = oci_database_management_execution.test_execution.id
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	sql_object_id = oci_objectstorage_object.test_object.id
	sql_tuning_advisor_task_id = oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id
}
```

## Argument Reference

The following arguments are supported:

* `execution_id` - (Required) The execution id for an execution of a SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `sql_object_id` - (Required) The SQL object id for the SQL tuning task. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `sql_tuning_advisor_task_id` - (Required) The SQL tuning task identifier. This is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `sql_tuning_advisor_task_recommendation_collection` - The list of sql_tuning_advisor_task_recommendation_collection.

### ManagedDatabaseSqlTuningAdvisorTasksRecommendation Reference

The following attributes are exported:

* `items` - A list of SQL Tuning Advisor recommendations.
	* `benefit` - The percentage benefit of this implementation.
	* `finding` - Summary of the issue found for the SQL statement.
	* `implement_action_sql` - Action sql to be implemented based on the recommendation result.
	* `rationale` - Describes the reasoning behind the recommendation and how it relates to the finding.
	* `recommendation` - Particular recommendation for the finding.
	* `recommendation_key` - Unique identifier of the recommendation in the scope of the task.
	* `recommendation_type` - Type of recommendation
	* `sql_tuning_advisor_task_id` - Unique identifier of the task. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	* `sql_tuning_advisor_task_object_id` - Key of the object to which these recommendations apply. It is not the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 

