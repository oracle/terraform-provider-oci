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
	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningSetDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name_contains":       acctest.Representation{RepType: acctest.Optional, Create: `${var.name_contains}`},
		"owner":               acctest.Representation{RepType: acctest.Optional, Create: `${var.owner}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningSetDataSourceNamedCredentialRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name_contains":           acctest.Representation{RepType: acctest.Optional, Create: `${var.name_contains}`},
		"owner":                   acctest.Representation{RepType: acctest.Optional, Create: `${var.owner}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}
	DatabaseManagementManagedDatabaseSqlTuningSetResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseSqlTuningSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseSqlTuningSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	nameContains := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user")
	nameContainsStr := fmt.Sprintf("variable \"name_contains\" { default = \"%s\" }\n", nameContains)

	owner := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user")
	ownerStr := fmt.Sprintf("variable \"owner\" { default = \"%s\" }\n", owner)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_sql_tuning_sets.test_managed_database_sql_tuning_sets"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_sets", "test_managed_database_sql_tuning_sets", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningSetDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdStr + DatabaseManagementManagedDatabaseSqlTuningSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.0.name"),
			),
		},
		// verify datasource with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_sets", "test_managed_database_sql_tuning_sets", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningSetDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdStr + nameContainsStr + ownerStr + DatabaseManagementManagedDatabaseSqlTuningSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "name_contains"),
				resource.TestCheckResourceAttrSet(datasourceName, "owner"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.0.name"),
			),
		},
		// verify datasource with optional namedparameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_sql_tuning_sets", "test_managed_database_sql_tuning_sets", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseSqlTuningSetDataSourceNamedCredentialRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdStr + nameContainsStr + ownerStr + opcNamedCredentialIdStr + DatabaseManagementManagedDatabaseSqlTuningSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "name_contains"),
				resource.TestCheckResourceAttrSet(datasourceName, "owner"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_tuning_set_collection.0.items.0.name"),
			),
		},
	})
}
