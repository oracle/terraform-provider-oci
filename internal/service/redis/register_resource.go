// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_redis_oci_cache_config_set", RedisOciCacheConfigSetResource())
	tfresource.RegisterResource("oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster", RedisOciCacheConfigSetlistAssociatedOciCacheClusterResource())
	tfresource.RegisterResource("oci_redis_oci_cache_user", RedisOciCacheUserResource())
	tfresource.RegisterResource("oci_redis_oci_cache_user_get_redis_cluster", RedisOciCacheUserGetRedisClusterResource())
	tfresource.RegisterResource("oci_redis_redis_cluster", RedisRedisClusterResource())
	tfresource.RegisterResource("oci_redis_redis_cluster_attach_oci_cache_user", RedisRedisClusterAttachOciCacheUserResource())
	tfresource.RegisterResource("oci_redis_redis_cluster_detach_oci_cache_user", RedisRedisClusterDetachOciCacheUserResource())
	tfresource.RegisterResource("oci_redis_redis_cluster_get_oci_cache_user", RedisRedisClusterGetOciCacheUserResource())
	tfresource.RegisterResource("oci_redis_redis_cluster_create_identity_token", RedisRedisClusterCreateIdentityTokenResource())
}
