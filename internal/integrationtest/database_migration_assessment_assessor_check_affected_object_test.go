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
	DatabaseMigrationAssessmentAssessorCheckAffectedObjectDataSourceRepresentation = map[string]interface{}{
		"assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.assessment_id}`},
		"assessor_name": acctest.Representation{RepType: acctest.Required, Create: `COMPATIBILITY_ASSESSOR`},
		"check_name":    acctest.Representation{RepType: acctest.Required, Create: `standard_traditional_audit_adb`},
	}

	DatabaseMigrationAssessmentAssessorCheckAffectedObjectResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentAssessorCheckAffectedObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentAssessorCheckAffectedObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	assessmentId := utils.GetEnvSettingWithBlankDefault("assessment_id")
	assessmentIdVariableStr := fmt.Sprintf("variable \"assessment_id\" { default = \"%s\" }\n", assessmentId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	datasourceName := "data.oci_database_migration_assessment_assessor_check_affected_objects.test_assessment_assessor_check_affected_objects"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_assessor_check_affected_objects", "test_assessment_assessor_check_affected_objects", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorCheckAffectedObjectDataSourceRepresentation) +
				compartmentIdVariableStr + assessmentIdVariableStr + DatabaseMigrationAssessmentAssessorCheckAffectedObjectResourceConfig + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assessment_id"),
				resource.TestCheckResourceAttr(datasourceName, "assessor_name", "COMPATIBILITY_ASSESSOR"),
				resource.TestCheckResourceAttr(datasourceName, "check_name", "standard_traditional_audit_adb"),

				resource.TestCheckResourceAttrSet(datasourceName, "affected_objects_collection.#"),
			),
		},
	})
}
