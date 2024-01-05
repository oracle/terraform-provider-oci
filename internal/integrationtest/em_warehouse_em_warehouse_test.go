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
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	EmWarehouseRequiredOnlyResource = EmWarehouseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseRepresentation)

	EmWarehouseResourceConfig = EmWarehouseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Update, emWarehouseRepresentation)

	emWarehouseSingularDataSourceRepresentation = map[string]interface{}{
		"em_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`},
	}

	emWarehouseDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"id":                               acctest.Representation{RepType: acctest.Optional, Create: `${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`},
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.operations_insights_warehouse_id}`},
		"state":                            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                           acctest.RepresentationGroup{RepType: acctest.Required, Group: emWarehouseDataSourceFilterRepresentation}}
	emWarehouseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_em_warehouse_em_warehouse.test_em_warehouse.id}`}},
	}

	emWarehouseRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"em_bridge_id":                     acctest.Representation{RepType: acctest.Required, Create: `${var.em_bridge_id}`},
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${var.operations_insights_warehouse_id}`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesEmWarehouseRepresentation},
	}

	ignoreChangesEmWarehouseRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	EmWarehouseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: em_warehouse/default
func TestEmWarehouseEmWarehouseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestEmWarehouseEmWarehouseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	operationsInsightWarehouseId := utils.GetEnvSettingWithBlankDefault("operations_insights_warehouse_id")
	operationsInsightWarehouseIdVariableStr := fmt.Sprintf("variable \"operations_insights_warehouse_id\" { default = \"%s\" }\n", operationsInsightWarehouseId)

	emBridgeId := utils.GetEnvSettingWithBlankDefault("em_bridge_id")
	emBridgeIdVariableStr := fmt.Sprintf("variable \"em_bridge_id\" { default = \"%s\" }\n", emBridgeId)

	resourceName := "oci_em_warehouse_em_warehouse.test_em_warehouse"
	datasourceName := "data.oci_em_warehouse_em_warehouses.test_em_warehouses"
	singularDatasourceName := "data.oci_em_warehouse_em_warehouse.test_em_warehouse"

	c := config + compartmentIdVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseRepresentation)
	print(c)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+EmWarehouseResourceDependencies+operationsInsightWarehouseIdVariableStr+emBridgeIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Create, emWarehouseRepresentation), "emwarehouse", "emWarehouse", t)

	acctest.ResourceTest(t, testAccCheckEmWarehouseEmWarehouseDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "em_bridge_id"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + EmWarehouseResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Create, emWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "em_bridge_id"),
				resource.TestCheckResourceAttrSet(resourceName, "em_warehouse_type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(emWarehouseRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "em_bridge_id"),
				resource.TestCheckResourceAttrSet(resourceName, "em_warehouse_type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Update, emWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "em_bridge_id"),
				resource.TestCheckResourceAttrSet(resourceName, "em_warehouse_type"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_em_warehouse_em_warehouses", "test_em_warehouses", acctest.Optional, acctest.Update, emWarehouseDataSourceRepresentation) +
				compartmentIdVariableStr + EmWarehouseResourceDependencies + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Optional, acctest.Update, emWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"), // This is hash generated value. so we cannot match exact value
				resource.TestCheckResourceAttrSet(datasourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "em_warehouse_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "em_warehouse_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_em_warehouse_em_warehouse", "test_em_warehouse", acctest.Required, acctest.Create, emWarehouseSingularDataSourceRepresentation) +
				compartmentIdVariableStr + operationsInsightWarehouseIdVariableStr + emBridgeIdVariableStr + EmWarehouseResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "em_warehouse_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "em_warehouse_type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + EmWarehouseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckEmWarehouseEmWarehouseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).EmWarehouseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_em_warehouse_em_warehouse" {
			noResourceFound = false
			request := oci_em_warehouse.GetEmWarehouseRequest{}

			tmp := rs.Primary.ID
			request.EmWarehouseId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "em_warehouse")

			response, err := client.GetEmWarehouse(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_em_warehouse.EmWarehouseLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("EmWarehouseEmWarehouse") {
		resource.AddTestSweepers("EmWarehouseEmWarehouse", &resource.Sweeper{
			Name:         "EmWarehouseEmWarehouse",
			Dependencies: acctest.DependencyGraph["emWarehouse"],
			F:            sweepEmWarehouseEmWarehouseResource,
		})
	}
}

func sweepEmWarehouseEmWarehouseResource(compartment string) error {
	emWarehouseClient := acctest.GetTestClients(&schema.ResourceData{}).EmWarehouseClient()
	emWarehouseIds, err := getEmWarehouseIds(compartment)
	if err != nil {
		return err
	}
	for _, emWarehouseId := range emWarehouseIds {
		if ok := acctest.SweeperDefaultResourceId[emWarehouseId]; !ok {
			deleteEmWarehouseRequest := oci_em_warehouse.DeleteEmWarehouseRequest{}

			deleteEmWarehouseRequest.EmWarehouseId = &emWarehouseId

			deleteEmWarehouseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "em_warehouse")
			_, error := emWarehouseClient.DeleteEmWarehouse(context.Background(), deleteEmWarehouseRequest)
			if error != nil {
				fmt.Printf("Error deleting EmWarehouse %s %s, It is possible that the resource is already deleted. Please verify manually \n", emWarehouseId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &emWarehouseId, emWarehouseSweepWaitCondition, time.Duration(3*time.Minute),
				emWarehouseSweepResponseFetchOperation, "em_warehouse", true)
		}
	}
	return nil
}

func getEmWarehouseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "EmWarehouseId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	emWarehouseClient := acctest.GetTestClients(&schema.ResourceData{}).EmWarehouseClient()

	listEmWarehousesRequest := oci_em_warehouse.ListEmWarehousesRequest{}
	listEmWarehousesRequest.CompartmentId = &compartmentId
	listEmWarehousesRequest.LifecycleState = oci_em_warehouse.EmWarehouseLifecycleStateActive
	listEmWarehousesResponse, err := emWarehouseClient.ListEmWarehouses(context.Background(), listEmWarehousesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting EmWarehouse list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, emWarehouse := range listEmWarehousesResponse.Items {
		id := *emWarehouse.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "EmWarehouseId", id)
	}
	return resourceIds, nil
}

func emWarehouseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if emWarehouseResponse, ok := response.Response.(oci_em_warehouse.GetEmWarehouseResponse); ok {
		return emWarehouseResponse.LifecycleState != oci_em_warehouse.EmWarehouseLifecycleStateDeleted
	}
	return false
}

func emWarehouseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.EmWarehouseClient().GetEmWarehouse(context.Background(), oci_em_warehouse.GetEmWarehouseRequest{
		EmWarehouseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
