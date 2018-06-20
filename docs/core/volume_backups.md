# oci_core_volume_backup

## VolumeBackup Resource

### VolumeBackup Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the volume backup.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `expiration_time` - The date and time the volume backup will expire and be automatically deleted. Format defined by RFC3339. This parameter will always be present for backups that were created automatically by a scheduled-backup policy. For manually created backups, it will be absent, signifying that there is no expiration time and the backup will last forever until manually deleted. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the volume backup.
* `size_in_gbs` - The size of the volume, in GBs. 
* `size_in_mbs` - The size of the volume in MBs. The value must be a multiple of 1024. This field is deprecated. Please use `size_in_gbs`. 
* `source_type` - Specifies whether the backup was created manually, or via scheduled backup policy.
* `state` - The current state of a volume backup.
* `time_created` - The date and time the volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by RFC3339. 
* `time_request_received` - The date and time the request to create the volume backup was received. Format defined by RFC3339. 
* `type` - The type of a volume backup. Supported values are 'FULL' or 'INCREMENTAL'.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than sizeInGBs, depending on the space consumed on the volume and whether the backup is full or incremental. 
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than sizeInMBs, depending on the space consumed on the volume and whether the backup is full or incremental. This field is deprecated. Please use uniqueSizeInGBs. 
* `volume_id` - The OCID of the volume.



### Create Operation
Creates a new backup of the specified volume. For general information about volume backups,
see [Overview of Block Volume Service Backups](https://docs.us-phoenix-1.oraclecloud.com/Content/Block/Concepts/blockvolumebackups.htm)

When the request is received, the backup object is in a REQUEST_RECEIVED state.
When the data is imaged, it goes into a CREATING state.
After the backup is fully uploaded to the cloud, it goes into an AVAILABLE state.


The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Optional) The type of backup to create. If omitted, defaults to INCREMENTAL. Supported values are 'FULL' or 'INCREMENTAL'.
* `volume_id` - (Required) The OCID of the volume that needs to be backed up.


### Update Operation
Updates the display name for the specified volume backup.
Avoid entering confidential information.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_volume_backup" "test_volume_backup" {
	#Required
	volume_id = "${oci_core_volume.test_volume.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.volume_backup_display_name}"
	freeform_tags = {"Department"= "Finance"}
	type = "${var.volume_backup_type}"
}
```

# oci_core_volume_backups

## VolumeBackup DataSource

Gets a list of volume_backups.

### List Operation
Lists the volume backups in the specified compartment. You can filter the results by volume.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `volume_id` - (Optional) The OCID of the volume.


The following attributes are exported:

* `volume_backups` - The list of volume_backups.

### Example Usage

```hcl
data "oci_core_volume_backups" "test_volume_backups" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.volume_backup_display_name}"
	state = "${var.volume_backup_state}"
	volume_id = "${oci_core_volume.test_volume.id}"
}
```