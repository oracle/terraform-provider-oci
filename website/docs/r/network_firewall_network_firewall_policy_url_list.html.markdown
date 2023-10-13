---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_url_list"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_url_list"
description: |-
  Provides the Network Firewall Policy Url List resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_url_list
This resource provides the Network Firewall Policy Url List resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Url List for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_url_list" "test_network_firewall_policy_url_list" {
	#Required
	name = var.network_firewall_policy_url_list_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	urls {
		#Required
		pattern = var.network_firewall_policy_url_list_urls_pattern
		type = var.network_firewall_policy_url_list_urls_type
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique name to identify the group of urls to be used in the policy rules.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `urls` - (Required) (Updatable) List of urls.
	* `pattern` - (Required) (Updatable) A string consisting of a concatenation of optional host component and optional path component. The host component may start with `*.` to match the case-insensitive domain and all its subdomains. The path component must start with a `/`, and may end with `*` to match all paths of which it is a case-sensitive prefix. A missing host component matches all request domains, and a missing path component matches all request paths. An empty value matches all requests. 
	* `type` - (Required) (Updatable) The type of pattern.
		* SIMPLE - The only accepted value is `SIMPLE`. A simple pattern with optional subdomain and/or path suffix wildcards. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `name` - Unique name identifier for the URL list.
* `parent_resource_id` - OCID of the Network Firewall Policy this URL List belongs to.
* `total_urls` - Total count of URLs in the URL List
* `urls` - List of urls.
	* `pattern` - A string consisting of a concatenation of optional host component and optional path component. The host component may start with `*.` to match the case-insensitive domain and all its subdomains. The path component must start with a `/`, and may end with `*` to match all paths of which it is a case-sensitive prefix. A missing host component matches all request domains, and a missing path component matches all request paths. An empty value matches all requests. 
	* `type` - The type of pattern.
		* SIMPLE - A simple pattern with optional subdomain and/or path suffix wildcards. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Url List
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Url List
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Url List


## Import

NetworkFirewallPolicyUrlLists can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list "networkFirewallPolicies/{networkFirewallPolicyId}/urlLists/{urlListName}" 
```

