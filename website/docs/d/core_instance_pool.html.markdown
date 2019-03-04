---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool"
sidebar_current: "docs-oci-datasource-core-instance_pool"
description: |-
  Provides details about a specific Instance Pool in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_pool
This data source provides details about a specific Instance Pool resource in Oracle Cloud Infrastructure Core service.

Gets the specified instance pool

## Example Usage

```hcl
data "oci_core_instance_pool" "test_instance_pool" {
	#Required
	instance_pool_id = "${oci_core_instance_pool.test_instance_pool.id}"
}
```

## Argument Reference

The following arguments are supported:

* `instance_pool_id` - (Required) The OCID of the instance pool.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the instance pool
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name.  Does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance pool
* `instance_configuration_id` - The OCID of the instance configuration associated to the intance pool.
* `load_balancers` - The load balancers attached to the instance pool. 
	* `backend_set_name` - The name of the backend set on the load balancer.
	* `id` - The OCID of the load balancer attachment.
	* `instance_pool_id` - The OCID of the instance pool of the load balancer attachment.
	* `load_balancer_id` - The OCID of the load balancer attached to the pool.
	* `port` - The port value used for the backends.
	* `state` - The status of the interaction between the pool and the load balancer.
	* `vnic_selection` - Indicates which vnic on each instance in the pool should be used to associate with the load balancer. possible values are "PrimaryVnic" or the displayName of one of the secondary VNICs on the instance configuration that is associated to the instance pool.
* `placement_configurations` - The placement configurations for the instance pool.
	* `availability_domain` - The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - The OCID of the primary subnet to place instances.
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The displayName of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
		* `subnet_id` - The subnet OCID for the secondary vnic
* `size` - The number of actual instances in the instance pool on the cloud. This attribute will be different when instance pool is used along with autoScaling Configuration.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

