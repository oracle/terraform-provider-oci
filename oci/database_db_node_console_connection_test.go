// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_database "github.com/oracle/oci-go-sdk/database"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	DbNodeConsoleConnectionResourceConfig = DbNodeConsoleConnectionResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", Optional, Update, dbNodeConsoleConnectionRepresentation)

	dbNodeConsoleConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"db_node_id": Representation{repType: Required, create: `${data.oci_database_db_node.test_db_node.id}`},
		"id":         Representation{repType: Required, create: `${oci_database_db_node_console_connection.test_db_node_console_connection.id}`},
	}

	dbNodeConsoleConnectionDataSourceRepresentation = map[string]interface{}{
		"db_node_id": Representation{repType: Required, create: `${data.oci_database_db_node.test_db_node.id}`},
		"filter":     RepresentationGroup{Required, dbNodeConsoleConnectionDataSourceFilterRepresentation}}
	dbNodeConsoleConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_database_db_node_console_connection.test_db_node_console_connection.id}`}},
	}

	dbNodeConsoleConnectionRepresentation = map[string]interface{}{
		"db_node_id": Representation{repType: Required, create: `${data.oci_database_db_node.test_db_node.id}`},
		"public_key": Representation{repType: Required, create: `ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin`},
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

func TestDatabaseDbNodeConsoleConnectionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeConsoleConnectionResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_node_console_connection.test_db_node_console_connection"
	datasourceName := "data.oci_database_db_node_console_connections.test_db_node_console_connections"
	singularDatasourceName := "data.oci_database_db_node_console_connection.test_db_node_console_connection"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatabaseDbNodeConsoleConnectionDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + DbNodeConsoleConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", Required, Create, dbNodeConsoleConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "db_node_id"),
					resource.TestCheckResourceAttr(resourceName, "public_key", "ssh-rsa KKKLK3NzaC1yc2EAAAADAQABAAABAQC+UC9MFNA55NIVtKPIBCNw7++ACXhD0hx+Zyj25JfHykjz/QU3Q5FAU3DxDbVXyubgXfb/GJnrKRY8O4QDdvnZZRvQFFEOaApThAmCAM5MuFUIHdFvlqP+0W+ZQnmtDhwVe2NCfcmOrMuaPEgOKO3DOW6I/qOOdO691Xe2S9NgT9HhN0ZfFtEODVgvYulgXuCCXsJs+NUqcHAOxxFUmwkbPvYi0P0e2DT8JKeiOOC8VKUEgvVx+GKmqasm+Y6zHFW7vv3g2GstE1aRs3mttHRoC/JPM86PRyIxeWXEMzyG5wHqUu4XZpDbnWNxi6ugxnAGiL3CrIFdCgRNgHz5qS1l MustWin"),
				),
			},

			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_db_node_console_connections", "test_db_node_console_connections", Optional, Update, dbNodeConsoleConnectionDataSourceRepresentation) +
					compartmentIdVariableStr + DbNodeConsoleConnectionResourceDependencies +
					generateResourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", Optional, Update, dbNodeConsoleConnectionRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_db_node_console_connection", "test_db_node_console_connection", Required, Create, dbNodeConsoleConnectionSingularDataSourceRepresentation) +
					compartmentIdVariableStr + DbNodeConsoleConnectionResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func testAccCheckDatabaseDbNodeConsoleConnectionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).databaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_db_node_console_connection" {
			noResourceFound = false
			request := oci_database.GetConsoleConnectionRequest{}

			if value, ok := rs.Primary.Attributes["db_node_id"]; ok {
				request.DbNodeId = &value
			}

			tmp := rs.Primary.ID
			request.ConsoleConnectionId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("DatabaseDbNodeConsoleConnection") {
		resource.AddTestSweepers("DatabaseDbNodeConsoleConnection", &resource.Sweeper{
			Name:         "DatabaseDbNodeConsoleConnection",
			Dependencies: DependencyGraph["dbNodeConsoleConnection"],
			F:            sweepDatabaseDbNodeConsoleConnectionResource,
		})
	}
}

func sweepDatabaseDbNodeConsoleConnectionResource(compartment string) error {
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()
	dbNodeConsoleConnectionIds, err := getDbNodeConsoleConnectionIds(compartment)
	if err != nil {
		return err
	}
	for _, dbNodeConsoleConnectionId := range dbNodeConsoleConnectionIds {
		if ok := SweeperDefaultResourceId[dbNodeConsoleConnectionId]; !ok {
			deleteConsoleConnectionRequest := oci_database.DeleteConsoleConnectionRequest{}

			dbNodeId, id, err := parseDbNodeConsoleConnectionCompositeId(dbNodeConsoleConnectionId)
			if err == nil {
				deleteConsoleConnectionRequest.DbNodeId = &dbNodeId
				deleteConsoleConnectionRequest.ConsoleConnectionId = &id
			} else {
				log.Printf("[WARN] sweepDatabaseDbNodeConsoleConnectionResource() unable to parse current ID: %s", dbNodeConsoleConnectionId)
			}

			deleteConsoleConnectionRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "database")
			_, error := databaseClient.DeleteConsoleConnection(context.Background(), deleteConsoleConnectionRequest)
			if error != nil {
				fmt.Printf("Error deleting DbNodeConsoleConnection %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbNodeConsoleConnectionId, error)
				continue
			}
			waitTillCondition(testAccProvider, &dbNodeConsoleConnectionId, dbNodeConsoleConnectionSweepWaitCondition, time.Duration(3*time.Minute),
				dbNodeConsoleConnectionSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDbNodeConsoleConnectionIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "DbNodeConsoleConnectionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := GetTestClients(&schema.ResourceData{}).databaseClient()

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
			id := getDbNodeConsoleConnectionCompositeId(dbNodeId, *dbNodeConsoleConnection.Id)
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "DbNodeConsoleConnectionId", id)
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

func dbNodeConsoleConnectionSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.databaseClient().GetConsoleConnection(context.Background(), oci_database.GetConsoleConnectionRequest{
		ConsoleConnectionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
