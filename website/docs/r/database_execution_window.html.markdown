---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_execution_window"
sidebar_current: "docs-oci-resource-database-execution_window"
description: |-
  Provides the Execution Window resource in Oracle Cloud Infrastructure Database service
---

# oci_database_execution_window
This resource provides the Execution Window resource in Oracle Cloud Infrastructure Database service.

Creates an execution window resource.


## Example Usage

```hcl
resource "oci_database_execution_window" "test_execution_window" {
	#Required
	compartment_id = var.compartment_id
	execution_resource_id = oci_cloud_guard_resource.test_resource.id
	time_scheduled = var.execution_window_time_scheduled
	window_duration_in_mins = var.execution_window_window_duration_in_mins

	#Optional
	defined_tags = var.execution_window_defined_tags
	freeform_tags = {"Department"= "Finance"}
	is_enforced_duration = var.execution_window_is_enforced_duration
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `execution_resource_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution resource the execution window belongs to.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_enforced_duration` - (Optional) (Updatable) Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
* `time_scheduled` - (Required) (Updatable) The scheduled start date and time of the execution window.
* `window_duration_in_mins` - (Required) (Updatable) Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Description of the execution window.
* `display_name` - The user-friendly name for the execution window. The name does not need to be unique.
* `estimated_time_in_mins` - The estimated time of the execution window in minutes.
* `execution_resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution resource the execution window belongs to.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the execution window.
* `is_enforced_duration` - Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `lifecycle_substate` - The current sub-state of the execution window. Valid states are DURATION_EXCEEDED, MAINTENANCE_IN_PROGRESS and WAITING. 
* `state` - The current state of the Schedule Policy. Valid states are CREATED, SCHEDULED, IN_PROGRESS, FAILED, CANCELED, UPDATING, DELETED, SUCCEEDED and PARTIAL_SUCCESS. 
* `time_created` - The date and time the execution window was created.
* `time_ended` - The date and time that the execution window ended.
* `time_scheduled` - The scheduled start date and time of the execution window.
* `time_started` - The date and time that the execution window was started.
* `time_updated` - The last date and time that the execution window was updated.
* `total_time_taken_in_mins` - The total time taken by corresponding resource activity in minutes.
* `window_duration_in_mins` - Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes. 
* `window_type` - The execution window is of PLANNED or UNPLANNED type.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Execution Window
	* `update` - (Defaults to 20 minutes), when updating the Execution Window
	* `delete` - (Defaults to 20 minutes), when destroying the Execution Window


## Import

ExecutionWindows can be imported using the `id`, e.g.

```
$ terraform import oci_database_execution_window.test_execution_window "id"
```

