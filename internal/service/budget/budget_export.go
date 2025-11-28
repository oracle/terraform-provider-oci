package budget

import (
	"fmt"

	oci_budget "github.com/oracle/oci-go-sdk/v65/budget"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportBudgetAlertRuleHints.GetIdFn = getBudgetAlertRuleId
	tf_export.RegisterTenancyGraphs("budget", budgetResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getBudgetAlertRuleId(resource *tf_export.OCIResource) (string, error) {

	alertRuleId, ok := resource.SourceAttributes["id"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find alertRuleId for Budget AlertRule")
	}
	budgetId := resource.Parent.Id
	return GetAlertRuleCompositeId(alertRuleId, budgetId), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportBudgetBudgetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_budget_budget",
	DatasourceClass:      "oci_budget_budgets",
	DatasourceItemsAttr:  "budgets",
	ResourceAbbreviation: "budget",
	DiscoverableLifecycleStates: []string{
		string(oci_budget.BudgetLifecycleStateActive),
	},
}

var exportBudgetAlertRuleHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_budget_alert_rule",
	DatasourceClass:      "oci_budget_alert_rules",
	DatasourceItemsAttr:  "alert_rules",
	ResourceAbbreviation: "alert_rule",
	DiscoverableLifecycleStates: []string{
		string(oci_budget.AlertRuleLifecycleStateActive),
	},
}

var exportBudgetCostAnomalyMonitorHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_budget_cost_anomaly_monitor",
	DatasourceClass:        "oci_budget_cost_anomaly_monitors",
	DatasourceItemsAttr:    "cost_anomaly_monitor_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_anomaly_monitor",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_budget.CostAnomalyMonitorLifecycleStateActive),
	},
}

var exportBudgetCostAnomalyEventHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_budget_cost_anomaly_event",
	DatasourceClass:        "oci_budget_cost_anomaly_events",
	DatasourceItemsAttr:    "cost_anomaly_event_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_anomaly_event",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_budget.CostAnomalyEventLifecycleStateActive),
	},
}

var exportBudgetCostAlertSubscriptionHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_budget_cost_alert_subscription",
	DatasourceClass:        "oci_budget_cost_alert_subscriptions",
	DatasourceItemsAttr:    "cost_alert_subscription_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "cost_alert_subscription",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_budget.CostAlertSubscriptionLifecycleStateActive),
	},
}

var budgetResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportBudgetBudgetHints},
		{TerraformResourceHints: exportBudgetCostAnomalyMonitorHints},
		{TerraformResourceHints: exportBudgetCostAnomalyEventHints},
		{TerraformResourceHints: exportBudgetCostAlertSubscriptionHints},
	},
	"oci_identity_tenancy": {
		{
			TerraformResourceHints: exportBudgetBudgetHints,
			DatasourceQueryParams: map[string]string{
				"target_type": "'ALL'",
			},
		},
	},
	"oci_budget_budget": {
		{
			TerraformResourceHints: exportBudgetAlertRuleHints,
			DatasourceQueryParams: map[string]string{
				"budget_id": "id",
			},
		},
	},
}
