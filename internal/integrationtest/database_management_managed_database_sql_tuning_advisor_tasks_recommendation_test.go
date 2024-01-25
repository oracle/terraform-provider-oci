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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation = map[string]interface{}{
		"execution_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.execution_id}`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.sql_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation = map[string]interface{}{
		"execution_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.execution_id}`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.sql_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"execution_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.execution_id}`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.sql_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_tuning_advisor_task_id}`},
		"opc_named_credential_id":    acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, ObjectStorageBucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, ObjectStorageObjectRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	executionId := utils.GetEnvSettingWithBlankDefault("dbmgmt_execution_id")
	executionIdVariableStr := fmt.Sprintf("variable \"execution_id\" { default = \"%s\" }\n", executionId)

	objectId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_object_id")
	objectIdVariableStr := fmt.Sprintf("variable \"sql_object_id\" { default = \"%s\" }\n", objectId)

	sqlTuningAdvisorTaskId := utils.GetEnvSettingWithBlankDefault("dbmgmt_sql_tuning_advisor_task_id")
	sqlTuningAdvisorTaskIdVariableStr := fmt.Sprintf("variable \"sql_tuning_advisor_task_id\" { default = \"%s\" }\n", sqlTuningAdvisorTaskId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations.test_managed_database_sql_tuning_advisor_tasks_recommendations"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation.test_managed_database_sql_tuning_advisor_tasks_recommendation"
	namedCredentialDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations.test_managed_database_sql_tuning_advisor_tasks_recommendations"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", "test_managed_database_sql_tuning_advisor_tasks_recommendations", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + executionIdVariableStr + objectIdVariableStr + sqlTuningAdvisorTaskIdVariableStr +
				DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "execution_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_recommendation_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_recommendation_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", "test_managed_database_sql_tuning_advisor_tasks_recommendation", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + executionIdVariableStr + objectIdVariableStr + sqlTuningAdvisorTaskIdVariableStr +
				DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", "test_managed_database_sql_tuning_advisor_tasks_recommendations", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + executionIdVariableStr + objectIdVariableStr + sqlTuningAdvisorTaskIdVariableStr + opcNamedCredentialIdStr +
				DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "execution_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
