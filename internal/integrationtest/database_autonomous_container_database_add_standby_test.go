// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseAutonomousContainerDatabaseAddStandbyRequiredOnlyResource = DatabaseAdbdAutonomousContainerDatabaseAddStandbyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Required, acctest.Create, DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation)

	DatabaseAddStandbyAdbdAcdRepresentation = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Required, Create: `FirstStandby`},
		"peer_cloud_autonomous_vm_cluster_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.peer_cloud_autonomous_vm_cluster.id}`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_AVAILABILITY`},
		//"peer_autonomous_container_database_backup_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAddStandbyAutonomousContainerDatabaseBackupConfigRepresentation},
	}

	// multi standby dg
	DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Required, Create: `FirstStandby`},
		"peer_cloud_autonomous_vm_cluster_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_autonomous_vm_cluster.peer_cloud_autonomous_vm_cluster.id}`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_AVAILABILITY`},
		"is_automatic_failover_enabled":                     acctest.Representation{RepType: acctest.Required, Create: `true`},
		"fast_start_fail_over_lag_limit_in_seconds":         acctest.Representation{RepType: acctest.Required, Create: `30`},
		"standby_maintenance_buffer_in_days":                acctest.Representation{RepType: acctest.Optional, Create: `7`},
	}
	DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Required, Create: `FirstStandby`},
		"peer_autonomous_vm_cluster_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Required, Create: `MAXIMUM_AVAILABILITY`},
		"standby_maintenance_buffer_in_days":                acctest.Representation{RepType: acctest.Optional, Create: `7`},
		"peer_autonomous_container_database_backup_config":  acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigRepresentation},
	}

	DatabaseAutonomousContainerDatabaseAddStandbyPeerAutonomousContainerDatabaseBackupConfigRepresentation = map[string]interface{}{
		"backup_destination_details": acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseAddStandbyPeerAutonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation},
		"recovery_window_in_days":    acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}
	DatabaseAutonomousContainerDatabaseAddStandbyPeerAutonomousContainerDatabaseBackupConfigBackupDestinationDetailsRepresentation = map[string]interface{}{
		"type":           acctest.Representation{RepType: acctest.Required, Create: `NFS`},
		"dbrs_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_identity_policy.test_policy.id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `id`},
		"internet_proxy": acctest.Representation{RepType: acctest.Optional, Create: `internetProxy`},
		"vpc_password":   acctest.Representation{RepType: acctest.Optional, Create: `vpcPassword`},
		"vpc_user":       acctest.Representation{RepType: acctest.Optional, Create: `vpcUser`},
	}

	// ACD with peer dependencies
	AdbccDgSetupDependencies = getAdbccDgSetupDependency()

	DatabaseAdbdAutonomousContainerDatabaseAddStandbyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Required, acctest.Create, DatabaseAutonomousContainerDatabaseRepresentation) +
		//AutonomousExadataInfrastructureResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "test_cloud_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseCloudExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, DatabaseExadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_identity_policy", "test_policy", acctest.Required, acctest.Create, IdentityPolicyRepresentation)
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseAddStandbyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseAddStandbyResource_basic")
	defer httpreplay.SaveScenario()

	dbName := utils.RandomStringOrHttpReplayValue(8, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "STANDBY1")
	adbdDgSetupDependencies := getAdbdAddStandbyDgSetupDependencies(dbName)

	config := acctest.ProviderTestConfig() + `
provider "oci" {
	ignore_defined_tags = [
		"Oracle-Tags.CreatedBy",
		"Oracle-Tags.CreatedOn",
	]
}
`

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	addStandbyRepresentation := acctest.RepresentationCopyWithNewProperties(DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation, map[string]interface{}{
		"is_automatic_failover_enabled": acctest.Representation{RepType: acctest.Required, Create: `false`},
	})
	standbyRepresentation := acctest.RepresentationCopyWithNewProperties(AdbdStandbyACDRepresentation, map[string]interface{}{
		"depends_on": acctest.Representation{RepType: acctest.Required, Create: []string{resourceName}},
	})
	standbyConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create, standbyRepresentation)
	standbyDeleteConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(standbyRepresentation, map[string]interface{}{
			"count": acctest.Representation{RepType: acctest.Required, Create: `0`},
		}))

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdbdAutonomousContainerDatabaseAddStandbyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation), "database", "autonomousContainerDatabaseAddStandby", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + adbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Required, acctest.Create, addStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "fast_start_fail_over_lag_limit_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
			),
		},
		// Refresh the Add Standby state before using its Data Guard group member to import the standby ACD.
		{
			RefreshState: true,
		},
		// Import the standby ACD created by Add Standby so Terraform can clean it up explicitly.
		{
			Config:             config + compartmentIdVariableStr + adbdDgSetupDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation) + standbyConfig,
			ImportState:        true,
			ImportStateIdFunc:  getStandbyAcdOcid(resourceName),
			ImportStatePersist: true,
			ResourceName:       standbyResourceName,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
			),
		},
		// Adopt the imported standby into the normal configuration. Terraform
		// import does not persist configuration dependencies in state; this
		// no-op apply records them before the standby is removed.
		{
			Config: config + compartmentIdVariableStr + adbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation) +
				standbyConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
			),
		},
		// Delete the imported standby before the primary ACD and its dependencies.
		{
			// Keep the standby block with count zero so the cleanup plan targets
			// only the imported instance. The adoption step recorded its dependency
			// edges in state, ordering this delete before the primary ACD and peer
			// VM cluster can be destroyed.
			Config: config + compartmentIdVariableStr + adbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation) +
				standbyDeleteConfig,
		},
		{
			Config: config + compartmentIdVariableStr + adbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "fast_start_fail_over_lag_limit_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
			),
		},
	})
}

