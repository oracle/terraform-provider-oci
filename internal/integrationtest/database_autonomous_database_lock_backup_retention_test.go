package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseAutonomousDatabaseRepresentationLockBckRetention = map[string]interface{}{
		"compartment_id":                       acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compute_model":                        acctest.Representation{RepType: acctest.Required, Create: `ECPU`},
		"compute_count":                        acctest.Representation{RepType: acctest.Required, Create: `4.0`},
		"data_storage_size_in_tbs":             acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_name":                              acctest.Representation{RepType: acctest.Required, Create: adbName},
		"admin_password":                       acctest.Representation{RepType: acctest.Required, Create: `BEstrO0ng_#11`, Update: `BEstrO0ng_#12`},
		"db_version":                           acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_autonomous_db_versions.test_autonomous_db_versions.autonomous_db_versions.0.version}`},
		"db_workload":                          acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"character_set":                        acctest.Representation{RepType: acctest.Optional, Create: `AL32UTF8`},
		"display_name":                         acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`},
		"freeform_tags":                        acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_auto_scaling_enabled":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_auto_scaling_for_storage_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":                         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_mtls_connection_required":          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_backup_retention_locked":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"autonomous_maintenance_schedule_type": acctest.Representation{RepType: acctest.Optional, Create: `REGULAR`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"customer_contacts":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousDatabaseCustomerContactsRepresentation},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"whitelisted_ips":            acctest.Representation{RepType: acctest.Optional, Create: []string{`1.1.1.1/28`}},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"timeouts":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: autonomousDatabaseTimeoutsRepresentation},
		"ncharacter_set":             acctest.Representation{RepType: acctest.Optional, Create: `AL16UTF16`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatabaseAutonomousDatabaseResourceDependenciesLockBckRetention = DefinedTagsDependencies + KeyResourceDependencyConfigDbaas +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))
)

func TestDatabaseAutonomousDatabaseResource_lock_backup_retention(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousDatabaseResource_lock_backup_retention")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	okvSecret = utils.GetEnvSettingWithBlankDefault("okv_secret")
	OkvSecretVariableStr = fmt.Sprintf("variable \"okv_secret\" { default = \"%s\" }\n", okvSecret)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAutonomousDatabaseResourceDependenciesLockBckRetention+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRepresentationLockBckRetention), "database", "autonomousDatabase", t)

	acctest.ResourceTest(t, testAccCheckDatabaseAutonomousDatabaseDestroy, []resource.TestStep{
		//0. Verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, DatabaseAutonomousDatabaseRepresentationLockBckRetention),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compute_model", "ECPU"),
				resource.TestCheckResourceAttr(resourceName, "is_backup_retention_locked", "false"),
			),
		},

		//1. update backup retention from false to true
		{
			Config: config + compartmentIdVariableStr + DatabaseAutonomousDatabaseResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithRemovedProperties(acctest.RepresentationCopyWithNewProperties(DatabaseAutonomousDatabaseRepresentationLockBckRetention, map[string]interface{}{
						"is_backup_retention_locked": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
					}), []string{"admin_password", "customer_contacts", "freeform_tags", "display_name"})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "is_backup_retention_locked", "true"),
			),
		},
	})
}
