# oci\_core\_volumes

[Volume Reference][ce7191fd]

  [ce7191fd]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/ "VolumeReference"

Create a volume.

## Example Usage

```
resource "oci_core_volume" "t" {
    availability_domain = "availability_domain"
    compartment_id = "compartment_id"
    volume_backup_id = "volume_id"
    size_in_gbs = 50
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain of the volume.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `compartment_id` - (Required) The OCID of the compartment.
* `volume_backup_id` - (Optional) The OCID of the volume backup from which the data should be restored on the newly created volume.
* `source_details` - (Optional) Specifies the volume source details for a new Block Volume. 
See [Source Details](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/CreateVolumeDetails) documentation.
Example usage: 
```
resource "oci_core_volume" "t" {
    
    source_details {
        type = "volume"
        id = "${var.volume_id}"
    }
     
    // or
     
    source_details {
        type = "volumeBackup"
        id = "${var.volume_backup_id}" // note: this requires an oci_core_volume_backup resource OCID
    }
    ...
}
```


## Attributes Reference
* `availability_domain` - The Availability Domain of the volume.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the Volume backup.
* `state` - The current state of the volume. Allowed values are: [PROVISIONING,RESTORING,AVAILABLE,TERMINATING,TERMINATED,FAULTY]
* `size_in_mbs` - (Deprecated) The size of the volume, in MBs.
* `size_in_gbs` - The size of the volume, in GBs.
* `time_created` - The date and time the Volume was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
* `source_details` - Specifies the volume source details for a new Block Volume.
