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
	managedDatabaseSqlTuningAdvisorTaskSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":                        acctest.Representation{RepType: acctest.Optional, Create: `INITIAL`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	managedDatabaseSqlTuningAdvisorTaskDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"name":                          acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"status":                        acctest.Representation{RepType: acctest.Optional, Create: `INITIAL`},
		"time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: `timeGreaterThanOrEqualTo`},
		"time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: `timeLessThanOrEqualTo`},
	}

	ManagedDatabaseSqlTuningAdvisorTaskResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningAdvisorTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_tasks.test_managed_database_sql_tuning_advisor_tasks"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_tuning_advisor_task.test_managed_database_sql_tuning_advisor_task"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_tasks", "test_managed_database_sql_tuning_advisor_tasks", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTaskDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_advisor_task_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_advisor_task", "test_managed_database_sql_tuning_advisor_task", acctest.Required, acctest.Create, managedDatabaseSqlTuningAdvisorTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseSqlTuningAdvisorTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
