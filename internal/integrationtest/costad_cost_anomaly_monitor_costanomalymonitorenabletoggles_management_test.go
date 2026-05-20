// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation = map[string]interface{}{
		"cost_anomaly_monitor_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_costad_cost_anomaly_monitor.test_cost_anomaly_monitor.id}`},
		"enable_costanomalymonitorenabletoggle": acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}

	CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_alert_subscription", "test_cost_alert_subscription", acctest.Required, acctest.Create, CostadCostAlertSubscriptionRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, CostadCostAnomalyMonitorRepresentation)
)

// issue-routing-tag: costad/default
func TestCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management.test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"
	parentResourceName := "oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management.test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation), "costad", "costAnomalyMonitorCostanomalymonitorenabletogglesManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// create with enable
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// verify enable
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Required, acctest.Create, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_costanomalymonitorenabletoggle", `true`),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies,
		},
		// create with enable and optional fields
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Create, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// update to disable
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Update, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "cost_anomaly_monitor_id"),
			),
		},
		// verify disable
		{
			Config: config + compartmentIdVariableStr + CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", "test_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", acctest.Optional, acctest.Update, CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(parentResourceName, "enable_costanomalymonitorenabletoggle", `false`),
			),
		},
	})
}
