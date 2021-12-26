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
	managedDatabaseSqlTuningAdvisorTasksFindingSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
		"begin_exec_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_begin_exec.test_begin_exec.id}`},
		"end_exec_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_end_exec.test_end_exec.id}`},
		"finding_filter":             acctest.Representation{RepType: acctest.Optional, Create: `none`},
		"index_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `indexHashFilter`},
		"search_period":              acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"stats_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `statsHashFilter`},
	}

	managedDatabaseSqlTuningAdvisorTasksFindingDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
		"begin_exec_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_begin_exec.test_begin_exec.id}`},
		"end_exec_id":                acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_management_end_exec.test_end_exec.id}`},
		"finding_filter":             acctest.Representation{RepType: acctest.Optional, Create: `none`},
		"index_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `indexHashFilter`},
		"search_period":              acctest.Representation{RepType: acctest.Optional, Create: `LAST_24HR`},
		"stats_hash_filter":          acctest.Representation{RepType: acctest.Optional, Create: `statsHashFilter`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_findings.test_managed_database_sql_tuning_advisor_tasks_findings"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_finding.test_managed_database_sql_tuning_advisor_tasks_finding"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_findings", "test_managed_database_sql_tuning_advisor_tasks_findings", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksFindingDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig,
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_finding", "test_managed_database_sql_tuning_advisor_tasks_finding", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksFindingSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksFindingResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
