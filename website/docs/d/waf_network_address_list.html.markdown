---
subcategory: "Waf"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waf_network_address_list"
sidebar_current: "docs-oci-datasource-waf-network_address_list"
description: |-
  Provides details about a specific Network Address List in Oracle Cloud Infrastructure Waf service
---

# Data Source: oci_waf_network_address_list
This data source provides details about a specific Network Address List resource in Oracle Cloud Infrastructure Waf service.

Gets a NetworkAddressList by OCID.

## Example Usage

```hcl
data "oci_waf_network_address_list" "test_network_address_list" {
	#Required
	network_address_list_id = oci_waf_network_address_list.test_network_address_list.id
}
```

## Argument Reference

The following arguments are supported:

* `network_address_list_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.


## Attributes Reference

The following attributes are exported:

* `addresses` - A list of IP address prefixes in CIDR notation. To specify all addresses, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - NetworkAddressList display name, can be renamed.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NetworkAddressList.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in FAILED state. 
* `state` - The current state of the NetworkAddressList.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the NetworkAddressList was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the NetworkAddressList was updated. An RFC3339 formatted datetime string.
* `type` - Type of NetworkAddressList.
* `vcn_addresses` - A list of private address prefixes, each associated with a particular VCN. To specify all addresses in a VCN, use "0.0.0.0/0" for IPv4 and "::/0" for IPv6. 
	* `addresses` - A private IP address or CIDR IP address range.
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.

