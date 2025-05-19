---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_nat_rule"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_nat_rule"
description: |-
  Provides details about a specific Network Firewall Policy Nat Rule in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_nat_rule
This data source provides details about a specific Network Firewall Policy Nat Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Get NAT Rule by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_nat_rule" "test_network_firewall_policy_nat_rule" {
	#Required
	nat_rule_name = oci_events_rule.test_rule.name
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `nat_rule_name` - (Required) Unique identifier for NAT Rules in the network firewall policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier


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

