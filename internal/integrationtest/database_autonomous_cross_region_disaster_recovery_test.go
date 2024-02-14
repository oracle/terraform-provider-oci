package integrationtest

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	drTypeADG         = "ADG"
	drTypeBackupBased = "BACKUP_BASED"

	remoteDisasterRecoveryConfigurationRepresentationADG = map[string]interface{}{
		"disaster_recovery_type": acctest.Representation{RepType: acctest.Optional, Update: `ADG`},
	}

	remoteDisasterRecoveryConfigurationRepresentationBackupBased = map[string]interface{}{
		"disaster_recovery_type": acctest.Representation{RepType: acctest.Optional, Update: `BACKUP_BASED`},
	}

	drPrimaryAutonomousDatabaseRepresentation = map[string]interface{}{
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":           acctest.Representation{RepType: acctest.Required, Create: `1`},
		"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Required, Create: `1`},
		"db_version":               acctest.Representation{RepType: acctest.Optional, Create: `19c`},
		"db_workload":              acctest.Representation{RepType: acctest.Optional, Create: `OLTP`},
		"display_name":             acctest.Representation{RepType: acctest.Optional, Create: `example_autonomous_database`, Update: `displayName2`},
		"is_auto_scaling_enabled":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_dedicated":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_preview_version_with_service_terms_accepted": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"kms_key_id":                 acctest.Representation{RepType: acctest.Optional, Create: ``},
		"license_model":              acctest.Representation{RepType: acctest.Optional, Create: `BRING_YOUR_OWN_LICENSE`},
		"vault_id":                   acctest.Representation{RepType: acctest.Optional, Create: ``},
		"operations_insights_status": acctest.Representation{RepType: acctest.Optional, Create: `NOT_ENABLED`, Update: `ENABLED`},
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `CROSS_REGION_DISASTER_RECOVERY`},
	}

	drStandbyAutonomousDatabaseRepresentation = map[string]interface{}{}

	StandbyDrAutonomousDatabaseResourceDependencies = acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseCrossRegionDisasterRecovery_basic(t *testing.T) {
	//Storing region and source_region for reset after test completes
	initialRegion := os.Getenv("TF_VAR_region")
	initialSourceRegion := os.Getenv("TF_VAR_source_region")

	//Hard coding the region as fake CrossRegion standbys can only be created in ashburn, make them to env variable if in future other regions are also supported
	os.Setenv("TF_VAR_region", "us-ashburn-1")
	os.Setenv("TF_VAR_source_region", "us-phoenix-1")

	currentRegion := os.Getenv("TF_VAR_region")
	sourceRegion := os.Getenv("TF_VAR_source_region")

	if currentRegion != "us-ashburn-1" || sourceRegion == "" {
		t.Skip("Skipping TestDatabaseCrossRegionDisasterRecovery_basic test.\n" +
			"Current TF_VAR_region=" + currentRegion + ", expected us-ashburn-1.\n" +
			"Current TF_VAR_source_region=" + sourceRegion + ", expected not to be empty.")
	}

	httpreplay.SetScenario("TestDatabaseCrossRegionDisasterRecovery_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()
	isSnapshotStandbyT := new(bool)
	isSnapshotStandbyF := new(bool)
	*isSnapshotStandbyT = true
	*isSnapshotStandbyF = false
	isReplicateBackupsEnabledT := new(bool)
	isReplicateBackupsEnabledF := new(bool)
	*isReplicateBackupsEnabledT = true
	*isReplicateBackupsEnabledF = false

	err := createPrimaryAdbInProvidedRegion(sourceRegion)
	if err != nil {
		t.Fatalf("Unable to create cross region primary ADB. Error: %v", err)
	}

	drStandbyAutonomousDatabaseRepresentation = acctest.RepresentationCopyWithNewProperties(
		drPrimaryAutonomousDatabaseRepresentation,
		map[string]interface{}{
			"source_id":                     acctest.Representation{RepType: acctest.Optional, Create: primaryId},
			"db_name":                       acctest.Representation{RepType: acctest.Optional, Create: primaryDbName},
			"remote_disaster_recovery_type": acctest.Representation{RepType: acctest.Optional, Create: `ADG`},
		})

	standbyAutonomousDatabaseConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, drStandbyAutonomousDatabaseRepresentation)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StandbyDrAutonomousDatabaseResourceDependencies+standbyAutonomousDatabaseConfig, "database", "autonomousDatabase", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//0. create dependencies
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies,
			},
			//1. create standby adb ADG
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"remote_disaster_recovery_configuration": acctest.Representation{RepType: acctest.Optional, Create: remoteDisasterRecoveryConfigurationRepresentationADG},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "1"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "1"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.disaster_recovery_type", drTypeADG),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			//2. Update Primary
			//Updating primary will also update standby. No API calls should be made. We are just checking if the refreshed state is the same as the updated config
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := updatePrimaryAdbInProvidedRegion(sourceRegion)
					if err != nil {
						t.Fatalf("Unable to update cross region primary ADB. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":           acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `2`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.disaster_recovery_type", drTypeADG),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//3. Updating Standby to new value
			//Changing standby Disaster Recovery Configuration from ADG to Backup Based.
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := triggerStandbyChangeDisasterRecoveryConfiguration(resId2, currentRegion, oci_database.ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased)
					if err != nil {
						t.Fatalf("Unable to update new values of cross region standby ADB. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId2, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":                         acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"remote_disaster_recovery_configuration": acctest.Representation{RepType: acctest.Optional, Update: drTypeBackupBased},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.disaster_recovery_type", drTypeBackupBased),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "BACKUP_COPY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//4. Updating Standby back to old value
			//Changing standby Disaster Recovery Configuration from Backup Based to ADG.
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := triggerStandbyChangeDisasterRecoveryConfiguration(resId2, currentRegion, oci_database.ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg)
					if err != nil {
						t.Fatalf("Unable to update back to old values of cross region standby ADB. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":                         acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"remote_disaster_recovery_configuration": acctest.Representation{RepType: acctest.Optional, Update: drTypeADG},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.disaster_recovery_type", drTypeADG),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//5. Converting regular standby to snapshot standby
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := triggerChangeSnapshotStandby(resId2, currentRegion, isSnapshotStandbyT)
					if err != nil {
						t.Fatalf("Unable to convert to snapshot standby. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":           acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `2`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "SNAPSHOT_STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//6. Converting snapshot standby back to standby
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := triggerChangeSnapshotStandby(resId2, currentRegion, isSnapshotStandbyF)
					if err != nil {
						t.Fatalf("Unable to convert from snapshot standby to standby. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":           acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `2`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//7. Delete standby
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies,
			},
			// replicate backups for cross region dr tests
			//8. enable adg with replicate backups as true
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies,
			},
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"is_replicate_automatic_backups": acctest.Representation{RepType: acctest.Optional, Create: `true`},
							"cpu_core_count":                 acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":       acctest.Representation{RepType: acctest.Optional, Create: `2`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.disaster_recovery_type", drTypeADG),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.is_replicate_automatic_backups", "true"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			//9. disable replicate backups on standby
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := triggerReplicateBackupsStandby(resId, currentRegion, isReplicateBackupsEnabledF)
					if err != nil {
						t.Fatalf("Unable to disable replicate standby backups. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(drStandbyAutonomousDatabaseRepresentation, map[string]interface{}{
							"cpu_core_count":           acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `2`},
						})),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "state", "AVAILABLE"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "cpu_core_count", "2"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "db_name", primaryDbName),
					resource.TestCheckResourceAttr(resourceName, "db_version", "19c"),
					resource.TestCheckResourceAttr(resourceName, "db_workload", "OLTP"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "example_autonomous_database"),
					resource.TestCheckResourceAttr(resourceName, "license_model", "BRING_YOUR_OWN_LICENSE"),
					resource.TestCheckResourceAttr(resourceName, "is_preview_version_with_service_terms_accepted", "false"),
					resource.TestCheckResourceAttr(resourceName, "remote_disaster_recovery_configuration.0.is_replicate_automatic_backups", "false"),
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DISASTER_RECOVERY"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			//10. Delete this standby
			{
				Config: config + compartmentIdVariableStr + StandbyDrAutonomousDatabaseResourceDependencies,
			},
		},
	})

	deletePrimaryAdb()

	//Resetting region and source_region after test
	os.Setenv("TF_VAR_region", initialRegion)
	os.Setenv("TF_VAR_source_region", initialSourceRegion)
}

