# oci\_core\_volume\_backups

**API:** [VolumeBackup Reference][1bc974b1]

  [1bc974b1]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/ "VolumeBackupReference"

Gets a list of volume backups in a compartment.

## Example Usage

```
data "oci_core_volume_backups" "t" {
  compartment_id = "compartmentid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `volume_id` - (Optional) The OCID of a volume.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The pagination token to continue listing from.


## Attributes Reference

The following attributes are exported:

* `volume_backups` - The list of volume backups.

## Volume Backups Reference
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the volume backup.
* `state` - The current state of the volume. Allowed values are: [CREATING, AVAILABLE, TERMINATING, TERMINATED, FAULTY, REQUEST_RECEIVED]
* `size_in_gbs` - The size of the volume, in GBs. The value must be a multiple of 1024.
* `time_created` - The date and time the volume was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `time_requested` - The date and time the request to create the volume backup was received, in the format defined by RFC3339.
* `unique_size_in_gbs` - The size used by the backup, in GBs. It is typically smaller than sizeInGBs, depending on the space consumed on the volume and whether the backup is full or incremental.
* `volume_id` - The OCID of the volume.
