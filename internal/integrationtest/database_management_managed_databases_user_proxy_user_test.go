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
	DatabaseManagementDatabaseManagementManagedDatabasesUserProxyUserSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.user_name}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DatabaseManagementDatabaseManagementManagedDatabasesUserProxyUserDataSourceRepresentation = map[string]interface{}{
		"managed_database_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"user_name":               acctest.Representation{RepType: acctest.Required, Create: `${var.user_name}`},
		"name":                    acctest.Representation{RepType: acctest.Optional, Create: `name`},
		"opc_named_credential_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.opc_named_credential_id}`},
	}

	ManagedDatabasesUserProxyUserResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesUserProxyUserResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesUserProxyUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	userName := utils.GetEnvSettingWithBlankDefault("dbmgmt_user_name")
	userNameStr := fmt.Sprintf("variable \"user_name\" { default = \"%s\" }\n", userName)

	opcNamedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	opcNamedCredentialIdStr := fmt.Sprintf("variable \"opc_named_credential_id\" { default = \"%s\" }\n", opcNamedCredentialId)

	datasourceName := "data.oci_database_management_managed_databases_user_proxy_users.test_managed_databases_user_proxy_users"
	singularDatasourceName := "data.oci_database_management_managed_databases_user_proxy_user.test_managed_databases_user_proxy_user"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_proxy_users", "test_managed_databases_user_proxy_users", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserProxyUserDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameStr + ManagedDatabasesUserProxyUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "proxy_user_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "proxy_user_collection.0.items.#", "0"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_proxy_user", "test_managed_databases_user_proxy_user", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserProxyUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameStr + ManagedDatabasesUserProxyUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "0"),
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_proxy_users", "test_managed_databases_user_proxy_users", acctest.Optional, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabasesUserProxyUserDataSourceRepresentation) +
				compartmentIdVariableStr + managedDatabaseIdVariableStr + userNameStr + opcNamedCredentialIdStr + ManagedDatabasesUserProxyUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "opc_named_credential_id"),
			),
		},
	})
}
