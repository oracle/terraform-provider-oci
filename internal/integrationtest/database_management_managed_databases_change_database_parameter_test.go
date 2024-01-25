// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseManagementManagedDatabasesChangeDatabaseParameterRequiredOnlyResource = DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentation)

	DatabaseManagementManagedDatabasesChangeDatabaseParameterParametersRepresentation = map[string]interface{}{
		"name":           acctest.Representation{RepType: acctest.Required, Create: `open_cursors`},
		"value":          acctest.Representation{RepType: acctest.Required, Create: `500`},
		"update_comment": acctest.Representation{RepType: acctest.Optional, Create: `Increase open cursors`},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentation = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"parameters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterParametersRepresentation},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `BOTH`},
		"credentials":         acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterCredentialsRepresentation},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithDatabaseCredentials = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"parameters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterParametersRepresentation},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `BOTH`},
		"database_credential": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterDatabaseCredentialsRepresentation},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithNamedDatabaseCredentials = map[string]interface{}{
		"managed_database_id": acctest.Representation{RepType: acctest.Required, Create: `${var.managed_database_id}`},
		"parameters":          acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterParametersRepresentation},
		"scope":               acctest.Representation{RepType: acctest.Required, Create: `BOTH`},
		"database_credential": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseManagementManagedDatabasesChangeDatabaseParameterNamedCredentialRepresentation},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterCredentialsRepresentation = map[string]interface{}{
		"role":      acctest.Representation{RepType: acctest.Required, Create: `${var.db_role}`},
		"secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vault_secret_id}`},
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `${var.db_user}`},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterDatabaseCredentialsRepresentation = map[string]interface{}{
		"credential_type":    acctest.Representation{RepType: acctest.Required, Create: `SECRET`},
		"password_secret_id": acctest.Representation{RepType: acctest.Required, Create: `${var.vault_secret_id}`},
		"role":               acctest.Representation{RepType: acctest.Required, Create: `${var.db_role}`},
		"username":           acctest.Representation{RepType: acctest.Required, Create: `${var.db_user}`},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterNamedCredentialRepresentation = map[string]interface{}{
		"credential_type":     acctest.Representation{RepType: acctest.Required, Create: `NAMED_CREDENTIAL`},
		"named_credential_id": acctest.Representation{RepType: acctest.Required, Create: `${var.named_credential_id}`},
	}

	DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies = ""
)

// issue-routing-tag: database_management/default
func TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic(t *testing.T) {
	//t.Skip("Skip this test till Database Management service provides a better way of testing this. It requires a live managed database instance")
	httpreplay.SetScenario("TestDatabaseManagementManagedDatabasesChangeDatabaseParameterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("dbmgmt_compartment_id")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managedDatabaseId := utils.GetEnvSettingWithBlankDefault("dbmgmt_managed_database_id")
	managedDatabaseIdVariableStr := fmt.Sprintf("variable \"managed_database_id\" { default = \"%s\" }\n", managedDatabaseId)

	dbRole := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user_role")
	dbRoleVariableStr := fmt.Sprintf("variable \"db_role\" { default = \"%s\" }\n", dbRole)

	vaultSecretId := utils.GetEnvSettingWithBlankDefault("dbmgmt_vault_secret_id")
	vaultSecretIdVariableStr := fmt.Sprintf("variable \"vault_secret_id\" { default = \"%s\" }\n", vaultSecretId)

	dbUser := utils.GetEnvSettingWithBlankDefault("dbmgmt_db_user")
	dbUserVariableStr := fmt.Sprintf("variable \"db_user\" { default = \"%s\" }\n", dbUser)

	namedCredentialId := utils.GetEnvSettingWithBlankDefault("dbmgmt_named_credential_id")
	namedCredentialIdVariableStr := fmt.Sprintf("variable \"named_credential_id\" { default = \"%s\" }\n", namedCredentialId)

	variableStr := managedDatabaseIdVariableStr + dbRoleVariableStr + vaultSecretIdVariableStr + dbUserVariableStr + namedCredentialIdVariableStr

	resourceName := "oci_database_management_managed_databases_change_database_parameter.test_managed_databases_change_database_parameter"

	var resId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Optional, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentation), "databasemanagement", "managedDatabasesChangeDatabaseParameter", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + variableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Optional, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "credentials.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.password", ""),
				resource.TestCheckResourceAttr(resourceName, "credentials.0.role", "SYSDBA"),
				resource.TestCheckResourceAttrSet(resourceName, "credentials.0.user_name"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.update_comment", "Increase open cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Create with database credential
		{
			Config: config + compartmentIdVariableStr + variableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithDatabaseCredentials),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Optional, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithDatabaseCredentials),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_credential.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_credential.0.password_secret_id", vaultSecretId),
				resource.TestCheckResourceAttrSet(resourceName, "database_credential.0.username"),
				resource.TestCheckResourceAttrSet(resourceName, "database_credential.0.role"),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.update_comment", "Increase open cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify Create with named credential
		{
			Config: config + compartmentIdVariableStr + variableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Required, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithNamedDatabaseCredentials),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + variableStr + DatabaseManagementManagedDatabasesChangeDatabaseParameterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_management_managed_databases_change_database_parameter", "test_managed_databases_change_database_parameter", acctest.Optional, acctest.Create, DatabaseManagementManagedDatabasesChangeDatabaseParameterRepresentationWithNamedDatabaseCredentials),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_credential.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database_credential.0.named_credential_id", namedCredentialId),
				resource.TestCheckResourceAttrSet(resourceName, "managed_database_id"),
				resource.TestCheckResourceAttr(resourceName, "parameters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.name", "open_cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.update_comment", "Increase open cursors"),
				resource.TestCheckResourceAttr(resourceName, "parameters.0.value", "500"),
				resource.TestCheckResourceAttr(resourceName, "scope", "BOTH"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
	})
}