func getAdbdAddStandbyDgSetupDependencies(dbName string) string {
	// Cloud Autonomous VM Cluster dependencies can receive automatic Oracle-Tags
	// after creation. Scope this ignore to this test so shared tag checks stay active.
	cavmIgnoreChanges := acctest.Representation{RepType: acctest.Optional, Create: []string{
		`autonomous_data_storage_size_in_tbs`,
		`cpu_core_count_per_node`,
		`defined_tags`,
	}}
	testCavmRepresentation := acctest.GetUpdatedRepresentationCopy("lifecycle.ignore_changes", cavmIgnoreChanges, ATPDCloudAutonomousVmClusterRepresentation)
	peerCavmRepresentation := acctest.GetUpdatedRepresentationCopy("cloud_exadata_infrastructure_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.peer_cloud_exadata_infrastructure.id}`}, testCavmRepresentation)
	primaryAcdRepresentation := acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
		"db_name":                      acctest.Representation{RepType: acctest.Optional, Create: dbName},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
		"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
	})

	return acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, primaryAcdRepresentation) +
		DatabaseCloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "peer_cloud_exadata_infrastructure", acctest.Required, acctest.Create, PeerCeiRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, testCavmRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "peer_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, peerCavmRepresentation)
}

func TestDatabaseAdbccAutonomousContainerDatabaseAddStandbyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbccAutonomousContainerDatabaseAddStandbyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby"
	mainAcdResourceName := "oci_database_autonomous_container_database.test_autonomous_container_database"
	standbyResourceName := "oci_database_autonomous_container_database.standby_acd"
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))

	dependencyPrefix := getAdbccAddStandbyDependencyPrefix()
	mainAcdRepresentation := getAdbccAddStandbyMainAcdRepresentation()
	addStandbyRepresentation := getAdbccAddStandbyRepresentation()
	standbyAcdRepresentation := getAdbccStandbyAcdRepresentation()

	if simulateDb {
		acctest.PreCheck(t)
		sharedDependencyAddresses := []string{
			"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster",
			"oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster",
			"oci_database_backup_destination.test_backup_destination",
			"oci_database_key_store.test_key_store",
		}
		sharedDependencyIDs, cleanup := ResolveOrCreateSharedDependenciesFromConfig(t, map[string]string{
			"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster": utils.GetEnvSettingWithBlankDefault("autonomous_vm_cluster_id"),
			"oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster": utils.GetEnvSettingWithBlankDefault("peer_autonomous_vm_cluster_id"),
			"oci_database_backup_destination.test_backup_destination":       utils.GetEnvSettingWithBlankDefault("backup_destination_id"),
			"oci_database_key_store.test_key_store":                         utils.GetEnvSettingWithBlankDefault("key_store_id"),
		}, config+compartmentIdVariableStr+ExaccAddStandbyACDWithDataGuardResourceDependencies, sharedDependencyAddresses)
		if cleanup != nil {
			t.Cleanup(cleanup)
		}
		t.Logf("[SHARED_DEP_SETUP] autonomous_vm_cluster_id=%s | peer_autonomous_vm_cluster_id=%s | backup_destination_id=%s | key_store_id=%s",
			sharedDependencyIDs["oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"],
			sharedDependencyIDs["oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster"],
			sharedDependencyIDs["oci_database_backup_destination.test_backup_destination"],
			sharedDependencyIDs["oci_database_key_store.test_key_store"],
		)

		dependencyPrefix = getExaccTagDependency()
		mainAcdRepresentation = adbccAddStandbyAutonomousContainerDatabaseRepresentationWithSharedDependencies(
			sharedDependencyIDs["oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"],
			sharedDependencyIDs["oci_database_backup_destination.test_backup_destination"],
			sharedDependencyIDs["oci_database_key_store.test_key_store"],
		)
		addStandbyRepresentation = adbccAutonomousContainerDatabaseAddStandbyRepresentationWithSharedDependencies(
			sharedDependencyIDs["oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster"],
			sharedDependencyIDs["oci_database_backup_destination.test_backup_destination"],
		)
		standbyAcdRepresentation = adbccStandbyAcdRepresentationWithSharedDependency(sharedDependencyIDs["oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster"])
	}

	mainAcdConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, mainAcdRepresentation)
	addStandbyConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, addStandbyRepresentation)
	standbyConfig := getAdbccStandbyAcdConfig(standbyAcdRepresentation, map[string]string{})
	standbySwitchoverConfig := getAdbccStandbyAcdConfig(standbyAcdRepresentation, map[string]string{"switchover_trigger": "1"})
	standbyReinstateConfig := getAdbccStandbyAcdConfig(standbyAcdRepresentation, map[string]string{"switchover_trigger": "1", "reinstate_trigger": "1"})
	mainFailoverRepresentation := getAdbccMainAcdRepresentationWithTriggers(mainAcdRepresentation, map[string]string{"failover_trigger": "1"})
	mainFailoverConfig := acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create, mainFailoverRepresentation)

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+dependencyPrefix+mainAcdConfig+addStandbyConfig, "database", "autonomousContainerDatabaseAddStandby", t)

	t.Run("Add-Standby-Switchover-Failover-Reinstate-Acd-Delete", func(t *testing.T) {
		var addStandbyResId, mainAcdResId, standbyAcdResId string

		acctest.ResourceTest(t, nil, []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + dependencyPrefix + mainAcdConfig + addStandbyConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
					resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
					resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
					resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
					resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
					exaccMainResourceLog(t, "create ADBCC ACD", mainAcdResourceName, nil, &mainAcdResId,
						"display_name", "role", "protection_mode", "service_level_agreement_type"),
					exaccMainResourceLog(t, "create ADBCC ACD add standby", resourceName, nil, &addStandbyResId,
						"autonomous_container_database_id", "peer_autonomous_container_database_display_name", "protection_mode",
						"standby_maintenance_buffer_in_days", "peer_autonomous_vm_cluster_id",
						"peer_autonomous_container_database_backup_config.0.recovery_window_in_days"),
				),
			},
			// refresh after standby creation
			{
				RefreshState: true,
			},
			// import standby ACD created by add-standby
			{
				Config:             config + compartmentIdVariableStr + dependencyPrefix + mainAcdConfig + addStandbyConfig + standbyConfig,
				ImportState:        true,
				ImportStateIdFunc:  getStandbyAcdOcid(mainAcdResourceName),
				ImportStatePersist: true,
				ResourceName:       standbyResourceName,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
					exaccMainResourceLog(t, "import ADBCC standby ACD", standbyResourceName, nil, &standbyAcdResId,
						"display_name", "role", "autonomous_vm_cluster_id"),
				),
			},
			// switchover using standby ACD
			{
				Config: config + compartmentIdVariableStr + dependencyPrefix + mainAcdConfig + addStandbyConfig + standbySwitchoverConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(standbyResourceName, "role", "PRIMARY"),
					exaccMainResourceLog(t, "switchover ADBCC standby ACD", standbyResourceName, &standbyAcdResId, &standbyAcdResId,
						"display_name", "role", "switchover_trigger"),
					exaccMainResourceLog(t, "verify ADBCC main ACD after standby switchover", mainAcdResourceName, &mainAcdResId, &mainAcdResId,
						"display_name", "role"),
				),
			},
			// failover using original ACD
			{
				Config: config + compartmentIdVariableStr + dependencyPrefix + mainFailoverConfig + addStandbyConfig + standbySwitchoverConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					exaccMainResourceLog(t, "failover ADBCC main ACD", mainAcdResourceName, &mainAcdResId, &mainAcdResId,
						"display_name", "role", "failover_trigger"),
				),
			},
			// reinstate old primary as standby
			{
				Config: config + compartmentIdVariableStr + dependencyPrefix + mainFailoverConfig + addStandbyConfig + standbyReinstateConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(standbyResourceName, "role", "STANDBY"),
					exaccMainResourceLog(t, "reinstate ADBCC standby ACD", standbyResourceName, &standbyAcdResId, &standbyAcdResId,
						"display_name", "role", "reinstate_trigger"),
				),
			},
			// remove imported standby ACD before post-test destroy
			{
				Config: config + compartmentIdVariableStr + dependencyPrefix + mainFailoverConfig + addStandbyConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					exaccMainResourceLog(t, "delete imported ADBCC standby ACD", standbyResourceName, &standbyAcdResId, nil),
				),
			},
		})
	})
}

