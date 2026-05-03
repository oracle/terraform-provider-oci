package database_tools_runtime

import (
	"fmt"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymId
	exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalId
	exportDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetId
	exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecId
	exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolId
	exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialId
	exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId
	exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeHints.GetIdFn = getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeId
	tf_export.RegisterCompartmentGraphs("database_tools_runtime", databaseToolsRuntimeResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymId(resource *tf_export.OCIResource) (string, error) {
	credentialKey, ok := resource.SourceAttributes["credential_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find credentialKey for DatabaseToolsRuntime DatabaseToolsConnectionCredentialPublicSynonym")
	}
	databaseToolsConnectionId := resource.Parent.Id
	publicSynonymKey, ok := resource.SourceAttributes["public_synonym_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find publicSynonymKey for DatabaseToolsRuntime DatabaseToolsConnectionCredentialPublicSynonym")
	}
	return GetDatabaseToolsConnectionCredentialPublicSynonymCompositeId(credentialKey, databaseToolsConnectionId, publicSynonymKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalId(resource *tf_export.OCIResource) (string, error) {
	databaseToolsDatabaseApiGatewayConfigId := resource.Parent.Id
	globalKey, ok := resource.SourceAttributes["global_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find globalKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigGlobal")
	}
	return GetDatabaseToolsDatabaseApiGatewayConfigGlobalCompositeId(databaseToolsDatabaseApiGatewayConfigId, globalKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetId(resource *tf_export.OCIResource) (string, error) {
	databaseToolsConnectionId := resource.Parent.Id
	propertySetKey, ok := resource.SourceAttributes["property_set_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find propertySetKey for DatabaseToolsRuntime DatabaseToolsConnectionPropertySet")
	}
	return GetDatabaseToolsConnectionPropertySetCompositeId(databaseToolsConnectionId, propertySetKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecId(resource *tf_export.OCIResource) (string, error) {
	apiSpecKey, ok := resource.SourceAttributes["api_spec_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find apiSpecKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	}
	databaseToolsDatabaseApiGatewayConfigId := resource.Parent.Id
	poolKey, ok := resource.SourceAttributes["pool_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find poolKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigPoolApiSpec")
	}
	return GetDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecCompositeId(apiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolId(resource *tf_export.OCIResource) (string, error) {
	databaseToolsDatabaseApiGatewayConfigId := resource.Parent.Id
	poolKey, ok := resource.SourceAttributes["pool_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find poolKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigPool")
	}
	return GetDatabaseToolsDatabaseApiGatewayConfigPoolCompositeId(databaseToolsDatabaseApiGatewayConfigId, poolKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialId(resource *tf_export.OCIResource) (string, error) {
	credentialKey, ok := resource.SourceAttributes["credential_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find credentialKey for DatabaseToolsRuntime DatabaseToolsConnectionCredential")
	}
	databaseToolsConnectionId := resource.Parent.Id
	return GetDatabaseToolsConnectionCredentialCompositeId(credentialKey, databaseToolsConnectionId), nil
}

func getDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecId(resource *tf_export.OCIResource) (string, error) {
	autoApiSpecKey, ok := resource.SourceAttributes["auto_api_spec_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find autoApiSpecKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	}
	databaseToolsDatabaseApiGatewayConfigId := resource.Parent.Id
	poolKey, ok := resource.SourceAttributes["pool_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find poolKey for DatabaseToolsRuntime DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec")
	}
	return GetDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecCompositeId(autoApiSpecKey, databaseToolsDatabaseApiGatewayConfigId, poolKey), nil
}

func getDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeId(resource *tf_export.OCIResource) (string, error) {
	credentialKey, ok := resource.SourceAttributes["credential_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find credentialKey for DatabaseToolsRuntime DatabaseToolsConnectionCredentialExecuteGrantee")
	}
	databaseToolsConnectionId := resource.Parent.Id
	executeGranteeKey, ok := resource.SourceAttributes["execute_grantee_key"].(string)
	if !ok {
		return "", fmt.Errorf("[ERROR] unable to find executeGranteeKey for DatabaseToolsRuntime DatabaseToolsConnectionCredentialExecuteGrantee")
	}
	return GetDatabaseToolsConnectionCredentialExecuteGranteeCompositeId(credentialKey, databaseToolsConnectionId, executeGranteeKey), nil
}

// Hints for discovering and exporting this resource to configuration and state files
var exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_connection_credential_public_synonym",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_connection_credential_public_synonyms",
	DatasourceItemsAttr:    "credential_public_synonym_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_connection_credential_public_synonym",
	RequireResourceRefresh: true,
}

var exportDatabaseToolsRuntimeDatabaseToolsConnectionPropertySetHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_tools_runtime_database_tools_connection_property_set",
	DatasourceClass:      "oci_database_tools_runtime_database_tools_connection_property_set",
	ResourceAbbreviation: "database_tools_connection_property_set",
}

var exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs",
	DatasourceItemsAttr:    "database_tools_database_api_gateway_config_pool_api_spec_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_database_api_gateway_config_pool_api_spec",
	RequireResourceRefresh: true,
}

var exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalHints = &tf_export.TerraformResourceHints{
	ResourceClass:        "oci_database_tools_runtime_database_tools_database_api_gateway_config_global",
	DatasourceClass:      "oci_database_tools_runtime_database_tools_database_api_gateway_config_global",
	ResourceAbbreviation: "database_tools_database_api_gateway_config_global",
}

var exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_database_api_gateway_config_pools",
	DatasourceItemsAttr:    "database_tools_database_api_gateway_config_pool_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_database_api_gateway_config_pool",
	RequireResourceRefresh: true,
}

var exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_connection_credential",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_connection_credentials",
	DatasourceItemsAttr:    "credential_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_connection_credential",
	RequireResourceRefresh: true,
}

var exportDatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_specs",
	DatasourceItemsAttr:    "database_tools_database_api_gateway_config_pool_auto_api_spec_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_database_api_gateway_config_pool_auto_api_spec",
	RequireResourceRefresh: true,
}

var exportDatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_database_tools_runtime_database_tools_connection_credential_execute_grantee",
	DatasourceClass:        "oci_database_tools_runtime_database_tools_connection_credential_execute_grantees",
	DatasourceItemsAttr:    "credential_execute_grantee_collection",
	IsDatasourceCollection: true,
	ResourceAbbreviation:   "database_tools_connection_credential_execute_grantee",
	RequireResourceRefresh: true,
}

var databaseToolsRuntimeResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {},
}
