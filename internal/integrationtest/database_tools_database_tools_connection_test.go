// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// ═══════════════════════════════════════════════════════════════════════════════
// Shared representations (data-sources, sub-resources, lifecycle)
// ═══════════════════════════════════════════════════════════════════════════════

var (
	DatabaseToolsDatabaseToolsConnectionSingularDataSourceRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
	}

	DatabaseToolsDatabaseToolsConnectionCollectionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `tf_connection_name`, Update: `displayName2`},
		"related_resource_identifier": acctest.Representation{RepType: acctest.Optional, Create: `${var.related_resource_id}`, Update: `identifier2`},
		"runtime_identity":            acctest.Representation{RepType: acctest.Optional, Create: []string{`AUTHENTICATED_PRINCIPAL`}},
		"runtime_support":             acctest.Representation{RepType: acctest.Optional, Create: []string{`SUPPORTED`}},
		"state":                       acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`ORACLE_DATABASE`}},
		"filter":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionDataSourceFilterRepresentation},
	}

	DatabaseToolsDatabaseToolsConnectionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`}},
	}

	// ── Shared sub-object representations ────────────────────────────────────

	DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStoreContentRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStorePasswordRepresentation = map[string]interface{}{
		"value_type": acctest.Representation{RepType: acctest.Required, Create: `SECRETID`},
		"secret_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.secret_id}`},
	}
	DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation = map[string]interface{}{
		"entity_type": acctest.Representation{RepType: acctest.Required, Create: `DATABASE`},
		"identifier":  acctest.Representation{RepType: acctest.Required, Create: `${var.related_resource_id}`, Update: `identifier2`},
	}

	DatabaseToolsIgnoreChangesDatabaseToolsConnectionRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`, `freeform_tags`}},
	}
	DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	// Shared key stores representation (same for all Oracle DB tests)
	DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation = map[string]interface{}{
		"key_store_content":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStoreContentRepresentation},
		"key_store_password": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresKeyStorePasswordRepresentation},
		"key_store_type":     acctest.Representation{RepType: acctest.Optional, Create: `JAVA_KEY_STORE`, Update: `JAVA_TRUST_STORE`},
	}

	DatabaseToolsDatabaseToolsConnectionResourceDependencies = ""

	// ═══════════════════════════════════════════════════════════════════════════
	// Doc 2.4 PASSWORD – proxy_client USER_NAME block, PASSWORD connection auth
	// ───────────────────────────────────────────────────────────────────────────
	// proxyClient: { proxyAuthenticationType: USER_NAME, userName, userPassword, roles }
	// Connection: plain userName + userPassword.
	// This is a double-session approach. Stays USER_NAME on Create and Update.
	// ═══════════════════════════════════════════════════════════════════════════

	DatabaseToolsDatabaseToolsConnectionProxyClientPasswordRepresentation = map[string]interface{}{
		"proxy_authentication_type": acctest.Representation{RepType: acctest.Required, Create: `USER_NAME`, Update: `USER_NAME`},
		"user_name":                 acctest.Representation{RepType: acctest.Required, Create: `proxy_client`, Update: `proxy_client`},
		"user_password":             acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		"roles":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`roles`}, Update: []string{`roles2`}},
	}

	DatabaseToolsDatabaseToolsConnectionPasswordRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `testuser`, Update: `updateduser`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		// Optional so it is only sent in Optional steps (API defaults to PASSWORD anyway).
		"authentication_type": acctest.Representation{RepType: acctest.Optional, Create: `PASSWORD`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"proxy_client":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionProxyClientPasswordRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity":    acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":     acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}

	DatabaseToolsDatabaseToolsConnectionRepresentation = DatabaseToolsDatabaseToolsConnectionPasswordRepresentation

	DatabaseToolsDatabaseToolsConnectionPasswordRequiredOnlyResource = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation)

	DatabaseToolsDatabaseToolsConnectionPasswordResourceConfig = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation)

	// ═══════════════════════════════════════════════════════════════════════════
	// Doc 2.2 TOKEN – bracketed userName only, no proxy_client block
	// ───────────────────────────────────────────────────────────────────────────
	// userName: "[proxy_client1]" — the bracket IS the inline proxy specification.
	// Mutually exclusive with a proxy_client block (API error if both supplied).
	// No userPassword (TOKEN auth forbids it at connection level).
	// ═══════════════════════════════════════════════════════════════════════════

	DatabaseToolsDatabaseToolsConnectionTokenRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		// Required so TOKEN is always sent (prevents defaulting to PASSWORD in step 1).
		"authentication_type": acctest.Representation{RepType: acctest.Required, Create: `TOKEN`},
		// Bracketed format = inline proxy specification. Mutually exclusive with proxy_client block.
		"user_name": acctest.Representation{RepType: acctest.Required, Create: `[testuser]`, Update: `[updateduser]`},
		// user_password omitted – TOKEN auth forbids it.
		// proxy_client omitted – mutually exclusive with bracketed user_name.
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity":    acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":     acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}

	DatabaseToolsDatabaseToolsConnectionTokenRequiredOnlyResource = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenRepresentation)

	DatabaseToolsDatabaseToolsConnectionTokenResourceConfig = DatabaseToolsDatabaseToolsConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenRepresentation)

	// ═══════════════════════════════════════════════════════════════════════════
	// Doc 2.1 – PASSWORD + inline "proxy_user[proxy_client]" in userName
	// ───────────────────────────────────────────────────────────────────────────
	// Single-session. The text before brackets is the proxy user that authenticates,
	// the text inside brackets is the proxy client. No proxy_client block needed.
	// ═══════════════════════════════════════════════════════════════════════════

	DatabaseToolsDatabaseToolsConnectionPasswordInlineProxyRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		// Inline proxy: "proxy_user[proxy_client]" — proxy_user authenticates with the password,
		// the session is opened on behalf of proxy_client. No proxy_client block needed.
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `proxy_user[proxy_client]`, Update: `proxy_user[proxy_client_updated]`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		"authentication_type": acctest.Representation{RepType: acctest.Optional, Create: `PASSWORD`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		// proxy_client omitted – inline bracket syntax is the proxy specification.
		"related_resource": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity": acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":  acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":        acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// Doc 2.4 TOKEN – proxy_client USER_NAME block, TOKEN connection auth
	// ───────────────────────────────────────────────────────────────────────────
	// TOKEN authentication: the IAM token identifies the proxy user at runtime.
	// The proxy_client block names the specific proxy client for a double session.
	// No connection-level userName or userPassword (IAM token provides identity).
	// ═══════════════════════════════════════════════════════════════════════════

	// Stays USER_NAME on Create and Update – only roles change.
	DatabaseToolsDatabaseToolsConnectionProxyClientTokenUserNameRepresentation = map[string]interface{}{
		"proxy_authentication_type": acctest.Representation{RepType: acctest.Required, Create: `USER_NAME`, Update: `USER_NAME`},
		"user_name":                 acctest.Representation{RepType: acctest.Required, Create: `proxy_client`, Update: `proxy_client`},
		// user_password is optional per doc 2.4 table.
		"roles": acctest.Representation{RepType: acctest.Optional, Create: []string{`roles`}, Update: []string{`roles2`}},
	}

	DatabaseToolsDatabaseToolsConnectionTokenProxyUserNameRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		// Required so TOKEN is always sent.
		"authentication_type": acctest.Representation{RepType: acctest.Required, Create: `TOKEN`},
		// user_name omitted – IAM token provides the proxy user identity at runtime.
		// user_password omitted – TOKEN auth forbids it.
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"proxy_client":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionProxyClientTokenUserNameRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity":    acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":     acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}

	// ═══════════════════════════════════════════════════════════════════════════
	// Doc 3.1 – USER_NAME_AUTO_DETECT proxy (works with PASSWORD and TOKEN)
	// ───────────────────────────────────────────────────────────────────────────
	// The DBTools service infers the proxy client from USER_PROXIES at runtime.
	// AUTO_DETECT accepts only roles – no userName, no userPassword in the block.
	// ═══════════════════════════════════════════════════════════════════════════

	// Shared AUTO_DETECT proxy block – only roles, stays AUTO_DETECT on Create and Update.
	DatabaseToolsDatabaseToolsConnectionProxyClientAutoDetectRepresentation = map[string]interface{}{
		"proxy_authentication_type": acctest.Representation{RepType: acctest.Required, Create: `USER_NAME_AUTO_DETECT`, Update: `USER_NAME_AUTO_DETECT`},
		"roles":                     acctest.Representation{RepType: acctest.Optional, Create: []string{`roles`}, Update: []string{`roles2`}},
		// userName and userPassword are NOT valid for USER_NAME_AUTO_DETECT per spec.
	}

	// Doc 3.1 PASSWORD variant: plain userName + userPassword + AUTO_DETECT proxy block.
	DatabaseToolsDatabaseToolsConnectionPasswordAutoDetectRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		"user_name":           acctest.Representation{RepType: acctest.Required, Create: `proxy_user`, Update: `proxy_user_updated`},
		"user_password":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
		"authentication_type": acctest.Representation{RepType: acctest.Optional, Create: `PASSWORD`},
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"proxy_client":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionProxyClientAutoDetectRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity":    acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":     acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}

	// Doc 3.1 TOKEN variant: IAM token provides proxy user identity; AUTO_DETECT
	// infers the proxy client. No connection-level userName or userPassword.
	DatabaseToolsDatabaseToolsConnectionTokenAutoDetectRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `ADMIN@DB202005191141_low`, Update: `displayName2`},
		"type":                acctest.Representation{RepType: acctest.Required, Create: `ORACLE_DATABASE`},
		"connection_string":   acctest.Representation{RepType: acctest.Required, Create: `mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com`, Update: `connectionString2`},
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${var.database_tools_private_endpoint_id}`},
		// Required so TOKEN is always sent.
		"authentication_type": acctest.Representation{RepType: acctest.Required, Create: `TOKEN`},
		// user_name omitted – IAM token provides identity.
		// user_password omitted – TOKEN auth forbids it.
		"advanced_properties": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"oracle.net.ssl_server_dn_match": "true"}, Update: map[string]string{"oracle.net.ssl_server_dn_match": "false"}},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"key_stores":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionKeyStoresRepresentation},
		"proxy_client":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionProxyClientAutoDetectRepresentation},
		"related_resource":    acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseToolsDatabaseToolsConnectionRelatedResourceRepresentation},
		"runtime_identity":    acctest.Representation{RepType: acctest.Optional, Create: `AUTHENTICATED_PRINCIPAL`},
		"runtime_support":     acctest.Representation{RepType: acctest.Optional, Create: `SUPPORTED`},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
	}
)

