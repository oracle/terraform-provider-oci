# baremetal\_core\_subnets

Gets a list of subnets.

## Example Usage

```
data "baremetal_identity_availability_domains" "ADs" {
  compartment_id = "${var.compartment_id}"
}

resource "baremetal_core_virtual_network" "t" {
  cidr_block     = "10.0.0.0/16"
  compartment_id = "${var.compartment_id}"
  display_name   = "network_name"
}

resource "baremetal_core_subnet" "t" {
  compartment_id = "${var.compartment_id}"

  availability_domain = "${data.baremetal_identity_availability_domains.ADs.availability_domains.0.name}"
  route_table_id      = "${baremetal_core_virtual_network.t.default_route_table_id}"
  vcn_id              = "${baremetal_core_virtual_network.t.id}"
  security_list_ids   = ["${baremetal_core_virtual_network.t.default_security_list_id}"]
  dhcp_options_id     = "${baremetal_core_virtual_network.t.default_dhcp_options_id}"

  display_name               = "display_name"
  cidr_block                 = "10.10.10.0/24"
  prohibit_public_ip_on_vnic = true
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain to contain the subnet.
* `cidr_block` - (Required) The CIDR IP address range of the subnet.
* `compartment_id` - (Required) The OCID of the compartment to contain the subnet.
* `dhcp_options_id` - (Required) The OCID of the set of DHCP options the subnet will use.
* `route_table_id` - (Required) The OCID of the route table the subnet will use.
* `security_list_ids` - (Required) OCIDs for the security lists to associate with the subnet. Remember that security lists are associated at the subnet level, but the rules are applied to the individual VNICs in the subnet.
* `vcn_id` - (Required) The OCID of the VCN to contain the subnet.

* `dns_label` - (Optional) DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (e.g., bminstance-1.subnet123.vcn1.oraclevcn.com). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed. The absence of this parameter means the Internet and VCN Resolver will not resolve hostnames of instances in this subnet.
* `display_name` - (Optional) User-friendly name. Does not have to be unique, and it's changeable.
* `prohibit_public_ip_on_vnic` - (Optional) Whether VNICs within this subnet can have public IP. If it is allowed, VNICs created in the subnet will automatically be assigned public IP unless otherwise specified in the VNIC. If it is prohibited, VNICs in the subnet cannot have public IP address assigned. The default value is false if unspecified.

WARNING: With some exceptions, changing these properties in a plan after the resources has been created results in destruction and recreation of the resources and all dependent resources. Display name can be changed non-destructively.  

## Attributes Reference

* `id` - The subnet's Oracle ID (OCID).
* `state` - The VCN's current state. [PROVISIONING, AVAILABLE, TERMINATING, TERMINATED]
* `time_created` - The date and time the VCN was created.
* `virtual_router_ip` - The IP address of the virtual router.
* `virtual_router_mac` - The MAC address of the virtual router.
