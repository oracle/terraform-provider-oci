---
subcategory: "Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_monitoring_alarm_statuses"
sidebar_current: "docs-oci-datasource-monitoring-alarm_statuses"
description: |-
  Provides the list of Alarm Statuses in Oracle Cloud Infrastructure Monitoring service
---

# Data Source: oci_monitoring_alarm_statuses
This data source provides the list of Alarm Statuses in Oracle Cloud Infrastructure Monitoring service.

List the status of each alarm in the specified compartment.
Status is collective, across all metric streams in the alarm.
To list alarm status for each metric stream, use [RetrieveDimensionStates](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/AlarmDimensionStatesCollection/RetrieveDimensionStates).
The alarm attribute `isNotificationsPerMetricDimensionEnabled` must be set to `true`.
For more information, see
[Listing Alarm Statuses](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/list-alarm-status.htm).
For important limits information, see
[Limits on Monitoring](https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#limits).

This call is subject to a Monitoring limit that applies to the total number of requests across all alarm operations.
Monitoring might throttle this call to reject an otherwise valid request when the total rate of alarm operations exceeds 10 requests,
or transactions, per second (TPS) for a given tenancy.


## Example Usage

```hcl
data "oci_monitoring_alarm_statuses" "test_alarm_statuses" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.alarm_status_compartment_id_in_subtree
	display_name = var.alarm_status_display_name
	entity_id = oci_monitoring_entity.test_entity.id
	resource_id = oci_monitoring_resource.test_resource.id
	service_name = oci_core_service.test_service.name
	status = var.alarm_status_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the resources monitored by the metric that you are searching for. Use tenancyId to search in the root compartment.  Example: `ocid1.compartment.oc1..exampleuniqueID` 
* `compartment_id_in_subtree` - (Optional) When true, returns resources from all compartments and subcompartments. The parameter can only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment). A true value requires the user to have tenancy-level permissions. If this requirement is not met, then the call is rejected. When false, returns resources from only the compartment specified in compartmentId. Default is false. 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. Use this filter to list an alarm by name. Alternatively, when you know the alarm OCID, use the GetAlarm operation. 
* `entity_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the entity monitored by the metric that you are searching for.  Example: `ocid1.instance.oc1.phx.exampleuniqueID` 
* `resource_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a resource that is monitored by the metric that you are searching for.  Example: `ocid1.instance.oc1.phx.exampleuniqueID` 
* `service_name` - (Optional) A filter to return only resources that match the given service name exactly. Use this filter to list all alarms containing metric streams that match the *exact* service-name dimension.  Example: `logging-analytics` 
* `status` - (Optional) The status of the metric stream to use for alarm filtering. For example, set `StatusQueryParam` to "FIRING" to filter results to metric streams of the alarm with that status. Default behaviour is to return alarms irrespective of metric streams' status.  Example: `FIRING` 


## Attributes Reference

The following attributes are exported:

* `alarm_statuses` - The list of alarm_statuses.

### AlarmStatus Reference

The following attributes are exported:

* `display_name` - The configured name of the alarm.  Example: `High CPU Utilization` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm. 
* `rule_name` - Identifier of the alarm's base values for alarm evaluation, for use when the alarm contains overrides.  A valid ruleName value starts with an alphabetic character and includes only alphanumeric characters, underscores and square brackets.  Minimum number of characters: 3. Default value is `BASE`. For information about alarm overrides, see [AlarmOverride](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/datatypes/AlarmOverride). 
* `severity` - The perceived type of response required when the alarm is in the "FIRING" state.  Example: `CRITICAL` 
* `status` - The status of this alarm. Status is collective, across all metric streams in the alarm. To list alarm status for each metric stream, use [RetrieveDimensionStates](https://docs.cloud.oracle.com/iaas/api/#/en/monitoring/latest/AlarmDimensionStatesCollection/RetrieveDimensionStates). The alarm attribute `isNotificationsPerMetricDimensionEnabled` must be set to `true`. Example: `FIRING` 
* `suppression` - The configuration details for suppressing an alarm. 
	* `description` - Human-readable reason for suppressing alarm notifications. It does not have to be unique, and it's changeable. Avoid entering confidential information.

		Oracle recommends including tracking information for the event or associated work, such as a ticket number.

		Example: `Planned outage due to change IT-1234.` 
	* `time_suppress_from` - The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
	* `time_suppress_until` - The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T02:02:29.600Z` 
* `timestamp_triggered` - Timestamp for the transition of the alarm state. For example, the time when the alarm transitioned from OK to Firing. Note: A three-minute lag for this value accounts for any late-arriving metrics.  Example: `2023-02-01T01:02:29.600Z` 

