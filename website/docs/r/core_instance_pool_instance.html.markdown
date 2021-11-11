---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool_instance"
sidebar_current: "docs-oci-resource-core-instance_pool_instance"
description: |-
  Provides the Instance Pool Instance resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_pool_instance
This resource provides the Instance Pool Instance resource in Oracle Cloud Infrastructure Core service.

Attaches an instance to an instance pool. For information about the prerequisites
that an instance must meet before you can attach it to a pool, see
[Attaching an Instance to an Instance Pool](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/updatinginstancepool.htm#attach-instance).

Using this resource will impact the size of the instance pool, attach will increment the size of the pool

## Example Usage

```hcl
resource "oci_core_instance_pool_instance" "test_instance_pool_instance" {
	#Required
	instance_id = oci_core_instance.test_instance.id
	instance_pool_id = oci_core_instance_pool.test_instance_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `instance_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the instance. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - The fault domain the instance is running in.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `instance_configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration used to create the instance. 
* `instance_pool_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
* `load_balancer_backends` - The load balancer backends that are configured for the instance pool instance. 
	* `backend_health_status` - The health of the backend as observed by the load balancer.
	* `backend_name` - The name of the backend in the backend set.
	* `backend_set_name` - The name of the backend set on the load balancer.
	* `load_balancer_id` - The OCID of the load balancer attached to the instance pool.
* `region` - The region that contains the availability domain the instance is running in.
* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Shape/ListShapes). 
* `state` - The lifecycle state of the instance. Refer to `lifecycleState` in the [Instance](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance) resource.
* `time_created` - The date and time the instance pool instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Instance Pool Instance
	* `update` - (Defaults to 20 minutes), when updating the Instance Pool Instance
	* `delete` - (Defaults to 20 minutes), when destroying the Instance Pool Instance


## Import

InstancePoolInstances can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_pool_instance.test_instance_pool_instance "instancePools/{instancePoolId}/instances/compartmentId/{compartmentId}" 
```

