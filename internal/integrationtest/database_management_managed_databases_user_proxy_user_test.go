// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	managedDatabasesUserProxyUserSingularDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.database.oc1.iad.aaaaaaaaptylhec05c6b998f279490b8984a61b02b9472f73837763ff31dc173c0ns2`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `DVSYS`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	managedDatabasesUserProxyUserDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `ocid1.database.oc1.iad.aaaaaaaaptylhec05c6b998f279490b8984a61b02b9472f73837763ff31dc173c0ns2`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `DVSYS`},
		"name":                acctest.Representation{RepType: acctest.Optional, Create: `name`},
	}

	ManagedDatabasesUserProxyUserResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesUserProxyUserResource_basic(t *testing.T) {
	t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesUserProxyUserResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_management_managed_databases_user_proxy_users.test_managed_databases_user_proxy_users"
	singularDatasourceName := "data.oci_database_management_managed_databases_user_proxy_user.test_managed_databases_user_proxy_user"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_proxy_users", "test_managed_databases_user_proxy_users", acctest.Required, acctest.Create, managedDatabasesUserProxyUserDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabasesUserProxyUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "user_name"),

				resource.TestCheckResourceAttrSet(datasourceName, "proxy_user_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "proxy_user_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_databases_user_proxy_user", "test_managed_databases_user_proxy_user", acctest.Required, acctest.Create, managedDatabasesUserProxyUserSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ManagedDatabasesUserProxyUserResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
			),
		},
	})
}
