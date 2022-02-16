// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	tf_database "github.com/terraform-providers/terraform-provider-oci/internal/service/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	DbNodeConsoleConnectionResourceConfig = DbNodeConsoleConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Update, dbNodeConsoleConnectionRepresentation)

	dbNodeConsoleConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"id":         acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_node_console_connection.test_db_node_console_connection.id}`},
	}

	dbNodeConsoleConnectionDataSourceRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: dbNodeConsoleConnectionDataSourceFilterRepresentation}}
	dbNodeConsoleConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_db_node_console_connection.test_db_node_console_connection.id}`}},
	}

	dbNodeConsoleConnectionRepresentation = map[string]interface{}{
		"db_node_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_node.test_db_node.id}`},
		"public_key": acctest.Representation{RepType: acctest.Required, Create: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`},
	}

	DbNodeConsoleConnectionResourceDependencies = DbSystemResourceConfig + `
		data "oci_database_db_nodes" "test_db_nodes" {
			compartment_id = "${var.compartment_id}"
			db_system_id = "${oci_database_db_system.test_db_system.id}"
		}
		data "oci_database_db_node" "test_db_node" {
			db_node_id = "${data.oci_database_db_nodes.test_db_nodes.db_nodes.0.id}"
		}`
)

// issue-routing-tag: database/default
func TestDatabaseDbNodeConsoleConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeConsoleConnectionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_node_console_connection.test_db_node_console_connection"
	datasourceName := "data.oci_database_db_node_console_connections.test_db_node_console_connections"
	singularDatasourceName := "data.oci_database_db_node_console_connection.test_db_node_console_connection"

	var resId string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DbNodeConsoleConnectionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, dbNodeConsoleConnectionRepresentation), "database", "dbNodeConsoleConnection", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDbNodeConsoleConnectionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DbNodeConsoleConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, dbNodeConsoleConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
				resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),

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

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_connections", "test_db_node_console_connections", acctest.Optional, acctest.Update, dbNodeConsoleConnectionDataSourceRepresentation) +
				compartmentIdVariableStr + DbNodeConsoleConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Optional, acctest.Update, dbNodeConsoleConnectionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "db_node_id"),

				resource.TestCheckResourceAttr(datasourceName, "console_connections.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.compartment_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.connection_string"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.db_node_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.fingerprint"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "console_connections.0.state"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", acctest.Required, acctest.Create, dbNodeConsoleConnectionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DbNodeConsoleConnectionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "connection_string"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fingerprint"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + DbNodeConsoleConnectionResourceConfig,
		},
		// verify resource import
		{
			Config:            config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"public_key",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseDbNodeConsoleConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_node_console_connection" {
			noResourceFound = false
			request := oci_database.GetConsoleConnectionRequest{}

			if value, ok := rs.Primary.Attributes["db_node_id"]; ok {
				request.DbNodeId = &value
			}

			tmp := rs.Primary.ID
			request.ConsoleConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetConsoleConnection(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.ConsoleConnectionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseDbNodeConsoleConnection") {
		resource.AddTestSweepers("DatabaseDbNodeConsoleConnection", &resource.Sweeper{
			Name:         "DatabaseDbNodeConsoleConnection",
			Dependencies: acctest.DependencyGraph["dbNodeConsoleConnection"],
			F:            sweepDatabaseDbNodeConsoleConnectionResource,
		})
	}
}

func sweepDatabaseDbNodeConsoleConnectionResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	dbNodeConsoleConnectionIds, err := getDbNodeConsoleConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, dbNodeConsoleConnectionId := range dbNodeConsoleConnectionIds {
		if ok := acctest.SweeperDefaultResourceId[dbNodeConsoleConnectionId]; !ok {
			deleteConsoleConnectionRequest := oci_database.DeleteConsoleConnectionRequest{}

			dbNodeId, id, err := tf_database.ParseDbNodeConsoleConnectionCompositeId(dbNodeConsoleConnectionId)
			if err == nil {
				deleteConsoleConnectionRequest.DbNodeId = &dbNodeId
				deleteConsoleConnectionRequest.ConsoleConnectionId = &id
			} else {
				log.Printf("[WARN] sweepDatabaseDbNodeConsoleConnectionResource() unable to parse current ID: %s", dbNodeConsoleConnectionId)
			}

			deleteConsoleConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteConsoleConnection(context.Background(), deleteConsoleConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting DbNodeConsoleConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbNodeConsoleConnectionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbNodeConsoleConnectionId, dbNodeConsoleConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				dbNodeConsoleConnectionSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDbNodeConsoleConnectionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbNodeConsoleConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listConsoleConnectionsRequest := oci_database.ListConsoleConnectionsRequest{}

	dbNodeIds, error := getDbNodeIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting dbNodeId required for DbNodeConsoleConnection resource requests \n")
	}
	for _, dbNodeId := range dbNodeIds {
		listConsoleConnectionsRequest.DbNodeId = &dbNodeId

		listConsoleConnectionsResponse, err := databaseClient.ListConsoleConnections(context.Background(), listConsoleConnectionsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting DbNodeConsoleConnection list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, dbNodeConsoleConnection := range listConsoleConnectionsResponse.Items {
			id := tf_database.GetDbNodeConsoleConnectionCompositeId(dbNodeId, *dbNodeConsoleConnection.Id)
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbNodeConsoleConnectionId", id)
		}

	}
	return resourceIds, nil
}

func dbNodeConsoleConnectionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbNodeConsoleConnectionResponse, ok := response.Response.(oci_database.GetConsoleConnectionResponse); ok {
		return dbNodeConsoleConnectionResponse.LifecycleState != oci_database.ConsoleConnectionLifecycleStateDeleted
	}
	return false
}

func dbNodeConsoleConnectionSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetConsoleConnection(context.Background(), oci_database.GetConsoleConnectionRequest{
		ConsoleConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
