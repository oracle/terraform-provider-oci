---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ips"
sidebar_current: "docs-oci-datasource-core-public_ips"
description: |-
  Provides the list of Public Ips in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_public_ips
This data source provides the list of Public Ips in Oracle Cloud Infrastructure Core service.

Lists the [PublicIp](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/) objects
in the specified compartment. You can filter the list by using query parameters.

To list your reserved public IPs:
  * Set `scope` = `REGION`  (required)
  * Leave the `availabilityDomain` parameter empty
  * Set `lifetime` = `RESERVED`

To list the ephemeral public IPs assigned to a regional entity such as a NAT gateway:
  * Set `scope` = `REGION`  (required)
  * Leave the `availabilityDomain` parameter empty
  * Set `lifetime` = `EPHEMERAL`

To list the ephemeral public IPs assigned to private IPs:
  * Set `scope` = `AVAILABILITY_DOMAIN` (required)
  * Set the `availabilityDomain` parameter to the desired availability domain (required)
  * Set `lifetime` = `EPHEMERAL`

**Note:** An ephemeral public IP assigned to a private IP
is always in the same availability domain and compartment as the private IP.


## Example Usage

```hcl
data "oci_core_public_ips" "test_public_ips" {
	#Required
	compartment_id = var.compartment_id
	scope = var.public_ip_scope

	#Optional
	availability_domain = var.public_ip_availability_domain
	lifetime = var.public_ip_lifetime
	public_ip_pool_id = oci_core_public_ip_pool.test_public_ip_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `lifetime` - (Optional) A filter to return only public IPs that match given lifetime. 
* `public_ip_pool_id` - (Optional) A filter to return only resources that belong to the given public IP pool. 
* `scope` - (Required) Whether the public IP is regional or specific to a particular availability domain.
	* `REGION`: The public IP exists within a region and is assigned to a regional entity (such as a [NatGateway](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NatGateway/)), or can be assigned to a private IP in any availability domain in the region. Reserved public IPs have `scope` = `REGION`, as do ephemeral public IPs assigned to a regional entity.
	* `AVAILABILITY_DOMAIN`: The public IP exists within the availability domain of the entity it's assigned to, which is specified by the `availabilityDomain` property of the public IP object. Ephemeral public IPs that are assigned to private IPs have `scope` = `AVAILABILITY_DOMAIN`. 


## Attributes Reference

The following attributes are exported:

* `public_ips` - The list of public_ips.

### PublicIp Reference

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

