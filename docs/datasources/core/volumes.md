# baremetal\_core\_volumes

Gets a list of volumes in a compartment.

## Example Usage

```
data "baremetal_core_volumes" "t" {
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
* `availability_domain` - The availability domain of the volume.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the Volume.
* `state` - The current state of the volume. [PROVISIONING,RESTORING,AVAILABLE,TERMINATING,TERMINATED,FAULTY]
* `size_in_mbs` - The size of the volume, in MBs.
* `time_created` - The date and time the Volume was created.
