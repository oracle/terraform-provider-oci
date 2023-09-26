---
subcategory: "Compute Cloud At Customer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_compute_cloud_at_customer_ccc_upgrade_schedules"
sidebar_current: "docs-oci-datasource-compute_cloud_at_customer-ccc_upgrade_schedules"
description: |-
  Provides the list of Ccc Upgrade Schedules in Oracle Cloud Infrastructure Compute Cloud At Customer service
---

# Data Source: oci_compute_cloud_at_customer_ccc_upgrade_schedules
This data source provides the list of Ccc Upgrade Schedules in Oracle Cloud Infrastructure Compute Cloud At Customer service.

Returns a list of Compute Cloud@Customer upgrade schedules.


## Example Usage

```hcl
data "oci_compute_cloud_at_customer_ccc_upgrade_schedules" "test_ccc_upgrade_schedules" {

	#Optional
	access_level = var.ccc_upgrade_schedule_access_level
	ccc_upgrade_schedule_id = oci_compute_cloud_at_customer_ccc_upgrade_schedule.test_ccc_upgrade_schedule.id
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.ccc_upgrade_schedule_compartment_id_in_subtree
	display_name = var.ccc_upgrade_schedule_display_name
	display_name_contains = var.ccc_upgrade_schedule_display_name_contains
	state = var.ccc_upgrade_schedule_state
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `ccc_upgrade_schedule_id` - (Optional) Compute Cloud@Customer upgrade schedule [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources. 
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and sub-compartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `display_name_contains` - (Optional) A filter to return only resources whose display name contains the substring. 
* `state` - (Optional) A filter to return resources only when their lifecycleState matches the given lifecycleState. 


## Attributes Reference

The following attributes are exported:

* `ccc_upgrade_schedule_collection` - The list of ccc_upgrade_schedule_collection.

### CccUpgradeSchedule Reference

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

