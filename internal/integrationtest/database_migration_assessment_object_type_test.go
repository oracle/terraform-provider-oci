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
	DatabaseMigrationAssessmentObjectTypeDataSourceRepresentation = map[string]interface{}{
		"connection_type": acctest.Representation{RepType: acctest.Required, Create: `ORACLE`},
	}

	DatabaseMigrationAssessmentObjectTypeResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_assessment", "test_assessment", acctest.Required, acctest.Create, DatabaseMigrationAssessmentRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationAssessmentObjectTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationAssessmentObjectTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	sourceConnectionOracleId := utils.GetEnvSettingWithBlankDefault("source_connection_oracle_id")
	sourceConnectionOracleIdVariableStr := fmt.Sprintf("variable \"source_connection_oracle_id\" { default = \"%s\" }\n", sourceConnectionOracleId)

	targetConnectionOracleId := utils.GetEnvSettingWithBlankDefault("target_connection_oracle_id")
	targetConnectionOracleIdVariableStr := fmt.Sprintf("variable \"target_connection_oracle_id\" { default = \"%s\" }\n", targetConnectionOracleId)

	datasourceName := "data.oci_database_migration_assessment_object_types.test_assessment_object_types"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_assessment_object_types", "test_assessment_object_types", acctest.Required, acctest.Create, DatabaseMigrationAssessmentObjectTypeDataSourceRepresentation) +
				compartmentIdVariableStr + sourceConnectionOracleIdVariableStr + targetConnectionOracleIdVariableStr + DatabaseMigrationAssessmentObjectTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "connection_type", "ORACLE"),

				resource.TestCheckResourceAttrSet(datasourceName, "assessment_object_type_summary_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "assessment_object_type_summary_collection.0.items.#", "32"),
			),
		},
	})
}
