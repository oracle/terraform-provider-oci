---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_security_lists"
sidebar_current: "docs-oci-datasource-core-security_lists"
description: |-
  Provides the list of Security Lists in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_security_lists
This data source provides the list of Security Lists in Oracle Cloud Infrastructure Core service.

Lists the security lists in the specified VCN and compartment.
If the VCN ID is not provided, then the list includes the security lists from all VCNs in the specified compartment.


## Example Usage

```hcl
data "oci_core_security_lists" "test_security_lists" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.security_list_display_name
	state = var.security_list_state
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
* `vcn_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN.


## Attributes Reference

The following attributes are exported:

* `security_lists` - The list of security_lists.

### SecurityList Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `egress_security_rules` - Rules for allowing egress IP packets.
	* `description` - An optional description of your choice for the rule. 
	* `destination` - Conceptually, this is the range of IP addresses that a packet originating from the instance can go to.

		Allowed values:
		* IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56` Note that IPv6 addressing is currently supported only in certain regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
		* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a security list rule for traffic destined for a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`. 
	* `destination_type` - Type of destination for the rule. The default is `CIDR_BLOCK`.

		Allowed values:
		* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
		* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway). 
	* `icmp_options` - Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
		* [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
		* [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)

		If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `protocol` - The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58"). 
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The security list's Oracle Cloud ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ingress_security_rules` - Rules for allowing ingress IP packets.
	* `description` - An optional description of your choice for the rule. 
	* `icmp_options` - Optional and valid only for ICMP and ICMPv6. Use to specify a particular ICMP type and code as defined in:
		* [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml)
		* [ICMPv6 Parameters](https://www.iana.org/assignments/icmpv6-parameters/icmpv6-parameters.xhtml)

		If you specify ICMP or ICMPv6 as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic via IPv4, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `protocol` - The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), UDP ("17"), and ICMPv6 ("58"). 
	* `source` - Conceptually, this is the range of IP addresses that a packet coming into the instance can come from.

		Allowed values:
		* IP address range in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`. IPv6 addressing is supported for all commercial and government regions. See [IPv6 Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
		* The `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/), if you're setting up a security list rule for traffic coming from a particular `Service` through a service gateway. For example: `oci-phx-objectstorage`. 
	* `source_type` - Type of source for the rule. The default is `CIDR_BLOCK`.
		* `CIDR_BLOCK`: If the rule's `source` is an IP address range in CIDR notation.
		* `SERVICE_CIDR_BLOCK`: If the rule's `source` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic coming from a particular `Service` through a service gateway). 
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if ingress traffic allows TCP destination port 80, there should be an egress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
* `state` - The security list's current state.
* `time_created` - The date and time the security list was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z`
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the security list belongs to.
