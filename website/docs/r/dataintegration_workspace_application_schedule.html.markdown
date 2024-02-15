---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_application_schedule"
sidebar_current: "docs-oci-resource-dataintegration-workspace_application_schedule"
description: |-
  Provides the Workspace Application Schedule resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_application_schedule
This resource provides the Workspace Application Schedule resource in Oracle Cloud Infrastructure Data Integration service.

Endpoint to create a new schedule

## Example Usage

```hcl
resource "oci_dataintegration_workspace_application_schedule" "test_workspace_application_schedule" {
	#Required
	application_key = var.workspace_application_schedule_application_key
	identifier = var.workspace_application_schedule_identifier
	name = var.workspace_application_schedule_name
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	description = var.workspace_application_schedule_description
	frequency_details {
		#Required
		model_type = var.workspace_application_schedule_frequency_details_model_type

		#Optional
		custom_expression = var.workspace_application_schedule_frequency_details_custom_expression
		day_of_week = var.workspace_application_schedule_frequency_details_day_of_week
		days = var.workspace_application_schedule_frequency_details_days
		frequency = var.workspace_application_schedule_frequency_details_frequency
		interval = var.workspace_application_schedule_frequency_details_interval
		time {

			#Optional
			hour = var.workspace_application_schedule_frequency_details_time_hour
			minute = var.workspace_application_schedule_frequency_details_time_minute
			second = var.workspace_application_schedule_frequency_details_time_second
		}
		week_of_month = var.workspace_application_schedule_frequency_details_week_of_month
	}
	is_daylight_adjustment_enabled = var.workspace_application_schedule_is_daylight_adjustment_enabled
	key = var.workspace_application_schedule_key
	model_version = var.workspace_application_schedule_model_version
	object_status = var.workspace_application_schedule_object_status
	object_version = var.workspace_application_schedule_object_version
	registry_metadata {

		#Optional
		aggregator_key = var.workspace_application_schedule_registry_metadata_aggregator_key
		is_favorite = var.workspace_application_schedule_registry_metadata_is_favorite
		key = var.workspace_application_schedule_registry_metadata_key
		labels = var.workspace_application_schedule_registry_metadata_labels
		registry_version = var.workspace_application_schedule_registry_metadata_registry_version
	}
	timezone = var.workspace_application_schedule_timezone
}
```

## Argument Reference

The following arguments are supported:

* `application_key` - (Required) The application key.
* `description` - (Optional) (Updatable) Detailed description for the object.
* `frequency_details` - (Optional) (Updatable) The model that holds the frequency details.
	* `custom_expression` - (Applicable when model_type=CUSTOM) (Updatable) This holds the complete cron expression for this schedule, for example, 10 0/5 * * * ? that fires every 5 minutes, at 10 seconds after the minute (i.e. 10:00:10 am, 10:05:10 am, etc.)
	* `day_of_week` - (Applicable when model_type=MONTHLY_RULE) (Updatable) This holds the day of the week on which the schedule should be triggered.
	* `days` - (Applicable when model_type=MONTHLY | WEEKLY) (Updatable) A list of days of the month to be scheduled. i.e. excute every 2nd,3rd, 10th of the month.
	* `frequency` - (Optional) (Updatable) the frequency of the schedule.
	* `interval` - (Applicable when model_type=DAILY | HOURLY | MONTHLY | MONTHLY_RULE) (Updatable) This hold the repeatability aspect of a schedule. i.e. in a monhtly frequency, a task can be scheduled for every month, once in two months, once in tree months etc.
	* `model_type` - (Required) (Updatable) The type of the model
	* `time` - (Applicable when model_type=DAILY | HOURLY | MONTHLY | MONTHLY_RULE | WEEKLY) (Updatable) A model to hold time in hour:minute:second format.
		* `hour` - (Applicable when model_type=DAILY | HOURLY | MONTHLY | MONTHLY_RULE | WEEKLY) (Updatable) The hour value.
		* `minute` - (Applicable when model_type=DAILY | HOURLY | MONTHLY | MONTHLY_RULE | WEEKLY) (Updatable) The minute value.
		* `second` - (Applicable when model_type=DAILY | HOURLY | MONTHLY | MONTHLY_RULE | WEEKLY) (Updatable) The second value.
	* `week_of_month` - (Applicable when model_type=MONTHLY_RULE) (Updatable) This holds the week of the month in which the schedule should be triggered.
* `identifier` - (Required) (Updatable) Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
* `is_daylight_adjustment_enabled` - (Optional) (Updatable) A flag to indicate whether daylight adjustment should be considered or not.
* `key` - (Optional) (Updatable) Generated key that can be used in API calls to identify schedule. On scenarios where reference to the schedule is needed, a value can be passed in create.
* `model_version` - (Optional) (Updatable) This is a version number that is used by the service to upgrade objects if needed through releases of the service.
* `name` - (Required) (Updatable) Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
* `object_status` - (Optional) (Updatable) The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
* `object_version` - (Optional) (Updatable) This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
* `registry_metadata` - (Optional) (Updatable) Information about the object and its parent.
	* `aggregator_key` - (Optional) (Updatable) The owning object's key for this object.
	* `is_favorite` - (Optional) (Updatable) Specifies whether this object is a favorite or not.
	* `key` - (Optional) (Updatable) The identifying key for the object.
	* `labels` - (Optional) (Updatable) Labels are keywords or labels that you can add to data assets, dataflows etc. You can define your own labels and use them to categorize content.
	* `registry_version` - (Optional) (Updatable) The registry version.
* `timezone` - (Optional) (Updatable) The timezone for the schedule.
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Application Schedule
	* `update` - (Defaults to 20 minutes), when updating the Workspace Application Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Application Schedule


## Import

WorkspaceApplicationSchedules can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_application_schedule.test_workspace_application_schedule "workspaces/{workspaceId}/applications/{applicationKey}/schedules/{scheduleKey}" 
```

