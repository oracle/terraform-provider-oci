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
	DataSafealertAnalyticSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Required, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"group_by":                  acctest.Representation{RepType: acctest.Required, Create: []string{`targetIds`}},
		"scim_query":                acctest.Representation{RepType: acctest.Required, Create: `status eq \"OPEN\"`},
		"time_ended":                acctest.Representation{RepType: acctest.Required, Create: `2022-01-30T16:02:08.000Z`},
		"time_started":              acctest.Representation{RepType: acctest.Required, Create: `2022-01-28T16:02:08.000Z`},
	}
	DataSafeAlertAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeAlertAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAlertAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_data_safe_alert_analytic.test_alert_analytic"
	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_alert_analytic", "test_alert_analytic", acctest.Required, acctest.Create, DataSafealertAnalyticSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAlertAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
			),
		},
	})
}
