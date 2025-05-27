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

Returns a list of all the Schedule Definitions in the specified compartment.
The query parameter `compartmentId` is required unless the query parameter `id` is specified.


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
	runbook_version_name = oci_fleet_apps_management_runbook_version.test_runbook_version.name
	state = var.scheduler_definition_state
	time_scheduled_greater_than_or_equal_to = var.scheduler_definition_time_scheduled_greater_than_or_equal_to
	time_scheduled_less_than = var.scheduler_definition_time_scheduled_less_than
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Optional) unique Fleet identifier
* `id` - (Optional) Unique identifier or OCID for listing a single Schedule Definition by id. Either compartmentId or id must be provided. 
* `maintenance_window_id` - (Optional) A filter to return only schedule definitions whose associated maintenanceWindowId matches the given maintenanceWindowId.
* `product` - (Optional) A filter to return only dchedule definitions whose assocaited product matches the given product
* `runbook_id` - (Optional) A filter to return only schedule definitions whose associated runbookId matches the given runbookId.
* `runbook_version_name` - (Optional) RunbookVersion Name filter
* `state` - (Optional) A filter to return only scheduleDefinitions whose lifecycleState matches the given lifecycleState.
* `time_scheduled_greater_than_or_equal_to` - (Optional) Scheduled Time
* `time_scheduled_less_than` - (Optional) Scheduled Time


## Attributes Reference

The following attributes are exported:

* `scheduler_definition_collection` - The list of scheduler_definition_collection.

### SchedulerDefinition Reference

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

