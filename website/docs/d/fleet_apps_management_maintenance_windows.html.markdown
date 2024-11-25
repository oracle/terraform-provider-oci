---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_maintenance_windows"
sidebar_current: "docs-oci-datasource-fleet_apps_management-maintenance_windows"
description: |-
  Provides the list of Maintenance Windows in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_maintenance_windows
This data source provides the list of Maintenance Windows in Oracle Cloud Infrastructure Fleet Apps Management service.

List maintenance windows for a specified tenancy in Fleet Application Management.


## Example Usage

```hcl
data "oci_fleet_apps_management_maintenance_windows" "test_maintenance_windows" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.maintenance_window_display_name
	id = var.maintenance_window_id
	state = var.maintenance_window_state
	time_schedule_start_greater_than_or_equal_to = var.maintenance_window_time_schedule_start_greater_than_or_equal_to
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) A filter to return only the Maintenance Windows whose identifier matches the given identifier.
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState.
* `time_schedule_start_greater_than_or_equal_to` - (Optional) A filter to return only resources whose timeScheduleStart is greater than or equal to the provided date and time.


## Attributes Reference

The following attributes are exported:

* `maintenance_window_collection` - The list of maintenance_window_collection.

### MaintenanceWindow Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `duration` - Duration of the maintenance window. Specify how long the maintenance window remains open. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `is_outage` - Does the maintenenace window cause outage? An outage indicates whether a maintenance window can consider operations that require downtime. It means a period when the application is not accessible. 
* `is_recurring` - Is this a recurring maintenance window?
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `maintenance_window_type` - Type of the MaintenanceWindow.
* `recurrences` - Recurrence rule specification if maintenance window recurring. Specify the frequency of running the maintenance window. 
* `resource_region` - Associated region
* `state` - The current state of the MaintenanceWindow.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `task_initiation_cutoff` - Task initiation cutoff time for the maintenance window.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_schedule_start` - Specify the date and time of the day that the maintenance window starts.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

