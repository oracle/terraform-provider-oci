---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster_detach_oci_cache_user"
sidebar_current: "docs-oci-resource-redis-redis_cluster_detach_oci_cache_user"
description: |-
  Provides the Redis Cluster Detach Oci Cache User resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_redis_cluster_detach_oci_cache_user
This resource provides the Redis Cluster Detach Oci Cache User resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/RedisClusterDetachOciCacheUser

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Detach existing Oracle Cloud Infrastructure cache users to a redis cluster.

## Example Usage

```hcl
resource "oci_redis_redis_cluster_detach_oci_cache_user" "test_redis_cluster_detach_oci_cache_user" {
	#Required
	oci_cache_users = var.redis_cluster_detach_oci_cache_user_oci_cache_users
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `oci_cache_users` - (Required) List of Oracle Cloud Infrastructure cache user unique IDs (OCIDs).
* `redis_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Redis Cluster Detach Oci Cache User
	* `update` - (Defaults to 20 minutes), when updating the Redis Cluster Detach Oci Cache User
	* `delete` - (Defaults to 20 minutes), when destroying the Redis Cluster Detach Oci Cache User


## Import

Import is not supported for this resource.

