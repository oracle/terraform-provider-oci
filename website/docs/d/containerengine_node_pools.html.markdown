---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_node_pools"
sidebar_current: "docs-oci-datasource-containerengine-node_pools"
description: |-
  Provides the list of Node Pools in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_node_pools
This data source provides the list of Node Pools in Oracle Cloud Infrastructure Container Engine service.

List all the node pools in a compartment, and optionally filter by cluster.

## Example Usage

```hcl
data "oci_containerengine_node_pools" "test_node_pools" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	cluster_id = oci_containerengine_cluster.test_cluster.id
	name = var.node_pool_name
	state = var.node_pool_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The OCID of the cluster.
* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) The name to filter on.
* `state` - (Optional) A list of nodepool lifecycle states on which to filter on, matching any of the list items (OR logic). eg. [ACTIVE, DELETING]


## Attributes Reference

The following attributes are exported:

* `node_pools` - The list of node_pools.

### NodePool Reference

The following attributes are exported:

* `cluster_id` - The OCID of the cluster to which this node pool is attached.
* `compartment_id` - The OCID of the compartment in which the node pool exists.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the node pool.
* `initial_node_labels` - A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	* `key` - The key of the pair.
	* `value` - The value of the pair.
* `kubernetes_version` - The version of Kubernetes running on the nodes in the node pool.
* `lifecycle_details` - Details about the state of the nodepool.
* `name` - The name of the node pool.
* `network_launch_type` - Emulation type for the physical network interface card (NIC) for nodes
* `node_config_details` - The configuration of nodes in the node pool.
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	* `kms_key_id` - The OCID of the Key Management Service key assigned to the boot volume.
	* `node_pool_pod_network_option_details` - The CNI related configuration of pods in the node pool. 
		* `cni_type` - The CNI plugin used by this node pool
		* `max_pods_per_node` - The max number of pods per node in the node pool. This value will be limited by the number of VNICs attachable to the node pool shape 
		* `pod_nsg_ids` - The OCIDs of the Network Security Group(s) to associate pods for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
		* `pod_subnet_ids` - The OCIDs of the subnets in which to place pods for this node pool. This can be one of the node pool subnet IDs 
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `nsg_ids` - The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `placement_configs` - The placement configurations for the node pool. Provide one placement configuration for each availability domain in which you intend to launch a node.

		To use the node pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
		* `availability_domain` - The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1` 
		* `capacity_reservation_id` - The OCID of the compute capacity reservation in which to place the compute instance.
		* `fault_domains` - A list of fault domains in which to place nodes. 
		* `preemptible_node_config` - Configuration options for preemptible nodes.
			* `preemption_action` - The action to run when the preemptible node is interrupted for eviction.
				* `is_preserve_boot_volume` - Whether to preserve the boot volume that was used to launch the preemptible instance when the instance is terminated. Defaults to false if not specified. 
				* `type` - The type of action to run when the instance is interrupted for eviction.
		* `subnet_id` - The OCID of the subnet in which to place nodes.
	* `size` - The number of nodes in the node pool. 
* `node_eviction_node_pool_settings` - Node Eviction Details configuration
	* `eviction_grace_duration` - Duration after which OKE will give up eviction of the pods on the node. PT0M will indicate you want to delete the node without cordon and drain. Default PT60M, Min PT0M, Max: PT60M. Format ISO 8601 e.g PT30M 
	* `is_force_action_after_grace_duration` - If the node action should be performed if not all the pods can be evicted in the grace period
	* `is_force_delete_after_grace_duration` - If the underlying compute instance should be deleted if you cannot evict all the pods in grace period
* `node_image_id` - Deprecated. see `nodeSource`. The OCID of the image running on the nodes in the node pool. 
* `node_image_name` - Deprecated. see `nodeSource`. The name of the image running on the nodes in the node pool. 
* `node_pool_cycling_details` - Node Pool Cycling Details
	* `cycle_modes` - An ordered list of cycle modes that should be performed on the OKE nodes.  
	* `is_node_cycling_enabled` - If cycling operation should be performed on the nodes in the node pool.
	* `maximum_surge` - Maximum additional new compute instances that would be temporarily created and added to nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 1, Ranges from 0 to Nodepool size or 0% to 100% 
	* `maximum_unavailable` - Maximum active nodes that would be terminated from nodepool during the cycling nodepool process. OKE supports both integer and percentage input. Defaults to 0, Ranges from 0 to Nodepool size or 0% to 100% 
* `node_shape` - The name of the node shape of the nodes in the node pool.
* `node_shape_config` - The shape configuration of the nodes.
	* `memory_in_gbs` - The total amount of memory available to each node, in gigabytes. 
	* `ocpus` - The total number of OCPUs available to each node in the node pool. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `node_source` - Deprecated. see `nodeSourceDetails`. Source running on the nodes in the node pool.
	* `image_id` - The OCID of the image.
	* `source_name` - The user-friendly name of the entity corresponding to the OCID. 
	* `source_type` - The source type of this option. `IMAGE` means the OCID is of an image. 
* `node_source_details` - Source running on the nodes in the node pool.
	* `boot_volume_size_in_gbs` - The size of the boot volume in GBs. Minimum value is 50 GB. See [here](https://docs.cloud.oracle.com/en-us/iaas/Content/Block/Concepts/bootvolumes.htm) for max custom boot volume sizing and OS-specific requirements.
	* `image_id` - The OCID of the image used to boot the node.
	* `source_type` - The source type for the node. Use `IMAGE` when specifying an OCID of an image.
* `quantity_per_subnet` - The number of nodes in each subnet.
* `secondary_vnics` - A list of secondary vnics to attach to nodes
	* `create_vnic_details` - The properties of the secondary vnics
		* `application_resources` - The application resource that corresponds to this secondary vnic. Used to map pods to this specific vnic for scheduling 
		* `assign_ipv6ip` - Whether to allocate an IPv6 address at instance and VNIC creation from an IPv6 enabled subnet 
		* `assign_public_ip` - Whether the VNIC should be assigned a public IP address 
		* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
		* `display_name` - Display name for secondary vnic
		* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
		* `ip_count` - The number of ip addresses to attach to secondary vnic 
		* `ipv6address_ipv6subnet_cidr_pair_details` - A list of IPv6 prefixes from which the VNIC should be assigned an IPv6 address. You can provide only the prefix  and Oracle Cloud Infrastructure selects an available address from the range. You can optionally choose to leave the prefix range empty  and instead provide the specific IPv6 address that should be used from within that range. 
			* `ipv6address` - An IPv6 address of your choice. Must be an available IPv6 address within the subnet's prefix 
			* `ipv6subnet_cidr` - The IPv6 prefix allocated to the subnet 
		* `nsg_ids` - A list of the OCIDs of the network security groups (NSGs) to add the VNIC to 
		* `skip_source_dest_check` - Whether the source/destination check is disabled on the VNIC 
		* `subnet_id` - the ocid of the subnet to create the vnic in
	* `display_name` - Display name for vnic attachment
	* `nic_index` - Which physical network interface card (NIC) the VNIC will use
* `ssh_public_key` - The SSH public key on each node in the node pool on launch.
* `state` - The state of the nodepool.
* `subnet_ids` - The OCIDs of the subnets in which to place nodes for this node pool.

