---
layout: "oci"
page_title: "OCI: oci_containerengine_node_pool"
sidebar_current: "docs-oci-datasource-containerengine-node_pool"
description: |-
  Provides details about a specific NodePool
---

# Data Source: oci_containerengine_node_pool
The `oci_containerengine_node_pool` data source provides details about a specific NodePool

Get the details of a node pool.

## Example Usage

```hcl
data "oci_containerengine_node_pool" "test_node_pool" {
	#Required
	node_pool_id = "${oci_containerengine_node_pool.test_node_pool.id}"
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

