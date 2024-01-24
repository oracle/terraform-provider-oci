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
	DataSafemaskingReportsMaskedColumnDataSourceRepresentation = map[string]interface{}{
		"masking_report_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.masking_report_id}`},
		"column_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`columnName`}},
		"masking_column_group": acctest.Representation{RepType: acctest.Optional, Create: []string{`maskingColumnGroup`}},
		"object":               acctest.Representation{RepType: acctest.Optional, Create: []string{`object`}},
		"object_type":          acctest.Representation{RepType: acctest.Optional, Create: []string{`objectType`}},
		"schema_name":          acctest.Representation{RepType: acctest.Optional, Create: []string{`schemaName`}},
		"sensitive_type_id":    acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_safe_sensitive_type.test_sensitive_type.id}`},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeMaskingReportsMaskedColumnResource_basic(t *testing.T) {
	t.Skip("Skipping this test as the report ocid is hardcoded and may not exist when the test runs")
	httpreplay.SetScenario("TestDataSafeMaskingReportsMaskedColumnResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	maskingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_masking_report_ocid")
	maskingReportIdVariableStr := fmt.Sprintf("variable \"masking_report_id\" { default = \"%s\" }\n", maskingReportId)

	datasourceName := "data.oci_data_safe_masking_reports_masked_columns.test_masking_reports_masked_columns"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + maskingReportIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_masking_reports_masked_columns", "test_masking_reports_masked_columns", acctest.Required, acctest.Create, DataSafemaskingReportsMaskedColumnDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "masking_report_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "masked_column_collection.#"),
			),
		},
	})
}
