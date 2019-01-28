---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarms"
sidebar_current: "docs-oci-datasource-monitoring-alarms"
description: |-
  Provides the list of Alarms in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_alarms
This data source provides the list of Alarms in Oracle Cloud Infrastructure Monitoring service.

Lists the alarms for the specified compartment.


## Example Usage

```hcl
data "oci_monitoring_alarms" "test_alarms" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	compartment_id_in_subtree = "${var.alarm_compartment_id_in_subtree}"
	display_name = "${var.alarm_display_name}"
	state = "${var.alarm_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment. 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list an alarm by name. Alternatively, when you know the alarm OCID, use the GetAlarm operation. 
* `state` - (Optional) A filter to return only alarms that match the given lifecycle state exactly. When not specified, only alarms in the ACTIVE lifecycle state are listed. 


## Attributes Reference

The following attributes are exported:

* `alarms` - The list of alarms.

### Alarm Reference

The following attributes are exported:

* `body` - The human-readable content of the notification delivered. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.  Example: `High CPU usage alert. Follow runbook instructions for resolution.` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `destinations` - An array of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to which the notifications for this alarm will be delivered. An example destination is an OCID for a topic managed by the Oracle Cloud Infrastructure Notification service. 
* `display_name` - A user-friendly name for the alarm. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	This name is sent as the title for notifications related to this alarm.

	Example: `High CPU Utilization` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm. 
* `is_enabled` - Whether the alarm is enabled.  Example: `true` 
* `metric_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric being evaluated by the alarm. 
* `metric_compartment_id_in_subtree` - When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified in metricCompartmentId. Default is false.  Example: `true` 
* `namespace` - The source service or application emitting the metric that is evaluated by the alarm.  Example: `oci_computeagent` 
* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state  changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the  alarm must persist in breaching the condition for five minutes before the alarm updates its  state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five  minutes before the alarm updates its state to "OK."

	The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.

	Under the default value of PT1M, the first evaluation that breaches the alarm updates the state to "FIRING" and the first evaluation that does not breach the alarm updates the state to "OK".

	Example: `PT5M` 
* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of  the Monitoring service interprets results for each returned time series as Boolean values,  where zero represents false and a non-zero value represents true. A true value means that the trigger  rule condition has been met. The query must specify a metric, statistic, interval, and trigger  rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally  specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.  For details about Monitoring Query Language (MQL), see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm).  For available dimensions, review the metric definition for the supported service.  See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

	Example of threshold alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

	-----

	Example of absence alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()

	----- 
* `repeat_notification_duration` - The frequency at which notifications are re-submitted, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours. Minimum: PT1M. Maximum: P30D.

	Default value: null (notifications are not re-submitted).

	Example: `PT2H` 
* `resolution` - The time between calculated aggregation windows for the alarm. Supported value: `1m` 
* `severity` - The perceived type of response required when the alarm is in the "FIRING" state.  Example: `CRITICAL` 
* `state` - The current lifecycle state of the alarm.  Example: `DELETED` 
* `suppression` - The configuration details for suppressing an alarm. 
	* `description` - Human-readable reason for suppressing alarm notifications. It does not have to be unique, and it's changeable. Avoid entering confidential information.

		Oracle recommends including tracking information for the event or associated work, such as a ticket number.

		Example: `Planned outage due to change IT-1234.` 
	* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
	* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2018-02-01T02:02:29.600Z` 
* `time_created` - The date and time the alarm was created. Format defined by RFC3339.  Example: `2018-02-01T01:02:29.600Z` 
* `time_updated` - The date and time the alarm was last updated. Format defined by RFC3339.  Example: `2018-02-03T01:02:29.600Z` 

