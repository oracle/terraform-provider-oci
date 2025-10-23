---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_nat_rule"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_nat_rule"
description: |-
  Provides the Network Firewall Policy Nat Rule resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_nat_rule
This resource provides the Network Firewall Policy Nat Rule resource in Oracle Cloud Infrastructure Network Firewall service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/network-firewall/latest/NetworkFirewallPolicyNatRule

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/network_firewall

Creates a new NAT Rule for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule" {
	#Required
	action = var.network_firewall_policy_nat_rule_action
	condition {

		#Optional
		destination_address = var.network_firewall_policy_nat_rule_condition_destination_address
		service = var.network_firewall_policy_nat_rule_condition_service
		source_address = var.network_firewall_policy_nat_rule_condition_source_address
	}
	name = var.network_firewall_policy_nat_rule_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	type = var.network_firewall_policy_nat_rule_type

	#Optional
	description = var.network_firewall_policy_nat_rule_description
	position {

		#Optional
		after_rule = var.network_firewall_policy_nat_rule_position_after_rule
		before_rule = var.network_firewall_policy_nat_rule_position_before_rule
	}
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required) (Updatable) action:
	* DIPP_SRC_NAT - Dynamic-ip-port source NAT. 
* `condition` - (Required) (Updatable) Match criteria used in NAT Rule used on the firewall policy.
	* `destination_address` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic destination address.
	* `service` - (Optional) (Updatable) A Service name to be evaluated against the traffic protocol and protocol-specific parameters.
	* `source_address` - (Optional) (Updatable) An array of IP address list names to be evaluated against the traffic source address.
* `description` - (Optional) (Updatable) Description of a NAT rule. This field can be used to add additional info.
* `name` - (Required) Name for the NAT rule, must be unique within the policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `position` - (Optional) (Updatable) An object which defines the position of the rule.
	* `after_rule` - (Optional) (Updatable) Identifier for rule after which this rule lies.
	* `before_rule` - (Optional) (Updatable) Identifier for rule before which this rule lies.
* `type` - (Required) (Updatable) NAT type:
	* NATV4 - NATV4 type NAT. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - action:
	* DIPP_SRC_NAT - Dynamic-ip-port source NAT. 
* `condition` - Match criteria used in NAT Rule used on the firewall policy.
	* `destination_address` - An array of IP address list names to be evaluated against the traffic destination address.
	* `service` - A Service name to be evaluated against the traffic protocol and protocol-specific parameters.
	* `source_address` - An array of IP address list names to be evaluated against the traffic source address.
* `description` - Description of a NAT rule. This field can be used to add additional info.
* `name` - Name for the NAT rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this decryption profile belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.
* `priority_order` - The priority order in which this rule should be evaluated
* `type` - NAT type:
	* NATV4 - NATV4 type NAT. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Nat Rule
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Nat Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Nat Rule


## Import

NetworkFirewallPolicyNatRules can be imported using the `id`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule "networkFirewallPolicies/{networkFirewallPolicyId}/natRules/{natRuleName}" 
```

