# oci\_core\_volume\_backups

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
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the Volume backup.
* `state` - The current state of the volume. [CREATING,AVAILABLE,TERMINATING,TERMINATED,FAULTY,REQUEST_RECEIVED]
* `size_in_mbs` - The size of the volume, in MBs.
* `time_created` - The date and time the Volume was created.
* `time_requested` - The date and time the request to create the volume backup was received.
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than sizeInMBs, depending on the space consumed on the volume and whether the backup is full or incremental.
* `volume_id` - The OCID of the Volume.

