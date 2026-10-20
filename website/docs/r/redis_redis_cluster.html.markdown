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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/RedisCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Creates a new Oracle Cloud Infrastructure Cache cluster. A cluster is a memory-based storage solution.
You can optionally initialize the cluster data by restoring from an Oracle Cloud Infrastructure Cache Backup (backupId) or by importing from Object Storage RDB file(s) (importFromObjectStorageDetails).
For more information, see [OCI Cache](https://docs.cloud.oracle.com/iaas/Content/ocicache/home.htm).

## Cross-Region Replication Switchover

Switchover isn't supported in Terraform. To complete a switchover, use the OCI Console, CLI, or SDK.

After the switchover completes, update the Terraform configuration for both affected clusters before you run any further `terraform apply` operations:

1. For the new primary cluster (formerly the secondary), remove `primary_cluster_id` from the resource configuration.
2. For the new secondary cluster (formerly the primary), set `primary_cluster_id` to the OCID of the new primary cluster.
3. Run `terraform plan` for both cluster resources. Proceed only if the plan reports no changes, confirming that the Terraform configuration matches the new cluster topology.

If you don't update these configurations, a later Terraform apply might revert to the previous topology and unintentionally change or break the cross-region replication relationship.


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
	backup_id = oci_database_backup.test_backup.id
	cluster_mode = var.redis_cluster_cluster_mode
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	import_from_object_storage_details {
		#Required
		bucket = var.redis_cluster_import_from_object_storage_details_bucket
		namespace = var.redis_cluster_import_from_object_storage_details_namespace
		objects {
			#Required
			object = var.redis_cluster_import_from_object_storage_details_objects_object
		}
	}
	nsg_ids = var.redis_cluster_nsg_ids
	oci_cache_config_set_id = oci_redis_oci_cache_config_set.test_oci_cache_config_set.id
	primary_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
	security_attributes = var.redis_cluster_security_attributes
	shard_count = var.redis_cluster_shard_count
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Optional) The ID of the Oracle Cloud Infrastructure Cache Backup from which this cluster was created.Mutually exclusive with 'importFromObjectStorageDetails'.
* `cluster_mode` - (Optional) Specifies whether the cluster is sharded or non-sharded.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the cluster.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `import_from_object_storage_details` - (Optional) Details for importing Oracle Cloud Infrastructure Cache data from Object Storage RDB file(s) during cluster creation.
	* `bucket` - (Required) The Object Storage bucket name.
	* `namespace` - (Required) The Object Storage namespace name.
	* `objects` - (Required) The list of objects to import from the specified bucket.
		* `object` - (Required) The name of the object in the bucket (for example, 'customerA/exports/backup_ocid/dump.rdb').
* `node_count` - (Required) (Updatable) The number of nodes per shard in the cluster when clusterMode is SHARDED. This is the total number of nodes when clusterMode is NONSHARDED.
* `node_memory_in_gbs` - (Required) (Updatable) The amount of memory allocated to the cluster's nodes, in gigabytes.
* `nsg_ids` - (Optional) (Updatable) A list of Network Security Group (NSG) [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this cluster. For more information, see [Using an NSG for Clusters](https://docs.cloud.oracle.com/iaas/Content/ocicache/connecttocluster.htm#connecttocluster__networksecuritygroup). 
* `oci_cache_config_set_id` - (Optional) (Updatable) The ID of the corresponding Oracle Cloud Infrastructure Cache Config Set for the cluster.
* `primary_cluster_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the primary cluster from which data will be replicated. Setting it on a standalone cluster converts that cluster to a secondary cluster; removing it from a secondary cluster converts that cluster to standalone. Changing directly from one primary cluster to another is not supported: remove it and apply before setting a different primary cluster.
* `security_attributes` - (Optional) (Updatable) Security attributes for redis cluster resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `shard_count` - (Optional) (Updatable) The number of shards in sharded cluster. Only applicable when clusterMode is SHARDED.
* `software_version` - (Required) (Updatable) The Oracle Cloud Infrastructure Cache engine version that the cluster is running.
* `subnet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster's subnet.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `backup_id` - The ID of the Oracle Cloud Infrastructure Cache Backup from which this cluster was created.
* `cluster_mode` - Specifies whether the cluster is sharded or non-sharded.
* `cluster_replication_topology` - Defines the replication topology of an Oracle Cloud Infrastructure cache cluster, including the primary cluster and associated secondary clusters participating in replication.
	* `primary_cluster` - The details of a cluster participating in the replication setup.
		* `oci_cache_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Oracle Cloud Infrastructure Cache cluster.
		* `region` - The Oracle Cloud Infrastructure region to which the cluster belongs.
	* `secondary_clusters` - The list of secondary clusters that replicate data from the primary cluster.
		* `oci_cache_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the Oracle Cloud Infrastructure Cache cluster.
		* `region` - The Oracle Cloud Infrastructure region to which the cluster belongs.
* `cluster_role` - The current role of the cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the compartment that contains the cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `discovery_endpoint_ip_address` - The private IP address of the API endpoint for sharded cluster discovery.
* `discovery_fqdn` - The fully qualified domain name (FQDN) of the API endpoint for sharded cluster discovery.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
* `import_from_object_storage_details` - Details for importing Oracle Cloud Infrastructure Cache data from Object Storage RDB file(s) during cluster creation.
	* `bucket` - The Object Storage bucket name.
	* `namespace` - The Object Storage namespace name.
	* `objects` - The list of objects to import from the specified bucket.
		* `object` - The name of the object in the bucket (for example, 'customerA/exports/backup_ocid/dump.rdb').
* `lifecycle_details` - A message describing the current state in more detail. For example, the message might provide actionable information for a resource in `FAILED` state.
* `node_collection` - The collection of  cluster nodes.
	* `items` - Collection of node objects.
		* `display_name` - A user-friendly name of a cluster node.
		* `private_endpoint_fqdn` - The fully qualified domain name (FQDN) of the API endpoint to access a specific node.
		* `private_endpoint_ip_address` - The private IP address of the API endpoint to access a specific node.
* `node_count` - The number of nodes per shard in the cluster when clusterMode is SHARDED. This is the total number of nodes when clusterMode is NONSHARDED.
* `node_memory_in_gbs` - The amount of memory allocated to the cluster's nodes, in gigabytes.
* `nsg_ids` - A list of Network Security Group (NSG) [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) associated with this cluster. For more information, see [Using an NSG for Clusters](https://docs.cloud.oracle.com/iaas/Content/ocicache/connecttocluster.htm#connecttocluster__networksecuritygroup). 
* `oci_cache_config_set_id` - The ID of the corresponding Oracle Cloud Infrastructure Cache Config Set for the cluster.
* `primary_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the primary cluster in CRR.
* `primary_endpoint_ip_address` - The private IP address of the API endpoint for the cluster's primary node.
* `primary_fqdn` - The fully qualified domain name (FQDN) of the API endpoint for the cluster's primary node.
* `replicas_endpoint_ip_address` - The private IP address of the API endpoint for the cluster's replica nodes.
* `replicas_fqdn` - The fully qualified domain name (FQDN) of the API endpoint for the cluster's replica nodes.
* `security_attributes` - Security attributes for redis cluster resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "enforce"}}}` 
* `shard_count` - The number of shards in a sharded cluster. Only applicable when clusterMode is SHARDED.
* `software_version` - The Oracle Cloud Infrastructure Cache engine version that the cluster is running.
* `state` - The current state of the cluster.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster's subnet.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cluster was created. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
* `time_updated` - The date and time the cluster was updated. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.

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
