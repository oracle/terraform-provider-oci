// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation = map[string]interface{}{
		"cost_anomaly_monitor_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_budget_cost_anomaly_monitor.test_cost_anomaly_monitor.id}`},
		"enable_costanomalymonitorenabletoggle": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	// Updated dependency chain
	CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, BudgetCostAlertSubscriptionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorRepresentation)
)

// issue-routing-tag: budget/default
func TestBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management.test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"
	parentResourceName := "oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management.test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation), "budget", "costAnomalyMonitorCostanomalymonitorenabletogglesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_costanomalymonitorenabletoggle", `true`),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Create, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + CostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Update, BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_costanomalymonitorenabletoggle", `false`),
			),
		},
	})
}
