package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	vaultReplicaRepresentation = map[string]interface{}{
		"replica_region": acctest.Representation{RepType: acctest.Required, Create: `uk-cardiff-1`, Update: `sa-santiago-1`},
		"vault_id":       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_kms_vault.test_vault.id}`},
	}

	KmsVaultReplicationResourceDependencies = KmsKeyResourceDependencies
)

// issue-routing-tag: kms/default
func TestKmsVaultReplicationResource_basic(t *testing.T) {
	t.Skip("Skip this test because virtual private vault is needed")
	httpreplay.SetScenario("TestKmsVaultReplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_kms_vault_replication.test_replica"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+KmsVaultReplicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", acctest.Required, acctest.Create, vaultReplicaRepresentation), "keymanagement", "vaultReplica", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", acctest.Required, acctest.Create, vaultReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "replica_region", "uk-cardiff-1"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + KmsVaultReplicationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_kms_vault_replication", "test_replica", acctest.Required, acctest.Update, vaultReplicaRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "replica_region", "sa-santiago-1"),
				resource.TestCheckResourceAttrSet(resourceName, "vault_id"),
			),
		},
	})
}