func adbccAddStandbyAutonomousContainerDatabaseRepresentationWithSharedDependencies(autonomousVmClusterID string, backupDestinationID string, keyStoreID string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation), map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: autonomousVmClusterID},
		"backup_config":            acctest.RepresentationGroup{RepType: acctest.Required, Group: exaccAutonomousContainerDatabaseBackupConfigWithSharedDependencies(backupDestinationID)},
		"key_store_id":             acctest.Representation{RepType: acctest.Optional, Create: keyStoreID},
		"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
		"freeform_tags":            acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"Department": "Accounting"}},
		"defined_tags":             acctest.Representation{RepType: acctest.Optional, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
	})
}

func adbccAutonomousContainerDatabaseAddStandbyRepresentationWithSharedDependencies(peerAutonomousVmClusterID string, backupDestinationID string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation, map[string]interface{}{
		"peer_autonomous_vm_cluster_id":                    acctest.Representation{RepType: acctest.Required, Create: peerAutonomousVmClusterID},
		"peer_autonomous_container_database_backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: exaccAutonomousContainerDatabaseBackupConfigWithSharedDependencies(backupDestinationID)},
	})
}

func adbccStandbyAcdRepresentationWithSharedDependency(peerAutonomousVmClusterID string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(AdbccStandbyACDRepresentation, map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: peerAutonomousVmClusterID},
	})
}

