// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
)

var (
	managedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation = map[string]interface{}{
		"execution_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.test_execution_id}`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.test_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_sql_tuning_advisor_task_id}`},
	}

	managedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation = map[string]interface{}{
		"execution_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.test_execution_id}`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.test_object_id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_sql_tuning_advisor_task_id}`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, objectRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	testExecutionId := utils.GetEnvSettingWithBlankDefault("test_execution_id")
	testExecutionIdVariableStr := fmt.Sprintf("variable \"test_execution_id\" { default = \"%s\" }\n", testExecutionId)

	testObjectId := utils.GetEnvSettingWithBlankDefault("test_object_id")
	testObjectIdVariableStr := fmt.Sprintf("variable \"test_object_id\" { default = \"%s\" }\n", testObjectId)

	testSqlTuningAdvisorTaskId := utils.GetEnvSettingWithBlankDefault("test_sql_tuning_advisor_task_id")
	testSqlTuningAdvisorTaskIdVariableStr := fmt.Sprintf("variable \"test_sql_tuning_advisor_task_id\" { default = \"%s\" }\n", testSqlTuningAdvisorTaskId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations.test_managed_database_sql_tuning_advisor_tasks_recommendations"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation.test_managed_database_sql_tuning_advisor_tasks_recommendation"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", "test_managed_database_sql_tuning_advisor_tasks_recommendations", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testExecutionIdVariableStr +
				testObjectIdVariableStr + testSqlTuningAdvisorTaskIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", "test_managed_database_sql_tuning_advisor_tasks_recommendation", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testExecutionIdVariableStr +
				testObjectIdVariableStr + testSqlTuningAdvisorTaskIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
