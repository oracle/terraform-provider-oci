// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	managedDatabaseOptimizerStatisticsAdvisorExecutionScriptSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
		"execution_name":      acctest.Representation{RepType: acctest.Required, Create: `${element(element(data.oci_database_management_managed_database_optimizer_statistics_advisor_executions.test_managed_database_optimizer_statistics_advisor_executions.optimizer_statistics_advisor_executions_collection, 0).items, 0).execution_name}`},
		"task_name":           acctest.Representation{RepType: acctest.Required, Create: `${element(element(data.oci_database_management_managed_database_optimizer_statistics_advisor_executions.test_managed_database_optimizer_statistics_advisor_executions.optimizer_statistics_advisor_executions_collection, 0).items, 0).task_name}`},
	}

	managedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.tenancy_ocid}testManagedDatabase0`},
		"start_time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: time.Now().UTC().AddDate(0, 0, -4).Format("2006-01-02T15:04:05.000Z")},
		"end_time_less_than_or_equal_to":      acctest.Representation{RepType: acctest.Optional, Create: time.Now().UTC().Format("2006-01-02T15:04:05.000Z")},
	}
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_management_managed_database_optimizer_statistics_advisor_execution_script.test_managed_database_optimizer_statistics_advisor_execution_script"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_advisor_executions", "test_managed_database_optimizer_statistics_advisor_executions", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsAdvisorExecutionsDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_optimizer_statistics_advisor_execution_script", "test_managed_database_optimizer_statistics_advisor_execution_script", acctest.Required, acctest.Create, managedDatabaseOptimizerStatisticsAdvisorExecutionScriptSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "task_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "script"),
			),
		},
	})
}
