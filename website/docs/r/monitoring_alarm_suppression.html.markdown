---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarm_suppression"
sidebar_current: "docs-oci-resource-monitoring-alarm_suppression"
description: |-
  Provides the Alarm Suppression resource in Oracle Cloud Infrastructure Monitoring service
---

# oci_monitoring_alarm_suppression
This resource provides the Alarm Suppression resource in Oracle Cloud Infrastructure Monitoring service.

Creates a new alarm suppression at the specified level (alarm-wide or dimension-specific).
For more information, see
[Adding an Alarm-wide Suppression](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/add-alarm-suppression.htm) and
[Adding a Dimension-Specific Alarm Suppression](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/create-alarm-suppression.htm).

For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
resource "oci_monitoring_alarm_suppression" "test_alarm_suppression" {
	#Required
	alarm_suppression_target {
		#Required
		target_type = var.alarm_suppression_alarm_suppression_target_target_type

		#Optional
		alarm_id = oci_monitoring_alarm.test_alarm.id
		compartment_id = var.compartment_id
		compartment_id_in_subtree = var.alarm_suppression_alarm_suppression_target_compartment_id_in_subtree
	}
	display_name = var.alarm_suppression_display_name
	time_suppress_from = var.alarm_suppression_time_suppress_from
	time_suppress_until = var.alarm_suppression_time_suppress_until

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.alarm_suppression_description
	dimensions = var.alarm_suppression_dimensions
	freeform_tags = {"Department"= "Finance"}
	level = var.alarm_suppression_level
	suppression_conditions {
		#Required
		condition_type = var.alarm_suppression_suppression_conditions_condition_type
		suppression_duration = var.alarm_suppression_suppression_conditions_suppression_duration
		suppression_recurrence = var.alarm_suppression_suppression_conditions_suppression_recurrence
	}
}
```

## Argument Reference

The following arguments are supported:

* `alarm_suppression_target` - (Required) The target of the alarm suppression.
	* `alarm_id` - (Required when target_type=ALARM) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
	* `compartment_id` - (Required when target_type=COMPARTMENT) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment or tenancy that is the  target of the alarm suppression. Example: `ocid1.compartment.oc1..exampleuniqueID` 
	* `compartment_id_in_subtree` - (Applicable when target_type=COMPARTMENT) When true, the alarm suppression targets all alarms under all compartments and subcompartments of  the tenancy specified. The parameter can only be set to true when compartmentId is the tenancy OCID  (the tenancy is the root compartment). When false, the alarm suppression targets only the alarms under the specified compartment. 
	* `target_type` - (Required) The type of the alarm suppression target.
* `defined_tags` - (Optional) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) Human-readable reason for this alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	Oracle recommends including tracking information for the event or associated work, such as a ticket number.

	Example: `Planned outage due to change IT-1234.` 

* `dimensions` - (Required) A filter to suppress only alarm state entries that include the set of specified dimension key-value pairs. If you specify {"availabilityDomain": "phx-ad-1"} and the alarm state entry corresponds to the set {"availabilityDomain": "phx-ad-1" and "resourceId": "instance.region1.phx.exampleuniqueID"}, then this alarm will be included for suppression.

	This is required only when the value of level is `DIMENSION`. If required, the value cannot be an empty object. Only a single value is allowed per key. No grouping of multiple values is allowed under the same key. Maximum characters (after serialization): 4000. This maximum satisfies typical use cases. The response for an exceeded maximum is `HTTP 400` with an "dimensions values are too long" message. 
* `display_name` - (Required) A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `level` - (Optional) The level of this alarm suppression. `ALARM` indicates a suppression of the entire alarm, regardless of dimension. `DIMENSION` indicates a suppression configured for specified dimensions.

	Defaut: `DIMENSION` 
* `suppression_conditions` - (Optional) Array of all preconditions for alarm suppression. Example: `[{ conditionType: "RECURRENCE", suppressionRecurrence: "FRQ=DAILY;BYHOUR=10", suppressionDuration: "PT1H" }]` 
	* `condition_type` - (Required) Type of suppression condition.
	* `suppression_duration` - (Required) Duration of the recurring suppression. Specified as a string in ISO 8601 format. Minimum: `PT1M` (1 minute). Maximum: `PT24H` (24 hours). 
	* `suppression_recurrence` - (Required) Frequency and start time of the recurring suppression. The format follows [the iCalendar specification (RFC 5545, section 3.3.10)](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). Supported rule parts:
		* `FREQ`: Frequency of the recurring suppression: `WEEKLY` or `DAILY` only.
		* `BYDAY`: Comma separated days. Use with weekly suppressions only. Supported values: `MO`, `TU`, `WE`, `TH`, `FR`, `SA` ,`SU`.
		* `BYHOUR`, `BYMINUTE`, `BYSECOND`: Start time in UTC, after `timeSuppressFrom` value. Default is 00:00:00 UTC after `timeSuppressFrom`. 
* `time_suppress_from` - (Required) The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
* `time_suppress_until` - (Required) The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T02:02:29.600Z` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `alarm_suppression_target` - The target of the alarm suppression.
	* `alarm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment or tenancy that is the  target of the alarm suppression. Example: `ocid1.compartment.oc1..exampleuniqueID` 
	* `compartment_id_in_subtree` - When true, the alarm suppression targets all alarms under all compartments and subcompartments of  the tenancy specified. The parameter can only be set to true when compartmentId is the tenancy OCID  (the tenancy is the root compartment). When false, the alarm suppression targets only the alarms under the specified compartment. 
	* `target_type` - The type of the alarm suppression target.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm suppression.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `description` - Human-readable reason for this alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	Oracle recommends including tracking information for the event or associated work, such as a ticket number.

	Example: `Planned outage due to change IT-1234.` 
* `dimensions` - Configured dimension filter for suppressing alarm state entries that include the set of specified dimension key-value pairs.  Example: `{"resourceId": "instance.region1.phx.exampleuniqueID"}` 
* `display_name` - A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm suppression.
* `level` - The level of this alarm suppression. `ALARM` indicates a suppression of the entire alarm, regardless of dimension. `DIMENSION` indicates a suppression configured for specified dimensions. 
* `state` - The current lifecycle state of the alarm suppression.  Example: `DELETED` 
* `suppression_conditions` - Array of all preconditions for alarm suppression. Example: `[{ conditionType: "RECURRENCE", suppressionRecurrence: "FRQ=DAILY;BYHOUR=10", suppressionDuration: "PT1H" }]` 
	* `condition_type` - Type of suppression condition.
	* `suppression_duration` - Duration of the recurring suppression. Specified as a string in ISO 8601 format. Minimum: `PT1M` (1 minute). Maximum: `PT24H` (24 hours). 
	* `suppression_recurrence` - Frequency and start time of the recurring suppression. The format follows [the iCalendar specification (RFC 5545, section 3.3.10)](https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10). Supported rule parts:
		* `FREQ`: Frequency of the recurring suppression: `WEEKLY` or `DAILY` only.
		* `BYDAY`: Comma separated days. Use with weekly suppressions only. Supported values: `MO`, `TU`, `WE`, `TH`, `FR`, `SA` ,`SU`.
		* `BYHOUR`, `BYMINUTE`, `BYSECOND`: Start time in UTC, after `timeSuppressFrom` value. Default is 00:00:00 UTC after `timeSuppressFrom`. 
* `time_created` - The date and time the alarm suppression was created. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T02:02:29.600Z` 
* `time_updated` - The date and time the alarm suppression was last updated (deleted). Format defined by RFC3339.  Example: `2018-02-03T01:02:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Alarm Suppression
	* `update` - (Defaults to 20 minutes), when updating the Alarm Suppression
	* `delete` - (Defaults to 20 minutes), when destroying the Alarm Suppression


## Import

AlarmSuppressions can be imported using the `id`, e.g.

```
$ terraform import oci_monitoring_alarm_suppression.test_alarm_suppression "id"
```

