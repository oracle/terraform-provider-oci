---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policies"
sidebar_current: "docs-oci-datasource-core-volume_backup_policies"
description: |-
  Provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backup_policies
This data source provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service.

Lists all the volume backup policies available in the specified compartment.

For more information about Oracle defined backup policies and user defined backup policies,
see [Policy-Based Backups](https://docs.cloud.oracle.com/iaas/Content/Block/Tasks/schedulingvolumebackups.htm).


## Example Usage

```hcl
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment. If no compartment is specified, the Oracle defined backup policies are listed. 


## Attributes Reference

The following attributes are exported:

* `volume_backup_policies` - The list of volume_backup_policies.

### VolumeBackupPolicy Reference

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
* `time_created` - The date and time the volume backup policy was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

