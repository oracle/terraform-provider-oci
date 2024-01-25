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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineJobDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}
	DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineJobDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseSqlPlanBaselineJobResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlPlanBaselineJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlPlanBaselineJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDbId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDbIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDbId)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_plan_baseline_jobs.test_managed_database_sql_plan_baseline_jobs"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baseline_jobs", "test_managed_database_sql_plan_baseline_jobs", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineJobDataSourceRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + DatabaseManagementManagedDatabaseSqlPlanBaselineJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_job_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_plan_baseline_job_collection.0.items.#"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_plan_baseline_jobs", "test_managed_database_sql_plan_baseline_jobs", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlPlanBaselineJobDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDbIdVariableStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlPlanBaselineJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
