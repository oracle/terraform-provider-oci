---
layout: "oci"
page_title: "OCI: oci_core_private_ip"
sidebar_current: "docs-oci-resource-core-private_ip"
description: |-
  Creates and manages an OCI PrivateIp
---

# oci_core_private_ip
The `oci_core_private_ip` resource creates and manages an OCI PrivateIp

Creates a secondary private IP for the specified VNIC.
For more information about secondary private IPs, see
[IP Addresses](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingIPaddresses.htm).


## Example Usage

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

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - (Optional) (Updatable) The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).

	Example: `bminstance-1` 
* `ip_address` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.  Example: `10.0.3.3` 
* `vnic_id` - (Required) (Updatable) The OCID of the VNIC to assign the private IP to. The VNIC and private IP must be in the same subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The private IP's Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment containing the private IP.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/dns.htm).

	Example: `bminstance-1` 
* `id` - The private IP's Oracle ID (OCID).
* `ip_address` - The private IP address of the `privateIp` object. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3` 
* `is_primary` - Whether this private IP is the primary one on the VNIC. Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.  Example: `true` 
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the private IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The OCID of the VNIC the private IP is assigned to. The VNIC and private IP must be in the same subnet. 

## Import

PrivateIps can be imported using the `id`, e.g.

```
$ terraform import oci_core_private_ip.test_private_ip "id"
```
