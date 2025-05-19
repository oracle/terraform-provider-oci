---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_nat_rules"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_nat_rules"
description: |-
  Provides the list of Network Firewall Policy Nat Rules in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_nat_rules
This data source provides the list of Network Firewall Policy Nat Rules in Oracle Cloud Infrastructure Network Firewall service.

Returns a list of NAT Rules for the Network Firewall Policy.


## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_nat_rules" "test_network_firewall_policy_nat_rules" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	display_name = var.network_firewall_policy_nat_rule_display_name
	nat_rule_priority_order = var.network_firewall_policy_nat_rule_nat_rule_priority_order
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `nat_rule_priority_order` - (Optional) Unique priority order for NAT Rules in the network firewall policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


## Attributes Reference

The following attributes are exported:

* `nat_rule_collection` - The list of nat_rule_collection.

### NetworkFirewallPolicyNatRule Reference

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

