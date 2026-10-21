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
	DatabaseDatabaseSnapshotSingularDataSourceRepresentation = map[string]interface{}{
		"database_snapshot_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database_snapshot.test_database_snapshot.id}`},
	}

	DatabaseDatabaseSnapshotDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_database_snapshot.test_database_snapshot.id}`}},
	}

	DatabaseDatabaseSnapshotRepresentation = map[string]interface{}{
		"database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.test_database.id}`},
		"name":          acctest.Representation{RepType: acctest.Required, Create: `snapshot-1`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	DatabaseDatabaseSnapshotSourceDbHomeRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${var.exadb_vm_cluster_id}`},
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_home_version}`},
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `tf-db-home-snapshot`},
		"source":       acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
	}

	DatabaseDatabaseSnapshotSourceDatabaseRepresentation = map[string]interface{}{
		"database":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabaseSnapshotSourceDatabaseDetailsRepresentation},
		"db_home_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home.id}`},
		"source":     acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"lifecycle":  acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseIgnoreDefinedTagsRepresentation},
	}

	DatabaseDatabaseSnapshotSourceDatabaseDetailsRepresentation = map[string]interface{}{
		"admin_password":       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"db_name":              acctest.Representation{RepType: acctest.Required, Create: `XSSNAP`},
		"storage_size_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabaseSnapshotSourceDatabaseStorageSizeDetailsRepresentation},
	}

	DatabaseDatabaseSnapshotSourceDatabaseStorageSizeDetailsRepresentation = map[string]interface{}{
		"data_storage_size_in_gb":  acctest.Representation{RepType: acctest.Required, Create: `60`},
		"reco_storage_size_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `44`},
	}
)

func databaseDatabaseSnapshotResourceDependencies(clusterVariableName string, dbName string, expectedShapeAttribute string) string {
	sourceDbHomeRepresentation := acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseSnapshotSourceDbHomeRepresentation, map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: fmt.Sprintf("${var.%s}", clusterVariableName)},
		"db_version":   acctest.Representation{RepType: acctest.Required, Create: `${var.db_home_version}`},
	})
	dbHomeVersion := "19.0.0.0"
	dbHomeDisplayName := "tf-db-home-snapshot-block"

	sourceDatabaseDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(DatabaseDatabaseSnapshotSourceDatabaseDetailsRepresentation, []string{"storage_size_details"})
	sourceDatabaseDetailsRepresentation = acctest.RepresentationCopyWithNewProperties(sourceDatabaseDetailsRepresentation, map[string]interface{}{
		"db_name": acctest.Representation{RepType: acctest.Required, Create: dbName},
	})
	if expectedShapeAttribute == "BLOCK_STORAGE" {
		sourceDatabaseDetailsRepresentation = acctest.RepresentationCopyWithNewProperties(sourceDatabaseDetailsRepresentation, map[string]interface{}{
			"storage_size_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabaseSnapshotSourceDatabaseStorageSizeDetailsRepresentation},
		})
	} else if expectedShapeAttribute == "SMART_STORAGE" {
		dbHomeVersion = "26.0.0.0.0"
		dbHomeDisplayName = "tf-db-home-snapshot-smart"
	}
	sourceDbHomeRepresentation = acctest.RepresentationCopyWithNewProperties(sourceDbHomeRepresentation, map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: dbHomeDisplayName},
	})

	sourceDatabaseRepresentation := acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseSnapshotSourceDatabaseRepresentation, map[string]interface{}{
		"database": acctest.RepresentationGroup{RepType: acctest.Required, Group: sourceDatabaseDetailsRepresentation},
	})

	return fmt.Sprintf(`
		variable "%s" {}
		variable "db_home_version" { default = "%s" }

		data "oci_database_exadb_vm_cluster" "test_exadb_vm_cluster" {
			exadb_vm_cluster_id = var.%s
		}
	`, clusterVariableName, dbHomeVersion, clusterVariableName) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Optional, acctest.Create, sourceDbHomeRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, sourceDatabaseRepresentation)
}

