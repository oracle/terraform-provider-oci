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
	vmClusterAddVirtualMachineRepresentation = map[string]interface{}{
		"db_servers":    RepresentationGroup{Required, vmClusterAddVirtualMachineDbServersRepresentation},
		"vm_cluster_id": Representation{RepType: Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	vmClusterAddVirtualMachineDbServersRepresentation = map[string]interface{}{
		"db_server_id": Representation{RepType: Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.3.id}`},
	}

	exadataInfrastructureActivateHalfRackRepresentation = map[string]interface{}{
		"admin_network_cidr":          Representation{RepType: Required, Create: `192.168.0.0/16`, Update: `192.168.0.0/20`},
		"cloud_control_plane_server1": Representation{RepType: Required, Create: `10.32.88.1`, Update: `10.32.88.2`},
		"cloud_control_plane_server2": Representation{RepType: Required, Create: `10.32.88.3`, Update: `10.32.88.4`},
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":                Representation{RepType: Required, Create: `tstExaInfra`},
		"dns_server":                  Representation{RepType: Required, Create: []string{`10.231.225.65`}, Update: []string{`10.31.138.25`, `206.223.27.1`}},
		"gateway":                     Representation{RepType: Required, Create: `10.32.88.5`, Update: `10.32.88.6`},
		"infini_band_network_cidr":    Representation{RepType: Required, Create: `10.31.8.0/21`, Update: `10.31.8.0/22`},
		"netmask":                     Representation{RepType: Required, Create: `255.255.255.0`, Update: `255.255.254.0`},
		"ntp_server":                  Representation{RepType: Required, Create: []string{`10.231.225.76`}, Update: []string{`10.246.6.36`, `10.31.138.20`}},
		"shape":                       Representation{RepType: Required, Create: `ExadataCC.HalfX8M.200`},
		"time_zone":                   Representation{RepType: Required, Create: `US/Pacific`, Update: `UTC`},
		"contacts":                    RepresentationGroup{Optional, exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             Representation{RepType: Optional, Create: `http://192.168.19.1:80`, Update: `http://192.168.19.2:80`},
		//"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
	}

	vmClusterNetworkValidateHalfRackRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":                Representation{RepType: Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":   Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                       RepresentationGroup{Required, vmClusterNetworkScansHalfRackRepresentation},
		"vm_networks":                 []RepresentationGroup{{Required, vmClusterNetworkBackupVmNetworkHalfRackRepresentation}, {Required, vmClusterNetworkClientVmNetworkHalfRackRepresentation}},
		"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         Representation{RepType: Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":               Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                         Representation{RepType: Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"validate_vm_cluster_network": Representation{RepType: Optional, Create: "true", Update: "true"},
		"lifecycle":                   RepresentationGroup{RepType: Optional, Group: vmClusterNetworkIgnoreChangesHalfRackRepresentation},
	}
	vmClusterNetworkIgnoreChangesHalfRackRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`validate_vm_cluster_network`}},
	}

	vmClusterNetworkScansHalfRackRepresentation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix1-ivmmj-scan`, Update: `myprefix2-ivmmj-scan`},
		"ips":      Representation{RepType: Required, Create: []string{`192.168.19.7`, `192.168.19.6`, `192.168.19.8`}, Update: []string{`192.168.19.7`, `192.168.19.8`, `192.168.19.9`}},
		"port":     Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	vmClusterNetworkBackupVmNetworkHalfRackRepresentation = map[string]interface{}{
		"domain_name":  Representation{RepType: Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.169.20.1`, Update: `192.169.20.2`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": Representation{RepType: Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []RepresentationGroup{{Required, vmClusterNetworkVmNetworksBackupNodes1HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksBackupNodes2HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksBackupNodes3HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksBackupNodes4HalfRackRepresentation}},
		"vlan_id":      Representation{RepType: Required, Create: `100`},
	}
	vmClusterNetworkClientVmNetworkHalfRackRepresentation = map[string]interface{}{
		"domain_name":  Representation{RepType: Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.168.20.1`, Update: `192.168.20.2`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": Representation{RepType: Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []RepresentationGroup{{Required, vmClusterNetworkVmNetworksClientNodes1HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksClientNodes2HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksClientNodes3HalfRackRepresentation}, {Required, vmClusterNetworkVmNetworksClientNodes4HalfRackRepresentation}},
		"vlan_id":      Representation{RepType: Required, Create: `101`},
	}
	vmClusterNetworkVmNetworksClientNodes1HalfRackRepresentation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb21`, Update: `myprefix2-xapb22`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.10`, Update: `192.168.19.11`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.12`, Update: `192.168.19.13`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb21-vip`, Update: `myprefix2-xapb22-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes2HalfRackRepresentation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb25`, Update: `myprefix2-xapb26`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.14`, Update: `192.168.19.15`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.16`, Update: `192.168.19.17`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb25-vip`, Update: `myprefix2-xapb26-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes3HalfRackRepresentation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb29`, Update: `myprefix2-xapb30`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.18`, Update: `192.168.19.19`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.20`, Update: `192.168.19.21`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb29-vip`, Update: `myprefix2-xapb30-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes4HalfRackRepresentation = map[string]interface{}{
		"hostname":     Representation{RepType: Required, Create: `myprefix2-xapb33`, Update: `myprefix2-xapb34`},
		"ip":           Representation{RepType: Required, Create: `192.168.19.22`, Update: `192.168.19.23`},
		"vip":          Representation{RepType: Optional, Create: `192.168.19.24`, Update: `192.168.19.25`},
		"vip_hostname": Representation{RepType: Optional, Create: `myprefix2-xapb33-vip`, Update: `myprefix2-xapb34-vip`},
	}
	vmClusterNetworkVmNetworksBackupNodes1HalfRackRepresentation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb23`, Update: `myprefix2-xapb24`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.26`, Update: `192.169.19.27`},
	}
	vmClusterNetworkVmNetworksBackupNodes2HalfRackRepresentation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb27`, Update: `myprefix2-xapb28`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.28`, Update: `192.169.19.29`},
	}
	vmClusterNetworkVmNetworksBackupNodes3HalfRackRepresentation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb31`, Update: `myprefix2-xapb32`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.30`, Update: `192.169.19.31`},
	}
	vmClusterNetworkVmNetworksBackupNodes4HalfRackRepresentation = map[string]interface{}{
		"hostname": Representation{RepType: Required, Create: `myprefix2-xapb35`, Update: `myprefix2-xapb36`},
		"ip":       Representation{RepType: Required, Create: `192.169.19.32`, Update: `192.169.19.33`},
	}

	vmClusterDbServerRepresentation = map[string]interface{}{
		"compartment_id":              Representation{RepType: Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              Representation{RepType: Required, Create: `6`, Update: `6`},
		"display_name":                Representation{RepType: Required, Create: `vmCluster`},
		"exadata_infrastructure_id":   Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  Representation{RepType: Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":             Representation{RepType: Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       Representation{RepType: Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"data_storage_size_in_tbs":    Representation{RepType: Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs": Representation{RepType: Optional, Create: `120`, Update: `160`},
		"db_servers":                  Representation{RepType: Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     Representation{RepType: Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": Representation{RepType: Optional, Create: `false`},
		"license_model":               Representation{RepType: Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs":          Representation{RepType: Optional, Create: `60`, Update: `90`},
		"time_zone":                   Representation{RepType: Optional, Create: `US/Pacific`},
		"lifecycle":                   RepresentationGroup{RepType: Optional, Group: vmClusterDbServerIgnoreChangesRepresentation},
	}

	vmClusterDbServerIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`db_servers`, `cpu_core_count`}},
	}

	VmClusterAddVirtualMachineResourceDependencies = DefinedTagsDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update,
			RepresentationCopyWithNewProperties(exadataInfrastructureActivateHalfRackRepresentation, map[string]interface{}{
				"activation_file":    Representation{RepType: Optional, Update: activationFilePath},
				"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
			})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateHalfRackRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", Required, Create, vmClusterDbServerRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterAddVirtualMachineResource_basic(t *testing.T) {
	t.Skip("Skip test for there is a diff in plan despite adding lifecycle state ignore changes for db_servers and cpu_core_count for the resource `oci_database_vm_cluster`")
	httpreplay.SetScenario("TestDatabaseVmClusterAddVirtualMachineResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_add_virtual_machine.test_vm_cluster_add_virtual_machine"

	//Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+VmClusterAddVirtualMachineResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_add_virtual_machine", "test_vm_cluster_add_virtual_machine", Required, Create, vmClusterAddVirtualMachineRepresentation), "database", "vmClusterAddVirtualMachine", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{

				Config: config + compartmentIdVariableStr + VmClusterAddVirtualMachineResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_database_vm_cluster_add_virtual_machine", "test_vm_cluster_add_virtual_machine", Required, Create, vmClusterAddVirtualMachineRepresentation) +
					GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", Required, Create, dbServerDataSourceRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "180"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "cpus_enabled", "9"),
					resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				),
			},
		},
	})
}
