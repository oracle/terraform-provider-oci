---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vcn"
sidebar_current: "docs-oci-resource-core-vcn"
description: |-
  Provides the Vcn resource in Oracle Cloud Infrastructure Core service
---

# oci_core_vcn
This resource provides the Vcn resource in Oracle Cloud Infrastructure Core service.

The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
For managing these resources, see [Managing Default VCN Resources](/docs/providers/oci/guides/managing_default_resources.html)

Creates a new Virtual Cloud Network (VCN). For more information, see
[VCNs and Subnets](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVCNs.htm).

For the VCN, you specify a list of one or more IPv4 CIDR blocks that meet the following criteria:

- The CIDR blocks must be valid.
- They must not overlap with each other or with the on-premises network CIDR block. 
- The number of CIDR blocks does not exceed the limit of CIDR blocks allowed per VCN.

For a CIDR block, Oracle recommends that you use one of the private IP address ranges specified in [RFC 1918](https://tools.ietf.org/html/rfc1918) (10.0.0.0/8, 172.16/12, and 192.168/16). Example:
172.16.0.0/16. The CIDR blocks can range from /16 to /30.

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want the VCN to
reside. Consult an Oracle Cloud Infrastructure administrator in your organization if you're not sure which
compartment to use. Notice that the VCN doesn't have to be in the same compartment as the subnets or other
Networking Service components. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the VCN, otherwise a default is provided. It does not have to
be unique, and you can change it. Avoid entering confidential information.

You can also add a DNS label for the VCN, which is required if you want the instances to use the
Interent and VCN Resolver option for DNS in the VCN. For more information, see
[DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

The VCN automatically comes with a default route table, default security list, and default set of DHCP options.
The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for each is returned in the response. You can't delete these default objects, but you can change their
contents (that is, change the route rules, security list rules, and so on).

The VCN and subnets you create are not accessible until you attach an internet gateway or set up a Site-to-Site VPN
or FastConnect. For more information, see
[Overview of the Networking Service](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm).

## Supported Aliases

* `oci_core_virtual_network`

## Example Usage

```hcl
resource "oci_core_vcn" "test_vcn" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cidr_block = var.vcn_cidr_block
	cidr_blocks = var.vcn_cidr_blocks
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.vcn_display_name
	dns_label = var.vcn_dns_label
	freeform_tags = {"Department"= "Finance"}
	is_ipv6enabled = var.vcn_is_ipv6enabled
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Optional) **Deprecated.** Do *not* set this value. Use `cidrBlocks` instead. Example: `10.0.0.0/16` 
* `cidr_blocks` - (Optional) (Updatable) The list of one or more IPv4 CIDR blocks for the VCN that meet the following criteria:
	* The CIDR blocks must be valid.
	* They must not overlap with each other or with the on-premises network CIDR block.
	* The number of CIDR blocks must not exceed the limit of CIDR blocks allowed per VCN. It is an error to set both cidrBlock and cidrBlocks. Note: cidr_blocks update must be restricted to one operation at a time (either add/remove or modify one single cidr_block) or the operation will be declined. new cidr_block to be added must be placed at the end of the list. Once you migrate to using `cidr_blocks` from `cidr_block`, you will not be able to switch back.
	**Important:** Do *not* specify a value for `cidrBlock`. Use this parameter instead. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the VCN.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_label` - (Optional) A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Not required to be unique, but it's a best practice to set unique DNS labels for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter. The value cannot be changed.

	You must set this value if you want instances to be able to use hostnames to resolve other instances in the VCN. Otherwise the Internet and VCN Resolver will not work.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_ipv6enabled` - (Optional) Whether IPv6 is enabled for the VCN. Default is `false`. If enabled, Oracle will assign the VCN a IPv6 /56 CIDR block. For important details about IPv6 addressing in a VCN, see [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).  Example: `true` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cidr_block` - Deprecated. The first CIDR IP address from cidrBlocks.  Example: `172.16.0.0/16` 
* `cidr_blocks` - The list of IPv4 CIDR blocks the VCN will use. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VCN.
* `default_dhcp_options_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default set of DHCP options. 
* `default_route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default route table.
* `default_security_list_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter. The value cannot be changed.

	The absence of this parameter means the Internet and VCN Resolver will not work for this VCN.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The VCN's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ipv6cidr_blocks` - For an IPv6-enabled VCN, this is the list of IPv6 CIDR blocks for the VCN's IP address space. The CIDRs are provided by Oracle and the sizes are always /56. 
* `state` - The VCN's current state.
* `time_created` - The date and time the VCN was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_domain_name` - The VCN's domain name, which consists of the VCN's DNS label, and the `oraclevcn.com` domain.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1.oraclevcn.com` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vcn
	* `update` - (Defaults to 20 minutes), when updating the Vcn
	* `delete` - (Defaults to 20 minutes), when destroying the Vcn


## Import

Vcns can be imported using the `id`, e.g.

```
$ terraform import oci_core_vcn.test_vcn "id"
```
