package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExaccPluggableDatabaseRequiredOnlyResource = DatabaseExaccPluggableDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabaseExaccPluggableDatabaseRepresentation)

	DatabaseExaccPluggableDatabaseResourceConfig = DatabaseExaccPluggableDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Update, DatabaseExaccPluggableDatabaseRepresentation)

	dbHomeExaccRepresentationSourceNonePdb = acctest.RepresentationCopyWithNewProperties(DatabaseExaccDbHomeRepresentationBase, map[string]interface{}{
		"database":      acctest.RepresentationGroup{RepType: acctest.Required, Group: dbHomeDatabaseRepresentationSourceNone},
		"db_system_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
		"db_version":    acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0`},
		"source":        acctest.Representation{RepType: acctest.Optional, Create: `NONE`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `createdDbHomeNone`},
	})

	DatabaseExaccDatabaseRepresentationPdb = map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster.id}`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"key_store_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
	}

	DatabaseExaccPluggableDatabaseRepresentation = map[string]interface{}{
		"container_database_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_database_database.test_database.id}`},
		"pdb_admin_password":                 acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"pdb_name":                           acctest.Representation{RepType: acctest.Required, Create: `SalesPdb`},
		"tde_wallet_password":                acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`},
		"defined_tags":                       acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"should_pdb_admin_account_be_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"lifecycle":                          acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesLBRepresentation},
		"rotate_key_trigger":                 acctest.Representation{RepType: acctest.Required, Create: `0`, Update: `1`},
	}

	DatabaseExaccPluggableDatabaseResourceDependencies = DatabaseVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) + KmsVaultIdVariableStr + OkvSecretVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster", acctest.Required, acctest.Create, dbHomeExaccRepresentationSourceNonePdb) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseExaccDatabaseRepresentationPdb)
)

func TestDatabaseExaccPluggableDatabaseResource(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccPluggableDatabaseResource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_pluggable_database.test_pluggable_database"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExaccPluggableDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Optional, acctest.Create, DatabaseExaccPluggableDatabaseRepresentation), "database", "pluggableDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabasePluggableDatabaseDestroy, []resource.TestStep{
		//Verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccPluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Create, DatabaseExaccPluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
			),
		},
		//Verify Update
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccPluggableDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_pluggable_database", "test_pluggable_database", acctest.Required, acctest.Update, DatabaseExaccPluggableDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "pdb_admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "pdb_name", "SalesPdb"),
				resource.TestCheckResourceAttr(resourceName, "tde_wallet_password", "BEstrO0ng_#11"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseExaccPluggableDatabaseRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"pdb_admin_password",
				"should_pdb_admin_account_be_locked",
				"tde_wallet_password",
				"rotate_key_trigger",
				"state",
			},
			ResourceName: resourceName,
		},
	})
}
