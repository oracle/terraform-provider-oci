---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_definition"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_definition"
description: |-
  Provides details about a specific Scheduler Definition in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_definition
This data source provides details about a specific Scheduler Definition resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a SchedulerDefinition by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_definition" "test_scheduler_definition" {
	#Required
	scheduler_definition_id = oci_fleet_apps_management_scheduler_definition.test_scheduler_definition.id
}
```

## Argument Reference

The following arguments are supported:

* `scheduler_definition_id` - (Required) unique SchedulerDefinition identifier


## Attributes Reference

The following attributes are exported:

* `action_group_types` - All ActionGroup Types part of the schedule.
* `action_groups` - Action Groups associated with the Schedule.
	* `application_type` - Application Type associated. Only applicable if type is ENVIRONMENT. 
	* `lifecycle_operation` - LifeCycle Operation
	* `product` - Product associated. Only applicable if type is PRODUCT. 
	* `resource_id` - Provide the ID of the resource; Ex- fleetId.
	* `runbook_id` - ID of the runbook
	* `subjects` - Provide subjects that need to be considered for the schedule.
	* `target_id` - Provide the target if schedule is created against the target
	* `type` - ActionGroup Type associated.
* `activity_initiation_cut_off` - Activity Initiation Cut Off
* `application_types` - All application types part of the schedule for ENVIRONMENT ActionGroup Type. 
* `compartment_id` - Tenancy OCID
* `count_of_affected_action_groups` - Count of Action Groups affected by the Schedule.
* `count_of_affected_resources` - Count of Resources affected by the Schedule
* `count_of_affected_targets` - Count of Targets affected by the Schedule
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `lifecycle_operations` - All LifeCycle Operations part of the schedule
* `products` - All products part of the schedule for PRODUCT ActionGroup Type.
* `resource_region` - Associated region
* `run_books` - Runbooks.
	* `id` - The ID of the Runbook
	* `input_parameters` - Input Parameters for the Task
		* `arguments` - Arguments for the Task
			* `name` - Name of the output variable
			* `value` - The task output
		* `step_name` - stepName for which the input parameters are provided
* `schedule` - Schedule Information.
	* `duration` - Duration if schedule type is Custom
	* `execution_startdate` - Start Date for the schedule. An RFC3339 formatted datetime string
	* `maintenance_window_id` - Provide MaintenanceWindowId if Schedule Type is Maintenance Window
	* `recurrences` - Recurrence rule specification if Schedule Type is Custom and Recurring
	* `type` - Schedule Type
* `state` - The current state of the SchedulerDefinition.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_of_next_run` - Scheduled date for the next run of the Job.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

