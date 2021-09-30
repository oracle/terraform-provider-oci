// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	jobOutputSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": Representation{RepType: Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	jobOutputDataSourceRepresentation = map[string]interface{}{
		"job_id": Representation{RepType: Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	JobOutputResourceConfig = GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobOutputResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobOutputResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_migration_job_output.test_job_output"
	singularDatasourceName := "data.oci_database_migration_job_output.test_job_output"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_migration_job_output", "test_job_output", Required, Create, jobOutputDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "job_output_summary_collection"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_migration_job_output", "test_job_output", Required, Create, jobOutputSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
