---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_url_lists"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_url_lists"
description: |-
  Provides the list of Network Firewall Policy Url Lists in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_url_lists
This data source provides the list of Network Firewall Policy Url Lists in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of URL lists for the Network Firewall Policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_url_lists" "test_network_firewall_policy_url_lists" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_url_list_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `url_list_summary_collection` - The list of url_list_summary_collection.

### NetworkFirewallPolicyUrlList Reference

The following attributes are exported:

* `name` - Unique name identifier for the URL list.
* `parent_resource_id` - OCID of the Network Firewall Policy this URL List belongs to.
* `total_urls` - Total count of URLs in the URL List
* `urls` - List of urls.
	* `pattern` - A string consisting of a concatenation of optional host component and optional path component. The host component may start with `*.` to match the case-insensitive domain and all its subdomains. The path component must start with a `/`, and may end with `*` to match all paths of which it is a case-sensitive prefix. A missing host component matches all request domains, and a missing path component matches all request paths. An empty value matches all requests. 
	* `type` - The type of pattern.
		* SIMPLE - A simple pattern with optional subdomain and/or path suffix wildcards. 

