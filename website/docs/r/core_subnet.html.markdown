---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_subnet"
sidebar_current: "docs-oci-resource-core-subnet"
description: |-
  Provides the Subnet resource in Oracle Cloud Infrastructure Core service
---

# oci_core_subnet
This resource provides the Subnet resource in Oracle Cloud Infrastructure Core service.

Creates a new subnet in the specified VCN. You can't change the size of the subnet after creation,
so it's important to think about the size of subnets you need before creating them.
For more information, see [VCNs and Subnets](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVCNs.htm).
For information on the number of subnets you can have in a VCN, see
[Service Limits](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the subnet
to reside. Notice that the subnet doesn't have to be in the same compartment as the VCN, route tables, or
other Networking Service components. If you're not sure which compartment to use, put the subnet in
the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs,
see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally associate a route table with the subnet. If you don't, the subnet will use the
VCN's default route table. For more information about route tables, see
[Route Tables](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm).

You may optionally associate a security list with the subnet. If you don't, the subnet will use the
VCN's default security list. For more information about security lists, see
[Security Lists](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securitylists.htm).

You may optionally associate a set of DHCP options with the subnet. If you don't, the subnet will use the
VCN's default set. For more information about DHCP options, see
[DHCP Options](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingDHCP.htm).

You may optionally specify a *display name* for the subnet, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

You can also add a DNS label for the subnet, which is required if you want the Internet and
VCN Resolver to resolve hostnames for instances in the subnet. For more information, see
[DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).


## Example Usage

```hcl
resource "oci_core_subnet" "test_subnet" {
	#Required
	cidr_block = "${var.subnet_cidr_block}"
	compartment_id = "${var.compartment_id}"
	security_list_ids = "${var.subnet_security_list_ids}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	availability_domain = "${var.subnet_availability_domain}"
	defined_tags = {"Operations.CostCenter"= "42"}
	dhcp_options_id = "${oci_core_dhcp_options.test_dhcp_options.id}"
	display_name = "${var.subnet_display_name}"
	dns_label = "${var.subnet_dns_label}"
	freeform_tags = {"Department"= "Finance"}
	prohibit_public_ip_on_vnic = "${var.subnet_prohibit_public_ip_on_vnic}"
	route_table_id = "${oci_core_route_table.test_route_table.id}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The availability domain to contain the subnet.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - (Required) The CIDR IP address range of the subnet.  Example: `172.16.1.0/24` 
* `compartment_id` - (Required) The OCID of the compartment to contain the subnet.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `dhcp_options_id` - (Optional) (Updatable) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet uses the VCN's default set of DHCP options. 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_label` - (Optional) A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.

	This value must be set if you want to use the Internet and VCN Resolver to resolve the hostnames of instances in the subnet. It can only be set if the VCN itself was created with a DNS label.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `subnet123` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `prohibit_public_ip_on_vnic` - (Optional) Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/)). If `prohibitPublicIpOnVnic` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).  Example: `true` 
* `route_table_id` - (Optional) (Updatable) The OCID of the route table the subnet will use. If you don't provide a value, the subnet uses the VCN's default route table. 
* `security_list_ids` - (Optional) (Updatable) The OCIDs of the security list or lists the subnet will use. If you don't provide a value, the subnet uses the VCN's default security list. Remember that security lists are associated *with the subnet*, but the rules are applied to the individual VNICs in the subnet. 
* `vcn_id` - (Required) The OCID of the VCN to contain the subnet.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The subnet's availability domain.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - The subnet's CIDR block.  Example: `172.16.1.0/24` 
* `compartment_id` - The OCID of the compartment containing the subnet.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `dhcp_options_id` - The OCID of the set of DHCP options that the subnet uses. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.

	The absence of this parameter means the Internet and VCN Resolver will not resolve hostnames of instances in this subnet.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `subnet123` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The subnet's Oracle ID (OCID).
* `prohibit_public_ip_on_vnic` - Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/CreateVnicDetails/)). If `prohibitPublicIpOnVnic` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).  Example: `true` 
* `route_table_id` - The OCID of the route table that the subnet uses.
* `security_list_ids` - The OCIDs of the security list or lists that the subnet uses. Remember that security lists are associated *with the subnet*, but the rules are applied to the individual VNICs in the subnet. 
* `state` - The subnet's current state.
* `subnet_domain_name` - The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the `oraclevcn.com` domain.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `subnet123.vcn1.oraclevcn.com` 
* `time_created` - The date and time the subnet was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the subnet is in.
* `virtual_router_ip` - The IP address of the virtual router.  Example: `10.0.14.1` 
* `virtual_router_mac` - The MAC address of the virtual router.  Example: `00:00:17:B6:4D:DD` 

## Import

Subnets can be imported using the `id`, e.g.

```
$ terraform import oci_core_subnet.test_subnet "id"
```

