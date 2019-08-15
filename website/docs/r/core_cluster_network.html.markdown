---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cluster_network"
sidebar_current: "docs-oci-resource-core-cluster_network"
description: |-
  Provides the Cluster Network resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cluster_network
This resource provides the Cluster Network resource in Oracle Cloud Infrastructure Core service.

Create a cluster network.

## Example Usage

```hcl
resource "oci_core_cluster_network" "test_cluster_network" {
	#Required
	compartment_id = "${var.compartment_id}"
	instance_pools {
		#Required
		instance_configuration_id = "${oci_core_instance_configuration.test_instance_configuration.id}"
		size = "${var.cluster_network_instance_pools_size}"

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		display_name = "${var.cluster_network_instance_pools_display_name}"
		freeform_tags = {"Department"= "Finance"}
	}
	placement_configuration {
		#Required
		availability_domain = "${var.cluster_network_placement_configuration_availability_domain}"
		primary_subnet_id = "${oci_core_primary_subnet.test_primary_subnet.id}"

		#Optional
		secondary_vnic_subnets {
			#Required
			subnet_id = "${oci_core_subnet.test_subnet.id}"

			#Optional
			display_name = "${var.cluster_network_placement_configuration_secondary_vnic_subnets_display_name}"
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.cluster_network_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the cluster network.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The user-friendly name.  Does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `instance_pools` - (Required) the data to create the instance pools in the cluster network.
	* `compartment_id` - (Optional) (Updatable) The OCID of the compartment containing the instance pool.
	* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - (Optional) The user-friendly name.  Does not have to be unique.
	* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `instance_configuration_id` - (Required) The OCID of the instance configuration associated to the instance pool.
	* `size` - (Required) The number of instances that should be in the instance pool.
* `placement_configuration` - (Required) the placement data for the instance pools in the cluster network
	* `availability_domain` - (Required) The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - (Required) The OCID of the primary subnet to place instances.
	* `secondary_vnic_subnets` - (Optional) The set of secondary VNIC data for instances in the pool.
		* `display_name` - (Optional) The display name of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
		* `subnet_id` - (Required) The subnet OCID for the secondary vnic


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the cluster netowrk.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name.  Does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the cluster network.
* `instance_pools` - the instance pools in the cluster network.
	* `compartment_id` - The OCID of the compartment containing the instance pool.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
	* `display_name` - The user-friendly name.  Does not have to be unique.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the instance pool.
	* `instance_configuration_id` - The OCID of the instance configuration associated with the instance pool.
	* `load_balancers` - The load balancers attached to the instance pool. 
		* `backend_set_name` - The name of the backend set on the load balancer.
		* `id` - The OCID of the load balancer attachment.
		* `instance_pool_id` - The OCID of the instance pool of the load balancer attachment.
		* `load_balancer_id` - The OCID of the load balancer attached to the instance pool.
		* `port` - The port value used for the backends.
		* `state` - The status of the interaction between the instance pool and the load balancer.
		* `vnic_selection` - Indicates which VNIC on each instance in the instance pool should be used to associate with the load balancer. Possible values are "PrimaryVnic" or the display name of one of the secondary VNICs on the instance configuration that is associated with the instance pool.
	* `placement_configurations` - The placement configurations for the instance pool.
		* `availability_domain` - The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
		* `primary_subnet_id` - The OCID of the primary subnet to place instances.
		* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
			* `display_name` - The display name of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
			* `subnet_id` - The subnet OCID for the secondary vnic
	* `size` - The number of instances that should be in the instance pool.
	* `state` - The current state of the instance pool.
	* `time_created` - The date and time the instance pool was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `placement_configuration` - the placement data for the intance pools in the cluster network
	* `availability_domain` - The availability domain to place instances. Example: `Uocm:PHX-AD-1` 
	* `primary_subnet_id` - The OCID of the primary subnet to place instances.
	* `secondary_vnic_subnets` - The set of secondary VNIC data for instances in the pool.
		* `display_name` - The display name of the vnic. This is also use to match against the Instance Configuration defined secondary vnic. 
		* `subnet_id` - The subnet OCID for the secondary vnic
* `state` - The current state of the cluster network.
* `time_created` - The date and time the resource was created, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the resource was updated, in the format defined by RFC3339. Example: `2016-08-25T21:10:29.600Z` 

## Import

ClusterNetworks can be imported using the `id`, e.g.

```
$ terraform import oci_core_cluster_network.test_cluster_network "id"
```

