---
layout: "oci"
page_title: "OCI: oci_containerengine_node_pool"
sidebar_current: "docs-oci-resource-containerengine-node_pool"
description: |-
  Creates and manages an OCI NodePool
---

# oci_containerengine_node_pool
The `oci_containerengine_node_pool` resource creates and manages an OCI NodePool

Create a new node pool.

## Example Usage

```hcl
resource "oci_containerengine_node_pool" "test_node_pool" {
	#Required
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"
	compartment_id = "${var.compartment_id}"
	kubernetes_version = "${var.node_pool_kubernetes_version}"
	name = "${var.node_pool_name}"
	node_image_name = "${var.node_pool_node_image_name}"
	node_shape = "${var.node_pool_node_shape}"
	subnet_ids = "${var.node_pool_subnet_ids}"

	#Optional
	initial_node_labels {

		#Optional
		key = "${var.node_pool_initial_node_labels_key}"
		value = "${var.node_pool_initial_node_labels_value}"
	}
	quantity_per_subnet = "${var.node_pool_quantity_per_subnet}"
	ssh_public_key = "${var.node_pool_ssh_public_key}"
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
* `node_image_name` - (Required) The name of the image running on the nodes in the node pool.
* `node_shape` - (Required) The name of the node shape of the nodes in the node pool.
* `quantity_per_subnet` - (Optional) (Updatable) The number of nodes to create in each subnet.
* `ssh_public_key` - (Optional) The SSH public key to add to each node in the node pool.
* `subnet_ids` - (Required) (Updatable) The OCIDs of the subnets in which to place nodes for this node pool.


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
* `node_image_id` - The OCID of the image running on the nodes in the node pool.
* `node_image_name` - The name of the image running on the nodes in the node pool.
* `node_shape` - The name of the node shape of the nodes in the node pool.
* `nodes` - The nodes in the node pool.
	* `availability_domain` - The name of the availability domain in which this node is placed.
	* `error` - An error that may be associated with the node.
		* `code` - A short error code that defines the error, meant for programmatic parsing. See [API Errors](https://docs.us-phoenix-1.oraclecloud.com/Content/API/References/apierrors.htm).
		* `message` - A human-readable error string.
	* `id` - The OCID of the compute instance backing this node.
	* `lifecycle_details` - Details about the state of the node.
	* `name` - The name of the node.
	* `node_pool_id` - The OCID of the node pool to which this node belongs.
	* `public_ip` - The public IP address of this node.
	* `state` - The state of the node.
	* `subnet_id` - The OCID of the subnet in which this node is placed.
* `quantity_per_subnet` - The number of nodes in each subnet.
* `ssh_public_key` - The SSH public key on each node in the node pool.
* `subnet_ids` - The OCIDs of the subnets in which to place nodes for this node pool.

## Import

NodePools can be imported using the `id`, e.g.

```
$ terraform import oci_containerengine_node_pool.test_node_pool "id"
```
