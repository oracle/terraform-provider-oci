---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group_backup"
sidebar_current: "docs-oci-resource-core-volume_group_backup"
description: |-
  Provides the Volume Group Backup resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_group_backup
This resource provides the Volume Group Backup resource in Oracle Cloud Infrastructure Core service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/iaas/latest/VolumeGroupBackup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/

Creates a new backup volume group of the specified volume group.
For more information, see [Volume Groups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).


## Example Usage

```hcl
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
	#Required
	volume_group_id = oci_core_volume_group.test_volume_group.id

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.volume_group_backup_display_name
	freeform_tags = {"Department"= "Finance"}
	is_indefinite_retention_enabled = var.volume_group_backup_is_indefinite_retention_enabled
	is_prevent_deletion_enabled = var.volume_group_backup_is_prevent_deletion_enabled
	is_retention_lock_enabled = var.volume_group_backup_is_retention_lock_enabled
	retention_period {
		#Required
		retention_time_amount = var.volume_group_backup_retention_period_retention_time_amount
		retention_time_unit = var.volume_group_backup_retention_period_retention_time_unit
	}
	type = var.volume_group_backup_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that will contain the volume group backup. This parameter is optional, by default backup will be created in the same compartment and source volume group. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_indefinite_retention_enabled` - (Optional) (Updatable) feature that preserves backup data from modification or deletion to ensure it remains available for legal or regulatory investigations or litigation, regardless of standard retention policies. This is an optional field. If it is not specified, it is set to null, no legal hold will be applied to the backups.
* `is_prevent_deletion_enabled` - (Optional) (Updatable) Prevent backups from being deleted during the configured retention period. This is an optional field. If it is not specified, it is set to null, prevent deletion will not be applied to the backups.
* `is_retention_lock_enabled` - (Optional) (Updatable) feature that prevents deletion or alteration of backup data for a specified period to ensure data protection and regulatory compliance. This is an optional field. If it is not specified, it is set to null, no retention lock will be applied to the backups. This feature should be used in conjunction with the retention-period field.
* `retention_period` - (Optional) (Updatable) This field is used to define the retention period for backups. This is an optional field. If it is not specified, it is set to null, no retention period will be applied to the backups.
	* `retention_time_amount` - (Required) (Updatable) The value to enter for the amount of retention time should be a numerical figure (such as 1, 7, 30, etc.) that corresponds to the period specified in the retention time unit property (such as YEARS, DAYS). The combination of these two properties determines the total length of the retention period.
	* `retention_time_unit` - (Required) (Updatable) The value you can assign to the Time Unit property for this Duration may be either "YEARS" or "DAYS".
* `type` - (Optional) The type of backup to create. If omitted, defaults to incremental.
	* Allowed values are :
		* FULL
		* INCREMENTAL
* `volume_group_id` - (Required) The OCID of the volume group that needs to be backed up.
* `source_details` - (Optional) Details of the volume group backup source in the cloud.
    * `kms_key_id` - (Optional) The OCID of the KMS key in the destination region which will be the master encryption key for the copied volume backup.
    * `region` - (Required) The region of the volume backup source.
    * `volume_group_backup_id` - (Required) The OCID of the source volume group backup.



** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume group backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume group backup will expire and be automatically deleted. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). This parameter will always be present for volume group backups that were created automatically by a scheduled-backup policy. For manually created volume group backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume group backup.
* `is_indefinite_retention_enabled` - feature that preserves backup data from modification or deletion to ensure it remains available for legal or regulatory investigations or litigation, regardless of standard retention policies. This is an optional field. If it is not specified, it is set to null, no legal hold will be applied to the backups.
* `is_prevent_deletion_enabled` - Prevent backups from being deleted during the configured retention period. This is an optional field. If it is not specified, it is set to null, prevent deletion will not be applied to the backups.
* `is_retention_lock_enabled` - feature that prevents deletion or alteration of backup data for a specified period to ensure data protection and regulatory compliance. This is an optional field. If it is not specified, it is set to null, no retention lock will be applied to the backups. This feature should be used in conjunction with the retention-period field.
* `retention_period` - This field is used to define the retention period for backups. This is an optional field. If it is not specified, it is set to null, no retention period will be applied to the backups.
	* `retention_time_amount` - The value to enter for the amount of retention time should be a numerical figure (such as 1, 7, 30, etc.) that corresponds to the period specified in the retention time unit property (such as YEARS, DAYS). The combination of these two properties determines the total length of the retention period.
	* `retention_time_unit` - The value you can assign to the Time Unit property for this Duration may be either "YEARS" or "DAYS".
* `size_in_gbs` - The aggregate size of the volume group backup, in GBs. 
* `size_in_mbs` - The aggregate size of the volume group backup, in MBs. 
* `source_type` - Specifies whether the volume group backup was created manually, or via scheduled backup policy. 
* `source_volume_group_backup_id` - The OCID of the source volume group backup.
* `state` - The current state of a volume group backup.
* `time_created` - The date and time the volume group backup was created. This is the time the actual point-in-time image of the volume group data was taken. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_request_received` - The date and time the request to create the volume group backup was received. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_retention_expires_at` - The date and time when a backup’s retention period ends and it is set to expire. This is an optional field. If it is not specified, it is set to null, no retention period will be applied to the backups.
* `type` - The type of backup.
* `unique_size_in_gbs` - The aggregate size used by the volume group backup, in GBs.  It is typically smaller than `size_in_gbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `unique_size_in_mbs` - The aggregate size used by the volume group backup, in MBs.  It is typically smaller than `size_in_mbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `volume_backup_ids` - OCIDs for the volume backups in this volume group backup.
* `volume_group_id` - The OCID of the source volume group.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Volume Group Backup
	* `update` - (Defaults to 20 minutes), when updating the Volume Group Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Volume Group Backup


## Import

VolumeGroupBackups can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_group_backup.test_volume_group_backup "id"
```

