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
	DatabaseAdvancedClusterFileSystemRequiredOnlyResource = DatabaseAdvancedClusterFileSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Required, acctest.Create, DatabaseAdvancedClusterFileSystemRepresentation)

	DatabaseAdvancedClusterFileSystemResourceConfig = DatabaseAdvancedClusterFileSystemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Update, DatabaseAdvancedClusterFileSystemRepresentation)

	DatabaseAdvancedClusterFileSystemSingularDataSourceRepresentation = map[string]interface{}{
		"advanced_cluster_file_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system.id}`},
	}

	DatabaseAdvancedClusterFileSystemDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `fileSystemName`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAdvancedClusterFileSystemDataSourceFilterRepresentation}}
	DatabaseAdvancedClusterFileSystemDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system.id}`}},
	}

	DatabaseAdvancedClusterFileSystemRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `fileSystemName`},
		"storage_in_gbs": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"vm_cluster_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseAdvancedClusterFileSystemResourceDependencies = DatabaseExascaleVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseExascaleVmClusterRepresentation)
)

// issue-routing-tag: database/default
func TestDatabaseAdvancedClusterFileSystemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdvancedClusterFileSystemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system"
	datasourceName := "data.oci_database_advanced_cluster_file_systems.test_advanced_cluster_file_systems"
	singularDatasourceName := "data.oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdvancedClusterFileSystemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Create, DatabaseAdvancedClusterFileSystemRepresentation), "database", "advancedClusterFileSystem", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAdvancedClusterFileSystemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Create, DatabaseAdvancedClusterFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "fileSystemName"),
				resource.TestCheckResourceAttr(resourceName, "storage_in_gbs", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Create, DatabaseAdvancedClusterFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_mounted"),
				resource.TestCheckResourceAttrSet(resourceName, "mount_point"),
				resource.TestCheckResourceAttr(resourceName, "name", "fileSystemName"),
				resource.TestCheckResourceAttr(resourceName, "storage_in_gbs", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_id"),

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
			Config: config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Update, DatabaseAdvancedClusterFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "is_mounted"),
				resource.TestCheckResourceAttrSet(resourceName, "mount_point"),
				resource.TestCheckResourceAttr(resourceName, "name", "fileSystemName"),
				resource.TestCheckResourceAttr(resourceName, "storage_in_gbs", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "vm_cluster_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_advanced_cluster_file_systems", "test_advanced_cluster_file_systems", acctest.Optional, acctest.Update, DatabaseAdvancedClusterFileSystemDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Optional, acctest.Update, DatabaseAdvancedClusterFileSystemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "fileSystemName"),
				resource.TestCheckResourceAttr(datasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_id"),

				resource.TestCheckResourceAttr(datasourceName, "advanced_cluster_file_system_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "advanced_cluster_file_system_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_advanced_cluster_file_system", "test_advanced_cluster_file_system", acctest.Required, acctest.Create, DatabaseAdvancedClusterFileSystemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "advanced_cluster_file_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_mounted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mount_point"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "fileSystemName"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "AVAILABLE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_in_gbs", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vault_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + compartmentIdVariableStr + DatabaseAdvancedClusterFileSystemRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckDatabaseAdvancedClusterFileSystemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_advanced_cluster_file_system" {
			noResourceFound = false
			request := oci_database.GetAdvancedClusterFileSystemRequest{}

			tmp := rs.Primary.ID
			request.AdvancedClusterFileSystemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetAdvancedClusterFileSystem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.AdvancedClusterFileSystemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseAdvancedClusterFileSystem") {
		resource.AddTestSweepers("DatabaseAdvancedClusterFileSystem", &resource.Sweeper{
			Name:         "DatabaseAdvancedClusterFileSystem",
			Dependencies: acctest.DependencyGraph["advancedClusterFileSystem"],
			F:            sweepDatabaseAdvancedClusterFileSystemResource,
		})
	}
}

func sweepDatabaseAdvancedClusterFileSystemResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	advancedClusterFileSystemIds, err := getDatabaseAdvancedClusterFileSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, advancedClusterFileSystemId := range advancedClusterFileSystemIds {
		if ok := acctest.SweeperDefaultResourceId[advancedClusterFileSystemId]; !ok {
			deleteAdvancedClusterFileSystemRequest := oci_database.DeleteAdvancedClusterFileSystemRequest{}

			deleteAdvancedClusterFileSystemRequest.AdvancedClusterFileSystemId = &advancedClusterFileSystemId

			deleteAdvancedClusterFileSystemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteAdvancedClusterFileSystem(context.Background(), deleteAdvancedClusterFileSystemRequest)
			if error != nil {
				fmt.Printf("Error deleting AdvancedClusterFileSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", advancedClusterFileSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &advancedClusterFileSystemId, DatabaseAdvancedClusterFileSystemSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseAdvancedClusterFileSystemSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseAdvancedClusterFileSystemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AdvancedClusterFileSystemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listAdvancedClusterFileSystemsRequest := oci_database.ListAdvancedClusterFileSystemsRequest{}
	listAdvancedClusterFileSystemsRequest.CompartmentId = &compartmentId
	listAdvancedClusterFileSystemsRequest.LifecycleState = oci_database.AdvancedClusterFileSystemLifecycleStateAvailable
	listAdvancedClusterFileSystemsResponse, err := databaseClient.ListAdvancedClusterFileSystems(context.Background(), listAdvancedClusterFileSystemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AdvancedClusterFileSystem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, advancedClusterFileSystem := range listAdvancedClusterFileSystemsResponse.Items {
		id := *advancedClusterFileSystem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AdvancedClusterFileSystemId", id)
	}
	return resourceIds, nil
}

func DatabaseAdvancedClusterFileSystemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if advancedClusterFileSystemResponse, ok := response.Response.(oci_database.GetAdvancedClusterFileSystemResponse); ok {
		return advancedClusterFileSystemResponse.LifecycleState != oci_database.AdvancedClusterFileSystemLifecycleStateDeleted
	}
	return false
}

func DatabaseAdvancedClusterFileSystemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetAdvancedClusterFileSystem(context.Background(), oci_database.GetAdvancedClusterFileSystemRequest{
		AdvancedClusterFileSystemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
