---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_address_lists"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_address_lists"
description: |-
  Provides the list of Network Firewall Policy Address Lists in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_address_lists
This data source provides the list of Network Firewall Policy Address Lists in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Network Firewall Policies.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_address_lists" "test_network_firewall_policy_address_lists" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_address_list_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `address_list_summary_collection` - The list of address_list_summary_collection.

### NetworkFirewallPolicyAddressList Reference

The following attributes are exported:

* `addresses` - List of addresses.
* `name` - Unique name to identify the group of addresses to be used in the policy rules.
* `parent_resource_id` - OCID of the Network Firewall Policy this Address List belongs to.
* `total_addresses` - Count of total Addresses in the AddressList
* `type` - Type of address list.

