# oci\_core\_private_ip

[PrivateIp Reference][b236f15a]

  [b236f15a]: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PrivateIp/ "PrivateIpReference"

Provides a private IP resource.

## Example Usage

```
resource "oci_core_private_ip" "testPrivateIP" {
	vnic_id = "${var.vnic_id}"
	display_name = "${var.display_name}"
	hostname_label = "${var.hostname_label}"
	ip_address = "${var.ip_address}"
}

```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `hostname_label` - (Optional) The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).  Example: `bminstance-1`
* `ip_address` - (Optional) A private IP address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns a private IP address from the subnet.  Example: `10.0.3.3`
* `vnic_id` - (Required) The OCID of the VNIC to assign the private IP to. The VNIC and private IP must be in the same subnet.


## PrivateIP Reference
* `availability_domain` - The private IP's Availability Domain.  Example: `Uocm:PHX-AD-1`
* `compartment_id` - The OCID of the compartment containing the private IP.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).  For more information, see [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).  Example: `bminstance-1`
* `id` - The private IP's Oracle ID (OCID).
* `ip_address` - The private IP address of the `privateIp` object. The address is within the CIDR of the VNIC's subnet.  Example: `10.0.3.3`
* `is_primary` - Whether this private IP is the primary one on the VNIC. Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.  Example: `true`
* `subnet_id` - The OCID of the subnet the VNIC is in.
* `time_created` - The date and time the private IP was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z`
* `vnic_id` - The OCID of the VNIC the private IP is assigned to. The VNIC and private IP must be in the same subnet.
