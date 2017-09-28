# oci\_core\_volumes

[Volume Reference][337fdb07]

  [337fdb07]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/ "VolumeReference"

Gets a list of volumes in a compartment.

## Example Usage

```
data "oci_core_volumes" "t" {
  compartment_id = "compartmentid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `availability_domain` - (Optional) The OCID of a volume.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The pagination token to continue listing from.


## Attributes Reference

The following attributes are exported:

* `volumes` - The list of volumes.

## Volume Backups Reference
* `availability_domain` - The Availability Domain of the volume.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `id` - The OCID of the volume.
* `state` - The current state of the volume. [PROVISIONING, RESTORING, AVAILABLE, TERMINATING, TERMINATED, FAULTY]
* `size_in_mbs` - The size of the volume, in MBs. The size must be a multiple of 1024.
* `time_created` - The date and time the Volume was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
