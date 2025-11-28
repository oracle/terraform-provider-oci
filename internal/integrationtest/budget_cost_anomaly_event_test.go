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
	BudgetCostAnomalyEventSingularDataSourceRepresentation = map[string]interface{}{
		"cost_anomaly_event_id": acctest.Representation{RepType: acctest.Required, Create: `${var.cost_anomaly_event_id}`},
	}

	BudgetCostAnomalyEventDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"region":           acctest.Representation{RepType: acctest.Optional, Create: []string{"us-ashburn-1"}},
		"target_tenant_id": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.compartment_id}`}},
	}

	BudgetCostAnomalyEventResourceDependencies = ""
	//BudgetCostAnomalyEventResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorRepresentation)
	//DefinedTagsDependencies
)

// issue-routing-tag: budget/default
func TestBudgetCostAnomalyEventResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetCostAnomalyEventResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	// These should be existing cost anomaly event IDs in your test environment
	// You may need to create these manually or fetch them dynamically
	costAnomalyEventId := utils.GetEnvSettingWithBlankDefault("cost_anomaly_event_id")
	costAnomalyEventIdVariableStr := fmt.Sprintf("variable \"cost_anomaly_event_id\" { default = \"%s\" }\n", costAnomalyEventId)

	costAnomalyMonitorId := utils.GetEnvSettingWithBlankDefault("cost_anomaly_monitor_id")
	costAnomalyMonitorIdVariableStr := fmt.Sprintf("variable \"cost_anomaly_monitor_id\" { default = \"%s\" }\n", costAnomalyMonitorId)

	datasourceName := "data.oci_budget_cost_anomaly_events.test_cost_anomaly_events"
	singularDatasourceName := "data.oci_budget_cost_anomaly_event.test_cost_anomaly_event"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+costAnomalyEventIdVariableStr+costAnomalyMonitorIdVariableStr+BudgetCostAnomalyEventResourceDependencies+
		acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_events", "test_cost_anomaly_events", acctest.Optional, acctest.Create, BudgetCostAnomalyEventDataSourceRepresentation), "budget", "costAnomalyEvent", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource - list all cost anomaly events
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_events", "test_cost_anomaly_events", acctest.Required, acctest.Create, BudgetCostAnomalyEventDataSourceRepresentation) +
				compartmentIdVariableStr + costAnomalyMonitorIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "cost_anomaly_event_collection.#"),
			),
		},
		// verify datasource with optional filters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_events", "test_cost_anomaly_events", acctest.Optional, acctest.Create, BudgetCostAnomalyEventDataSourceRepresentation) +
				compartmentIdVariableStr + costAnomalyMonitorIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "region.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "cost_anomaly_event_collection.#"),
			),
		},
		// verify singular datasource - only if cost_anomaly_event_id is provided
		{
			PreConfig: func() {
				if costAnomalyEventId == "" {
					t.Skip("Skipping singular datasource test - cost_anomaly_event_id not provided")
				}
			},
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_event", "test_cost_anomaly_event", acctest.Required, acctest.Create, BudgetCostAnomalyEventSingularDataSourceRepresentation) +
				compartmentIdVariableStr + costAnomalyEventIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_anomaly_event_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_anomaly_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_impact"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_monitor_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_monitor_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_monitor_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "cost_variance_percentage"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "root_cause_detail"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_resource_filter"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_anomaly_event_date"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
