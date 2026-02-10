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
	DatabaseMigrationAssessmentAssessorCheckSingularDataSourceRepresentation = map[string]interface{}{
		"assessment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.assessment_id}`},
		"assessor_name":  acctest.Representation{RepType: acctest.Required, Create: `COMPATIBILITY_ASSESSOR`},
		"check_name":     acctest.Representation{RepType: acctest.Required, Create: `standard_traditional_audit_adb`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `testAllAgain`},
	}

	DatabaseMigrationAssessmentAssessorCheckDataSourceRepresentation = map[string]interface{}{
		"assessment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.assessment_id}`},
		"assessor_name":  acctest.Representation{RepType: acctest.Required, Create: `COMPATIBILITY_ASSESSOR`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `testAllAgain`},
	}

	DatabaseMigrationAssessmentAssessorCheckResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentAssessorCheckResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentAssessorCheckResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	assessmentId := utils.GetEnvSettingWithBlankDefault("assessment_id")
	assessmentIdVariableStr := fmt.Sprintf("variable \"assessment_id\" { default = \"%s\" }\n", assessmentId)

	datasourceName := "data.oci_database_migration_assessment_assessor_checks.test_assessment_assessor_checks"
	singularDatasourceName := "data.oci_database_migration_assessment_assessor_check.test_assessment_assessor_check"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_assessor_checks", "test_assessment_assessor_checks", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorCheckDataSourceRepresentation) +
				compartmentIdVariableStr + assessmentIdVariableStr + DatabaseMigrationAssessmentAssessorCheckResourceConfig + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assessment_id"),
				resource.TestCheckResourceAttr(datasourceName, "assessor_name", "COMPATIBILITY_ASSESSOR"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "assessor_check_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_assessor_check", "test_assessment_assessor_check", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorCheckSingularDataSourceRepresentation) +
				compartmentIdVariableStr + assessmentIdVariableStr + DatabaseMigrationAssessmentAssessorCheckResourceConfig + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assessment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assessor_name", "COMPATIBILITY_ASSESSOR"),
				resource.TestCheckResourceAttr(singularDatasourceName, "check_name", "standard_traditional_audit_adb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "action"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assessor_check_group.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assessor_check_state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "check_action.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "columns.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "help_link_text"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "help_link_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "impact"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "issue"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_location.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metadata.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_count"),
			),
		},
	})
}
