# baremetal\_core\_virtual_networks

Provides a virtual network resource.


## Example Usage

```
resource "baremetal_core_virtual_network" "t" {
    cidr_block = "cidr_block"
    compartment_id = "compartment_id"
    display_name = "display_name"
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Required) The CIDR IP address block of the VCN.
* `compartment_id` - (Required) The OCID of the compartment to contain the VCN.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable.
* `dns_label` - (Optional) A DNS label for the VCN.

## Attributes Reference
* `compartment_id` - The OCID of the compartment.
* `cidr_block` - The CIDR IP address block of the VCN.
* `default_dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `default_route_table_id` - The OCID for the VCN's default route table.
* `default_security_list_id` - The OCID for the VCN's default security list.
* `display_name` - A user-friendly name. Does not have to be unique.
* `id` - The OCID of the VNIC.
* `state` - The current state of the VNIC. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VNIC was created.
