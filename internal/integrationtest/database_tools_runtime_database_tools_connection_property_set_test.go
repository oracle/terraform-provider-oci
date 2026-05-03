// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthOCIAMRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"identity_provider":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthOCIAMIdentityProviderRepresentation},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthOCIAMIdentityProviderRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `OCI_IAM`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthAzureADRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`, Update: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"identity_provider":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthAzureADIdentityProviderRepresentation},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthAzureADIdentityProviderRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `AZURE_AD`},
		"configs": acctest.Representation{RepType: acctest.Required, Create: map[string]string{
			"tenant_id":          `${var.database_tools_property_set_azure_ad_tenant_id}`,
			"application_id":     `${var.database_tools_property_set_azure_ad_application_id}`,
			"application_id_uri": `${var.database_tools_property_set_azure_ad_application_id_uri}`,
		}},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`, Update: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
		"identity_provider":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneIdentityProviderRepresentation},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneIdentityProviderRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `NONE`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexFaIntegrationRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `APEX_FA_INTEGRATION`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `APEX_FA_INTEGRATION`},
		"authentication_substitutions": acctest.Representation{RepType: acctest.Required, Create: map[string]string{
			"FA_IANYQY_DISCOVERY_URL":        `${var.database_tools_property_set_fa_discovery_url}`,
			"FA_IANYQY_OAUTH_SCOPE":          `${var.database_tools_property_set_fa_oauth_scope}`,
			"INTERNAL$FA_IANYQY_PUBLIC_URL":  `${var.database_tools_property_set_fa_public_url}`,
			"INTERNAL$FA_IANYQY_ID":          `${var.database_tools_property_set_fa_id}`,
			"INTERNAL$FA_IANYQY_APEX_APP_ID": `${var.database_tools_property_set_fa_apex_app_id}`,
		}},
		"instance_dbms_credential_enabled": acctest.Representation{RepType: acctest.Required, Create: `Y`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexDocumentGeneratorRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `APEX_DOCUMENT_GENERATOR`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `APEX_DOCUMENT_GENERATOR`},
		"autonomous_database_resource_principal_status": acctest.Representation{RepType: acctest.Required, Create: `ENABLED`},
		"credential_key":                       acctest.Representation{RepType: acctest.Required, Create: `OCI$RESOURCE_PRINCIPAL`},
		"function_id":                          acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_property_set_apex_document_generator_function_id}`},
		"invoke_endpoint":                      acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_property_set_apex_document_generator_invoke_endpoint}`},
		"object_storage_bucket_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_property_set_apex_document_generator_object_storage_bucket_compartment_id}`},
		"object_storage_endpoint":              acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_property_set_apex_document_generator_object_storage_endpoint}`},
		"object_storage_namespace":             acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_property_set_apex_document_generator_object_storage_namespace}`},
		"print_server_type":                    acctest.Representation{RepType: acctest.Required, Create: `DOCUMENT_GENERATOR`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE_EXTERNAL_AUTHENTICATION`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexFaIntegrationSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `APEX_FA_INTEGRATION`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexDocumentGeneratorSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${var.existing_database_tools_connection_id}`},
		"property_set_key":             acctest.Representation{RepType: acctest.Required, Create: `APEX_DOCUMENT_GENERATOR`},
	}

	DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies = ""
)

