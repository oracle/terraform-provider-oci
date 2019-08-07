---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_private_ips"
sidebar_current: "docs-oci-datasource-core-private_ips"
description: |-
  Provides the list of Private Ips in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_private_ips
This data source provides the list of Private Ips in Oracle Cloud Infrastructure Core service.

Lists the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/) objects based
on one of these filters:

  - Subnet OCID.
  - VNIC OCID.
  - Both private IP address and subnet OCID: This lets
  you get a `privateIP` object based on its private IP
  address (for example, 10.0.3.3) and not its OCID. For comparison,
  [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/PrivateIp/GetPrivateIp)
  requires the OCID.

If you're listing all the private IPs associated with a given subnet
or VNIC, the response includes both primary and secondary private IPs.


## Example Usage

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

## Argument Reference

The following arguments are supported:

* `ip_address` - (Optional) An IP address. This could be either IPv4 or IPv6, depending on the resource. Example: `10.0.3.3` 
* `subnet_id` - (Optional) The OCID of the subnet.
* `vnic_id` - (Optional) The OCID of the VNIC.


## Attributes Reference

The following attributes are exported:

* `private_ips` - The list of private_ips.

### PrivateIp Reference

The following attributes are exported:

* `availability_domain` - The private IP's availability domain. This attribute will be null if this is a *secondary* private IP assigned to a VNIC that is in a *regional* subnet.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the private IP.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `bminstance-1` 
* `id` - The private IP's Oracle ID (OCID).
* `ip_address` - The private IP address of the `privateIp` object. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3` 
* `is_primary` - Whether this private IP is the primary one on the VNIC. Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.  Example: `true` 
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the private IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The OCID of the VNIC the private IP is assigned to. The VNIC and private IP must be in the same subnet. 

