---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster_nodes"
sidebar_current: "docs-oci-datasource-redis-redis_cluster_nodes"
description: |-
  Provides the list of Redis Cluster Nodes in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_redis_cluster_nodes
This data source provides the list of Redis Cluster Nodes in Oracle Cloud Infrastructure Redis service.

Gets the list of all nodes in a cluster.


## Example Usage

```hcl
data "oci_redis_redis_cluster_nodes" "test_redis_cluster_nodes" {
	#Required
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

	#Optional
	display_name = var.redis_cluster_node_display_name
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `redis_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.


## Attributes Reference

The following attributes are exported:

* `redis_node_collection` - The list of redis_node_collection.

### RedisClusterNode Reference

The following attributes are exported:

* `items` - The list of nodes in a cluster.
	* `display_name` - A user-friendly name of a cluster node.
	* `private_endpoint_fqdn` - The fully qualified domain name (FQDN) of the API endpoint to access a specific node.
	* `private_endpoint_ip_address` - The private IP address of the API endpoint to access a specific node.
	* `redis_cluster_id` - The OCID of the cluster
	* `shard_number` - The shard number to which the node belongs to.

