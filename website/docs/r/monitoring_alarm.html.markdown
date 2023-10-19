---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarm"
sidebar_current: "docs-oci-resource-monitoring-alarm"
description: |-
  Provides the Alarm resource in Oracle Cloud Infrastructure Monitoring service
---

# oci_monitoring_alarm
This resource provides the Alarm resource in Oracle Cloud Infrastructure Monitoring service.

Creates a new alarm in the specified compartment.
For more information, see
[Creating an Alarm](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/create-alarm.htm).
For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
resource "oci_monitoring_alarm" "test_alarm" {
	#Required
	compartment_id = var.compartment_id
	destinations = [oci_ons_notification_topic.test_notification_topic.id]
	display_name = var.alarm_display_name
	is_enabled = var.alarm_is_enabled
	metric_compartment_id = var.alarm_metric_compartment_id
	namespace = var.alarm_namespace
	query = var.alarm_query
	severity = var.alarm_severity

	#Optional
	body = var.alarm_body
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	is_notifications_per_metric_dimension_enabled = var.alarm_is_notifications_per_metric_dimension_enabled
	message_format = var.alarm_message_format
	metric_compartment_id_in_subtree = var.alarm_metric_compartment_id_in_subtree
	pending_duration = var.alarm_pending_duration
	repeat_notification_duration = var.alarm_repeat_notification_duration
	resolution = var.alarm_resolution
	resource_group = var.alarm_resource_group
	suppression {
		#Required
		time_suppress_from = var.alarm_suppression_time_suppress_from
		time_suppress_until = var.alarm_suppression_time_suppress_until

		#Optional
		description = var.alarm_suppression_description
	}
}
```

## Argument Reference

The following arguments are supported:

* `body` - (Optional) (Updatable) The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.  Example: `High CPU usage alert. Follow runbook instructions for resolution.` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm. 
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `destinations` - (Required) (Updatable) A list of destinations for alarm notifications. Each destination is represented by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a related resource, such as a [topic](https://docs.cloud.oracle.com/iaas/api/#/en/notification/latest/NotificationTopic). Supported destination services: Notifications , Streaming.           Limit: One destination per supported destination service. 
* `display_name` - (Required) (Updatable) A user-friendly name for the alarm. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	This value determines the title of each alarm notification.

	Example: `High CPU Utilization` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `is_enabled` - (Required) (Updatable) Whether the alarm is enabled.  Example: `true` 
* `is_notifications_per_metric_dimension_enabled` - (Optional) (Updatable) When set to `true`, splits alarm notifications per metric stream. When set to `false`, groups alarm notifications across metric streams. Example: `true` 
* `message_format` - (Optional) (Updatable) The format to use for alarm notifications. The formats are:
	* `RAW` - Raw JSON blob. Default value. When the `destinations` attribute specifies `Streaming`, all alarm notifications use this format.
	* `PRETTY_JSON`: JSON with new lines and indents. Available when the `destinations` attribute specifies `Notifications` only.
	* `ONS_OPTIMIZED`: Simplified, user-friendly layout. Available when the `destinations` attribute specifies `Notifications` only. Applies to Email subscription types only. 
* `metric_compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric being evaluated by the alarm. 
* `metric_compartment_id_in_subtree` - (Optional) (Updatable) When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified in metricCompartmentId. Default is false.  Example: `true` 
* `namespace` - (Required) (Updatable) The source service or application emitting the metric that is evaluated by the alarm.  Example: `oci_computeagent` 
* `pending_duration` - (Optional) (Updatable) The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING". For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING".

	The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.

	Under the default value of PT1M, the first evaluation that breaches the alarm updates the state to "FIRING".

	The alarm updates its status to "OK" when the breaching condition has been clear for the most recent minute.

	Example: `PT5M` 
