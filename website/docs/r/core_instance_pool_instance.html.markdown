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

Attach an instance to the instance pool.

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

* `instance_id` - (Required) the instance ocid to attach.
* `instance_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.
* `compartment_id` - The OCID of the compartment that contains the instance.
* `display_name` - The user-friendly name. Does not have to be unique.
* `fault_domain` - The fault domain the instance is running in.
* `id` - The OCID of the instance.
* `instance_configuration_id` - The OCID of the instance confgiuration used to create the instance.
* `instance_pool_id` - The OCID of the instance pool.
* `load_balancer_backends` - The load balancer backends that are configured for the instance pool instance. 
	* `backend_health_status` - The health of the backend as observed by the load balancer.
	* `backend_name` - The name of the backend in the backend set.
	* `backend_set_name` - The name of the backend set on the load balancer.
	* `load_balancer_id` - The OCID of the load balancer attached to the instance pool.
* `region` - The region that contains the availability domain the instance is running in.
* `shape` - The shape of an instance. The shape determines the number of CPUs, amount of memory, and other resources allocated to the instance.

	You can enumerate all available shapes by calling [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Shape/ListShapes). 
* `state` - The lifecycleState of the underlying instance. Refer lifecycleState in [Instance details](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Instance)
* `time_created` - The date and time the instance pool instance was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

## Import

InstancePoolInstances can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_pool_instance.test_instance_pool_instance "instancePools/{instancePoolId}/instances/compartmentId/{compartmentId}" 
```

