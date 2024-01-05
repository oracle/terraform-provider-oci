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
	DataSafeauditEventAnalyticSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_by":                  acctest.Representation{RepType: acctest.Optional, Create: []string{`groupBy`}},
		"query_time_zone":           acctest.Representation{RepType: acctest.Optional, Create: `queryTimeZone`},
		"scim_query":                acctest.Representation{RepType: acctest.Optional, Create: `scimQuery`},
		"summary_field":             acctest.Representation{RepType: acctest.Optional, Create: []string{`summaryField`}},
		"time_ended":                acctest.Representation{RepType: acctest.Optional, Create: `timeEnded`},
		"time_started":              acctest.Representation{RepType: acctest.Optional, Create: `timeStarted`},
	}

	DataSafeAuditEventAnalyticResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeAuditEventAnalyticResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeAuditEventAnalyticResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_data_safe_audit_event_analytic.test_audit_event_analytic"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_audit_event_analytic", "test_audit_event_analytic", acctest.Required, acctest.Create, DataSafeauditEventAnalyticSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeAuditEventAnalyticResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
			),
		},
	})
}
