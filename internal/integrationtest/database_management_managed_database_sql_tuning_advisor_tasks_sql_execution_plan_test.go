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
	managedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanSingularDataSourceRepresentation = map[string]interface{}{
		"attribute":                  acctest.Representation{RepType: acctest.Required, Create: `ORIGINAL`},
		"managed_database_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"sql_object_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_object.test_object.id}`},
		"sql_tuning_advisor_task_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_sql_tuning_advisor_task.test_sql_tuning_advisor_task.id}`},
	}

	ManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, namespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "test_object", acctest.Required, acctest.Create, objectRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan.test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", "test_managed_database_sql_tuning_advisor_tasks_sql_execution_plan", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "attribute", "ORIGINAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_object_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_tuning_advisor_task_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan.#"),
			),
		},
	})
}
