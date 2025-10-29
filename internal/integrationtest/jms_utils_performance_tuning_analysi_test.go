// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
)

var (
	// Due to unfortunate naming Analysi is the singular name, Analysis the plural name
	PerformanceTuningAnalysiProjectName = "DO_NOT_DELETE_TERRAFORM_TEST"
	PerformanceTuningAnalysiId          = JmsUtilsPerformanceTuningReportId

	JmsUtilsPerformanceTuningAnalysiDataSourceRepresentation = map[string]interface{}{
		"analysis_project_name": acctest.Representation{RepType: acctest.Optional, Create: PerformanceTuningAnalysiProjectName},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: JmsTenancyId},
	}

	JmsUtilsPerformanceTuningAnalysiSingularDataSourceRepresentation = map[string]interface{}{
		"performance_tuning_analysis_id": acctest.Representation{RepType: acctest.Required, Create: PerformanceTuningAnalysiId},
	}
)

// issue-routing-tag: jms_utils/default
// issue-routing-tag: jms_utils/default
func TestJmsUtilsPerformanceTuningAnalysiResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsPerformanceTuningAnalysiResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_utils_performance_tuning_analysis.test_performance_tuning_analysis"
	singularDatasourceName := "data.oci_jms_utils_performance_tuning_analysi.test_performance_tuning_analysi"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify create
		// note: we cannot write test for this case because
		// we don't have create API.

		// verify update
		// note: we cannot write test for this case because
		// we don't have update API.

		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_performance_tuning_analysis",
					"test_performance_tuning_analysis",
					acctest.Optional,
					acctest.Create,
					JmsUtilsPerformanceTuningAnalysiDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsTenancyId),
				resource.TestCheckResourceAttr(datasourceName, "analysis_project_name", PerformanceTuningAnalysiProjectName),

				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.#"),
				// we can verify there's 1 report with name "DO_NOT_DELETE_TERRAFORM_TEST"
				resource.TestCheckResourceAttr(datasourceName, "performance_tuning_analysis_collection.0.items.#", "1"),

				// check actual data matches data used for the List API
				resource.TestCheckResourceAttr(datasourceName, "performance_tuning_analysis_collection.0.items.0.analysis_project_name", PerformanceTuningAnalysiProjectName),
				resource.TestCheckResourceAttr(datasourceName, "performance_tuning_analysis_collection.0.items.0.compartment_id", JmsTenancyId),
				resource.TestCheckResourceAttr(datasourceName, "performance_tuning_analysis_collection.0.items.0.id", PerformanceTuningAnalysiId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.result"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.warning_count"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.created_by.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.created_by.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.created_by.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.result_object_storage_path"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.artifact_object_storage_path"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.time_finished"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "performance_tuning_analysis_collection.0.items.0.work_request_id"),
			),
		},

		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_performance_tuning_analysi",
					"test_performance_tuning_analysi",
					acctest.Optional,
					acctest.Create,
					JmsUtilsPerformanceTuningAnalysiSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) error {
					// Add debug logging or print statements here
					log.Printf("Singular Data Source Attributes: %#v", s.RootModule().Resources[singularDatasourceName].Primary.Attributes)
					return nil
				},

				// check actual data matches data used for the GET API
				resource.TestCheckResourceAttr(singularDatasourceName, "id", PerformanceTuningAnalysiId),
				resource.TestCheckResourceAttr(singularDatasourceName, "analysis_project_name", PerformanceTuningAnalysiProjectName),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", JmsTenancyId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "warning_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result_object_storage_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "artifact_object_storage_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
			),
		},
	})
}
