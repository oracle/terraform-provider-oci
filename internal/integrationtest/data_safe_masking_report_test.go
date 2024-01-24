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
	DataSafemaskingReportSingularDataSourceRepresentation = map[string]interface{}{
		"masking_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.masking_report_id}`},
	}

	DataSafemaskingReportDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingReportResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the report ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeMaskingReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	maskingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_report_ocid")
	maskingReportIdVariableStr := fmt.Sprintf("variable \"masking_report_id\" { default = \"%s\" }\n", maskingReportId)

	datasourceName := "data.oci_data_safe_masking_reports.test_masking_reports"
	singularDatasourceName := "data.oci_data_safe_masking_report.test_masking_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + maskingReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_reports", "test_masking_reports", acctest.Required, acctest.Create, DataSafemaskingReportDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "masking_report_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config + maskingReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_report", "test_masking_report", acctest.Required, acctest.Create, DataSafemaskingReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_report_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_drop_temp_tables_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_redo_logging_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_refresh_stats_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "masking_work_request_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parallel_degree"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recompile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_masking_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_columns"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_objects"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_schemas"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_sensitive_types"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_masked_values"),
			),
		},
	})
}
