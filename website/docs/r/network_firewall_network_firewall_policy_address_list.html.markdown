---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_address_list"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_address_list"
description: |-
  Provides the Network Firewall Policy Address List resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_address_list
This resource provides the Network Firewall Policy Address List resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Address List for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_address_list" "test_network_firewall_policy_address_list" {
	#Required
	name = var.network_firewall_policy_address_list_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	type = var.network_firewall_policy_address_list_type

	#Optional
	addresses = var.network_firewall_policy_address_list_addresses
}
```

## Argument Reference

The following arguments are supported:

* `addresses` - (Required) (Updatable) List of addresses.
* `name` - (Required) Unique name to identify the group of addresses to be used in the policy rules.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `type` - (Required) Type of address List. The accepted values are - * FQDN * IP


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `addresses` - List of addresses.
* `name` - Unique name to identify the group of addresses to be used in the policy rules.
* `parent_resource_id` - OCID of the Network Firewall Policy this Address List belongs to.
* `total_addresses` - Count of total addresses in the AddressList
* `type` - Type of address list. The accepted values are - * FQDN * IP. The usage FQDN is disabled by default. To get access to use FQDNs (only public FQDNs allowed) please contact support.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Address List
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Address List
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Address List


## Import

NetworkFirewallPolicyAddressLists can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list "networkFirewallPolicies/{networkFirewallPolicyId}/addressLists/{addressListName}" 
```

