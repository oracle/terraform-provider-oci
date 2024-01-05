// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	DataSafeSqlFirewallViolationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `targetId eq \"${var.target_id}\"`},
	}

	DataSafeSqlFirewallViolationResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallViolationResource_basic(t *testing.T) {
	t.Skip("Create operation is not available for SQL Firewall Violation resource")
	httpreplay.SetScenario("TestDataSafeSqlFirewallViolationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_sql_firewall_violations.test_sql_firewall_violations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_violations", "test_sql_firewall_violations", acctest.Required, acctest.Create, DataSafeSqlFirewallViolationDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlFirewallViolationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_firewall_violations_collection.#"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_violations", "test_sql_firewall_violations", acctest.Optional, acctest.Create, DataSafeSqlFirewallViolationDataSourceRepresentation) +
				compartmentIdVariableStr + targetIdVariableStr + DataSafeSqlFirewallViolationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_firewall_violations_collection.#"),
			),
		},
	})
}
