---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_subnets"
sidebar_current: "docs-oci-datasource-core-subnets"
description: |-
  Provides the list of Subnets in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_subnets
This data source provides the list of Subnets in Oracle Cloud Infrastructure Core service.

Lists the subnets in the specified VCN and the specified compartment.
If the VCN ID is not provided, then the list includes the subnets from all VCNs in the specified compartment.


## Example Usage

```hcl
data "oci_core_subnets" "test_subnets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.subnet_display_name
	state = var.subnet_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `subnets` - The list of subnets.

### Subnet Reference

The following attributes are exported:

* `availability_domain` - The subnet's availability domain. This attribute will be null if this is a regional subnet instead of an AD-specific subnet. Oracle recommends creating regional subnets.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - The subnet's CIDR block.  Example: `10.0.1.0/24` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the subnet.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `dhcp_options_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the set of DHCP options that the subnet uses. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.

	The absence of this parameter means the Internet and VCN Resolver will not resolve hostnames of instances in this subnet.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `subnet123` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The subnet's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ipv6cidr_block` - For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's IP address space. The subnet size is always /64. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).  Example: `2001:0db8:0123:1111::/64` 
* `ipv6virtual_router_ip` - For an IPv6-enabled subnet, this is the IPv6 address of the virtual router.  Example: `2001:0db8:0123:1111:89ab:cdef:1234:5678` 
* `prohibit_internet_ingress` - Whether to disallow ingress internet traffic to VNICs within this subnet. Defaults to false.

	For IPV4, `prohibitInternetIngress` behaves similarly to `prohibitPublicIpOnVnic`. If it is set to false, VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/)). If `prohibitInternetIngress` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a privatesubnet).

	For IPv6, if `prohibitInternetIngress` is set to `true`, internet access is not allowed for any IPv6s assigned to VNICs in the subnet. Otherwise, ingress internet traffic is allowed by default.

	Example: `true` 
* `prohibit_public_ip_on_vnic` - Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CreateVnicDetails/)). If `prohibitPublicIpOnVnic` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).  Example: `true` 
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table that the subnet uses.
* `security_list_ids` - The OCIDs of the security list or lists that the subnet uses. Remember that security lists are associated *with the subnet*, but the rules are applied to the individual VNICs in the subnet. 
* `state` - The subnet's current state.
* `subnet_domain_name` - The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the `oraclevcn.com` domain.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `subnet123.vcn1.oraclevcn.com` 
* `time_created` - The date and time the subnet was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the subnet is in.
* `virtual_router_ip` - The IP address of the virtual router.  Example: `10.0.14.1` 
* `virtual_router_mac` - The MAC address of the virtual router.  Example: `00:00:00:00:00:01` 

