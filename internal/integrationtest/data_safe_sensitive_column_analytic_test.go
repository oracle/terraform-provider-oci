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
	DataSafeSensitiveColumnAnalyticDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"column_name":               acctest.Representation{RepType: acctest.Optional, Create: []string{`FIRST_NAME`}},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`targetId`}},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${var.target_id}`},
		"object":                    acctest.Representation{RepType: acctest.Optional, Create: []string{`EMPLOYEES`}},
		"schema_name":               acctest.Representation{RepType: acctest.Optional, Create: []string{`HR`}},
	}

	DataSafeSensitiveColumnAnalyticResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sensitive_column_analytics", "test_sensitive_column_analytics", acctest.Optional, acctest.Create, DataSafeSensitiveColumnAnalyticDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSensitiveColumnAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSensitiveColumnAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	targetId := utils.GetEnvSettingWithBlankDefault("data_safe_target_ocid")
	targetIdVariableStr := fmt.Sprintf("variable \"target_id\" { default = \"%s\" }\n", targetId)

	datasourceName := "data.oci_data_safe_sensitive_column_analytics.test_sensitive_column_analytics"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + targetIdVariableStr + DataSafeSensitiveColumnAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "RESTRICTED"),
				resource.TestCheckResourceAttr(datasourceName, "column_name.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
				resource.TestCheckResourceAttr(datasourceName, "group_by.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "target_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "sensitive_column_analytics_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "sensitive_column_analytics_collection.0.items.#", "1"),
			),
		},
	})
}
