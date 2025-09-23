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

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// Due to unfortunate naming Analysi is the singular name, Analysis the plural name
	JmsUtilsJavaMigrationAnalysiCompartmentId = utils.GetEnvSettingWithBlankDefault("tenancy_ocid")
	JavaMigrationAnalysiProjectName           = "DO_NOT_DELETE_TERRAFORM_TEST"
	JavaMigrationAnalysiId                    = utils.GetEnvSettingWithBlankDefault("java_migration_report_ocid")

	JmsUtilsJavaMigrationAnalysiDataSourceRepresentation = map[string]interface{}{
		"analysis_project_name": acctest.Representation{RepType: acctest.Optional, Create: JavaMigrationAnalysiProjectName},
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: JmsUtilsJavaMigrationAnalysiCompartmentId},
	}

	JmsUtilsJavaMigrationAnalysiSingularDataSourceRepresentation = map[string]interface{}{
		"java_migration_analysis_id": acctest.Representation{RepType: acctest.Required, Create: JavaMigrationAnalysiId},
	}
)

// issue-routing-tag: jms_utils/default
func TestJmsUtilsJavaMigrationAnalysiResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestJmsUtilsJavaMigrationAnalysiResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	datasourceName := "data.oci_jms_utils_java_migration_analysis.test_java_migration_analysis"
	singularDatasourceName := "data.oci_jms_utils_java_migration_analysi.test_java_migration_analysi"

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
					"oci_jms_utils_java_migration_analysis",
					"test_java_migration_analysis",
					acctest.Optional,
					acctest.Create,
					JmsUtilsJavaMigrationAnalysiDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(datasourceName, "compartment_id", JmsUtilsJavaMigrationAnalysiCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "analysis_project_name", JavaMigrationAnalysiProjectName),

				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.#"),
				// we can verify there's 1 report with name "DO_NOT_DELETE_TERRAFORM_TEST"
				resource.TestCheckResourceAttr(datasourceName, "java_migration_analysis_collection.0.items.#", "1"),

				// check actual data matches data used for the List API
				resource.TestCheckResourceAttr(datasourceName, "java_migration_analysis_collection.0.items.0.analysis_project_name", JavaMigrationAnalysiProjectName),
				resource.TestCheckResourceAttr(datasourceName, "java_migration_analysis_collection.0.items.0.compartment_id", JmsUtilsJavaMigrationAnalysiCompartmentId),
				resource.TestCheckResourceAttr(datasourceName, "java_migration_analysis_collection.0.items.0.id", JavaMigrationAnalysiId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.analysis_result_files.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.analysis_result_object_storage_path"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.bucket"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.created_by.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.created_by.0.display_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.created_by.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.input_applications_object_storage_paths.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.input_applications_object_storage_paths.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.namespace"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.target_jdk_version"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.time_finished"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.time_started"),
				resource.TestCheckResourceAttrSet(datasourceName, "java_migration_analysis_collection.0.items.0.work_request_id"),
			),
		},

		// verify singular datasource

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_jms_utils_java_migration_analysi",
					"test_java_migration_analysi",
					acctest.Optional,
					acctest.Create,
					JmsUtilsJavaMigrationAnalysiSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) error {
					// Add debug logging or print statements here
					log.Printf("Singular Data Source Attributes: %#v", s.RootModule().Resources[singularDatasourceName].Primary.Attributes)
					return nil
				},

				// check actual data matches data used for the GET API
				resource.TestCheckResourceAttr(singularDatasourceName, "id", JavaMigrationAnalysiId),
				resource.TestCheckResourceAttr(singularDatasourceName, "analysis_project_name", JavaMigrationAnalysiProjectName),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", JmsUtilsJavaMigrationAnalysiCompartmentId),
				// check actual data is set (doesn't make much sense to hardcode more values)
				resource.TestCheckResourceAttrSet(singularDatasourceName, "analysis_result_files.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "analysis_result_object_storage_path"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "bucket"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.0.display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by.0.id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "input_applications_object_storage_paths.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "input_applications_object_storage_paths.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "namespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "target_jdk_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "work_request_id"),
			),
		},
	})
}
