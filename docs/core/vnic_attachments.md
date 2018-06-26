# oci_core_vnic_attachment

## VnicAttachment Resource

### VnicAttachment Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of the instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in. 
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information. 
* `id` - The OCID of the VNIC attachment.
* `instance_id` - The OCID of the instance.
* `nic_index` - Which physical network interface card (NIC) the VNIC uses. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm). 
* `state` - The current state of the VNIC attachment.
* `subnet_id` - The OCID of the VNIC's subnet.
* `time_created` - The date and time the VNIC attachment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vlan_tag` - The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.  Example: `0` 
* `vnic_id` - The OCID of the VNIC. Available after the attachment process is complete.



### Create Operation
Creates a secondary VNIC and attaches it to the specified instance.
For more information about secondary VNICs, see
[Virtual Network Interface Cards (VNICs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm).


The following arguments are supported:

* `create_vnic_details` - (Required) Details for creating a new VNIC. 
	* `assign_public_ip` - (Optional) Whether the VNIC should be assigned a public IP address. Defaults to whether the subnet is public or private. If not set and the VNIC is being created in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the [Subnet](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/)), then no public IP address is assigned. If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then a public IP address is assigned. If set to true and `prohibitPublicIpOnVnic` = true, an error is returned.  **Note:** This public IP address is associated with the primary private IP on the VNIC. For more information, see [IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPaddresses.htm).  **Note:** There's a limit to the number of [public IPs](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PublicIp/) a VNIC or instance can have. If you try to create a secondary VNIC with an assigned public IP for an instance that has already reached its public IP limit, an error is returned. For information about the public IP limits, see [Public IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingpublicIPs.htm).  Example: `false` 
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) A user-friendly name for the VNIC. Does not have to be unique. Avoid entering confidential information. 
	* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `hostname_label` - (Optional) The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname portion of the primary private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123). The value appears in the [Vnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  When launching an instance, use this `hostnameLabel` instead of the deprecated `hostnameLabel` in [LaunchInstanceDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). If you provide both, the values must match.  Example: `bminstance-1` 
	* `private_ip` - (Optional) A private IP address of your choice to assign to the VNIC. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet. This is the VNIC's *primary* private IP address. The value appears in the [Vnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/) object and also the [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/) object returned by [ListPrivateIps](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/ListPrivateIps) and [GetPrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp).  Example: `10.0.3.3` 
	* `skip_source_dest_check` - (Optional) Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. For information about why you would skip the source/destination check, see [Using a Private IP as a Route Target](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingroutetables.htm#privateip).  Example: `true` 
	* `subnet_id` - (Required) The OCID of the subnet to create the VNIC in. When launching an instance, use this `subnetId` instead of the deprecated `subnetId` in [LaunchInstanceDetails](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/requests/LaunchInstanceDetails). At least one of them is required; if you provide both, the values must match. 
* `display_name` - (Optional) A user-friendly name for the attachment. Does not have to be unique, and it cannot be changed. 
* `instance_id` - (Required) The OCID of the instance.
* `nic_index` - (Optional) Which physical network interface card (NIC) the VNIC will use. Defaults to 0. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm). 


### Update Operation


The following arguments support updates:
* NO arguments in this resource support updates

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_vnic_attachment" "test_vnic_attachment" {
	#Required
	create_vnic_details {
		#Required
		subnet_id = "${oci_core_subnet.test_subnet.id}"

		#Optional
		assign_public_ip = "${var.vnic_attachment_create_vnic_details_assign_public_ip}"
		defined_tags = "${var.vnic_attachment_create_vnic_details_defined_tags}"
		display_name = "${var.vnic_attachment_create_vnic_details_display_name}"
		freeform_tags = "${var.vnic_attachment_create_vnic_details_freeform_tags}"
		hostname_label = "${var.vnic_attachment_create_vnic_details_hostname_label}"
		private_ip = "${var.vnic_attachment_create_vnic_details_private_ip}"
		skip_source_dest_check = "${var.vnic_attachment_create_vnic_details_skip_source_dest_check}"
	}
	instance_id = "${oci_core_instance.test_instance.id}"

	#Optional
	display_name = "${var.vnic_attachment_display_name}"
	nic_index = "${var.vnic_attachment_nic_index}"
}
```

# oci_core_vnic_attachments

## VnicAttachment DataSource

Gets a list of vnic_attachments.

### List Operation
Lists the VNIC attachments in the specified compartment. A VNIC attachment
resides in the same compartment as the attached instance. The list can be
filtered by instance, VNIC, or Availability Domain.

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `vnic_id` - (Optional) The OCID of the VNIC.


The following attributes are exported:

* `vnic_attachments` - The list of vnic_attachments.

### Example Usage

```hcl
data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.vnic_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	vnic_id = "${oci_core_vnic.test_vnic.id}"
}
```