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
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalMySqlDatabaseResourceConfig = DatabaseManagementExternalMySqlDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseRepresentation)

	DatabaseManagementExternalMySqlDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_my_sql_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_my_sql_database.test_external_my_sql_database.id}`},
	}

	DatabaseManagementExternalMySqlDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`}}

	DatabaseManagementExternalMySqlDatabaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"db_name":        acctest.Representation{RepType: acctest.Required, Create: `ExternalDB45`, Update: `dbName44`},
	}

	DatabaseManagementExternalMySqlDatabaseResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalMySqlDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalMySqlDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_management_external_my_sql_database.test_external_my_sql_database"
	datasourceName := "data.oci_database_management_external_my_sql_databases.test_external_my_sql_databases"
	singularDatasourceName := "data.oci_database_management_external_my_sql_database.test_external_my_sql_database"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementExternalMySqlDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseRepresentation), "databasemanagement", "externalMySqlDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalMySqlDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalMySqlDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_name", "ExternalDB45"),

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
			Config: config + compartmentIdVariableStr + DatabaseManagementExternalMySqlDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "db_name", "dbName44"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_my_sql_databases", "test_external_my_sql_databases", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalMySqlDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Optional, acctest.Update, DatabaseManagementExternalMySqlDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttr(datasourceName, "external_my_sql_database_collection.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_my_sql_database", "test_external_my_sql_database", acctest.Required, acctest.Create, DatabaseManagementExternalMySqlDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseManagementExternalMySqlDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_my_sql_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_name", "dbName44"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_database_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseManagementExternalMySqlDatabaseResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalMySqlDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_my_sql_database" {
			noResourceFound = false
			request := oci_database_management.GetExternalMySqlDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalMySqlDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			_, err := client.GetExternalMySqlDatabase(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalMySqlDatabase") {
		resource.AddTestSweepers("DatabaseManagementExternalMySqlDatabase", &resource.Sweeper{
			Name:         "DatabaseManagementExternalMySqlDatabase",
			Dependencies: acctest.DependencyGraph["externalMySqlDatabase"],
			F:            sweepDatabaseManagementExternalMySqlDatabaseResource,
		})
	}
}

func sweepDatabaseManagementExternalMySqlDatabaseResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalMySqlDatabaseIds, err := getDatabaseManagementExternalMySqlDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalMySqlDatabaseId := range externalMySqlDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[externalMySqlDatabaseId]; !ok {
			deleteExternalMySqlDatabaseRequest := oci_database_management.DeleteExternalMySqlDatabaseRequest{}

			deleteExternalMySqlDatabaseRequest.ExternalMySqlDatabaseId = &externalMySqlDatabaseId

			deleteExternalMySqlDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalMySqlDatabase(context.Background(), deleteExternalMySqlDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalMySqlDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalMySqlDatabaseId, error)
				continue
			}
		}
	}
	return nil
}

func getDatabaseManagementExternalMySqlDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalMySqlDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalMySqlDatabasesRequest := oci_database_management.ListExternalMySqlDatabasesRequest{}
	listExternalMySqlDatabasesRequest.CompartmentId = &compartmentId
	listExternalMySqlDatabasesResponse, err := dbManagementClient.ListExternalMySqlDatabases(context.Background(), listExternalMySqlDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalMySqlDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalMySqlDatabase := range listExternalMySqlDatabasesResponse.Items {
		id := *externalMySqlDatabase.ExternalDatabaseId
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalMySqlDatabaseId", id)
	}
	return resourceIds, nil
}
