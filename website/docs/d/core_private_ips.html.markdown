---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_private_ips"
sidebar_current: "docs-oci-datasource-core-private_ips"
description: |-
  Provides the list of Private Ips in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_private_ips
This data source provides the list of Private Ips in Oracle Cloud Infrastructure Core service.

Lists the [PrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/) objects based
on one of these filters:

  - Subnet [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
  - VNIC [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
  - Both private IP address and subnet OCID: This lets
  you get a `privateIP` object based on its private IP
  address (for example, 10.0.3.3) and not its [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). For comparison,
  [GetPrivateIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PrivateIp/GetPrivateIp)
  requires the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

If you're listing all the private IPs associated with a given subnet
or VNIC, the response includes both primary and secondary private IPs.

If you are an Oracle Cloud VMware Solution customer and have VLANs
in your VCN, you can filter the list by VLAN [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).


## Example Usage

```hcl
# Filter on Subnet OCID
data "oci_core_private_ips" "test_private_ips_by_subnet" {
	#Optional
	subnet_id = var.private_ip_subnet_id
}
```
```hcl
# Filter on VNIC OCID
data "oci_core_private_ips" "test_private_ips_by_vnic" {
	#Optional
	vnic_id = oci_core_vnic.test_vnic.id
}
```
```hcl
# Filter on private IP address and Subnet OCID
data "oci_core_private_ips" "test_private_ips_by_ip_address" {
	#Optional
	ip_address = var.private_ip_ip_address
	ip_state = var.private_ip_ip_state
	lifetime = var.private_ip_lifetime
	subnet_id = oci_core_subnet.test_subnet.id
	vlan_id = oci_core_vlan.test_vlan.id
	vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `ip_address` - (Optional) An IP address. This could be either IPv4 or IPv6, depending on the resource. Example: `10.0.3.3` 
* `ip_state` - (Optional) State of the IP address. If an IP address is assigned to a VNIC it is ASSIGNED otherwise AVAILABLE
* `lifetime` - (Optional) Lifetime of the IP address. There are two types of IPs:
	* Ephemeral
	* Reserved 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet.
* `vlan_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN.
* `vnic_id` - (Optional) The OCID of the VNIC.


## Attributes Reference

The following attributes are exported:

* `private_ips` - The list of private_ips.

### PrivateIp Reference

The following attributes are exported:

* `availability_domain` - The private IP's availability domain. This attribute will be null if this is a *secondary* private IP assigned to a VNIC that is in a *regional* subnet.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the private IP.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname_label` - The hostname for the private IP. Used for DNS. The value is the hostname portion of the private IP's fully qualified domain name (FQDN) (for example, `bminstance1` in FQDN `bminstance1.subnet123.vcn1.oraclevcn.com`). Must be unique across all VNICs in the subnet and comply with [RFC 952](https://tools.ietf.org/html/rfc952) and [RFC 1123](https://tools.ietf.org/html/rfc1123).

	For more information, see [DNS in Your Virtual Cloud Network](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).

	Example: `bminstance1` 
* `id` - The private IP's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ip_address` - The private IP address of the `privateIp` object. The address is within the CIDR of the VNIC's subnet.

	However, if the `PrivateIp` object is being used with a VLAN as part of the Oracle Cloud VMware Solution, the address is from the range specified by the `cidrBlock` attribute for the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

	Example: `10.0.3.3` 
* `ip_state` - State of the IP address. If an IP address is assigned to a VNIC it is ASSIGNED, otherwise it is AVAILABLE. 
* `is_primary` - Whether this private IP is the primary one on the VNIC. Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.  Example: `true` 
* `lifetime` - Lifetime of the IP address. There are two types of IPv6 IPs:
	* Ephemeral
	* Reserved 
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see [Source Based Routing](https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing).
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.

	However, if the `PrivateIp` object is being used with a VLAN as part of the Oracle Cloud VMware Solution, the `subnetId` is null. 
* `time_created` - The date and time the private IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vlan_id` - Applicable only if the `PrivateIp` object is being used with a VLAN as part of the Oracle Cloud VMware Solution. The `vlanId` is the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan). 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the private IP is assigned to. The VNIC and private IP must be in the same subnet. However, if the `PrivateIp` object is being used with a VLAN as part of the Oracle Cloud VMware Solution, the `vnicId` is null. 

