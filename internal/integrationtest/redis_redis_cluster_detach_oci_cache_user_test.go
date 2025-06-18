// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisRedisClusterDetachOciCacheUserRepresentation = map[string]interface{}{
		"oci_cache_users":  acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_user.test_oci_cache_user.id}`}},
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster.id}`},
	}

	RedisRedisClusterDetachOciCacheUserResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterDetachOciCacheUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterDetachOciCacheUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	attachResourceName := "oci_redis_redis_cluster_attach_oci_cache_user.test_redis_cluster_attach_oci_cache_user"
	detachResourceName := "oci_redis_redis_cluster_detach_oci_cache_user.test_redis_cluster_detach_oci_cache_user"

	// Save TF content for attach step first
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisRedisClusterDetachOciCacheUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_attach_oci_cache_user", "test_redis_cluster_attach_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterDetachOciCacheUserRepresentation), "redis", "redisClusterAttachOciCacheUser", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// **Step 1: Attach and verify attach**
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterDetachOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_attach_oci_cache_user", "test_redis_cluster_attach_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterDetachOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(attachResourceName, "oci_cache_users.#", "1"),
				resource.TestCheckResourceAttrSet(attachResourceName, "redis_cluster_id"),
			),
		},
		// **Step 2: Detach and verify detach**
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterDetachOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_detach_oci_cache_user", "test_redis_cluster_detach_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterDetachOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(detachResourceName, "oci_cache_users.#", "1"),
				resource.TestCheckResourceAttrSet(detachResourceName, "redis_cluster_id"),
			),
		},
	})
}