func getAdbccAddStandbyDependencyPrefix() string {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return ExaccAddStandbyACDWithDataGuardResourceDependencies
	}

	autonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("autonomous_vm_cluster_id")
	peerAutonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("peer_autonomous_vm_cluster_id")
	backupDestinationId := utils.GetEnvSettingWithBlankDefault("backup_destination_id")
	return fmt.Sprintf("variable \"autonomous_vm_cluster_id\" { default = \"%s\" }\n", autonomousVmClusterId) +
		fmt.Sprintf("variable \"backup_destination_id\" { default = \"%s\" }\n", backupDestinationId) +
		fmt.Sprintf("variable \"peer_autonomous_vm_cluster_id\" { default = \"%s\" }\n", peerAutonomousVmClusterId)
}

func getAdbccAddStandbyMainAcdRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	mainAcdRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation)
	if simulateDb {
		return acctest.RepresentationCopyWithNewProperties(mainAcdRepresentation, map[string]interface{}{
			"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"Department": "Accounting"}},
			"defined_tags":  acctest.Representation{RepType: acctest.Optional, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		})
	}

	return acctest.RepresentationCopyWithNewProperties(
		acctest.RepresentationCopyWithRemovedProperties(mainAcdRepresentation, []string{"defined_tags", "freeform_tags", "okv_end_point_group_name", "key_store_id"}),
		map[string]interface{}{
			"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.autonomous_vm_cluster_id}`},
			"backup_config":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigRealRepresentation},
			"lifecycle":                acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
		})
}

