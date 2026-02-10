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
	DatabaseMigrationAssessmentAssessorSingularDataSourceRepresentation = map[string]interface{}{
		"assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_assessment.test_assessment.id}`},
		"assessor_name": acctest.Representation{RepType: acctest.Required, Create: `VIABILITY_ASSESSOR`},
	}

	DatabaseMigrationAssessmentAssessorDataSourceRepresentation = map[string]interface{}{
		"assessment_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_assessment.test_assessment.id}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"state":         acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
	}

	DatabaseMigrationAssessmentAssessorResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentAssessorResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentAssessorResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	datasourceName := "data.oci_database_migration_assessment_assessors.test_assessment_assessors"
	singularDatasourceName := "data.oci_database_migration_assessment_assessor.test_assessment_assessor"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_assessors", "test_assessment_assessors", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseMigrationAssessmentAssessorResourceConfig + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "assessment_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "assessor_summary_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_assessor", "test_assessment_assessor", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseMigrationAssessmentAssessorResourceConfig + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "assessment_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assessor_name", "VIABILITY_ASSESSOR"),

				resource.TestCheckResourceAttr(singularDatasourceName, "actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "assessor_group.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "VIABILITY_ASSESSOR"),
			),
		},
	})
}
