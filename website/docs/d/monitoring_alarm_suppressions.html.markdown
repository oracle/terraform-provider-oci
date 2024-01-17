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

Lists alarm suppressions for the specified alarm.
Only dimension-level suppressions are listed. Alarm-level suppressions are not listed.

For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
data "oci_monitoring_alarm_suppressions" "test_alarm_suppressions" {
	#Required
	alarm_id = oci_monitoring_alarm.test_alarm.id

	#Optional
	display_name = var.alarm_suppression_display_name
	state = var.alarm_suppression_state
}
```

## Argument Reference

The following arguments are supported:

* `alarm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list a alarm suppression by name. Alternatively, when you know the alarm suppression OCID, use the GetAlarmSuppression operation. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly. When not specified, only resources in the ACTIVE lifecycle state are listed. 


## Attributes Reference

The following attributes are exported:

* `alarm_suppression_collection` - The list of alarm_suppression_collection.

### AlarmSuppression Reference

The following attributes are exported:

* `alarm_suppression_target` - The target of the alarm suppression.
	* `alarm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
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
* `state` - The current lifecycle state of the alarm suppression.  Example: `DELETED` 
* `time_created` - The date and time the alarm suppression was created. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T02:02:29.600Z` 
* `time_updated` - The date and time the alarm suppression was last updated (deleted). Format defined by RFC3339.  Example: `2018-02-03T01:02:29.600Z` 

