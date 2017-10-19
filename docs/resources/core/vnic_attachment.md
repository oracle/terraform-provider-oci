# oci\_core\_vnic\_attachment

[VnicAttachment Reference][46f9706c]

  [46f9706c]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VnicAttachment/ "VnicAttachmentReference"

Provides a VNIC attachment resource.

## Example Usage

```
resource "oci_core_vnic_attachment" "t" {
  instance_id = "${var.instance_id}"
  display_name = "secondary_vnic_attachment"
  create_vnic_details {
    subnet_id = "${var.subnet_id}"
    display_name = "secondary_vnic"
    assign_public_ip = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `instance_id` - (Required) The OCID of the instance.
* `create_vnic_details` - (Required) Details for creating a new VNIC. See [Create Vnic Details](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/CreateVnicDetails).

## Create VNIC Details Argument Reference

* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address. Example: `true`
* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information.
* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP.
* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC.
* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in.
* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip).

## Attributes Reference

* `availability_domain` - The VNIC Attachment's Availability Domain.
* `compartment_id` - The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in.
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information.
* `id` - The OCID of the VNIC Attachment.
* `instance_id` - The OCID of the instance.
* `state` - The current state of the VNIC attachment. Allowed values are: [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `subnet_id` - The OCID of the VNIC's subnet.
* `time_created` - The date and time the VNIC attachment was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z`.
* `vlan_tag` - The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.
* `vnic_id` - The OCID of the VNIC.
