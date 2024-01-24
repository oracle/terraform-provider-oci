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
	DatabaseMigrationjobOutputSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	DatabaseMigrationjobOutputDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	DatabaseMigrationJobOutputResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobOutputResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobOutputResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_migration_job_output.test_job_output"
	singularDatasourceName := "data.oci_database_migration_job_output.test_job_output"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job_output", "test_job_output", acctest.Required, acctest.Create, DatabaseMigrationjobOutputDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "job_output_summary_collection"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job_output", "test_job_output", acctest.Required, acctest.Create, DatabaseMigrationjobOutputSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
