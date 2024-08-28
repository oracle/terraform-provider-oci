---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_maintenance_window"
sidebar_current: "docs-oci-resource-stack_monitoring-maintenance_window"
description: |-
  Provides the Maintenance Window resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_maintenance_window
This resource provides the Maintenance Window resource in Oracle Cloud Infrastructure Stack Monitoring service.

Creates a new Maintenance Window for the given resources. It will create also the 
Alarms Suppression for each alarm that the resource migth trigger.


## Example Usage

```hcl
resource "oci_stack_monitoring_maintenance_window" "test_maintenance_window" {
	#Required
	compartment_id = var.compartment_id
	name = var.maintenance_window_name
	resources {
		#Required
		resource_id = oci_cloud_guard_resource.test_resource.id

		#Optional
		are_members_included = var.maintenance_window_resources_are_members_included
	}
	schedule {
		#Required
		schedule_type = var.maintenance_window_schedule_schedule_type

		#Optional
		maintenance_window_duration = var.maintenance_window_schedule_maintenance_window_duration
		maintenance_window_recurrences = var.maintenance_window_schedule_maintenance_window_recurrences
		time_maintenance_window_end = var.maintenance_window_schedule_time_maintenance_window_end
		time_maintenance_window_start = var.maintenance_window_schedule_time_maintenance_window_start
	}

	#Optional
	description = var.maintenance_window_description
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `description` - (Optional) (Updatable) Maintenance Window description.
* `name` - (Required) Maintenance Window name.
* `resources` - (Required) (Updatable) List of resource Ids which are part of the Maintenance Window 
	* `are_members_included` - (Optional) (Updatable) Flag to indicate if the members of the resource has to be include in the Maintenance Window. 
	* `resource_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource part of the Maintenance window. 
* `schedule` - (Required) (Updatable) Schedule information of the Maintenance Window 
	* `maintenance_window_duration` - (Applicable when schedule_type=RECURRENT) (Updatable) Duration time of each recurrence of each Maintenance Window. It must be specified as a string in ISO 8601 extended format. 
	* `maintenance_window_recurrences` - (Required when schedule_type=RECURRENT) (Updatable) A RFC5545 formatted recurrence string which represents the Maintenance Window Recurrence. Please refer this for details:https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10 FREQ: Frequency of the Maintenance Window. The supported values are: DAILY and WEEKLY. BYDAY: Comma separated days for Weekly Maintenance Window. BYHOUR: Specifies the start hour of each recurrence after `timeMaintenanceWindowStart` value. BYMINUTE: Specifies the start minute of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 BYSECOND: Specifies the start second of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 Other Rules are not supported. 
	* `schedule_type` - (Required) (Updatable) Property to identify the type of the Maintenance Window. 
	* `time_maintenance_window_end` - (Optional) (Updatable) Start time of Maintenance window. A RFC3339 formatted datetime string 
	* `time_maintenance_window_start` - (Optional) (Updatable) Start time of Maintenance window. A RFC3339 formatted datetime string 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `description` - Maintenance Window description.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of maintenance window. 
* `lifecycle_details` - Lifecycle Details of the Maintenance Window.
* `name` - Maintenance Window name.
* `resources` - List of resource Ids which are part of the Maintenance Window 
	* `are_members_included` - Flag to indicate if the members of the resource has to be include in the Maintenance Window. 
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource part of the Maintenance window. 
* `resources_details` - List of resource details that are part of the Maintenance Window. 
	* `name` - Name of the monitored resource 
	* `number_of_members` - Number of members of the resource 
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource part of the Maintenance window. 
	* `type` - Type of the monitored resource 
* `schedule` - Schedule information of the Maintenance Window 
	* `maintenance_window_duration` - Duration time of each recurrence of each Maintenance Window. It must be specified as a string in ISO 8601 extended format. 
	* `maintenance_window_recurrences` - A RFC5545 formatted recurrence string which represents the Maintenance Window Recurrence. Please refer this for details:https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10 FREQ: Frequency of the Maintenance Window. The supported values are: DAILY and WEEKLY. BYDAY: Comma separated days for Weekly Maintenance Window. BYHOUR: Specifies the start hour of each recurrence after `timeMaintenanceWindowStart` value. BYMINUTE: Specifies the start minute of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 BYSECOND: Specifies the start second of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 Other Rules are not supported. 
	* `schedule_type` - Property to identify the type of the Maintenance Window. 
	* `time_maintenance_window_end` - Start time of Maintenance window. A RFC3339 formatted datetime string 
	* `time_maintenance_window_start` - Start time of Maintenance window. A RFC3339 formatted datetime string 
* `state` - Lifecycle state of the monitored resource.
* `time_created` - The time the the maintenance window was created. An RFC3339 formatted datetime string 
* `time_updated` - The time the the mainteance window was updated. An RFC3339 formatted datetime string 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Maintenance Window
	* `update` - (Defaults to 20 minutes), when updating the Maintenance Window
	* `delete` - (Defaults to 20 minutes), when destroying the Maintenance Window


## Import

MaintenanceWindows can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_maintenance_window.test_maintenance_window "id"
```

