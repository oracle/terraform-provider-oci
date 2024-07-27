---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_maintenance_event"
sidebar_current: "docs-oci-resource-core-instance_maintenance_event"
description: |-
  Provides the Instance Maintenance Event resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_maintenance_event
This resource provides the Instance Maintenance Event resource in Oracle Cloud Infrastructure Core service.

Updates the maintenance event for the given instance.


## Example Usage

```hcl
resource "oci_core_instance_maintenance_event" "test_instance_maintenance_event" {
	#Required
	instance_maintenance_event_id = oci_core_instance_maintenance_event.test_instance_maintenance_event.id

	#Optional
	alternative_resolution_action = var.instance_maintenance_event_alternative_resolution_action
	can_delete_local_storage = var.instance_maintenance_event_can_delete_local_storage
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.instance_maintenance_event_display_name
	freeform_tags = {"Department"= "Finance"}
	time_window_start = var.instance_maintenance_event_time_window_start
}
```

## Argument Reference

The following arguments are supported:

* `alternative_resolution_action` - (Optional) (Updatable) One of the alternativeResolutionActions that was provided in the InstanceMaintenanceEvent. 
* `can_delete_local_storage` - (Optional) (Updatable) This field is only applicable when setting the alternativeResolutionAction.

	For Instances that have local storage, this must be set to true to verify that the local storage will be deleted during the migration. For instances without, this parameter has no effect.

	In cases where the local storage will be lost, this parameter must be set or the request will fail. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_maintenance_event_id` - (Required) The OCID of the instance maintenance event.
* `time_window_start` - (Optional) (Updatable) The beginning of the time window when Maintenance is scheduled to begin. The Maintenance will not begin before this time.

	The timeWindowEnd is automatically calculated based on the maintenanceReason and the instanceAction. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - Additional details of the maintenance in the form of json. 
* `alternative_resolution_actions` - These are alternative actions to the requested instanceAction that can be taken to resolve the Maintenance. 
* `can_delete_local_storage` - For Instances that have local storage, this field is set to true when local storage will be deleted as a result of the Maintenance. 
* `can_reschedule` - Indicates if this MaintenanceEvent is capable of being rescheduled up to the timeHardDueDate. 
* `compartment_id` - The OCID of the compartment that contains the instance. 
* `correlation_token` - A unique identifier that will group Instances that have a relationship with one another and must be scheduled together for the Maintenance to proceed. Any Instances that have a relationship with one another from a Maintenance perspective will have a matching correlationToken. 
* `created_by` - The creator of the maintenance event. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - It is the descriptive information about the maintenance taking place on the customer instance. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `estimated_duration` - This is the estimated duration of the Maintenance, once the Maintenance has entered the STARTED state. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance event. 
* `instance_action` - This is the action that will be performed on the Instance by Oracle Cloud Infrastructure when the Maintenance begins. 
* `instance_id` - The OCID of the instance.
* `maintenance_category` - This indicates the priority and allowed actions for this Maintenance. Higher priority forms of Maintenance have tighter restrictions and may not be rescheduled, while lower priority/severity Maintenance can be rescheduled, deferred, or even cancelled. Please see the [Instance Maintenance](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/placeholder.htm) documentation for details. 
* `maintenance_reason` - This is the reason that Maintenance is being performed. See [Instance Maintenance](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/placeholder.htm) documentation for details. 
* `start_window_duration` - The duration of the time window Maintenance is scheduled to begin within. 
* `state` - The current state of the maintenance event. 
* `time_created` - The date and time the maintenance event was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 
* `time_finished` - The time at which the Maintenance actually finished. 
* `time_hard_due_date` - It is the scheduled hard due date and time of the maintenance event. The maintenance event will happen at this time and the due date will not be extended. 
* `time_started` - The time at which the Maintenance actually started. 
* `time_window_start` - The beginning of the time window when Maintenance is scheduled to begin. The Maintenance will not begin before this time. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Instance Maintenance Event
	* `update` - (Defaults to 20 minutes), when updating the Instance Maintenance Event
	* `delete` - (Defaults to 20 minutes), when destroying the Instance Maintenance Event


## Import

InstanceMaintenanceEvents can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_maintenance_event.test_instance_maintenance_event "id"
```

