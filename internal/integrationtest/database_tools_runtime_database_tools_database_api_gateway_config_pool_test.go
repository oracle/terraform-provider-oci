// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRequiredOnlyResource = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceConfig = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"pool_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `HR Pool`, Update: `HR Pool Updated`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceFilterRepresentation}}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `key`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`}},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation = map[string]interface{}{
		"database_tools_connection_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name":            acctest.Representation{RepType: acctest.Required, Create: `HR Pool`, Update: `HR Pool Updated`},
		"pool_route_value":        acctest.Representation{RepType: acctest.Required, Create: `hr`, Update: `hr`},
		"type":                    acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"initial_pool_size":       acctest.Representation{RepType: acctest.Optional, Create: `5`, Update: `5`},
		"max_pool_size":           acctest.Representation{RepType: acctest.Optional, Create: `100`, Update: `100`},
		"min_pool_size":           acctest.Representation{RepType: acctest.Optional, Create: `0`, Update: `0`},
		"rest_enabled_sql_status": acctest.Representation{RepType: acctest.Optional, Create: `ENABLED`, Update: `ENABLED`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	databaseToolsDatabaseApiGatewayConfigId := utils.GetEnvSettingWithDefault(
		"database_tools_cloud_api_gateway_config_id",
		utils.GetEnvSettingWithBlankDefault("database_tools_database_api_gateway_config_id"),
	)
	if databaseToolsDatabaseApiGatewayConfigId == "" {
		t.Skip("set database_tools_cloud_api_gateway_config_id or database_tools_database_api_gateway_config_id to run this test")
	}
	existingDatabaseToolsConnectionId := utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("database_tools_connection_id",
			utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))
	if existingDatabaseToolsConnectionId == "" {
		t.Skip("set existing_database_tools_connection_id, database_tools_connection_id, or database_tools_connection_ocid to run this test")
	}

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	databaseToolsDatabaseApiGatewayConfigIdVariableStr := terraformStringVariable("database_tools_cloud_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId)
	existingDatabaseToolsConnectionIdVariableStr := terraformStringVariable("existing_database_tools_connection_id", existingDatabaseToolsConnectionId)

	resourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool"
	datasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pools.test_database_tools_database_api_gateway_config_pools"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseToolsDatabaseApiGatewayConfigIdVariableStr+existingDatabaseToolsConnectionIdVariableStr+DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation), "databasetoolsruntime", "databaseToolsDatabaseApiGatewayConfigPool", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "HR Pool"),
				resource.TestCheckResourceAttr(resourceName, "pool_route_value", "hr"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "HR Pool"),
				resource.TestCheckResourceAttr(resourceName, "initial_pool_size", "5"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "max_pool_size", "100"),
				resource.TestCheckResourceAttr(resourceName, "min_pool_size", "0"),
				resource.TestCheckResourceAttr(resourceName, "pool_route_value", "hr"),
				resource.TestCheckResourceAttr(resourceName, "rest_enabled_sql_status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "HR Pool Updated"),
				resource.TestCheckResourceAttr(resourceName, "initial_pool_size", "5"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "max_pool_size", "100"),
				resource.TestCheckResourceAttr(resourceName, "min_pool_size", "0"),
				resource.TestCheckResourceAttr(resourceName, "pool_route_value", "hr"),
				resource.TestCheckResourceAttr(resourceName, "rest_enabled_sql_status", "ENABLED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pools", "test_database_tools_database_api_gateway_config_pools", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "HR Pool Updated"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool", "test_database_tools_database_api_gateway_config_pool", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "database_tools_database_api_gateway_config_id", databaseToolsDatabaseApiGatewayConfigId),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "pool_key", resourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "HR Pool Updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "initial_pool_size", "5"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_pool_size", "100"),
				resource.TestCheckResourceAttr(singularDatasourceName, "min_pool_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "pool_route_value", "hr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rest_enabled_sql_status", "ENABLED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}

			if value, ok := rs.Primary.Attributes["database_tools_database_api_gateway_config_id"]; ok {
				request.DatabaseToolsDatabaseApiGatewayConfigId = &value
			}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.PoolKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetDatabaseToolsDatabaseApiGatewayConfigPool(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
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
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPool") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPool", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPool",
			Dependencies: acctest.DependencyGraph["databaseToolsDatabaseApiGatewayConfigPool"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsDatabaseApiGatewayConfigPoolIds, err := getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsDatabaseApiGatewayConfigPoolId := range databaseToolsDatabaseApiGatewayConfigPoolIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsDatabaseApiGatewayConfigPoolId]; !ok {
			deleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest{}

			deleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteDatabaseToolsDatabaseApiGatewayConfigPool(context.Background(), deleteDatabaseToolsDatabaseApiGatewayConfigPoolRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsDatabaseApiGatewayConfigPool %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsDatabaseApiGatewayConfigPoolId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsDatabaseApiGatewayConfigPoolId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listDatabaseToolsDatabaseApiGatewayConfigPoolsRequest := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolsRequest{}

	databaseToolsDatabaseApiGatewayConfigIds, error := getConfigIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting databaseToolsDatabaseApiGatewayConfigId required for DatabaseToolsDatabaseApiGatewayConfigPool resource requests \n")
	}
	for _, databaseToolsDatabaseApiGatewayConfigId := range databaseToolsDatabaseApiGatewayConfigIds {
		listDatabaseToolsDatabaseApiGatewayConfigPoolsRequest.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId

		listDatabaseToolsDatabaseApiGatewayConfigPoolsResponse, err := databaseToolsRuntimeClient.ListDatabaseToolsDatabaseApiGatewayConfigPools(context.Background(), listDatabaseToolsDatabaseApiGatewayConfigPoolsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DatabaseToolsDatabaseApiGatewayConfigPool list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, databaseToolsDatabaseApiGatewayConfigPool := range listDatabaseToolsDatabaseApiGatewayConfigPoolsResponse.Items {
			if databaseToolsDatabaseApiGatewayConfigPool.GetKey() == nil {
				continue
			}
			id := databaseToolsDatabaseApiGatewayConfigId + "/" + *databaseToolsDatabaseApiGatewayConfigPool.GetKey()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsDatabaseApiGatewayConfigPoolId", id)
		}

	}
	return resourceIds, nil
}
