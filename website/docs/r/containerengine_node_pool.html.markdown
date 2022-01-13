---
subcategory: "Container Engine"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_containerengine_node_pool"
sidebar_current: "docs-oci-resource-containerengine-node_pool"
description: |-
  Provides the Node Pool resource in Oracle Cloud Infrastructure Container Engine service
---

# oci_containerengine_node_pool
This resource provides the Node Pool resource in Oracle Cloud Infrastructure Container Engine service.

Create a new node pool.

## Example Usage

```hcl
resource "oci_containerengine_node_pool" "test_node_pool" {
	#Required
	cluster_id = oci_containerengine_cluster.test_cluster.id
	compartment_id = var.compartment_id
	kubernetes_version = var.node_pool_kubernetes_version
	name = var.node_pool_name
	node_shape = var.node_pool_node_shape
	subnet_ids = var.node_pool_subnet_ids

	#Optional
	initial_node_labels {

		#Optional
		key = var.node_pool_initial_node_labels_key
		value = var.node_pool_initial_node_labels_value
	}
	node_config_details {
		#Required
		placement_configs {
			#Required
			availability_domain = var.node_pool_node_config_details_placement_configs_availability_domain
			subnet_id = oci_core_subnet.test_subnet.id
		}
		size = var.node_pool_node_config_details_size

		#Optional
		is_pv_encryption_in_transit_enabled = var.node_pool_node_config_details_is_pv_encryption_in_transit_enabled
		kms_key_id = oci_kms_key.test_key.id
		nsg_ids = var.node_pool_node_config_details_nsg_ids
	}
	node_image_name = oci_core_image.test_image.name
	node_metadata = var.node_pool_node_metadata
	node_shape_config {

		#Optional
		memory_in_gbs = var.node_pool_node_shape_config_memory_in_gbs
		ocpus = var.node_pool_node_shape_config_ocpus
	}
	node_source_details {
		#Required
		image_id = oci_core_image.test_image.id
		source_type = var.node_pool_node_source_details_source_type

		#Optional
		boot_volume_size_in_gbs = var.node_pool_node_source_details_boot_volume_size_in_gbs
	}
	quantity_per_subnet = var.node_pool_quantity_per_subnet
	ssh_public_key = var.node_pool_ssh_public_key
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required) The OCID of the cluster to which this node pool is attached.
* `compartment_id` - (Required) The OCID of the compartment in which the node pool exists.
* `initial_node_labels` - (Optional) (Updatable) A list of key/value pairs to add to nodes after they join the Kubernetes cluster.
	* `key` - (Optional) (Updatable) The key of the pair.
	* `value` - (Optional) (Updatable) The value of the pair.
* `kubernetes_version` - (Required) (Updatable) The version of Kubernetes to install on the nodes in the node pool.
* `name` - (Required) (Updatable) The name of the node pool. Avoid entering confidential information.
* `node_config_details` - (Optional) (Updatable) The configuration of nodes in the node pool. Exactly one of the subnetIds or nodeConfigDetails properties must be specified. 
	* `is_pv_encryption_in_transit_enabled` - (Optional) (Updatable) Whether to enable in-transit encryption for the data volume's paravirtualized attachment. This field applies to both block volumes and boot volumes. The default value is false.
	* `kms_key_id` - (Optional) (Updatable) The OCID of the Key Management Service key assigned to the boot volume.
	* `nsg_ids` - (Optional) (Updatable) The OCIDs of the Network Security Group(s) to associate nodes for this node pool with. For more information about NSGs, see [NetworkSecurityGroup](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/NetworkSecurityGroup/). 
	* `placement_configs` - (Required) (Updatable) The placement configurations for the node pool. Provide one placement configuration for each availability domain in which you intend to launch a node.

		To use the node pool with a regional subnet, provide a placement configuration for each availability domain, and include the regional subnet in each placement configuration. 
		* `availability_domain` - (Required) (Updatable) The availability domain in which to place nodes. Example: `Uocm:PHX-AD-1` 
		* `subnet_id` - (Required) (Updatable) The OCID of the subnet in which to place nodes.
	* `size` - (Required) (Updatable) The number of nodes that should be in the node pool. 
* `node_image_id` - (Optional) Deprecated. Use `node_source_details` instead. The OCID of the image running on the nodes in the node pool. Cannot be used when `node_image_name` is specified.
* `node_image_name` - (Optional) Deprecated. Use `nodeSourceDetails` instead. If you specify values for both, this value is ignored. The name of the image running on the nodes in the node pool. Cannot be used when `node_image_id` is specified.
* `node_metadata` - (Optional) (Updatable) A list of key/value pairs to add to each underlying Oracle Cloud Infrastructure instance in the node pool on launch.
* `node_shape` - (Required) (Updatable) The name of the node shape of the nodes in the node pool.
* `node_shape_config` - (Optional) (Updatable) Specify the configuration of the shape to launch nodes in the node pool. 
	* `memory_in_gbs` - (Optional) (Updatable) The total amount of memory available to each node, in gigabytes. 
	* `ocpus` - (Optional) (Updatable) The total number of OCPUs available to each node in the node pool. See [here](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/iaas/20160918/Shape/) for details. 
* `node_source_details` - (Optional) (Updatable) Specify the source to use to launch nodes in the node pool. Currently, image is the only supported source. 
	* `boot_volume_size_in_gbs` - (Optional) (Updatable) The size of the boot volume in GBs. Minimum value is 50 GB. See [here](https://docs.cloud.oracle.com/en-us/iaas/Content/Block/Concepts/bootvolumes.htm) for max custom boot volume sizing and OS-specific requirements.
	* `image_id` - (Required) (Updatable) The OCID of the image used to boot the node.
	* `source_type` - (Required) (Updatable) The source type for the node. Use `IMAGE` when specifying an OCID of an image. 
* `quantity_per_subnet` - (Optional) (Updatable) Optional, default to 1. The number of nodes to create in each subnet specified in subnetIds property. When used, subnetIds is required. This property is deprecated, use nodeConfigDetails instead. 
* `ssh_public_key` - (Optional) (Updatable) The SSH public key on each node in the node pool on launch.
* `subnet_ids` - (Optional) (Updatable) The OCIDs of the subnets in which to place nodes for this node pool. When used, quantityPerSubnet can be provided. This property is deprecated, use nodeConfigDetails. Exactly one of the subnetIds or nodeConfigDetails properties must be specified. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 50 minutes), when creating the Node Pool
	* `update` - (Defaults to 50 minutes), when updating the Node Pool
	* `delete` - (Defaults to 50 minutes), when destroying the Node Pool


## Import

NodePools can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_node_pool.test_node_pool "id"
```

