---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group_backups"
sidebar_current: "docs-oci-datasource-core-volume_group_backups"
description: |-
  Provides the list of Volume Group Backups in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_group_backups
This data source provides the list of Volume Group Backups in Oracle Cloud Infrastructure Core service.

Lists the volume group backups in the specified compartment. You can filter the results by volume group.
For more information, see [Volume Groups](https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/volumegroups.htm).


## Example Usage

```hcl
data "oci_core_volume_group_backups" "test_volume_group_backups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.volume_group_backup_display_name
	volume_group_id = oci_core_volume_group.test_volume_group.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `volume_group_id` - (Optional) The OCID of the volume group.


## Attributes Reference

The following attributes are exported:

* `volume_group_backups` - The list of volume_group_backups.

### VolumeGroupBackup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume group backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume group backup will expire and be automatically deleted. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). This parameter will always be present for volume group backups that were created automatically by a scheduled-backup policy. For manually created volume group backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume group backup.
* `size_in_gbs` - The aggregate size of the volume group backup, in GBs. 
* `size_in_mbs` - The aggregate size of the volume group backup, in MBs. 
* `source_type` - Specifies whether the volume group backup was created manually, or via scheduled backup policy. 
* `source_volume_group_backup_id` - The OCID of the source volume group backup.
* `state` - The current state of a volume group backup.
* `time_created` - The date and time the volume group backup was created. This is the time the actual point-in-time image of the volume group data was taken. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_request_received` - The date and time the request to create the volume group backup was received. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `type` - The type of backup.
* `unique_size_in_gbs` - The aggregate size used by the volume group backup, in GBs.  It is typically smaller than `size_in_gbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `unique_size_in_mbs` - The aggregate size used by the volume group backup, in MBs.  It is typically smaller than `size_in_mbs`, depending on the space consumed on the volume group and whether the volume backup is full or incremental. 
* `volume_backup_ids` - OCIDs for the volume backups in this volume group backup.
* `volume_group_id` - The OCID of the source volume group.

