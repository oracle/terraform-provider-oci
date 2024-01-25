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
	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseManagementExternalDbSystemDiscoveryRequiredOnlyResource = DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryRepresentation)

	DatabaseManagementExternalDbSystemDiscoveryResourceConfig = DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDiscoveryRepresentation)

	DatabaseManagementDatabaseManagementExternalDbSystemDiscoverySingularDataSourceRepresentation = map[string]interface{}{
		"external_db_system_discovery_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_external_db_system_discovery.test_external_db_system_discovery.id}`},
	}

	DatabaseManagementDatabaseManagementExternalDbSystemDiscoveryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `EXAMPLE-displayName-Value`, Update: `displayName2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementExternalDbSystemDiscoveryDataSourceFilterRepresentation}}
	DatabaseManagementExternalDbSystemDiscoveryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_management_external_db_system_discovery.test_external_db_system_discovery.id}`}},
	}

	DatabaseManagementExternalDbSystemDiscoveryRepresentation = map[string]interface{}{
		"agent_id":       acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `discoveryDisplayName`, Update: `displayName2`},
	}

	DatabaseManagementExternalDbSystemDiscoveryResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementExternalDbSystemDiscoveryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementExternalDbSystemDiscoveryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_database_management_external_db_system_discovery.test_external_db_system_discovery"
	datasourceName := "data.oci_database_management_external_db_system_discoveries.test_external_db_system_discoveries"
	singularDatasourceName := "data.oci_database_management_external_db_system_discovery.test_external_db_system_discovery"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseManagementExternalDbSystemDiscoveryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryRepresentation), "databasemanagement", "externalDbSystemDiscovery", t)

	acctest.ResourceTest(t, testAccCheckDatabaseManagementExternalDbSystemDiscoveryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Optional, acctest.Create, DatabaseManagementExternalDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "discoveryDisplayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "agent_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system_discoveries", "test_external_db_system_discoveries", acctest.Optional, acctest.Update, DatabaseManagementDatabaseManagementExternalDbSystemDiscoveryDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Optional, acctest.Update, DatabaseManagementExternalDbSystemDiscoveryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),

				resource.TestCheckResourceAttr(datasourceName, "external_db_system_discovery_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "external_db_system_discovery_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_external_db_system_discovery", "test_external_db_system_discovery", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementExternalDbSystemDiscoverySingularDataSourceRepresentation) +
				compartmentIdVariableStr + agentIdVariableStr + DatabaseManagementExternalDbSystemDiscoveryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "discovered_components.#"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "grid_home"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseManagementExternalDbSystemDiscoveryRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"patch_operations",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseManagementExternalDbSystemDiscoveryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbManagementClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_management_external_db_system_discovery" {
			noResourceFound = false
			request := oci_database_management.GetExternalDbSystemDiscoveryRequest{}

			tmp := rs.Primary.ID
			request.ExternalDbSystemDiscoveryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")

			response, err := client.GetExternalDbSystemDiscovery(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database_management.ExternalDbSystemDiscoveryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("DatabaseManagementExternalDbSystemDiscovery") {
		resource.AddTestSweepers("DatabaseManagementExternalDbSystemDiscovery", &resource.Sweeper{
			Name:         "DatabaseManagementExternalDbSystemDiscovery",
			Dependencies: acctest.DependencyGraph["externalDbSystemDiscovery"],
			F:            sweepDatabaseManagementExternalDbSystemDiscoveryResource,
		})
	}
}

func sweepDatabaseManagementExternalDbSystemDiscoveryResource(compartment string) error {
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()
	externalDbSystemDiscoveryIds, err := getDatabaseManagementExternalDbSystemDiscoveryIds(compartment)
	if err != nil {
		return err
	}
	for _, externalDbSystemDiscoveryId := range externalDbSystemDiscoveryIds {
		if ok := acctest.SweeperDefaultResourceId[externalDbSystemDiscoveryId]; !ok {
			deleteExternalDbSystemDiscoveryRequest := oci_database_management.DeleteExternalDbSystemDiscoveryRequest{}

			deleteExternalDbSystemDiscoveryRequest.ExternalDbSystemDiscoveryId = &externalDbSystemDiscoveryId

			deleteExternalDbSystemDiscoveryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database_management")
			_, error := dbManagementClient.DeleteExternalDbSystemDiscovery(context.Background(), deleteExternalDbSystemDiscoveryRequest)
			if error != nil {
				fmt.Printf("Error deleting ExternalDbSystemDiscovery %s %s, It is possible that the resource is already deleted. Please verify manually \n", externalDbSystemDiscoveryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &externalDbSystemDiscoveryId, DatabaseManagementExternalDbSystemDiscoverySweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseManagementExternalDbSystemDiscoverySweepResponseFetchOperation, "database_management", true)
		}
	}
	return nil
}

func getDatabaseManagementExternalDbSystemDiscoveryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ExternalDbSystemDiscoveryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dbManagementClient := acctest.GetTestClients(&schema.ResourceData{}).DbManagementClient()

	listExternalDbSystemDiscoveriesRequest := oci_database_management.ListExternalDbSystemDiscoveriesRequest{}
	listExternalDbSystemDiscoveriesRequest.CompartmentId = &compartmentId
	listExternalDbSystemDiscoveriesResponse, err := dbManagementClient.ListExternalDbSystemDiscoveries(context.Background(), listExternalDbSystemDiscoveriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ExternalDbSystemDiscovery list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, externalDbSystemDiscovery := range listExternalDbSystemDiscoveriesResponse.Items {
		id := *externalDbSystemDiscovery.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ExternalDbSystemDiscoveryId", id)
	}
	return resourceIds, nil
}

func DatabaseManagementExternalDbSystemDiscoverySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if externalDbSystemDiscoveryResponse, ok := response.Response.(oci_database_management.GetExternalDbSystemDiscoveryResponse); ok {
		return externalDbSystemDiscoveryResponse.LifecycleState != oci_database_management.ExternalDbSystemDiscoveryLifecycleStateDeleted
	}
	return false
}

func DatabaseManagementExternalDbSystemDiscoverySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbManagementClient().GetExternalDbSystemDiscovery(context.Background(), oci_database_management.GetExternalDbSystemDiscoveryRequest{
		ExternalDbSystemDiscoveryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
