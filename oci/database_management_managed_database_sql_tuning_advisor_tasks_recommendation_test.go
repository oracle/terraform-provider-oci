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
	managedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation = map[string]interface{}{
		"execution_id":               Representation{RepType: Required, Create: `${oci_database_management_execution.test_execution.id}`},
		"managed_database_id":        Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_object_id":              Representation{RepType: Required, Create: `${oci_objectstorage_object.test_object.id}`},
		"sql_tuning_advisor_task_id": Representation{RepType: Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
	}

	managedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation = map[string]interface{}{
		"execution_id":               Representation{RepType: Required, Create: `${oci_database_management_execution.test_execution.id}`},
		"managed_database_id":        Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_object_id":              Representation{RepType: Required, Create: `${oci_objectstorage_object.test_object.id}`},
		"sql_tuning_advisor_task_id": Representation{RepType: Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig = GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", Required, Create, bucketRepresentation) +
		GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation) +
		GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", Required, Create, objectRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations.test_managed_database_sql_tuning_advisor_tasks_recommendations"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation.test_managed_database_sql_tuning_advisor_tasks_recommendation"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations", "test_managed_database_sql_tuning_advisor_tasks_recommendations", Required, Create, managedDatabaseSqlTuningAdvisorTasksRecommendationDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation", "test_managed_database_sql_tuning_advisor_tasks_recommendation", Required, Create, managedDatabaseSqlTuningAdvisorTasksRecommendationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksRecommendationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
