// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseMigrationAssessmentAssessorActionRepresentation = map[string]interface{}{
		"assessment_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_assessment.test_assessment.id}`},
		"assessor_action": acctest.Representation{RepType: acctest.Required, Create: `RUN`},
		"assessor_name":   acctest.Representation{RepType: acctest.Required, Create: `VIABILITY_ASSESSOR`},
		"items":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseMigrationAssessmentAssessorActionItemsRepresentation},
	}
	DatabaseMigrationAssessmentAssessorActionItemsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `value`},
	}

	DatabaseMigrationAssessmentAssessorActionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentAssessorActionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentAssessorActionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	resourceName := "oci_database_migration_assessment_assessor_action.test_assessment_assessor_action"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseMigrationAssessmentAssessorActionResourceDependencies+compartmentIdUVariableStr+sourceConnectionOracleIdVariableStr+targetConnectionOracleIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment_assessor_action", "test_assessment_assessor_action", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorActionRepresentation), "databasemigration", "assessmentAssessorAction", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseMigrationAssessmentAssessorActionResourceDependencies + compartmentIdUVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment_assessor_action", "test_assessment_assessor_action", acctest.Required, acctest.Create, DatabaseMigrationAssessmentAssessorActionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "assessment_id"),
				resource.TestCheckResourceAttr(resourceName, "assessor_action", "RUN"),
				resource.TestCheckResourceAttr(resourceName, "assessor_name", "VIABILITY_ASSESSOR"),
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "items.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "items.0.value", "value"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
	})
}
