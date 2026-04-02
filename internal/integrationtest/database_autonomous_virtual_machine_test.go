// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseAutonomousVirtualMachineSingularDataSourceRepresentation = map[string]interface{}{
		"autonomous_virtual_machine_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_autonomous_virtual_machines.test_autonomous_virtual_machines.autonomous_virtual_machines.0.id}`},
	}

	DatabaseDatabaseAutonomousVirtualMachineDataSourceRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"compartment_id":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"state":                    acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
	}

	DatabaseAutonomousVirtualMachineResourceConfig = AvailabilityDomainConfig + acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create,
		acctest.RepresentationCopyWithNewProperties(exadataInfrastructureRepresentationWithContacts, map[string]interface{}{"activation_file": acctest.Representation{RepType: acctest.Required, Create: activationFilePath}})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(vmClusterNetwork2Representation, map[string]interface{}{"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Required, Create: "true"}})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation) +
		DefinedTagsDependencies

	DatabaseAutonomousVirtualMachineRepresentation = map[string]interface{}{
		"autonomous_vm_cluster_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id}`},
		"client_ip_address":           acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              acctest.Representation{RepType: acctest.Required, Create: `4`, Update: `6`},
		"db_node_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `120`, Update: `160`},
		"db_server_id":                acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}}, "id": acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `90`},
		"state":              acctest.Representation{RepType: acctest.Optional, Create: `AVAILABLE`},
		"vm_name":            acctest.Representation{RepType: acctest.Optional, Create: `id`},
	}
)

// issue-routing-tag: database/default
func TestDatabaseAutonomousVirtualMachineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseAutonomousVirtualMachineResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_database_autonomous_virtual_machines.test_autonomous_virtual_machines"
	singularDatasourceName := "data.oci_database_autonomous_virtual_machine.test_autonomous_virtual_machine"
	simulateDb, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("simulate_db", "false"))

	dependencyConfig := DatabaseAutonomousVirtualMachineResourceConfig
	dataSourceRepresentation := DatabaseDatabaseAutonomousVirtualMachineDataSourceRepresentation
	singularDataSourceRepresentation := DatabaseDatabaseAutonomousVirtualMachineSingularDataSourceRepresentation

	if simulateDb {
		acctest.PreCheck(t)

		sharedAutonomousVmClusterID := utils.GetEnvSettingWithBlankDefault("autonomous_vm_cluster_id")
		sharedDependencyAddresses := []string{
			"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster",
		}
		sharedDependencyIDs, cleanup := ResolveOrCreateSharedDependenciesFromConfig(
			t,
			map[string]string{
				"oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster": sharedAutonomousVmClusterID,
			},
			config+compartmentIdVariableStr+DatabaseAVMClusterWithSingleNetworkResourceDependencies+
				acctest.GenerateResourceFromRepresentationMap("oci_database_autonomous_vm_cluster", "test_autonomous_vm_cluster", acctest.Required, acctest.Create, DatabaseAutonomousVmClusterRepresentation),
			sharedDependencyAddresses,
		)
		sharedAutonomousVmClusterID = sharedDependencyIDs["oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster"]
		if cleanup != nil {
			t.Cleanup(cleanup)
		}
		t.Logf("[SHARED_DEP_SETUP] autonomous_vm_cluster_id=%s", sharedAutonomousVmClusterID)

		dependencyConfig = fmt.Sprintf("variable \"autonomous_vm_cluster_id\" { default = \"%s\" }\n", sharedAutonomousVmClusterID)
		dataSourceRepresentation = acctest.RepresentationCopyWithNewProperties(DatabaseDatabaseAutonomousVirtualMachineDataSourceRepresentation, map[string]interface{}{
			"autonomous_vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: sharedAutonomousVmClusterID},
		})
	}

	acctest.SaveConfigContent("", "", "", t)

	t.Run("list-and-read-autonomous-virtual-machines", func(t *testing.T) {
		resource.Test(t, resource.TestCase{
			PreCheck: func() { acctest.PreCheck(t) },
			Providers: map[string]*schema.Provider{
				"oci": provider,
			},
			Steps: []resource.TestStep{
				// verify datasource
				{
					Config: config +
						acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_virtual_machines", "test_autonomous_virtual_machines", acctest.Required, acctest.Create, dataSourceRepresentation) +
						compartmentIdVariableStr + dependencyConfig + getExaccTagDependency(),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_vm_cluster_id"),
						resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.#"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.autonomous_vm_cluster_id"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.client_ip_address"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.compartment_id"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.cpu_core_count"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.db_node_storage_size_in_gbs"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.db_server_display_name"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.db_server_id"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.id"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.memory_size_in_gbs"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.state"),
						resource.TestCheckResourceAttrSet(datasourceName, "autonomous_virtual_machines.0.vm_name"),
						func(s *terraform.State) error {
							attrs := s.RootModule().Resources[datasourceName].Primary.Attributes
							t.Logf("[DATA_SOURCE_STATE] action=list autonomous virtual machines | autonomous_vm_cluster_id=%s | item_count=%s | first_id=%s | first_db_server_id=%s | first_cpu_core_count=%s | first_memory_size_in_gbs=%s | first_state=%s",
								attrs["autonomous_vm_cluster_id"],
								attrs["autonomous_virtual_machines.#"],
								attrs["autonomous_virtual_machines.0.id"],
								attrs["autonomous_virtual_machines.0.db_server_id"],
								attrs["autonomous_virtual_machines.0.cpu_core_count"],
								attrs["autonomous_virtual_machines.0.memory_size_in_gbs"],
								attrs["autonomous_virtual_machines.0.state"],
							)
							return nil
						},
					),
				},
				// verify singular datasource
				{
					Config: config +
						acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_virtual_machines", "test_autonomous_virtual_machines", acctest.Required, acctest.Create, dataSourceRepresentation) +
						acctest.GenerateDataSourceFromRepresentationMap("oci_database_autonomous_virtual_machine", "test_autonomous_virtual_machine", acctest.Required, acctest.Create, singularDataSourceRepresentation) +
						compartmentIdVariableStr + dependencyConfig + getExaccTagDependency(),
					Check: acctest.ComposeAggregateTestCheckFuncWrapper(
						resource.TestCheckResourceAttrSet(singularDatasourceName, "autonomous_virtual_machine_id"),

						resource.TestCheckResourceAttrSet(singularDatasourceName, "client_ip_address"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "cpu_core_count"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "db_node_storage_size_in_gbs"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "db_server_display_name"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "db_server_id"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "memory_size_in_gbs"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
						resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_name"),
						func(s *terraform.State) error {
							attrs := s.RootModule().Resources[singularDatasourceName].Primary.Attributes
							t.Logf("[DATA_SOURCE_STATE] action=read autonomous virtual machine | autonomous_virtual_machine_id=%s | db_server_id=%s | cpu_core_count=%s | memory_size_in_gbs=%s | state=%s | vm_name=%s",
								attrs["autonomous_virtual_machine_id"],
								attrs["db_server_id"],
								attrs["cpu_core_count"],
								attrs["memory_size_in_gbs"],
								attrs["state"],
								attrs["vm_name"],
							)
							return nil
						},
					),
				},
			},
		})
	})
}
