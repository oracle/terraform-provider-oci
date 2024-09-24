---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_filesystem_snapshot_policy"
sidebar_current: "docs-oci-resource-file_storage-filesystem_snapshot_policy"
description: |-
  Provides the Filesystem Snapshot Policy resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_filesystem_snapshot_policy
This resource provides the Filesystem Snapshot Policy resource in Oracle Cloud Infrastructure File Storage service.

Creates a new file system snapshot policy in the specified compartment and
availability domain.

After you create a file system snapshot policy, you can associate it with
file systems.


## Example Usage

```hcl
resource "oci_file_storage_filesystem_snapshot_policy" "test_filesystem_snapshot_policy" {
	#Required
	availability_domain = var.filesystem_snapshot_policy_availability_domain
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.filesystem_snapshot_policy_display_name
	freeform_tags = {"Department"= "Finance"}
	locks {
		#Required
		type = var.filesystem_snapshot_policy_locks_type

		#Optional
		message = var.filesystem_snapshot_policy_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.filesystem_snapshot_policy_locks_time_created
	}
	policy_prefix = var.filesystem_snapshot_policy_policy_prefix
	schedules {
		#Required
		period = var.filesystem_snapshot_policy_schedules_period
		time_zone = var.filesystem_snapshot_policy_schedules_time_zone

		#Optional
		day_of_month = var.filesystem_snapshot_policy_schedules_day_of_month
		day_of_week = var.filesystem_snapshot_policy_schedules_day_of_week
		hour_of_day = var.filesystem_snapshot_policy_schedules_hour_of_day
		month = var.filesystem_snapshot_policy_schedules_month
		retention_duration_in_seconds = var.filesystem_snapshot_policy_schedules_retention_duration_in_seconds
		schedule_prefix = var.filesystem_snapshot_policy_schedules_schedule_prefix
		time_schedule_start = var.filesystem_snapshot_policy_schedules_time_schedule_start
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain that the file system snapshot policy is in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system snapshot policy.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `policy1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `policy_prefix` - (Optional) (Updatable) The prefix to apply to all snapshots created by this policy.  Example: `acme` 
* `schedules` - (Optional) (Updatable) The list of associated snapshot schedules. A maximum of 10 schedules can be associated with a policy.

	If using the CLI, provide the schedule as a list of JSON strings, with the list wrapped in quotation marks, i.e. ``` --schedules '[{"timeZone":"UTC","period":"DAILY","hourOfDay":18},{"timeZone":"UTC","period":"HOURLY"}]' ``` 
	* `day_of_month` - (Optional) (Updatable) The day of the month to create a scheduled snapshot. If the day does not exist for the month, snapshot creation will be skipped. Used for MONTHLY and YEARLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `day_of_week` - (Optional) (Updatable) The day of the week to create a scheduled snapshot. Used for WEEKLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `hour_of_day` - (Optional) (Updatable) The hour of the day to create a DAILY, WEEKLY, MONTHLY, or YEARLY snapshot. If not set, the system chooses a value at creation time. 
	* `month` - (Optional) (Updatable) The month to create a scheduled snapshot. Used only for YEARLY snapshot schedules. If not set, the system chooses a value at creation time. 
	* `period` - (Required) (Updatable) The frequency of scheduled snapshots.
	* `retention_duration_in_seconds` - (Optional) (Updatable) The number of seconds to retain snapshots created with this schedule. Snapshot expiration time will not be set if this value is empty. 
	* `schedule_prefix` - (Optional) (Updatable) A name prefix to be applied to snapshots created by this schedule.  Example: `compliance1` 
	* `time_schedule_start` - (Optional) (Updatable) The starting point used to begin the scheduling of the snapshots based upon recurrence string in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. If no `timeScheduleStart` is provided, the value will be set to the time when the schedule was created. 
	* `time_zone` - (Required) (Updatable) Time zone used for scheduling the snapshot.
* `state` - (Optional) (Updatable) The target state for the Filesystem Snapshot Policy. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain that the file system snapshot policy is in. May be unset using a blank or NULL value.  Example: `Uocm:PHX-AD-2` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system snapshot policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My Filesystem Snapshot Policy` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system snapshot policy.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Filesystem Snapshot Policy
	* `update` - (Defaults to 20 minutes), when updating the Filesystem Snapshot Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Filesystem Snapshot Policy


## Import

FilesystemSnapshotPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy "id"
```

