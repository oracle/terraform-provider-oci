# oci\_core\_volume\_attachments

Gets a list of volume attachments.

## Example Usage

```
data "oci_core_volume_attachments" "t" {
  availability_domain = "availability_domain"
  compartment_id = "compartment_id"
  limit = 1
  page = "page"
  instance_id = "instance_id"
  volume_id = "volume_id"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `availability_domain` - (Optional) The name of the availability domain.
* `instance_id` - (Optional) The OCID of the instance.
* `volume_id` - (Optional) The OCID of the volume.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The pagination token to continue listing from.


## Attributes Reference

The following attributes are exported:

* `virtual_networks` - The list of virtual networks.

## Volume Attachment Reference
* `attachment_type` - The type of volume attachment.
* `availability_domain` - The Availability Domain of an instance.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique, and it cannot be changed.
* `id` - The OCID of the volume attachment.
* `instance_id` - The OCID of the instance the volume is attached to.
* `state` - The current state of the volume attachment: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `time_created` - The date and time the volume was created
* `volume_id` - The OCID of the volume.


