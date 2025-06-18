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
	RedisRedisClusterAttachOciCacheUserRepresentation = map[string]interface{}{
		"oci_cache_users":  acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_user.test_oci_cache_user.id}`}},
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster.id}`},
	}

	RedisRedisClusterAttachOciCacheUserResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterAttachOciCacheUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterAttachOciCacheUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_redis_cluster_attach_oci_cache_user.test_redis_cluster_attach_oci_cache_user"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisRedisClusterAttachOciCacheUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_attach_oci_cache_user", "test_redis_cluster_attach_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterAttachOciCacheUserRepresentation), "redis", "redisClusterAttachOciCacheUser", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterAttachOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_attach_oci_cache_user", "test_redis_cluster_attach_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterAttachOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "oci_cache_users.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "redis_cluster_id"),
			),
		},
	})
}
