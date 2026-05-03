// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialSingularDataSourceRepresentation = map[string]interface{}{
		"credential_key":               acctest.Representation{RepType: acctest.Required, Create: `${var.user_credential_key}`},
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"user_key":                     acctest.Representation{RepType: acctest.Required, Create: `${var.user_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"user_key":                     acctest.Representation{RepType: acctest.Required, Create: `${var.user_key}`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialResourceConfig = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	_ = compartmentId

	datasourceName := "data.oci_database_tools_runtime_database_tools_connection_user_credentials.test_database_tools_connection_user_credentials"
	singularDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_user_credential.test_database_tools_connection_user_credential"

	acctest.SaveConfigContent("", "", "", t)

	allVars := databaseToolsRuntimeExistingUserCredentialVariables()

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_user_credentials", "test_database_tools_connection_user_credentials", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(datasourceName, "user_key", utils.GetEnvSettingWithDefault("user_key", "APEX_240200")),

				resource.TestCheckResourceAttrSet(datasourceName, "user_credential_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_user_credential", "test_database_tools_connection_user_credential", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialSingularDataSourceRepresentation) +
				allVars + DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "credential_key", utils.GetEnvSettingWithDefault("user_credential_key", "OCI$RESOURCE_PRINCIPAL")),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_key", utils.GetEnvSettingWithDefault("user_key", "APEX_240200")),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "enabled"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "owner"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
			),
		},
	})
}

func databaseToolsRuntimeExistingUserCredentialVariables() string {
	return terraformStringVariable("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
			utils.GetEnvSettingWithDefault("database_tools_connection_id",
				utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))) +
		terraformStringVariable("user_key", utils.GetEnvSettingWithDefault("user_key", "APEX_240200")) +
		terraformStringVariable("user_credential_key", utils.GetEnvSettingWithDefault("user_credential_key", "OCI$RESOURCE_PRINCIPAL"))
}
