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
	managedDatabaseUserSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"user_name":           Representation{RepType: Required, Create: `{}`},
	}

	managedDatabaseUserDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": Representation{RepType: Required, Create: `${oci_database_management_managed_database.test_managed_database.id}`},
		"name":                Representation{RepType: Optional, Create: `name`},
	}

	ManagedDatabaseUserResourceConfig = GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases", "test_managed_databases", Required, Create, managedDatabaseDataSourceRepresentation)
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabaseUserResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabaseUserResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_database_users.test_managed_database_users"
	singularDatasourceName := "data.oci_database_management_managed_database_user.test_managed_database_user"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_users", "test_managed_database_users", Required, Create, managedDatabaseUserDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "user_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_user", "test_managed_database_user", Required, Create, managedDatabaseUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabaseUserResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
