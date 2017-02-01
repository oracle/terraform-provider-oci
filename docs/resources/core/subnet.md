# baremetal\_core\_subnets

Gets a list of subnets.

## Example Usage

```
resource "baremetal_core_subnet" "t" {
    availability_domain = "availabilitydomainid"
    compartment_id = "compartmentid"
    display_name = "display_name"
    cidr_block = "10.10.10.0/24"
    route_table_id = "routetableid"
    vcn_id = "vcnid"
    security_list_ids = ["slid1", "slid2"]
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain to contain the subnet.
* `compartment_id` - (Required) The OCID of the compartment to contain the subnet.
* `cidr_block` - (Required) The CIDR IP address range of the subnet.
* `vcn_id` - (Required) The OCID of the VCN to contain the subnet.
* `dhcp_options_id` - (Optional) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet will use the VCN's default set of DHCP options.
* `display_name` - (Optional) The maximum number of items to return in a paginated "List" call.
* `route_table_id` - (Optional) The OCID of the route table the subnet will use. If you don't provide a value, the subnet will use the VCN's default route table.
* `security_list_ids` - (Optional) OCIDs for the security lists to associate with the subnet. If you don't provide a value, the VCN's default security list will be associated with the subnet. Remember that security lists are associated at the subnet level, but the rules are applied to the individual VNICs in the subnet.



## Attributes Reference

* `availability_domain` - The subnet's Availability Domain.
* `cidr_block` - The CIDR IP address block of the VCN.
* `compartment_id` - The OCID of the compartment containing the VCN.
* `dhcp_options_id` - The OCID for the VCN's default set of DHCP options.
* `route_table_id` - The OCID for the VCN's default route table.
* `security_list_ids` - OCIDs for the security lists to use for VNICs in this subnet.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `id` - The subnet's Oracle ID (OCID).
* `vcn_id` - The OCID of the VCN the subnet is in.
* `state` - The VCN's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created.
* `virtual_router_ip` - The IP address of the virtual router.
* `virtual_router_mac` - The MAC address of the virtual router.