---
subcategory: "Compute Cloud At Customer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_compute_cloud_at_customer_ccc_upgrade_schedule"
sidebar_current: "docs-oci-resource-compute_cloud_at_customer-ccc_upgrade_schedule"
description: |-
  Provides the Ccc Upgrade Schedule resource in Oracle Cloud Infrastructure Compute Cloud At Customer service
---

# oci_compute_cloud_at_customer_ccc_upgrade_schedule
This resource provides the Ccc Upgrade Schedule resource in Oracle Cloud Infrastructure Compute Cloud At Customer service.

Creates a new Compute Cloud@Customer upgrade schedule.


## Example Usage

```hcl
resource "oci_compute_cloud_at_customer_ccc_upgrade_schedule" "test_ccc_upgrade_schedule" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.ccc_upgrade_schedule_display_name
	events {
		#Required
		description = var.ccc_upgrade_schedule_events_description
		schedule_event_duration = var.ccc_upgrade_schedule_events_schedule_event_duration
		time_start = var.ccc_upgrade_schedule_events_time_start

		#Optional
		schedule_event_recurrences = var.ccc_upgrade_schedule_events_schedule_event_recurrences
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.ccc_upgrade_schedule_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Compute Cloud@Customer Upgrade Schedule. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) An optional description of the Compute Cloud@Customer upgrade schedule. Avoid entering confidential information. 
* `display_name` - (Required) (Updatable) Compute Cloud@Customer upgrade schedule display name. Avoid entering confidential information. 
* `events` - (Required) (Updatable) List of preferred times for Compute Cloud@Customer infrastructure to be upgraded. 
	* `description` - (Required) (Updatable) A description of the Compute Cloud@Customer upgrade schedule time block.
	* `schedule_event_duration` - (Required) (Updatable) The duration of this block of time. The duration must be specified and be of the ISO-8601 format for durations. 
	* `schedule_event_recurrences` - (Optional) (Updatable) Frequency of recurrence of schedule block. When this field is not included, the event is assumed to be a one time occurrence. The frequency field is strictly parsed and must conform to RFC-5545 formatting for recurrences. 
	* `time_start` - (Required) (Updatable) The date and time when the Compute Cloud@Customer upgrade schedule event starts, inclusive. An RFC3339 formatted UTC datetime string. For an event with recurrences, this is the date that a recurrence can start being applied. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the Compute Cloud@Customer upgrade schedule. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - An optional description of the Compute Cloud@Customer upgrade schedule. Avoid entering confidential information. 
* `display_name` - Compute Cloud@Customer upgrade schedule display name. Avoid entering confidential information. 
* `events` - List of preferred times for Compute Cloud@Customer infrastructures associated with this schedule to be upgraded. 
	* `description` - A description of the Compute Cloud@Customer upgrade schedule time block.
	* `name` - Generated name associated with the event.
	* `schedule_event_duration` - The duration of this block of time. The duration must be specified and be of the ISO-8601 format for durations. 
	* `schedule_event_recurrences` - Frequency of recurrence of schedule block. When this field is not included, the event is assumed to be a one time occurrence. The frequency field is strictly parsed and must conform to RFC-5545 formatting for recurrences. 
	* `time_start` - The date and time when the Compute Cloud@Customer upgrade schedule event starts, inclusive. An RFC3339 formatted UTC datetime string. For an event with recurrences, this is the date that a recurrence can start being applied. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Upgrade schedule [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). This cannot be changed once created. 
* `infrastructure_ids` - List of Compute Cloud@Customer infrastructure [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that are using this upgrade schedule. 
* `lifecycle_details` - A message describing the current state in more detail. For example, the message can be used to provide actionable information for a resource in a Failed state. 
* `state` - Lifecycle state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the upgrade schedule was created, using an RFC3339 formatted datetime string. 
* `time_updated` - The time the upgrade schedule was updated, using an RFC3339 formatted datetime string. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ccc Upgrade Schedule
	* `update` - (Defaults to 20 minutes), when updating the Ccc Upgrade Schedule
	* `delete` - (Defaults to 20 minutes), when destroying the Ccc Upgrade Schedule


## Import

CccUpgradeSchedules can be imported using the `id`, e.g.

```
$ terraform import oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule "id"
```

