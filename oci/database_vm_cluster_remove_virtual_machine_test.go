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
	vmClusterRemoveVirtualMachineRepresentation = map[string]interface{}{
		"db_servers":    RepresentationGroup{Required, vmClusterRemoveVirtualMachineDbServersRepresentation},
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	vmClusterRemoveVirtualMachineDbServersRepresentation = map[string]interface{}{
		"db_server_id": Representation{RepType: Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.2.id}`},
	}

	vmClusterRemoveDbServerRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              Representation{RepType: Required, Create: `6`, Update: `6`},
		"display_name":                Representation{RepType: Required, Create: `vmCluster`},
		"exadata_infrastructure_id":   Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  Representation{RepType: Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":             Representation{RepType: Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       Representation{RepType: Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"data_storage_size_in_tbs":    Representation{RepType: Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs": Representation{RepType: Optional, Create: `120`, Update: `160`},
		"db_servers":                  Representation{RepType: Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.2.id}`}},
		"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     Representation{RepType: Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": Representation{RepType: Optional, Create: `false`},
		"license_model":               Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs":          Representation{RepType: Optional, Create: `60`, Update: `90`},
		"time_zone":                   Representation{RepType: Optional, Create: `US/Pacific`},
	}

	VmClusterRemoveVirtualMachineResourceDependencies = GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation) +
		DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
			RepresentationCopyWithNewProperties(exadataInfrastructureActivateHalfRackRepresentation, map[string]interface{}{
				"activation_file":    Representation{RepType: Optional, Update: activationFilePath},
				"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
			})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateHalfRackRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterRemoveDbServerRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterRemoveVirtualMachineResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterRemoveVirtualMachineResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_remove_virtual_machine.test_vm_cluster_remove_virtual_machine"

	// Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VmClusterRemoveVirtualMachineResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_remove_virtual_machine", "test_vm_cluster_remove_virtual_machine", Required, Create, vmClusterRemoveVirtualMachineRepresentation), "database", "vmClusterRemoveVirtualMachine", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create

			{
				Config: config + compartmentIdVariableStr + VmClusterRemoveVirtualMachineResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_database_vm_cluster_remove_virtual_machine", "test_vm_cluster_remove_virtual_machine", Required, Create, vmClusterRemoveVirtualMachineRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "180"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "cpus_enabled", "6"),
					resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				),
			},
		},
	})
}
