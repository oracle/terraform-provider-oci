---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_scheduler_definition"
sidebar_current: "docs-oci-resource-fleet_apps_management-scheduler_definition"
description: |-
  Provides the Scheduler Definition resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_scheduler_definition
This resource provides the Scheduler Definition resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new SchedulerDefinition.


## Example Usage

```hcl
resource "oci_fleet_apps_management_scheduler_definition" "test_scheduler_definition" {
	#Required
	action_groups {
		#Required
		resource_id = oci_cloud_guard_resource.test_resource.id
		runbook_id = oci_fleet_apps_management_runbook.test_runbook.id

		#Optional
		application_type = var.scheduler_definition_action_groups_application_type
		lifecycle_operation = var.scheduler_definition_action_groups_lifecycle_operation
		product = var.scheduler_definition_action_groups_product
		subjects = var.scheduler_definition_action_groups_subjects
		target_id = oci_cloud_guard_target.test_target.id
		type = var.scheduler_definition_action_groups_type
	}
	compartment_id = var.compartment_id
	schedule {
		#Required
		execution_startdate = var.scheduler_definition_schedule_execution_startdate
		type = var.scheduler_definition_schedule_type

		#Optional
		duration = var.scheduler_definition_schedule_duration
		maintenance_window_id = oci_fleet_apps_management_maintenance_window.test_maintenance_window.id
		recurrences = var.scheduler_definition_schedule_recurrences
	}

	#Optional
	activity_initiation_cut_off = var.scheduler_definition_activity_initiation_cut_off
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.scheduler_definition_description
	display_name = var.scheduler_definition_display_name
	freeform_tags = {"bar-key"= "value"}
	run_books {
		#Required
		id = var.scheduler_definition_run_books_id

		#Optional
		input_parameters {
			#Required
			step_name = var.scheduler_definition_run_books_input_parameters_step_name

			#Optional
			arguments {
				#Required
				name = var.scheduler_definition_run_books_input_parameters_arguments_name

				#Optional
				value = var.scheduler_definition_run_books_input_parameters_arguments_value
			}
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `action_groups` - (Required) (Updatable) Action Groups associated with the Schedule.
	* `application_type` - (Optional) (Updatable) Application Type associated. Only applicable if type is ENVIRONMENT. 
	* `lifecycle_operation` - (Optional) (Updatable) LifeCycle Operation
	* `product` - (Optional) (Updatable) Product associated. Only applicable if type is PRODUCT. 
	* `resource_id` - (Required) (Updatable) Provide the ID of the resource; Ex- fleetId.
	* `runbook_id` - (Required) (Updatable) ID of the runbook
	* `subjects` - (Optional) (Updatable) Provide subjects that need to be considered for the schedule.
	* `target_id` - (Optional) (Updatable) Provide the target if schedule is created against the target
	* `type` - (Optional) (Updatable) ActionGroup Type associated.
* `activity_initiation_cut_off` - (Optional) (Updatable) Activity Initiation Cut Off
* `compartment_id` - (Required) Tenancy OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `run_books` - (Optional) (Updatable) Runbooks.
	* `id` - (Required) (Updatable) The ID of the Runbook
	* `input_parameters` - (Optional) (Updatable) Input Parameters for the Task
		* `arguments` - (Optional) (Updatable) Arguments for the Task
			* `name` - (Required) (Updatable) Name of the output variable
			* `value` - (Optional) (Updatable) The task output
		* `step_name` - (Required) (Updatable) stepName for which the input parameters are provided
* `schedule` - (Required) (Updatable) Schedule Information.
	* `duration` - (Optional) (Updatable) Duration if schedule type is Custom
	* `execution_startdate` - (Required) (Updatable) Start Date for the schedule. An RFC3339 formatted datetime string
	* `maintenance_window_id` - (Optional) (Updatable) Provide MaintenanceWindowId if Schedule Type is Maintenance Window
	* `recurrences` - (Optional) (Updatable) Recurrence rule specification if Schedule Type is Custom and Recurring
	* `type` - (Required) (Updatable) Schedule Type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduler Definition
	* `update` - (Defaults to 20 minutes), when updating the Scheduler Definition
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduler Definition


## Import

SchedulerDefinitions can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_scheduler_definition.test_scheduler_definition "id"
```

