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
	DataSafeSecurityPolicyReportSingularDataSourceRepresentation = map[string]interface{}{
		"security_policy_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.security_policy_report_id}`},
	}

	DataSafeSecurityPolicyReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Required, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"security_policy_report_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_security_policy_report.test_security_policy_report.id}`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"target_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_cloud_guard_target.test_target.id}`},
	}

	DataSafeSecurityPolicyReportResourceConfig = DefinedTagsDependencies
)

// issue-routing-tag: data_safe/default
func TestDataSafeSecurityPolicyReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSecurityPolicyReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	securityPolicyReportId := utils.GetEnvSettingWithBlankDefault("security_policy_report_ocid")
	securityPolicyReportIdVariableStr := fmt.Sprintf("variable \"security_policy_report_id\" { default = \"%s\" }\n", securityPolicyReportId)

	datasourceName := "data.oci_data_safe_security_policy_reports.test_security_policy_reports"
	singularDatasourceName := "data.oci_data_safe_security_policy_report.test_security_policy_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_reports", "test_security_policy_reports", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSecurityPolicyReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "security_policy_report_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_security_policy_report", "test_security_policy_report", acctest.Required, acctest.Create, DataSafeSecurityPolicyReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + securityPolicyReportIdVariableStr + DataSafeSecurityPolicyReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "security_policy_report_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
