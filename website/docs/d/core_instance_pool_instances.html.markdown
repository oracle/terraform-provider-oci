---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool_instances"
sidebar_current: "docs-oci-datasource-core-instance_pool_instances"
description: |-
  Provides the list of Instance Pool Instances in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_pool_instances
This data source provides the list of Instance Pool Instances in Oracle Cloud Infrastructure Core service.

List the instances in the specified instance pool.

## Example Usage

```hcl
data "oci_core_instance_pool_instances" "test_instance_pool_instances" {
	#Required
	compartment_id = var.compartment_id
	instance_pool_id = oci_core_instance_pool.test_instance_pool.id

	#Optional
	display_name = var.instance_pool_instance_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `instance_pool_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.


## Attributes Reference

The following attributes are exported:

* `instances` - The list of instances.

### InstancePoolInstance Reference

The following attributes are exported:

* `availability_domain` - The availability domain the instance is running in.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the instance. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `fault_domain` - The fault domain the instance is running in.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `instance_configuration_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration used to create the instance. 
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