func databaseDatabaseSnapshotRepresentation(snapshotName string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseSnapshotRepresentation, map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: snapshotName, Update: snapshotName},
	})
}

func databaseDatabaseSnapshotDataSourceRepresentation(clusterVariableName string, snapshotName string) map[string]interface{} {
	return map[string]interface{}{
		"cluster_id":  acctest.Representation{RepType: acctest.Optional, Create: fmt.Sprintf("${var.%s}", clusterVariableName)},
		"database_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_database.test_database.id}`},
		"name":        acctest.Representation{RepType: acctest.Optional, Create: snapshotName, Update: snapshotName},
		"state":       acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"filter":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseDatabaseSnapshotDataSourceFilterRepresentation},
	}
}

// Pre-requisite to run the TestDatabaseDatabaseSnapshotResource_exadbxs_block_storage_basic:
//
//	set env variable TF_VAR_exadb_vm_cluster_block_storage_id to the ocid of a ExaDB-XS VM Cluster with BLOCK_STORAGE shapeAttribute
//
// issue-routing-tag: database/default
func TestDatabaseDatabaseSnapshotResource_exadbxs_block_storage_basic(t *testing.T) {
	testDatabaseDatabaseSnapshotResourceExadbXs(t, "TestDatabaseDatabaseSnapshotResource_exadbxs_block_storage_basic", "exadb_vm_cluster_block_storage_id", "BLOCK_STORAGE", "tfsnapb", "tf-snapshot-block-required", "tf-snapshot-block-optional")
}

// Pre-requisite to run the TestDatabaseDatabaseSnapshotResource_exadbxs_smart_storage_basic:
//
//	set env variable TF_VAR_exadb_vm_cluster_smart_storage_id to the ocid of a ExaDB-XS VM Cluster with SMART_STORAGE shapeAttribute
//
// issue-routing-tag: database/default
func TestDatabaseDatabaseSnapshotResource_exadbxs_smart_storage_basic(t *testing.T) {
	testDatabaseDatabaseSnapshotResourceExadbXs(t, "TestDatabaseDatabaseSnapshotResource_exadbxs_smart_storage_basic", "exadb_vm_cluster_smart_storage_id", "SMART_STORAGE", "tfsnaps", "tf-snapshot-smart-required", "tf-snapshot-smart-optional")
}

