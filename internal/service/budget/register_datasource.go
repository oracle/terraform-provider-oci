// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package budget

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_budget_alert_rule", BudgetAlertRuleDataSource())
	tfresource.RegisterDatasource("oci_budget_alert_rules", BudgetAlertRulesDataSource())
	tfresource.RegisterDatasource("oci_budget_budget", BudgetBudgetDataSource())
	tfresource.RegisterDatasource("oci_budget_budgets", BudgetBudgetsDataSource())
}
