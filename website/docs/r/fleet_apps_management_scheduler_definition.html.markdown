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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/SchedulerDefinition

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Create a SchedulerDefinition to perform lifecycle operations.


## Example Usage

```hcl
resource "oci_fleet_apps_management_scheduler_definition" "test_scheduler_definition" {
	#Required
	action_groups {
		#Required
		fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
		kind = var.scheduler_definition_action_groups_kind
		runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
		runbook_version_name = oci_fleet_apps_management_runbook_version.test_runbook_version.name

		#Optional
		display_name = var.scheduler_definition_action_groups_display_name
		sequence = var.scheduler_definition_action_groups_sequence
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
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.scheduler_definition_description
	display_name = var.scheduler_definition_display_name
	freeform_tags = {"bar-key"= "value"}
	run_books {
		#Required
		runbook_id = oci_fleet_apps_management_runbook.test_runbook.id
		runbook_version_name = oci_fleet_apps_management_runbook_version.test_runbook_version.name

		#Optional
		input_parameters {
			#Required
			step_name = var.scheduler_definition_run_books_input_parameters_step_name

			#Optional
			arguments {
				#Required
				kind = var.scheduler_definition_run_books_input_parameters_arguments_kind
				name = var.scheduler_definition_run_books_input_parameters_arguments_name

				#Optional
				content {
					#Required
					bucket = var.scheduler_definition_run_books_input_parameters_arguments_content_bucket
					checksum = var.scheduler_definition_run_books_input_parameters_arguments_content_checksum
					namespace = var.scheduler_definition_run_books_input_parameters_arguments_content_namespace
					object = var.scheduler_definition_run_books_input_parameters_arguments_content_object
					source_type = var.scheduler_definition_run_books_input_parameters_arguments_content_source_type
				}
				value = var.scheduler_definition_run_books_input_parameters_arguments_value
			}
		}
	}
}
```

## Argument Reference

The following arguments are supported:

* `action_groups` - (Required) (Updatable) Action Groups associated with the Schedule.
	* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_id` - (Required) (Updatable) ID of the fleet
	* `kind` - (Required) (Updatable) Action Group kind
	* `runbook_id` - (Required) (Updatable) ID of the runbook
	* `runbook_version_name` - (Required) (Updatable) Name of the runbook version
	* `sequence` - (Optional) (Updatable) Sequence of the Action Group. Action groups will be executed in a seuential order. All Action Groups having the same sequence will be executed parallely. If no value is provided a default value of 1 will be given. 
* `compartment_id` - (Required) Compartment OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `run_books` - (Optional) (Updatable) Runbooks.
	* `input_parameters` - (Optional) (Updatable) Input Parameters for the Task
		* `arguments` - (Optional) (Updatable) Arguments for the Task
			* `content` - (Applicable when kind=FILE) (Updatable) Content Source details.
				* `bucket` - (Required) (Updatable) Bucket Name.
				* `checksum` - (Required) (Updatable) md5 checksum of the artifact.
				* `namespace` - (Required) (Updatable) Namespace.
				* `object` - (Required) (Updatable) Object Name.
				* `source_type` - (Required) (Updatable) Content Source type details. 
			* `kind` - (Required) (Updatable) Task argument kind
			* `name` - (Required) (Updatable) Name of the input variable
			* `value` - (Applicable when kind=STRING) (Updatable) The task input
		* `step_name` - (Required) (Updatable) stepName for which the input parameters are provided
	* `runbook_id` - (Required) (Updatable) The ID of the Runbook
	* `runbook_version_name` - (Required) (Updatable) The runbook version name
* `schedule` - (Required) (Updatable) Schedule Information.
	* `duration` - (Required when type=CUSTOM) (Updatable) Duration of the schedule.
	* `execution_startdate` - (Required) (Updatable) Start Date for the schedule. An RFC3339 formatted datetime string
	* `maintenance_window_id` - (Required when type=MAINTENANCE_WINDOW) (Updatable) Provide MaintenanceWindowId
	* `recurrences` - (Applicable when type=CUSTOM) (Updatable) Recurrence rule specification if recurring
	* `type` - (Required) (Updatable) Schedule Type


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action_groups` - Action Groups associated with the Schedule.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
	* `fleet_id` - ID of the fleet
	* `kind` - Action Group kind
	* `runbook_id` - ID of the runbook
	* `runbook_version_name` - Name of the runbook version
	* `sequence` - Sequence of the Action Group. Action groups will be executed in a seuential order. All Action Groups having the same sequence will be executed parallely. If no value is provided a default value of 1 will be given. 
* `compartment_id` - Compartment OCID
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
	* `input_parameters` - Input Parameters for the Task
		* `arguments` - Arguments for the Task
			* `content` - Content Source details.
				* `bucket` - Bucket Name.
				* `checksum` - md5 checksum of the artifact.
				* `namespace` - Namespace.
				* `object` - Object Name.
				* `source_type` - Content Source type details. 
			* `kind` - Task argument kind
			* `name` - Name of the input variable
			* `value` - The task input
		* `step_name` - stepName for which the input parameters are provided
	* `runbook_id` - The ID of the Runbook
	* `runbook_version_name` - The runbook version name
* `schedule` - Schedule Information.
	* `duration` - Duration of the schedule.
	* `execution_startdate` - Start Date for the schedule. An RFC3339 formatted datetime string
	* `maintenance_window_id` - Provide MaintenanceWindowId
	* `recurrences` - Recurrence rule specification if recurring
	* `type` - Schedule Type
* `state` - The current state of the SchedulerDefinition.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_of_next_run` - The scheduled date for the next run of the Job.
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

