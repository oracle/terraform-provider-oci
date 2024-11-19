---
subcategory: "Recovery"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_recovery_long_term_backups"
sidebar_current: "docs-oci-datasource-recovery-long_term_backups"
description: |-
  Provides the list of Long Term Backups in Oracle Cloud Infrastructure Recovery service
---

# Data Source: oci_recovery_long_term_backups
This data source provides the list of Long Term Backups in Oracle Cloud Infrastructure Recovery service.

Lists the long-term backups associated with a protected database. You can filter the results using the unique identifier (OCID) of a specific compartment, a protected database, or a long-term backup.


## Example Usage

```hcl
data "oci_recovery_long_term_backups" "test_long_term_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.long_term_backup_display_name
	id = var.long_term_backup_id
	protected_database_id = oci_recovery_protected_database.test_protected_database.id
	state = var.long_term_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment OCID.
* `display_name` - (Optional) A filter to return only resources that match the entire 'displayname' given.
* `id` - (Optional) The long-term backup OCID. Use longTermBackupId to filter a long-term backup based on its unique identifier.
* `protected_database_id` - (Optional) The protected database OCID. Use protectedDatabaseId to list the long-term backups of a specific protected database.
* `state` - (Optional) A filter to return only the resources that match the specified lifecycle state.


## Attributes Reference

The following attributes are exported:

* `long_term_backup_collection` - The list of long_term_backup_collection.

### LongTermBackup Reference

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

