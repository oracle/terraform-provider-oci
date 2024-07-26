// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseDatabaseAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
		"peer_cloud_autonomous_vm_cluster_id":               acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_cloud_autonomous_vm_cluster.peer_cloud_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Optional, Create: `StandbyAcd`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseAutonomousContainerDatabaseDataguardAssociationUpdateRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Update: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Update: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"fast_start_fail_over_lag_limit_in_seconds":              acctest.Representation{RepType: acctest.Optional, Update: `30`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Update: `true`},
		"protection_mode":                                        acctest.Representation{RepType: acctest.Optional, Update: `MAXIMUM_PERFORMANCE`},
	}

	DatabaseAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_dataguard_association_id}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_id}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `SNAPSHOT_STANDBY`},
		})

	DatabaseAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_dataguard_association_id}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_id}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `STANDBY`},
		})

	DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdateRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"fast_start_fail_over_lag_limit_in_seconds":              acctest.Representation{RepType: acctest.Optional, Create: `30`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"protection_mode":                                        acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
	}

	DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdateRepresentation2 = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"fast_start_fail_over_lag_limit_in_seconds":         acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"peer_autonomous_container_database_backup_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAddStandbyAutonomousContainerDatabaseBackupConfigRepresentation},
		"peer_autonomous_vm_cluster_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_db_unique_name":                               acctest.Representation{RepType: acctest.Optional, Create: acbDBName2},
		"standby_maintenance_buffer_in_days":                acctest.Representation{RepType: acctest.Optional, Create: `7`},
		"is_automatic_failover_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Optional, Create: `StandbyAcd`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2Representation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		"protection_mode":                                        acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
	}

	DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew = map[string]interface{}{
		"autonomous_container_database_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                     acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `false`},
		"protection_mode":                                   acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`, Update: `MAXIMUM_AVAILABILITY`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Optional, Create: `StandbyAcd`},
		"peer_autonomous_vm_cluster_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `SNAPSHOT_STANDBY`},
		})

	DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfigNew = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_dataguard_association_id}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_id}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `SNAPSHOT_STANDBY`},
		})

	DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_dataguard_association_id"]}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["peer_autonomous_container_database_id"]}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `STANDBY`},
		})

	DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfigNew = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_role_change", "test_oci_database_autonomous_container_database_dataguard_role_change", acctest.Optional, acctest.Create,
		map[string]interface{}{
			"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_dataguard_association_id}`},
			"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association.peer_autonomous_container_database_id}`},
			"connection_strings_type":                                acctest.Representation{RepType: acctest.Optional, Create: `PRIMARY_SERVICES`},
			"role":                                                   acctest.Representation{RepType: acctest.Optional, Create: `STANDBY`},
		})

	DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(DatabaseAutonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
		})) +
		DatabaseCloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "peer_cloud_exadata_infrastructure", acctest.Required, acctest.Create, PeerCeiRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, ATPDCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "peer_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, PeerCloudAvmRepresentation)

	ignoreDataguardChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`peer_autonomous_container_database_compartment_id`,
			`peer_autonomous_container_database_display_name`,
			`peer_autonomous_exadata_infrastructure_id`,
			`peer_autonomous_vm_cluster_id`,
			`peer_cloud_autonomous_vm_cluster_id`,
			`peer_db_unique_name`,
			`service_level_agreement_type`,
			`protection_mode`,
			`peer_autonomous_container_database_backup_config`}},
	}

	ATPDCloudAutonomousVmClusterRepresentation = acctest.RepresentationCopyWithRemovedProperties(DatabaseCloudAutonomousVmClusterRepresentation, []string{"nsg_ids"})
	PeerCloudAvmRepresentation                 = acctest.GetUpdatedRepresentationCopy("cloud_exadata_infrastructure_id", acctest.Representation{RepType: acctest.Required, Create: `${oci_database_cloud_exadata_infrastructure.peer_cloud_exadata_infrastructure.id}`}, ATPDCloudAutonomousVmClusterRepresentation)

	ExaCCACDResourceDependencies = DatabaseAVMClusterWithSingleNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseOCPUAutonomousVmClusterRepresentation)

	peerExadataInfraNewProperties = map[string]interface{}{
		"activation_file": acctest.Representation{RepType: acctest.Required, Create: activationFilePath},
		"display_name":    acctest.Representation{RepType: acctest.Required, Create: `PeerExadataInfra`},
	}
	peerExadataInfraRepresentation = acctest.RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, peerExadataInfraNewProperties)

	peerAutonomousVmClusterNewProperties = map[string]interface{}{
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `peerAutonomousVmCluster`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
		"is_local_backup_enabled":   acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"vm_cluster_network_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.peer_vm_cluster_network.id}`},
	}

	peerAutonomousVmClusterRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseOCPUAutonomousVmClusterRepresentation, peerAutonomousVmClusterNewProperties)

	ExaccACDWithDataGuardResourceDependencies = ExaCCACDResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "peer_exadata_infrastructure", acctest.Required, acctest.Create, peerExadataInfraRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "peer_vm_cluster_network", acctest.Required, acctest.Create, DatabasePeerVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "peer_autonomous_vm_cluster", acctest.Required, acctest.Create, peerAutonomousVmClusterRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "peer_db_servers", acctest.Required, acctest.Create, DatabaseDatabasePeerExaInfraDbServerDataSourceRepresentation)

	ExaccAddStandbyACDWithDataGuardResourceDependencies = ExaCCACDResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "peer_exadata_infrastructure", acctest.Required, acctest.Create, peerExadataInfraRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "peer_vm_cluster_network", acctest.Required, acctest.Create, DatabasePeerVmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "peer_autonomous_vm_cluster", acctest.Required, acctest.Create, peerAutonomousVmClusterRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "peer_db_servers", acctest.Required, acctest.Create, DatabaseDatabasePeerExaInfraDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", acctest.Optional, acctest.Create, backupDestinationNFSRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_key_store", "test_key_store", acctest.Optional, acctest.Create, DatabaseKeyStoreRepresentation) +
		KmsVaultIdVariableStr + OkvSecretVariableStr

	ExaccACDWithDataGuardRepresentation = map[string]interface{}{
		"display_name":                 acctest.Representation{RepType: acctest.Required, Create: `ACD-DG-TF-TEST`},
		"patch_model":                  acctest.Representation{RepType: acctest.Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                acctest.RepresentationGroup{RepType: acctest.Required, Group: acdBackupConfigLocalRepresentation},
		"compartment_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               acctest.Representation{RepType: acctest.Optional, Create: dgDbUniqueName},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `AUTONOMOUS_DATAGUARD`},
		"peer_autonomous_container_database_backup_config":  acctest.RepresentationGroup{RepType: acctest.Optional, Group: acdBackupConfigLocalRepresentation},
		"peer_autonomous_container_database_compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"peer_autonomous_vm_cluster_id":                     acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_display_name":   acctest.Representation{RepType: acctest.Optional, Create: `PEER-ACD-DG`},
		"protection_mode":               acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
		"is_automatic_failover_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig = ExaccACDWithDataGuardResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
			acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDataGuardRepresentation))

	ExaccACDFSFOResourceConfig = ExaccACDWithDataGuardResourceDependencies + acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(ExaccACDWithDataGuardRepresentation, map[string]interface{}{
			"is_automatic_failover_enabled": acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
		}))

	DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(AddStandbyACDatabaseRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
		})) + ExaccAddStandbyACDWithDataGuardResourceDependencies
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"
	acctest.SaveConfigContent("", "", "", t)

	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})

	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)

	DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(AutonomousContainerDatabaseDedicatedRepresentation, []string{"vault_id", "kms_key_id"}), map[string]interface{}{
			"service_level_agreement_type": acctest.Representation{RepType: acctest.Optional, Create: `STANDARD`},
			"protection_mode":              acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_PERFORMANCE`},
			"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDataguardChangesRep},
		})) +
		DatabaseCloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "peer_cloud_exadata_infrastructure", acctest.Required, acctest.Create, PeerCeiRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, ATPDCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "peer_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, PeerCloudAvmRepresentation)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		// verify create datasource
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return nil
				},
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_PERFORMANCE"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		// verify updates data guard association
		{
			Config: config +
				// acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		// convert to snapshot standby
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig + DatabaseAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
		},

		// verify convert to snapshot standby result
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig + DatabaseAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "SNAPSHOT_STANDBY"),
			),
		},

		// convert to physical standby
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig + DatabaseAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
		},

		// verify convert to physical standby result
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAutonomousContainerDatabaseDataguardAssociationRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig + DatabaseAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
			),
		},
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"
	resourceName := "oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//exacc dg ds
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.peer_role"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.protection_mode"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.role"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.time_created"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "role"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},

		// verify Create DG Association (add standby)
		{
			Config: config + compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdateRepresentation2),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.internet_proxy", "internetProxy"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.type", "NFS"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_password", "vpcPassword"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.backup_destination_details.0.vpc_user", "bkupUser1"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_backup_config.0.recovery_window_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_compartment_id"),
				resource.TestCheckResourceAttr(resourceName, "peer_autonomous_container_database_display_name", "StandbyAcd"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_vm_cluster_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttr(resourceName, "standby_maintenance_buffer_in_days", "7"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		// verify updates data guard association with optional parameters
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew) +
				compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
			),
		},

		// convert to snapshot standby
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew) +
				compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 + DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfigNew,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
		},

		// verify convert to snapshot standby result
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew) +
				compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 + DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangeSnapshotResourceConfigNew,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "SNAPSHOT_STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
			),
		},

		// convert to physical standby
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew) +
				compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 + DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfigNew,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					time.Sleep(1 * time.Minute)
					return nil
				},
			),
		},

		// verify convert to physical standby result
		{
			Config: config +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseExaccAutonomousContainerDatabaseDataguardAssociationUpdate2RepresentationNew) +
				compartmentIdVariableStr + DatabaseAddStandbyAutonomousContainerExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig1 + DatabaseExaccAutonomousContainerDatabaseDataguardRoleChangePhysicalResourceConfigNew,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttr(resourceName, "peer_role", "STANDBY"),
				resource.TestCheckResourceAttr(resourceName, "role", "PRIMARY"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "protection_mode", "MAXIMUM_AVAILABILITY"),
			),
		},
	})
}
