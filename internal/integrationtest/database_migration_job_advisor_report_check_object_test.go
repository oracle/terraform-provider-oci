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
	DatabaseMigrationJobAdvisorReportCheckObjectDataSourceRepresentation = map[string]interface{}{
		"advisor_report_check_id": acctest.Representation{RepType: acctest.Required, Create: `has_users_lack_create_privileges`},
		"job_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.job_advisor_id}`},
	}
)

// issue-routing-tag: database_migration/default
func TestDatabaseMigrationJobAdvisorReportCheckObjectResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseMigrationJobAdvisorReportCheckObjectResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	jobId := utils.GetEnvSettingWithBlankDefault("job_advisor_id")
	jobIdVariableStr := fmt.Sprintf("variable \"job_advisor_id\" { default = \"%s\" }\n", jobId)

	datasourceName := "data.oci_database_migration_job_advisor_report_check_objects.test_job_advisor_report_check_objects"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_migration_job_advisor_report_check_objects", "test_job_advisor_report_check_objects", acctest.Required, acctest.Create, DatabaseMigrationJobAdvisorReportCheckObjectDataSourceRepresentation) +
				compartmentIdVariableStr + jobIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "advisor_report_check_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "job_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "advisor_report_check_objects_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "advisor_report_check_objects_collection.0.items.#", "21"),
			),
		},
	})
}
