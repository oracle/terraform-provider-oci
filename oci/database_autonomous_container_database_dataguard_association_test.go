// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": Representation{repType: Required, create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       Representation{repType: Required, create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation = map[string]interface{}{
		"autonomous_container_database_id": Representation{repType: Required, create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
	}

	AutonomousContainerDatabaseDataguardAssociationResourceConfig = generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", Optional, Create,
		representationCopyWithNewProperties(representationCopyWithRemovedProperties(autonomousContainerDatabaseRepresentation, []string{"vault_id", "kms_key_id", "peer_autonomous_container_database_backup_config", "peer_autonomous_container_database_compartment_id", "peer_autonomous_vm_cluster_id"}), map[string]interface{}{
			"service_level_agreement_type":              Representation{repType: Optional, create: `AUTONOMOUS_DATAGUARD`},
			"protection_mode":                           Representation{repType: Optional, create: `MAXIMUM_AVAILABILITY`},
			"peer_autonomous_exadata_infrastructure_id": Representation{repType: Optional, create: `${oci_database_autonomous_exadata_infrastructure.peer_autonomous_exadata_infrastructure.id}`},
		})) +
		AutonomousExadataInfrastructureResourceConfig +
		generateResourceFromRepresentationMap("oci_database_autonomous_exadata_infrastructure", "peer_autonomous_exadata_infrastructure", Optional, Create, autonomousExadataInfrastructureRepresentation)

	ExaCCACDResourceDependencies = generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Required, Create,
		representationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": Representation{repType: Required, create: activationFilePath}})) +
		generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", Required, Create, autonomousVmClusterRepresentation) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Required, Create,
			representationCopyWithNewProperties(vmClusterNetworkRepresentation, map[string]interface{}{"validate_vm_cluster_network": Representation{repType: Required, create: "true"}}))

	peerExadataInfraNewProperties = map[string]interface{}{
		"activation_file": Representation{repType: Required, create: activationFilePath},
		"display_name":    Representation{repType: Required, create: `PeerExadataInfra`},
	}
	peerExadataInfraRepresentation = representationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, peerExadataInfraNewProperties)

	peerVmClusterNetworkNewProperties = map[string]interface{}{
		"validate_vm_cluster_network": Representation{repType: Required, create: "true"},
		"display_name":                Representation{repType: Required, create: `peerVmClusterNw`},
		"exadata_infrastructure_id":   Representation{repType: Required, create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
	}

	peerVmClusterNetworkRepresentation = representationCopyWithNewProperties(vmClusterNetworkRepresentation, peerVmClusterNetworkNewProperties)

	peerAutonomousVmClusterNewProperties = map[string]interface{}{
		"display_name":              Representation{repType: Required, create: `peerAutonomousVmCluster`},
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
		"is_local_backup_enabled":   Representation{repType: Optional, create: `true`},
		"vm_cluster_network_id":     Representation{repType: Required, create: `${oci_database_vm_cluster_network.peer_vm_cluster_network.id}`},
	}

	peerAutonomousVmClusterRepresentation = representationCopyWithNewProperties(autonomousVmClusterRepresentation, peerAutonomousVmClusterNewProperties)

	ExaccACDWithDataGuardResourceDependencies = ExaCCACDResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "peer_exadata_infrastructure", Required, Create, peerExadataInfraRepresentation) +
		generateResourceFromRepresentationMap("oci_database_vm_cluster_network", "peer_vm_cluster_network", Required, Create, peerVmClusterNetworkRepresentation) +
		generateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "peer_autonomous_vm_cluster", Required, Create, peerAutonomousVmClusterRepresentation) +
		DefinedTagsDependencies

	ExaccACDWithDataGuardRepresentation = map[string]interface{}{
		"display_name":                 Representation{repType: Required, create: `ACD-DG-TF-TEST`},
		"patch_model":                  Representation{repType: Required, create: `RELEASE_UPDATES`, update: `RELEASE_UPDATE_REVISIONS`},
		"autonomous_vm_cluster_id":     Representation{repType: Required, create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"backup_config":                RepresentationGroup{Required, acdBackupConfigLocalRepresentation},
		"compartment_id":               Representation{repType: Optional, create: `${var.compartment_id}`},
		"db_unique_name":               Representation{repType: Optional, create: dgDbUniqueName},
		"freeform_tags":                Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"maintenance_window_details":   RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsRepresentation},
		"service_level_agreement_type": Representation{repType: Optional, create: `AUTONOMOUS_DATAGUARD`},
		"peer_autonomous_container_database_backup_config":  RepresentationGroup{Optional, acdBackupConfigLocalRepresentation},
		"peer_autonomous_container_database_compartment_id": Representation{repType: Optional, create: `${var.compartment_id}`},
		"peer_autonomous_vm_cluster_id":                     Representation{repType: Optional, create: `${oci_database_autonomous_vm_cluster.peer_autonomous_vm_cluster.id}`},
		"peer_autonomous_container_database_display_name":   Representation{repType: Optional, create: `PEER-ACD-DG`},
		"protection_mode": Representation{repType: Optional, create: `MAXIMUM_PERFORMANCE`},
	}

	ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig = ExaccACDWithDataGuardResourceDependencies +
		generateResourceFromRepresentationMap("oci_database_autonomous_container_database", "exacc_test_autonomous_container_database", Optional, Create,
			getUpdatedRepresentationCopy("maintenance_window_details", RepresentationGroup{Optional, autonomousContainerDatabaseMaintenanceWindowDetailsNoPreferenceRepresentation}, ExaccACDWithDataGuardRepresentation))
)

func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", Optional, Create, autonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", Optional, Create, autonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + AutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}

func TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExaccAutonomousContainerDatabaseDataguardAssociationResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.exacc_test_autonomous_container_database_dataguard_association"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			//exacc dg ds
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "exacc_test_autonomous_container_database_dataguard_associations", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
					generateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "exacc_test_autonomous_container_database_dataguard_association", Optional, Create, exaccAutonomousContainerDatabaseDataguardAssociationSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ExaccAutonomousContainerDatabaseDataguardAssociationResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
		},
	})
}
