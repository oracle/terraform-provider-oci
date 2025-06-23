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
	RedisOciCacheConfigSetlistAssociatedOciCacheClusterDataSourceRepresentation = map[string]interface{}{
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheConfigSetlistAssociatedOciCacheClusterDataSourceFilterRepresentation}}
	RedisOciCacheConfigSetlistAssociatedOciCacheClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster.test_oci_cache_config_setlist_associated_oci_cache_cluster.id}`}},
	}

	RedisOciCacheConfigSetlistAssociatedOciCacheClusterRepresentation = map[string]interface{}{
		"oci_cache_config_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_config_set.test_oci_cache_config_set.id}`},
	}

	RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Required, acctest.Create, RedisOciCacheConfigSetRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisOciCacheConfigSetlistAssociatedOciCacheClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheConfigSetlistAssociatedOciCacheClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster.test_oci_cache_config_setlist_associated_oci_cache_cluster"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster", "test_oci_cache_config_setlist_associated_oci_cache_cluster", acctest.Required, acctest.Create, RedisOciCacheConfigSetlistAssociatedOciCacheClusterRepresentation), "redis", "ociCacheConfigSetlistAssociatedOciCacheCluster", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheConfigSetlistAssociatedOciCacheClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster", "test_oci_cache_config_setlist_associated_oci_cache_cluster", acctest.Required, acctest.Create, RedisOciCacheConfigSetlistAssociatedOciCacheClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "oci_cache_config_set_id"),
			),
		},
	})
}
