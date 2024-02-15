---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_application_schedules"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_application_schedules"
description: |-
  Provides the list of Workspace Application Schedules in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_application_schedules
This data source provides the list of Workspace Application Schedules in Oracle Cloud Infrastructure Data Integration service.

Use this endpoint to list schedules.


## Example Usage

```hcl
data "oci_dataintegration_workspace_application_schedules" "test_workspace_application_schedules" {
	#Required
	application_key = var.workspace_application_schedule_application_key
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	identifier = var.workspace_application_schedule_identifier
	key = var.workspace_application_schedule_key
	name = var.workspace_application_schedule_name
	type = var.workspace_application_schedule_type
}
```

## Argument Reference

The following arguments are supported:

* `application_key` - (Required) The application key.
* `identifier` - (Optional) Used to filter by the identifier of the object.
* `key` - (Optional) Used to filter by the key of the object.
* `name` - (Optional) Used to filter by the name of the object.
* `type` - (Optional) Used to filter by the object type of the object. It can be suffixed with an optional filter operator InSubtree. If this operator is not specified, then exact match is considered. <br><br><B>Examples:</B><br> <ul> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=false</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=true</B> returns all objects of type data loader task</li> </ul>
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

The following attributes are exported:

* `schedule_summary_collection` - The list of schedule_summary_collection.

### WorkspaceApplicationSchedule Reference

The following attributes are exported:

* `description` - Detailed description for the object.
* `frequency_details` - The model that holds the frequency details.
	* `custom_expression` - This holds the complete cron expression for this schedule, for example, 10 0/5 * * * ? that fires every 5 minutes, at 10 seconds after the minute (i.e. 10:00:10 am, 10:05:10 am, etc.)
	* `day_of_week` - This holds the day of the week on which the schedule should be triggered.
	* `days` - A list of days of the month to be scheduled. i.e. excute every 2nd,3rd, 10th of the month.
	* `frequency` - the frequency of the schedule.
	* `interval` - This hold the repeatability aspect of a schedule. i.e. in a monhtly frequency, a task can be scheduled for every month, once in two months, once in tree months etc.
	* `model_type` - The type of the model
	* `time` - A model to hold time in hour:minute:second format.
		* `hour` - The hour value.
		* `minute` - The minute value.
		* `second` - The second value.
	* `week_of_month` - This holds the week of the month in which the schedule should be triggered.
* `identifier` - Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `is_daylight_adjustment_enabled` - A flag to indicate daylight saving.
* `key` - Generated key that can be used in API calls to identify schedule. On scenarios where reference to the schedule is needed, a value can be passed in create.
* `metadata` - A summary type containing information about the object including its key, name and when/who created/updated it.
	* `aggregator` - A summary type containing information about the object's aggregator including its type, key, name and description.
		* `description` - The description of the aggregator.
		* `identifier` - The identifier of the aggregator.
		* `key` - The key of the aggregator object.
		* `name` - The name of the aggregator.
		* `type` - The type of the aggregator.
	* `aggregator_key` - The owning object key for this object.
	* `count_statistics` - A count statistics.
		* `object_type_count_list` - The array of statistics.
			* `object_count` - The value for the count statistic object.
			* `object_type` - The type of object for the count statistic object.
	* `created_by` - The user that created the object.
	* `created_by_name` - The user that created the object.
	* `identifier_path` - The full path to identify this object.
	* `info_fields` - Information property fields.
	* `is_favorite` - Specifies whether this object is a favorite or not.
	* `labels` - Labels are keywords or tags that you can add to data assets, dataflows and so on. You can define your own labels and use them to categorize content.
	* `registry_version` - The registry version of the object.
	* `time_created` - The date and time that the object was created.
	* `time_updated` - The date and time that the object was updated.
	* `updated_by` - The user that updated the object.
	* `updated_by_name` - The user that updated the object.
* `model_type` - The type of the object.
* `model_version` - This is a version number that is used by the service to upgrade objects if needed through releases of the service.
* `name` - Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
* `parent_ref` - A reference to the object's parent.
	* `parent` - Key of the parent object.
	* `root_doc_id` - Key of the root document object.
* `timezone` - The timezone for the schedule.

