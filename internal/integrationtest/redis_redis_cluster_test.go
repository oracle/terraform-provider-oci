// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RedisRedisClusterRequiredOnlyResource = RedisRedisClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation)

	RedisRedisClusterResourceConfig = RedisRedisClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Update, RedisRedisClusterRepresentation)

	RedisRedisClusterSingularDataSourceRepresentation = map[string]interface{}{
		"redis_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_redis_cluster.test_redis_cluster.id}`},
	}

	RedisRedisClusterDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayNameNonSharded2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_redis_redis_cluster.test_redis_cluster.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisRedisClusterDataSourceFilterRepresentation}}

	RedisRedisClusterDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_redis_cluster.test_redis_cluster.id}`}},
	}

	RedisRedisClusterRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayNameNonSharded`, Update: `displayNameNonSharded2`},
		"node_count":         acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `5`},
		"node_memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"software_version":   acctest.Representation{RepType: acctest.Required, Create: `REDIS_7_0`},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"cluster_mode":       acctest.Representation{RepType: acctest.Optional, Create: `NONSHARDED`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":            acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRedisTagsChangesRepresentation},
	}

	RedisRedisShardedClusterRepresentation = map[string]interface{}{
		"compartment_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":       acctest.Representation{RepType: acctest.Required, Create: `displayNameSharded`, Update: `displayNameSharded2`},
		"node_count":         acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"node_memory_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"shard_count":        acctest.Representation{RepType: acctest.Required, Create: `3`, Update: `5`},
		"software_version":   acctest.Representation{RepType: acctest.Required, Create: `V7_0_5`},
		"subnet_id":          acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"cluster_mode":       acctest.Representation{RepType: acctest.Required, Create: `SHARDED`},
		"defined_tags":       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"nsg_ids":            acctest.Representation{RepType: acctest.Optional, Create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}, Update: []string{}},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRedisTagsChangesRepresentation},
	}

	ignoreRedisTagsChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`, `defined_tags`}, Update: []string{`defined_tags`, `freeform_tags`}},
	}

	RedisRedisClusterResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "redis_security_list", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSecurityListRepresentation, map[string]interface{}{
				"display_name": acctest.Representation{RepType: acctest.Required, Create: `redis-security-list`},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreSubnetRepresentation, map[string]interface{}{
				"security_list_ids": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.redis_security_list.id}`}}})) +
		DefinedTagsDependencies
)

