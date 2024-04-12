---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster"
sidebar_current: "docs-oci-datasource-redis-redis_cluster"
description: |-
  Provides details about a specific Redis Cluster in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_redis_cluster
This data source provides details about a specific Redis Cluster resource in Oracle Cloud Infrastructure Redis service.

Retrieves the specified Redis cluster. A Redis cluster is a memory-based storage solution. For more information, see [OCI Caching Service with Redis](https://docs.cloud.oracle.com/iaas/Content/redis/home.htm).

## Example Usage

```hcl
data "oci_redis_redis_cluster" "test_redis_cluster" {
	#Required
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `redis_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Redis cluster.


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

