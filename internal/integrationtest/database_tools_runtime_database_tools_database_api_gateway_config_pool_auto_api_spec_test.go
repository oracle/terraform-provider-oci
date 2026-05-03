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
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequiredOnlyResource = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceConfig = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecUpdateRepresentation)

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSingularDataSourceRepresentation = map[string]interface{}{
		"auto_api_spec_key":                             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec.test_database_tools_database_api_gateway_config_pool_auto_api_spec.key}`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"pool_key": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceRepresentation = map[string]interface{}{
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"pool_key":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `Employees Auto API`, Update: `displayName2`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecPoolRepresentation = map[string]interface{}{
		"database_tools_connection_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name":     acctest.Representation{RepType: acctest.Required, Create: `HR Pool`},
		"pool_route_value": acctest.Representation{RepType: acctest.Required, Create: `hr`},
		"type":             acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
	}

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation = map[string]interface{}{
		"database_object_name":                          acctest.Representation{RepType: acctest.Required, Create: `Some_Cool_Table_Name`, Update: `Some_Cool_Table_Name`},
		"database_object_type":                          acctest.Representation{RepType: acctest.Required, Create: `TABLE`, Update: `TABLE`},
		"database_tools_database_api_gateway_config_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_cloud_api_gateway_config_id}`},
		"display_name":                                  acctest.Representation{RepType: acctest.Required, Create: `Employees Auto API`, Update: `displayName2`},
		"pool_key":                                      acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool.key}`},
		"type":                                          acctest.Representation{RepType: acctest.Required, Create: `DEFAULT`},
		"alias":                                         acctest.Representation{RepType: acctest.Optional, Create: `emp`, Update: `alias2`},
		"description":                                   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
	}
	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecUpdateRepresentation = DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation

	DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_database_tools_runtime_database_tools_database_api_gateway_config_pool",
		"test_database_tools_database_api_gateway_config_pool",
		acctest.Required,
		acctest.Create,
		DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecPoolRepresentation,
	)
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource_basic")
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

	resourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec.test_database_tools_database_api_gateway_config_pool_auto_api_spec"
	poolResourceName := "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool.test_database_tools_database_api_gateway_config_pool"
	datasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_specs.test_database_tools_database_api_gateway_config_pool_auto_api_specs"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec.test_database_tools_database_api_gateway_config_pool_auto_api_spec"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+databaseToolsDatabaseApiGatewayConfigIdVariableStr+existingDatabaseToolsConnectionIdVariableStr+DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation), "databasetoolsruntime", "databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", t)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_object_name", "Some_Cool_Table_Name"),
				resource.TestCheckResourceAttr(resourceName, "database_object_type", "TABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Employees Auto API"),
				resource.TestCheckResourceAttrPair(resourceName, "pool_key", poolResourceName, "key"),
				resource.TestCheckResourceAttr(resourceName, "type", "DEFAULT"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Optional, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias", "emp"),
				resource.TestCheckResourceAttr(resourceName, "database_object_name", "Some_Cool_Table_Name"),
				resource.TestCheckResourceAttr(resourceName, "database_object_type", "TABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "Employees Auto API"),
				resource.TestCheckResourceAttrSet(resourceName, "key"),
				resource.TestCheckResourceAttrPair(resourceName, "pool_key", poolResourceName, "key"),
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
			Config: config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "alias", "alias2"),
				resource.TestCheckResourceAttr(resourceName, "database_object_name", "Some_Cool_Table_Name"),
				resource.TestCheckResourceAttr(resourceName, "database_object_type", "TABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_specs", "test_database_tools_database_api_gateway_config_pool_auto_api_specs", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Optional, acctest.Update, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecUpdateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrPair(datasourceName, "pool_key", poolResourceName, "key"),

				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_auto_api_spec_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_database_api_gateway_config_pool_auto_api_spec_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec", "test_database_tools_database_api_gateway_config_pool_auto_api_spec", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecSingularDataSourceRepresentation) +
				compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_api_spec_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_database_api_gateway_config_id"),
				resource.TestCheckResourceAttrPair(singularDatasourceName, "pool_key", poolResourceName, "key"),

				resource.TestCheckResourceAttr(singularDatasourceName, "alias", "alias2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_object_name", "Some_Cool_Table_Name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_object_type", "TABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "DEFAULT"),
			),
		},
		// verify resource import
		{
			Config:            config + compartmentIdVariableStr + databaseToolsDatabaseApiGatewayConfigIdVariableStr + existingDatabaseToolsConnectionIdVariableStr + DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeDatabaseApiGatewayConfigPoolAutoApiSpecCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseToolsRuntimeClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec" {
			noResourceFound = false
			request := oci_database_tools_runtime.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}

			if value, ok := rs.Primary.Attributes["key"]; ok {
				request.AutoApiSpecKey = &value
			}

			if value, ok := rs.Primary.Attributes["database_tools_database_api_gateway_config_id"]; ok {
				request.DatabaseToolsDatabaseApiGatewayConfigId = &value
			}

			if value, ok := rs.Primary.Attributes["pool_key"]; ok {
				request.PoolKey = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")

			_, err := client.GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec") {
		resource.AddTestSweepers("DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec", &resource.Sweeper{
			Name:         "DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec",
			Dependencies: acctest.DependencyGraph["databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec"],
			F:            sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource,
		})
	}
}

func sweepDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource(compartment string) error {
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()
	databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecIds, err := getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId := range databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecIds {
		if ok := acctest.SweeperDefaultResourceId[databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId]; !ok {
			deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest := oci_database_tools_runtime.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest{}

			deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_tools_runtime")
			_, error := databaseToolsRuntimeClient.DeleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec(context.Background(), deleteDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseToolsRuntimeClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseToolsRuntimeClient()

	listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest := oci_database_tools_runtime.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest{}

	databaseToolsDatabaseApiGatewayConfigIds, error := getConfigIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting databaseToolsDatabaseApiGatewayConfigId required for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec resource requests \n")
	}
	for _, databaseToolsDatabaseApiGatewayConfigId := range databaseToolsDatabaseApiGatewayConfigIds {
		listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest.DatabaseToolsDatabaseApiGatewayConfigId = &databaseToolsDatabaseApiGatewayConfigId

		poolKeys, error := getDatabaseToolsRuntimePoolKeys(compartment)
		if error != nil {
			return resourceIds, fmt.Errorf("Error getting poolKey required for DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec resource requests \n")
		}
		for _, poolKey := range poolKeys {
			listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest.PoolKey = &poolKey

			listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse, err := databaseToolsRuntimeClient.ListDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecs(context.Background(), listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsRequest)

			if err != nil {
				return resourceIds, fmt.Errorf("Error getting DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec list for compartment id : %s , %s \n", compartmentId, err)
			}
			for _, databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec := range listDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsResponse.Items {
				if databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.GetKey() == nil {
					continue
				}
				id := databaseToolsDatabaseApiGatewayConfigId + "/" + poolKey + "/" + *databaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.GetKey()
				resourceIds = append(resourceIds, id)
				acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId", id)
			}

		}
	}
	return resourceIds, nil
}
