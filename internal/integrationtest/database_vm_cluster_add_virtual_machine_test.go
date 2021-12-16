// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vmClusterAddVirtualMachineRepresentation = map[string]interface{}{
		"db_servers":    acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterAddVirtualMachineDbServersRepresentation},
		"vm_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster.test_vm_cluster.id}`},
	}
	vmClusterAddVirtualMachineDbServersRepresentation = map[string]interface{}{
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.3.id}`},
	}

	exadataInfrastructureActivateHalfRackRepresentation = map[string]interface{}{
		"admin_network_cidr":          acctest.Representation{RepType: acctest.Required, Create: `192.168.0.0/16`, Update: `192.168.0.0/20`},
		"cloud_control_plane_server1": acctest.Representation{RepType: acctest.Required, Create: `10.32.88.1`, Update: `10.32.88.2`},
		"cloud_control_plane_server2": acctest.Representation{RepType: acctest.Required, Create: `10.32.88.3`, Update: `10.32.88.4`},
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `tstExaInfra`},
		"dns_server":                  acctest.Representation{RepType: acctest.Required, Create: []string{`10.231.225.65`}, Update: []string{`10.31.138.25`, `206.223.27.1`}},
		"gateway":                     acctest.Representation{RepType: acctest.Required, Create: `10.32.88.5`, Update: `10.32.88.6`},
		"infini_band_network_cidr":    acctest.Representation{RepType: acctest.Required, Create: `10.31.8.0/21`, Update: `10.31.8.0/22`},
		"netmask":                     acctest.Representation{RepType: acctest.Required, Create: `255.255.255.0`, Update: `255.255.254.0`},
		"ntp_server":                  acctest.Representation{RepType: acctest.Required, Create: []string{`10.231.225.76`}, Update: []string{`10.246.6.36`, `10.31.138.20`}},
		"shape":                       acctest.Representation{RepType: acctest.Required, Create: `ExadataCC.HalfX8M.200`},
		"time_zone":                   acctest.Representation{RepType: acctest.Required, Create: `US/Pacific`, Update: `UTC`},
		"contacts":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureContactsRepresentation},
		"corporate_proxy":             acctest.Representation{RepType: acctest.Optional, Create: `http://192.168.19.1:80`, Update: `http://192.168.19.2:80`},
		//"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
	}

	vmClusterNetworkValidateHalfRackRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkScansHalfRackRepresentation},
		"vm_networks":                 []acctest.RepresentationGroup{{RepType: acctest.Required, Group: vmClusterNetworkBackupVmNetworkHalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkClientVmNetworkHalfRackRepresentation}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Optional, Create: "true", Update: "true"},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: vmClusterNetworkIgnoreChangesHalfRackRepresentation},
	}
	vmClusterNetworkIgnoreChangesHalfRackRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`validate_vm_cluster_network`}},
	}

	vmClusterNetworkScansHalfRackRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix1-ivmmj-scan`, Update: `myprefix2-ivmmj-scan`},
		"ips":      acctest.Representation{RepType: acctest.Required, Create: []string{`192.168.19.7`, `192.168.19.6`, `192.168.19.8`}, Update: []string{`192.168.19.7`, `192.168.19.8`, `192.168.19.9`}},
		"port":     acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	vmClusterNetworkBackupVmNetworkHalfRackRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.169.20.1`, Update: `192.169.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: vmClusterNetworkVmNetworksBackupNodes1HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksBackupNodes2HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksBackupNodes3HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksBackupNodes4HalfRackRepresentation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `100`},
	}
	vmClusterNetworkClientVmNetworkHalfRackRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.168.20.1`, Update: `192.168.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: vmClusterNetworkVmNetworksClientNodes1HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksClientNodes2HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksClientNodes3HalfRackRepresentation}, {RepType: acctest.Required, Group: vmClusterNetworkVmNetworksClientNodes4HalfRackRepresentation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `101`},
	}
	vmClusterNetworkVmNetworksClientNodes1HalfRackRepresentation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb21`, Update: `myprefix2-xapb22`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.10`, Update: `192.168.19.11`},
		"vip":          acctest.Representation{RepType: acctest.Optional, Create: `192.168.19.12`, Update: `192.168.19.13`},
		"vip_hostname": acctest.Representation{RepType: acctest.Optional, Create: `myprefix2-xapb21-vip`, Update: `myprefix2-xapb22-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes2HalfRackRepresentation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb25`, Update: `myprefix2-xapb26`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.14`, Update: `192.168.19.15`},
		"vip":          acctest.Representation{RepType: acctest.Optional, Create: `192.168.19.16`, Update: `192.168.19.17`},
		"vip_hostname": acctest.Representation{RepType: acctest.Optional, Create: `myprefix2-xapb25-vip`, Update: `myprefix2-xapb26-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes3HalfRackRepresentation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb29`, Update: `myprefix2-xapb30`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.18`, Update: `192.168.19.19`},
		"vip":          acctest.Representation{RepType: acctest.Optional, Create: `192.168.19.20`, Update: `192.168.19.21`},
		"vip_hostname": acctest.Representation{RepType: acctest.Optional, Create: `myprefix2-xapb29-vip`, Update: `myprefix2-xapb30-vip`},
	}
	vmClusterNetworkVmNetworksClientNodes4HalfRackRepresentation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb33`, Update: `myprefix2-xapb34`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.22`, Update: `192.168.19.23`},
		"vip":          acctest.Representation{RepType: acctest.Optional, Create: `192.168.19.24`, Update: `192.168.19.25`},
		"vip_hostname": acctest.Representation{RepType: acctest.Optional, Create: `myprefix2-xapb33-vip`, Update: `myprefix2-xapb34-vip`},
	}
	vmClusterNetworkVmNetworksBackupNodes1HalfRackRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb23`, Update: `myprefix2-xapb24`},
		"ip":       acctest.Representation{RepType: acctest.Required, Create: `192.169.19.26`, Update: `192.169.19.27`},
	}
	vmClusterNetworkVmNetworksBackupNodes2HalfRackRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb27`, Update: `myprefix2-xapb28`},
		"ip":       acctest.Representation{RepType: acctest.Required, Create: `192.169.19.28`, Update: `192.169.19.29`},
	}
	vmClusterNetworkVmNetworksBackupNodes3HalfRackRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb31`, Update: `myprefix2-xapb32`},
		"ip":       acctest.Representation{RepType: acctest.Required, Create: `192.169.19.30`, Update: `192.169.19.31`},
	}
	vmClusterNetworkVmNetworksBackupNodes4HalfRackRepresentation = map[string]interface{}{
		"hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb35`, Update: `myprefix2-xapb36`},
		"ip":       acctest.Representation{RepType: acctest.Required, Create: `192.169.19.32`, Update: `192.169.19.33`},
	}

	vmClusterDbServerRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"cpu_core_count":              acctest.Representation{RepType: acctest.Required, Create: `6`, Update: `6`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `vmCluster`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"gi_version":                  acctest.Representation{RepType: acctest.Required, Create: `19.0.0.0.0`},
		"ssh_public_keys":             acctest.Representation{RepType: acctest.Required, Create: []string{`ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDOuBJgh6lTmQvQJ4BA3RCJdSmxRtmiXAQEEIP68/G4gF3XuZdKEYTFeputacmRq9yO5ZnNXgO9akdUgePpf8+CfFtveQxmN5xo3HVCDKxu/70lbMgeu7+wJzrMOlzj+a4zNq2j0Ww2VWMsisJ6eV3bJTnO/9VLGCOC8M9noaOlcKcLgIYy4aDM724MxFX2lgn7o6rVADHRxkvLEXPVqYT4syvYw+8OVSnNgE4MJLxaw8/2K0qp19YlQyiriIXfQpci3ThxwLjymYRPj+kjU1xIxv6qbFQzHR7ds0pSWp1U06cIoKPfCazU9hGWW8yIe/vzfTbWrt2DK6pLwBn/G0x3 sample`}},
		"vm_cluster_network_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"data_storage_size_in_tbs":    acctest.Representation{RepType: acctest.Optional, Create: `84`, Update: `86`},
		"db_node_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `120`, Update: `160`},
		"db_servers":                  acctest.Representation{RepType: acctest.Required, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"is_local_backup_enabled":     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"is_sparse_diskgroup_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"license_model":               acctest.Representation{RepType: acctest.Optional, Create: `LICENSE_INCLUDED`},
		"memory_size_in_gbs":          acctest.Representation{RepType: acctest.Optional, Create: `60`, Update: `90`},
		"time_zone":                   acctest.Representation{RepType: acctest.Optional, Create: `US/Pacific`},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: vmClusterDbServerIgnoreChangesRepresentation},
	}

	vmClusterDbServerIgnoreChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`db_servers`, `cpu_core_count`}},
	}

	VmClusterAddVirtualMachineResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateHalfRackRepresentation, map[string]interface{}{
				"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
				"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, vmClusterNetworkValidateHalfRackRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterDbServerRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterAddVirtualMachineResource_basic(t *testing.T) {
	t.Skip("Skip test for there is a diff in plan despite adding lifecycle state ignore changes for db_servers and cpu_core_count for the resource `oci_database_vm_cluster`")
	httpreplay.SetScenario("TestDatabaseVmClusterAddVirtualMachineResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_add_virtual_machine.test_vm_cluster_add_virtual_machine"

	//Save TF content to create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+VmClusterAddVirtualMachineResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_add_virtual_machine", "test_vm_cluster_add_virtual_machine", acctest.Required, acctest.Create, vmClusterAddVirtualMachineRepresentation), "database", "vmClusterAddVirtualMachine", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify create
			{

				Config: config + compartmentIdVariableStr + VmClusterAddVirtualMachineResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_add_virtual_machine", "test_vm_cluster_add_virtual_machine", acctest.Required, acctest.Create, vmClusterAddVirtualMachineRepresentation) +
					acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, dbServerDataSourceRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "db_node_storage_size_in_gbs", "180"),
					resource.TestCheckResourceAttr(resourceName, "data_storage_size_in_tbs", "2"),
					resource.TestCheckResourceAttr(resourceName, "cpus_enabled", "9"),
					resource.TestCheckResourceAttr(resourceName, "memory_size_in_gbs", "90"),
				),
			},
		},
	})
}