func getAdbccMainAcdRepresentationWithTriggers(baseRepresentation map[string]interface{}, triggers map[string]string) map[string]interface{} {
	properties := map[string]interface{}{}
	for name, value := range triggers {
		properties[name] = acctest.Representation{RepType: acctest.Optional, Create: value}
	}
	if len(properties) == 0 {
		return baseRepresentation
	}
	return acctest.RepresentationCopyWithNewProperties(baseRepresentation, properties)
}

func getAdbccStandbyAcdRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return AdbccStandbyACDRepresentation
	}

	return acctest.RepresentationCopyWithNewProperties(AdbccStandbyACDRepresentation, map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${var.peer_autonomous_vm_cluster_id}`},
	})
}

func getAdbccStandbyAcdConfig(baseRepresentation map[string]interface{}, triggers map[string]string) string {
	properties := map[string]interface{}{
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
	}
	for name, value := range triggers {
		properties[name] = acctest.Representation{RepType: acctest.Required, Create: value}
	}

	return acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "standby_acd", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(baseRepresentation, properties))
}

func getAdbccDgSetupDependency() string {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return ExaccAddStandbyACDWithDataGuardResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation), map[string]interface{}{
				"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
				"freeform_tags": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"Department": "Accounting"}},
				"defined_tags":  acctest.Representation{RepType: acctest.Optional, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
			}))
	} else {
		autonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("autonomous_vm_cluster_id")
		peerAutonomousVmClusterId := utils.GetEnvSettingWithBlankDefault("peer_autonomous_vm_cluster_id")
		backupDestinationId := utils.GetEnvSettingWithBlankDefault("backup_destination_id")
		return fmt.Sprintf("variable \"autonomous_vm_cluster_id\" { default = \"%s\" }\n", autonomousVmClusterId) +
			fmt.Sprintf("variable \"backup_destination_id\" { default = \"%s\" }\n", backupDestinationId) +
			fmt.Sprintf("variable \"peer_autonomous_vm_cluster_id\" { default = \"%s\" }\n", peerAutonomousVmClusterId) +
			acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
				acctest.RepresentationCopyWithNewProperties(
					acctest.RepresentationCopyWithRemovedProperties(acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation), []string{"defined_tags", "okv_end_point_group_name", "key_store_id"}),
					map[string]interface{}{
						"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.autonomous_vm_cluster_id}`},
						"backup_config":            acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigRealRepresentation},
					}))
	}
}

func getAdbccAddStandbyRepresentation() map[string]interface{} {
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))
	if simulateDb {
		return DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation
	} else {
		// Add the dynamic properties
		return acctest.RepresentationCopyWithNewProperties(DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation,
			map[string]interface{}{
				"peer_autonomous_vm_cluster_id":                    acctest.Representation{RepType: acctest.Optional, Create: `${var.peer_autonomous_vm_cluster_id}`},
				"peer_autonomous_container_database_backup_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseAutonomousContainerDatabaseBackupConfigRealRepresentation},
			})
	}
}
