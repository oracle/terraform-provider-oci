---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduled_actions"
sidebar_current: "docs-oci-datasource-database-scheduled_actions"
description: |-
  Provides the list of Scheduled Actions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduled_actions
This data source provides the list of Scheduled Actions in Oracle Cloud Infrastructure Database service.

Lists the Scheduled Action resources in the specified compartment.


## Example Usage

```hcl
data "oci_database_scheduled_actions" "test_scheduled_actions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.scheduled_action_display_name
	id = var.scheduled_action_id
	scheduling_plan_id = oci_database_scheduling_plan.test_scheduling_plan.id
	service_type = var.scheduled_action_service_type
	state = var.scheduled_action_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `id` - (Optional) A filter to return only resources that match the given Scheduled Action id exactly.
* `scheduling_plan_id` - (Optional) A filter to return only resources that match the given scheduling policy id exactly.
* `service_type` - (Optional) A filter to return only resources that match the given service type exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `scheduled_action_collection` - The list of scheduled_action_collection.

### ScheduledAction Reference

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

