// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementDatabaseManagementManagedDatabasesUserSystemPrivilegeSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DatabaseManagementDatabaseManagementManagedDatabasesUserSystemPrivilegeDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}

	DatabaseManagementManagedDatabasesUserSystemPrivilegeResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesUserSystemPrivilegeResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesUserSystemPrivilegeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_databases_user_system_privileges.test_managed_databases_user_system_privileges"
	singularDatasourceName := "data.oci_database_management_managed_databases_user_system_privilege.test_managed_databases_user_system_privilege"

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	userName := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_name")
	userNameVariableStr := fmt.Sprintf("variable \"db_monitoring_user_name\" { default = \"%s\" }\n", userName)

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	userSystemPrivilageCount := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_system_privleage_count")

	log.Printf("[INFO] named credential is %v", namedCredentialId)

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_system_privileges", "test_managed_databases_user_system_privileges", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserSystemPrivilegeDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabasesUserSystemPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "system_privilege_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "system_privilege_collection.0.items.#", userSystemPrivilageCount),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_system_privilege", "test_managed_databases_user_system_privilege", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserSystemPrivilegeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabasesUserSystemPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", userSystemPrivilageCount),
			),
		},
		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_system_privileges", "test_managed_databases_user_system_privileges", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserSystemPrivilegeDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + namedCredentialIdVariableStr + DatabaseManagementManagedDatabasesUserSystemPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "system_privilege_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "system_privilege_collection.0.items.#", userSystemPrivilageCount),
			),
		},
	})
}
