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
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	tenantId                                     = utils.GetEnvSettingWithBlankDefault("tenant_ocid")
	BudgetCostAnomalyMonitorRequiredOnlyResource = BudgetCostAnomalyMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorRepresentation)

	BudgetCostAnomalyMonitorResourceConfig = BudgetCostAnomalyMonitorResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorRepresentation)

	BudgetCostAnomalyMonitorSingularDataSourceRepresentation = map[string]interface{}{
		"cost_anomaly_monitor_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor.id}`},
	}

	BudgetCostAnomalyMonitorDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: BudgetCostAnomalyMonitorDataSourceFilterRepresentation}}
	BudgetCostAnomalyMonitorDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor.id}`}},
	}

	BudgetCostAnomalyMonitorRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                        acctest.Representation{RepType: acctest.Required, Create: `name`},
		"target_resource_filter":      acctest.Representation{RepType: acctest.Required, Create: fmt.Sprintf(`{\"operator\": \"AND\", \"dimensions\": [], \"tags\": [], \"filters\": [{\"operator\": \"AND\", \"dimensions\": [{\"key\": \"tenantId\", \"value\": \"%s\"}], \"tags\": [], \"filters\": []}, {\"operator\": \"AND\", \"dimensions\": [{\"key\": \"region\", \"value\": \"us-phoenix-1\"}], \"tags\": [], \"filters\": []}, {\"operator\": \"AND\", \"dimensions\": [{\"key\": \"service\", \"value\": \"COMPUTE\"}], \"tags\": [], \"filters\": []}]}`, tenantId)},
		"cost_alert_subscription_map": acctest.RepresentationGroup{RepType: acctest.Required, Group: BudgetCostAnomalyMonitorCostAlertSubscriptionMapRepresentation},
		"description":                 acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
	BudgetCostAnomalyMonitorCostAlertSubscriptionMapRepresentation = map[string]interface{}{
		"cost_alert_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_cost_alert_subscription.test_cost_alert_subscription.id}`},
		"operator":                   acctest.Representation{RepType: acctest.Required, Create: `AND`, Update: `OR`},
		"threshold_absolute_value":   acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"threshold_relative_percent": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	BudgetCostAnomalyMonitorResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, BudgetCostAlertSubscriptionRepresentation)
	//+ DefinedTagsDependencies
)

// issue-routing-tag: budget/default
func TestBudgetCostAnomalyMonitorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetCostAnomalyMonitorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor"
	datasourceName := "data.oci_budget_cost_anomaly_monitors.test_cost_anomaly_monitors"
	singularDatasourceName := "data.oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+BudgetCostAnomalyMonitorResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Optional, acctest.Create, BudgetCostAnomalyMonitorRepresentation), "budget", "costAnomalyMonitor", t)

	acctest.ResourceTest(t, testAccCheckBudgetCostAnomalyMonitorDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_filter", fmt.Sprintf(`{"operator":"AND","dimensions":[],"tags":[],"filters":[{"operator":"AND","dimensions":[{"key":"tenantId","value":"%s"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"region","value":"us-phoenix-1"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"service","value":"COMPUTE"}],"tags":[],"filters":[]}]}`, tenantId)),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cost_alert_subscription_map.0.cost_alert_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.operator", "AND"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Optional, acctest.Create, BudgetCostAnomalyMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cost_alert_subscription_map.0.cost_alert_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.operator", "AND"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.threshold_absolute_value", "10"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.threshold_relative_percent", "10"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_filter", fmt.Sprintf(`{"operator":"AND","dimensions":[],"tags":[],"filters":[{"operator":"AND","dimensions":[{"key":"tenantId","value":"%s"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"region","value":"us-phoenix-1"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"service","value":"COMPUTE"}],"tags":[],"filters":[]}]}`, tenantId)),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					//if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
					//	if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
					//		return errExport
					//	}
					//}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "cost_alert_subscription_map.0.cost_alert_subscription_id"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.operator", "OR"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.threshold_absolute_value", "11"),
				resource.TestCheckResourceAttr(resourceName, "cost_alert_subscription_map.0.threshold_relative_percent", "11"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "target_resource_filter", fmt.Sprintf(`{"operator":"AND","dimensions":[],"tags":[],"filters":[{"operator":"AND","dimensions":[{"key":"tenantId","value":"%s"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"region","value":"us-phoenix-1"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"service","value":"COMPUTE"}],"tags":[],"filters":[]}]}`, tenantId)),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_monitors", "test_cost_anomaly_monitors", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "name", "name"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "cost_anomaly_monitor_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "cost_anomaly_monitor_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + BudgetCostAnomalyMonitorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_anomaly_monitor_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "cost_alert_subscription_map.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cost_alert_subscription_map.0.operator", "OR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cost_alert_subscription_map.0.threshold_absolute_value", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cost_alert_subscription_map.0.threshold_relative_percent", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "target_resource_filter", fmt.Sprintf(`{"operator":"AND","dimensions":[],"tags":[],"filters":[{"operator":"AND","dimensions":[{"key":"tenantId","value":"%s"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"region","value":"us-phoenix-1"}],"tags":[],"filters":[]},{"operator":"AND","dimensions":[{"key":"service","value":"COMPUTE"}],"tags":[],"filters":[]}]}`, tenantId)),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
		// verify resource import
		{
			Config:                  config + BudgetCostAnomalyMonitorRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckBudgetCostAnomalyMonitorDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).CostAdClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_budget_cost_anomaly_monitor" {
			noResourceFound = false
			request := oci_budget.GetCostAnomalyMonitorRequest{}

			tmp := rs.Primary.ID
			request.CostAnomalyMonitorId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")

			response, err := client.GetCostAnomalyMonitor(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_budget.CostAnomalyMonitorLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("BudgetCostAnomalyMonitor") {
		resource.AddTestSweepers("BudgetCostAnomalyMonitor", &resource.Sweeper{
			Name:         "BudgetCostAnomalyMonitor",
			Dependencies: acctest.DependencyGraph["costAnomalyMonitor"],
			F:            sweepBudgetCostAnomalyMonitorResource,
		})
	}
}

