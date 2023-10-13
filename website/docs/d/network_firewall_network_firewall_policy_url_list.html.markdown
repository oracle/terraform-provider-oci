---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_url_list"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_url_list"
description: |-
  Provides details about a specific Network Firewall Policy Url List in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_url_list
This data source provides details about a specific Network Firewall Policy Url List resource in Oracle Cloud Infrastructure Network Firewall service.

Get Url List by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_url_list" "test_network_firewall_policy_url_list" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	url_list_name = var.network_firewall_policy_url_list_url_list_name
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `url_list_name` - (Required) Unique name identifier for url lists in the scope of Network Firewall Policy.


## Attributes Reference

The following attributes are exported:

* `name` - Unique name identifier for the URL list.
* `parent_resource_id` - OCID of the Network Firewall Policy this URL List belongs to.
* `total_urls` - Total count of URLs in the URL List
* `urls` - List of urls.
	* `pattern` - A string consisting of a concatenation of optional host component and optional path component. The host component may start with `*.` to match the case-insensitive domain and all its subdomains. The path component must start with a `/`, and may end with `*` to match all paths of which it is a case-sensitive prefix. A missing host component matches all request domains, and a missing path component matches all request paths. An empty value matches all requests. 
	* `type` - The type of pattern.
		* SIMPLE - A simple pattern with optional subdomain and/or path suffix wildcards. 

