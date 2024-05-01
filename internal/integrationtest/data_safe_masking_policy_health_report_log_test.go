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
	DataSafeMaskingPolicyHealthReportLogDataSourceRepresentation = map[string]interface{}{
		"masking_policy_health_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.masking_policy_health_report_id}`},
		"message_type":                    acctest.Representation{RepType: acctest.Optional, Create: `PASS`},
	}

	DataSafeMaskingPolicyHealthReportLogResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy_health_report", "test_masking_policy_health_report", acctest.Required, acctest.Create, DataSafeMaskingPolicyHealthReportLogDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingPolicyHealthReportLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingPolicyHealthReportLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	healthReportId := utils.GetEnvSettingWithBlankDefault("masking_health_report_id")
	healthReportIdVariableStr := fmt.Sprintf("variable \"masking_policy_health_report_id\" { default = \"%s\" }\n", healthReportId)

	datasourceName := "data.oci_data_safe_masking_policy_health_report_logs.test_masking_policy_health_report_logs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_policy_health_report_logs", "test_masking_policy_health_report_logs", acctest.Required, acctest.Create, DataSafeMaskingPolicyHealthReportLogDataSourceRepresentation) +
				compartmentIdVariableStr + healthReportIdVariableStr + DataSafeMaskingPolicyHealthReportLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "masking_policy_health_report_id"),
			),
		},
	})
}
