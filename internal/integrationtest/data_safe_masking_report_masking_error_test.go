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
	DataSafeMaskingReportMaskingErrorDataSourceRepresentation = map[string]interface{}{
		"masking_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.masking_report_id}`},
		"step_name":         acctest.Representation{RepType: acctest.Optional, Create: `EXECUTE_MASKING`},
	}

	DataSafeMaskingReportMaskingErrorResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingReportMaskingErrorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeMaskingReportMaskingErrorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	maskingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_report_id")
	maskingReportIdVariableStr := fmt.Sprintf("variable \"masking_report_id\" { default = \"%s\" }\n", maskingReportId)

	datasourceName := "data.oci_data_safe_masking_report_masking_errors.test_masking_report_masking_errors"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_report_masking_errors", "test_masking_report_masking_errors", acctest.Required, acctest.Create, DataSafeMaskingReportMaskingErrorDataSourceRepresentation) +
				compartmentIdVariableStr + maskingReportIdVariableStr + DataSafeMaskingReportMaskingErrorResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "masking_report_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "masking_error_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "masking_error_collection.0.items.#", "2"),
			),
		},
	})
}
