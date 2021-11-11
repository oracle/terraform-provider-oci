---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dhcp_options"
sidebar_current: "docs-oci-datasource-core-dhcp_options"
description: |-
  Provides the list of Dhcp Options in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_dhcp_options
This data source provides the list of Dhcp Options in Oracle Cloud Infrastructure Core service.

Lists the sets of DHCP options in the specified VCN and specified compartment.
If the VCN ID is not provided, then the list includes the sets of DHCP options from all VCNs in the specified compartment.
The response includes the default set of options that automatically comes with each VCN,
plus any other sets you've created.


## Example Usage

```hcl
data "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.dhcp_options_display_name
	state = var.dhcp_options_state
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

* `options` - The list of options.

### DhcpOptions Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the set of DHCP options.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `domain_name_type` - The search domain name type of DHCP options
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) for the set of DHCP options.
* `options` - The collection of individual DHCP options.
	* `custom_dns_servers` - If you set `serverType` to `CustomDnsServer`, specify the IP address of at least one DNS server of your choice (three maximum). 
	* `search_domain_names` - A single search domain name according to [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). During a DNS query, the OS will append this search domain name to the value being queried.

		If you set [DhcpDnsOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpDnsOption/) to `VcnLocalPlusInternet`, and you assign a DNS label to the VCN during creation, the search domain name in the VCN's default set of DHCP options is automatically set to the VCN domain (for example, `vcn1.oraclevcn.com`).

		If you don't want to use a search domain name, omit this option from the set of DHCP options. Do not include this option with an empty list of search domain names, or with an empty string as the value for any search domain name. 
	* `server_type` - 
		* **VcnLocal:** Reserved for future use.
		* **VcnLocalPlusInternet:** Also referred to as "Internet and VCN Resolver". Instances can resolve internet hostnames (no internet gateway is required), and can resolve hostnames of instances in the VCN. This is the default value in the default set of DHCP options in the VCN. For the Internet and VCN Resolver to work across the VCN, there must also be a DNS label set for the VCN, a DNS label set for each subnet, and a hostname for each instance. The Internet and VCN Resolver also enables reverse DNS lookup, which lets you determine the hostname corresponding to the private IP address. For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
		* **CustomDnsServer:** Instances use a DNS server of your choice (three maximum). 
	* `type` - The specific DHCP option. Either `DomainNameServer` (for [DhcpDnsOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpDnsOption/)) or `SearchDomain` (for [DhcpSearchDomainOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpSearchDomainOption/)). 
* `state` - The current state of the set of DHCP options.
* `time_created` - Date and time the set of DHCP options was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the set of DHCP options belongs to.

