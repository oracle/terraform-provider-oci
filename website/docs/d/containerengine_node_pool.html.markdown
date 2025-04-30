---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_node_pool"
sidebar_current: "docs-oci-datasource-containerengine-node_pool"
description: |-
  Provides details about a specific Node Pool in Oracle Cloud Infrastructure Container Engine service
---

# Data Source: oci_containerengine_node_pool
This data source provides details about a specific Node Pool resource in Oracle Cloud Infrastructure Container Engine service.

Get the details of a node pool.

## Example Usage

```hcl
data "oci_containerengine_node_pool" "test_node_pool" {
	#Required
	node_pool_id = oci_containerengine_node_pool.test_node_pool.id
}
```

## Argument Reference

The following arguments are supported:

* `node_pool_id` - (Required) The OCID of the node pool.


## Attributes Reference

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
* `node_metadata` - A list of key/value pairs to add to each underlying Oracle Cloud Infrastructure instance in the node pool on launch.
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
* `nodes` - The nodes in the node pool.
	* `availability_domain` - The name of the availability domain in which this node is placed.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
	* `error` - An error that may be associated with the node.
		* `code` - A short error code that defines the upstream error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
		* `message` - A human-readable error string of the upstream error.
		* `status` - The status of the HTTP response encountered in the upstream error.
	* `fault_domain` - The fault domain of this node.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
	* `id` - The OCID of the compute instance backing this node.
	* `kubernetes_version` - The version of Kubernetes this node is running.
	* `lifecycle_details` - Details about the state of the node.
	* `name` - The name of the node.
	* `node_pool_id` - The OCID of the node pool to which this node belongs.
	* `private_ip` - The private IP address of this node.
	* `public_ip` - The public IP address of this node.
	* `state` - The state of the node.
	* `subnet_id` - The OCID of the subnet in which this node is placed.
* `quantity_per_subnet` - The number of nodes in each subnet.
* `ssh_public_key` - The SSH public key on each node in the node pool on launch.
* `state` - The state of the nodepool.
* `subnet_ids` - The OCIDs of the subnets in which to place nodes for this node pool.

