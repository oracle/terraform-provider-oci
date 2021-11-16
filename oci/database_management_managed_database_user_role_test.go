// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedDatabaseUserRoleSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"user_name":           Representation{RepType: Required, Create: `${oci_identity_user.test_user.name}`},
		"name":                Representation{RepType: Optional, Create: `name`},
	}

	managedDatabaseUserRoleDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"user_name":           Representation{RepType: Required, Create: `${oci_identity_user.test_user.name}`},
		"name":                Representation{RepType: Optional, Create: `name`},
	}

	ManagedDatabaseUserRoleResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserRoleResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserRoleResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_user_roles.test_managed_database_user_roles"
	singularDatasourceName := "data.oci_database_management_managed_database_user_role.test_managed_database_user_role"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_roles", "test_managed_database_user_roles", Required, Create, managedDatabaseUserRoleDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserRoleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "role_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "role_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user_role", "test_managed_database_user_role", Required, Create, managedDatabaseUserRoleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserRoleResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
