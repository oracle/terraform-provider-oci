---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_maintenance_window"
sidebar_current: "docs-oci-resource-fleet_apps_management-maintenance_window"
description: |-
  Provides the Maintenance Window resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_maintenance_window
This resource provides the Maintenance Window resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Creates a new MaintenanceWindow.


## Example Usage

```hcl
resource "oci_fleet_apps_management_maintenance_window" "test_maintenance_window" {
	#Required
	compartment_id = var.compartment_id
	duration = var.maintenance_window_duration

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.maintenance_window_description
	display_name = var.maintenance_window_display_name
	freeform_tags = {"bar-key"= "value"}
	is_outage = var.maintenance_window_is_outage
	is_recurring = var.maintenance_window_is_recurring
	maintenance_window_type = var.maintenance_window_maintenance_window_type
	recurrences = var.maintenance_window_recurrences
	task_initiation_cutoff = var.maintenance_window_task_initiation_cutoff
	time_schedule_start = var.maintenance_window_time_schedule_start
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `duration` - (Required) (Updatable) Duration if schedule type is Custom
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_outage` - (Optional) (Updatable) Does the maintenenace window cause outage?
* `is_recurring` - (Optional) (Updatable) Is this is a recurring maintenance window
* `maintenance_window_type` - (Optional) (Updatable) Type of maintenenace window
* `recurrences` - (Optional) (Updatable) Recurrence rule specification if recurring
* `task_initiation_cutoff` - (Optional) (Updatable) Task initiation cutoff
* `time_schedule_start` - (Optional) (Updatable) Start time of schedule


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description. To provide some insight about the resource. Avoid entering confidential information. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `duration` - Duration if schedule type is Custom
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource.
* `is_outage` - Does the maintenenace window cause outage?
* `is_recurring` - Is this is a recurring maintenance window
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `maintenance_window_type` - Type of the MaintenanceWindow.
* `recurrences` - Recurrence rule specification if recurring
* `resource_region` - Associated region
* `state` - The current state of the MaintenanceWindow.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `task_initiation_cutoff` - Task initiation cutoff
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_schedule_start` - Start time of schedule
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Window
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Window
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Window


## Import

MaintenanceWindows can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_maintenance_window.test_maintenance_window "id"
```

