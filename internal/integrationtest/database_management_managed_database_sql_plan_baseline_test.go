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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"plan_name":           acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_management_managed_database_sql_plan_baselines.test_managed_database_sql_plan_baselines.sql_plan_baseline_collection.0.items.0.plan_name}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"is_accepted":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_adaptive":         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_enabled":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_fixed":            acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_reproduced":       acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"origin":              acctest.Representation{RepType: acctest.Optional, Create: `ADDM_SQLTUNE`},
		"plan_name":           acctest.Representation{RepType: acctest.Optional, Create: `planName`},
		"sql_handle":          acctest.Representation{RepType: acctest.Optional, Create: `sqlHandle`},
		"sql_text":            acctest.Representation{RepType: acctest.Optional, Create: `sqlText`},
		"limit":               acctest.Representation{RepType: acctest.Required, Create: `1000`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlPlanBaselineResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlPlanBaselineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlPlanBaselineResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDbIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDbId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_plan_baselines.test_managed_database_sql_plan_baselines"
	singularDatasourceName := "data.oci_database_management_managed_database_sql_plan_baseline.test_managed_database_sql_plan_baseline"
	namedCredentialDatasourceName := "data.oci_database_management_managed_database_sql_plan_baselines.test_managed_database_sql_plan_baselines"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baselines", "test_managed_database_sql_plan_baselines", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + DatabaseManagementManagedDatabaseSqlPlanBaselineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.accepted"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.adaptive"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.enabled"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.fixed"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.reproduced"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.origin"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.plan_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.sql_handle"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.0.items.0.sql_text"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baselines", "test_managed_database_sql_plan_baselines", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baseline", "test_managed_database_sql_plan_baseline", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + DatabaseManagementManagedDatabaseSqlPlanBaselineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "adaptive"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "auto_purge"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "execution_plan"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fixed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "origin"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plan_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "reproduced"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_handle"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_text"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_last_modified"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baselines", "test_managed_database_sql_plan_baselines", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlPlanBaselineResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(namedCredentialDatasourceName, "opc_named_credential_id"),
			),
		},
	})
}
