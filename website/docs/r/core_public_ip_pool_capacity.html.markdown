---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_public_ip_pool_capacity"
sidebar_current: "docs-oci-resource-core-public_ip_pool_capacity"
description: |-
  Provides the Public Ip Pool Capacity resource in Oracle Cloud Infrastructure Core service
---

# oci_core_public_ip_pool_capacity
This resource is used to manage the `cidr_blocks` of Public Ip Pool resource in Oracle Cloud Infrastructure Core service. 
Adds a Cidr from the named Byoip Range prefix to the referenced Public IP Pool. The cidr must be a subset of the Byoip Range in question. The cidr must not overlap with any other cidr already added to this or any other Public Ip Pool.

**Note:** When a new `oci_core_public_ip_pool_capacity` resource is created or removed, terraform needs to be refreshed to update the `cidr_blocks` of `oci_core_public_ip_pool` resource in state file.

## Example Usage

```hcl
resource "oci_core_public_ip_pool_capacity" "test_public_ip_pool_capacity" {
	#Required
	public_ip_pool_id = "${var.public_ip_pool_id}"
	byoip_id = "${var.byoip_id}"
	cidr_block = "${var.cidr_block}"
}
```

## Argument Reference

The following arguments are supported:

* `public_ip_pool_id` - (Required) The OCID of the pool object created by the current tenancy
* `byoip_id` - (Required) The OCID of the Byoip Range Id object to which the cidr block belongs.
* `cidr_block` - (Required) The CIDR IP address range to be added to the Public Ip Pool. Example: `10.0.1.0/24`

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `public_ip_pool_id` - (Required) The OCID of the pool object created by the current tenancy
* `byoip_id` - (Required) The OCID of the Byoip Range Id object to which the cidr block belongs.
* `cidr_block` - (Required) The CIDR IP address range to be added to the Public Ip Pool. Example: `10.0.1.0/24`

## Import

PublicIpPoolCapacity can be imported using the `id`, e.g.

```
$ terraform import oci_core_public_ip_pool_capacity.test_public_ip_pool_capacity "publicIpPoolId/{publicIpPoolId}/byoipId/{byoipId}/cidrBlock/{cidrBlock}"
```
