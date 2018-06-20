# oci_core_subnet

## Subnet Resource

### Subnet Reference

The following attributes are exported:

* `availability_domain` - The subnet's Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - The subnet's CIDR block.  Example: `172.16.1.0/24` 
* `compartment_id` - The OCID of the compartment containing the subnet.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `dhcp_options_id` - The OCID of the set of DHCP options associated with the subnet. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `dns_label` - A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.  The absence of this parameter means the Internet and VCN Resolver will not resolve hostnames of instances in this subnet.  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `subnet123` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The subnet's Oracle ID (OCID).
* `prohibit_public_ip_on_vnic` - Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CreateVnicDetails/)). If `prohibitPublicIpOnVnic` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).  Example: `true` 
* `route_table_id` - The OCID of the route table the subnet is using.
* `security_list_ids` - OCIDs for the security lists to use for VNICs in this subnet.
* `state` - The subnet's current state.
* `subnet_domain_name` - The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the `oraclevcn.com` domain.  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `subnet123.vcn1.oraclevcn.com` 
* `time_created` - The date and time the subnet was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the subnet is in.
* `virtual_router_ip` - The IP address of the virtual router.  Example: `10.0.14.1` 
* `virtual_router_mac` - The MAC address of the virtual router.  Example: `00:00:17:B6:4D:DD` 



### Create Operation
Creates a new subnet in the specified VCN. You can't change the size of the subnet after creation,
so it's important to think about the size of subnets you need before creating them.
For more information, see [VCNs and Subnets](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVCNs.htm).
For information on the number of subnets you can have in a VCN, see
[Service Limits](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/servicelimits.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want the subnet
to reside. Notice that the subnet doesn't have to be in the same compartment as the VCN, route tables, or
other Networking Service components. If you're not sure which compartment to use, put the subnet in
the same compartment as the VCN. For more information about compartments and access control, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm). For information about OCIDs,
see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally associate a route table with the subnet. If you don't, the subnet will use the
VCN's default route table. For more information about route tables, see
[Route Tables](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm).

You may optionally associate a security list with the subnet. If you don't, the subnet will use the
VCN's default security list. For more information about security lists, see
[Security Lists](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).

You may optionally associate a set of DHCP options with the subnet. If you don't, the subnet will use the
VCN's default set. For more information about DHCP options, see
[DHCP Options](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDHCP.htm).

You may optionally specify a *display name* for the subnet, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.

You can also add a DNS label for the subnet, which is required if you want the Internet and
VCN Resolver to resolve hostnames for instances in the subnet. For more information, see
[DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).


The following arguments are supported:

* `availability_domain` - (Required) The Availability Domain to contain the subnet.  Example: `Uocm:PHX-AD-1` 
* `cidr_block` - (Required) The CIDR IP address range of the subnet.  Example: `172.16.1.0/24` 
* `compartment_id` - (Required) The OCID of the compartment to contain the subnet.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `dhcp_options_id` - (Optional) The OCID of the set of DHCP options the subnet will use. If you don't provide a value, the subnet will use the VCN's default set of DHCP options. 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `dns_label` - (Optional) A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be an alphanumeric string that begins with a letter and is unique within the VCN. The value cannot be changed.  This value must be set if you want to use the Internet and VCN Resolver to resolve the hostnames of instances in the subnet. It can only be set if the VCN itself was created with a DNS label.  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `subnet123` 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `prohibit_public_ip_on_vnic` - (Optional) Whether VNICs within this subnet can have public IP addresses. Defaults to false, which means VNICs created in this subnet will automatically be assigned public IP addresses unless specified otherwise during instance launch or VNIC creation (with the `assignPublicIp` flag in [CreateVnicDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/CreateVnicDetails/)). If `prohibitPublicIpOnVnic` is set to true, VNICs created in this subnet cannot have public IP addresses (that is, it's a private subnet).  Example: `true` 
* `route_table_id` - (Optional) The OCID of the route table the subnet will use. If you don't provide a value, the subnet will use the VCN's default route table. 
* `security_list_ids` - (Required) OCIDs for the security lists to associate with the subnet. If you don't provide a value, the VCN's default security list will be associated with the subnet. Remember that security lists are associated at the subnet level, but the rules are applied to the individual VNICs in the subnet. 
* `vcn_id` - (Required) The OCID of the VCN to contain the subnet.


### Update Operation
Updates the specified subnet's display name. Avoid entering confidential information.


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_subnet" "test_subnet" {
	#Required
	availability_domain = "${var.subnet_availability_domain}"
	cidr_block = "${var.subnet_cidr_block}"
	compartment_id = "${var.compartment_id}"
	security_list_ids = "${var.subnet_security_list_ids}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	dhcp_options_id = "${oci_core_dhcp_options.test_dhcp_options.id}"
	display_name = "${var.subnet_display_name}"
	dns_label = "${var.subnet_dns_label}"
	freeform_tags = {"Department"= "Finance"}
	prohibit_public_ip_on_vnic = "${var.subnet_prohibit_public_ip_on_vnic}"
	route_table_id = "${oci_core_route_table.test_route_table.id}"
}
```

# oci_core_subnets

## Subnet DataSource

Gets a list of subnets.

### List Operation
Lists the subnets in the specified VCN and the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


The following attributes are exported:

* `subnets` - The list of subnets.

### Example Usage

```hcl
data "oci_core_subnets" "test_subnets" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.subnet_display_name}"
	state = "${var.subnet_state}"
}
```