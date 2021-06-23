---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_nat_gateway"
sidebar_current: "docs-oci-resource-core-nat_gateway"
description: |-
  Provides the Nat Gateway resource in Oracle Cloud Infrastructure Core service
---

# oci_core_nat_gateway
This resource provides the Nat Gateway resource in Oracle Cloud Infrastructure Core service.

Creates a new NAT gateway for the specified VCN. You must also set up a route rule with the
NAT gateway as the rule's target. See [Route Table](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/RouteTable/).


## Example Usage

```hcl
resource "oci_core_nat_gateway" "test_nat_gateway" {
	#Required
	compartment_id = var.compartment_id
	vcn_id = oci_core_vcn.test_vcn.id

	#Optional
	block_traffic = var.nat_gateway_block_traffic
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.nat_gateway_display_name
	freeform_tags = {"Department"= "Finance"}
	public_ip_id = oci_core_public_ip.test_public_ip.id
}
```

## Argument Reference

The following arguments are supported:

* `block_traffic` - (Optional) (Updatable) Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `true` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the NAT gateway. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `public_ip_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway. 
* `vcn_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the gateway belongs to. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `block_traffic` - Whether the NAT gateway blocks traffic through it. The default is `false`.  Example: `true` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the NAT gateway. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the NAT gateway. 
* `nat_ip` - The IP address associated with the NAT gateway. 
* `public_ip_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway. 
* `state` - The NAT gateway's current state.
* `time_created` - The date and time the NAT gateway was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the NAT gateway belongs to. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Nat Gateway
	* `update` - (Defaults to 20 minutes), when updating the Nat Gateway
	* `delete` - (Defaults to 20 minutes), when destroying the Nat Gateway


## Import

NatGateways can be imported using the `id`, e.g.

```
$ terraform import oci_core_nat_gateway.test_nat_gateway "id"
```

