---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_scheduling_policy_scheduling_window"
sidebar_current: "docs-oci-datasource-database-scheduling_policy_scheduling_window"
description: |-
  Provides details about a specific Scheduling Policy Scheduling Window in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_scheduling_policy_scheduling_window
This data source provides details about a specific Scheduling Policy Scheduling Window resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Scheduling Window.


## Example Usage

```hcl
data "oci_database_scheduling_policy_scheduling_window" "test_scheduling_policy_scheduling_window" {
	#Required
	scheduling_policy_id = oci_database_scheduling_policy.test_scheduling_policy.id
	scheduling_window_id = oci_database_scheduling_window.test_scheduling_window.id
}
```

## Argument Reference

The following arguments are supported:

* `scheduling_policy_id` - (Required) The Scheduling Policy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `scheduling_window_id` - (Required) The Scheduling Window [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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

