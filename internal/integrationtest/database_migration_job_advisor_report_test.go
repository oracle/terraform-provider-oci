// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	jobAdvisorReportSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_migration_job.test_job.id}`},
	}

	JobAdvisorReportResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job", "test_job", acctest.Required, acctest.Create, jobRepresentation)
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobAdvisorReportResource_basic(t *testing.T) {
	t.Skip("Skip this job creation is an independent operation after validating the migration")
	httpreplay.SetScenario("TestDatabaseMigrationJobAdvisorReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_migration_job_advisor_report.test_job_advisor_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job_advisor_report", "test_job_advisor_report", acctest.Required, acctest.Create, jobAdvisorReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "number_of_fatal_blockers", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "number_of_warnings"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "result"),
			),
		},
	})
}
