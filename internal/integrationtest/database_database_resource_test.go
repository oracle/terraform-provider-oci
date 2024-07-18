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
	DatabaseExaccRequiredOnlyResource = DatabaseExaccDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseExaccDatabaseRepresentation)

	DatabaseExaccDatabaseResourceConfig = DatabaseExaccDatabaseResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Update, DatabaseExaccDatabaseRepresentation)

	DatabaseExaccDatabaseRepresentation = map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster.id}`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"db_version":   acctest.Representation{RepType: acctest.Optional, Create: `19.0.0.0`},
		"key_store_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_key_store.test_key_store.id}`},
	}

	databaseExaccRepresentationMigration = map[string]interface{}{
		"database":     acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster.id}`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"key_store_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
	}

	databaseExaccRepresentationRotation = map[string]interface{}{
		"database":         acctest.RepresentationGroup{RepType: acctest.Required, Group: databaseDatabaseRepresentation},
		"db_home_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_db_home.test_db_home_vm_cluster.id}`},
		"source":           acctest.Representation{RepType: acctest.Required, Create: `NONE`},
		"kms_key_rotation": acctest.Representation{RepType: acctest.Required, Update: `1`},
		"key_store_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_key_store.test_key_store.id}`},
	}

	DatabaseExaccDbHomeRepresentationBase = map[string]interface{}{
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}

	DatabaseExaccDatabaseResourceDependencies = DatabaseVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, DatabaseVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) + KmsVaultIdVariableStr + OkvSecretVariableStr +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home_vm_cluster", acctest.Required, acctest.Create, dbHomeRepresentationSourceVmClusterExacc)
)

func TestDatabaseExaccDatabaseResource(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccDatabaseResource")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_database.test_database"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseExaccDatabaseResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseExaccDatabaseRepresentation), "database", "database", t)

	acctest.ResourceTest(t, testAccCheckDatabaseDatabaseDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, DatabaseExaccDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// verify change key store type
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Update, databaseExaccRepresentationMigration),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// verify database key rotation
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Update, databaseExaccRepresentationRotation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies,
		},
		// verify create optional
		{
			Config: config + compartmentIdVariableStr + DatabaseExaccDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Optional, acctest.Create, DatabaseExaccDatabaseRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "database.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "database.0.admin_password", "BEstrO0ng_#11"),
				resource.TestCheckResourceAttr(resourceName, "database.0.db_name", "myTestDb"),
				resource.TestCheckResourceAttrSet(resourceName, "db_home_id"),
				resource.TestCheckResourceAttrSet(resourceName, "key_store_id"),
				resource.TestCheckResourceAttr(resourceName, "source", "NONE"),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseExaccRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"database",
				"db_version",
				"kms_key_rotation",
				"source",
				"key_store_id",
			},
			ResourceName: resourceName,
		},
	})
}
