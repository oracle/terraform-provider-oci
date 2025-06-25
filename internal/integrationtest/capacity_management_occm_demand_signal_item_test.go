// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CapacityManagementOccmDemandSignalItemRequiredOnlyResource = CapacityManagementOccmDemandSignalItemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalItemRepresentation)

	CapacityManagementOccmDemandSignalItemResourceConfig = CapacityManagementOccmDemandSignalItemResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalItemRepresentation)

	CapacityManagementOccmDemandSignalItemSingularDataSourceRepresentation = map[string]interface{}{
		"occm_demand_signal_item_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item.id}`},
	}
	//CapacityManagementOccmDemandSignalItemDataSourceRepresentation = map[string]interface{}{
	//	"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	//	"demand_signal_namespace": acctest.Representation{RepType: acctest.Optional, Create: `COMPUTE`},
	//	"occm_demand_signal_id":   acctest.Representation{RepType: acctest.Optional, Create: `${oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id}`},
	//	"resource_name":           acctest.Representation{RepType: acctest.Optional, Create: `BM.Standard.E5.192`},
	//	"filter":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: CapacityManagementOccmDemandSignalItemDataSourceFilterRepresentation}}
	//CapacityManagementOccmDemandSignalItemDataSourceFilterRepresentation = map[string]interface{}{
	//	"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
	//	"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item.id}`}},
	//}

	CapacityManagementOccmDemandSignalItemRepresentation = map[string]interface{}{
		"compartment_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"demand_quantity":                   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"demand_signal_catalog_resource_id": acctest.Representation{RepType: acctest.Required, Create: `${var.catalogresource_id}`},
		"demand_signal_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id}`},
		"region":                            acctest.Representation{RepType: acctest.Required, Create: `region`, Update: `region2`},
		"request_type":                      acctest.Representation{RepType: acctest.Required, Create: `DEMAND`},
		"resource_properties":               acctest.Representation{RepType: acctest.Required, Create: map[string]string{"resourceProperties": "resourceProperties"}, Update: map[string]string{"resourceProperties2": "resourceProperties2"}},
		"time_needed_before":                acctest.Representation{RepType: acctest.Required, Create: `2025-06-01T00:00:00Z`, Update: `2025-06-02T00:00:00Z`},
		"notes":                             acctest.Representation{RepType: acctest.Optional, Create: `notes`, Update: `notes2`},
	}

	CapacityManagementOccmDemandSignalItemResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal", "test_occm_demand_signal", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalRepresentation) +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