func sweepBudgetCostAnomalyMonitorResource(compartment string) error {
	costAdClient := acctest.GetTestClients(&schema.ResourceData{}).CostAdClient()
	costAnomalyMonitorIds, err := getBudgetCostAnomalyMonitorIds(compartment)
	if err != nil {
		return err
	}
	for _, costAnomalyMonitorId := range costAnomalyMonitorIds {
		if ok := acctest.SweeperDefaultResourceId[costAnomalyMonitorId]; !ok {
			deleteCostAnomalyMonitorRequest := oci_budget.DeleteCostAnomalyMonitorRequest{}

			deleteCostAnomalyMonitorRequest.CostAnomalyMonitorId = &costAnomalyMonitorId

			deleteCostAnomalyMonitorRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "budget")
			_, error := costAdClient.DeleteCostAnomalyMonitor(context.Background(), deleteCostAnomalyMonitorRequest)
			if error != nil {
				fmt.Printf("Error deleting CostAnomalyMonitor %s %s, It is possible that the resource is already deleted. Please verify manually \n", costAnomalyMonitorId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &costAnomalyMonitorId, BudgetCostAnomalyMonitorSweepWaitCondition, time.Duration(3*time.Minute),
				BudgetCostAnomalyMonitorSweepResponseFetchOperation, "budget", true)
		}
	}
	return nil
}

func getBudgetCostAnomalyMonitorIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "CostAnomalyMonitorId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	costAdClient := acctest.GetTestClients(&schema.ResourceData{}).CostAdClient()

	listCostAnomalyMonitorsRequest := oci_budget.ListCostAnomalyMonitorsRequest{}
	listCostAnomalyMonitorsRequest.CompartmentId = &compartmentId
	listCostAnomalyMonitorsRequest.LifecycleState = oci_budget.CostAnomalyMonitorLifecycleStateActive
	listCostAnomalyMonitorsResponse, err := costAdClient.ListCostAnomalyMonitors(context.Background(), listCostAnomalyMonitorsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting CostAnomalyMonitor list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, costAnomalyMonitor := range listCostAnomalyMonitorsResponse.Items {
		id := *costAnomalyMonitor.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "CostAnomalyMonitorId", id)
	}
	return resourceIds, nil
}

func BudgetCostAnomalyMonitorSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if costAnomalyMonitorResponse, ok := response.Response.(oci_budget.GetCostAnomalyMonitorResponse); ok {
		return costAnomalyMonitorResponse.LifecycleState != oci_budget.CostAnomalyMonitorLifecycleStateDeleted
	}
	return false
}

func BudgetCostAnomalyMonitorSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.CostAdClient().GetCostAnomalyMonitor(context.Background(), oci_budget.GetCostAnomalyMonitorRequest{
		CostAnomalyMonitorId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
