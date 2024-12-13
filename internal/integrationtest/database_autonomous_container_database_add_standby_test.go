// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
		"is_automatic_failover_enabled":                     acctest.Representation{RepType: acctest.Required, Create: `true`},
		"fast_start_fail_over_lag_limit_in_seconds":         acctest.Representation{RepType: acctest.Required, Create: `0`},
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
	AdbdDgSetupDependencies = DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig

	AdbccDgSetupDependencies = ExaccAddStandbyACDWithDataGuardResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ACDatabaseRepresentation), map[string]interface{}{
			"lifecycle":     acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
			"freeform_tags": acctest.Representation{RepType: acctest.Optional, Update: map[string]string{"Department": "Accounting"}},
			"defined_tags":  acctest.Representation{RepType: acctest.Optional, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		}))

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

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdbdAutonomousContainerDatabaseAddStandbyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation), "database", "autonomousContainerDatabaseAddStandby", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AdbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Required, acctest.Create, DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "fast_start_fail_over_lag_limit_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		{
			Config: config + compartmentIdVariableStr + AdbdDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, DatabaseAdbdAutonomousContainerDatabaseAddStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "fast_start_fail_over_lag_limit_in_seconds", "30"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_cloud_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
			),
		},
	})
}

func TestDatabaseAdbccAutonomousContainerDatabaseAddStandbyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAdbccAutonomousContainerDatabaseAddStandbyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_add_standby.test_autonomous_container_database_add_standby"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseAdbdAutonomousContainerDatabaseAddStandbyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation), "database", "autonomousContainerDatabaseAddStandby", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AdbccDgSetupDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_add_standby", "test_autonomous_container_database_add_standby", acctest.Optional, acctest.Create, DatabaseAdbccAutonomousContainerDatabaseAddStandbyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "FirstStandby"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.id"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AdbccDgSetupDependencies,
		},
	})
}
