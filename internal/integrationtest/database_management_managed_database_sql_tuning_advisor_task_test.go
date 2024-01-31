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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":                        acctest.Representation{RepType: acctest.Optional, Create: `INITIAL`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":                        acctest.Representation{RepType: acctest.Optional, Create: `INITIAL`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":                  acctest.Representation{RepType: acctest.Optional, Create: `INITIAL`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks.test_managed_database_sql_tuning_advisor_tasks"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_task.test_managed_database_sql_tuning_advisor_task"
	namedCredentialDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks.test_managed_database_sql_tuning_advisor_tasks"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks", "test_managed_database_sql_tuning_advisor_tasks", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_task", "test_managed_database_sql_tuning_advisor_task", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks", "test_managed_database_sql_tuning_advisor_tasks", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
