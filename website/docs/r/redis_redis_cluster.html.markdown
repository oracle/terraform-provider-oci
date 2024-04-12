---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster"
sidebar_current: "docs-oci-resource-redis-redis_cluster"
description: |-
  Provides the Redis Cluster resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_redis_cluster
This resource provides the Redis Cluster resource in Oracle Cloud Infrastructure Redis service.

Creates a new Redis cluster. A Redis cluster is a memory-based storage solution. For more information, see [OCI Caching Service with Redis](https://docs.cloud.oracle.com/iaas/Content/redis/home.htm).


## Example Usage

```hcl
resource "oci_redis_redis_cluster" "test_redis_cluster" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.redis_cluster_display_name
	node_count = var.redis_cluster_node_count
	node_memory_in_gbs = var.redis_cluster_node_memory_in_gbs
	software_version = var.redis_cluster_software_version
	subnet_id = oci_core_subnet.test_subnet.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	nsg_ids = var.redis_cluster_nsg_ids
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the Redis cluster.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `node_count` - (Required) (Updatable) The number of nodes in the Redis cluster.
* `node_memory_in_gbs` - (Required) (Updatable) The amount of memory allocated to the Redis cluster's nodes, in gigabytes.
* `nsg_ids` - (Optional) (Updatable) OCIDs of the NSGs to control access in the customer network
* `software_version` - (Required) The Redis version that the cluster is running.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster's subnet.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the Redis cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster.
* `lifecycle_details` - A message describing the current state in more detail. For example, the message might provide actionable information for a resource in `FAILED` state.
* `node_collection` - The collection of Redis cluster nodes.
	* `items` - Collection of node objects.
		* `display_name` - A user-friendly name of a Redis cluster node.
		* `private_endpoint_fqdn` - The fully qualified domain name (FQDN) of the API endpoint to access a specific node.
		* `private_endpoint_ip_address` - The private IP address of the API endpoint to access a specific node.
* `node_count` - The number of nodes in the Redis cluster.
* `node_memory_in_gbs` - The amount of memory allocated to the Redis cluster's nodes, in gigabytes.
* `nsg_ids` - OCIDs of the NSGs to control access in the customer network
* `primary_endpoint_ip_address` - The private IP address of the API endpoint for the Redis cluster's primary node.
* `primary_fqdn` - The fully qualified domain name (FQDN) of the API endpoint for the Redis cluster's primary node.
* `replicas_endpoint_ip_address` - The private IP address of the API endpoint for the Redis cluster's replica nodes.
* `replicas_fqdn` - The fully qualified domain name (FQDN) of the API endpoint for the Redis cluster's replica nodes.
* `software_version` - The Redis version that the cluster is running.
* `state` - The current state of the Redis cluster.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster's subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Redis cluster was created. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
* `time_updated` - The date and time the Redis cluster was updated. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Redis Cluster
	* `update` - (Defaults to 20 minutes), when updating the Redis Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Redis Cluster


## Import

RedisClusters can be imported using the `id`, e.g.

```
$ terraform import oci_redis_redis_cluster.test_redis_cluster "id"
```

