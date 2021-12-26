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
	managedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_tuning_advisor_task_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
		"begin_exec_id_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"end_exec_id_less_than_or_equal_to":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"search_period":                          acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"time_greater_than_or_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":             acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report.test_managed_database_sql_tuning_advisor_tasks_summary_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", "test_managed_database_sql_tuning_advisor_tasks_summary_report", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "index_findings.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_stat_findings.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "statistics.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_info.#"),
			),
		},
	})
}
