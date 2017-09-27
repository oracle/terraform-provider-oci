# oci\_core\_volumes

Gets a list of volumes in a compartment.

## Example Usage

```
resource "oci_core_volume" "t" {
    availability_domain = "availability_domain"
    compartment_id = "compartment_id"
    size_in_gbs = 50
    volume_backup_id = "volume_id"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain of the volume.
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `volume_backup_id` - (Optional) The OCID of the volume backup from which the data should be restored on the newly created volume.

## Attributes Reference
* `availability_domain` - The availability domain of the volume.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the Volume backup.
* `state` - The current state of the volume. [PROVISIONING,RESTORING,AVAILABLE,TERMINATING,TERMINATED,FAULTY]
* `size_in_mbs` - (Deprecated) The size of the volume, in MBs.
* `size_in_gbs` - The size of the volume, in GBs.
* `time_created` - The date and time the Volume was created.
