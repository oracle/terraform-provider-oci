// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

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
	DatabaseAtpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	exaccAutonomousContainerDatabaseDataguardAssociationResourceRepresentation = map[string]interface{}{
		"autonomous_container_database_dataguard_association_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_container_database_dataguard_associations.exacc_test_autonomous_container_database_dataguard_associations.autonomous_container_database_dataguard_associations[0]["id"]}`},
		"autonomous_container_database_id":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_container_database.exacc_test_autonomous_container_database.id}`},
		"is_automatic_failover_enabled":                          acctest.Representation{RepType: acctest.Optional, Create: `true`, Update: `false`},
	}
)

// issue-routing-tag: database/dbaas-atp-d
func TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_update(t *testing.T) {

	httpreplay.SetScenario("TestDatabaseAutonomousContainerDatabaseDataguardAssociationResource_update")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_container_database_dataguard_associations.test_autonomous_container_database_dataguard_associations"
	singularDatasourceName := "data.oci_database_autonomous_container_database_dataguard_association.test_autonomous_container_database_dataguard_association"
	resourceName := "oci_database_autonomous_container_database_dataguard_association.test_update_autonomous_container_database_dataguard_association"
	acctest.SaveConfigContent("", "", "", t)

	AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation := acctest.RepresentationCopyWithRemovedProperties(
		acctest.GetUpdatedRepresentationCopy("months",
			[]acctest.RepresentationGroup{{RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation2}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation3}, {RepType: acctest.Optional, Group: DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsMonthsRepresentation4}},
			DatabaseAutonomousContainerDatabaseMaintenanceWindowDetailsRepresentation), []string{"lead_time_in_weeks"})

	AutonomousContainerDatabaseDedicatedRepresentation := acctest.GetUpdatedRepresentationCopy("maintenance_window_details", acctest.RepresentationGroup{RepType: acctest.Optional, Group: AutonomousContainerDatabaseDedicatedMaintenanceWindowDetailsRepresentation}, DatabaseAutonomousContainerDatabaseRepresentation)

	DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database", "test_autonomous_container_database", acctest.Optional, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(AutonomousContainerDatabaseDedicatedRepresentation, []string{"vault_id", "kms_key_id", "peer_autonomous_container_database_backup_config", "peer_autonomous_container_database_compartment_id", "peer_autonomous_vm_cluster_id"}), map[string]interface{}{
			"service_level_agreement_type":        acctest.Representation{RepType: acctest.Optional, Create: `AUTONOMOUS_DATAGUARD`},
			"protection_mode":                     acctest.Representation{RepType: acctest.Optional, Create: `MAXIMUM_AVAILABILITY`},
			"peer_cloud_autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_database_cloud_autonomous_vm_cluster.peer_cloud_autonomous_vm_cluster.id}`},
		})) +
		DatabaseCloudAutonomousVmClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_exadata_infrastructure", "peer_cloud_exadata_infrastructure", acctest.Required, acctest.Create, PeerCeiRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "test_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, ATPDCloudAutonomousVmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_cloud_autonomous_vm_cluster", "peer_cloud_autonomous_vm_cluster", acctest.Optional, acctest.Create, PeerCloudAvmRepresentation)

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//create datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "autonomous_container_database_dataguard_associations.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "autonomous_container_database_dataguard_associations.0.is_automatic_failover_enabled", "false"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, DatabaseAtpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_automatic_failover_enabled", "false"),
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

		// verify create with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_update_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Create, DatabaseAtpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},

		// verify updates with optional parameters
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_associations", "test_autonomous_container_database_dataguard_associations", acctest.Optional, acctest.Create, DatabaseautonomousContainerDatabaseDataguardAssociationDataSourceRepresentation) +
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_container_database_dataguard_association", "test_update_autonomous_container_database_dataguard_association", acctest.Optional, acctest.Update, DatabaseAtpdAutonomousContainerDatabaseDataguardAssociationResourceRepresentation) +
				compartmentIdVariableStr + DatabaseAutonomousContainerDatabaseDataguardAssociationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "autonomous_container_database_id"),

				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_automatic_failover_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_dataguard_association_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_autonomous_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_lifecycle_state"),
				resource.TestCheckResourceAttrSet(resourceName, "peer_role"),
				resource.TestCheckResourceAttrSet(resourceName, "protection_mode"),
				resource.TestCheckResourceAttrSet(resourceName, "role"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
			),
		},
	})
}
