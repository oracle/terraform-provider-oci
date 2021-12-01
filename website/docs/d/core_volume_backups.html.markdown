---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backups"
sidebar_current: "docs-oci-datasource-core-volume_backups"
description: |-
  Provides the list of Volume Backups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backups
This data source provides the list of Volume Backups in Oracle Cloud Infrastructure Core service.

Lists the volume backups in the specified compartment. You can filter the results by volume.


## Example Usage

```hcl
data "oci_core_volume_backups" "test_volume_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.volume_backup_display_name
	source_volume_backup_id = oci_core_volume_backup.test_volume_backup.id
	state = var.volume_backup_state
	volume_id = oci_core_volume.test_volume.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `source_volume_backup_id` - (Optional) A filter to return only resources that originated from the given source volume backup. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `volume_id` - (Optional) The OCID of the volume.


## Attributes Reference

The following attributes are exported:

* `volume_backups` - The list of volume_backups.

### VolumeBackup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume backup will expire and be automatically deleted. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). This parameter will always be present for backups that were created automatically by a scheduled-backup policy. For manually created backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume backup.
* `kms_key_id` - The OCID of the Key Management key which is the master encryption key for the volume backup. For more information about the Key Management service and encryption keys, see [Overview of Key Management](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
* `size_in_gbs` - The size of the volume, in GBs. 
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy.
* `source_volume_backup_id` - The OCID of the source volume backup.
* `state` - The current state of a volume backup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `time_created` - The date and time the volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_request_received` - The date and time the request to create the volume backup was received. Format defined by [RFC3339]https://tools.ietf.org/html/rfc3339. 
* `type` - The type of a volume backup. Supported values are 'FULL' or 'INCREMENTAL'.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than sizeInGBs, depending on the space consumed on the volume and whether the backup is full or incremental. 
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than sizeInMBs, depending on the space consumed on the volume and whether the backup is full or incremental. This field is deprecated. Please use uniqueSizeInGBs. 
* `volume_id` - The OCID of the volume.

