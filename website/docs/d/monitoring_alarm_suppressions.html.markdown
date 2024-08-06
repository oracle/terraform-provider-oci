---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarm_suppressions"
sidebar_current: "docs-oci-datasource-monitoring-alarm_suppressions"
description: |-
  Provides the list of Alarm Suppressions in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_alarm_suppressions
This data source provides the list of Alarm Suppressions in Oracle Cloud Infrastructure Monitoring service.

Lists alarm suppressions for the specified alarm. For more information, see
[Listing Alarm Suppressions](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm-suppression.htm).

For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
data "oci_monitoring_alarm_suppressions" "test_alarm_suppressions" {

	#Optional
	alarm_id = oci_monitoring_alarm.test_alarm.id
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.alarm_suppression_compartment_id_in_subtree
	display_name = var.alarm_suppression_display_name
	is_all_suppressions = var.alarm_suppression_is_all_suppressions
	level = var.alarm_suppression_level
	state = var.alarm_suppression_state
	target_type = var.alarm_suppression_target_type
}
```

## Argument Reference

The following arguments are supported:

* `alarm_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for searching.  Use the tenancy OCID to search in the root compartment.

	If targetType is not specified, searches all suppressions defined under the compartment.  If targetType is `COMPARTMENT`, searches suppressions in the specified compartment only.

	Example: `ocid1.compartment.oc1..exampleuniqueID` 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list an alarm suppression by name. Alternatively, when you know the alarm suppression OCID, use the GetAlarmSuppression operation. 
* `is_all_suppressions` - (Optional) Setting this parameter to true requires the query to specify the alarm (`alarmId`).

	When true, lists all alarm suppressions that affect the specified alarm, including suppressions that target the corresponding compartment or tenancy. When false, lists only the alarm suppressions that target the specified alarm.

	Default is false. 
* `level` - (Optional) The level of this alarm suppression. `ALARM` indicates a suppression of the entire alarm, regardless of dimension. `DIMENSION` indicates a suppression configured for specified dimensions. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly. When not specified, only resources in the ACTIVE lifecycle state are listed. 
* `target_type` - (Optional) The target type to use when listing alarm suppressions.     `ALARM` lists all suppression records for the specified alarm. `COMPARTMENT` lists all suppression records for the specified compartment or tenancy. 


## Attributes Reference

The following attributes are exported:

* `alarm_suppression_collection` - The list of alarm_suppression_collection.

### AlarmSuppression Reference

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

