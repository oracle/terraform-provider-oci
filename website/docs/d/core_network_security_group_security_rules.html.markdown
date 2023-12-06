---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_network_security_group_security_rules"
sidebar_current: "docs-oci-datasource-core-network_security_group_security_rules"
description: |-
  Provides the list of Network Security Group Security Rules in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_network_security_group_security_rules
This data source provides the list of Network Security Group Security Rules in Oracle Cloud Infrastructure Core service.

Lists the security rules in the specified network security group.


## Example Usage

```hcl
data "oci_core_network_security_group_security_rules" "test_network_security_group_security_rules" {
	#Required
	network_security_group_id = oci_core_network_security_group.test_network_security_group.id

	#Optional
	direction = var.network_security_group_security_rule_direction
}
```

## Argument Reference

The following arguments are supported:

* `direction` - (Optional) Direction of the security rule. Set to `EGRESS` for rules that allow outbound IP packets, or `INGRESS` for rules that allow inbound IP packets. 
* `network_security_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.


## Attributes Reference

The following attributes are exported:

* `security_rules` - The list of security_rules.

### NetworkSecurityGroupSecurityRule Reference

The following attributes are exported:

* `description` - An optional description of your choice for the rule. 
* `destination` - Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.

	Allowed values:
	* An IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56` IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a security rule for traffic destined for a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`.
	* The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/) in the same VCN. The value can be the NSG that the rule belongs to if the rule's intent is to control traffic between VNICs in the same NSG. 
* `destination_type` - Type of destination for the rule. Required if `direction` = `EGRESS`.

	Allowed values:
	* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway).
	* `NETWORK_SECURITY_GROUP`: If the rule's `destination` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
* `direction` - Direction of the security rule. Set to `EGRESS` for rules to allow outbound IP packets, or `INGRESS` for rules to allow inbound IP packets. 
* `icmp_options` - Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
	* [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
	* [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)

	If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
	* `code` - The ICMP code (optional).
	* `type` - The ICMP type.
* `id` - An Oracle-assigned identifier for the security rule. You specify this ID when you want to update or delete the rule.  Example: `04ABEC` 
* `is_valid` - Whether the rule is valid. The value is `True` when the rule is first created. If the rule's `source` or `destination` is a network security group, the value changes to `False` if that network security group is deleted. 
* `protocol` - The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58"). 
* `source` - Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.

	Allowed values:
	* An IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56` IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a security rule for traffic coming from a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`.
	* The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/) in the same VCN. The value can be the NSG that the rule belongs to if the rule's intent is to control traffic between VNICs in the same NSG. 
* `source_type` - Type of source for the rule. Required if `direction` = `INGRESS`.
	* `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
	* `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic coming from a particular `Service` through a service gateway).
	* `NETWORK_SECURITY_GROUP`: If the rule's `source` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroup/). 
* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
	* `destination_port_range` - 
		* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
		* `min` - The minimum port number. Must not be greater than the maximum port number. 
	* `source_port_range` - 
		* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
		* `min` - The minimum port number. Must not be greater than the maximum port number. 
* `time_created` - The date and time the security rule was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
	* `destination_port_range` - 
		* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
		* `min` - The minimum port number. Must not be greater than the maximum port number. 
	* `source_port_range` - 
		* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
		* `min` - The minimum port number. Must not be greater than the maximum port number. 