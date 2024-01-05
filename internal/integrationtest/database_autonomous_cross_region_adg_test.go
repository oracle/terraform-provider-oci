// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	primaryDbName string

	primaryAutonomousDatabaseRepresentation = map[string]interface{}{
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
		"source":                     acctest.Representation{RepType: acctest.Optional, Create: `CROSS_REGION_DATAGUARD`},
	}

	standbyAutonomousDatabaseRepresentationCrossRegionAdg = map[string]interface{}{}
	primaryId                                             string
	sourceRegion                                          = utils.GetEnvSettingWithBlankDefault("source_region")

	StandbyAutonomousDatabaseResourceDependencies = KeyResourceDependencyConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_db_versions", acctest.Required, acctest.Create, DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_db_versions", "test_autonomous_dw_versions", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousDbVersionDataSourceRepresentation, map[string]interface{}{
				"db_workload": acctest.Representation{RepType: acctest.Required, Create: `DW`}}))
)

// issue-routing-tag: database/dbaas-adb
func TestDatabaseCrossRegionAdg_basic(t *testing.T) {

	//	Fake CrossRegion standbys can only be created in ashburn
	currentRegion := utils.GetEnvSettingWithBlankDefault("region")
	if currentRegion != "us-ashburn-1" || sourceRegion == "" {
		t.Skip("Skipping TestDatabaseCrossRegionAdg_basic test.\n" +
			"Current TF_VAR_region=" + currentRegion + ", expected us-ashburn-1.\n" +
			"Current TF_VAR_source_region=" + sourceRegion + ", expected not to be empty.")
	}

	httpreplay.SetScenario("TestDatabaseCrossRegionAdg_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	err := createPrimaryAdb()
	if err != nil {
		t.Fatalf("Unable to create cross region primary ADB. Error: %v", err)
	}

	standbyAutonomousDatabaseRepresentationCrossRegionAdg = acctest.RepresentationCopyWithNewProperties(
		primaryAutonomousDatabaseRepresentation,
		map[string]interface{}{
			"source_id": acctest.Representation{RepType: acctest.Optional, Create: primaryId},
			"db_name":   acctest.Representation{RepType: acctest.Optional, Create: primaryDbName},
		})
	standbyAutonomousDatabaseConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create, standbyAutonomousDatabaseRepresentationCrossRegionAdg)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_database.test_autonomous_database"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+StandbyAutonomousDatabaseResourceDependencies+standbyAutonomousDatabaseConfig, "database", "autonomousDatabase", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// 0 create dependencies
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies,
			},
			// 1 create standby adb
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies + standbyAutonomousDatabaseConfig,
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
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
			// 2 Update Primary
			//	Updating primary will also update standby. No API calls should be made. We are just checking if the refreshed state is the same as the updated config
			{
				PreConfig: func() {
					acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
						getAdbFromSourceRegion, "database", true)()
					err := updatePrimaryAdb()
					if err != nil {
						t.Fatalf("Unable to update cross region primary ADB. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(standbyAutonomousDatabaseRepresentationCrossRegionAdg, map[string]interface{}{
							"cpu_core_count":           acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs": acctest.Representation{RepType: acctest.Optional, Create: `2`},
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
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
			// 3 Switchover to Standby - Valid PeerId
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(standbyAutonomousDatabaseRepresentationCrossRegionAdg, map[string]interface{}{
							"cpu_core_count":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":     acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"switchover_to_remote_peer_id": acctest.Representation{RepType: acctest.Optional, Create: primaryId},
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),
					resource.TestCheckResourceAttrSet(resourceName, "time_data_guard_role_changed"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// 4 Switchover to Standby - Valid PeerId with different case
			//	No API calls should be made.
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(standbyAutonomousDatabaseRepresentationCrossRegionAdg, map[string]interface{}{
							"cpu_core_count":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":     acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"switchover_to_remote_peer_id": acctest.Representation{RepType: acctest.Optional, Create: strings.ToUpper(primaryId)},
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),
					resource.TestCheckResourceAttrSet(resourceName, "time_data_guard_role_changed"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// 5 Switchover to Standby - Empty string
			//	No API calls should be made. Simply checking if the refreshed state has switchover_to_remote_peer_id as ""
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(standbyAutonomousDatabaseRepresentationCrossRegionAdg, map[string]interface{}{
							"cpu_core_count":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":     acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"switchover_to_remote_peer_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),
					resource.TestCheckResourceAttrSet(resourceName, "time_data_guard_role_changed"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// 6 Switchover back to Primary
			//	Switched-over cross region standbys cannot be terminated. Hence we need to switch-over back so that we can delete this standby
			//	No API calls should be made. Simply checking if the refreshed state has role as STANDBY
			{
				PreConfig: func() {
					err := switchoverToPrimaryAdb(resId)
					if err != nil {
						t.Fatalf("Unable to switchover to primary adb. Error: %v", err)
					}
					acctest.WaitTillCondition(acctest.TestAccProvider, &resId, adbWaitTillLifecycleStateStandbyCondition, 10*time.Minute,
						getAdbFromCurrentRegion, "database", true)()
				},
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_database", "test_autonomous_database", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(standbyAutonomousDatabaseRepresentationCrossRegionAdg, map[string]interface{}{
							"cpu_core_count":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"data_storage_size_in_tbs":     acctest.Representation{RepType: acctest.Optional, Create: `2`},
							"switchover_to_remote_peer_id": acctest.Representation{RepType: acctest.Optional, Create: ``},
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
					resource.TestCheckResourceAttr(resourceName, "source", "CROSS_REGION_DATAGUARD"),
					resource.TestCheckResourceAttr(resourceName, "source_id", primaryId),
					resource.TestCheckResourceAttr(resourceName, "dataguard_region_type", "REMOTE_STANDBY_DG_REGION"),
					resource.TestCheckResourceAttr(resourceName, "role", "STANDBY"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "peer_db_ids.0", primaryId),
					resource.TestCheckResourceAttrSet(resourceName, "time_data_guard_role_changed"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// 7 Delete standby
			{
				Config: config + compartmentIdVariableStr + StandbyAutonomousDatabaseResourceDependencies,
			},
		},
	})
	deletePrimaryAdb()
}

func deletePrimaryAdb() error {
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()

	err := deleteAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, primaryId)
	if err != nil {
		log.Printf("[WARN] failed to delete cross region primary ADB with the error %v", err)
		return err
	}
	return nil
}

func createPrimaryAdb() error {
	var err error
	primaryId, primaryDbName, err = createAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion)
	if err != nil {
		log.Printf("[WARN] failed to create cross region primary ADB with the error %v", err)
		return err
	}
	return nil
}

func updatePrimaryAdb() error {
	_, err := updateAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, primaryId)
	if err != nil {
		log.Printf("[WARN] failed to update cross region primary ADB with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}

func switchoverToPrimaryAdb(standbyId string) error {
	err := triggerSwitchoverOnAdbInRegion(acctest.GetTestClients(&schema.ResourceData{}), sourceRegion, primaryId, standbyId)
	if err != nil {
		log.Printf("[WARN] failed to switchover to primary ADB with the error %v", err)
		return err
	}
	acctest.WaitTillCondition(acctest.TestAccProvider, &primaryId, adbWaitTillLifecycleStateAvailableCondition, 10*time.Minute,
		getAdbFromSourceRegion, "database", true)()
	return nil
}
