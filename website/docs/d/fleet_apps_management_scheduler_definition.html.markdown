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

Get the details of a SchedulerDefinition that performs lifecycle management operations.

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

