package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vaultReplicaRepresentation = map[string]interface{}{
		"replica_region": Representation{repType: Required, create: `uk-cardiff-1`, update: `sa-santiago-1`},
		"vault_id":       Representation{repType: Required, create: `${data.oci_kms_vault.test_vault.id}`},
	}

	KmsVaultReplicationResourceDependencies = KeyResourceDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultReplicationResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsVaultReplicationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault_replication.test_replica"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	saveConfigContent(config+compartmentIdVariableStr+KmsVaultReplicationResourceDependencies+
		generateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Create, vaultReplicaRepresentation), "keymanagement", "vaultReplica", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Create, vaultReplicaRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "replica_region", "uk-cardiff-1"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
					generateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", Required, Update, vaultReplicaRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "replica_region", "sa-santiago-1"),
					resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
				),
			},
		},
	})
}
