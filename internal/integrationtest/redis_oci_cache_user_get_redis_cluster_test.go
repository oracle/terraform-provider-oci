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
	RedisOciCacheUserGetRedisClusterRequiredOnlyResource = RedisOciCacheUserGetRedisClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user_get_redis_cluster", "test_oci_cache_user_get_redis_cluster", acctest.Required, acctest.Create, RedisOciCacheUserGetRedisClusterRepresentation)

	RedisOciCacheUserGetRedisClusterDataSourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheUserGetRedisClusterDataSourceFilterRepresentation}}
	RedisOciCacheUserGetRedisClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_user_get_redis_cluster.test_oci_cache_user_get_redis_cluster.id}`}},
	}

	RedisOciCacheUserGetRedisClusterRepresentation = map[string]interface{}{
		"oci_cache_user_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_user.test_oci_cache_user.id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}

	RedisOciCacheUserGetRedisClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user", "test_oci_cache_user", acctest.Required, acctest.Create, RedisOciCacheUserRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisOciCacheUserGetRedisClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheUserGetRedisClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_oci_cache_user_get_redis_cluster.test_oci_cache_user_get_redis_cluster"
	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheUserGetRedisClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user_get_redis_cluster", "test_oci_cache_user_get_redis_cluster", acctest.Optional, acctest.Create, RedisOciCacheUserGetRedisClusterRepresentation), "redis", "ociCacheUserGetRedisCluster", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserGetRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user_get_redis_cluster", "test_oci_cache_user_get_redis_cluster", acctest.Required, acctest.Create, RedisOciCacheUserGetRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "oci_cache_user_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserGetRedisClusterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheUserGetRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_user_get_redis_cluster", "test_oci_cache_user_get_redis_cluster", acctest.Optional, acctest.Create, RedisOciCacheUserGetRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "oci_cache_clusters.#", "0"),

				resource.TestCheckResourceAttrSet(resourceName, "oci_cache_user_id"),

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
