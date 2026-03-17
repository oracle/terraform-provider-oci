---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_backup"
sidebar_current: "docs-oci-resource-redis-oci_cache_backup"
description: |-
  Provides the Oci Cache Backup resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_backup
This resource provides the Oci Cache Backup resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/ocicache/latest/OciCacheBackup

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Creates a new Oracle Cloud Infrastructure Cache Backup.

## Example Usage

```hcl
resource "oci_redis_oci_cache_backup" "test_oci_cache_backup" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.oci_cache_backup_display_name
	source_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

	#Optional
	backup_source = var.oci_cache_backup_backup_source
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.oci_cache_backup_description
	freeform_tags = {"bar-key"= "value"}
	retention_period_in_days = var.oci_cache_backup_retention_period_in_days
}
```

## Argument Reference

The following arguments are supported:

* `backup_source` - (Optional) Specifies whether the backup was created from a replica or primary node
* `compartment_id` - (Required) (Updatable) Compartment identifier
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) Backup description
* `display_name` - (Required) (Updatable) Backup display name.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `retention_period_in_days` - (Optional) (Updatable) Backup retention period in days.
* `source_cluster_id` - (Required) Oracle Cloud Infrastructure Cache cluster identifier
* `export_to_object_storage_trigger` - (Optional) (Updatable) An optional property when incremented triggers Export To Object Storage. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache Backup
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache Backup


## Import

OciCacheBackups can be imported using the `id`, e.g.

```
$ terraform import oci_redis_oci_cache_backup.test_oci_cache_backup "id"
```

