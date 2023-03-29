---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_virtual_node_pool"
sidebar_current: "docs-oci-resource-containerengine-virtual_node_pool"
description: |-
  Provides the Virtual Node Pool resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_virtual_node_pool
This resource provides the Virtual Node Pool resource in Oracle Cloud Infrastructure Container Engine service.

Create a new virtual node pool.

## Example Usage

```hcl
resource "oci_containerengine_virtual_node_pool" "test_virtual_node_pool" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compartment_id = var.compartment_id
	display_name = var.virtual_node_pool_display_name
	placement_configurations {

		#Required
		availability_domain = var.virtual_node_pool_placement_configurations_availability_domain
		fault_domain = var.virtual_node_pool_placement_configurations_fault_domain
		subnet_id = oci_core_subnet.test_subnet.id
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	initial_virtual_node_labels {

		#Optional
		key = var.virtual_node_pool_initial_virtual_node_labels_key
		value = var.virtual_node_pool_initial_virtual_node_labels_value
	}
	nsg_ids = var.virtual_node_pool_nsg_ids
	#Required
	pod_configuration {
		#Required
		shape = var.virtual_node_pool_pod_configuration_shape
		subnet_id = oci_core_subnet.test_subnet.id

		#Optional
		nsg_ids = var.virtual_node_pool_pod_configuration_nsg_ids
	}
	size = var.virtual_node_pool_size
	#Optional
	taints {

		#Optional
		effect = var.virtual_node_pool_taints_effect
		key = var.virtual_node_pool_taints_key
		value = var.virtual_node_pool_taints_value
	}
	virtual_node_tags {

		#Optional
		defined_tags = {"Operations.CostCenter"= "42"}
		freeform_tags = {"Department"= "Finance"}
	}
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The cluster the virtual node pool is associated with. A virtual node pool can only be associated with one cluster.
* `compartment_id` - (Required) Compartment of the virtual node pool.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `display_name` - (Required) (Updatable) Display name of the virtual node pool. This is a non-unique value.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `initial_virtual_node_labels` - (Optional) (Updatable) Initial labels that will be added to the Kubernetes Virtual Node object when it registers.
	* `key` - (Optional) (Updatable) The key of the pair.
	* `value` - (Optional) (Updatable) The value of the pair.
* `nsg_ids` - (Optional) (Updatable) List of network security group id's applied to the Virtual Node VNIC.
* `placement_configurations` - (Required) (Updatable) The list of placement configurations which determines where Virtual Nodes will be provisioned across as it relates to the subnet and availability domains. The size attribute determines how many we evenly spread across these placement configurations
	* `availability_domain` - (Required) (Updatable) The availability domain in which to place virtual nodes. Example: `Uocm:PHX-AD-1`
	* `fault_domain` - (Required) (Updatable) The fault domain of this virtual node.
	* `subnet_id` - (Required) (Updatable) The OCID of the subnet in which to place virtual nodes.
* `pod_configuration` - (Required) (Updatable) The pod configuration for pods run on virtual nodes of this virtual node pool.
	* `nsg_ids` - (Optional) (Updatable) List of network security group IDs applied to the Pod VNIC.
	* `shape` - (Required) (Updatable) Shape of the pods.
	* `subnet_id` - (Required) (Updatable) The regional subnet where pods' VNIC will be placed.
* `size` - (Required) (Updatable) The number of Virtual Nodes that should be in the Virtual Node Pool. The placement configurations determine where these virtual nodes are placed.
* `taints` - (Optional) (Updatable) A taint is a collection of <key, value, effect>. These taints will be applied to the Virtual Nodes of this Virtual Node Pool for Kubernetes scheduling.
	* `effect` - (Optional) (Updatable) The effect of the pair.
	* `key` - (Optional) (Updatable) The key of the pair.
	* `value` - (Optional) (Updatable) The value of the pair.
* `virtual_node_tags` - (Optional) (Updatable) The tags associated to the virtual nodes in this virtual node pool.
	* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `cluster_id` - The cluster the virtual node pool is associated with. A virtual node pool can only be associated with one cluster.
* `compartment_id` - Compartment of the virtual node pool.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
* `display_name` - Display name of the virtual node pool. This is a non-unique value.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`
* `id` - The OCID of the virtual node pool.
* `initial_virtual_node_labels` - Initial labels that will be added to the Kubernetes Virtual Node object when it registers. This is the same as virtualNodePool resources.
	* `key` - The key of the pair.
	* `value` - The value of the pair.
* `kubernetes_version` - The version of Kubernetes running on the nodes in the node pool.
* `lifecycle_details` - Details about the state of the Virtual Node Pool.
* `nsg_ids` - List of network security group id's applied to the Virtual Node VNIC.
* `placement_configurations` - The list of placement configurations which determines where Virtual Nodes will be provisioned across as it relates to the subnet and availability domains. The size attribute determines how many we evenly spread across these placement configurations
	* `availability_domain` - The availability domain in which to place virtual nodes. Example: `Uocm:PHX-AD-1`
	* `fault_domain` - The fault domain of this virtual node.
	* `subnet_id` - The OCID of the subnet in which to place virtual nodes.
* `pod_configuration` - The pod configuration for pods run on virtual nodes of this virtual node pool.
	* `nsg_ids` - List of network security group IDs applied to the Pod VNIC.
	* `shape` - Shape of the pods.
	* `subnet_id` - The regional subnet where pods' VNIC will be placed.
* `size` - The number of Virtual Nodes that should be in the Virtual Node Pool. The placement configurations determine where these virtual nodes are placed.
* `state` - The state of the Virtual Node Pool.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `taints` - A taint is a collection of <key, value, effect>. These taints will be applied to the Virtual Nodes of this Virtual Node Pool for Kubernetes scheduling.
	* `effect` - The effect of the pair.
	* `key` - The key of the pair.
	* `value` - The value of the pair.
* `time_created` - The time the virtual node pool was created.
* `time_updated` - The time the virtual node pool was updated.
* `virtual_node_tags` - The tags associated to the virtual nodes in this virtual node pool.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}`
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Virtual Node Pool
	* `update` - (Defaults to 20 minutes), when updating the Virtual Node Pool
	* `delete` - (Defaults to 20 minutes), when destroying the Virtual Node Pool


## Import

VirtualNodePools can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_virtual_node_pool.test_virtual_node_pool "id"
```

