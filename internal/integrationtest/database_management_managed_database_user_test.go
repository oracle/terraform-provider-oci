// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"terraform-provider-oci/internal/acctest"
	"terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementDatabaseManagementManagedDatabaseUserSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `${var.test_user_name}`},
	}

	DatabaseManagementDatabaseManagementManagedDatabaseUserDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	DatabaseManagementManagedDatabaseUserResourceConfig = acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	testUserName := utils.GetEnvSettingWithBlankDefault("test_user_name")
	testUserNameVariableStr := fmt.Sprintf("variable \"test_user_name\" { default = \"%s\" }\n", testUserName)

	datasourceName := "data.oci_database_management_managed_database_users.test_managed_database_users"
	singularDatasourceName := "data.oci_database_management_managed_database_user.test_managed_database_user"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_users", "test_managed_database_users", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "user_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user", "test_managed_database_user", acctest.Required, acctest.Create, DatabaseManagementDatabaseManagementManagedDatabaseUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testUserNameVariableStr + DatabaseManagementManagedDatabaseUserResourceConfig,
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
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_locked"),
			),
		},
	})
}
