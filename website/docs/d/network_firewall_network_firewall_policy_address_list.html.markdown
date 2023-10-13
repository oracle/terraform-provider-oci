---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_address_list"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_address_list"
description: |-
  Provides details about a specific Network Firewall Policy Address List in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_address_list
This data source provides details about a specific Network Firewall Policy Address List resource in Oracle Cloud Infrastructure Network Firewall service.

Get Address List by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_address_list" "test_network_firewall_policy_address_list" {
	#Required
	address_list_name = oci_waas_address_list.test_address_list.name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `address_list_name` - (Required) Unique identifier for address lists in the scope of Network Firewall Policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `addresses` - List of addresses.
* `name` - Unique name to identify the group of addresses to be used in the policy rules.
* `parent_resource_id` - OCID of the Network Firewall Policy this Address List belongs to.
* `total_addresses` - Count of total Addresses in the AddressList
* `type` - Type of address List.

