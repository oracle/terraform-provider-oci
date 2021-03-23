---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipv6s"
sidebar_current: "docs-oci-datasource-core-ipv6s"
description: |-
  Provides the list of Ipv6s in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_ipv6s
This data source provides the list of Ipv6s in Oracle Cloud Infrastructure Core service.

Lists the [IPv6](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Ipv6/) objects based
on one of these filters:

  * Subnet OCID.
  * VNIC OCID.
  * Both IPv6 address and subnet OCID: This lets you get an `Ipv6` object based on its private
  IPv6 address (for example, 2001:0db8:0123:1111:abcd:ef01:2345:6789) and not its OCID. For comparison,
  [GetIpv6](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Ipv6/GetIpv6) requires the OCID.


## Example Usage

```hcl
data "oci_core_ipv6s" "test_ipv6s" {

	#Optional
	ip_address = var.ipv6_ip_address
	subnet_id = oci_core_subnet.test_subnet.id
	vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `ip_address` - (Optional) An IP address. This could be either IPv4 or IPv6, depending on the resource. Example: `10.0.3.3` 
* `subnet_id` - (Optional) The OCID of the subnet.
* `vnic_id` - (Optional) The OCID of the VNIC.


## Attributes Reference

The following attributes are exported:

* `ipv6s` - The list of ipv6s.

### Ipv6 Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPv6. This is the same as the VNIC's compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPv6.
* `ip_address` - The IPv6 address of the `IPv6` object. The address is within the IPv6 CIDR block of the VNIC's subnet (see the `ipv6CidrBlock` attribute for the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/) object.  Example: `2001:0db8:0123:1111:abcd:ef01:2345:6789` 
* `state` - The IPv6's current state.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
* `time_created` - The date and time the IPv6 was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the IPv6 is assigned to. The VNIC and IPv6 must be in the same subnet. 

