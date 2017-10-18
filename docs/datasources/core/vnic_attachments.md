# oci\_core\_vnic_attachments

**API:** [VnicAttachment Reference][42bf6f8c]

  [42bf6f8c]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/ "VnicAttachmentReference"

Gets information about a specific VNIC attachment.

## Example Usage

```
data "oci_core_vnic_attachments" "s" {
    compartment_id = "compartmentid"
    availability_domain = "availabilityid"
    vnic_id = "vnicid"
    instance_id = "instanceid"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `vnic_id` - (Optional) The OCID of the VNIC.
* `instance_id` - (Optional) The OCID of the instance.
* `availability_domain` - (Optional) The name of the Availability Domain.
* `limit` - (Optional) The maximum number of items to return in a paginated "List" call.
* `page` - (Optional) The value of the opc-next-page response header from the previous "List" call.


## Attributes Reference

The following attributes are exported:

* `vnic_attachments` - A list of vnic attachments.

## Vnic Attachment Reference
* `availability_domain` - The VNIC Attachment's Availability Domain.
* `compartment_id` - The OCID of the compartment.
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the VNIC Attachment.
* `instance_id` - The OCID of the instance.
* `state` - The current state of the VNIC. Allowed values are: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `vnic_id` - The OCID of the VNIC.
* `subnet_id` - The OCID of the subnet the VNIC Attachment is in.
* `time_created` - The date and time the VNIC was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`.
