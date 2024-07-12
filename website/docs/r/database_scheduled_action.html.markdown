---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduled_action"
sidebar_current: "docs-oci-resource-database-scheduled_action"
description: |-
  Provides the Scheduled Action resource in Oracle Cloud Infrastructure Database service
---

# oci_database_scheduled_action
This resource provides the Scheduled Action resource in Oracle Cloud Infrastructure Database service.

Creates a Scheduled Action resource.


## Example Usage

```hcl
resource "oci_database_scheduled_action" "test_scheduled_action" {
	#Required
	action_type = var.scheduled_action_action_type
	compartment_id = var.compartment_id
	scheduling_plan_id = oci_database_scheduling_plan.test_scheduling_plan.id
	scheduling_window_id = oci_database_scheduling_window.test_scheduling_window.id

	#Optional
	action_members {
		#Required
		member_id = oci_database_member.test_member.id
		member_order = var.scheduled_action_action_members_member_order

		#Optional
		estimated_time_in_mins = var.scheduled_action_action_members_estimated_time_in_mins
	}
	action_params = var.scheduled_action_action_params
	defined_tags = var.scheduled_action_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `action_members` - (Optional) (Updatable) The list of action members in a scheduled action.
	* `estimated_time_in_mins` - (Optional) (Updatable) The estimated time for the intended action member.
	* `member_id` - (Required) (Updatable) The ocid of the action member.
	* `member_order` - (Required) (Updatable) The order of the action member in a scheduled action.
* `action_params` - (Optional) (Updatable) Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
* `action_type` - (Required) The type of the scheduled action being performed
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `scheduling_plan_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
* `scheduling_window_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Window.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action_members` - The list of action members in a scheduled action.
	* `estimated_time_in_mins` - The estimated time for the intended action member.
	* `member_id` - The ocid of the action member.
	* `member_order` - The order of the action member in a scheduled action.
* `action_order` - The order of the scheduled action.
* `action_params` - Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
* `action_type` - The type of the scheduled action being performed
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The display name of the Scheduled Action.
* `estimated_time_in_mins` - The estimated patching time for the scheduled action.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduled Action.
* `scheduling_plan_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Plan.
* `scheduling_window_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
* `state` - The current state of the Scheduled Action. Valid states are CREATING, NEEDS_ATTENTION, AVAILABLE, UPDATING, FAILED, DELETING and DELETED. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the Scheduled Action Resource was created.
* `time_updated` - The date and time the Scheduled Action Resource was updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduled Action
	* `update` - (Defaults to 20 minutes), when updating the Scheduled Action
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduled Action


## Import

ScheduledActions can be imported using the `id`, e.g.

```
$ terraform import oci_database_scheduled_action.test_scheduled_action "id"
```

