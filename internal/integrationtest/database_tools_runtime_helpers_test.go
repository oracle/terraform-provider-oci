package integrationtest

import (
	"fmt"
	"strconv"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func databaseToolsRuntimeBaseVariables() string {
	return databaseToolsStandardVariables()
}

func databaseToolsRuntimeStandardVariables() string {
	image := terraformStringVariable("image", utils.GetEnvSettingWithBlankDefault("image"))
	pbfListingId := terraformStringVariable("pbf_listing_id", utils.GetEnvSettingWithBlankDefault("pbf_listing_id"))
	keyId := terraformStringVariable("key_id", utils.GetEnvSettingWithDefault("key_id", utils.GetEnvSettingWithBlankDefault("kms_key_ocid")))
	vaultId := terraformStringVariable("vault_id", utils.GetEnvSettingWithDefault("vault_id", utils.GetEnvSettingWithBlankDefault("kms_vault_ocid")))
	return databaseToolsStandardVariables() + image + pbfListingId + keyId + vaultId
}

func terraformStringVariable(name string, value string) string {
	return fmt.Sprintf("variable %q { default = %s }\n", name, strconv.Quote(value))
}

var (
	DatabaseToolsRuntimeMinimalConnectionRepresentation = acctest.RepresentationCopyWithNewProperties(
		acctest.GetRepresentationCopyWithMultipleRemovedProperties([]string{
			"advanced_properties",
			"defined_tags",
			"proxy_client",
			"freeform_tags",
			"key_stores",
			"private_endpoint_id",
			"related_resource",
			"runtime_identity",
			"runtime_support",
			"lifecycle",
		}, DatabaseToolsDatabaseToolsConnectionRepresentation),
		map[string]interface{}{
			"connection_string": acctest.Representation{RepType: acctest.Required, Create: `${var.connection_string}`},
		},
	)

	DatabaseToolsRuntimeMinimalCredentialRepresentation = map[string]interface{}{
		"database_tools_connection_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_tools_database_tools_connection.test_database_tools_connection.id}`},
		"key":                          acctest.Representation{RepType: acctest.Required, Create: `credentialKey`, Update: `credentialKey`},
		"password":                     acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"type":                         acctest.Representation{RepType: acctest.Required, Create: `BASIC`, Update: `BASIC`},
		"user_name":                    acctest.Representation{RepType: acctest.Required, Create: `testuser`, Update: `testuser`},
	}

	DatabaseToolsRuntimeMinimalConnectionResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_tools_database_tools_connection", "test_database_tools_connection", acctest.Required, acctest.Create, DatabaseToolsRuntimeMinimalConnectionRepresentation)

	DatabaseToolsRuntimeMinimalCredentialResourceDependencies = DatabaseToolsRuntimeMinimalConnectionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_tools_runtime_database_tools_connection_credential", "test_database_tools_connection_credential", acctest.Required, acctest.Create, DatabaseToolsRuntimeMinimalCredentialRepresentation)
)
