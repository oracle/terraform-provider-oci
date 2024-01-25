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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDbIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDbId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	singularDatasourceName := "data.oci_database_management_managed_database_sql_plan_baseline_configuration.test_managed_database_sql_plan_baseline_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baseline_configuration", "test_managed_database_sql_plan_baseline_configuration", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_capture_filters.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_spm_evolve_task_parameters.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_auto_spm_evolve_task_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_automatic_initial_plan_capture_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_high_frequency_auto_spm_evolve_task_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_sql_plan_baselines_usage_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_retention_weeks"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "space_budget_mb"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "space_budget_percent"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "space_used_mb"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baseline_configuration", "test_managed_database_sql_plan_baseline_configuration", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
