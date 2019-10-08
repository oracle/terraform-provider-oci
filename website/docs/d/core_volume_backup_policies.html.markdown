---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policies"
sidebar_current: "docs-oci-datasource-core-volume_backup_policies"
description: |-
  Provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backup_policies
This data source provides the list of Volume Backup Policies in Oracle Cloud Infrastructure Core service.

Lists all volume backup policies available to the caller.

## Example Usage

```hcl
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {

	#Optional
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment to list. If no compartment is specified, list the predefined (Gold, Silver, Bronze) backup policies. 


## Attributes Reference

The following attributes are exported:

* `volume_backup_policies` - The list of volume_backup_policies.

### VolumeBackupPolicy Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume backup policy. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume backup policy.
* `schedules` - The collection of schedules that this policy will apply.
	* `backup_type` - The type of backup to create.
	* `day_of_month` - The day of the month to schedule the backup
	* `day_of_week` - The day of the week to schedule the backup
	* `hour_of_day` - The hour of the day to schedule the backup
	* `month` - The month of the year to schedule the backup
	* `offset_seconds` - The number of seconds that the backup time should be shifted from the default interval boundaries specified by the period. Backup time = Frequency start time + Offset.
	* `offset_type` - Indicates how offset is defined. If value is `STRUCTURED`, then `hourOfDay`, `dayOfWeek`, `dayOfMonth`, and `month` fields are used and `offsetSeconds` will be ignored in requests and users should ignore its value from the respones. `hourOfDay` is applicable for periods `ONE_DAY`, `ONE_WEEK`, `ONE_MONTH` and `ONE_YEAR`. `dayOfWeek` is applicable for period `ONE_WEEK`. `dayOfMonth` is applicable for periods `ONE_MONTH` and `ONE_YEAR`. 'month' is applicable for period 'ONE_YEAR'. They will be ignored in the requests for inapplicable periods. If value is `NUMERIC_SECONDS`, then `offsetSeconds` will be used for both requests and responses and the structured fields will be ignored in the requests and users should ignore their values from the respones. For clients using older versions of Apis and not sending `offsetType` in their requests, the behaviour is just like `NUMERIC_SECONDS`.
	* `period` - How often the backup should occur.
	* `retention_seconds` - How long, in seconds, backups created by this schedule should be kept until being automatically deleted.
	* `time_zone` - Specifies what time zone is the schedule in
* `time_created` - The date and time the volume backup policy was created. Format defined by RFC3339. 

