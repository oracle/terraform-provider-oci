---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_maintenance_window"
sidebar_current: "docs-oci-datasource-fleet_apps_management-maintenance_window"
description: |-
  Provides details about a specific Maintenance Window in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_maintenance_window
This data source provides details about a specific Maintenance Window resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a MaintenanceWindow by identifier

## Example Usage

```hcl
data "oci_fleet_apps_management_maintenance_window" "test_maintenance_window" {
	#Required
	maintenance_window_id = oci_fleet_apps_management_maintenance_window.test_maintenance_window.id
}
```

## Argument Reference

The following arguments are supported:

* `maintenance_window_id` - (Required) unique MaintenanceWindow identifier


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

