---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_backups"
sidebar_current: "docs-oci-datasource-redis-oci_cache_backups"
description: |-
  Provides the list of Oci Cache Backups in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_backups
This data source provides the list of Oci Cache Backups in Oracle Cloud Infrastructure Redis service.

Returns a list of Oracle Cloud Infrastructure Cache Backups.

## Example Usage

```hcl
data "oci_redis_oci_cache_backups" "test_oci_cache_backups" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.oci_cache_backup_display_name
	oci_cache_backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup.id
	source_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
	state = var.oci_cache_backup_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `oci_cache_backup_id` - (Optional) Unique Oracle Cloud Infrastructure Cache Backup identifier.
* `source_cluster_id` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Backup resources, whose source cluster ID matches with the given source cluster ID.
* `state` - (Optional) A filter to return the Oracle Cloud Infrastructure Cache Backup resources, whose lifecycle state matches with the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `oci_cache_backup_collection` - The list of oci_cache_backup_collection.

### OciCacheBackup Reference

The following attributes are exported:

* `backup_size_in_gbs` - Backup size in GB.
* `backup_source` - Specifies whether the backup was created from a replica or primary node
* `backup_type` - Backup Type.
* `cluster_memory_in_gbs` - The amount of memory allocated to the cluster, in gigabytes.
* `cluster_mode` - Specifies whether the cluster is sharded or non-sharded.
* `compartment_id` - Backup compartment identifier
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Backup description
* `display_name` - Backup display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - Unique identifier that is immutable on creation
* `retention_period_in_days` - Backup retention period in days.
* `shard_count` - The number of shards in a sharded cluster. Only applicable when clusterMode is SHARDED.
* `software_version` - The Oracle Cloud Infrastructure Cache engine version that the cluster is running.
* `source_cluster_id` - The source Oracle Cloud Infrastructure Cache Cluster OCID.
* `state` - The current state of the backup.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the backup was created. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
* `time_updated` - The date and time the backup was updated. An [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.

