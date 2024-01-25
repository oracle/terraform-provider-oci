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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanSingularDataSourceRepresentation = map[string]interface{}{
		"attribute":                  acctest.Representation{RepType: acctest.Required, Create: `ORIGINAL`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.sql_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, ObjectStorageObjectRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	sqlObjectId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_object_id")
	sqlObjectIdVariableStr := fmt.Sprintf("variable \"sql_object_id\" { default = \"%s\" }\n", sqlObjectId)

	sqlTuningAdvisorId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_tuning_advisor_task_id")
	sqlTuningAdvisorIdVariableStr := fmt.Sprintf("variable \"sql_tuning_advisor_task_id\" { default = \"%s\" }\n", sqlTuningAdvisorId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan.test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlObjectIdVariableStr + sqlTuningAdvisorIdVariableStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute", "ORIGINAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan.#"),
			),
		},
		// verify singular datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + sqlObjectIdVariableStr + sqlTuningAdvisorIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute", "ORIGINAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
