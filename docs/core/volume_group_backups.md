# oci_core_volume_group_backup

## VolumeGroupBackup Resource

### VolumeGroupBackup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume group backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume group backup. Does not have to be unique and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume group backup (unique).
* `size_in_mbs` - The aggregate size of the volume group backup, in MBs. 
* `state` - The current state of a volume group backup.
* `time_created` - The date and time the volume group backup was created. This is the time the actual point-in-time image of the volume group data was taken. Format defined by RFC3339. 
* `time_request_received` - The date and time the request to create the volume group backup was received. Format defined by RFC3339. 
* `type` - The type of backup.
* `unique_size_in_mbs` - The aggregate size used by the volume group backup, in MBs.  It is typically smaller than sizeInMBs, depending on the space consumed on the volume group and whether the backup is full or incremental. 
* `volume_backup_ids` - OCIDs for the backups in this volume group backup.
* `volume_group_id` - The OCID of the source volume group.



### Create Operation
Creates a new group backup of the specified volume group.


The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that will contain the volume group backup. This parameter is optional, by default backup will be created in the same compartment and source volume group.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name for the volume group backup. Does not have to be unique and it's changeable.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Optional) The type of backup to create. If omitted, defaults to incremental.
  * Allowed values are :
    * FULL
    * INCREMENTAL
* `volume_group_id` - (Required) The OCID of the volume group that needs to be backed up.


### Update Operation
Updates the display name for the specified volume group backup.

The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume group backup. Does not have to be unique and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_volume_group_backup" "test_volume_group_backup" {
	#Required
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"

	#Optional
	compartment_id = "${var.compartment_id}"
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.volume_group_backup_display_name}"
	freeform_tags = {"Department"= "Finance"}
	type = "${var.volume_group_backup_type}"
}
```

# oci_core_volume_group_backups

## VolumeGroupBackup DataSource

Gets a list of volume_group_backups.

### List Operation
Lists the backups for volume groups in the specified compartment. You can filter the results by volume group.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `volume_group_id` - (Optional) The OCID of the volume group.


The following attributes are exported:

* `volume_group_backups` - The list of volume_group_backups.

### Example Usage

```hcl
data "oci_core_volume_group_backups" "test_volume_group_backups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.volume_group_backup_display_name}"
	volume_group_id = "${oci_core_volume_group.test_volume_group.id}"
}
```