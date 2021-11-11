---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume_backups"
sidebar_current: "docs-oci-datasource-core-boot_volume_backups"
description: |-
  Provides the list of Boot Volume Backups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_boot_volume_backups
This data source provides the list of Boot Volume Backups in Oracle Cloud Infrastructure Core service.

Lists the boot volume backups in the specified compartment. You can filter the results by boot volume.


## Example Usage

```hcl
data "oci_core_boot_volume_backups" "test_boot_volume_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	boot_volume_id = oci_core_boot_volume.test_boot_volume.id
	display_name = var.boot_volume_backup_display_name
	source_boot_volume_backup_id = oci_core_boot_volume_backup.test_boot_volume_backup.id
	state = var.boot_volume_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_id` - (Optional) The OCID of the boot volume.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `source_boot_volume_backup_id` - (Optional) A filter to return only resources that originated from the given source boot volume backup. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `boot_volume_backups` - The list of boot_volume_backups.

### BootVolumeBackup Reference

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

