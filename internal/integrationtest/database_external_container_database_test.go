// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExternalContainerDatabaseRequiredOnlyResource = DatabaseExternalContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation)

	DatabaseExternalContainerDatabaseResourceConfig = DatabaseExternalContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseRepresentation)

	DatabaseDatabaseExternalContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
	}

	DatabaseDatabaseExternalContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestExternalCdb`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `NOT_CONNECTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExternalContainerDatabaseDataSourceFilterRepresentation}}
	DatabaseExternalContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_external_container_database.test_external_container_database.id}`}},
	}

	DatabaseExternalContainerDatabaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `myTestExternalCdb`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseExternalContainerDatabaseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExternalContainerDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_external_container_database.test_external_container_database"
	datasourceName := "data.oci_database_external_container_databases.test_external_container_databases"
	singularDatasourceName := "data.oci_database_external_container_database.test_external_container_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExternalContainerDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseRepresentation), "database", "externalContainerDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExternalContainerDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + DatabaseExternalContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseExternalContainerDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_container_databases", "test_external_container_databases", acctest.Optional, acctest.Update, DatabaseDatabaseExternalContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NOT_CONNECTED"),

				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.database_management_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.ncharacter_set"),
				resource.TestCheckResourceAttr(datasourceName, "external_container_databases.0.stack_monitoring_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_container_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseDatabaseExternalContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseExternalContainerDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_container_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "myTestExternalCdb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "ncharacter_set"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stack_monitoring_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseExternalContainerDatabaseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExternalContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_container_database" {
			noResourceFound = false
			request := oci_database.GetExternalContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExternalContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalContainerDatabaseLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExternalContainerDatabase") {
		resource.AddTestSweepers("DatabaseExternalContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseExternalContainerDatabase",
			Dependencies: acctest.DependencyGraph["externalContainerDatabase"],
			F:            sweepDatabaseExternalContainerDatabaseResource,
		})
	}
}

func sweepDatabaseExternalContainerDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	externalContainerDatabaseIds, err := getDatabaseExternalContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalContainerDatabaseId := range externalContainerDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[externalContainerDatabaseId]; !ok {
			deleteExternalContainerDatabaseRequest := oci_database.DeleteExternalContainerDatabaseRequest{}

			deleteExternalContainerDatabaseRequest.ExternalContainerDatabaseId = &externalContainerDatabaseId

			deleteExternalContainerDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalContainerDatabase(context.Background(), deleteExternalContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalContainerDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalContainerDatabaseId, externalContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseExternalContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseExternalContainerDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExternalContainerDatabasesRequest := oci_database.ListExternalContainerDatabasesRequest{}
	listExternalContainerDatabasesRequest.CompartmentId = &compartmentId
	listExternalContainerDatabasesRequest.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateAvailable
	listExternalContainerDatabasesResponse, err := databaseClient.ListExternalContainerDatabases(context.Background(), listExternalContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalContainerDatabase := range listExternalContainerDatabasesResponse.Items {
		id := *externalContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func externalContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalContainerDatabaseResponse, ok := response.Response.(oci_database.GetExternalContainerDatabaseResponse); ok {
		return externalContainerDatabaseResponse.LifecycleState != oci_database.ExternalContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func DatabaseExternalContainerDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExternalContainerDatabase(context.Background(), oci_database.GetExternalContainerDatabaseRequest{
		ExternalContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
