---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarm_history_collection"
sidebar_current: "docs-oci-datasource-monitoring-alarm_history_collection"
description: |-
  Provides details about a specific Alarm History Collection in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_alarm_history_collection
This data source provides details about a specific Alarm History Collection resource in Oracle Cloud Infrastructure Monitoring service.

Get the history of the specified alarm.
For more information, see
[Getting History of an Alarm](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/get-alarm-history.htm).
For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
data "oci_monitoring_alarm_history_collection" "test_alarm_history_collection" {
	#Required
	alarm_id = oci_monitoring_alarm.test_alarm.id

	#Optional
	alarm_historytype = var.alarm_history_collection_alarm_historytype
	timestamp_greater_than_or_equal_to = var.alarm_history_collection_timestamp_greater_than_or_equal_to
	timestamp_less_than = var.alarm_history_collection_timestamp_less_than
}
```

## Argument Reference

The following arguments are supported:

* `alarm_historytype` - (Optional) The type of history entries to retrieve. State history (STATE_HISTORY), state transition history (STATE_TRANSITION_HISTORY), rule history (RULE_HISTORY) or rule transition history (RULE_TRANSITION_HISTORY). If not specified, entries of all types are retrieved.  Example: `STATE_HISTORY` 
* `alarm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of an alarm. 
* `timestamp_greater_than_or_equal_to` - (Optional) A filter to return only alarm history entries with timestamps occurring on or after the specified date and time. Format defined by RFC3339.  Example: `2023-01-01T01:00:00.789Z` 
* `timestamp_less_than` - (Optional) A filter to return only alarm history entries with timestamps occurring before the specified date and time. Format defined by RFC3339.  Example: `2023-01-02T01:00:00.789Z` 


## Attributes Reference

The following attributes are exported:

* `alarm_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm to retrieve history for. 
* `entries` - The set of history entries retrieved for the alarm. 
	* `summary` - Description for this alarm history entry.

		Example 1 - alarm state history entry: `The alarm state is FIRING`

		Example 2 - alarm state transition history entry: `State transitioned from OK to Firing` 
	* `timestamp` - Timestamp for this alarm history entry. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
	* `timestamp_triggered` - Timestamp for the transition of the alarm state. For example, the time when the alarm transitioned from OK to Firing. Available for state transition entries only. Note: A three-minute lag for this value accounts for any late-arriving metrics.  Example: `2023-02-01T0:59:00.789Z` 
* `is_enabled` - Whether the alarm is enabled.  Example: `true` 

