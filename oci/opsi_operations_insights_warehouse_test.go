// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v53/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v53/opsi"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	OperationsInsightsWarehouseRequiredOnlyResource = OperationsInsightsWarehouseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation)

	OperationsInsightsWarehouseResourceConfig = OperationsInsightsWarehouseResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Optional, Update, operationsInsightsWarehouseRepresentation)

	operationsInsightsWarehouseSingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": Representation{RepType: Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
	}

	operationsInsightsWarehouseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"display_name":   Representation{RepType: Optional, Create: `displayName`, Update: `displayName2`},
		"id":             Representation{RepType: Optional, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"state":          Representation{RepType: Optional, Create: []string{`ACTIVE`}},
		"filter":         RepresentationGroup{Required, operationsInsightsWarehouseDataSourceFilterRepresentation}}
	operationsInsightsWarehouseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `id`},
		"values": Representation{RepType: Required, Create: []string{`${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`}},
	}

	operationsInsightsWarehouseRepresentation = map[string]interface{}{
		"compartment_id":           Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_allocated":            Representation{RepType: Required, Create: `1.0`, Update: `2.0`},
		"display_name":             Representation{RepType: Required, Create: `displayName`, Update: `displayName2`},
		"defined_tags":             Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":            Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"storage_allocated_in_gbs": Representation{RepType: Optional, Create: `1.0`, Update: `1.0`},
		"lifecycle":                RepresentationGroup{Required, ignoreChangesOperationsInsightsWarehouseRepresentation},
	}

	ignoreChangesOperationsInsightsWarehouseRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}

	OperationsInsightsWarehouseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse"
	datasourceName := "data.oci_opsi_operations_insights_warehouses.test_operations_insights_warehouses"
	singularDatasourceName := "data.oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+OperationsInsightsWarehouseResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Optional, Create, operationsInsightsWarehouseRepresentation), "operationsinsights", "operationsInsightsWarehouse", t)

	ResourceTest(t, testAccCheckOpsiOperationsInsightsWarehouseDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Optional, Create, operationsInsightsWarehouseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "1"),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_allocated_in_gbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Optional, Update, operationsInsightsWarehouseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "2"),
				//resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_allocated_in_gbs", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouses", "test_operations_insights_warehouses", Optional, Update, operationsInsightsWarehouseDataSourceRepresentation) +
				compartmentIdVariableStr + OperationsInsightsWarehouseResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Optional, Update, operationsInsightsWarehouseRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "operations_insights_warehouse_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "operations_insights_warehouse_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", Required, Create, operationsInsightsWarehouseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OperationsInsightsWarehouseResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_allocated", "2"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_used"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "dynamic_group_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_tenancy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_allocated_in_gbs", "1"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "storage_used_in_gbs"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_wallet_rotated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + OperationsInsightsWarehouseResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiOperationsInsightsWarehouseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).operationsInsightsClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_operations_insights_warehouse" {
			noResourceFound = false
			request := oci_opsi.GetOperationsInsightsWarehouseRequest{}

			tmp := rs.Primary.ID
			request.OperationsInsightsWarehouseId = &tmp

			request.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")

			response, err := client.GetOperationsInsightsWarehouse(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_opsi.OperationsInsightsWarehouseLifecycleStateDeleted): true,
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
	if !InSweeperExcludeList("OpsiOperationsInsightsWarehouse") {
		resource.AddTestSweepers("OpsiOperationsInsightsWarehouse", &resource.Sweeper{
			Name:         "OpsiOperationsInsightsWarehouse",
			Dependencies: DependencyGraph["operationsInsightsWarehouse"],
			F:            sweepOpsiOperationsInsightsWarehouseResource,
		})
	}
}

func sweepOpsiOperationsInsightsWarehouseResource(compartment string) error {
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()
	operationsInsightsWarehouseIds, err := getOperationsInsightsWarehouseIds(compartment)
	if err != nil {
		return err
	}
	for _, operationsInsightsWarehouseId := range operationsInsightsWarehouseIds {
		if ok := SweeperDefaultResourceId[operationsInsightsWarehouseId]; !ok {
			deleteOperationsInsightsWarehouseRequest := oci_opsi.DeleteOperationsInsightsWarehouseRequest{}

			deleteOperationsInsightsWarehouseRequest.OperationsInsightsWarehouseId = &operationsInsightsWarehouseId

			deleteOperationsInsightsWarehouseRequest.RequestMetadata.RetryPolicy = GetRetryPolicy(true, "opsi")
			_, error := operationsInsightsClient.DeleteOperationsInsightsWarehouse(context.Background(), deleteOperationsInsightsWarehouseRequest)
			if error != nil {
				fmt.Printf("Error deleting OperationsInsightsWarehouse %s %s, It is possible that the resource is already deleted. Please verify manually \n", operationsInsightsWarehouseId, error)
				continue
			}
			WaitTillCondition(testAccProvider, &operationsInsightsWarehouseId, operationsInsightsWarehouseSweepWaitCondition, time.Duration(3*time.Minute),
				operationsInsightsWarehouseSweepResponseFetchOperation, "opsi", true)
		}
	}
	return nil
}

func getOperationsInsightsWarehouseIds(compartment string) ([]string, error) {
	ids := GetResourceIdsToSweep(compartment, "OperationsInsightsWarehouseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := GetTestClients(&schema.ResourceData{}).operationsInsightsClient()

	listOperationsInsightsWarehousesRequest := oci_opsi.ListOperationsInsightsWarehousesRequest{}
	listOperationsInsightsWarehousesRequest.CompartmentId = &compartmentId
	listOperationsInsightsWarehousesRequest.LifecycleState = []oci_opsi.OperationsInsightsWarehouseLifecycleStateEnum{oci_opsi.OperationsInsightsWarehouseLifecycleStateActive}
	listOperationsInsightsWarehousesResponse, err := operationsInsightsClient.ListOperationsInsightsWarehouses(context.Background(), listOperationsInsightsWarehousesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting OperationsInsightsWarehouse list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, operationsInsightsWarehouse := range listOperationsInsightsWarehousesResponse.Items {
		id := *operationsInsightsWarehouse.Id
		resourceIds = append(resourceIds, id)
		AddResourceIdToSweeperResourceIdMap(compartmentId, "OperationsInsightsWarehouseId", id)
	}
	return resourceIds, nil
}

func operationsInsightsWarehouseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operationsInsightsWarehouseResponse, ok := response.Response.(oci_opsi.GetOperationsInsightsWarehouseResponse); ok {
		return operationsInsightsWarehouseResponse.LifecycleState != oci_opsi.OperationsInsightsWarehouseLifecycleStateDeleted
	}
	return false
}

func operationsInsightsWarehouseSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.operationsInsightsClient().GetOperationsInsightsWarehouse(context.Background(), oci_opsi.GetOperationsInsightsWarehouseRequest{
		OperationsInsightsWarehouseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
