// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": Representation{RepType: Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       Representation{RepType: Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": Representation{RepType: Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       Representation{RepType: Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{RepType: Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	AutonomousContainerDatabaseDataguardAssociationResourceConfig = GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
		RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(autonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id", "peer_autonomous_container_database_backup_config", "peer_autonomous_container_database_compartment_id", "peer_autonomous_vm_cluster_id"}), map[string]interface{}{
			"service_level_agreement_type":              Representation{RepType: Optional, Create: `AUTONOMOUS_DATAGUARD`},
			"protection_mode":                           Representation{RepType: Optional, Create: `MAXIMUM_AVAILABILITY`},
			"peer_autonomous_exadata_infrastructure_id": Representation{RepType: Optional, Create: `${oci_database_autonomous_exadata_infrastructure.peer_autonomous_exadata_infrastructure.id}`},
		})) +
		AutonomousExadataInfrastructureResourceConfig +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "peer_autonomous_exadata_infrastructure", Optional, Create, autonomousExadataInfrastructureRepresentation)

	ExaCCACDResourceDependencies = GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create,
		RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": Representation{RepType: Required, Create: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create,
			RepresentationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": Representation{RepType: Required, Create: "true"}}))

	peerExadataInfraNewProperties = map[string]interface{}{
		"activation_file": Representation{RepType: Required, Create: activationFilePath},
		"display_name":    Representation{RepType: Required, Create: `PeerExadataInfra`},
	}
	peerExadataInfraRepresentation = RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, peerExadataInfraNewProperties)

	peerVmClusterNetworkNewProperties = map[string]interface{}{
		"validate_vm_cluster_network": Representation{RepType: Required, Create: "true"},
		"display_name":                Representation{RepType: Required, Create: `peerVmClusterNw`},
		"exadata_infrastructure_id":   Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
	}

	peerVmClusterNetworkRepresentation = RepresentationCopyWithNewProperties(vmClusterNetworkRepresentation, peerVmClusterNetworkNewProperties)

	peerAutonomousVmClusterNewProperties = map[string]interface{}{
		"display_name":              Representation{RepType: Required, Create: `peerAutonomousVmCluster`},
		"exadata_infrastructure_id": Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
		"is_local_backup_enabled":   Representation{RepType: Optional, Create: `true`},
		"vm_cluster_network_id":     Representation{RepType: Required, Create: `${oci_database_vm_cluster_network.peer_vm_cluster_network.id}`},
	}

	peerAutonomousVmClusterRepresentation = RepresentationCopyWithNewProperties(autonomousVmClusterRepresentation, peerAutonomousVmClusterNewProperties)

	ExaccACDWithDataGuardResourceDependencies = ExaCCACDResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "peer_exadata_infrastructure", Required, Create, peerExadataInfraRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "peer_vm_cluster_network", Required, Create, peerVmClusterNetworkRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "peer_autonomous_vm_cluster", Required, Create, peerAutonomousVmClusterRepresentation) +
		DefinedTagsDependencies

	ExaccACDWithDataGuardRepresentation = map[string]interface{}{
		"display_name":                 Representation{RepType: Required, Create: `ACD-DG-TF-TEST`},
		"patch_model":                  Representation{RepType: Required, Create: `RELEASE_UPDATES`, Update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     Representation{RepType: Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                RepresentationGroup{Required, acdBackupConfigLocalRepresentation},
		"compartment_id":               Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"db_unique_name":               Representation{RepType: Optional, Create: dgDbUniqueName},
		"freeform_tags":                Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": Representation{RepType: Optional, Create: `AUTONOMOUS_DATAGUARD`},
		"peer_autonomous_container_database_backup_config":  RepresentationGroup{Optional, acdBackupConfigLocalRepresentation},
		"peer_autonomous_container_database_compartment_id": Representation{RepType: Optional, Create: `${var.compartment_id}`},
		"peer_autonomous_vm_cluster_id":                     Representation{RepType: Optional, Create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_display_name":   Representation{RepType: Optional, Create: `PEER-ACD-DG`},
		"protection_mode": Representation{RepType: Optional, Create: `MAXIMUM_PERFORMANCE`},
	}

	ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig = ExaccACDWithDataGuardResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", Optional, Create,
			GetUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDataGuardRepresentation))
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	// sshaagar: Commenting AEI test as not needed post migration.
	// DISALLOWED_API.launchAutonomousExadataInfrastructure
	t.Skip("Skipping Test from execution on a regular test run")

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", Optional, Create, autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"

	ResourceTest(t, nil, []resource.TestStep{
		//exacc dg ds
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
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
	})
}
