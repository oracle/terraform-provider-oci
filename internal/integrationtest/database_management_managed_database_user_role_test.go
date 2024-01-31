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
	managedDatabaseUserRoleSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	managedDatabaseUserRoleDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}

	ManagedDatabaseUserRoleResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserRoleResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserRoleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	userName := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_name")
	userNameVariableStr := fmt.Sprintf("variable \"db_monitoring_user_name\" { default = \"%s\" }\n", userName)

	monitoringUserRoleCount := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_role_count")

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_user_roles.test_managed_database_user_roles"
	singularDatasourceName := "data.oci_database_management_managed_database_user_role.test_managed_database_user_role"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_roles", "test_managed_database_user_roles", acctest.Required, acctest.Create, managedDatabaseUserRoleDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + ManagedDatabaseUserRoleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "role_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "role_collection.0.items.#", monitoringUserRoleCount),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_role", "test_managed_database_user_role", acctest.Required, acctest.Create, managedDatabaseUserRoleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + ManagedDatabaseUserRoleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", monitoringUserRoleCount),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_roles", "test_managed_database_user_roles", acctest.Optional, acctest.Create, managedDatabaseUserRoleDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + namedCredentialIdVariableStr + ManagedDatabaseUserRoleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "role_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "role_collection.0.items.#", monitoringUserRoleCount),
			),
		},
	})
}
