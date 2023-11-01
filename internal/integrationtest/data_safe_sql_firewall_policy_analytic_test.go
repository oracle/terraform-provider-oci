// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSqlFirewallPolicyAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`groupBy`}},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"time_ended":                acctest.Representation{RepType: acctest.Optional, Create: `timeEnded`},
		"time_started":              acctest.Representation{RepType: acctest.Optional, Create: `timeStarted`},
	}

	DataSafeSqlFirewallPolicyAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallPolicyAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlFirewallPolicyAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_sql_firewall_policy_analytics.test_sql_firewall_policy_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_policy_analytics", "test_sql_firewall_policy_analytics", acctest.Required, acctest.Create, DataSafeSqlFirewallPolicyAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlFirewallPolicyAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
			),
		},
	})
}
