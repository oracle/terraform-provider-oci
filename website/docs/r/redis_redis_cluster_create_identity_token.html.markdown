---
subcategory: "Redis"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_redis_redis_cluster_create_identity_token"
sidebar_current: "docs-oci-resource-redis-redis_cluster_create_identity_token"
description: |-
  Provides the Redis Cluster Create Identity Token resource in Oracle Cloud Infrastructure Redis service
---

# oci_redis_redis_cluster_create_identity_token
This resource provides the Redis Cluster Create Identity Token resource in Oracle Cloud Infrastructure Redis service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/redis/latest/RedisClusterCreateIdentityToken

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/redis

Generates an identity token to sign in with the specified redis user for the redis cluster

## Example Usage

```hcl
resource "oci_redis_redis_cluster_create_identity_token" "test_redis_cluster_create_identity_token" {
	#Required
	public_key = var.redis_cluster_create_identity_token_public_key
	redis_cluster_id = oci_redis_redis_cluster.test_redis_cluster.id
	redis_user = var.redis_cluster_create_identity_token_redis_user

	#Optional
	defined_tags = var.redis_cluster_create_identity_token_defined_tags
	freeform_tags = var.redis_cluster_create_identity_token_freeform_tags
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `public_key` - (Required) User public key pair
* `redis_cluster_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm#Oracle) of the cluster.
* `redis_user` - (Required) Redis User generating identity token.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `identity_token` - Generated Identity token
* `redis_user` - Redis user for the newly created identity token

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Redis Cluster Create Identity Token
	* `update` - (Defaults to 20 minutes), when updating the Redis Cluster Create Identity Token
	* `delete` - (Defaults to 20 minutes), when destroying the Redis Cluster Create Identity Token


## Import

Import is not supported for this resource.

