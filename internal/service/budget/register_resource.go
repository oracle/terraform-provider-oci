// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_budget_alert_rule", BudgetAlertRuleResource())
	tfresource.RegisterResource("oci_budget_budget", BudgetBudgetResource())
	tfresource.RegisterResource("oci_budget_cost_alert_subscription", BudgetCostAlertSubscriptionResource())
	tfresource.RegisterResource("oci_budget_cost_anomaly_event", BudgetCostAnomalyEventResource())
	tfresource.RegisterResource("oci_budget_cost_anomaly_monitor", BudgetCostAnomalyMonitorResource())
	tfresource.RegisterResource("oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management", BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource())
}
