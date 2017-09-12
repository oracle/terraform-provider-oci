# oci\_core\_vnic_attachments

Gets information about a specific vnic.

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
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the VNIC Attachment.
* `instance_id` - The OCID of the instance.
* `state` - The current state of the VNIC. [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `vnic_id` - The OCID of the VNIC.
* `subnet_id` - The OCID of the subnet the VNIC Attachment is in.
* `time_created` - The date and time the VNIC was created.
