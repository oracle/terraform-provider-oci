// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDbNodeSnapshotResourceConfig = DatabaseDbNodeSnapshotDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot", "test_db_node_snapshot", acctest.Required, acctest.Create, DatabaseDbNodeSnapshotRepresentation)

	DatabaseDbNodeSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"dbnode_snapshot_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_node_snapshot_id}`},
	}

	DatabaseDbNodeSnapshotDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cluster_id":       acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_node.test_db_node.db_system_id}`},
		"name":             acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_node.test_db_node.hostname}-snapshot-test`},
		"source_dbnode_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.db_node_id}`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbNodeSnapshotDataSourceFilterRepresentation},
	}

	DatabaseDbNodeSnapshotDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_db_node_snapshot.test_db_node_snapshot.id}`}},
	}

	DatabaseDbNodeSnapshotRepresentation = map[string]interface{}{
		"dbnode_snapshot_id": acctest.Representation{RepType: acctest.Required, Create: `${var.db_node_snapshot_id}`},
		"mount_dbnode_id":    acctest.Representation{RepType: acctest.Required, Create: `null`, Update: `${var.db_node_id}`},
		"lifecycle":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDbNodeSnapshotIgnoreTagsRepresentation},
	}

	DatabaseDbNodeSnapshotIgnoreTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`system_tags`}},
	}

	// Pre-requisite of this integ test: create a dbnode snapshot with name "snapshot-test" and set the dbnodeId / dbnodeSnapshotId as env variablies
	DatabaseDbNodeSnapshotDependencies = `
        variable db_node_id {}
        variable db_node_snapshot_id {}
        data "oci_database_db_node" "test_db_node" {
            db_node_id = var.db_node_id
        }
    `
)

// issue-routing-tag: database/ExaCS
func TestDatabaseDbNodeSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseDbNodeSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	dbNodeId := utils.GetRequiredEnvSetting("db_node_id")
	dbNodeSnapshotId := utils.GetRequiredEnvSetting("db_node_snapshot_id")

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_db_node_snapshot.test_db_node_snapshot"
	datasourceName := "data.oci_database_db_node_snapshots.test_db_node_snapshots"
	singularDatasourceName := "data.oci_database_db_node_snapshot.test_db_node_snapshot"

	//var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseDbNodeSnapshotDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot", "test_db_node_snapshot", acctest.Optional, acctest.Create, DatabaseDbNodeSnapshotRepresentation), "database", "dbnodesnapshot", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dbnode_snapshot_id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "mount_dbnode_id", "null"),
				resource.TestCheckResourceAttr(resourceName, "mount_points.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "source_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				// commenting out this code due to make test-compile failure
				//func(s *terraform.State) (err error) {
				//	resId, err = acctest.FromInstanceState(s, resourceName, "id")
				//	return err
				//},
			),
		},
		// verify Update - Mount
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_db_node_snapshot", "test_db_node_snapshot", acctest.Required, acctest.Update, DatabaseDbNodeSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dbnode_snapshot_id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "mount_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(resourceName, "mount_points.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "mount_points.0.db_node_id", dbNodeId),
				resource.TestCheckResourceAttrSet(resourceName, "mount_points.0.name"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "source_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(resourceName, "state", "MOUNTED"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				// ommenting out this code due to make test-compile failure
				//func(s *terraform.State) (err error) {
				//	resId2, err = acctest.FromInstanceState(s, resourceName, "id")
				//	if resId != resId2 {
				//		return fmt.Errorf("resource recreated when it was supposed to be updated")
				//	}
				//	return err
				//},
			),
		},
		// verify Update - Unmount
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "dbnode_snapshot_id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(resourceName, "mount_dbnode_id", "null"),
				resource.TestCheckResourceAttr(resourceName, "mount_points.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "name"),
				resource.TestCheckResourceAttr(resourceName, "source_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				// commenting out this code due to make test-compile failure
				//func(s *terraform.State) (err error) {
				//	resId2, err = acctest.FromInstanceState(s, resourceName, "id")
				//	if resId != resId2 {
				//		return fmt.Errorf("resource recreated when it was supposed to be updated")
				//	}
				//	return err
				//},
			),
		},
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_snapshots", "test_db_node_snapshots", acctest.Optional, acctest.Update, DatabaseDbNodeSnapshotDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "name"),
				resource.TestCheckResourceAttrSet(datasourceName, "source_dbnode_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttrSet(datasourceName, "dbnode_snapshots.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "dbnode_snapshots.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.dbnode_snapshot_id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.mount_dbnode_id", "null"),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.mount_points.#", "0"),
				resource.TestCheckResourceAttrSet(datasourceName, "dbnode_snapshots.0.name"),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.source_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(datasourceName, "dbnode_snapshots.0.state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "dbnode_snapshots.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + DatabaseDbNodeSnapshotResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_node_snapshot", "test_db_node_snapshot", acctest.Required, acctest.Create, DatabaseDbNodeSnapshotSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dbnode_snapshot_id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(singularDatasourceName, "id", dbNodeSnapshotId),
				resource.TestCheckResourceAttr(singularDatasourceName, "mount_dbnode_id", "null"),
				resource.TestCheckResourceAttr(singularDatasourceName, "mount_points.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "source_dbnode_id", dbNodeId),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabaseDbNodeSnapshotResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
		// delete
		{
			Config: config + compartmentIdVariableStr,
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseDbNodeSnapshot") {
		resource.AddTestSweepers("DatabaseDbNodeSnapshot", &resource.Sweeper{
			Name:         "DatabaseDbNodeSnapshot",
			Dependencies: acctest.DependencyGraph["dbNodeSnapshot"],
			F:            sweepDatabaseDbNodeSnapshotResource,
		})
	}
}

func sweepDatabaseDbNodeSnapshotResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	dbNodeSnapshotIds, err := getDatabaseDbNodeSnapshotIds(compartment)
	if err != nil {
		return err
	}
	for _, dbNodeSnapshotId := range dbNodeSnapshotIds {
		if ok := acctest.SweeperDefaultResourceId[dbNodeSnapshotId]; !ok {
			deleteDbnodeSnapshotRequest := oci_database.DeleteDbnodeSnapshotRequest{}

			deleteDbnodeSnapshotRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDbnodeSnapshot(context.Background(), deleteDbnodeSnapshotRequest)
			if error != nil {
				fmt.Printf("Error deleting DbNodeSnapshot %s %s, It is possible that the resource is already deleted. Please verify manually \n", dbNodeSnapshotId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &dbNodeSnapshotId, DatabaseDbNodeSnapshotSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseDbNodeSnapshotSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseDbNodeSnapshotIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DbNodeSnapshotId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listDbnodeSnapshotsRequest := oci_database.ListDbnodeSnapshotsRequest{}
	listDbnodeSnapshotsRequest.CompartmentId = &compartmentId
	listDbnodeSnapshotsRequest.LifecycleState = oci_database.DbnodeSnapshotLifecycleStateAvailable
	listDbnodeSnapshotsResponse, err := databaseClient.ListDbnodeSnapshots(context.Background(), listDbnodeSnapshotsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DbNodeSnapshot list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, dbNodeSnapshot := range listDbnodeSnapshotsResponse.Items {
		id := *dbNodeSnapshot.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DbNodeSnapshotId", id)
	}
	return resourceIds, nil
}

func DatabaseDbNodeSnapshotSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if dbNodeSnapshotResponse, ok := response.Response.(oci_database.GetDbnodeSnapshotResponse); ok {
		return dbNodeSnapshotResponse.LifecycleState != oci_database.DbnodeSnapshotLifecycleStateTerminated
	}
	return false
}

func DatabaseDbNodeSnapshotSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDbnodeSnapshot(context.Background(), oci_database.GetDbnodeSnapshotRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