// ═══════════════════════════════════════════════════════════════════════════════
// Shared test helpers
// ═══════════════════════════════════════════════════════════════════════════════

func dbToolsConnectionTestVars(t *testing.T) (
	config string,
	compartmentIdVariableStr string,
	compartmentIdUVariableStr string,
	privateEndpointIdVariableStr string,
	secretIdVariableStr string,
	relatedResourceIdVariableStr string,
	connectionStringVariableStr string,
	compartmentId string,
	compartmentIdU string,
) {
	config = acctest.ProviderTestConfig()

	compartmentId = utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	if compartmentId == "" || compartmentId == "compartment_id" {
		compartmentId = utils.GetEnvSettingWithBlankDefault("TF_VAR_compartment_id")
	}
	compartmentIdVariableStr = fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU = utils.GetEnvSettingWithDefault("TF_VAR_compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr = fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	privateEndpointId := utils.GetEnvSettingWithBlankDefault("database_tools_private_endpoint_ocid")
	privateEndpointIdVariableStr = fmt.Sprintf("variable \"database_tools_private_endpoint_id\" { default = \"%s\" }\n", privateEndpointId)

	secretId := utils.GetEnvSettingWithBlankDefault("TF_VAR_secret_id")
	secretIdVariableStr = fmt.Sprintf("variable \"secret_id\" { default = \"%s\" }\n", secretId)

	relatedResourceId := utils.GetEnvSettingWithBlankDefault("TF_VAR_related_resource_id")
	relatedResourceIdVariableStr = fmt.Sprintf("variable \"related_resource_id\" { default = \"%s\" }\n", relatedResourceId)

	connectionString := utils.GetEnvSettingWithBlankDefault("TF_VAR_connection_string")
	connectionStringVariableStr = fmt.Sprintf("variable \"connection_string\" { default = \"%s\" }\n", connectionString)

	return
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 2.4 PASSWORD – TestDatabaseToolsConnectionOracleResource_Password
//
// PASSWORD auth + explicit proxy_client USER_NAME block (double-session).
// Connection: plain userName + userPassword.
// Proxy:      proxyClient { USER_NAME, userName, userPassword, roles }.
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_Password(t *testing.T) {
	config, compartmentIdVariableStr, compartmentIdUVariableStr,
		privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, compartmentIdU := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr

	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_connection"

	var resId, resId2 string

	acctest.SaveConfigContent(
		config+allVars+DatabaseToolsDatabaseToolsConnectionResourceDependencies+
			acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation),
		"databasetools", "databaseToolsConnectionPassword", t,
	)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies,
		},

		// Step 3 – create with optionals (proxy_client USER_NAME with userName, userPassword, roles)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "related_resource.0.identifier"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
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

		// Step 4 – update to a different compartment
		{
			Config: config + allVars + compartmentIdUVariableStr + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsConnectionPasswordRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – update parameters (key_store_type → JAVA_TRUST_STORE, roles → roles2)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.user_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "identifier2"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 6 – datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_connections", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionCollectionDataSourceRepresentation) +
				allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "runtime_identity.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "runtime_support.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.#", "1"),
			),
		},

		// Step 7 – singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionSingularDataSourceRepresentation) +
				allVars + DatabaseToolsDatabaseToolsConnectionPasswordResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_client.0.user_password.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_client.0.user_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.identifier", "identifier2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "user_password.0.value_type", "SECRETID"),
			),
		},

		// Step 8 – import
		{
			Config:                  config + DatabaseToolsDatabaseToolsConnectionPasswordRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 2.2 TOKEN – TestDatabaseToolsConnectionOracleResource_Token
//
// TOKEN auth + bracketed userName as inline proxy specification (single-session).
// userName: "[proxy_client]" – bracket syntax is mutually exclusive with proxy_client block.
// No connection-level userPassword (TOKEN auth forbids it).
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_Token(t *testing.T) {
	config, compartmentIdVariableStr, compartmentIdUVariableStr,
		privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, compartmentIdU := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr

	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	datasourceName := "data.oci_database_tools_database_tools_connections.test_database_tools_connections"
	singularDatasourceName := "data.oci_database_tools_database_tools_connection.test_database_tools_connection"

	var resId, resId2 string

	acctest.SaveConfigContent(
		config+allVars+DatabaseToolsDatabaseToolsConnectionResourceDependencies+
			acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenRepresentation),
		"databasetools", "databaseToolsConnectionToken", t,
	)

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				// user_password intentionally not checked – TOKEN auth has no password.
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies,
		},

		// Step 3 – create with optionals
		// No proxy_client block – bracketed user_name IS the proxy specification.
		// Combining both is an API error: "userName in [x] format cannot be combined with proxyClient details".
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				// proxy_client intentionally not checked – TOKEN uses bracketed user_name instead.
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "related_resource.0.identifier"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				// user_password intentionally not checked – TOKEN auth has no password.
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

		// Step 4 – update to a different compartment
		{
			Config: config + allVars + compartmentIdUVariableStr + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(DatabaseToolsDatabaseToolsConnectionTokenRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "mydbsystem.mysubnet.myvcn.oraclevcn.com:1521/mydb_phx1ds.mysubnet.myvcn.oraclevcn.com"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – update parameters (key_store_type → JAVA_TRUST_STORE, user_name → [updateduser])
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_content.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "key_stores.0.key_store_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.0.identifier", "identifier2"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 6 – datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connections", "test_database_tools_connections", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionCollectionDataSourceRepresentation) +
				allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "runtime_identity.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "runtime_support.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "type.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "database_tools_connection_collection.0.items.#", "1"),
			),
		},

		// Step 7 – singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionSingularDataSourceRepresentation) +
				allVars + DatabaseToolsDatabaseToolsConnectionTokenResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "database_tools_connection_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "connection_string", "connectionString2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_content.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_password.0.value_type", "SECRETID"),
				resource.TestCheckResourceAttr(singularDatasourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				// proxy_client intentionally not checked – TOKEN uses bracketed user_name instead.
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.entity_type", "DATABASE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "related_resource.0.identifier", "identifier2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "user_name"),
				// user_password intentionally not checked – TOKEN auth has no password.
			),
		},

		// Step 8 – import
		{
			Config:                  config + DatabaseToolsDatabaseToolsConnectionTokenRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 2.1 – TestDatabaseToolsConnectionOracleResource_PasswordInlineProxy
//
// PASSWORD auth + inline "proxy_user[proxy_client]" in userName (single-session).
// The text before brackets authenticates with the password; the bracketed text is
// the proxy client the session opens on behalf of. No proxy_client block.
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_PasswordInlineProxy(t *testing.T) {
	config, compartmentIdVariableStr, _, privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, _ := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordInlineProxyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name contains the inline "proxy_user[proxy_client]" form.
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies},

		// Step 3 – create with optionals (no proxy_client block – inline bracket is the proxy)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordInlineProxyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				// The API always returns a proxyClient object; inline bracket userName results in NO_PROXY.
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "NO_PROXY"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name should contain the "proxy_user[proxy_client]" form.
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 4 – update (user_name proxy client updated, key_store_type rotated)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordInlineProxyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "NO_PROXY"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – import
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordInlineProxyRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 2.4 TOKEN – TestDatabaseToolsConnectionOracleResource_TokenProxyUserName
//
// TOKEN auth + explicit proxy_client USER_NAME block (double-session).
// IAM token provides the proxy user identity at runtime.
// No connection-level userName or userPassword.
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_TokenProxyUserName(t *testing.T) {
	config, compartmentIdVariableStr, _, privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, _ := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create (TOKEN, no user_name, no user_password)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenProxyUserNameRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name and user_password absent – IAM token is the identity.
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies},

		// Step 3 – create with optionals (proxy_client USER_NAME, userName, roles)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenProxyUserNameRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				// user_password absent from proxy_client (optional per doc 2.4 table).
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name and user_password absent at connection level – IAM token is the identity.
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 4 – update (roles → roles2, key_store_type → JAVA_TRUST_STORE)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenProxyUserNameRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME"),
				resource.TestCheckResourceAttrSet(resourceName, "proxy_client.0.user_name"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – import
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenProxyUserNameRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 3.1 PASSWORD – TestDatabaseToolsConnectionOracleResource_PasswordAutoDetect
//
// PASSWORD auth + proxy_client USER_NAME_AUTO_DETECT (double-session).
// The DBTools service queries USER_PROXIES at runtime to infer the proxy client.
// Connection: plain userName + userPassword.
// Proxy block: only roles – no userName, no userPassword in the proxy block.
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_PasswordAutoDetect(t *testing.T) {
	config, compartmentIdVariableStr, _, privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, _ := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies},

		// Step 3 – create with optionals (USER_NAME_AUTO_DETECT proxy – roles only, no proxy userName/password)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionPasswordAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME_AUTO_DETECT"),
				// user_name and user_password absent from proxy_client per spec.
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_password.0.secret_id"),
				resource.TestCheckResourceAttr(resourceName, "user_password.0.value_type", "SECRETID"),
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 4 – update (roles → roles2, key_store_type → JAVA_TRUST_STORE)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "PASSWORD"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME_AUTO_DETECT"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "user_name"),
				resource.TestCheckResourceAttr(resourceName, "user_password.#", "1"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – import
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionPasswordAutoDetectRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

// ═══════════════════════════════════════════════════════════════════════════════
// Doc 3.1 TOKEN – TestDatabaseToolsConnectionOracleResource_TokenAutoDetect
//
// TOKEN auth + proxy_client USER_NAME_AUTO_DETECT (double-session).
// IAM token provides the proxy user identity; DBTools auto-detects the proxy client.
// No connection-level userName or userPassword.
// Proxy block: only roles – no userName, no userPassword in the proxy block.
// ═══════════════════════════════════════════════════════════════════════════════

// issue-routing-tag: database_tools/default
func TestDatabaseToolsConnectionOracleResource_TokenAutoDetect(t *testing.T) {
	config, compartmentIdVariableStr, _, privateEndpointIdVariableStr, secretIdVariableStr,
		relatedResourceIdVariableStr, connectionStringVariableStr, compartmentId, _ := dbToolsConnectionTestVars(t)

	allVars := compartmentIdVariableStr + privateEndpointIdVariableStr + secretIdVariableStr + relatedResourceIdVariableStr + connectionStringVariableStr
	resourceName := "oci_database_tools_database_tools_connection.test_database_tools_connection"
	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy, []resource.TestStep{
		// Step 1 – Required-only create (TOKEN, no user_name, no user_password)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name and user_password absent – IAM token is the identity.
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 2 – delete before next create
		{Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies},

		// Step 3 – create with optionals (USER_NAME_AUTO_DETECT proxy – roles only)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Create, DatabaseToolsDatabaseToolsConnectionTokenAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "advanced_properties.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "ADMIN@DB202005191141_low"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_KEY_STORE"),
				resource.TestCheckResourceAttrSet(resourceName, "private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME_AUTO_DETECT"),
				// user_name and user_password absent from proxy_client per spec.
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "related_resource.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "runtime_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "runtime_identity", "AUTHENTICATED_PRINCIPAL"),
				resource.TestCheckResourceAttr(resourceName, "runtime_support", "SUPPORTED"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				// user_name and user_password absent at connection level – IAM token is the identity.
				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// Step 4 – update (roles → roles2, key_store_type → JAVA_TRUST_STORE)
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Optional, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenAutoDetectRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "authentication_type", "TOKEN"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "key_stores.0.key_store_type", "JAVA_TRUST_STORE"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.proxy_authentication_type", "USER_NAME_AUTO_DETECT"),
				resource.TestCheckResourceAttr(resourceName, "proxy_client.0.roles.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "ORACLE_DATABASE"),
				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Step 5 – import
		{
			Config: config + allVars + DatabaseToolsDatabaseToolsConnectionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Update, DatabaseToolsDatabaseToolsConnectionTokenAutoDetectRepresentation),
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func TestDatabaseToolsConnectionGenericJdbcResource_basic(t *testing.T) {

	DBToolsVars := databaseToolsStandardVariables()
	ConnectionResourceType := "oci_database_tools_database_tools_connection"
	ConnectionName := "test_jdbc_connection"
	ResourceReference := fmt.Sprintf("%s.%s", ConnectionResourceType, ConnectionName)
	DataReference := fmt.Sprintf("data.%s.%s", ConnectionResourceType, ConnectionName)

	var (
		ResourceRepresentation = map[string]interface{}{
			"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
			"display_name":    acctest.Representation{RepType: acctest.Required, Create: `tf_generic_jdbc_connection`},
			"type":            acctest.Representation{RepType: acctest.Required, Create: `GENERIC_JDBC`},
			"runtime_support": acctest.Representation{RepType: acctest.Required, Create: `UNSUPPORTED`},
			"user_name":       acctest.Representation{RepType: acctest.Required, Create: `testuser`},
			"url":             acctest.Representation{RepType: acctest.Required, Create: `jdbc:oracle:thin:@my.db.server:1521:my_sid`},
			"user_password":   acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsDatabaseToolsConnectionUserPasswordRepresentation},
			"lifecycle":       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseToolsIgnoreDefinedTagsChangesDatabaseToolsConnectionRepresentation},
		}
		DataRepresentation = map[string]interface{}{
			"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: fmt.Sprintf("${%s.%s.id}", ConnectionResourceType, ConnectionName)},
		}
	)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t) },
		Providers:    databaseToolsOciProvider(),
		CheckDestroy: testAccCheckDatabaseToolsDatabaseToolsConnectionDestroy,
		Steps: []resource.TestStep{
			// Step 1. Verify create
			{
				Config: DBToolsVars + acctest.GenerateResourceFromRepresentationMap(ConnectionResourceType, ConnectionName, acctest.Required, acctest.Create, ResourceRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(ResourceReference, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(ResourceReference, "display_name", "tf_generic_jdbc_connection"),
					resource.TestCheckResourceAttr(ResourceReference, "type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(ResourceReference, "url", "jdbc:oracle:thin:@my.db.server:1521:my_sid"),
				),
			},
			// Step 2. Verify singular datasource
			{
				Config: DBToolsVars +
					acctest.GenerateResourceFromRepresentationMap(ConnectionResourceType, ConnectionName, acctest.Required, acctest.Create, ResourceRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap(ConnectionResourceType, ConnectionName, acctest.Required, acctest.Create, DataRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(DataReference, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(DataReference, "display_name", "tf_generic_jdbc_connection"),
					resource.TestCheckResourceAttr(DataReference, "type", "GENERIC_JDBC"),
					resource.TestCheckResourceAttr(DataReference, "url", "jdbc:oracle:thin:@my.db.server:1521:my_sid"),
				),
			},
			// Step 3. Verify resource import
			{
				Config:                  DBToolsVars + acctest.GenerateResourceFromRepresentationMap(ConnectionResourceType, ConnectionName, acctest.Required, acctest.Create, ResourceRepresentation),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            ResourceReference,
			},
		},
	})
}

func init() {
	databaseToolsInitDependencyGraphAndSweeper("DatabaseToolsDatabaseToolsConnection")
}
