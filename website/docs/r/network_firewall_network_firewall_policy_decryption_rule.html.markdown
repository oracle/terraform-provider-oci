---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_rule"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_decryption_rule"
description: |-
  Provides the Network Firewall Policy Decryption Rule resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_decryption_rule
This resource provides the Network Firewall Policy Decryption Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Decryption Rule for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_decryption_rule" "test_network_firewall_policy_decryption_rule" {
	lifecycle = {
		ignore_changes = [position]
	}
	#Required
	name = var.network_firewall_policy_decryption_rule_name
	action = var.network_firewall_policy_decryption_rule_action
	condition {
		destination_address = var.network_firewall_policy_decryption_rule_condition_destination_address
		source_address = var.network_firewall_policy_decryption_rule_condition_source_address
	}
	position {
		#Optional
		after_rule = var.network_firewall_policy_decryption_rule_position_after_rule
		before_rule = var.network_firewall_policy_decryption_rule_position_before_rule
	}
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	decryption_profile = var.network_firewall_policy_decryption_rule_decryption_profile
	secret = var.network_firewall_policy_decryption_rule_secret
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required) (Updatable) Action:
	* NO_DECRYPT - Matching traffic is not decrypted.
	* DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`. 
* `condition` - (Required) (Updatable) Match criteria used in Decryption Rule used on the firewall policy rules. The resources mentioned must already be present in the policy before being referenced in the rule.
	* `destination_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic destination address.
	* `source_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic source address.
* `decryption_profile` - (Required only when action is `DECRYPT`) (Updatable) The name of the decryption profile to use.
* `secret` - (Required only when action is `DECRYPT`) (Updatable) The name of a mapped secret. Its `type` must match that of the specified decryption profile.
* `name` - (Required) Name for the decryption rule, must be unique within the policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `position` - (Optional) (Updatable) An object which defines the position of the rule. Only one of `after_rule` or `before_rule` should be provided.
	* `after_rule` - (Optional) (Updatable) Identifier for rule after which this rule lies.
	* `before_rule` - (Optional) (Updatable) Identifier for rule before which this rule lies.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - Action:
	* NO_DECRYPT - Matching traffic is not decrypted.
	* DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`. 
* `condition` - Match criteria used in Decryption Rule used on the firewall policy rules.
	* `destination_address` - An array of address list names to be evaluated against the traffic destination address.
	* `source_address` - An array of address list names to be evaluated against the traffic source address.
* `decryption_profile` - The name of the decryption profile to use.
* `secret` - The name of a mapped secret. Its `type` must match that of the specified decryption profile.
* `name` - Name for the decryption rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this decryption rule belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Decryption Rule
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Decryption Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Decryption Rule


## Import

NetworkFirewallPolicyDecryptionRules can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_decryption_rule.test_network_firewall_policy_decryption_rule "networkFirewallPolicies/{networkFirewallPolicyId}/decryptionRules/{decryptionRuleName}" 
```

