// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabasePluggableDatabaseSnapshotRequiredOnlyResource = DatabasePluggableDatabaseSnapshotResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Required, acctest.Create, DatabasePluggableDatabaseSnapshotRepresentation)

	DatabasePluggableDatabaseSnapshotResourceConfig = DatabasePluggableDatabaseSnapshotResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Optional, acctest.Update, DatabasePluggableDatabaseSnapshotRepresentation)

	DatabasePluggableDatabaseSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"pluggable_database_snapshot_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot.id}`},
	}

	DatabasePluggableDatabaseSnapshotDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"pluggable_database_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.pluggable_database_id}`},
		"name":                  acctest.Representation{RepType: acctest.Optional, Create: `TFSnapshot2`},
		"state":                 acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableDatabaseSnapshotDataSourceFilterRepresentation}}

	DatabasePluggableDatabaseSnapshotDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot.id}`}},
	}

	DatabasePluggableDatabaseSnapshotRepresentation = map[string]interface{}{
		"name":                  acctest.Representation{RepType: acctest.Required, Create: `TFSnapshot1`, Update: `TFSnapshot2`},
		"pluggable_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.pluggable_database_id}`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"example-tag-namespace-all.example-tag": "value"}},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}},
		"lifecycle":             acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabasePluggableDatabaseSnapshotIgnoreDefinedTagsRepresentation},
	}

	DatabasePluggableDatabaseSnapshotIgnoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	// Note: set env variable TF_VAR_pluggable_database_id before running this test
	DatabasePluggableDatabaseSnapshotResourceDependencies = `variable "pluggable_database_id" {}`
)

// issue-routing-tag: database/ExaCS
func TestDatabasePluggableDatabaseSnapshotResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabasePluggableDatabaseSnapshotResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot"
	datasourceName := "data.oci_database_pluggable_database_snapshots.test_pluggable_database_snapshots"
	singularDatasourceName := "data.oci_database_pluggable_database_snapshot.test_pluggable_database_snapshot"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabasePluggableDatabaseSnapshotResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Optional, acctest.Create, DatabasePluggableDatabaseSnapshotRepresentation), "database", "pluggableDatabaseSnapshot", t)

	acctest.ResourceTest(t, testAccCheckDatabasePluggableDatabaseSnapshotDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Required, acctest.Create, DatabasePluggableDatabaseSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "TFSnapshot1"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseSnapshotResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabasePluggableDatabaseSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Optional, acctest.Update, DatabasePluggableDatabaseSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "TFSnapshot2"),
				resource.TestCheckResourceAttrSet(resourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "0"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_database_snapshots", "test_pluggable_database_snapshots", acctest.Optional, acctest.Update, DatabasePluggableDatabaseSnapshotDataSourceRepresentation) +
				compartmentIdVariableStr + DatabasePluggableDatabaseSnapshotResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Optional, acctest.Update, DatabasePluggableDatabaseSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_id"),

				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_snapshots.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_snapshots.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_snapshots.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_snapshots.0.cluster_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.0.name", "TFSnapshot2"),
				resource.TestCheckResourceAttrSet(datasourceName, "pluggable_database_snapshots.0.pluggable_database_id"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.0.freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(datasourceName, "pluggable_database_snapshots.0.system_tags.%", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_pluggable_database_snapshot", "test_pluggable_database_snapshot", acctest.Required, acctest.Create, DatabasePluggableDatabaseSnapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabasePluggableDatabaseSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pluggable_database_snapshot_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cluster_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "TFSnapshot2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "pluggable_database_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.Department", "Finance"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "0"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatabasePluggableDatabaseSnapshotRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabasePluggableDatabaseSnapshotDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_pluggable_database_snapshot" {
			noResourceFound = false
			request := oci_database.GetPluggableDatabaseSnapshotRequest{}

			tmp := rs.Primary.ID
			request.PluggableDatabaseSnapshotId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetPluggableDatabaseSnapshot(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.PluggableDatabaseSnapshotLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabasePluggableDatabaseSnapshot") {
		resource.AddTestSweepers("DatabasePluggableDatabaseSnapshot", &resource.Sweeper{
			Name:         "DatabasePluggableDatabaseSnapshot",
			Dependencies: acctest.DependencyGraph["pluggableDatabaseSnapshot"],
			F:            sweepDatabasePluggableDatabaseSnapshotResource,
		})
	}
}

func sweepDatabasePluggableDatabaseSnapshotResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	pluggableDatabaseSnapshotIds, err := getDatabasePluggableDatabaseSnapshotIds(compartment)
	if err != nil {
		return err
	}
	for _, pluggableDatabaseSnapshotId := range pluggableDatabaseSnapshotIds {
		if ok := acctest.SweeperDefaultResourceId[pluggableDatabaseSnapshotId]; !ok {
			deletePluggableDatabaseSnapshotRequest := oci_database.DeletePluggableDatabaseSnapshotRequest{}

			deletePluggableDatabaseSnapshotRequest.PluggableDatabaseSnapshotId = &pluggableDatabaseSnapshotId

			deletePluggableDatabaseSnapshotRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeletePluggableDatabaseSnapshot(context.Background(), deletePluggableDatabaseSnapshotRequest)
			if error != nil {
				fmt.Printf("Error deleting PluggableDatabaseSnapshot %s %s, It is possible that the resource is already deleted. Please verify manually \n", pluggableDatabaseSnapshotId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &pluggableDatabaseSnapshotId, DatabasePluggableDatabaseSnapshotSweepWaitCondition, time.Duration(3*time.Minute),
				DatabasePluggableDatabaseSnapshotSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabasePluggableDatabaseSnapshotIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "PluggableDatabaseSnapshotId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listPluggableDatabaseSnapshotsRequest := oci_database.ListPluggableDatabaseSnapshotsRequest{}
	listPluggableDatabaseSnapshotsRequest.CompartmentId = &compartmentId
	listPluggableDatabaseSnapshotsRequest.LifecycleState = oci_database.PluggableDatabaseSnapshotLifecycleStateAvailable
	listPluggableDatabaseSnapshotsResponse, err := databaseClient.ListPluggableDatabaseSnapshots(context.Background(), listPluggableDatabaseSnapshotsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting PluggableDatabaseSnapshot list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, pluggableDatabaseSnapshot := range listPluggableDatabaseSnapshotsResponse.Items {
		id := *pluggableDatabaseSnapshot.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "PluggableDatabaseSnapshotId", id)
	}
	return resourceIds, nil
}

func DatabasePluggableDatabaseSnapshotSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if pluggableDatabaseSnapshotResponse, ok := response.Response.(oci_database.GetPluggableDatabaseSnapshotResponse); ok {
		return pluggableDatabaseSnapshotResponse.LifecycleState != oci_database.PluggableDatabaseSnapshotLifecycleStateTerminated
	}
	return false
}

func DatabasePluggableDatabaseSnapshotSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetPluggableDatabaseSnapshot(context.Background(), oci_database.GetPluggableDatabaseSnapshotRequest{
		PluggableDatabaseSnapshotId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
