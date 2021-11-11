---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ip"
sidebar_current: "docs-oci-resource-core-public_ip"
description: |-
  Provides the Public Ip resource in Oracle Cloud Infrastructure Core service
---

# oci_core_public_ip
This resource provides the Public Ip resource in Oracle Cloud Infrastructure Core service.

Creates a public IP. Use the `lifetime` property to specify whether it's an ephemeral or
reserved public IP. For information about limits on how many you can create, see
[Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).

* **For an ephemeral public IP assigned to a private IP:** You must also specify a `privateIpId`
with the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the primary private IP you want to assign the public IP to. The public IP is
created in the same availability domain as the private IP. An ephemeral public IP must always be
assigned to a private IP, and only to the *primary* private IP on a VNIC, not a secondary
private IP. Exception: If you create a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/), Oracle
automatically assigns the NAT gateway a regional ephemeral public IP that you cannot remove.

* **For a reserved public IP:** You may also optionally assign the public IP to a private
IP by specifying `privateIpId`. Or you can later assign the public IP with
[UpdatePublicIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/UpdatePublicIp).

**Note:** When assigning a public IP to a private IP, the private IP must not already have
a public IP with `lifecycleState` = ASSIGNING or ASSIGNED. If it does, an error is returned.

Also, for reserved public IPs, the optional assignment part of this operation is
asynchronous. Poll the public IP's `lifecycleState` to determine if the assignment
succeeded.


## Example Usage

```hcl
resource "oci_core_public_ip" "test_public_ip" {
	#Required
	compartment_id = var.compartment_id
	lifetime = var.public_ip_lifetime

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.public_ip_display_name
	freeform_tags = {"Department"= "Finance"}
	private_ip_id = oci_core_private_ip.test_private_ip.id
	public_ip_pool_id = oci_core_public_ip_pool.test_public_ip_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the public IP. For ephemeral public IPs, you must set this to the private IP's compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `lifetime` - (Required) Defines when the public IP is deleted and released back to the Oracle Cloud Infrastructure public IP pool. For more information, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm). 
* `private_ip_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private IP to assign the public IP to.

	Required for an ephemeral public IP because it must always be assigned to a private IP (specifically a *primary* private IP).

	Optional for a reserved public IP. If you don't provide it, the public IP is created but not assigned to a private IP. You can later assign the public IP with [UpdatePublicIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/UpdatePublicIp). 
* `public_ip_pool_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `assigned_entity_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the entity the public IP is assigned to, or in the process of being assigned to. 
* `assigned_entity_type` - The type of entity the public IP is assigned to, or in the process of being assigned to. 
* `availability_domain` - The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP (that is, when the `scope` of the public IP is set to AVAILABILITY_DOMAIN). The value is the availability domain of the assigned private IP.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the public IP. For an ephemeral public IP, this is the compartment of its assigned entity (which can be a private IP or a regional entity such as a NAT gateway). For a reserved public IP that is currently assigned, its compartment can be different from the assigned private IP's. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The public IP's Oracle ID ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
* `ip_address` - The public IP address of the `publicIp` object.  Example: `203.0.113.2` 
* `lifetime` - Defines when the public IP is deleted and released back to Oracle's public IP pool.
	* `EPHEMERAL`: The lifetime is tied to the lifetime of its assigned entity. An ephemeral public IP must always be assigned to an entity. If the assigned entity is a private IP, the ephemeral public IP is automatically deleted when the private IP is deleted, when the VNIC is terminated, or when the instance is terminated. If the assigned entity is a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/), the ephemeral public IP is automatically deleted when the NAT gateway is terminated.
	* `RESERVED`: You control the public IP's lifetime. You can delete a reserved public IP whenever you like. It does not need to be assigned to a private IP at all times.

	For more information and comparison of the two types, see [Public IP Addresses](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm). 
* `private_ip_id` - Deprecated. Use `assignedEntityId` instead.

	The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private IP that the public IP is currently assigned to, or in the process of being assigned to.

	**Note:** This is `null` if the public IP is not assigned to a private IP, or is in the process of being assigned to one. 
* `public_ip_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pool object created in the current tenancy.
* `scope` - Whether the public IP is regional or specific to a particular availability domain.
	* `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs and ephemeral public IPs assigned to a regional entity have `scope` = `REGION`.
	* `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`. 
* `state` - The public IP's current state.
* `time_created` - The date and time the public IP was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Public Ip
	* `update` - (Defaults to 20 minutes), when updating the Public Ip
	* `delete` - (Defaults to 20 minutes), when destroying the Public Ip


## Import

PublicIps can be imported using the `id`, e.g.

```
$ terraform import oci_core_public_ip.test_public_ip "id"
```

