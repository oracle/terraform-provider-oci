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
	DataSafeSubsettingPolicyHealthReportLogDataSourceRepresentation = map[string]interface{}{
		"subsetting_policy_health_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subsetting_policy_health_report_id}`},
		"message_type":                       acctest.Representation{RepType: acctest.Optional, Create: `PASS`},
	}

	DataSafeSubsettingPolicyHealthReportLogResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_health_report", "test_subsetting_policy_health_report", acctest.Required, acctest.Create, DataSafeSubsettingPolicyHealthReportLogDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingPolicyHealthReportLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingPolicyHealthReportLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	healthReportId := utils.GetEnvSettingWithBlankDefault("subsetting_health_report_id")
	healthReportIdVariableStr := fmt.Sprintf("variable \"subsetting_policy_health_report_id\" { default = \"%s\" }\n", healthReportId)

	datasourceName := "data.oci_data_safe_subsetting_policy_health_report_logs.test_subsetting_policy_health_report_logs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_policy_health_report_logs", "test_subsetting_policy_health_report_logs", acctest.Optional, acctest.Create, DataSafeSubsettingPolicyHealthReportLogDataSourceRepresentation) +
				compartmentIdVariableStr + healthReportIdVariableStr + DataSafeSubsettingPolicyHealthReportLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "message_type"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_policy_health_report_log_collection.0.items.0.health_check_type"),
			),
		},
	})
}