func createPrimaryAdbInProvidedRegion(region string) error {
	var err error
	primaryId, primaryDbName, err = createAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), region)
	if err != nil {
		log.Printf("[WARN] failed to create cross region primary ADB with the error %v", err)
		return err
	}
	return nil
}

func updatePrimaryAdbInProvidedRegion(region string) error {
	_, err := updateAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), region, primaryId)
	if err != nil {
		log.Printf("[WARN] failed to update cross region primary ADB with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}

func triggerStandbyChangeDisasterRecoveryConfiguration(standbyId string, currentRegion string, disasterRecoveryType oci_database.ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum) error {
	_, err := changeDisasterRecoveryConfiguration(acctest.GetTestClients(&schema.ResourceData{}), currentRegion, standbyId, disasterRecoveryType)
	if err != nil {
		log.Printf("[WARN] failed to change cross region disaster recovery configuration of standby ADB with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}

func triggerChangeSnapshotStandby(standbyId string, currentRegion string, isSnapshotStandby *bool) error {
	_, err := changeSnapshotStandby(acctest.GetTestClients(&schema.ResourceData{}), currentRegion, standbyId, isSnapshotStandby)
	if err != nil {
		log.Printf("[WARN] failed to connect/disconnect with snapshot standby with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}

func triggerReplicateBackupsStandby(standbyId string, currentRegion string, isReplicateBackupsEnabled *bool) error {
	_, err := replicateBackupsStandby(acctest.GetTestClients(&schema.ResourceData{}), currentRegion, standbyId, isReplicateBackupsEnabled)
	if err != nil {
		log.Printf("[WARN] failed to enable/disable replicate auto backups for standby with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}
