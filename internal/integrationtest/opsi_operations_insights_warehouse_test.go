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
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OpsiOperationsInsightsWarehouseRequiredOnlyResource = OpsiOperationsInsightsWarehouseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation)
	OpsiOperationsInsightsWarehouseResourceConfig = OpsiOperationsInsightsWarehouseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Optional, acctest.Update, OpsiOperationsInsightsWarehouseRepresentation)

	OpsiOpsiOperationsInsightsWarehouseSingularDataSourceRepresentation = map[string]interface{}{
		"operations_insights_warehouse_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
	}

	OpsiOpsiOperationsInsightsWarehouseDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OpsiOperationsInsightsWarehouseDataSourceFilterRepresentation},
	}

	OpsiOperationsInsightsWarehouseDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse.id}`}},
	}

	OpsiOperationsInsightsWarehouseRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_allocated":            acctest.Representation{RepType: acctest.Required, Create: `2.0`, Update: `3.0`},
		"display_name":             acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"compute_model":            acctest.Representation{RepType: acctest.Required, Create: `ECPU`, Update: `ECPU`},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"storage_allocated_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `1.0`, Update: `1.0`},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesOperationsInsightsWarehouseRepresentation},
	}

	ignoreChangesOperationsInsightsWarehouseRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	OpsiOperationsInsightsWarehouseResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: opsi/controlPlane
func TestOpsiOperationsInsightsWarehouseResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOpsiOperationsInsightsWarehouseResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse"
	singularDatasourceName := "data.oci_opsi_operations_insights_warehouse.test_operations_insights_warehouse"

	var resId, resId2 string

	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OpsiOperationsInsightsWarehouseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Optional, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation), "operationsinsights", "operationsInsightsWarehouse", t)

	acctest.ResourceTest(t, testAccCheckOpsiOperationsInsightsWarehouseDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next create
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Optional, acctest.Create, OpsiOperationsInsightsWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_allocated_in_gbs", "1"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OpsiOperationsInsightsWarehouseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Optional, acctest.Update, OpsiOperationsInsightsWarehouseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "cpu_allocated", "3"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "storage_allocated_in_gbs", "1"),
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
		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + OpsiOperationsInsightsWarehouseResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_opsi_operations_insights_warehouse", "test_operations_insights_warehouse", acctest.Required, acctest.Create, OpsiOpsiOperationsInsightsWarehouseSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "operations_insights_warehouse_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cpu_allocated", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "storage_allocated_in_gbs", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_wallet_rotated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + OpsiOperationsInsightsWarehouseRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOpsiOperationsInsightsWarehouseDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).OperationsInsightsClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_opsi_operations_insights_warehouse" {
			noResourceFound = false
			request := oci_opsi.GetOperationsInsightsWarehouseRequest{}
			tmp := rs.Primary.ID
			request.OperationsInsightsWarehouseId = &tmp
			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}

	if !acctest.InSweeperExcludeList("OpsiOperationsInsightsWarehouse") {
		resource.AddTestSweepers("OpsiOperationsInsightsWarehouse", &resource.Sweeper{
			Name:         "OpsiOperationsInsightsWarehouse",
			Dependencies: acctest.DependencyGraph["operationsInsightsWarehouse"],
			F:            sweepOpsiOperationsInsightsWarehouseResource,
		})
	}
}

func sweepOpsiOperationsInsightsWarehouseResource(compartment string) error {
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()
	operationsInsightsWarehouseIds, err := getOpsiOperationsInsightsWarehouseIds(compartment)
	if err != nil {
		return err
	}

	for _, operationsInsightsWarehouseId := range operationsInsightsWarehouseIds {
		if ok := acctest.SweeperDefaultResourceId[operationsInsightsWarehouseId]; !ok {
			deleteOperationsInsightsWarehouseRequest := oci_opsi.DeleteOperationsInsightsWarehouseRequest{}
			deleteOperationsInsightsWarehouseRequest.OperationsInsightsWarehouseId = &operationsInsightsWarehouseId
			deleteOperationsInsightsWarehouseRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "opsi")

			_, error := operationsInsightsClient.DeleteOperationsInsightsWarehouse(context.Background(), deleteOperationsInsightsWarehouseRequest)
			if error != nil {
				fmt.Printf("Error deleting OperationsInsightsWarehouse %s %s, It is possible that the resource is already deleted. Please verify manually \n", operationsInsightsWarehouseId, error)
				continue
			}

			acctest.WaitTillCondition(acctest.TestAccProvider, &operationsInsightsWarehouseId, OpsiOperationsInsightsWarehouseSweepWaitCondition, time.Duration(3*time.Minute), OpsiOperationsInsightsWarehouseSweepResponseFetchOperation, "opsi", true)
		}
	}

	return nil
}

func getOpsiOperationsInsightsWarehouseIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "OperationsInsightsWarehouseId")
	if ids != nil {
		return ids, nil
	}

	var resourceIds []string
	compartmentId := compartment
	operationsInsightsClient := acctest.GetTestClients(&schema.ResourceData{}).OperationsInsightsClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "OperationsInsightsWarehouseId", id)
	}

	return resourceIds, nil
}

func OpsiOperationsInsightsWarehouseSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if operationsInsightsWarehouseResponse, ok := response.Response.(oci_opsi.GetOperationsInsightsWarehouseResponse); ok {
		return operationsInsightsWarehouseResponse.LifecycleState != oci_opsi.OperationsInsightsWarehouseLifecycleStateDeleted
	}

	return false
}

func OpsiOperationsInsightsWarehouseSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.OperationsInsightsClient().GetOperationsInsightsWarehouse(context.Background(), oci_opsi.GetOperationsInsightsWarehouseRequest{
		OperationsInsightsWarehouseId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})

	return err
}
