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
	DataSafeSqlFirewallViolationAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`targetName`}},
		"query_time_zone":           acctest.Representation{RepType: acctest.Optional, Create: `UTC`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `scimQuery`},
		"summary_field":             acctest.Representation{RepType: acctest.Optional, Create: []string{`targetName`}},
		"time_ended":                acctest.Representation{RepType: acctest.Optional, Create: `timeEnded`},
		"time_started":              acctest.Representation{RepType: acctest.Optional, Create: `timeStarted`},
	}

	DataSafeSqlFirewallViolationAnalyticResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallViolationAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlFirewallViolationAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_data_safe_sql_firewall_violation_analytics.test_sql_firewall_violation_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_violation_analytics", "test_sql_firewall_violation_analytics", acctest.Required, acctest.Create, DataSafeSqlFirewallViolationAnalyticDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlFirewallViolationAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sql_firewall_violation_analytics_collection.#"),
			),
		},
	})
}
