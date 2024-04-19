---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vcn"
sidebar_current: "docs-oci-datasource-core-vcn"
description: |-
  Provides details about a specific Vcn in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vcn
This data source provides details about a specific Vcn resource in Oracle Cloud Infrastructure Core service.

Gets the specified VCN's information.

## Example Usage

```hcl
data "oci_core_vcn" "test_vcn" {
	#Required
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `vcn_id` - (Required) Specify the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `byoipv6cidr_blocks` - The list of BYOIPv6 prefixes required to create a VCN that uses BYOIPv6 ranges. 
* `cidr_block` - Deprecated. The first CIDR IP address from cidrBlocks.  Example: `172.16.0.0/16` 
* `cidr_blocks` - The list of IPv4 CIDR blocks the VCN will use. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the VCN.
* `default_dhcp_options_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default set of DHCP options. 
* `default_route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default route table.
* `default_security_list_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the VCN's default security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter. The value cannot be changed.

	The absence of this parameter means the Internet and VCN Resolver will not work for this VCN.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The VCN's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ipv6cidr_blocks` - For an IPv6-enabled VCN, this is the list of IPv6 prefixes for the VCN's IP address space. The prefixes are provided by Oracle and the sizes are always /56. 
* `ipv6private_cidr_blocks` - For an IPv6-enabled VCN, this is the list of Private IPv6 prefixes for the VCN's IP address space. 
* `state` - The VCN's current state.
* `time_created` - The date and time the VCN was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_domain_name` - The VCN's domain name, which consists of the VCN's DNS label, and the `oraclevcn.com` domain.

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `vcn1.oraclevcn.com` 

