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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                    acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_tuning_advisor_task_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"begin_exec_id_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"end_exec_id_less_than_or_equal_to":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"search_period":                          acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"time_greater_than_or_equal_to":          acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":             acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	sqlTuningAdvisorTaskId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_tuning_advisor_task_id")
	sqlTuningAdvisorTaskIdVariableStr := fmt.Sprintf("variable \"sql_tuning_advisor_task_id\" { default = \"%s\" }\n", sqlTuningAdvisorTaskId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report.test_managed_database_sql_tuning_advisor_tasks_summary_report"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", "test_managed_database_sql_tuning_advisor_tasks_summary_report", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlTuningAdvisorTaskIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "index_findings.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "object_stat_findings.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "statistics.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_info.#"),
			),
		},

		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report", "test_managed_database_sql_tuning_advisor_tasks_summary_report", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportSingularDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlTuningAdvisorTaskIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
