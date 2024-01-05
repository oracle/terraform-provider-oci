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
	DataSafeSqlFirewallAllowedSqlDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DataSafeSqlFirewallAllowedSqlResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallAllowedSqlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlFirewallAllowedSqlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	datasourceName := "data.oci_data_safe_sql_firewall_allowed_sqls.test_sql_firewall_allowed_sqls"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_allowed_sqls", "test_sql_firewall_allowed_sqls", acctest.Optional, acctest.Create, DataSafeSqlFirewallAllowedSqlDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlFirewallAllowedSqlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_firewall_allowed_sql_collection.#"),
			),
		},
	})
}
