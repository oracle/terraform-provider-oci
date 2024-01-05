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
	DataSafealertPolicySingularDataSourceRepresentation = map[string]interface{}{
		"alert_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${var.policy_id}`},
	}

	DataSafealertPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id_in_subtree":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"display_name":                          acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_alert_policy.test_alert_policy.display_name}`},
		"is_user_defined":                       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"state":                                 acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"time_created_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `2018-01-01T00:00:00.000Z`},
		"time_created_less_than":                acctest.Representation{RepType: acctest.Optional, Create: `2038-01-01T00:00:00.000Z`},
		"alert_policy_type":                     acctest.Representation{RepType: acctest.Optional, Create: `AUDITING`},
	}

	DataSafeAlertPolicyResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policies", "test_alert_policies", acctest.Required, acctest.Create, DataSafealertPolicyDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertPolicyResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for Alert Policy resource")
	httpreplay.SetScenario("TestDataSafeAlertPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	policyId := utils.GetEnvSettingWithBlankDefault("data_safe_alert_policy_ocid")
	policyIdVariableStr := fmt.Sprintf("variable \"policy_id\" { default = \"%s\" }\n", policyId)

	datasourceName := "data.oci_data_safe_alert_policies.test_alert_policies"
	singularDatasourceName := "data.oci_data_safe_alert_policy.test_alert_policy"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + policyIdVariableStr +
				DataSafeAlertPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_collection.0.items.0.display_name"),
				resource.TestCheckResourceAttr(datasourceName, "alert_policy_collection.0.items.0.is_user_defined", "false"),
				resource.TestCheckResourceAttr(datasourceName, "alert_policy_collection.0.items.0.state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "alert_policy_collection.0.items.0.time_updated"),
				resource.TestCheckResourceAttr(datasourceName, "alert_policy_collection.0.items.0.alert_policy_type", "AUDITING"),
			),
		},
		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_policy", "test_alert_policy", acctest.Required, acctest.Create, DataSafealertPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + policyIdVariableStr + DataSafeAlertPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "alert_policy_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
