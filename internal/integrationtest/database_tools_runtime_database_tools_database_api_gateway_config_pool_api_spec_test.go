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
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceConfig = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSingularDataSourceRepresentation = map[string]interface{}{
		"api_spec_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec.test_database_tools_database_api_gateway_config_pool_api_spec.key}`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"pool_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"pool_key":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `Database API Gateway Config Pool Api Spec Canary Test`, Update: `Database API Gateway Config Pool Api Spec Canary Test Updated`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceFilterRepresentation}}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec.test_database_tools_database_api_gateway_config_pool_api_spec.id}`}},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecPoolRepresentation = map[string]interface{}{
		"database_tools_connection_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `HR Pool`},
		"pool_route_value": acctest.Representation{RepType: acctest.Required, Create: `hr`},
		"type":             acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation = map[string]interface{}{
		"content": acctest.Representation{
			RepType: acctest.Required,
			Create:  `{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"unit test spec\",\"version\":\"1.0\"},\"paths\":{\"/test\":{\"get\":{\"responses\":{},\"x-dbtools-operation\":{\"sourceType\":\"query\",\"source\":\"select 1 from dual\"}}}}}`,
			Update:  `{\"openapi\":\"3.0.0\",\"info\":{\"title\":\"unit test spec updated\",\"version\":\"1.1\"},\"paths\":{\"/test\":{\"get\":{\"responses\":{},\"x-dbtools-operation\":{\"sourceType\":\"query\",\"source\":\"select sysdate from dual\"}}}}}`,
		},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name": acctest.Representation{RepType: acctest.Required, Create: `Database API Gateway Config Pool Api Spec Canary Test`, Update: `Database API Gateway Config Pool Api Spec Canary Test Updated`},
		"pool_key":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
		"type":         acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_database_tools_runtime_database_tools_database_api_gateway_config_pool",
		"test_database_tools_database_api_gateway_config_pool",
		acctest.Required,
		acctest.Create,
		DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecPoolRepresentation,
	)
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource_basic")
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

	resourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec.test_database_tools_database_api_gateway_config_pool_api_spec"
	poolResourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool"
	datasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs.test_database_tools_database_api_gateway_config_pool_api_specs"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec.test_database_tools_database_api_gateway_config_pool_api_spec"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseToolsDatabaseApiGatewayConfigIdVariableStr+existingDatabaseToolsConnectionIdVariableStr+DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation), "databasetoolsruntime", "databaseToolsDatabaseApiGatewayConfigPoolApiSpec", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", `{"openapi":"3.0.0","info":{"title":"unit test spec","version":"1.0"},"paths":{"/test":{"get":{"responses":{},"x-dbtools-operation":{"sourceType":"query","source":"select 1 from dual"}}}}}`),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Database API Gateway Config Pool Api Spec Canary Test"),
				resource.TestCheckResourceAttrPair(resourceName, "pool_key", poolResourceName, "key"),
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
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "content", `{"openapi":"3.0.0","info":{"title":"unit test spec updated","version":"1.1"},"paths":{"/test":{"get":{"responses":{},"x-dbtools-operation":{"sourceType":"query","source":"select sysdate from dual"}}}}}`),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Database API Gateway Config Pool Api Spec Canary Test Updated"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrPair(resourceName, "pool_key", poolResourceName, "key"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs", "test_database_tools_database_api_gateway_config_pool_api_specs", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "Database API Gateway Config Pool Api Spec Canary Test Updated"),
				resource.TestCheckResourceAttrPair(datasourceName, "pool_key", poolResourceName, "key"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_api_spec_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_api_spec_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec", "test_database_tools_database_api_gateway_config_pool_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrPair(singularDatasourceName, "api_spec_key", resourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "pool_key", poolResourceName, "key"),

				resource.TestCheckResourceAttr(singularDatasourceName, "content", `{"openapi":"3.0.0","info":{"title":"unit test spec updated","version":"1.1"},"paths":{"/test":{"get":{"responses":{},"x-dbtools-operation":{"sourceType":"query","source":"select sysdate from dual"}}}}}`),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "Database API Gateway Config Pool Api Spec Canary Test Updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolApiSpecCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.ApiSpecKey = &value
			}

			if value, ok := rs.Primary.Attributes["database_tools_database_api_gateway_config_id"]; ok {
				request.DatabaseToolsDatabaseApiGatewayConfigId = &value
			}

			if value, ok := rs.Primary.Attributes["pool_key"]; ok {
				request.PoolKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec",
			Dependencies: acctest.DependencyGraph["databaseToolsDatabaseApiGatewayConfigPoolApiSpec"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsDatabaseApiGatewayConfigPoolApiSpecIds, err := getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsDatabaseApiGatewayConfigPoolApiSpecId := range databaseToolsDatabaseApiGatewayConfigPoolApiSpecIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsDatabaseApiGatewayConfigPoolApiSpecId]; !ok {
			deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest{}

			deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpec(context.Background(), deleteDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsDatabaseApiGatewayConfigPoolApiSpecId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest{}

	databaseToolsDatabaseApiGatewayConfigIds, error := getConfigIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting databaseToolsDatabaseApiGatewayConfigId required for DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec resource requests \n")
	}
	for _, databaseToolsDatabaseApiGatewayConfigId := range databaseToolsDatabaseApiGatewayConfigIds {
		listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId

		poolKeys, error := getDatabaseToolsRuntimePoolKeys(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting poolKey required for DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec resource requests \n")
		}
		for _, poolKey := range poolKeys {
			listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest.PoolKey = &poolKey

			listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse, err := databaseToolsRuntimeClient.ListDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecs(context.Background(), listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, databaseToolsDatabaseApiGatewayConfigPoolApiSpec := range listDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsResponse.Items {
				if databaseToolsDatabaseApiGatewayConfigPoolApiSpec.GetKey() == nil {
					continue
				}
				id := databaseToolsDatabaseApiGatewayConfigId + "/" + poolKey + "/" + *databaseToolsDatabaseApiGatewayConfigPoolApiSpec.GetKey()
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsDatabaseApiGatewayConfigPoolApiSpecId", id)
			}

		}
	}
	return resourceIds, nil
}

func getDatabaseToolsRuntimePoolKeys(compartment string) ([]string, error) {
	return []string{"poolKey"}, nil
}
