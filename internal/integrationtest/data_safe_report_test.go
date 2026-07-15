// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataSafereportSingularDataSourceRepresentation = map[string]interface{}{
		"report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.rep_identifier}`},
	}

	DataSafereportDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.report_compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `RESTRICTED`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"report_definition_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.report_ocid}`},
	}

	DataSafeReportRepresentation = map[string]interface{}{
		"report_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.rep_identifier}`},
		"defined_tags":  acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}
)

// issue-routing-tag: data_safe/default
func TestDataSafeReportResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	reportCompartmentId := utils.GetEnvSettingWithDefault("report_compartment_ocid", compartmentId)
	reportCompartmentIdVariableStr := fmt.Sprintf("variable \"report_compartment_id\" { default = \"%s\" }\n", reportCompartmentId)

	reportDefId := utils.GetEnvSettingWithBlankDefault("report_ocid")
	reportDefIdVariableStr := fmt.Sprintf("variable \"report_ocid\" { default = \"%s\" }\n", reportDefId)

	reportIdentifier := utils.GetEnvSettingWithBlankDefault("rep_identifier")
	reportIdentifierStr := fmt.Sprintf("variable \"rep_identifier\" { default = \"%s\" }\n", reportIdentifier)

	datasourceName := "data.oci_data_safe_reports.test_reports"
	singularDatasourceName := "data.oci_data_safe_report.test_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_reports", "test_reports", acctest.Required, acctest.Create, DataSafereportDataSourceRepresentation) +
				reportCompartmentIdVariableStr + reportDefIdVariableStr + reportIdentifierStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", reportCompartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "report_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_report", "test_report", acctest.Required, acctest.Create, DataSafereportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + reportDefIdVariableStr + reportIdentifierStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "report_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", reportCompartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "mime_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_generated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
			),
		},
	})
}
