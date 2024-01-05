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
	DataSafereportContentSingularDataSourceRepresentation = map[string]interface{}{
		"report_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_data_safe_report.test_report.id}`},
	}

	DataSafeReportContentResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_reports", "test_reports", acctest.Required, acctest.Create, DataSafereportDataSourceRepresentation)
)

// issue-routing-tag: data_safe/default
func TestDataSafeReportContentResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeReportContentResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_data_safe_report_content.test_report_content"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_report_content", "test_report_content", acctest.Required, acctest.Create, DataSafereportContentSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeReportContentResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report_id"),
			),
		},
	})
}
