---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_backup_export_to_object_storage"
sidebar_current: "docs-oci-resource-redis-oci_cache_backup_export_to_object_storage"
description: |-
  Provides the Oci Cache Backup Export To Object Storage resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_backup_export_to_object_storage
This resource provides the Oci Cache Backup Export To Object Storage resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/ocicache/latest/OciCacheBackup/ExportToObjectStorage

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Initiates an asynchronous export of the backup’s RDB file(s) to the specified Object Storage bucket. The service generates the object names. For sharded backups, one object is written per shard under the optional prefix.


## Example Usage

```hcl
resource "oci_redis_oci_cache_backup_export_to_object_storage" "test_oci_cache_backup_export_to_object_storage" {
	#Required
	bucket = var.oci_cache_backup_export_to_object_storage_bucket
	namespace = var.oci_cache_backup_export_to_object_storage_namespace
	oci_cache_backup_id = oci_redis_oci_cache_backup.test_oci_cache_backup.id

	#Optional
	prefix = var.oci_cache_backup_export_to_object_storage_prefix
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) The target Object Storage bucket name.
* `namespace` - (Required) The Object Storage namespace name.
* `oci_cache_backup_id` - (Required) Unique Oracle Cloud Infrastructure Cache Backup identifier.
* `prefix` - (Optional) Optional prefix under which the service will place the exported object(s).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache Backup Export To Object Storage
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache Backup Export To Object Storage
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache Backup Export To Object Storage


## Import

Import is not supported for this resource.

