---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_security_rule"
sidebar_current: "docs-oci-resource-network_firewall-network_firewall_policy_security_rule"
description: |-
  Provides the Network Firewall Policy Security Rule resource in Oracle Cloud Infrastructure Network Firewall service
---

# oci_network_firewall_network_firewall_policy_security_rule
This resource provides the Network Firewall Policy Security Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Creates a new Security Rule for the Network Firewall Policy.


## Example Usage

```hcl
resource "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule" {
	lifecycle {
		ignore_changes = [position]
	}
	#Required
	action = var.network_firewall_policy_security_rule_action
	name = var.network_firewall_policy_security_rule_name
	condition {
		application = var.network_firewall_policy_security_rule_condition_application
		destination_address = var.network_firewall_policy_security_rule_condition_destination_address
		service = var.network_firewall_policy_security_rule_condition_service
		source_address = var.network_firewall_policy_security_rule_condition_source_address
		url = var.network_firewall_policy_security_rule_condition_url
	}
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id

	#Optional
	inspection = var.network_firewall_policy_security_rule_inspection
	position {

		#Optional
		after_rule = var.network_firewall_policy_security_rule_position_after_rule
		before_rule = var.network_firewall_policy_security_rule_position_before_rule
	}
}
```

## Argument Reference

The following arguments are supported:

* `action` - (Required) (Updatable) Types of Action on the Traffic flow.
	* ALLOW - Allows the traffic.
	* DROP - Silently drops the traffic, e.g. without sending a TCP reset.
	* REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
	* INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection. 
* `condition` - (Required) (Updatable) Criteria to evaluate against network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. The resources mentioned must already be present in the policy before being referenced in the rule. 
	* `application` - (Optional) (Updatable) An array of application group names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `destination_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic destination address.
	* `service` - (Optional) (Updatable) An array of service list names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `source_address` - (Optional) (Updatable) An array of address list names to be evaluated against the traffic source address.
	* `url` - (Optional) (Updatable) An array of URL list names to be evaluated against the HTTP(S) request target.
* `inspection` - (Optional) (Updatable) Type of inspection to affect the traffic flow. This is only applicable if action is INSPECT.
	* INTRUSION_DETECTION - Intrusion Detection.
	* INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`. 
* `name` - (Required) Name for the Security rule, must be unique within the policy.
* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `position` - (Optional) (Updatable) An object which defines the position of the rule. Only one of the following position references should be provided.
	* `after_rule` - (Optional) (Updatable) Identifier for rule after which this rule lies.
	* `before_rule` - (Optional) (Updatable) Identifier for rule before which this rule lies.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `action` - Types of Action on the Traffic flow.
	* ALLOW - Allows the traffic.
	* DROP - Silently drops the traffic, e.g. without sending a TCP reset.
	* REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
	* INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection. 
* `condition` - Criteria to evaluate against network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
	* `application` - An array of application list names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `destination_address` - An array of address list names to be evaluated against the traffic destination address.
	* `service` - An array of service list names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `source_address` - An array of address list names to be evaluated against the traffic source address.
	* `url` - An array of URL list names to be evaluated against the HTTP(S) request target.
* `inspection` - Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
	* INTRUSION_DETECTION - Intrusion Detection.
	* INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`. 
* `name` - Name for the Security rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this security rule belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Network Firewall Policy Security Rule
	* `update` - (Defaults to 20 minutes), when updating the Network Firewall Policy Security Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Network Firewall Policy Security Rule


## Import

NetworkFirewallPolicySecurityRules can be imported using the `name`, e.g.

```
$ terraform import oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule "networkFirewallPolicies/{networkFirewallPolicyId}/securityRules/{securityRuleName}" 
```

