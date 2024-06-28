---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_filesystem_snapshot_policies"
sidebar_current: "docs-oci-datasource-file_storage-filesystem_snapshot_policies"
description: |-
  Provides the list of Filesystem Snapshot Policies in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_filesystem_snapshot_policies
This data source provides the list of Filesystem Snapshot Policies in Oracle Cloud Infrastructure File Storage service.

Lists file system snapshot policies in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_filesystem_snapshot_policies" "test_filesystem_snapshot_policies" {
	#Required
	availability_domain = var.filesystem_snapshot_policy_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.filesystem_snapshot_policy_display_name
	id = var.filesystem_snapshot_policy_id
	state = var.filesystem_snapshot_policy_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `filesystem_snapshot_policies` - The list of filesystem_snapshot_policies.

### FilesystemSnapshotPolicy Reference

The following attributes are exported:

* `availability_domain` - The availability domain that the file system snapshot policy is in. May be unset using a blank or NULL value.  Example: `Uocm:PHX-AD-2` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system snapshot policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My Filesystem Snapshot Policy` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system snapshot policy.
* `policy_prefix` - The prefix to apply to all snapshots created by this policy.  Example: `acme` 
* `schedules` - The list of associated snapshot schedules. A maximum of 10 schedules can be associated with a policy. 
	* `day_of_month` - The day of the month to create a scheduled snapshot. If the day does not exist for the month, snapshot creation will be skipped. Used for MONTHLY and YEARLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `day_of_week` - The day of the week to create a scheduled snapshot. Used for WEEKLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `hour_of_day` - The hour of the day to create a DAILY, WEEKLY, MONTHLY, or YEARLY snapshot. If not set, the system chooses a value at creation time. 
	* `month` - The month to create a scheduled snapshot. Used only for YEARLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `period` - The frequency of scheduled snapshots.
	* `retention_duration_in_seconds` - The number of seconds to retain snapshots created with this schedule. Snapshot expiration time will not be set if this value is empty. 
	* `schedule_prefix` - A name prefix to be applied to snapshots created by this schedule.  Example: `compliance1` 
	* `time_schedule_start` - The starting point used to begin the scheduling of the snapshots based upon recurrence string in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. If no `timeScheduleStart` is provided, the value will be set to the time when the schedule was created. 
	* `time_zone` - Time zone used for scheduling the snapshot.
* `state` - The current state of this file system snapshot policy. 
* `time_created` - The date and time the file system snapshot policy was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

