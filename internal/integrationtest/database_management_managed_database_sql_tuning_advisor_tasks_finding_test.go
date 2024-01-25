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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"begin_exec_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_begin_exec.test_begin_exec.id}`},
		"end_exec_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_end_exec.test_end_exec.id}`},
		"finding_filter":             acctest.Representation{RepType: acctest.Optional, Create: `none`},
		"index_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `indexHashFilter`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_management_named_credential.test_named_credential_id}`},
		"search_period":              acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"stats_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `statsHashFilter`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"begin_exec_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_begin_exec.test_begin_exec.id}`},
		"end_exec_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_end_exec.test_end_exec.id}`},
		"finding_filter":             acctest.Representation{RepType: acctest.Optional, Create: `none`},
		"index_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `indexHashFilter`},
		"search_period":              acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"stats_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `statsHashFilter`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	sqlTuningAdvisorId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_tuning_advisor_task_id")
	sqlTuningAdvisorIdVariableStr := fmt.Sprintf("variable \"sql_tuning_advisor_task_id\" { default = \"%s\" }\n", sqlTuningAdvisorId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_findings.test_managed_database_sql_tuning_advisor_tasks_findings"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_finding.test_managed_database_sql_tuning_advisor_tasks_finding"
	namedCredentialDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_findings.test_managed_database_sql_tuning_advisor_tasks_findings"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_findings", "test_managed_database_sql_tuning_advisor_tasks_findings", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlTuningAdvisorIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_finding_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_finding_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_finding", "test_managed_database_sql_tuning_advisor_tasks_finding", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlTuningAdvisorIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_findings", "test_managed_database_sql_tuning_advisor_tasks_findings", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlTuningAdvisorIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
