---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_tunnel_inspection_rule"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_tunnel_inspection_rule"
description: |-
  Provides details about a specific Network Firewall Policy Tunnel Inspection Rule in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_tunnel_inspection_rule
This data source provides details about a specific Network Firewall Policy Tunnel Inspection Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Get tunnel inspection rule by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" "test_network_firewall_policy_tunnel_inspection_rule" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	tunnel_inspection_rule_name = oci_events_rule.test_rule.name
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `tunnel_inspection_rule_name` - (Required) Unique identifier for Tunnel Inspection Rules in the network firewall policy.


## Attributes Reference

The following attributes are exported:

* `action` - Types of Inspect Action on the Traffic flow.
	* INSPECT - Inspect the traffic.
	* INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic. 
* `condition` - Criteria to evaluate against incoming network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
	* `destination_address` - An array of address list names to be evaluated against the traffic destination address.
	* `source_address` - An array of address list names to be evaluated against the traffic source address.
* `name` - Name for the Tunnel Inspection Rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this Tunnel Inspection Rule belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.
* `priority_order` - The priority order in which this rule should be evaluated
* `profile` - Vxlan Inspect profile used in Vxlan Tunnel Inspection Rules. 
	* `must_return_traffic_to_source` - Return scanned VXLAN tunnel traffic to source.
* `protocol` - Types of Tunnel Inspection Protocol to be applied on the traffic.
	* VXLAN - VXLAN Tunnel Inspection Protocol will be applied on the traffic. 

