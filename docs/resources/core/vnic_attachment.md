# baremetal\_core\_vnic\_attachment

Provides a VNIC attachment resource

## Example Usage

```
resource "baremetal_core_vnic_attachment" "t" {
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

* `display_name` - A user-friendly name. Does not have to be unique.
* `instance_id` - (Required) The OCID of the instance.
* `create_vnic_details` - (Required) Details for creating a new VNIC. See [Create Vnic Details](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/CreateVnicDetails).

## Create VNIC Details Argument Reference

* `assign_public_ip` - Whether the VNIC should be assigned a public IP address.
* `display_name` - A user-friendly name for the VNIC. Does not have to be unique.
* `hostname_label` - The hostname for the VNIC's primary private IP.
* `private_p` - A private IP address of your choice to assign to the VNIC.
* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in.

## Attributes Reference

* `availability_domain` - The VNIC Attachment's Availability Domain.
* `compartment_id` - The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the VNIC Attachment.
* `instance_id` - The OCID of the instance.
* `state` - The current state of the VNIC attachment. [ATTACHING, ATTACHED, DETACHING, DETACHED].
* `subnet_id` - The OCID of the VNIC's subnet.
* `time_created` - The date and time the VNIC attachment was created, in the format defined by RFC3339.
* `vlan_tag` - The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.
* `vnic_id` - The OCID of the VNIC.
