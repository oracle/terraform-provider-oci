// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafealertPolicyRuleSingularDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
	}

	DataSafealertPolicyRuleDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
	}

	DataSafeAlertPolicyRuleResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policies", "test_alert_policies", acctest.Required, acctest.Create, DataSafealertPolicyRuleDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertPolicyRuleResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Alert Policy resource")
	httpreplay.SetScenario("TestDataSafeAlertPolicyRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	datasourceName := "data.oci_data_safe_alert_policy_rules.test_alert_policy_rules"
	singularDatasourceName := "data.oci_data_safe_alert_policy_rule.test_alert_policy_rule"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy_rules", "test_alert_policy_rules", acctest.Required, acctest.Create, DataSafealertPolicyRuleDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_rule_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "alert_policy_rule_collection.0.items.0.key", "6"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_rule_collection.0.items.0.expression"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy_rule", "test_alert_policy_rule", acctest.Required, acctest.Create, DataSafealertPolicyRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.0.key", "6"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.0.expression"),
			),
		},
	})
}
