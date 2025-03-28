---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_definitions"
sidebar_current: "docs-oci-datasource-fleet_apps_management-scheduler_definitions"
description: |-
  Provides the list of Scheduler Definitions in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_scheduler_definitions
This data source provides the list of Scheduler Definitions in Oracle Cloud Infrastructure Fleet Apps Management service.

List all lifecycle management schedules in Fleet Application Management.


## Example Usage

```hcl
data "oci_fleet_apps_management_scheduler_definitions" "test_scheduler_definitions" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.scheduler_definition_display_name
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	id = var.scheduler_definition_id
	maintenance_window_id = oci_fleet_apps_management_maintenance_window.test_maintenance_window.id
	product = var.scheduler_definition_product
	runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
	state = var.scheduler_definition_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Optional) unique Fleet identifier
* `id` - (Optional) A filter to return only schedule definitions whose identifier matches the given identifier.
* `maintenance_window_id` - (Optional) A filter to return only schedule definitions whose associated maintenanceWindowId matches the given maintenanceWindowId.
* `product` - (Optional) A filter to return only dchedule definitions whose assocaited product matches the given product
* `runbook_id` - (Optional) A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
* `state` - (Optional) A filter to return only scheduleDefinitions whose lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `scheduler_definition_collection` - The list of scheduler_definition_collection.

### SchedulerDefinition Reference

The following attributes are exported:

* `action_group_types` - All ActionGroup Types that are part of the schedule.
* `action_groups` - Action Groups associated with the Schedule.
	* `application_type` - Application Type associated. Only applicable if type is ENVIRONMENT. 
	* `lifecycle_operation` - LifeCycle Operation
	* `product` - Product associated. Only applicable if type is PRODUCT. 
	* `resource_id` - Provide the ID of the resource. Example fleet ID.
	* `runbook_id` - ID of the runbook
	* `subjects` - Provide subjects that need to be considered for the schedule.
	* `target_id` - Provide the target if schedule is created against the target
	* `type` - ActionGroup Type associated.
* `activity_initiation_cut_off` - Activity Initiation Cut Off.
* `application_types` - All application types that are part of the schedule for ENVIRONMENT ActionGroup Type. 
* `compartment_id` - Tenancy OCID
* `count_of_affected_action_groups` - Count of Action Groups affected by the Schedule.
* `count_of_affected_resources` - Count of Resources affected by the Schedule.
* `count_of_affected_targets` - Count of Targets affected by the Schedule.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `lifecycle_operations` - All LifeCycle Operations that are part of the schedule.
* `products` - All products that are part of the schedule for PRODUCT ActionGroup Type.
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
* `time_of_next_run` - The scheduled date for the next run of the Job.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

