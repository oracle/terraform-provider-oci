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
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DbmulticloudOracleDbAwsKeyPoolRequiredOnlyResource = DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeyPoolRepresentation)

	DbmulticloudOracleDbAwsKeyPoolResourceConfig = DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyPoolRepresentation)

	DbmulticloudOracleDbAwsKeyPoolSingularDataSourceRepresentation = map[string]interface{}{
		"oracle_db_aws_key_pool_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool.id}`},
	}

	DbmulticloudOracleDbAwsKeyPoolDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `TestDbAwsVault`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DbmulticloudOracleDbAwsKeyPoolDataSourceFilterRepresentation}}
	DbmulticloudOracleDbAwsKeyPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool.id}`}},
	}

	DbmulticloudOracleDbAwsKeyPoolRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `TestDbAwsVault`, Update: `displayName2`},
		//"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "value"})}`, Update: `${tomap({"${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}" = "updatedValue"})}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_multi_region_key_pool": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_reuse_allowed":         acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	DbmulticloudOracleDbAwsKeyPoolResourceDependencies = ""
)

// issue-routing-tag: dbmulticloud/default
func TestDbmulticloudOracleDbAwsKeyPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDbmulticloudOracleDbAwsKeyPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool"
	datasourceName := "data.oci_dbmulticloud_oracle_db_aws_key_pools.test_oracle_db_aws_key_pools"
	singularDatasourceName := "data.oci_dbmulticloud_oracle_db_aws_key_pool.test_oracle_db_aws_key_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbmulticloudOracleDbAwsKeyPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsKeyPoolRepresentation), "dbmulticloud", "oracleDbAwsKeyPool", t)

	acctest.ResourceTest(t, testAccCheckDbmulticloudOracleDbAwsKeyPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeyPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAwsVault"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Create, DbmulticloudOracleDbAwsKeyPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAwsVault"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_multi_region_key_pool", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_reuse_allowed", "false"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DbmulticloudOracleDbAwsKeyPoolRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "TestDbAwsVault"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_multi_region_key_pool", "false"),
				resource.TestCheckResourceAttr(resourceName, "is_reuse_allowed", "false"),

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
			Config: config + compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_multi_region_key_pool", "true"),
				resource.TestCheckResourceAttr(resourceName, "is_reuse_allowed", "true"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pools", "test_oracle_db_aws_key_pools", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyPoolDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Optional, acctest.Update, DbmulticloudOracleDbAwsKeyPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "oracle_db_aws_key_pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "oracle_db_aws_key_pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dbmulticloud_oracle_db_aws_key_pool", "test_oracle_db_aws_key_pool", acctest.Required, acctest.Create, DbmulticloudOracleDbAwsKeyPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbmulticloudOracleDbAwsKeyPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_multi_region_key_pool", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_reuse_allowed", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "replication_metadata.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + DbmulticloudOracleDbAwsKeyPoolRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDbmulticloudOracleDbAwsKeyPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbMulticloudAwsProviderClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_dbmulticloud_oracle_db_aws_key_pool" {
			noResourceFound = false
			request := oci_dbmulticloud.GetOracleDbAwsKeyPoolRequest{}

			tmp := rs.Primary.ID
			request.OracleDbAwsKeyPoolId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")

			response, err := client.GetOracleDbAwsKeyPool(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DbmulticloudOracleDbAwsKeyPool") {
		resource.AddTestSweepers("DbmulticloudOracleDbAwsKeyPool", &resource.Sweeper{
			Name:         "DbmulticloudOracleDbAwsKeyPool",
			Dependencies: acctest.DependencyGraph["oracleDbAwsKeyPool"],
			F:            sweepDbmulticloudOracleDbAwsKeyPoolResource,
		})
	}
}

func sweepDbmulticloudOracleDbAwsKeyPoolResource(compartment string) error {
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()
	oracleDbAwsKeyPoolIds, err := getDbmulticloudOracleDbAwsKeyPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, oracleDbAwsKeyPoolId := range oracleDbAwsKeyPoolIds {
		if ok := acctest.SweeperDefaultResourceId[oracleDbAwsKeyPoolId]; !ok {
			deleteOracleDbAwsKeyPoolRequest := oci_dbmulticloud.DeleteOracleDbAwsKeyPoolRequest{}

			deleteOracleDbAwsKeyPoolRequest.OracleDbAwsKeyPoolId = &oracleDbAwsKeyPoolId

			deleteOracleDbAwsKeyPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "dbmulticloud")
			_, error := dbMulticloudAwsProviderClient.DeleteOracleDbAwsKeyPool(context.Background(), deleteOracleDbAwsKeyPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting OracleDbAwsKeyPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", oracleDbAwsKeyPoolId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &oracleDbAwsKeyPoolId, DbmulticloudOracleDbAwsKeyPoolSweepWaitCondition, time.Duration(3*time.Minute),
				DbmulticloudOracleDbAwsKeyPoolSweepResponseFetchOperation, "dbmulticloud", true)
		}
	}
	return nil
}

func getDbmulticloudOracleDbAwsKeyPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OracleDbAwsKeyPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbMulticloudAwsProviderClient := acctest.GetTestClients(&schema.ResourceData{}).DbMulticloudAwsProviderClient()

	listOracleDbAwsKeyPoolsRequest := oci_dbmulticloud.ListOracleDbAwsKeyPoolsRequest{}
	listOracleDbAwsKeyPoolsRequest.CompartmentId = &compartmentId
	listOracleDbAwsKeyPoolsRequest.LifecycleState = oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateActive
	listOracleDbAwsKeyPoolsResponse, err := dbMulticloudAwsProviderClient.ListOracleDbAwsKeyPools(context.Background(), listOracleDbAwsKeyPoolsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OracleDbAwsKeyPool list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, oracleDbAwsKeyPool := range listOracleDbAwsKeyPoolsResponse.Items {
		id := *oracleDbAwsKeyPool.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OracleDbAwsKeyPoolId", id)
	}
	return resourceIds, nil
}

func DbmulticloudOracleDbAwsKeyPoolSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if oracleDbAwsKeyPoolResponse, ok := response.Response.(oci_dbmulticloud.GetOracleDbAwsKeyPoolResponse); ok {
		return oracleDbAwsKeyPoolResponse.LifecycleState != oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateDeleted
	}
	return false
}

func DbmulticloudOracleDbAwsKeyPoolSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbMulticloudAwsProviderClient().GetOracleDbAwsKeyPool(context.Background(), oci_dbmulticloud.GetOracleDbAwsKeyPoolRequest{
		OracleDbAwsKeyPoolId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
