# oci_core_private_ip

## PrivateIp Resource

### PrivateIp Reference

The following attributes are exported:

* `availability_domain` - The private IP's Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the private IP.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `bminstance-1` 
* `id` - The private IP's Oracle ID (OCID).
* `ip_address` - The private IP address of the `privateIp` object. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3` 
* `is_primary` - Whether this private IP is the primary one on the VNIC. Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.  Example: `true` 
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the private IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The OCID of the VNIC the private IP is assigned to. The VNIC and private IP must be in the same subnet. 



### Create Operation
Creates a secondary private IP for the specified VNIC.
For more information about secondary private IPs, see
[IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPaddresses.htm).


The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `bminstance-1` 
* `ip_address` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.  Example: `10.0.3.3` 
* `vnic_id` - (Required) The OCID of the VNIC to assign the private IP to. The VNIC and private IP must be in the same subnet. 


### Update Operation
Updates the specified private IP. You must specify the object's OCID.
Use this operation if you want to:

  - Move a secondary private IP to a different VNIC in the same subnet.
  - Change the display name for a secondary private IP.
  - Change the hostname for a secondary private IP.

This operation cannot be used with primary private IPs.
To update the hostname for the primary IP on a VNIC, use
[UpdateVnic](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/UpdateVnic).


The following arguments support updates:
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).  Example: `bminstance-1` 
* `vnic_id` - The OCID of the VNIC to assign the private IP to. The VNIC and private IP must be in the same subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```hcl
resource "oci_core_private_ip" "test_private_ip" {
	#Required
	vnic_id = "${oci_core_vnic.test_vnic.id}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.private_ip_display_name}"
	freeform_tags = {"Department"= "Finance"}
	hostname_label = "${var.private_ip_hostname_label}"
	ip_address = "${var.private_ip_ip_address}"
}
```

# oci_core_private_ips

## PrivateIp DataSource

Gets a list of private_ips.

### List Operation
Lists the [PrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/) objects based
on one of these filters:

  - Subnet OCID.
  - VNIC OCID.
  - Both private IP address and subnet OCID: This lets
  you get a `privateIP` object based on its private IP
  address (for example, 10.0.3.3) and not its OCID. For comparison,
  [GetPrivateIp](https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp)
  requires the OCID.

If you're listing all the private IPs associated with a given subnet
or VNIC, the response includes both primary and secondary private IPs.

The following arguments are supported:

* `ip_address` - (Optional) An IP address.  Example: `10.0.3.3` 
* `subnet_id` - (Optional) The OCID of the subnet.
* `vnic_id` - (Optional) The OCID of the VNIC.


The following attributes are exported:

* `private_ips` - The list of private_ips.

### Example Usage

```hcl
# Filter on Subnet OCID
data "oci_core_private_ips" "test_private_ips_by_subnet" {
	#Optional
	subnet_id = "${var.private_ip_subnet_id}"
}
```
```hcl
# Filter on VNIC OCID
data "oci_core_private_ips" "test_private_ips_by_vnic" {
	#Optional
	vnic_id = "${oci_core_vnic.test_vnic.id}"
}
```
```hcl
# Filter on private IP address and Subnet OCID
data "oci_core_private_ips" "test_private_ips_by_ip_address" {
	#Optional
	ip_address = "${var.private_ip_ip_address}"
	subnet_id = "${oci_core_subnet.test_subnet.id}"
}
```