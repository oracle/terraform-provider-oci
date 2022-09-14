package vault

import (
	oci_vault "github.com/oracle/oci-go-sdk/v65/vault"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

func init() {
	tf_export.RegisterCompartmentGraphs("vault", vaultResourceGraph)
}

// Custom overrides for generating composite IDs within the resource discovery framework

// Hints for discovering and exporting this resource to configuration and state files
var exportVaultSecretHints = &tf_export.TerraformResourceHints{
	ResourceClass:          "oci_vault_secret",
	DatasourceClass:        "oci_vault_secrets",
	DatasourceItemsAttr:    "secrets",
	ResourceAbbreviation:   "secret",
	RequireResourceRefresh: true,
	DiscoverableLifecycleStates: []string{
		string(oci_vault.SecretLifecycleStateActive),
	},
}

var vaultResourceGraph = tf_export.TerraformResourceGraph{
	"oci_identity_compartment": {
		{TerraformResourceHints: exportVaultSecretHints},
	},
}
