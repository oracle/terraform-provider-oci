---
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
	compartment_id = "${var.compartment_id}"

	#Optional
	cluster_id = "${oci_containerengine_cluster.test_cluster.id}"
	name = "${var.node_pool_name}"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Optional) The OCID of the cluster.
* `compartment_id` - (Required) The OCID of the compartment.
* `name` - (Optional) The name to filter on.


## Attributes Reference

The following attributes are exported:

* `node_pools` - The list of node_pools.

### NodePool Reference

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
		* `code` - A short error code that defines the upstream error, meant for programmatic parsing. See [API Errors](https://docs.cloud.oracle.com/iaas/Content/API/References/apierrors.htm).
		* `message` - A human-readable error string of the upstream error.
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

