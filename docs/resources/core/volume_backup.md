# oci\_core\_volume\_backups

[VolumeBackup Reference][aa478c03]

  [aa478c03]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/ "VolumeBackupReference"

Gets a list of volume backups in a compartment. Volume backups are a point-in-time copy of a volume that can then be used to create a new block volume or recover a block volume.

## Example Usage

```
resource "oci_core_volume_backup" "t" {
    volume_id = "volume_id"
    display_name = "display_name"
}
```

## Argument Reference

The following arguments are supported:

* `volume_id` - (Optional) The OCID of a volume.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique. Avoid entering confidential information.


## Attributes Reference
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name for the volume backup. Does not have to be unique and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the Volume backup.
* `state` - The current state of the volume. Allowed values are: [CREATING, AVAILABLE, TERMINATING, TERMINATED, FAULTY, REQUEST_RECEIVED]
* `size_in_mbs` - The size of the volume, in MBs. Must be a multiple of 1024.
* `time_created` - The date and time the volume backup was created. This is the time the actual point-in-time image of the volume data was taken. Format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `time_requested` - The date and time the request to create the volume backup was received, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than `sizeInMBs`, depending on the space consumed on the volume and whether the backup is full or incremental.
* `volume_id` - The OCID of the Volume.
