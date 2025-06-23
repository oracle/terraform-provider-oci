// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	RedisOciCacheConfigSetRequiredOnlyResource = RedisOciCacheConfigSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Required, acctest.Create, RedisOciCacheConfigSetRepresentation)

	RedisOciCacheConfigSetResourceConfig = RedisOciCacheConfigSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Update, RedisOciCacheConfigSetRepresentation)

	RedisOciCacheConfigSetSingularDataSourceRepresentation = map[string]interface{}{
		"oci_cache_config_set_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_redis_oci_cache_config_set.test_oci_cache_config_set.id}`},
	}

	RedisOciCacheConfigSetDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_redis_oci_cache_config_set.test_oci_cache_config_set.id}`},
		"software_version": acctest.Representation{RepType: acctest.Optional, Create: `REDIS_7_0`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheConfigSetDataSourceFilterRepresentation}}
	RedisOciCacheConfigSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_redis_oci_cache_config_set.test_oci_cache_config_set.id}`}},
	}

	RedisOciCacheConfigSetRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheConfigSetConfigurationDetailsRepresentation},
		"display_name":          acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"software_version":      acctest.Representation{RepType: acctest.Required, Create: `REDIS_7_0`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":           acctest.Representation{RepType: acctest.Required, Create: `description`, Update: `description2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreRedisTagsChangesRepresentation},
	}
	RedisOciCacheConfigSetConfigurationDetailsRepresentation = map[string]interface{}{
		"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: RedisOciCacheConfigSetConfigurationDetailsItemsRepresentation},
	}
	RedisOciCacheConfigSetConfigurationDetailsItemsRepresentation = map[string]interface{}{
		"config_key":   acctest.Representation{RepType: acctest.Required, Create: `notify-keyspace-events`},
		"config_value": acctest.Representation{RepType: acctest.Required, Create: `KEA`},
	}

	RedisOciCacheConfigSetResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: redis/default
func TestRedisOciCacheConfigSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestRedisOciCacheConfigSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_redis_oci_cache_config_set.test_oci_cache_config_set"
	datasourceName := "data.oci_redis_oci_cache_config_sets.test_oci_cache_config_sets"
	singularDatasourceName := "data.oci_redis_oci_cache_config_set.test_oci_cache_config_set"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RedisOciCacheConfigSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Create, RedisOciCacheConfigSetRepresentation), "redis", "ociCacheConfigSet", t)

	acctest.ResourceTest(t, testAccCheckRedisOciCacheConfigSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheConfigSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Required, acctest.Create, RedisOciCacheConfigSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_key", "notify-keyspace-events"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_value", "KEA"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheConfigSetResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + RedisOciCacheConfigSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Create, RedisOciCacheConfigSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_key", "notify-keyspace-events"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_value", "KEA"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + RedisOciCacheConfigSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(RedisOciCacheConfigSetRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_key", "notify-keyspace-events"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_value", "KEA"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

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
			Config: config + compartmentIdVariableStr + RedisOciCacheConfigSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Update, RedisOciCacheConfigSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_key", "notify-keyspace-events"),
				resource.TestCheckResourceAttr(resourceName, "configuration_details.0.items.0.config_value", "KEA"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_config_sets", "test_oci_cache_config_sets", acctest.Optional, acctest.Update, RedisOciCacheConfigSetDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheConfigSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Optional, acctest.Update, RedisOciCacheConfigSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oci_cache_config_set_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oci_cache_config_set_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_redis_oci_cache_config_set", "test_oci_cache_config_set", acctest.Required, acctest.Create, RedisOciCacheConfigSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RedisOciCacheConfigSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "oci_cache_config_set_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.items.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.items.0.config_key", "notify-keyspace-events"),
				resource.TestCheckResourceAttr(singularDatasourceName, "configuration_details.0.items.0.config_value", "KEA"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_config_set_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "software_version", "REDIS_7_0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + RedisOciCacheConfigSetRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckRedisOciCacheConfigSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OciCacheConfigSetClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_redis_oci_cache_config_set" {
			noResourceFound = false
			request := oci_redis.GetOciCacheConfigSetRequest{}

			tmp := rs.Primary.ID
			request.OciCacheConfigSetId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")

			response, err := client.GetOciCacheConfigSet(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_redis.OciCacheConfigSetLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("RedisOciCacheConfigSet") {
		resource.AddTestSweepers("RedisOciCacheConfigSet", &resource.Sweeper{
			Name:         "RedisOciCacheConfigSet",
			Dependencies: acctest.DependencyGraph["ociCacheConfigSet"],
			F:            sweepRedisOciCacheConfigSetResource,
		})
	}
}

func sweepRedisOciCacheConfigSetResource(compartment string) error {
	ociCacheConfigSetClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheConfigSetClient()
	ociCacheConfigSetIds, err := getRedisOciCacheConfigSetIds(compartment)
	if err != nil {
		return err
	}
	for _, ociCacheConfigSetId := range ociCacheConfigSetIds {
		if ok := acctest.SweeperDefaultResourceId[ociCacheConfigSetId]; !ok {
			deleteOciCacheConfigSetRequest := oci_redis.DeleteOciCacheConfigSetRequest{}

			deleteOciCacheConfigSetRequest.OciCacheConfigSetId = &ociCacheConfigSetId

			deleteOciCacheConfigSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "redis")
			_, error := ociCacheConfigSetClient.DeleteOciCacheConfigSet(context.Background(), deleteOciCacheConfigSetRequest)
			if error != nil {
				fmt.Printf("Error deleting OciCacheConfigSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", ociCacheConfigSetId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &ociCacheConfigSetId, RedisOciCacheConfigSetSweepWaitCondition, time.Duration(3*time.Minute),
				RedisOciCacheConfigSetSweepResponseFetchOperation, "redis", true)
		}
	}
	return nil
}

func getRedisOciCacheConfigSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OciCacheConfigSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	ociCacheConfigSetClient := acctest.GetTestClients(&schema.ResourceData{}).OciCacheConfigSetClient()

	listOciCacheConfigSetsRequest := oci_redis.ListOciCacheConfigSetsRequest{}
	listOciCacheConfigSetsRequest.CompartmentId = &compartmentId
	listOciCacheConfigSetsRequest.LifecycleState = oci_redis.OciCacheConfigSetLifecycleStateActive
	listOciCacheConfigSetsResponse, err := ociCacheConfigSetClient.ListOciCacheConfigSets(context.Background(), listOciCacheConfigSetsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OciCacheConfigSet list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, ociCacheConfigSet := range listOciCacheConfigSetsResponse.Items {
		id := *ociCacheConfigSet.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OciCacheConfigSetId", id)
		acctest.SweeperDefaultResourceId[*ociCacheConfigSet.DefaultConfigSetId] = true

	}
	return resourceIds, nil
}

func RedisOciCacheConfigSetSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if ociCacheConfigSetResponse, ok := response.Response.(oci_redis.GetOciCacheConfigSetResponse); ok {
		return ociCacheConfigSetResponse.LifecycleState != oci_redis.OciCacheConfigSetLifecycleStateDeleted
	}
	return false
}

func RedisOciCacheConfigSetSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OciCacheConfigSetClient().GetOciCacheConfigSet(context.Background(), oci_redis.GetOciCacheConfigSetRequest{
		OciCacheConfigSetId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
