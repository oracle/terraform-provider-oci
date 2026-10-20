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
	DataSafeSubsettingReportSubsettedObjectDataSourceRepresentation = map[string]interface{}{
		"subsetting_report_id": acctest.Representation{RepType: acctest.Required, Create: `${var.subsetting_report_id}`},
		"object":               acctest.Representation{RepType: acctest.Required, Create: []string{`EMPLOYEES`}},
		"schema_name":          acctest.Representation{RepType: acctest.Required, Create: []string{`HR_TEST`}},
	}

	DataSafeSubsettingReportSubsettedObjectResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSubsettingReportSubsettedObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSubsettingReportSubsettedObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	subsettingReportId := utils.GetEnvSettingWithBlankDefault("data_safe_subsetting_report_ocid")
	if subsettingReportId == "" {
		subsettingReportId = utils.GetEnvSettingWithBlankDefault("data_safe_subsetting_report_id")
	}
	subsettingReportIdVariableStr := fmt.Sprintf("variable \"subsetting_report_id\" { default = \"%s\" }\n", subsettingReportId)

	datasourceName := "data.oci_data_safe_subsetting_report_subsetted_objects.test_subsetting_report_subsetted_objects"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_subsetting_report_subsetted_objects", "test_subsetting_report_subsetted_objects", acctest.Required, acctest.Create, DataSafeSubsettingReportSubsettedObjectDataSourceRepresentation) +
				compartmentIdVariableStr + subsettingReportIdVariableStr + DataSafeSubsettingReportSubsettedObjectResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "object.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "schema_name.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subsetting_report_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "subsetted_object_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "subsetted_object_collection.0.items.#", "1"),
			),
		},
	})
}
