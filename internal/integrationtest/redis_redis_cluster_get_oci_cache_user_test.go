// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisRedisClusterGetOciCacheUserRequiredOnlyResource = RedisRedisClusterGetOciCacheUserResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_get_oci_cache_user", "test_redis_cluster_get_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterGetOciCacheUserRepresentation)

	RedisRedisClusterGetOciCacheUserDataSourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisRedisClusterGetOciCacheUserDataSourceFilterRepresentation}}
	RedisRedisClusterGetOciCacheUserDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_redis_cluster_get_oci_cache_user.test_redis_cluster_get_oci_cache_user.id}`}},
	}

	RedisRedisClusterGetOciCacheUserRepresentation = map[string]interface{}{
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	RedisRedisClusterGetOciCacheUserResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterGetOciCacheUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterGetOciCacheUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_redis_cluster_get_oci_cache_user.test_redis_cluster_get_oci_cache_user"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisRedisClusterGetOciCacheUserResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_get_oci_cache_user", "test_redis_cluster_get_oci_cache_user", acctest.Optional, acctest.Create, RedisRedisClusterGetOciCacheUserRepresentation), "redis", "redisClusterGetOciCacheUser", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterGetOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_get_oci_cache_user", "test_redis_cluster_get_oci_cache_user", acctest.Required, acctest.Create, RedisRedisClusterGetOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "redis_cluster_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterGetOciCacheUserResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterGetOciCacheUserResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster_get_oci_cache_user", "test_redis_cluster_get_oci_cache_user", acctest.Optional, acctest.Create, RedisRedisClusterGetOciCacheUserRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "oci_cache_users.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "redis_cluster_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
