---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_decryption_rule"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_decryption_rule"
description: |-
  Provides details about a specific Network Firewall Policy Decryption Rule in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_decryption_rule
This data source provides details about a specific Network Firewall Policy Decryption Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Get Decryption Rule by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_decryption_rule" "test_network_firewall_policy_decryption_rule" {
	#Required
	decryption_rule_name = var.oci_network_firewall_network_firewall_policy_decryption_rule_name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `decryption_rule_name` - (Required) Unique identifier for Decryption Rules in the scope of Network Firewall Policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

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

