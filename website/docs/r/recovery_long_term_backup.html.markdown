---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_long_term_backup"
sidebar_current: "docs-oci-resource-recovery-long_term_backup"
description: |-
  Provides the Long Term Backup resource in Oracle Cloud Infrastructure Recovery service
---

# oci_recovery_long_term_backup
This resource provides the Long Term Backup resource in Oracle Cloud Infrastructure Recovery service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/recovery-service/latest/LongTermBackup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/recovery

Creates a long-term backup of a specified protected database.


## Example Usage

```hcl
resource "oci_recovery_long_term_backup" "test_long_term_backup" {
	#Required
	protected_database_id = oci_recovery_protected_database.test_protected_database.id
	retention_period {
		#Required
		retention_count = var.long_term_backup_retention_period_retention_count
		retention_period_type = var.long_term_backup_retention_period_retention_period_type
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.long_term_backup_display_name
	freeform_tags = {"bar-key"= "value"}
	retention_point_in_time = var.long_term_backup_retention_point_in_time
	retention_scn = var.long_term_backup_retention_scn
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - (Optional) (Updatable) A user provided name for the long term backup. The 'displayName' does not have to be unique, and it can be modified. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `protected_database_id` - (Required) The OCID of the protected database for which you want to create the long-term backup.
* `retention_period` - (Required) (Updatable) The maximum period to retain the long-term backup. Specify the retention period type and the duration for the long-term backup. If you have chosen the retention period type as 'DAYS', then specify a duration ranging from 90 days to 3650 days. If you have chosen the retention period type as 'YEARS', then specify a duration ranging from 1 year to 10 years.
	* `retention_count` - (Required) (Updatable) Specifies the duration (in days or years) to retain the long-term backup. If you have chosen the retentionPeriodType as 'DAYS', then specify a duration ranging from 90 days to 3650 days. If you have chosen the retentionPeriodType as 'YEARS', then specify a duration ranging from 1 year to 10 years.
	* `retention_period_type` - (Required) (Updatable) Specifies the retention period type for the long-term backup. Allowed values are DAYS or YEARS.
* `retention_point_in_time` - (Optional) An RFC3339 formatted datetime string that indicates the desired target point in time in the database at which you want to create the long-term backup. For example, if you want the long-term backup to include all the changes until May 22 at 9:10 PM, then specify the value as, '2020-05-22T21:10:00.000Z'. If you want to specify the target point as an SCN value instead of the target time, then use the databaseSCN parameter.
* `retention_scn` - (Optional) The desired target point (SCN) at which you want to create the long-term backup of the database.For example, specify the value as 1000 if you want to create the long-term backup until SCN 1000. If you want to specify the target point as a time expression instead of the SCN value, then use the longTermRetentionPointInTime parameter.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the long-term backup.
* `database_identifier` - The Oracle Database ID, which identifies an Oracle Database located outside of Oracle Cloud.
* `database_size_in_gbs` - The size of the database, in gigabytes.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `display_name` - The user-provided name for the long-term backup. You can change the displayName. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the long-term backup.
* `lifecycle_details` - A detailed message about the current lifecycle state of the long-term backup. For example, it can be used to provide actionable information for a resource in a Failed state.
* `lifecycle_substate` - More details on the state of the backup when it is in Creating lifecycleState. 
* `protected_database_id` - The OCID of the protected database associated with the long-term backup.
* `retention_period` - The maximum number of DAYS or YEARS to store the long-term backup. You can retain a long-term backup for a period ranging from 90-3650 days or 1-10 years.
	* `retention_count` - Specifies the duration (in days or years) to retain the long-term backup. If you have chosen the retentionPeriodType as 'DAYS', then specify a duration ranging from 90 days to 3650 days. If you have chosen the retentionPeriodType as 'YEARS', then specify a duration ranging from 1 year to 10 years.
	* `retention_period_type` - Specifies the retention period type for the long-term backup. Allowed values are DAYS or YEARS.
* `retention_point_in_time` - An RFC3339 formatted datetime string that indicates the target point in time until which the long-term backup is consistent. For example, '2020-05-22T21:10:29.600Z'.
* `retention_scn` - The unique system change number (SCN) or the target point in the database until which the long-term backup is consistent.
* `retention_until_date_time` - Indicates that Recovery Service must retain the backup for the specified long-term retention period.
* `rman_tag` - Recovery Manager (RMAN) assigned unique identifier for the long-term backup.
* `state` - The current state of the long term backup. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`. For more information, see [Resource Tags](https://docs.oracle.com/en-us/iaas/Content/General/Concepts/resourcetags.htm) 
* `time_backup_completed` - An RFC3339 formatted datetime string that indicates the time when the long-term backup completed. For example, '2020-05-22T21:10:29.600Z'.
* `time_backup_initiated` - An RFC3339 formatted datetime string that indicates the time when the long-term backup was created. For example, '2020-05-22T21:10:29.600Z'.
* `time_created` - An RFC3339 formatted datetime string that indicates the time when the long-term backup was created. For example: '2020-05-22T21:10:29.600Z'. 
* `time_updated` - An RFC3339 formatted datetime string that indicates the time when the long term backup was last updated. For example: '2020-05-22T21:10:29.600Z'. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 24 hours), when creating the Long Term Backup
	* `update` - (Defaults to 20 minutes), when updating the Long Term Backup
	* `delete` - (Defaults to 24 hours), when destroying the Long Term Backup


## Import

LongTermBackups can be imported using the `id`, e.g.

```
$ terraform import oci_recovery_long_term_backup.test_long_term_backup "id"
```

