---
layout: "oci"
page_title: "OCI: oci_core_security_lists"
sidebar_current: "docs-oci-datasource-core-security_lists"
description: |-
Provides a list of SecurityLists
---
# Data Source: oci_core_security_lists
The SecurityLists data source allows access to the list of OCI security_lists

Lists the security lists in the specified VCN and compartment.


## Example Usage

```hcl
data "oci_core_security_lists" "test_security_lists" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.security_list_display_name}"
	state = "${var.security_list_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


## Attributes Reference

The following attributes are exported:

* `security_lists` - The list of security_lists.

### SecurityList Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `egress_security_rules` - Rules for allowing egress IP packets.
	* `destination` - The destination service cidrBlock or destination IP address range in CIDR notation for the egress rule. This is the range of IP addresses that a packet originating from the instance can go to. 
	* `destination_type` - Type of destination for EgressSecurityRule. SERVICE_CIDR_BLOCK should be used if destination is a service cidrBlock. CIDR_BLOCK should be used if destination is IP address range in CIDR notation. It defaults to CIDR_BLOCK, if not specified. 
	* `icmp_options` - Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `protocol` - The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), and UDP ("17"). 
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The security list's Oracle Cloud ID (OCID).
* `ingress_security_rules` - Rules for allowing ingress IP packets.
	* `icmp_options` - Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
		* `type` - The ICMP type.
	* `protocol` - The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), and UDP ("17"). 
	* `source` - The source service cidrBlock or source IP address range in CIDR notation for the ingress rule. This is the range of IP addresses that a packet coming into the instance can come from.  Examples: `10.12.0.0/16`           `oci-phx-objectstorage` 
	* `source_type` - Type of source for IngressSecurityRule. SERVICE_CIDR_BLOCK should be used if source is a service cidrBlock. CIDR_BLOCK should be used if source is IP address range in CIDR notation. It defaults to CIDR_BLOCK, if not specified. 
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if ingress traffic allows TCP destination port 80, there should be an egress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - The minimum port number. Must not be greater than the maximum port number.
* `state` - The security list's current state.
* `time_created` - The date and time the security list was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the security list belongs to.

