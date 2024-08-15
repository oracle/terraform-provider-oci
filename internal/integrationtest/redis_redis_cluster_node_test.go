// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisRedisClusterNodeDataSourceRepresentation = map[string]interface{}{
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_sharded_redis_cluster.id}`},
	}

	RedisRedisClusterNodeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "redis_security_list", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
				"display_name": acctest.Representation{RepType: acctest.Required, Create: `redis-security-list`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.redis_security_list.id}`}}})) +
		DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_sharded_redis_cluster", acctest.Required, acctest.Create, RedisRedisShardedClusterRepresentation)
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterNodeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterNodeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_redis_redis_cluster_nodes.test_redis_cluster_nodes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_redis_cluster_nodes", "test_redis_cluster_nodes", acctest.Required, acctest.Create, RedisRedisClusterNodeDataSourceRepresentation) +
				compartmentIdVariableStr + RedisRedisClusterNodeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "redis_cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "redis_node_collection.0.items.#", "6"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.0.items.0.private_endpoint_fqdn"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.0.items.0.shard_number"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.0.items.0.private_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(datasourceName, "redis_node_collection.0.items.0.redis_cluster_id"),
			),
		},
	})
}
