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
	jobAdvisorReportSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": Representation{RepType: Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	JobAdvisorReportResourceConfig = GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", Required, Create, jobRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobAdvisorReportResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobAdvisorReportResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_migration_job_advisor_report.test_job_advisor_report"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_migration_job_advisor_report", "test_job_advisor_report", Required, Create, jobAdvisorReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_fatal_blockers", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_warnings"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result"),
			),
		},
	})
}
