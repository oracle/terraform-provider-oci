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

Creates a dimension-specific suppression for an alarm.

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
		alarm_id = oci_monitoring_alarm.test_alarm.id
		target_type = var.alarm_suppression_alarm_suppression_target_target_type
	}
	dimensions = var.alarm_suppression_dimensions
	display_name = var.alarm_suppression_display_name
	time_suppress_from = var.alarm_suppression_time_suppress_from
	time_suppress_until = var.alarm_suppression_time_suppress_until

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.alarm_suppression_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `alarm_suppression_target` - (Required) The target of the alarm suppression.
	* `alarm_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
	* `target_type` - (Required) The type of the alarm suppression target.
* `defined_tags` - (Optional) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) Human-readable reason for this alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.

	Oracle recommends including tracking information for the event or associated work, such as a ticket number.

	Example: `Planned outage due to change IT-1234.` 
* `dimensions` - (Required) A filter to suppress only alarm state entries that include the set of specified dimension key-value pairs. If you specify {"availabilityDomain": "phx-ad-1"} and the alarm state entry corresponds to the set {"availabilityDomain": "phx-ad-1" and "resourceId": "instance.region1.phx.exampleuniqueID"}, then this alarm will be included for suppression.

	The value cannot be an empty object. Only a single value is allowed per key. No grouping of multiple values is allowed under the same key. Maximum characters (after serialization): 4000. This maximum satisfies typical use cases. The response for an exceeded maximum is `HTTP 400` with an "dimensions values are too long" message. 
* `display_name` - (Required) A user-friendly name for the alarm suppression. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"Department": "Finance"}` 
* `time_suppress_from` - (Required) The start date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T01:02:29.600Z` 
* `time_suppress_until` - (Required) The end date and time for the suppression to take place, inclusive. Format defined by RFC3339.  Example: `2023-02-01T02:02:29.600Z` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

