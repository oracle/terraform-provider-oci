package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vaultReplicaRepresentation = map[string]interface{}{
		"replica_region": Representation{RepType: Required, Create: `uk-cardiff-1`, Update: `sa-santiago-1`},
		"vault_id":       Representation{RepType: Required, Create: `${data.oci_kms_vault.test_vault.id}`},
	}

	KmsVaultReplicationResourceDependencies = KeyResourceDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultReplicationResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsVaultReplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault_replication.test_replica"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+KmsVaultReplicationResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Create, vaultReplicaRepresentation), "keymanagement", "vaultReplica", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Create, vaultReplicaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "replica_region", "uk-cardiff-1"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Update, vaultReplicaRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "replica_region", "sa-santiago-1"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
			),
		},
	})
}
