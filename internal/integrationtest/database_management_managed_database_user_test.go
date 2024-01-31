// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementDatabaseManagementManagedDatabaseUserSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.db_user}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseUserDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseUserResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	userName := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_name")
	userNameVariableStr := fmt.Sprintf("variable \"db_user\" { default = \"%s\" }\n", userName)

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_users.test_managed_database_users"
	singularDatasourceName := "data.oci_database_management_managed_database_user.test_managed_database_user"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_users", "test_managed_database_users", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user", "test_managed_database_user", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "authentication"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumer_group"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_tablespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "editions_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "password_versions"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "temp_tablespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expiring"),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_users", "test_managed_database_users", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + namedCredentialIdVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_collection.#"),
			),
		},
		// verify singular datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user", "test_managed_database_user", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + namedCredentialIdVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "authentication"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumer_group"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "default_tablespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "editions_enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "password_versions"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "profile"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "temp_tablespace"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_expiring"),
			),
		},
	})
}
