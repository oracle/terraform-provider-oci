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
	BudgetCostAnomalyEventAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"region":           acctest.Representation{RepType: acctest.Optional, Create: []string{"us-ashburn-1"}},
		"target_tenant_id": acctest.Representation{RepType: acctest.Optional, Create: []string{`${var.compartment_id}`}},
	}

	//BudgetCostAnomalyEventAnalyticResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_budget_cost_anomaly_monitor", "test_cost_anomaly_monitor", acctest.Required, acctest.Create, BudgetCostAnomalyMonitorRepresentation)
)

// issue-routing-tag: budget/default
func TestBudgetCostAnomalyEventAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestBudgetCostAnomalyEventAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_budget_cost_anomaly_event_analytics.test_cost_anomaly_event_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_budget_cost_anomaly_event_analytics", "test_cost_anomaly_event_analytics", acctest.Optional, acctest.Create, BudgetCostAnomalyEventAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "cost_anomaly_event_analytic_collection.#"),
			),
		},
	})
}
