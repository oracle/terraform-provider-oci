---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_tunnel_inspection_rule"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_tunnel_inspection_rule"
description: |-
  Provides the Network Firewall Policy Tunnel Inspection Rule resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_tunnel_inspection_rule
This resource provides the Network Firewall Policy Tunnel Inspection Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new tunnel inspection rule for the network firewall policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" "test_network_firewall_policy_tunnel_inspection_rule" {
	#Required
	condition {

		#Optional
		destination_address = var.network_firewall_policy_tunnel_inspection_rule_condition_destination_address
		source_address = var.network_firewall_policy_tunnel_inspection_rule_condition_source_address
	}
	name = var.network_firewall_policy_tunnel_inspection_rule_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	protocol = var.network_firewall_policy_tunnel_inspection_rule_protocol

	#Optional
	action = var.network_firewall_policy_tunnel_inspection_rule_action
	position {

		#Optional
		after_rule = var.network_firewall_policy_tunnel_inspection_rule_position_after_rule
		before_rule = var.network_firewall_policy_tunnel_inspection_rule_position_before_rule
	}
	profile {

		#Optional
		must_return_traffic_to_source = var.network_firewall_policy_tunnel_inspection_rule_profile_must_return_traffic_to_source
	}
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Optional) (Updatable) Types of Inspect Action on the traffic flow.
	* INSPECT - Inspect the traffic.
	* INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic. 
* `condition` - (Required) (Updatable) Criteria to evaluate against incoming network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
	* `destination_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic destination address.
	* `source_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic source address.
* `name` - (Required) Name for the Tunnel Inspection Rule, must be unique within the policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `position` - (Optional) (Updatable) An object which defines the position of the rule.
	* `after_rule` - (Optional) (Updatable) Identifier for rule after which this rule lies.
	* `before_rule` - (Optional) (Updatable) Identifier for rule before which this rule lies.
* `profile` - (Optional) (Updatable) Vxlan Inspect profile used in Vxlan Tunnel Inspection Rules. 
	* `must_return_traffic_to_source` - (Optional) (Updatable) Return scanned VXLAN tunnel traffic to source.
* `protocol` - (Required) (Updatable) Types of Tunnel Inspection Protocol to be applied on the traffic.
	* VXLAN - VXLAN Tunnel Inspection Protocol will be applied on the traffic. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Tunnel Inspection Rule
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Tunnel Inspection Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Tunnel Inspection Rule


## Import

NetworkFirewallPolicyTunnelInspectionRules can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_tunnel_inspection_rule.test_network_firewall_policy_tunnel_inspection_rule "networkFirewallPolicies/{networkFirewallPolicyId}/tunnelInspectionRules/{tunnelInspectionRuleName}" 
```

