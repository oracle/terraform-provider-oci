// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	ExternalNonContainerDatabaseRequiredOnlyResource = ExternalNonContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, externalNonContainerDatabaseRepresentation)

	ExternalNonContainerDatabaseResourceConfig = ExternalNonContainerDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Update, externalNonContainerDatabaseRepresentation)

	externalNonContainerDatabaseSingularDataSourceRepresentation = map[string]interface{}{
		"external_non_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_non_container_database.test_external_non_container_database.id}`},
	}

	externalNonContainerDatabaseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `myTestExternalNonCdb`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `NOT_CONNECTED`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: externalNonContainerDatabaseDataSourceFilterRepresentation}}
	externalNonContainerDatabaseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_external_non_container_database.test_external_non_container_database.id}`}},
	}

	externalNonContainerDatabaseRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `myTestExternalNonCdb`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	ExternalNonContainerDatabaseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: database/default
func TestDatabaseExternalNonContainerDatabaseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalNonContainerDatabaseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_database_external_non_container_database.test_external_non_container_database"
	datasourceName := "data.oci_database_external_non_container_databases.test_external_non_container_databases"
	singularDatasourceName := "data.oci_database_external_non_container_database.test_external_non_container_database"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ExternalNonContainerDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Create, externalNonContainerDatabaseRepresentation), "database", "externalNonContainerDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseExternalNonContainerDatabaseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, externalNonContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Create, externalNonContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ExternalNonContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(externalNonContainerDatabaseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Update, externalNonContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "myTestExternalNonCdb"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_non_container_databases", "test_external_non_container_databases", acctest.Optional, acctest.Update, externalNonContainerDatabaseDataSourceRepresentation) +
				compartmentIdVariableStr + ExternalNonContainerDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Optional, acctest.Update, externalNonContainerDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "myTestExternalNonCdb"),
				resource.TestCheckResourceAttr(datasourceName, "state", "NOT_CONNECTED"),

				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.database_management_config.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.display_name", "myTestExternalNonCdb"),
				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "external_non_container_databases.0.operations_insights_config.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "external_non_container_databases.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_external_non_container_database", "test_external_non_container_database", acctest.Required, acctest.Create, externalNonContainerDatabaseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExternalNonContainerDatabaseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "external_non_container_database_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "database_management_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "myTestExternalNonCdb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "operations_insights_config.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + ExternalNonContainerDatabaseResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseExternalNonContainerDatabaseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_external_non_container_database" {
			noResourceFound = false
			request := oci_database.GetExternalNonContainerDatabaseRequest{}

			tmp := rs.Primary.ID
			request.ExternalNonContainerDatabaseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetExternalNonContainerDatabase(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ExternalNonContainerDatabaseLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseExternalNonContainerDatabase") {
		resource.AddTestSweepers("DatabaseExternalNonContainerDatabase", &resource.Sweeper{
			Name:         "DatabaseExternalNonContainerDatabase",
			Dependencies: acctest.DependencyGraph["externalNonContainerDatabase"],
			F:            sweepDatabaseExternalNonContainerDatabaseResource,
		})
	}
}

func sweepDatabaseExternalNonContainerDatabaseResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	externalNonContainerDatabaseIds, err := getExternalNonContainerDatabaseIds(compartment)
	if err != nil {
		return err
	}
	for _, externalNonContainerDatabaseId := range externalNonContainerDatabaseIds {
		if ok := acctest.SweeperDefaultResourceId[externalNonContainerDatabaseId]; !ok {
			deleteExternalNonContainerDatabaseRequest := oci_database.DeleteExternalNonContainerDatabaseRequest{}

			deleteExternalNonContainerDatabaseRequest.ExternalNonContainerDatabaseId = &externalNonContainerDatabaseId

			deleteExternalNonContainerDatabaseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteExternalNonContainerDatabase(context.Background(), deleteExternalNonContainerDatabaseRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalNonContainerDatabase %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalNonContainerDatabaseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalNonContainerDatabaseId, externalNonContainerDatabaseSweepWaitCondition, time.Duration(3*time.Minute),
				externalNonContainerDatabaseSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getExternalNonContainerDatabaseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalNonContainerDatabaseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listExternalNonContainerDatabasesRequest := oci_database.ListExternalNonContainerDatabasesRequest{}
	listExternalNonContainerDatabasesRequest.CompartmentId = &compartmentId
	listExternalNonContainerDatabasesRequest.LifecycleState = oci_database.ExternalDatabaseBaseLifecycleStateAvailable
	listExternalNonContainerDatabasesResponse, err := databaseClient.ListExternalNonContainerDatabases(context.Background(), listExternalNonContainerDatabasesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalNonContainerDatabase list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalNonContainerDatabase := range listExternalNonContainerDatabasesResponse.Items {
		id := *externalNonContainerDatabase.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalNonContainerDatabaseId", id)
	}
	return resourceIds, nil
}

func externalNonContainerDatabaseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalNonContainerDatabaseResponse, ok := response.Response.(oci_database.GetExternalNonContainerDatabaseResponse); ok {
		return externalNonContainerDatabaseResponse.LifecycleState != oci_database.ExternalNonContainerDatabaseLifecycleStateTerminated
	}
	return false
}

func externalNonContainerDatabaseSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetExternalNonContainerDatabase(context.Background(), oci_database.GetExternalNonContainerDatabaseRequest{
		ExternalNonContainerDatabaseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
