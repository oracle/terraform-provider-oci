---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policy"
sidebar_current: "docs-oci-resource-core-volume_backup_policy"
description: |-
  Provides the Volume Backup Policy resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_backup_policy
This resource provides the Volume Backup Policy resource in Oracle Cloud Infrastructure Core service.

Creates a new user defined backup policy.

For more information about Oracle defined backup policies and user defined backup policies,
see [Policy-Based Backups](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm).


## Example Usage

```hcl
resource "oci_core_volume_backup_policy" "test_volume_backup_policy" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	destination_region = var.volume_backup_policy_destination_region
	display_name = var.volume_backup_policy_display_name
	freeform_tags = {"Department"= "Finance"}
	schedules {
		#Required
		backup_type = var.volume_backup_policy_schedules_backup_type
		period = var.volume_backup_policy_schedules_period
		retention_seconds = var.volume_backup_policy_schedules_retention_seconds

		#Optional
		day_of_month = var.volume_backup_policy_schedules_day_of_month
		day_of_week = var.volume_backup_policy_schedules_day_of_week
		hour_of_day = var.volume_backup_policy_schedules_hour_of_day
		month = var.volume_backup_policy_schedules_month
		offset_seconds = var.volume_backup_policy_schedules_offset_seconds
		offset_type = var.volume_backup_policy_schedules_offset_type
		time_zone = var.volume_backup_policy_schedules_time_zone
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `destination_region` - (Optional) (Updatable) The paired destination region for copying scheduled backups to. Example: `us-ashburn-1`. See [Region Pairs](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm#RegionPairs) for details about paired regions. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `schedules` - (Optional) (Updatable) The collection of schedules for the volume backup policy. See see [Schedules](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm#schedules) in [Policy-Based Backups](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm) for more information. 
	* `backup_type` - (Required) (Updatable) The type of volume backup to create.
	* `day_of_month` - (Optional) (Updatable) The day of the month to schedule the volume backup.
	* `day_of_week` - (Optional) (Updatable) The day of the week to schedule the volume backup.
	* `hour_of_day` - (Optional) (Updatable) The hour of the day to schedule the volume backup.
	* `month` - (Optional) (Updatable) The month of the year to schedule the volume backup.
	* `offset_seconds` - (Optional) (Updatable) The number of seconds that the volume backup start time should be shifted from the default interval boundaries specified by the period. The volume backup start time is the frequency start time plus the offset. 
	* `offset_type` - (Optional) (Updatable) Indicates how the offset is defined. If value is `STRUCTURED`, then `hourOfDay`, `dayOfWeek`, `dayOfMonth`, and `month` fields are used and `offsetSeconds` will be ignored in requests and users should ignore its value from the responses.

		`hourOfDay` is applicable for periods `ONE_DAY`, `ONE_WEEK`, `ONE_MONTH` and `ONE_YEAR`.

		`dayOfWeek` is applicable for period `ONE_WEEK`.

		`dayOfMonth` is applicable for periods `ONE_MONTH` and `ONE_YEAR`.

		'month' is applicable for period 'ONE_YEAR'.

		They will be ignored in the requests for inapplicable periods.

		If value is `NUMERIC_SECONDS`, then `offsetSeconds` will be used for both requests and responses and the structured fields will be ignored in the requests and users should ignore their values from the responses.

		For clients using older versions of Apis and not sending `offsetType` in their requests, the behaviour is just like `NUMERIC_SECONDS`. 
	* `period` - (Required) (Updatable) The volume backup frequency.
	* `retention_seconds` - (Required) (Updatable) How long, in seconds, to keep the volume backups created by this schedule.
	* `time_zone` - (Optional) (Updatable) Specifies what time zone is the schedule in
        enum:
        - `UTC`
        - `REGIONAL_DATA_CENTER_TIME`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `destination_region` - The paired destination region for copying scheduled backups to. Example `us-ashburn-1`. See [Region Pairs](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm#RegionPairs) for details about paired regions. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume backup policy.
* `schedules` - The collection of schedules that this policy will apply.
	* `backup_type` - The type of volume backup to create.
	* `day_of_month` - The day of the month to schedule the volume backup.
	* `day_of_week` - The day of the week to schedule the volume backup.
	* `hour_of_day` - The hour of the day to schedule the volume backup.
	* `month` - The month of the year to schedule the volume backup.
	* `offset_seconds` - The number of seconds that the volume backup start time should be shifted from the default interval boundaries specified by the period. The volume backup start time is the frequency start time plus the offset. 
	* `offset_type` - Indicates how the offset is defined. If value is `STRUCTURED`, then `hourOfDay`, `dayOfWeek`, `dayOfMonth`, and `month` fields are used and `offsetSeconds` will be ignored in requests and users should ignore its value from the responses.

		`hourOfDay` is applicable for periods `ONE_DAY`, `ONE_WEEK`, `ONE_MONTH` and `ONE_YEAR`.

		`dayOfWeek` is applicable for period `ONE_WEEK`.

		`dayOfMonth` is applicable for periods `ONE_MONTH` and `ONE_YEAR`.

		'month' is applicable for period 'ONE_YEAR'.

		They will be ignored in the requests for inapplicable periods.

		If value is `NUMERIC_SECONDS`, then `offsetSeconds` will be used for both requests and responses and the structured fields will be ignored in the requests and users should ignore their values from the responses.

		For clients using older versions of Apis and not sending `offsetType` in their requests, the behaviour is just like `NUMERIC_SECONDS`. 
	* `period` - The volume backup frequency.
	* `retention_seconds` - How long, in seconds, to keep the volume backups created by this schedule.
	* `time_zone` - Specifies what time zone is the schedule in
        enum:
        - `UTC`
        - `REGIONAL_DATA_CENTER_TIME`
* `time_created` - The date and time the volume backup policy was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Volume Backup Policy
	* `update` - (Defaults to 20 minutes), when updating the Volume Backup Policy
	* `delete` - (Defaults to 20 minutes), when destroying the Volume Backup Policy


## Import

VolumeBackupPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_backup_policy.test_volume_backup_policy "id"
```

