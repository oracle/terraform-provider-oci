---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_dhcp_options"
sidebar_current: "docs-oci-resource-core-dhcp_options"
description: |-
  Provides the Dhcp Options resource in Oracle Cloud Infrastructure Core service
---

# oci_core_dhcp_options
This resource provides the Dhcp Options resource in Oracle Cloud Infrastructure Core service.

Creates a new set of DHCP options for the specified VCN. For more information, see
[DhcpOptions](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpOptions/).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want the set of
DHCP options to reside. Notice that the set of options doesn't have to be in the same compartment as the VCN,
subnets, or other Networking Service components. If you're not sure which compartment to use, put the set
of DHCP options in the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the set of DHCP options, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

For more information on configuring a VCN's default DHCP options, see [Managing Default VCN Resources](/docs/providers/oci/guides/managing_default_resources.html)

## Example Usage

### VCN Local with Internet

```hcl
resource "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = var.compartment_id
	options {
        type = "DomainNameServer"
        server_type = "VcnLocalPlusInternet"
	}
	
    options {
        type = "SearchDomain"
        search_domain_names = [ "test.com" ]
    }
	
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	display_name = var.dhcp_options_display_name
}
```

### Custom DNS Server

```hcl
resource "oci_core_dhcp_options" "test_dhcp_options" {
	#Required
	compartment_id = var.compartment_id
	options {
        type = "DomainNameServer"
        server_type = "CustomDnsServer"
        custom_dns_servers = [ "192.168.0.2", "192.168.0.11", "192.168.0.19" ]
	}
	
    options {
        type = "SearchDomain"
        search_domain_names = [ "test.com" ]
    }
	
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.dhcp_options_display_name
	domain_name_type = var.dhcp_options_domain_name_type
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the set of DHCP options.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `domain_name_type` - (Optional) (Updatable) The search domain name type of DHCP options
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `options` - (Required) (Updatable) A set of DHCP options.
	* `custom_dns_servers` - (Applicable when type=DomainNameServer) (Updatable) If you set `serverType` to `CustomDnsServer`, specify the IP address of at least one DNS server of your choice (three maximum). 
	* `search_domain_names` - (Required when type=SearchDomain) (Updatable) A single search domain name according to [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). During a DNS query, the OS will append this search domain name to the value being queried.

		If you set [DhcpDnsOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpDnsOption/) to `VcnLocalPlusInternet`, and you assign a DNS label to the VCN during creation, the search domain name in the VCN's default set of DHCP options is automatically set to the VCN domain (for example, `vcn1.oraclevcn.com`).

		If you don't want to use a search domain name, omit this option from the set of DHCP options. Do not include this option with an empty list of search domain names, or with an empty string as the value for any search domain name. 
	* `server_type` - (Required when type=DomainNameServer) (Updatable) 
		* **VcnLocal:** Reserved for future use.
		* **VcnLocalPlusInternet:** Also referred to as "Internet and VCN Resolver". Instances can resolve internet hostnames (no internet gateway is required), and can resolve hostnames of instances in the VCN. This is the default value in the default set of DHCP options in the VCN. For the Internet and VCN Resolver to work across the VCN, there must also be a DNS label set for the VCN, a DNS label set for each subnet, and a hostname for each instance. The Internet and VCN Resolver also enables reverse DNS lookup, which lets you determine the hostname corresponding to the private IP address. For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
		* **CustomDnsServer:** Instances use a DNS server of your choice (three maximum). 
	* `type` - (Required) (Updatable) The specific DHCP option. Either `DomainNameServer` (for [DhcpDnsOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpDnsOption/)) or `SearchDomain` (for [DhcpSearchDomainOption](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/DhcpSearchDomainOption/)). 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the set of DHCP options belongs to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dhcp Options
	* `update` - (Defaults to 20 minutes), when updating the Dhcp Options
	* `delete` - (Defaults to 20 minutes), when destroying the Dhcp Options


## Import

DhcpOptions can be imported using the `id`, e.g.

```
$ terraform import oci_core_dhcp_options.test_dhcp_options "id"
```