// issue-routing-tag: capacity_management/default
func TestCapacityManagementOccmDemandSignalItemResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCapacityManagementOccmDemandSignalItemResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("sp_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	catalogresourceId := utils.GetEnvSettingWithBlankDefault("demand_signal_catalog_resource_id")
	catalogresourceIdVariableStr := fmt.Sprintf("variable \"catalogresource_id\" { default = \"%s\" }\n", catalogresourceId)

	resourceName := "oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item"
	// datasourceName := "data.oci_capacity_management_occm_demand_signal_items.test_occm_demand_signal_items"
	singularDatasourceName := "data.oci_capacity_management_occm_demand_signal_item.test_occm_demand_signal_item"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+catalogresourceIdVariableStr+CapacityManagementOccmDemandSignalItemResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Optional, acctest.Create, CapacityManagementOccmDemandSignalItemRepresentation), "capacitymanagement", "occmDemandSignalItem", t)

	acctest.ResourceTest(t, testAccCheckCapacityManagementOccmDemandSignalItemDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + catalogresourceIdVariableStr + CapacityManagementOccmDemandSignalItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "demand_quantity", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_catalog_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttr(resourceName, "region", "region"),
				resource.TestCheckResourceAttr(resourceName, "request_type", "DEMAND"),
				resource.TestCheckResourceAttr(resourceName, "resource_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "time_needed_before", "2025-06-01T00:00:00Z"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + catalogresourceIdVariableStr + CapacityManagementOccmDemandSignalItemResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + catalogresourceIdVariableStr + CapacityManagementOccmDemandSignalItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Optional, acctest.Create, CapacityManagementOccmDemandSignalItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "demand_quantity", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_catalog_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notes", "notes"),
				resource.TestCheckResourceAttr(resourceName, "region", "region"),
				resource.TestCheckResourceAttr(resourceName, "request_type", "DEMAND"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_properties.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_needed_before", "2025-06-01T00:00:00Z"),

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
			Config: config + compartmentIdVariableStr + catalogresourceIdVariableStr + CapacityManagementOccmDemandSignalItemResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalItemRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "demand_quantity", "11"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_catalog_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_id"),
				resource.TestCheckResourceAttrSet(resourceName, "demand_signal_namespace"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "notes", "notes2"),
				resource.TestCheckResourceAttr(resourceName, "region", "region2"),
				resource.TestCheckResourceAttr(resourceName, "request_type", "DEMAND"),
				resource.TestCheckResourceAttrSet(resourceName, "resource_name"),
				resource.TestCheckResourceAttr(resourceName, "resource_properties.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "time_needed_before", "2025-06-02T00:00:00Z"),
				resource.TestCheckResourceAttrSet("oci_capacity_management_occm_demand_signal.test_occm_demand_signal", "id"), // Verify demand signal exists
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					log.Printf("[DEBUG] Resource state: %+v", s.RootModule().Resources[resourceName].Primary.Attributes)
					return err
				},
			),
		},
		// verify datasource
		//{
		//	Config: config +
		//		acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_items", "test_occm_demand_signal_items", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalItemDataSourceRepresentation) +
		//		compartmentIdVariableStr + CapacityManagementOccmDemandSignalItemResourceDependencies +
		//		acctest.GenerateResourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Optional, acctest.Update, CapacityManagementOccmDemandSignalItemRepresentation),
		//	Check: acctest.ComposeAggregateTestCheckFuncWrapper(
		//		resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
		//		resource.TestCheckResourceAttr(datasourceName, "demand_signal_namespace", "COMPUTE"),
		//		resource.TestCheckResourceAttrSet(datasourceName, "occm_demand_signal_id"),
		//		resource.TestCheckResourceAttrSet(datasourceName, "resource_name"),
		//		resource.TestCheckResourceAttr(datasourceName, "occm_demand_signal_item_collection.#", "1"),
		//	),
		//},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_capacity_management_occm_demand_signal_item", "test_occm_demand_signal_item", acctest.Required, acctest.Create, CapacityManagementOccmDemandSignalItemSingularDataSourceRepresentation) +
				compartmentIdVariableStr + catalogresourceIdVariableStr + CapacityManagementOccmDemandSignalItemResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "occm_demand_signal_item_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "demand_quantity", "11"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "demand_signal_namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "notes", "notes2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "region", "region2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "request_type", "DEMAND"),
				resource.TestCheckResourceAttr(singularDatasourceName, "resource_properties.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_needed_before"),
			),
		},
		// verify resource import
		{
			Config:                  config + CapacityManagementOccmDemandSignalItemRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckCapacityManagementOccmDemandSignalItemDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DemandSignalClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_capacity_management_occm_demand_signal_item" {
			noResourceFound = false
			request := oci_capacity_management.GetOccmDemandSignalItemRequest{}

			tmp := rs.Primary.ID
			request.OccmDemandSignalItemId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")

			response, err := client.GetOccmDemandSignalItem(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("CapacityManagementOccmDemandSignalItem") {
		resource.AddTestSweepers("CapacityManagementOccmDemandSignalItem", &resource.Sweeper{
			Name:         "CapacityManagementOccmDemandSignalItem",
			Dependencies: acctest.DependencyGraph["occmDemandSignalItem"],
			F:            sweepCapacityManagementOccmDemandSignalItemResource,
		})
	}
}

func sweepCapacityManagementOccmDemandSignalItemResource(compartment string) error {
	demandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).DemandSignalClient()
	occmDemandSignalItemIds, err := getCapacityManagementOccmDemandSignalItemIds(compartment)
	if err != nil {
		return err
	}
	for _, occmDemandSignalItemId := range occmDemandSignalItemIds {
		if ok := acctest.SweeperDefaultResourceId[occmDemandSignalItemId]; !ok {
			deleteOccmDemandSignalItemRequest := oci_capacity_management.DeleteOccmDemandSignalItemRequest{}

			deleteOccmDemandSignalItemRequest.OccmDemandSignalItemId = &occmDemandSignalItemId

			deleteOccmDemandSignalItemRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "capacity_management")
			_, error := demandSignalClient.DeleteOccmDemandSignalItem(context.Background(), deleteOccmDemandSignalItemRequest)
			if error != nil {
				fmt.Printf("Error deleting OccmDemandSignalItem %s %s, It is possible that the resource is already deleted. Please verify manually \n", occmDemandSignalItemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &occmDemandSignalItemId, CapacityManagementOccmDemandSignalItemSweepWaitCondition, time.Duration(3*time.Minute),
				CapacityManagementOccmDemandSignalItemSweepResponseFetchOperation, "capacity_management", true)
		}
	}
	return nil
}

func getCapacityManagementOccmDemandSignalItemIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OccmDemandSignalItemId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	demandSignalClient := acctest.GetTestClients(&schema.ResourceData{}).DemandSignalClient()

	listOccmDemandSignalItemsRequest := oci_capacity_management.ListOccmDemandSignalItemsRequest{}
	listOccmDemandSignalItemsRequest.CompartmentId = &compartmentId
	listOccmDemandSignalItemsResponse, err := demandSignalClient.ListOccmDemandSignalItems(context.Background(), listOccmDemandSignalItemsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OccmDemandSignalItem list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, occmDemandSignalItem := range listOccmDemandSignalItemsResponse.Items {
		id := *occmDemandSignalItem.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OccmDemandSignalItemId", id)
	}
	return resourceIds, nil
}

func CapacityManagementOccmDemandSignalItemSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if occmDemandSignalItemResponse, ok := response.Response.(oci_capacity_management.GetOccmDemandSignalItemResponse); ok {
		return occmDemandSignalItemResponse.LifecycleState != oci_capacity_management.OccmDemandSignalItemLifecycleStateDeleted
	}
	return false
}

func CapacityManagementOccmDemandSignalItemSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DemandSignalClient().GetOccmDemandSignalItem(context.Background(), oci_capacity_management.GetOccmDemandSignalItemRequest{
		OccmDemandSignalItemId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
