---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_policy_scheduling_window"
sidebar_current: "docs-oci-resource-database-scheduling_policy_scheduling_window"
description: |-
  Provides the Scheduling Policy Scheduling Window resource in Oracle Cloud Infrastructure Database service
---

# oci_database_scheduling_policy_scheduling_window
This resource provides the Scheduling Policy Scheduling Window resource in Oracle Cloud Infrastructure Database service.

Creates a Scheduling Window resource.


## Example Usage

```hcl
resource "oci_database_scheduling_policy_scheduling_window" "test_scheduling_policy_scheduling_window" {
	#Required
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
	window_preference {
		#Required
		days_of_week {
			#Required
			name = var.scheduling_policy_scheduling_window_window_preference_days_of_week_name
		}
		duration = var.scheduling_policy_scheduling_window_window_preference_duration
		is_enforced_duration = var.scheduling_policy_scheduling_window_window_preference_is_enforced_duration
		start_time = var.scheduling_policy_scheduling_window_window_preference_start_time
		weeks_of_month = var.scheduling_policy_scheduling_window_window_preference_weeks_of_month

		#Optional
		months {
			#Required
			name = var.scheduling_policy_scheduling_window_window_preference_months_name
		}
	}

	#Optional
	compartment_id = var.compartment_id
	defined_tags = var.scheduling_policy_scheduling_window_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `scheduling_policy_id` - (Required) The Scheduling Policy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `window_preference` - (Required) (Updatable) The Single Scheduling Window details. 
	* `days_of_week` - (Required) (Updatable) Days during the week when scheduling window should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `duration` - (Required) (Updatable) Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes. 
	* `is_enforced_duration` - (Required) (Updatable) Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
	* `months` - (Optional) (Updatable) Months during the year when scheduled window should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `start_time` - (Required) (Updatable) The scheduling window start time. The value must use the ISO-8601 format "hh:mm".
	* `weeks_of_month` - (Required) (Updatable) Weeks during the month when scheduled window should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow scheduling window during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Scheduling window cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and startTime parameters to allow you to specify specific days of the week and hours that scheduled window will be performed. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Scheduling Window. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Window.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `scheduling_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Scheduling Policy.
* `state` - The current state of the Scheduling Window. Valid states are CREATING, ACTIVE, UPDATING, FAILED, DELETING and DELETED. 
* `time_created` - The date and time the Scheduling Window was created.
* `time_next_scheduling_window_starts` - The date and time of the next upcoming window associated within the schedulingWindow is planned to start.
* `time_updated` - The last date and time that the Scheduling Window was updated.
* `window_preference` - The Single Scheduling Window details. 
	* `days_of_week` - Days during the week when scheduling window should be performed.
		* `name` - Name of the day of the week.
	* `duration` - Duration window allows user to set a duration they plan to allocate for Scheduling window. The duration is in minutes. 
	* `is_enforced_duration` - Indicates if duration the user plans to allocate for scheduling window is strictly enforced. The default value is `FALSE`.
	* `months` - Months during the year when scheduled window should be performed.
		* `name` - Name of the month of the year.
	* `start_time` - The scheduling window start time. The value must use the ISO-8601 format "hh:mm".
	* `weeks_of_month` - Weeks during the month when scheduled window should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow scheduling window during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Scheduling window cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and startTime parameters to allow you to specify specific days of the week and hours that scheduled window will be performed. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Scheduling Policy Scheduling Window
	* `update` - (Defaults to 20 minutes), when updating the Scheduling Policy Scheduling Window
	* `delete` - (Defaults to 20 minutes), when destroying the Scheduling Policy Scheduling Window


## Import

SchedulingPolicySchedulingWindows can be imported using the `id`, e.g.

```
$ terraform import oci_database_scheduling_policy_scheduling_window.test_scheduling_policy_scheduling_window "schedulingPolicies/{schedulingPolicyId}/schedulingWindows/{schedulingWindowId}" 
```

