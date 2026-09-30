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
	DataSafeSubsettingReportSubsettingErrorDataSourceRepresentation = map[string]interface{}{
		"subsetting_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subsetting_report_id}`},
		"step_name":            acctest.Representation{RepType: acctest.Optional, Create: nil},
	}

	DataSafeSubsettingReportSubsettingErrorResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingReportSubsettingErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingReportSubsettingErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	subsettingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_subsetting_report_ocid")
	if subsettingReportId == "" {
		subsettingReportId = utils.GetEnvSettingWithBlankDefault("data_safe_subsetting_report_id")
	}
	subsettingReportIdVariableStr := fmt.Sprintf("variable \"subsetting_report_id\" { default = \"%s\" }\n", subsettingReportId)

	datasourceName := "data.oci_data_safe_subsetting_report_subsetting_errors.test_subsetting_report_subsetting_errors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_report_subsetting_errors", "test_subsetting_report_subsetting_errors", acctest.Required, acctest.Create, DataSafeSubsettingReportSubsettingErrorDataSourceRepresentation) +
				compartmentIdVariableStr + subsettingReportIdVariableStr + DataSafeSubsettingReportSubsettingErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_report_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_error_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "subsetting_error_collection.0.items.#", "0"),
			),
		},
	})
}
