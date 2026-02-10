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
	DatabaseMigrationJobAdvisorReportCheckRepresentation = map[string]interface{}{
		"advisor_report_check_id": acctest.Representation{RepType: acctest.Required, Create: `has_users_lack_create_privileges`},
		"is_reviewed":             acctest.Representation{RepType: acctest.Required, Create: `false`, Update: `true`},
		"job_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.job_advisor_id}`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobAdvisorReportCheckResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationJobAdvisorReportCheckResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	jobId := utils.GetEnvSettingWithBlankDefault("job_advisor_id")
	jobIdVariableStr := fmt.Sprintf("variable \"job_advisor_id\" { default = \"%s\" }\n", jobId)

	datasourceName := "oci_database_migration_job_advisor_report_check.test_job_advisor_report_check"

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify singular datasource
		{
			Config: config + compartmentIdVariableStr + jobIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_migration_job_advisor_report_check", "test_job_advisor_report_check", acctest.Required, acctest.Create, DatabaseMigrationJobAdvisorReportCheckRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "advisor_report_check_id"),
				resource.TestCheckResourceAttr(datasourceName, "is_reviewed", "false"),
				resource.TestCheckResourceAttrSet(datasourceName, "job_id"),
			),
		},
	})
}
