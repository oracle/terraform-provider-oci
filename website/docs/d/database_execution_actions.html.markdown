---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_execution_actions"
sidebar_current: "docs-oci-datasource-database-execution_actions"
description: |-
  Provides the list of Execution Actions in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_execution_actions
This data source provides the list of Execution Actions in Oracle Cloud Infrastructure Database service.

Lists the execution action resources in the specified compartment.


## Example Usage

```hcl
data "oci_database_execution_actions" "test_execution_actions" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.execution_action_display_name
	execution_window_id = oci_database_execution_window.test_execution_window.id
	state = var.execution_action_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `execution_window_id` - (Optional) A filter to return only resources that match the given execution wondow id.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `execution_actions` - The list of execution_actions.

### ExecutionAction Reference

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

