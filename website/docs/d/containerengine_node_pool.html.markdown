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
* `id` - The OCID of the node pool.
* `initial_node_labels` - A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	* `key` - The key of the pair.
	* `value` - The value of the pair.
* `kubernetes_version` - The version of Kubernetes running on the nodes in the node pool.
* `name` - The name of the node pool.
* `node_config_details` - The configuration of nodes in the node pool.
	* `is_pv_encryption_in_transit_enabled` - Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	* `kms_key_id` - The OCID of the Key Management Service key assigned to the boot volume.
	* `nsg_ids` - The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `placement_configs` - The placement configurations for the node pool. Provide one placement configuration for each availability domain in which you intend to launch a node.

		To use the node pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
		* `availability_domain` - The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1` 
		* `subnet_id` - The OCID of the subnet in which to place nodes.
	* `size` - The number of nodes in the node pool. 
* `node_image_id` - Deprecated. see `nodeSource`. The OCID of the image running on the nodes in the node pool. 
* `node_image_name` - Deprecated. see `nodeSource`. The name of the image running on the nodes in the node pool. 
* `node_metadata` - A list of key/value pairs to add to each underlying Oracle Cloud Infrastructure instance in the node pool on launch.
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
	* `error` - An error that may be associated with the node.
		* `code` - A short error code that defines the upstream error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
		* `message` - A human-readable error string of the upstream error.
		* `status` - The status of the HTTP response encountered in the upstream error.
	* `fault_domain` - The fault domain of this node.
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
* `subnet_ids` - The OCIDs of the subnets in which to place nodes for this node pool.

