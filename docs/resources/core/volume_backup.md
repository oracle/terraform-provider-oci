# oci\_core\_volume\_backups

Gets a list of volume backups in a compartment.

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
* `display_name` - (Optional) A user-friendly name. Does not have to be unique.


## Attributes Reference
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the Volume backup.
* `state` - The current state of the volume. [CREATING,AVAILABLE,TERMINATING,TERMINATED,FAULTY,REQUEST_RECEIVED]
* `size_in_mbs` - The size of the volume, in MBs.
* `time_created` - The date and time the Volume was created.
* `time_requested` - The date and time the request to create the volume backup was received.
* `unique_size_in_mbs` - The size used by the backup, in MBs. It is typically smaller than sizeInMBs, depending on the space consumed on the volume and whether the backup is full or incremental.
* `volume_id` - The OCID of the Volume.

