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
	DatabaseManagementManagedDatabasePreferredCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"credential_name":     acctest.Representation{RepType: acctest.Required, Create: `${var.test_preferred_credential_name}`},
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
	}

	DatabaseManagementManagedDatabasePreferredCredentialDataSourceRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.test_managed_database_id}`},
	}

	DatabaseManagementManagedDatabasePreferredCredentialResourceConfig = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasePreferredCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasePreferredCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	testManagedDatabaseId := utils.GetEnvSettingWithBlankDefault("test_managed_database_id")
	testManagedDatabaseIdVariableStr := fmt.Sprintf("variable \"test_managed_database_id\" { default = \"%s\" }\n", testManagedDatabaseId)

	testPreferredCredentialName := utils.GetEnvSettingWithBlankDefault("test_preferred_credential_name")
	testPreferredCredentialNameVariableStr := fmt.Sprintf("variable \"test_preferred_credential_name\" { default = \"%s\" }\n", testPreferredCredentialName)

	datasourceName := "data.oci_database_management_managed_database_preferred_credentials.test_managed_database_preferred_credentials"
	singularDatasourceName := "data.oci_database_management_managed_database_preferred_credential.test_managed_database_preferred_credential"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_preferred_credentials", "test_managed_database_preferred_credentials", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasePreferredCredentialDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + DatabaseManagementManagedDatabasePreferredCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "preferred_credential_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_management_managed_database_preferred_credential", "test_managed_database_preferred_credential", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasePreferredCredentialSingularDataSourceRepresentation) +
				compartmentIdVariableStr + testManagedDatabaseIdVariableStr + testPreferredCredentialNameVariableStr + DatabaseManagementManagedDatabasePreferredCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "credential_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "managed_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "credential_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_accessible"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "password_secret_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "status"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
			),
		},
	})
}
