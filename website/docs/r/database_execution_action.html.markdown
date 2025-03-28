---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_execution_action"
sidebar_current: "docs-oci-resource-database-execution_action"
description: |-
  Provides the Execution Action resource in Oracle Cloud Infrastructure Database service
---

# oci_database_execution_action
This resource provides the Execution Action resource in Oracle Cloud Infrastructure Database service.

Creates an execution action resource.


## Example Usage

```hcl
resource "oci_database_execution_action" "test_execution_action" {
	#Required
	action_type = var.execution_action_action_type
	execution_window_id = oci_database_execution_window.test_execution_window.id

	#Optional
	action_members {
		#Required
		member_id = oci_database_member.test_member.id
		member_order = var.execution_action_action_members_member_order

		#Optional
		estimated_time_in_mins = var.execution_action_action_members_estimated_time_in_mins
		status = var.execution_action_action_members_status
		total_time_taken_in_mins = var.execution_action_action_members_total_time_taken_in_mins
	}
	action_params = var.execution_action_action_params
	compartment_id = var.compartment_id
	defined_tags = var.execution_action_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `action_members` - (Optional) (Updatable) List of action members of this execution action.
	* `estimated_time_in_mins` - (Optional) (Updatable) The estimated time of the execution action member in minutes.
	* `member_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent resource the execution action belongs to.
	* `member_order` - (Required) (Updatable) The priority order of the execution action member.
	* `status` - (Optional) (Updatable) The current status of the execution action member. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, DURATION_EXCEEDED, RESCHEDULED and COMPLETED. enum:
		* SCHEDULED
		* IN_PROGRESS
		* FAILED
		* CANCELED
		* DURATION_EXCEEDED
		* RESCHEDULED
		* SUCCEEDED 
	* `total_time_taken_in_mins` - (Optional) (Updatable) The total time taken by corresponding resource activity in minutes.
* `action_params` - (Optional) (Updatable) Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
* `action_type` - (Required) The action type of the execution action being performed
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `execution_window_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution window resource the execution action belongs to.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action_members` - List of action members of this execution action.
	* `estimated_time_in_mins` - The estimated time of the execution action member in minutes.
	* `member_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent resource the execution action belongs to.
	* `member_order` - The priority order of the execution action member.
	* `status` - The current status of the execution action member. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, DURATION_EXCEEDED, RESCHEDULED and COMPLETED. enum:
		* SCHEDULED
		* IN_PROGRESS
		* FAILED
		* CANCELED
		* DURATION_EXCEEDED
		* RESCHEDULED
		* SUCCEEDED 
	* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.
* `action_params` - Map<ParamName, ParamValue> where a key value pair describes the specific action parameter. Example: `{"count": "3"}` 
* `action_type` - The action type of the execution action being performed
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Description of the execution action.
* `display_name` - The user-friendly name for the execution action. The name does not need to be unique.
* `estimated_time_in_mins` - The estimated time of the execution action in minutes.
* `execution_action_order` - The priority order of the execution action.
* `execution_window_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution window resource the execution action belongs to.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution action.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `lifecycle_substate` - The current sub-state of the execution action. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING. 
* `state` - The current state of the execution action. Valid states are SCHEDULED, IN_PROGRESS, FAILED, CANCELED, UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS. 
* `time_created` - The date and time the execution action was created.
* `time_updated` - The last date and time that the execution action was updated.
* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Execution Action
	* `update` - (Defaults to 20 minutes), when updating the Execution Action
	* `delete` - (Defaults to 20 minutes), when destroying the Execution Action


## Import

ExecutionActions can be imported using the `id`, e.g.

```
$ terraform import oci_database_execution_action.test_execution_action "id"
```

