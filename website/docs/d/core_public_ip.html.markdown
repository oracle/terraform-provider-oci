---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ip"
sidebar_current: "docs-oci-datasource-core-public_ip"
description: |-
  Provides details about a specific Public Ip in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_public_ip
This data source provides details about a specific Public Ip resource in Oracle Cloud Infrastructure Core service.

Gets the specified public IP. You must specify the object's [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

Alternatively, you can get the object by using [GetPublicIpByIpAddress](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/GetPublicIpByIpAddress)
with the public IP address (for example, 203.0.113.2).

Or you can use [GetPublicIpByPrivateIpId](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/PublicIp/GetPublicIpByPrivateIpId)
with the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private IP that the public IP is assigned to.

**Note:** If you're fetching a reserved public IP that is in the process of being
moved to a different private IP, the service returns the public IP object with
`lifecycleState` = ASSIGNING and `assignedEntityId` = [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target private IP.


## Example Usage

### Get a public ip by public ip id
```hcl
data "oci_core_public_ip" "test_oci_core_public_ip_by_id" {
    id = var.test_public_ip_id
}
```

### Get a public ip by private ip id
```hcl
data "oci_core_public_ip" "test_oci_core_public_ip_by_private_ip_id" {
    private_ip_id = var.test_public_ip_private_ip_id
}
```

### Get a public ip by public ip address
```hcl
data "oci_core_public_ip" "test_oci_core_public_ip_by_ip" {
    ip_address = var.test_public_ip_ip_address
}
```

## Argument Reference

The following arguments are supported:

_Only one of the following values will be used. If multiple arguments are passed, the first non-empty value will be used based on the order below._
  
* `id` - (Optional) The OCID of the public IP.
* `private_ip_id` - (Optional) Gets the public IP assigned to the specified private IP. You must specify the OCID of the private IP. If no public IP is assigned, a 404 is returned.
* `ip_address` - (Optional) Gets the public IP based on the public IP address (for example, 129.146.2.1).


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