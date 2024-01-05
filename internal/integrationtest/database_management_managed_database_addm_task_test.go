// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	managedDatabaseAddmTaskSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"time_end":            acctest.Representation{RepType: acctest.Required, Create: `${var.test_timeEnd}`},
		"time_start":          acctest.Representation{RepType: acctest.Required, Create: `${var.test_timeStart}`},
	}

	managedDatabaseAddmTaskDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"time_end":            acctest.Representation{RepType: acctest.Required, Create: `${var.test_timeEnd}`},
		"time_start":          acctest.Representation{RepType: acctest.Required, Create: `${var.test_timeStart}`},
	}

	ManagedDatabaseAddmTaskResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseAddmTaskResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseAddmTaskResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	testTimeStart := utils.GetEnvSettingWithBlankDefault("test_timeStart")
	testTimeStartVariableStr := fmt.Sprintf("variable \"test_timeStart\" { default = \"%s\" }\n", testTimeStart)

	testTimeEnd := utils.GetEnvSettingWithBlankDefault("test_timeEnd")
	testTimeEndVariableStr := fmt.Sprintf("variable \"test_timeEnd\" { default = \"%s\" }\n", testTimeEnd)

	datasourceName := "data.oci_database_management_managed_database_addm_tasks.test_managed_database_addm_tasks"
	singularDatasourceName := "data.oci_database_management_managed_database_addm_task.test_managed_database_addm_task"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_addm_tasks", "test_managed_database_addm_tasks", acctest.Required, acctest.Create, managedDatabaseAddmTaskDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testTimeEndVariableStr + testTimeStartVariableStr + ManagedDatabaseAddmTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_start"),

				resource.TestCheckResourceAttrSet(datasourceName, "addm_tasks_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "addm_tasks_collection.0.items.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_addm_task", "test_managed_database_addm_task", acctest.Required, acctest.Create, managedDatabaseAddmTaskSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testTimeEndVariableStr + testTimeStartVariableStr + ManagedDatabaseAddmTaskResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_end"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_start"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "items.#"),
			),
		},
	})
}