* `query` - (Required) (Updatable) The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval depend on the specified time range. More interval values are supported for smaller time ranges. You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`. For information about writing MQL expressions, see [Editing the MQL Expression for a Query](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-mql.htm). For details about MQL, see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm). For available dimensions, review the metric definition for the supported service. See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

	Example of threshold alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

	-----

	Example of absence alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()

	----- 
* `repeat_notification_duration` - (Optional) (Updatable) The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours. Minimum: PT1M. Maximum: P30D.

	Default value: null (notifications are not re-submitted).

	Example: `PT2H` 
* `resolution` - (Optional) (Updatable) The time between calculated aggregation windows for the alarm. Supported value: `1m` 
* `resource_group` - (Optional) (Updatable) Resource group that you want to match. A null value returns only metric data that has no resource groups. The alarm retrieves metric data associated with the specified resource group only. Only one resource group can be applied per metric. A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($). Avoid entering confidential information.  Example: `frontend-fleet` 
* `severity` - (Required) (Updatable) The perceived type of response required when the alarm is in the "FIRING" state.  Example: `CRITICAL` 
* `suppression` - (Optional) (Updatable) The configuration details for suppressing an alarm. 
	* `description` - (Optional) (Updatable) Human-readable reason for suppressing alarm notifications. It does not have to be unique, and it's changeable. Avoid entering confidential information.

		Oracle recommends including tracking information for the event or associated work, such as a ticket number.

		Example: `Planned outage due to change IT-1234.` 
	* `time_suppress_from` - (Required) (Updatable) The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2019-02-01T01:02:29.600Z` 
	* `time_suppress_until` - (Required) (Updatable) The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2019-02-01T02:02:29.600Z` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `body` - The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices.  Example: `High CPU usage alert. Follow runbook instructions for resolution.` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm. 
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `destinations` - A list of destinations for alarm notifications. Each destination is represented by the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a related resource, such as a [topic](https://docs.cloud.oracle.com/iaas/api/#/en/notification/latest/NotificationTopic). Supported destination services: Notifications , Streaming.           Limit: One destination per supported destination service. 
* `display_name` - A user-friendly name for the alarm. It does not have to be unique, and it's changeable.

	This value determines the title of each alarm notification.

	Example: `High CPU Utilization` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm. 
* `is_enabled` - Whether the alarm is enabled.  Example: `true` 
* `is_notifications_per_metric_dimension_enabled` - When set to `true`, splits alarm notifications per metric stream. When set to `false`, groups alarm notifications across metric streams. 
* `message_format` - The format to use for alarm notifications. The formats are:
	* `RAW` - Raw JSON blob. Default value. When the `destinations` attribute specifies `Streaming`, all alarm notifications use this format.
	* `PRETTY_JSON`: JSON with new lines and indents. Available when the `destinations` attribute specifies `Notifications` only.
	* `ONS_OPTIMIZED`: Simplified, user-friendly layout. Available when the `destinations` attribute specifies `Notifications` only. Applies to Email subscription types only. 
* `metric_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric being evaluated by the alarm. 
* `metric_compartment_id_in_subtree` - When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified in metricCompartmentId. Default is false.  Example: `true` 
* `namespace` - The source service or application emitting the metric that is evaluated by the alarm.  Example: `oci_computeagent` 
* `pending_duration` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING". For example, a value of 5 minutes means that the alarm must persist in breaching the condition for five minutes before the alarm updates its state to "FIRING".

	The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H` for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.

	Under the default value of PT1M, the first evaluation that breaches the alarm updates the state to "FIRING".

	The alarm updates its status to "OK" when the breaching condition has been clear for the most recent minute.

	Example: `PT5M` 
* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value represents true. A true value means that the trigger rule condition has been met. The query must specify a metric, statistic, interval, and trigger rule (threshold or absence). Supported values for interval depend on the specified time range. More interval values are supported for smaller time ranges. You can optionally specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`. For information about writing MQL expressions, see [Editing the MQL Expression for a Query](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-mql.htm). For details about MQL, see [Monitoring Query Language (MQL) Reference](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm). For available dimensions, review the metric definition for the supported service. See [Supported Services](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).

	Example of threshold alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85

	-----

	Example of absence alarm:

	-----

	CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()

	----- 
* `repeat_notification_duration` - The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours. Minimum: PT1M. Maximum: P30D.

	Default value: null (notifications are not re-submitted).

	Example: `PT2H` 
* `resolution` - The time between calculated aggregation windows for the alarm. Supported value: `1m` 
* `resource_group` - Resource group to match for metric data retrieved by the alarm. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric. A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).  Example: `frontend-fleet` 
* `severity` - The perceived type of response required when the alarm is in the "FIRING" state.  Example: `CRITICAL` 
* `state` - The current lifecycle state of the alarm.  Example: `DELETED` 
* `suppression` - The configuration details for suppressing an alarm. 
	* `description` - Human-readable reason for suppressing alarm notifications. It does not have to be unique, and it's changeable. Avoid entering confidential information.

		Oracle recommends including tracking information for the event or associated work, such as a ticket number.

		Example: `Planned outage due to change IT-1234.` 
	* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2019-02-01T01:02:29.600Z` 
	* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2019-02-01T02:02:29.600Z` 
* `time_created` - The date and time the alarm was created. Format defined by RFC3339.  Example: `2019-02-01T01:02:29.600Z` 
* `time_updated` - The date and time the alarm was last updated. Format defined by RFC3339.  Example: `2019-02-03T01:02:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Alarm
	* `update` - (Defaults to 20 minutes), when updating the Alarm
	* `delete` - (Defaults to 20 minutes), when destroying the Alarm


## Import

Alarms can be imported using the `id`, e.g.

```
$ terraform import oci_monitoring_alarm.test_alarm "id"
```

