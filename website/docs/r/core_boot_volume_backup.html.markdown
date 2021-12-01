---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume_backup"
sidebar_current: "docs-oci-resource-core-boot_volume_backup"
description: |-
  Provides the Boot Volume Backup resource in Oracle Cloud Infrastructure Core service
---

# oci_core_boot_volume_backup
This resource provides the Boot Volume Backup resource in Oracle Cloud Infrastructure Core service.

Creates a new boot volume backup of the specified boot volume. For general information about boot volume backups,
see [Overview of Boot Volume Backups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/bootvolumebackups.htm)

When the request is received, the backup object is in a REQUEST_RECEIVED state.
When the data is imaged, it goes into a CREATING state.
After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.


## Example Usage

```hcl
resource "oci_core_boot_volume_backup" "test_boot_volume_backup" {
	#Required
	boot_volume_id = oci_core_boot_volume.test_boot_volume.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.boot_volume_backup_display_name
	freeform_tags = {"Department"= "Finance"}
	type = var.boot_volume_backup_type
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_id` - (Optional) The OCID of the boot volume that needs to be backed up. Cannot be defined if `source_details` is defined.
* `compartment_id` - (Optional) (Updatable) The OCID of the compartment that contains the boot volume backup.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Optional) The type of backup to create. If omitted, defaults to incremental. Supported values are 'FULL' or 'INCREMENTAL'.
* `source_details` - (Optional) Details of the volume backup source in the cloud. Cannot be defined if `boot_volume_id` is defined.
    * `kms_key_id` - (Optional) The OCID of the KMS key in the destination region which will be the master encryption key for the copied volume backup.
    * `region` - (Required) The region of the volume backup source.
    * `volume_backup_id` - (Required) The OCID of the source volume backup.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `boot_volume_id` - The OCID of the boot volume.
* `compartment_id` - The OCID of the compartment that contains the boot volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume backup will expire and be automatically deleted. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). This parameter will always be present for backups that were created automatically by a scheduled-backup policy. For manually created backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the boot volume backup.
* `image_id` - The image OCID used to create the boot volume the backup is taken from. 
* `kms_key_id` - The OCID of the Key Management master encryption assigned to the boot volume backup. For more information about the Key Management service and encryption keys, see [Overview of Key Management](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
* `size_in_gbs` - The size of the boot volume, in GBs. 
* `source_boot_volume_backup_id` - The OCID of the source boot volume backup.
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy. 
* `state` - The current state of a boot volume backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the boot volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_request_received` - The date and time the request to create the boot volume backup was received. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `type` - The type of a volume backup. Supported values are 'FULL' or 'INCREMENTAL'.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than sizeInGBs, depending on the space consumed on the boot volume and whether the backup is full or incremental. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Boot Volume Backup
	* `update` - (Defaults to 20 minutes), when updating the Boot Volume Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Boot Volume Backup


## Import

BootVolumeBackups can be imported using the `id`, e.g.

```
$ terraform import oci_core_boot_volume_backup.test_boot_volume_backup "id"
```

