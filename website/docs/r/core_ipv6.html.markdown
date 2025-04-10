---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_ipv6"
sidebar_current: "docs-oci-resource-core-ipv6"
description: |-
  Provides the Ipv6 resource in Oracle Cloud Infrastructure Core service
---

# oci_core_ipv6
This resource provides the Ipv6 resource in Oracle Cloud Infrastructure Core service.

Creates an IPv6 for the specified VNIC.


## Example Usage

```hcl
resource "oci_core_ipv6" "test_ipv6" {

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.ipv6_display_name
	freeform_tags = {"Department"= "Finance"}
	ip_address = var.ipv6_ip_address
	ipv6subnet_cidr = var.ipv6_ipv6subnet_cidr
	lifetime = var.ipv6_lifetime
	route_table_id = oci_core_route_table.test_route_table.id
	subnet_id = oci_core_subnet.test_subnet.id
	vnic_id = oci_core_vnic_attachment.test_vnic_attachment.id
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ip_address` - (Optional) An IPv6 address of your choice. Must be an available IP address within the subnet's CIDR. If you don't specify a value, Oracle automatically assigns an IPv6 address from the subnet. The subnet is the one that contains the VNIC you specify in `vnicId`.  Example: `2001:DB8::` 
* `ipv6subnet_cidr` - (Optional) The IPv6 prefix allocated to the subnet. This is required if more than one IPv6 prefix exists on the subnet. 
* `lifetime` - (Optional) (Updatable) Lifetime of the IP address. There are two types of IPv6 IPs:
	* Ephemeral
	* Reserved 
* `route_table_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see [Source Based Routing](https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing). 
* `subnet_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet from which the IPv6 is to be drawn. The IP address, *if supplied*, must be valid for the given subnet, only valid for reserved IPs currently. 
* `vnic_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC to assign the IPv6 to. The IPv6 will be in the VNIC's subnet. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the IPv6. This is the same as the VNIC's compartment. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the IPv6.
* `ip_address` - The IPv6 address of the `IPv6` object. The address is within the IPv6 prefix of the VNIC's subnet (see the `ipv6CidrBlock` attribute for the [Subnet](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Subnet/) object.  Example: `2001:0db8:0123:1111:abcd:ef01:2345:6789` 
* `ip_state` - State of the IP address. If an IP address is assigned to a VNIC it is ASSIGNED, otherwise it is AVAILABLE. 
* `lifetime` - Lifetime of the IP address. There are two types of IPv6 IPs:
	* Ephemeral
	* Reserved 
* `route_table_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the IP address or VNIC will use. For more information, see [Source Based Routing](https://docs.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#Overview_of_Routing_for_Your_VCN__source_routing). 
* `state` - The IPv6's current state.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the VNIC is in.
* `time_created` - The date and time the IPv6 was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC the IPv6 is assigned to. The VNIC and IPv6 must be in the same subnet. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Ipv6
	* `update` - (Defaults to 20 minutes), when updating the Ipv6
	* `delete` - (Defaults to 20 minutes), when destroying the Ipv6


## Import

Ipv6 can be imported using the `id`, e.g.

```
$ terraform import oci_core_ipv6.test_ipv6 "id"
```

