---
subcategory: "Network Firewall"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_network_firewall_network_firewall_policy_security_rule"
sidebar_current: "docs-oci-datasource-network_firewall-network_firewall_policy_security_rule"
description: |-
  Provides details about a specific Network Firewall Policy Security Rule in Oracle Cloud Infrastructure Network Firewall service
---

# Data Source: oci_network_firewall_network_firewall_policy_security_rule
This data source provides details about a specific Network Firewall Policy Security Rule resource in Oracle Cloud Infrastructure Network Firewall service.

Get Security Rule by the given name in the context of network firewall policy.

## Example Usage

```hcl
data "oci_network_firewall_network_firewall_policy_security_rule" "test_network_firewall_policy_security_rule" {
	#Required
	network_firewall_policy_id = oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id
	security_rule_name = var.oci_network_firewall_network_firewall_policy_security_rule_name
}
```

## Argument Reference

The following arguments are supported:

* `network_firewall_policy_id` - (Required) Unique Network Firewall Policy identifier
* `security_rule_name` - (Required) Unique identifier for Security Rules in the scope of Network Firewall Policy.


## Attributes Reference

The following attributes are exported:

* `action` - Types of Action on the Traffic flow.
	* ALLOW - Allows the traffic.
	* DROP - Silently drops the traffic, e.g. without sending a TCP reset.
	* REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
	* INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection. 
* `condition` - Criteria to evaluate against network traffic. A match occurs when at least one item in the array associated with each specified property corresponds with the relevant aspect of the traffic. 
	* `application` - An array of application list names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `destination_address` - An array of IP address list names to be evaluated against the traffic destination address.
	* `service` - An array of service list names to be evaluated against the traffic protocol and protocol-specific parameters.
	* `source_address` - An array of IP address list names to be evaluated against the traffic source address.
	* `url` - An array of URL pattern list names to be evaluated against the HTTP(S) request target.
* `inspection` - Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
	* INTRUSION_DETECTION - Intrusion Detection.
	* INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`. 
* `name` - Name for the Security rule, must be unique within the policy.
* `parent_resource_id` - OCID of the Network Firewall Policy this security rule belongs to.
* `position` - An object which defines the position of the rule.
	* `after_rule` - Identifier for rule after which this rule lies.
	* `before_rule` - Identifier for rule before which this rule lies.