// issue-routing-tag: database_tools_runtime/default
func TestDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()
	allVars := databaseToolsRuntimeExistingConnectionPropertySetVariables(t)
	existingDatabaseToolsConnectionId := utils.GetEnvSettingWithDefault("existing_database_tools_connection_id",
		utils.GetEnvSettingWithDefault("database_tools_connection_id",
			utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	resourceName := "oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_external_auth"
	apexFaIntegrationResourceName := "oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_apex_fa_integration"
	apexDocumentGeneratorResourceName := "oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_apex_document_generator"
	externalAuthDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_external_auth"
	apexFaIntegrationDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_apex_fa_integration"
	apexDocumentGeneratorDatasourceName := "data.oci_database_tools_runtime_database_tools_connection_property_set.test_database_tools_connection_property_set_apex_document_generator"

	var resId, resId2 string
	acctest.SaveConfigContent(
		config+compartmentIdVariableStr+allVars+DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies+
			acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthOCIAMRepresentation),
		"databasetoolsruntime",
		"databaseToolsConnectionPropertySet",
		t,
	)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthOCIAMRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database_tools_connection_id", existingDatabaseToolsConnectionId),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.type", "OCI_IAM"),
				resource.TestCheckResourceAttrSet(resourceName, "is_mutable"),
				resource.TestCheckResourceAttr(resourceName, "key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"),
				resource.TestCheckResourceAttr(resourceName, "property_set_key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthAzureADRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.type", "AZURE_AD"),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.configs.%", "3"),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.configs.application_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_application_id")),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.configs.application_id_uri", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_application_id_uri")),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.configs.tenant_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_tenant_id")),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		{
			Config: config + compartmentIdVariableStr + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_fa_integration", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexFaIntegrationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_document_generator", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexDocumentGeneratorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "identity_provider.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "identity_provider.0.type", "NONE"),
				resource.TestCheckResourceAttr(resourceName, "key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"),
				resource.TestCheckResourceAttr(apexFaIntegrationResourceName, "authentication_substitutions.%", "5"),
				resource.TestCheckResourceAttr(apexFaIntegrationResourceName, "instance_dbms_credential_enabled", "Y"),
				resource.TestCheckResourceAttr(apexFaIntegrationResourceName, "key", "APEX_FA_INTEGRATION"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "autonomous_database_resource_principal_status", "ENABLED"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "credential_key", "OCI$RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "function_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_function_id")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "invoke_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_invoke_endpoint")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "key", "APEX_DOCUMENT_GENERATOR"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "object_storage_bucket_compartment_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_bucket_compartment_id")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "object_storage_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_endpoint")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "object_storage_namespace", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_namespace")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorResourceName, "print_server_type", "DOCUMENT_GENERATOR"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId2, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_fa_integration", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexFaIntegrationSingularDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_document_generator", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexDocumentGeneratorSingularDataSourceRepresentation) +
				compartmentIdVariableStr + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_fa_integration", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexFaIntegrationRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_apex_document_generator", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetApexDocumentGeneratorRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(externalAuthDatasourceName, "property_set_key", "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"),
				resource.TestCheckResourceAttr(externalAuthDatasourceName, "identity_provider.#", "1"),
				resource.TestCheckResourceAttr(externalAuthDatasourceName, "identity_provider.0.type", "NONE"),
				resource.TestCheckResourceAttr(apexFaIntegrationDatasourceName, "property_set_key", "APEX_FA_INTEGRATION"),
				resource.TestCheckResourceAttr(apexFaIntegrationDatasourceName, "authentication_substitutions.%", "5"),
				resource.TestCheckResourceAttr(apexFaIntegrationDatasourceName, "instance_dbms_credential_enabled", "Y"),
				resource.TestCheckResourceAttrSet(apexFaIntegrationDatasourceName, "is_mutable"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "property_set_key", "APEX_DOCUMENT_GENERATOR"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "autonomous_database_resource_principal_status", "ENABLED"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "credential_key", "OCI$RESOURCE_PRINCIPAL"),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "function_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_function_id")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "invoke_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_invoke_endpoint")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "object_storage_bucket_compartment_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_bucket_compartment_id")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "object_storage_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_endpoint")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "object_storage_namespace", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_namespace")),
				resource.TestCheckResourceAttr(apexDocumentGeneratorDatasourceName, "print_server_type", "DOCUMENT_GENERATOR"),
			),
		},
		{
			Config:            config + compartmentIdVariableStr + allVars + DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_property_set", "test_database_tools_connection_property_set_external_auth", acctest.Required, acctest.Create, DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetExternalAuthNoneRepresentation),
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateCheck: checkImportedDatabaseToolsRuntimeCompositeID(
				resourceName,
				parseDatabaseToolsRuntimeConnectionPropertySetCompositeIDToAttributes,
			),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func databaseToolsRuntimeExistingConnectionPropertySetVariables(t *testing.T) string {
	required := []struct {
		name  string
		value string
	}{
		{"existing_database_tools_connection_id", utils.GetEnvSettingWithDefault("existing_database_tools_connection_id", utils.GetEnvSettingWithDefault("database_tools_connection_id", utils.GetEnvSettingWithBlankDefault("database_tools_connection_ocid")))},
		{"database_tools_property_set_azure_ad_tenant_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_tenant_id")},
		{"database_tools_property_set_azure_ad_application_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_application_id")},
		{"database_tools_property_set_azure_ad_application_id_uri", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_azure_ad_application_id_uri")},
		{"database_tools_property_set_fa_discovery_url", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_fa_discovery_url")},
		{"database_tools_property_set_fa_oauth_scope", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_fa_oauth_scope")},
		{"database_tools_property_set_fa_public_url", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_fa_public_url")},
		{"database_tools_property_set_fa_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_fa_id")},
		{"database_tools_property_set_fa_apex_app_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_fa_apex_app_id")},
		{"database_tools_property_set_apex_document_generator_object_storage_namespace", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_namespace")},
		{"database_tools_property_set_apex_document_generator_object_storage_bucket_compartment_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_bucket_compartment_id")},
		{"database_tools_property_set_apex_document_generator_object_storage_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_object_storage_endpoint")},
		{"database_tools_property_set_apex_document_generator_function_id", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_function_id")},
		{"database_tools_property_set_apex_document_generator_invoke_endpoint", utils.GetEnvSettingWithBlankDefault("database_tools_property_set_apex_document_generator_invoke_endpoint")},
	}

	for _, variable := range required {
		if variable.value == "" {
			t.Skipf("set %s to run this test", variable.name)
		}
	}

	result := ""
	for _, variable := range required {
		result += terraformStringVariable(variable.name, variable.value)
	}
	return result
}
