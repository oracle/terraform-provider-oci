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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CloudBridgeInventoryRequiredOnlyResource = CloudBridgeInventoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeInventoryRepresentation)

	CloudBridgeInventoryResourceConfig = CloudBridgeInventoryResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Optional, acctest.Update, CloudBridgeInventoryRepresentation)

	CloudBridgeCloudBridgeInventorySingularDataSourceRepresentation = map[string]interface{}{
		"inventory_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_cloud_bridge_inventory.test_inventory.id}`},
	}

	CloudBridgeCloudBridgeInventoryDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: CloudBridgeInventoryDataSourceFilterRepresentation}}
	CloudBridgeInventoryDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_cloud_bridge_inventory.test_inventory.id}`}},
	}

	CloudBridgeInventoryRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreSystemTagsChangesRep},
	}

	CloudBridgeInventoryResourceDependencies = ""
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeInventoryResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeInventoryResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_cloud_bridge_inventory.test_inventory"
	datasourceName := "data.oci_cloud_bridge_inventories.test_inventories"
	singularDatasourceName := "data.oci_cloud_bridge_inventory.test_inventory"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CloudBridgeInventoryResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Optional, acctest.Create, CloudBridgeInventoryRepresentation), "cloudbridge", "inventory", t)

	acctest.ResourceTest(t, testAccCheckCloudBridgeInventoryDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeInventoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeInventoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CloudBridgeInventoryResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CloudBridgeInventoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Optional, acctest.Create, CloudBridgeInventoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
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
			Config: config + compartmentIdVariableStr + CloudBridgeInventoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Optional, acctest.Update, CloudBridgeInventoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_inventories", "test_inventories", acctest.Optional, acctest.Update, CloudBridgeCloudBridgeInventoryDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeInventoryResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Optional, acctest.Update, CloudBridgeInventoryRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "inventory_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "inventory_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_inventory", "test_inventory", acctest.Required, acctest.Create, CloudBridgeCloudBridgeInventorySingularDataSourceRepresentation) +
				compartmentIdVariableStr + CloudBridgeInventoryResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "inventory_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + CloudBridgeInventoryRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCloudBridgeInventoryDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).InventoryClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_cloud_bridge_inventory" {
			noResourceFound = false
			request := oci_cloud_bridge.GetInventoryRequest{}

			tmp := rs.Primary.ID
			request.InventoryId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")

			response, err := client.GetInventory(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_cloud_bridge.InventoryLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CloudBridgeInventory") {
		resource.AddTestSweepers("CloudBridgeInventory", &resource.Sweeper{
			Name:         "CloudBridgeInventory",
			Dependencies: acctest.DependencyGraph["inventory"],
			F:            sweepCloudBridgeInventoryResource,
		})
	}
}

func sweepCloudBridgeInventoryResource(compartment string) error {
	inventoryClient := acctest.GetTestClients(&schema.ResourceData{}).InventoryClient()
	inventoryIds, err := getCloudBridgeInventoryIds(compartment)
	if err != nil {
		return err
	}
	for _, inventoryId := range inventoryIds {
		if ok := acctest.SweeperDefaultResourceId[inventoryId]; !ok {
			deleteInventoryRequest := oci_cloud_bridge.DeleteInventoryRequest{}

			deleteInventoryRequest.InventoryId = &inventoryId

			deleteInventoryRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "cloud_bridge")
			_, error := inventoryClient.DeleteInventory(context.Background(), deleteInventoryRequest)
			if error != nil {
				fmt.Printf("Error deleting Inventory %s %s, It is possible that the resource is already deleted. Please verify manually \n", inventoryId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &inventoryId, CloudBridgeInventorySweepWaitCondition, time.Duration(3*time.Minute),
				CloudBridgeInventorySweepResponseFetchOperation, "cloud_bridge", true)
		}
	}
	return nil
}

func getCloudBridgeInventoryIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "InventoryId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	inventoryClient := acctest.GetTestClients(&schema.ResourceData{}).InventoryClient()

	listInventoriesRequest := oci_cloud_bridge.ListInventoriesRequest{}
	listInventoriesRequest.CompartmentId = &compartmentId
	listInventoriesRequest.LifecycleState = oci_cloud_bridge.InventoryLifecycleStateActive
	listInventoriesResponse, err := inventoryClient.ListInventories(context.Background(), listInventoriesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Inventory list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, inventory := range listInventoriesResponse.Items {
		id := *inventory.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "InventoryId", id)
	}
	return resourceIds, nil
}

func CloudBridgeInventorySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if inventoryResponse, ok := response.Response.(oci_cloud_bridge.GetInventoryResponse); ok {
		return inventoryResponse.LifecycleState != oci_cloud_bridge.InventoryLifecycleStateDeleted
	}
	return false
}

func CloudBridgeInventorySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.InventoryClient().GetInventory(context.Background(), oci_cloud_bridge.GetInventoryRequest{
		InventoryId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
