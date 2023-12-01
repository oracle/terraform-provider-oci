// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDbNodeConsoleHistoryRequiredOnlyResource = DatabaseDbNodeConsoleHistoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Required, acctest.Create, DatabaseDbNodeConsoleHistoryRepresentation)

	DatabaseDbNodeConsoleHistoryResourceConfig = DatabaseDbNodeConsoleHistoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleHistoryRepresentation)

	DatabaseDbNodeConsoleHistorySingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id":         acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"console_history_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_node_console_history.test_db_node_console_history.id}`},
	}

	DatabaseDbNodeConsoleHistoryDataSourceRepresentation = map[string]interface{}{
		"db_node_id":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `console-history-20221202-1943`, Update: `displayName2`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `SUCCEEDED`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbNodeConsoleHistoryDataSourceFilterRepresentation}}
	DatabaseDbNodeConsoleHistoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `db_node_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_node.test_db_node.id}`}},
	}

	DatabaseDbNodeConsoleHistoryRepresentation = map[string]interface{}{
		"db_node_id":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Required, Create: `console-history-20221202-1943`, Update: `displayName2`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseDbNodeConsoleHistoryResourceDependencies = DatabaseVmClusterRequiredOnlyResource + AvailabilityDomainConfig +
		DefinedTagsDependencies +
		`
     data "oci_database_db_nodes" "test_db_nodes" {
        compartment_id = "${var.compartment_id}"
        vm_cluster_id = "${oci_database_vm_cluster.test_vm_cluster.id}"
     }
     data "oci_database_db_node" "test_db_node" {
        db_node_id = "${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}"
     }`
)

// issue-routing-tag: database/default
func TestDatabaseDbNodeConsoleHistoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeConsoleHistoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_node_console_history.test_db_node_console_history"
	datasourceName := "data.oci_database_db_node_console_histories.test_db_node_console_histories"
	singularDatasourceName := "data.oci_database_db_node_console_history.test_db_node_console_history"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDbNodeConsoleHistoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Optional, acctest.Create, DatabaseDbNodeConsoleHistoryRepresentation), "database", "dbNodeConsoleHistory", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDbNodeConsoleHistoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Required, acctest.Create, DatabaseDbNodeConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "console-history-20221202-1943"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Optional, acctest.Create, DatabaseDbNodeConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "console-history-20221202-1943"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_histories", "test_db_node_console_histories", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleHistoryDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Optional, acctest.Update, DatabaseDbNodeConsoleHistoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_node_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "SUCCEEDED"),

				resource.TestCheckResourceAttr(datasourceName, "console_history_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "console_history_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_history", "test_db_node_console_history", acctest.Required, acctest.Create, DatabaseDbNodeConsoleHistorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseDbNodeConsoleHistoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "console_history_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseDbNodeConsoleHistoryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getDatabaseConsoleHistoryCompositeId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getDatabaseConsoleHistoryCompositeId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("dbNodes/" + rs.Primary.Attributes["db_node_id"] + "/consoleHistories/" + rs.Primary.Attributes["id"]), nil
	}
}

func testAccCheckDatabaseDbNodeConsoleHistoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_node_console_history" {
			noResourceFound = false
			request := oci_database.GetConsoleHistoryRequest{}

			if value, ok := rs.Primary.Attributes["id"]; ok {
				request.ConsoleHistoryId = &value
			}

			if value, ok := rs.Primary.Attributes["db_node_id"]; ok {
				request.DbNodeId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetConsoleHistory(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ConsoleHistoryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseDbNodeConsoleHistory") {
		resource.AddTestSweepers("DatabaseDbNodeConsoleHistory", &resource.Sweeper{
			Name:         "DatabaseDbNodeConsoleHistory",
			Dependencies: acctest.DependencyGraph["dbNodeConsoleHistory"],
			F:            sweepDatabaseDbNodeConsoleHistoryResource,
		})
	}
}

func sweepDatabaseDbNodeConsoleHistoryResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	dbNodeConsoleHistoryIds, err := getDatabaseDbNodeConsoleHistoryIds(compartment)
	if err != nil {
		return err
	}
	for _, dbNodeConsoleHistoryId := range dbNodeConsoleHistoryIds {
		if ok := acctest.SweeperDefaultResourceId[dbNodeConsoleHistoryId]; !ok {
			deleteConsoleHistoryRequest := oci_database.DeleteConsoleHistoryRequest{}

			deleteConsoleHistoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteConsoleHistory(context.Background(), deleteConsoleHistoryRequest)
			if error != nil {
				fmt.Printf("Error deleting DbNodeConsoleHistory %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbNodeConsoleHistoryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbNodeConsoleHistoryId, DatabaseDbNodeConsoleHistorySweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseDbNodeConsoleHistorySweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseDbNodeConsoleHistoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbNodeConsoleHistoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listConsoleHistoriesRequest := oci_database.ListConsoleHistoriesRequest{}
	//listConsoleHistoriesRequest.CompartmentId = &compartmentId

	dbNodeIds, error := getDbNodeIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting dbNodeId required for DbNodeConsoleHistory resource requests \n")
	}
	for _, dbNodeId := range dbNodeIds {
		listConsoleHistoriesRequest.DbNodeId = &dbNodeId

		listConsoleHistoriesRequest.LifecycleState = oci_database.ConsoleHistorySummaryLifecycleStateSucceeded
		listConsoleHistoriesResponse, err := databaseClient.ListConsoleHistories(context.Background(), listConsoleHistoriesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DbNodeConsoleHistory list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dbNodeConsoleHistory := range listConsoleHistoriesResponse.Items {
			id := *dbNodeConsoleHistory.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbNodeConsoleHistoryId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseDbNodeConsoleHistorySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbNodeConsoleHistoryResponse, ok := response.Response.(oci_database.GetConsoleHistoryResponse); ok {
		return dbNodeConsoleHistoryResponse.LifecycleState != oci_database.ConsoleHistoryLifecycleStateDeleted
	}
	return false
}

func DatabaseDbNodeConsoleHistorySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetConsoleHistory(context.Background(), oci_database.GetConsoleHistoryRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
