---
subcategory: "Fusion Apps"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fusion_apps_fusion_environment_scheduled_activities"
sidebar_current: "docs-oci-datasource-fusion_apps-fusion_environment_scheduled_activities"
description: |-
  Provides the list of Fusion Environment Scheduled Activities in Oracle Cloud Infrastructure Fusion Apps service
---

# Data Source: oci_fusion_apps_fusion_environment_scheduled_activities
This data source provides the list of Fusion Environment Scheduled Activities in Oracle Cloud Infrastructure Fusion Apps service.

Returns a list of ScheduledActivities.


## Example Usage

```hcl
data "oci_fusion_apps_fusion_environment_scheduled_activities" "test_fusion_environment_scheduled_activities" {
	#Required
	fusion_environment_id = oci_fusion_apps_fusion_environment.test_fusion_environment.id

	#Optional
	display_name = var.fusion_environment_scheduled_activity_display_name
	run_cycle = var.fusion_environment_scheduled_activity_run_cycle
	state = var.fusion_environment_scheduled_activity_state
	time_expected_finish_less_than_or_equal_to = var.fusion_environment_scheduled_activity_time_expected_finish_less_than_or_equal_to
	time_scheduled_start_greater_than_or_equal_to = var.fusion_environment_scheduled_activity_time_scheduled_start_greater_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fusion_environment_id` - (Required) unique FusionEnvironment identifier
* `run_cycle` - (Optional) A filter that returns all resources that match the specified run cycle.
* `state` - (Optional) A filter that returns all resources that match the specified status
* `time_expected_finish_less_than_or_equal_to` - (Optional) A filter that returns all resources that end before this date
* `time_scheduled_start_greater_than_or_equal_to` - (Optional) A filter that returns all resources that are scheduled after this date


## Attributes Reference

The following attributes are exported:

* `scheduled_activity_collection` - The list of scheduled_activity_collection.

### FusionEnvironmentScheduledActivity Reference

The following attributes are exported:

* `actions` - List of actions
	* `action_type` - Type of action
	* `artifact` - patch that delivered the vertex update prerequisite
	* `category` - patch artifact category
	* `description` - A string that describes the details of the action. It does not have to be unique, and you can change it. Avoid entering confidential information.
	* `mode` - A string that describeds whether the change is applied hot or cold
	* `qualifier` - month qualifier
	* `reference_key` - Unique identifier of the object that represents the action
	* `state` - A string that describes whether the change is applied hot or cold
	* `version` - name of the repo
* `delay_in_hours` - Cumulative delay hours
* `display_name` - scheduled activity display name, can be renamed.
* `fusion_environment_id` - FAaaS Environment Identifier.
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `run_cycle` - run cadence.
* `service_availability` - Service availability / impact during scheduled activity execution up down
* `state` - The current state of the scheduledActivity.
* `time_created` - The time the scheduled activity record was created. An RFC3339 formatted datetime string.
* `time_expected_finish` - Current time the scheduled activity is scheduled to end. An RFC3339 formatted datetime string.
* `time_finished` - The time the scheduled activity actually completed / cancelled / failed. An RFC3339 formatted datetime string.
* `time_scheduled_start` - Current time the scheduled activity is scheduled to start. An RFC3339 formatted datetime string.
* `time_updated` - The time the scheduled activity record was updated. An RFC3339 formatted datetime string.

