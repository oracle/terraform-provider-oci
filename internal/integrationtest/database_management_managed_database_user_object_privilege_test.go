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
	DatabaseManagementDatabaseManagementManagedDatabaseUserObjectPrivilegeSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.db_monitoring_user_name}`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.named_credential_id}`},
	}

	DatabaseManagementManagedDatabaseUserObjectPrivilegeResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserObjectPrivilegeResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserObjectPrivilegeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	userName := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_name")
	userNameVariableStr := fmt.Sprintf("variable \"db_monitoring_user_name\" { default = \"%s\" }\n", userName)

	userObjectPrivilageCount := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_monitoring_user_name_object_privileage_count")

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)
	log.Printf("[INFO] named credential is %v", namedCredentialId)

	datasourceName := "data.oci_database_management_managed_database_user_object_privileges.test_managed_database_user_object_privileges"
	singularDatasourceName := "data.oci_database_management_managed_database_user_object_privilege.test_managed_database_user_object_privilege"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_object_privileges", "test_managed_database_user_object_privileges", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabaseUserObjectPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "object_privilege_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "object_privilege_collection.0.items.#", userObjectPrivilageCount),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_object_privilege", "test_managed_database_user_object_privilege", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserObjectPrivilegeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + DatabaseManagementManagedDatabaseUserObjectPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", userObjectPrivilageCount),
			),
		},

		// verify datasource with named credential
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_object_privileges", "test_managed_database_user_object_privileges", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserObjectPrivilegeDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameVariableStr + namedCredentialIdVariableStr + DatabaseManagementManagedDatabaseUserObjectPrivilegeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "object_privilege_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "object_privilege_collection.0.items.#", "0"),
			),
		},
	})
}
