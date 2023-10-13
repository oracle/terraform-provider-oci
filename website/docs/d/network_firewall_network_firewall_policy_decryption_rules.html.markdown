---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_rules"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_decryption_rules"
description: |-
  Provides the list of Network Firewall Policy Decryption Rules in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_decryption_rules
This data source provides the list of Network Firewall Policy Decryption Rules in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of Decryption Rule for the Network Firewall Policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_decryption_rules" "test_network_firewall_policy_decryption_rules" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional, only one of the following should be provided
	decryption_rule_priority_order = var.network_firewall_policy_decryption_rule_decryption_rule_priority_order
	display_name = var.network_firewall_policy_decryption_rule_display_name
}
```

## Argument Reference

The following arguments are supported:

* `decryption_rule_priority_order` - (Optional) Unique priority order for Decryption Rules in the network firewall policy.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `decryption_rule_summary_collection` - The list of decryption_rule_summary_collection.

### NetworkFirewallPolicyDecryptionRule Reference

The following attributes are exported:

* `action` - Action:
	* NO_DECRYPT - Matching traffic is not decrypted.
	* DECRYPT - Matching traffic is decrypted with the specified `secret` according to the specified `decryptionProfile`. 
* `condition` - Match criteria used in Decryption Rule used on the firewall policy rules.
	* `destination_address` - An array of IP address list names to be evaluated against the traffic destination address.
	* `source_address` - An array of IP address list names to be evaluated against the traffic source address.
* `decryption_profile` - The name of the decryption profile to use.
* `name` - Name for the decryption rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this decryption rule belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.
* `secret` - The name of a mapped secret. Its `type` must match that of the specified decryption profile.

