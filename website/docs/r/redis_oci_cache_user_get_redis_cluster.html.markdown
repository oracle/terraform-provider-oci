---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_oci_cache_user_get_redis_cluster"
sidebar_current: "docs-oci-resource-redis-oci_cache_user_get_redis_cluster"
description: |-
  Provides the Oci Cache User Get Redis Cluster resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_oci_cache_user_get_redis_cluster
This resource provides the Oci Cache User Get Redis Cluster resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/OciCacheUserGetRedisCluster

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Gets a list of associated redis cluster for an Oracle Cloud Infrastructure cache user.

## Example Usage

```hcl
resource "oci_redis_oci_cache_user_get_redis_cluster" "test_oci_cache_user_get_redis_cluster" {
	#Required
	oci_cache_user_id = oci_redis_oci_cache_user.test_oci_cache_user.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.oci_cache_user_get_redis_cluster_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `oci_cache_user_id` - (Required) A filter to return only resources, that match with the given Oracle Cloud Infrastructure cache user ID (OCID).


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `oci_cache_cluster_id` - OCID of the OciCacheCluster 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Oci Cache User Get Redis Cluster
	* `update` - (Defaults to 20 minutes), when updating the Oci Cache User Get Redis Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Oci Cache User Get Redis Cluster


## Import

Import is not supported for this resource.

