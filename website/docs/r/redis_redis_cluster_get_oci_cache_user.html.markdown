---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster_get_oci_cache_user"
sidebar_current: "docs-oci-resource-redis-redis_cluster_get_oci_cache_user"
description: |-
  Provides the Redis Cluster Get Oci Cache User resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_redis_cluster_get_oci_cache_user
This resource provides the Redis Cluster Get Oci Cache User resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/RedisClusterGetOciCacheUser

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Gets a list of associated Oracle Cloud Infrastructure cache users for a redis cluster.

## Example Usage

```hcl
resource "oci_redis_redis_cluster_get_oci_cache_user" "test_redis_cluster_get_oci_cache_user" {
	#Required
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.redis_cluster_get_oci_cache_user_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `redis_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `oci_cache_user_id` - OCID of the OciCacheUser 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Redis Cluster Get Oci Cache User
	* `update` - (Defaults to 20 minutes), when updating the Redis Cluster Get Oci Cache User
	* `delete` - (Defaults to 20 minutes), when destroying the Redis Cluster Get Oci Cache User


## Import

Import is not supported for this resource.

