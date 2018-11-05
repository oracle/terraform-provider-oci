---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_pool"
sidebar_current: "docs-oci-resource-core-instance_pool"
description: |-
  Provides the Instance Pool resource in Oracle Cloud Infrastructure Core service
---

# oci_core_instance_pool
This resource provides the Instance Pool resource in Oracle Cloud Infrastructure Core service.

Create an instance pool.

## Example Usage

```hcl
resource "oci_core_instance_pool" "test_instance_pool" {
	#Required
	compartment_id = "${var.compartment_id}"
	instance_configuration_id = "${oci_core_instance_configuration.test_instance_configuration.id}"
	placement_configurations {
		#Required
		availability_domain = "${var.instance_pool_placement_configurations_availability_domain}"
		primary_subnet_id = "${oci_core_primary_subnet.test_primary_subnet.id}"

		#Optional
		secondary_vnic_subnets {
			#Required
			subnet_id = "${oci_core_subnet.test_subnet.id}"

			#Optional
			display_name = "${var.instance_pool_placement_configurations_secondary_vnic_subnets_display_name}"
		}
	}
	size = "${var.instance_pool_size}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.instance_pool_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the instance pool
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name.  Does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_configuration_id` - (Required) (Updatable) The OCID of the instance configuration associated to the instance pool.
* `placement_configurations` - (Required) (Updatable) The placement configurations for the instance pool. There should be 1 placement configuration for each desired AD. 
	* `availability_domain` - (Required) (Updatable) The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - (Required) (Updatable) The OCID of the primary subnet to place instances.
	* `secondary_vnic_subnets` - (Optional) (Updatable) The set of secondary VNIC data for instances in the pool.
		* `display_name` - (Optional) (Updatable) The displayName of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
		* `subnet_id` - (Required) (Updatable) The subnet OCID for the secondary vnic
* `size` - (Required) (Updatable) The number of instances that should be in the instance pool.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the instance pool
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name.  Does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the instance pool
* `instance_configuration_id` - The OCID of the instance configuration associated to the intance pool.
* `placement_configurations` - The placement configurations for the instance pool.
	* `availability_domain` - The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - The OCID of the primary subnet to place instances.
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The displayName of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
		* `subnet_id` - The subnet OCID for the secondary vnic
* `size` - The number of instances that should be in the instance pool.
* `state` - The current state of the instance pool.
* `time_created` - The date and time the instance pool was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Import

InstancePools can be imported using the `id`, e.g.

```
$ terraform import oci_core_instance_pool.test_instance_pool "id"
```

