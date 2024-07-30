---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_tunnel_inspection_rules"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_tunnel_inspection_rules"
description: |-
  Provides the list of Network Firewall Policy Tunnel Inspection Rules in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_tunnel_inspection_rules
This data source provides the list of Network Firewall Policy Tunnel Inspection Rules in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of tunnel inspection rules for the network firewall policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_tunnel_inspection_rules" "test_network_firewall_policy_tunnel_inspection_rules" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_tunnel_inspection_rule_display_name
	tunnel_inspection_rule_priority_order = var.network_firewall_policy_tunnel_inspection_rule_tunnel_inspection_rule_priority_order
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `tunnel_inspection_rule_priority_order` - (Optional) Unique priority order for Tunnel Inspection rules in the network firewall policy.


## Attributes Reference

The following attributes are exported:

* `tunnel_inspection_rule_summary_collection` - The list of tunnel_inspection_rule_summary_collection.

### NetworkFirewallPolicyTunnelInspectionRule Reference

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

