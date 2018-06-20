# oci_core_security_list

## SecurityList Resource

For more information on configuring a VCN's default security list, 
see [Managing Default VCN Resources](https://github.com/oracle/terraform-provider-oci/blob/master/docs/Managing%20Default%20Resources.md)

### SecurityList Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the security list.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `egress_security_rules` - Rules for allowing egress IP packets.
	* `destination` - The destination CIDR block for the egress rule. This is the range of IP addresses that a packet originating from the instance can go to. 
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
	* `source` - The source CIDR block for the ingress rule. This is the range of IP addresses that a packet coming into the instance can come from. 
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



### Create Operation
Creates a new security list for the specified VCN. For more information
about security lists, see [Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).
For information on the number of rules you can have in a security list, see
[Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the security
list to reside. Notice that the security list doesn't have to be in the same compartment as the VCN, subnets,
or other Networking Service components. If you're not sure which compartment to use, put the security
list in the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs, see
[Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the security list, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the security list.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `egress_security_rules` - (Optional) Rules for allowing egress IP packets.
	* `destination` - (Required) The destination CIDR block for the egress rule. This is the range of IP addresses that a packet originating from the instance can go to. 
	* `icmp_options` - (Optional) Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - (Optional) The ICMP code (optional).
		* `type` - (Required) The ICMP type.
	* `protocol` - (Required) The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), and UDP ("17"). 
	* `stateless` - (Optional) A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - (Optional) Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
            * `max` - (Optional) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
            * `min` - (Optional) The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - (Optional) An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Required) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Required) The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - (Optional) Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Optional) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Optional) The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - (Optional) An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Required) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Required) The minimum port number. Must not be greater than the maximum port number.
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ingress_security_rules` - (Optional) Rules for allowing ingress IP packets.
	* `icmp_options` - (Optional) Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - (Optional) The ICMP code (optional).
		* `type` - (Required) The ICMP type.
	* `protocol` - (Required) The transport protocol. Specify either `all` or an IPv4 protocol number as defined in [Protocol Numbers](http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml). Options are supported only for ICMP ("1"), TCP ("6"), and UDP ("17"). 
	* `source` - (Required) The source CIDR block for the ingress rule. This is the range of IP addresses that a packet coming into the instance can come from. 
	* `stateless` - (Optional) A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if ingress traffic allows TCP destination port 80, there should be an egress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - (Optional) Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Optional) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Optional) The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - (Optional) An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Required) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Required) The minimum port number. Must not be greater than the maximum port number.
	* `udp_options` - (Optional) Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Optional) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Optional) The minimum port number. Must not be greater than the maximum port number.
		* `source_port_range` - (Optional) An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
			* `max` - (Required) The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
			* `min` - (Required) The minimum port number. Must not be greater than the maximum port number.
* `vcn_id` - (Required) The OCID of the VCN the security list belongs to.


### Update Operation
Updates the specified security list's display name or rules.
Avoid entering confidential information.

Note that the `egressSecurityRules` or `ingressSecurityRules` objects you provide replace the entire
existing objects.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `egress_security_rules` - Rules for allowing egress IP packets.
	* `icmp_options` - Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if egress traffic allows TCP destination port 80, there should be an ingress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
        	* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
        	* `min` - The minimum port number. Must not be greater than the maximum port number. 
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
		* The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
            * `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
            * `min` - The minimum port number. Must not be greater than the maximum port number. 
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ingress_security_rules` - Rules for allowing ingress IP packets.
	* `icmp_options` - Optional and valid only for ICMP. Use to specify a particular ICMP type and code as defined in [ICMP Parameters](http://www.iana.org/assignments/icmp-parameters/icmp-parameters.xhtml). If you specify ICMP as the protocol but omit this object, then all ICMP types and codes are allowed. If you do provide this object, the type is required and the code is optional. To enable MTU negotiation for ingress internet traffic, make sure to allow type 3 ("Destination Unreachable") code 4 ("Fragmentation Needed and Don't Fragment was Set"). If you need to specify multiple codes for a single type, create a separate security list rule for each. 
		* `code` - The ICMP code (optional).
	* `stateless` - A stateless rule allows traffic in one direction. Remember to add a corresponding stateless rule in the other direction if you need to support bidirectional traffic. For example, if ingress traffic allows TCP destination port 80, there should be an egress rule to allow TCP source port 80. Defaults to false, which means the rule is stateful and a corresponding rule is not necessary for bidirectional traffic. 
	* `tcp_options` - Optional and valid only for TCP. Use to specify particular destination ports for TCP rules. If you specify TCP as the protocol but omit this object, then all destination ports are allowed. 
        * The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
        	* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
        	* `min` - The minimum port number. Must not be greater than the maximum port number. 
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
	* `udp_options` - Optional and valid only for UDP. Use to specify particular destination ports for UDP rules. If you specify UDP as the protocol but omit this object, then all destination ports are allowed. 
        * The following 2 attributes specify an inclusive range of allowed destination ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 
        	* `max` - The maximum port number. Must not be lower than the minimum port number. To specify a single port number, set both the min and max to the same value. 
        	* `min` - The minimum port number. Must not be greater than the maximum port number. 
		* `source_port_range` - An inclusive range of allowed source ports. Use the same number for the min and max to indicate a single port. Defaults to all ports if not specified. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_security_list" "test_security_list" {
	#Required
	compartment_id = "${var.compartment_id}"
	egress_security_rules {
		#Required
		destination = "${var.security_list_egress_security_rules_destination}"
		protocol = "${var.security_list_egress_security_rules_protocol}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_egress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_egress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_egress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_tcp_options_source_port_range_min}"
			}
		}
		udp_options {

			#Optional
			max = "${var.security_list_egress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_egress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_egress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_egress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	ingress_security_rules {
		#Required
		protocol = "${var.security_list_ingress_security_rules_protocol}"
		source = "${var.security_list_ingress_security_rules_source}"

		#Optional
		icmp_options {
			#Required
			type = "${var.security_list_ingress_security_rules_icmp_options_type}"

			#Optional
			code = "${var.security_list_ingress_security_rules_icmp_options_code}"
		}
		stateless = "${var.security_list_ingress_security_rules_stateless}"
		tcp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_tcp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_tcp_options_source_port_range_min}"
			}
		}
		udp_options {

			#Optional
			max = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_max}"
			min = "${var.security_list_ingress_security_rules_udp_options_destination_port_range_min}"
			source_port_range {
				#Required
				max = "${var.security_list_ingress_security_rules_udp_options_source_port_range_max}"
				min = "${var.security_list_ingress_security_rules_udp_options_source_port_range_min}"
			}
		}
	}
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.security_list_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

# oci_core_security_lists

## SecurityList DataSource

Gets a list of security_lists.

### List Operation
Lists the security lists in the specified VCN and compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


The following attributes are exported:

* `security_lists` - The list of security_lists.

### Example Usage

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