// issue-routing-tag: redis/default
func TestRedisRedisClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisRedisClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_redis_redis_cluster.test_redis_cluster"
	shardedResourceName := "oci_redis_redis_cluster.test_sharded_redis_cluster"
	datasourceName := "data.oci_redis_redis_clusters.test_redis_clusters"
	singularDatasourceName := "data.oci_redis_redis_cluster.test_redis_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisRedisClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Create, RedisRedisClusterRepresentation), "redis", "redisCluster", t)

	acctest.ResourceTest(t, testAccCheckRedisRedisClusterDestroy, []resource.TestStep{
		// verify Create non Sharded cluster
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameNonSharded"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies,
		},

		//verify Create non sharded cluster with optionals
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Create, RedisRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_mode", "NONSHARDED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameNonSharded"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "node_collection.0.items.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RedisRedisClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_mode", "NONSHARDED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameNonSharded"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "node_collection.0.items.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "3"),
				resource.TestCheckResourceAttr(resourceName, "node_memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Update, RedisRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_mode", "NONSHARDED"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayNameNonSharded2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "node_collection.0.items.#", "5"),
				resource.TestCheckResourceAttr(resourceName, "node_count", "5"),
				resource.TestCheckResourceAttr(resourceName, "node_memory_in_gbs", "3"),
				resource.TestCheckResourceAttr(resourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "primary_fqdn"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(resourceName, "replicas_fqdn"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		//verify Create sharded cluster with optionals
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_sharded_redis_cluster", acctest.Optional, acctest.Create, RedisRedisShardedClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(shardedResourceName, "cluster_mode", "SHARDED"),
				resource.TestCheckResourceAttr(shardedResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(shardedResourceName, "display_name", "displayNameSharded"),
				resource.TestCheckResourceAttr(shardedResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "id"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_collection.0.items.#", "0"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_count", "2"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(shardedResourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(shardedResourceName, "shard_count", "3"),
				resource.TestCheckResourceAttr(shardedResourceName, "software_version", "V7_0_5"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, shardedResourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, shardedResourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_sharded_redis_cluster", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RedisRedisShardedClusterRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(shardedResourceName, "cluster_mode", "SHARDED"),
				resource.TestCheckResourceAttr(shardedResourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(shardedResourceName, "display_name", "displayNameSharded"),
				resource.TestCheckResourceAttr(shardedResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "id"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_collection.0.items.#", "0"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_count", "2"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_memory_in_gbs", "2"),
				resource.TestCheckResourceAttr(shardedResourceName, "nsg_ids.#", "1"),
				resource.TestCheckResourceAttr(shardedResourceName, "shard_count", "3"),
				resource.TestCheckResourceAttr(shardedResourceName, "software_version", "V7_0_5"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, shardedResourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_sharded_redis_cluster", acctest.Optional, acctest.Update, RedisRedisShardedClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(shardedResourceName, "cluster_mode", "SHARDED"),
				resource.TestCheckResourceAttr(shardedResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(shardedResourceName, "display_name", "displayNameSharded2"),
				resource.TestCheckResourceAttr(shardedResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "id"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_collection.0.items.#", "0"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_count", "3"),
				resource.TestCheckResourceAttr(shardedResourceName, "node_memory_in_gbs", "3"),
				resource.TestCheckResourceAttr(shardedResourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttr(shardedResourceName, "shard_count", "5"),
				resource.TestCheckResourceAttr(shardedResourceName, "software_version", "V7_0_5"),
				resource.TestCheckResourceAttrSet(shardedResourceName, "subnet_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, shardedResourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_redis_clusters", "test_redis_clusters", acctest.Optional, acctest.Create, RedisRedisClusterDataSourceRepresentation) +
				compartmentIdVariableStr + RedisRedisClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Optional, acctest.Update, RedisRedisClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayNameNonSharded2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "redis_cluster_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "redis_cluster_collection.0.items.#", "1"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_redis_cluster", "test_redis_cluster", acctest.Required, acctest.Create, RedisRedisClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RedisRedisClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "redis_cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_mode", "NONSHARDED"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayNameNonSharded2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_collection.0.items.#", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_count", "5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "node_memory_in_gbs", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "nsg_ids.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "primary_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "primary_fqdn"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replicas_endpoint_ip_address"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "replicas_fqdn"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},

		// verify resource import
		{
			Config:                  config + RedisRedisClusterRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRedisRedisClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).RedisClusterClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_redis_redis_cluster" {
			noResourceFound = false
			request := oci_redis.GetRedisClusterRequest{}

			tmp := rs.Primary.ID
			request.RedisClusterId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")

			response, err := client.GetRedisCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_redis.RedisClusterLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("RedisRedisCluster") {
		resource.AddTestSweepers("RedisRedisCluster", &resource.Sweeper{
			Name:         "RedisRedisCluster",
			Dependencies: acctest.DependencyGraph["redisCluster"],
			F:            sweepRedisRedisClusterResource,
		})
	}
}

func sweepRedisRedisClusterResource(compartment string) error {
	redisClusterClient := acctest.GetTestClients(&schema.ResourceData{}).RedisClusterClient()
	redisClusterIds, err := getRedisRedisClusterIds(compartment)
	if err != nil {
		return err
	}
	for _, redisClusterId := range redisClusterIds {
		if ok := acctest.SweeperDefaultResourceId[redisClusterId]; !ok {
			deleteRedisClusterRequest := oci_redis.DeleteRedisClusterRequest{}

			deleteRedisClusterRequest.RedisClusterId = &redisClusterId

			deleteRedisClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")
			_, error := redisClusterClient.DeleteRedisCluster(context.Background(), deleteRedisClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting RedisCluster %s %s, It is possible that the resource is already deleted. Please verify manually \n", redisClusterId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &redisClusterId, RedisRedisClusterSweepWaitCondition, time.Duration(3*time.Minute),
				RedisRedisClusterSweepResponseFetchOperation, "redis", true)
		}
	}
	return nil
}

func getRedisRedisClusterIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RedisClusterId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	redisClusterClient := acctest.GetTestClients(&schema.ResourceData{}).RedisClusterClient()

	listRedisClustersRequest := oci_redis.ListRedisClustersRequest{}
	listRedisClustersRequest.CompartmentId = &compartmentId
	listRedisClustersRequest.LifecycleState = oci_redis.RedisClusterLifecycleStateActive
	listRedisClustersResponse, err := redisClusterClient.ListRedisClusters(context.Background(), listRedisClustersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting RedisCluster list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, redisCluster := range listRedisClustersResponse.Items {
		id := *redisCluster.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RedisClusterId", id)
	}
	return resourceIds, nil
}

func RedisRedisClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if redisClusterResponse, ok := response.Response.(oci_redis.GetRedisClusterResponse); ok {
		return redisClusterResponse.LifecycleState != oci_redis.RedisClusterLifecycleStateDeleted
	}
	return false
}

func RedisRedisClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.RedisClusterClient().GetRedisCluster(context.Background(), oci_redis.GetRedisClusterRequest{
		RedisClusterId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
