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
	managedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                    Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_tuning_advisor_task_id":             Representation{RepType: Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
		"begin_exec_id_greater_than_or_equal_to": Representation{RepType: Optional, Create: `10`},
		"end_exec_id_less_than_or_equal_to":      Representation{RepType: Optional, Create: `10`},
		"search_period":                          Representation{RepType: Optional, Create: `LAST_24HR`},
		"time_greater_than_or_equal_to":          Representation{RepType: Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":             Representation{RepType: Optional, Create: `timeLessThanOrEqualTo`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig = GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report.test_managed_database_sql_tuning_advisor_tasks_summary_report"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", "test_managed_database_sql_tuning_advisor_tasks_summary_report", Required, Create, managedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