func testDatabaseDatabaseSnapshotResourceExadbXs(t *testing.T, scenarioName string, clusterVariableName string, expectedShapeAttribute string, dbName string, requiredOnlySnapshotName string, optionalSnapshotName string) {
	httpreplay.SetScenario(scenarioName)
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database_snapshot.test_database_snapshot"
	datasourceName := "data.oci_database_database_snapshots.test_database_snapshots"
	singularDatasourceName := "data.oci_database_database_snapshot.test_database_snapshot"
	clusterDatasourceName := "data.oci_database_exadb_vm_cluster.test_exadb_vm_cluster"

	resourceDependencies := databaseDatabaseSnapshotResourceDependencies(clusterVariableName, dbName, expectedShapeAttribute)
	requiredOnlySnapshotRepresentation := databaseDatabaseSnapshotRepresentation(requiredOnlySnapshotName)
	optionalSnapshotRepresentation := databaseDatabaseSnapshotRepresentation(optionalSnapshotName)
	dataSourceRepresentation := databaseDatabaseSnapshotDataSourceRepresentation(clusterVariableName, optionalSnapshotName)
	requiredOnlyResource := resourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Required, acctest.Create, optionalSnapshotRepresentation)
	resourceConfig := resourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Optional, acctest.Update, optionalSnapshotRepresentation)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+resourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Optional, acctest.Create, optionalSnapshotRepresentation), "database", "databaseSnapshot", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseSnapshotDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + resourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Required, acctest.Create, requiredOnlySnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(clusterDatasourceName, "shape_attribute", expectedShapeAttribute),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "name", requiredOnlySnapshotName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + resourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + resourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Optional, acctest.Create, optionalSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(clusterDatasourceName, "shape_attribute", expectedShapeAttribute),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", optionalSnapshotName),

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
			Config: config + compartmentIdVariableStr + resourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Optional, acctest.Update, optionalSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(clusterDatasourceName, "shape_attribute", expectedShapeAttribute),
				resource.TestCheckResourceAttrSet(resourceName, "database_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", optionalSnapshotName),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_snapshots", "test_database_snapshots", acctest.Optional, acctest.Update, dataSourceRepresentation) +
				compartmentIdVariableStr + resourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Optional, acctest.Update, optionalSnapshotRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(clusterDatasourceName, "shape_attribute", expectedShapeAttribute),
				resource.TestCheckResourceAttrSet(datasourceName, "cluster_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "database_id"),
				resource.TestCheckResourceAttr(datasourceName, "name", optionalSnapshotName),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),

				resource.TestCheckResourceAttr(datasourceName, "database_snapshot_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_snapshot_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_database_snapshot", "test_database_snapshot", acctest.Required, acctest.Create, DatabaseDatabaseSnapshotSingularDataSourceRepresentation) +
				compartmentIdVariableStr + resourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(clusterDatasourceName, "shape_attribute", expectedShapeAttribute),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_snapshot_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", optionalSnapshotName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + requiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseDatabaseSnapshotDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_database_snapshot" {
			noResourceFound = false
			request := oci_database.GetDatabaseSnapshotRequest{}

			tmp := rs.Primary.ID
			request.DatabaseSnapshotId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetDatabaseSnapshot(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.DatabaseSnapshotLifecycleStateTerminated): true,
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
	if !acctest.InSweeperExcludeList("DatabaseDatabaseSnapshot") {
		resource.AddTestSweepers("DatabaseDatabaseSnapshot", &resource.Sweeper{
			Name:         "DatabaseDatabaseSnapshot",
			Dependencies: acctest.DependencyGraph["databaseSnapshot"],
			F:            sweepDatabaseDatabaseSnapshotResource,
		})
	}
}

func sweepDatabaseDatabaseSnapshotResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	databaseSnapshotIds, err := getDatabaseDatabaseSnapshotIds(compartment)
	if err != nil {
		return err
	}
	for _, databaseSnapshotId := range databaseSnapshotIds {
		if ok := acctest.SweeperDefaultResourceId[databaseSnapshotId]; !ok {
			deleteDatabaseSnapshotRequest := oci_database.DeleteDatabaseSnapshotRequest{}

			deleteDatabaseSnapshotRequest.DatabaseSnapshotId = &databaseSnapshotId

			deleteDatabaseSnapshotRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteDatabaseSnapshot(context.Background(), deleteDatabaseSnapshotRequest)
			if error != nil {
				fmt.Printf("Error deleting DatabaseSnapshot %s %s, It is possible that the resource is already deleted. Please verify manually \n", databaseSnapshotId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &databaseSnapshotId, DatabaseDatabaseSnapshotSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseDatabaseSnapshotSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseDatabaseSnapshotIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DatabaseSnapshotId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listDatabaseSnapshotsRequest := oci_database.ListDatabaseSnapshotsRequest{}
	listDatabaseSnapshotsRequest.CompartmentId = &compartmentId
	listDatabaseSnapshotsRequest.LifecycleState = oci_database.DatabaseSnapshotLifecycleStateAvailable
	listDatabaseSnapshotsResponse, err := databaseClient.ListDatabaseSnapshots(context.Background(), listDatabaseSnapshotsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DatabaseSnapshot list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, databaseSnapshot := range listDatabaseSnapshotsResponse.Items {
		id := *databaseSnapshot.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DatabaseSnapshotId", id)
	}
	return resourceIds, nil
}

func DatabaseDatabaseSnapshotSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if databaseSnapshotResponse, ok := response.Response.(oci_database.GetDatabaseSnapshotResponse); ok {
		return databaseSnapshotResponse.LifecycleState != oci_database.DatabaseSnapshotLifecycleStateTerminated
	}
	return false
}

func DatabaseDatabaseSnapshotSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetDatabaseSnapshot(context.Background(), oci_database.GetDatabaseSnapshotRequest{
		DatabaseSnapshotId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
