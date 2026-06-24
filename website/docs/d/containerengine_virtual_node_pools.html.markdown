---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_virtual_node_pools"
sidebar_current: "docs-oci-datasource-containerengine-virtual_node_pools"
description: |-
  Provides the list of Virtual Node Pools in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_virtual_node_pools
This data source provides the list of Virtual Node Pools in Oracle Cloud Infrastructure Container Engine service.

List all the virtual node pools in a compartment, and optionally filter by cluster.

## Example Usage

```hcl
data "oci_containerengine_virtual_node_pools" "test_virtual_node_pools" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cluster_id = oci_containerengine_cluster.test_cluster.id
	name = var.virtual_node_pool_name
	state = var.virtual_node_pool_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The OCID of the cluster.
* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) The name to filter on.
* `state` - (Optional) A virtual node pool lifecycle state to filter on. Can have multiple parameters of this name.


## Attributes Reference

The following attributes are exported:

* `virtual_node_pools` - The list of virtual_node_pools.

### VirtualNodePool Reference

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

