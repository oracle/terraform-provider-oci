---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_services"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_services"
description: |-
  Provides the list of Network Firewall Policy Services in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_services
This data source provides the list of Network Firewall Policy Services in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Services for the policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_services" "test_network_firewall_policy_services" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_service_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `service_summary_collection` - The list of service_summary_collection.

### NetworkFirewallPolicyService Reference

The following attributes are exported:

* `name` - Name of the service.
* `parent_resource_id` - OCID of the Network Firewall Policy this service belongs to.
* `port_ranges` - List of port-ranges used.
	* `maximum_port` - The maximum port in the range (inclusive), which may be absent for a single-port range.
	* `minimum_port` - The minimum port in the range (inclusive), or the sole port of a single-port range.
* `type` - Describes the type of Service.

