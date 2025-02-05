---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_execution_action"
sidebar_current: "docs-oci-datasource-database-execution_action"
description: |-
  Provides details about a specific Execution Action in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_execution_action
This data source provides details about a specific Execution Action resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified execution action.


## Example Usage

```hcl
data "oci_database_execution_action" "test_execution_action" {
	#Required
	execution_action_id = oci_database_execution_action.test_execution_action.id
}
```

## Argument Reference

The following arguments are supported:

* `execution_action_id` - (Required) The execution action [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

