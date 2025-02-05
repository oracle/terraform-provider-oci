---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_execution_windows"
sidebar_current: "docs-oci-datasource-database-execution_windows"
description: |-
  Provides the list of Execution Windows in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_execution_windows
This data source provides the list of Execution Windows in Oracle Cloud Infrastructure Database service.

Lists the execution window resources in the specified compartment.


## Example Usage

```hcl
data "oci_database_execution_windows" "test_execution_windows" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.execution_window_display_name
	execution_resource_id = oci_cloud_guard_resource.test_resource.id
	state = var.execution_window_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `execution_resource_id` - (Optional) A filter to return only resources that match the given resource id exactly.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `execution_windows` - The list of execution_windows.

### ExecutionWindow Reference

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

