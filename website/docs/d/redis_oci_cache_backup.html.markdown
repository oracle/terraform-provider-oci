---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_backup"
sidebar_current: "docs-oci-datasource-redis-oci_cache_backup"
description: |-
  Provides details about a specific Oci Cache Backup in Oracle Cloud Infrastructure Redis service
---

# Data Source: oci_redis_oci_cache_backup
This data source provides details about a specific Oci Cache Backup resource in Oracle Cloud Infrastructure Redis service.

Gets an Oracle Cloud Infrastructure Cache Backup by identifier

## Example Usage

```hcl
data "oci_redis_oci_cache_backup" "test_oci_cache_backup" {
	#Required
	oci_cache_backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup.id
}
```

## Argument Reference

The following arguments are supported:

* `oci_cache_backup_id` - (Required) Unique Oracle Cloud Infrastructure Cache Backup identifier.


## Attributes Reference

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

