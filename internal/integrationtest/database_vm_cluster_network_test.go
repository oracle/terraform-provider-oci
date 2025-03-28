// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DatabaseVmClusterNetworkRequiredOnlyResource = DatabaseVmClusterNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseVmClusterNetworkRepresentation)

	DatabaseVmClusterNetworkResourceConfig = DatabaseVmClusterNetworkResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DatabaseVmClusterNetworkRepresentation)

	DatabaseDatabaseVmClusterNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
	}

	DatabaseDatabaseVmClusterNetworkDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"display_name":              acctest.Representation{RepType: acctest.Optional, Create: `testVmClusterNw`},
		"state":                     acctest.Representation{RepType: acctest.Optional, Create: `REQUIRES_VALIDATION`},
		"filter":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterNetworkDataSourceFilterRepresentation}}
	DatabaseVmClusterNetworkDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_database_vm_cluster_network.test_vm_cluster_network.id}`}},
	}

	DatabaseVmClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterNetworkScansRepresentation},
		"vm_networks":               []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetworkBackupVmNetworkRepresentation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkIgnoreNetworkRepresentation},
	}

	vmClusterNetwork2Representation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":              acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"scans":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterNetwork2ScansRepresentation},
		"vm_networks":               []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetwork2BackupVmNetworkRepresentation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetwork2ClientVmNetworkRepresentation}},
		"defined_tags":              acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":             acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                       acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"lifecycle":                 acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkIgnoreNetworkRepresentation},
	}

	DatabaseVmClusterNetworkScansRepresentation = map[string]interface{}{
		"hostname":                   acctest.Representation{RepType: acctest.Required, Create: `myprefix1-ivmmj-scan`, Update: `myprefix2-ivmmj-scan`},
		"ips":                        acctest.Representation{RepType: acctest.Required, Create: []string{`192.168.19.7`, `192.168.19.6`, `192.168.19.8`}, Update: []string{`192.168.19.7`, `192.168.19.8`, `192.168.19.9`}},
		"port":                       acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp":     acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp_ssl": acctest.Representation{RepType: acctest.Required, Create: `2484`, Update: `2484`},
	}

	DatabaseVmClusterNetwork2ScansRepresentation = map[string]interface{}{
		"hostname":                   acctest.Representation{RepType: acctest.Required, Create: `myprefix4-ivmmj-scan`, Update: `myprefix3-ivmmj-scan`},
		"ips":                        acctest.Representation{RepType: acctest.Required, Create: []string{`192.168.19.26`, `192.168.19.27`, `192.168.19.28`}, Update: []string{`192.168.19.27`, `192.168.19.28`, `192.168.19.29`}},
		"port":                       acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp":     acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"scan_listener_port_tcp_ssl": acctest.Representation{RepType: acctest.Required, Create: `2484`, Update: `2484`},
	}

	DatabaseVmClusterNetworkBackupVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.169.20.1`, Update: `192.169.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetworkVmNetworksBackupNodes1Representation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetworkVmNetworksBackupNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `100`},
	}

	DatabaseVmClusterNetwork2BackupVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.169.20.2`, Update: `192.169.20.3`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetwork2VmNetworksBackupNodes1Representation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetwork2VmNetworksBackupNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `100`},
	}

	DatabaseVmClusterNetworkClientVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.168.20.1`, Update: `192.168.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetworkVmNetworksClientNodes1Representation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetworkVmNetworksClientNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `101`},
	}

	DatabaseVmClusterNetwork2ClientVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.168.20.2`, Update: `192.168.20.3`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterNetwork2VmNetworksClientNodes1Representation}, {RepType: acctest.Required, Group: DatabaseVmClusterNetwork2VmNetworksClientNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `101`},
	}

	DatabaseVmClusterNetworkVmNetworksClientNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb21`, Update: `myprefix2-xapb22`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.10`, Update: `192.168.19.11`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.12`, Update: `192.168.19.13`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb21-vip`, Update: `myprefix2-xapb22-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
	}

	DatabaseVmClusterNetwork2VmNetworksClientNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb31`, Update: `myprefix2-xapb32`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.18`, Update: `192.168.19.19`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.20`, Update: `192.168.19.21`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb31-vip`, Update: `myprefix2-xapb32-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
	}

	DatabaseVmClusterNetworkVmNetworksClientNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb25`, Update: `myprefix2-xapb26`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.14`, Update: `192.168.19.15`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.16`, Update: `192.168.19.17`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb25-vip`, Update: `myprefix2-xapb26-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`},
	}

	DatabaseVmClusterNetwork2VmNetworksClientNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb35`, Update: `myprefix2-xapb36`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.22`, Update: `192.168.19.23`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.24`, Update: `192.168.19.25`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb35-vip`, Update: `myprefix2-xapb36-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`},
	}

	DatabaseVmClusterNetworkVmNetworksBackupNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb23`, Update: `myprefix2-xapb24`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.18`, Update: `192.169.19.19`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
	}
	DatabaseVmClusterNetworkVmNetworksBackupNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb27`, Update: `myprefix2-xapb28`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.20`, Update: `192.169.19.21`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`},
	}

	DatabaseVmClusterNetwork2VmNetworksBackupNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb33`, Update: `myprefix2-xapb34`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.22`, Update: `192.169.19.23`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`},
	}

	DatabaseVmClusterNetwork2VmNetworksBackupNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb37`, Update: `myprefix2-xapb38`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.24`, Update: `192.169.19.25`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`},
	}

	vmClusterNetworkIgnoreNetworkRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`vm_networks`}},
	}

	// DG peer network related representations below
	DatabasePeerVmClusterNetworkVmNetworksClientNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb21`, Update: `myprefix2-xapb22`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.10`, Update: `192.168.19.11`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.12`, Update: `192.168.19.13`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb21-vip`, Update: `myprefix2-xapb22-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.peer_db_servers.db_servers.0.id}`},
	}

	DatabasePeerVmClusterNetworkVmNetworksClientNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb25`, Update: `myprefix2-xapb26`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.168.19.14`, Update: `192.168.19.15`},
		"vip":          acctest.Representation{RepType: acctest.Required, Create: `192.168.19.16`, Update: `192.168.19.17`},
		"vip_hostname": acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb25-vip`, Update: `myprefix2-xapb26-vip`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.peer_db_servers.db_servers.1.id}`},
	}

	DatabasePeerVmClusterNetworkVmNetworksBackupNodes1Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb23`, Update: `myprefix2-xapb24`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.18`, Update: `192.169.19.19`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.peer_db_servers.db_servers.0.id}`},
	}
	DatabasePeerVmClusterNetworkVmNetworksBackupNodes2Representation = map[string]interface{}{
		"hostname":     acctest.Representation{RepType: acctest.Required, Create: `myprefix2-xapb27`, Update: `myprefix2-xapb28`},
		"ip":           acctest.Representation{RepType: acctest.Required, Create: `192.169.19.20`, Update: `192.169.19.21`},
		"db_server_id": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_database_db_servers.peer_db_servers.db_servers.1.id}`},
	}

	DatabasePeerVmClusterNetworkBackupVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.169.20.1`, Update: `192.169.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`, Update: `BACKUP`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkVmNetworksBackupNodes1Representation}, {RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkVmNetworksBackupNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `100`},
	}

	DatabasePeerVmClusterNetworkClientVmNetworkRepresentation = map[string]interface{}{
		"domain_name":  acctest.Representation{RepType: acctest.Required, Create: `oracle.com`, Update: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.168.20.1`, Update: `192.168.20.2`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`, Update: `255.255.192.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `CLIENT`, Update: `CLIENT`},
		"nodes":        []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkVmNetworksClientNodes1Representation}, {RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkVmNetworksClientNodes2Representation}},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `101`},
	}

	DatabasePeerVmClusterNetworkRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Required, Create: `peerVmClusterNw`},
		"exadata_infrastructure_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.peer_exadata_infrastructure.id}`},
		"scans":                       acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseVmClusterNetworkScansRepresentation},
		"vm_networks":                 []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkBackupVmNetworkRepresentation}, {RepType: acctest.Required, Group: DatabasePeerVmClusterNetworkClientVmNetworkRepresentation}},
		"defined_tags":                acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}, Update: []string{`192.168.10.12`}},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                         acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}, Update: []string{`192.168.10.22`}},
		"lifecycle":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: vmClusterNetworkIgnoreNetworkRepresentation},
		"validate_vm_cluster_network": acctest.Representation{RepType: acctest.Required, Create: "true"},
	}
	// End DG peer network related representations

	activationFilePath, _                        = createTmpActivationFile()
	DatabaseVmClusterNetworkResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update,
			acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
				"activation_file":    acctest.Representation{RepType: acctest.Optional, Update: activationFilePath},
				"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
			}))
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_database_vm_cluster_network.test_vm_cluster_network"
	datasourceName := "data.oci_database_vm_cluster_networks.test_vm_cluster_networks"
	singularDatasourceName := "data.oci_database_vm_cluster_network.test_vm_cluster_network"

	var resId, resId2, compositeId string

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+DatabaseVmClusterNetworkResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, DatabaseVmClusterNetworkRepresentation), "database", "vmClusterNetwork", t)

	acctest.ResourceTest(t, testAccCheckDatabaseVmClusterNetworkDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetwork2Representation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname": "myprefix4-ivmmj-scan",
					"ips.#":    "3",
					"port":     "1521",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"network_type": "CLIENT",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterNetworkResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Create, DatabaseVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname":                   "myprefix1-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1521",
					"scan_listener_port_tcp":     "1521",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"network_type": "CLIENT",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					exadataInfrastructureId, _ := acctest.FromInstanceState(s, resourceName, "exadata_infrastructure_id")
					compositeId = "exadataInfrastructures/" + exadataInfrastructureId + "/vmClusterNetworks/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + DatabaseVmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DatabaseVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(resourceName, "dns.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "vm_networks", map[string]string{
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_networks", "test_vm_cluster_networks", acctest.Optional, acctest.Update, DatabaseDatabaseVmClusterNetworkDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterNetworkResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Optional, acctest.Update, DatabaseVmClusterNetworkRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttrSet(datasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "REQUIRES_VALIDATION"),

				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.dns.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.ntp.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "vm_cluster_networks.0.scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "vm_cluster_networks.0.time_created"),
				resource.TestCheckResourceAttr(datasourceName, "vm_cluster_networks.0.vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "vm_cluster_networks.0.vm_networks", map[string]string{
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{
						"vlan_id",
					}),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, DatabaseDatabaseVmClusterNetworkSingularDataSourceRepresentation) +
				compartmentIdVariableStr + DatabaseVmClusterNetworkResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scans.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "scans", map[string]string{
					"hostname":                   "myprefix2-ivmmj-scan",
					"ips.#":                      "3",
					"port":                       "1522",
					"scan_listener_port_tcp":     "1522",
					"scan_listener_port_tcp_ssl": "2484",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttr(singularDatasourceName, "vm_networks.#", "2"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "vm_networks", map[string]string{
					"network_type": "BACKUP",
					"nodes.#":      "2",
				},
					[]string{}),
			),
		},
		// verify resource import
		{
			Config:            config + DatabaseVmClusterNetworkRequiredOnlyResource,
			ImportState:       true,
			ImportStateIdFunc: getVmClusterNetworkImportId(resourceName),
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"validate_vm_cluster_network",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatabaseVmClusterNetworkDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*client.OracleClients).DatabaseClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_database_vm_cluster_network" {
			noResourceFound = false
			request := oci_database.GetVmClusterNetworkRequest{}

			if value, ok := rs.Primary.Attributes["exadata_infrastructure_id"]; ok {
				request.ExadataInfrastructureId = &value
			}

			tmp := rs.Primary.ID
			request.VmClusterNetworkId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")

			response, err := client.GetVmClusterNetwork(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_database.VmClusterNetworkLifecycleStateTerminated): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatabaseVmClusterNetwork") {
		resource.AddTestSweepers("DatabaseVmClusterNetwork", &resource.Sweeper{
			Name:         "DatabaseVmClusterNetwork",
			Dependencies: acctest.DependencyGraph["vmClusterNetwork"],
			F:            sweepDatabaseVmClusterNetworkResource,
		})
	}
}

func sweepDatabaseVmClusterNetworkResource(compartment string) error {
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()
	vmClusterNetworkIds, err := getDatabaseVmClusterNetworkIds(compartment)
	if err != nil {
		return err
	}
	for _, vmClusterNetworkId := range vmClusterNetworkIds {
		if ok := acctest.SweeperDefaultResourceId[vmClusterNetworkId]; !ok {
			deleteVmClusterNetworkRequest := oci_database.DeleteVmClusterNetworkRequest{}

			deleteVmClusterNetworkRequest.VmClusterNetworkId = &vmClusterNetworkId

			deleteVmClusterNetworkRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "database")
			_, error := databaseClient.DeleteVmClusterNetwork(context.Background(), deleteVmClusterNetworkRequest)
			if error != nil {
				fmt.Printf("Error deleting VmClusterNetwork %s %s, It is possible that the resource is already deleted. Please verify manually \n", vmClusterNetworkId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &vmClusterNetworkId, DatabaseVmClusterNetworkSweepWaitCondition, time.Duration(3*time.Minute),
				DatabaseVmClusterNetworkSweepResponseFetchOperation, "database", true)
		}
	}
	return nil
}

func getDatabaseVmClusterNetworkIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "VmClusterNetworkId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	databaseClient := acctest.GetTestClients(&schema.ResourceData{}).DatabaseClient()

	listVmClusterNetworksRequest := oci_database.ListVmClusterNetworksRequest{}
	listVmClusterNetworksRequest.CompartmentId = &compartmentId

	exadataInfrastructureIds, error := getDatabaseExadataInfrastructureIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting exadataInfrastructureId required for VmClusterNetwork resource requests \n")
	}
	for _, exadataInfrastructureId := range exadataInfrastructureIds {
		listVmClusterNetworksRequest.ExadataInfrastructureId = &exadataInfrastructureId
		listVmClusterNetworksRequest.LifecycleState = oci_database.VmClusterNetworkSummaryLifecycleStateRequiresValidation
		listVmClusterNetworksResponse, err := databaseClient.ListVmClusterNetworks(context.Background(), listVmClusterNetworksRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting VmClusterNetwork list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, vmClusterNetwork := range listVmClusterNetworksResponse.Items {
			id := *vmClusterNetwork.Id
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "VmClusterNetworkId", id)
		}

	}
	return resourceIds, nil
}

func DatabaseVmClusterNetworkSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if vmClusterNetworkResponse, ok := response.Response.(oci_database.GetVmClusterNetworkResponse); ok {
		return vmClusterNetworkResponse.LifecycleState != oci_database.VmClusterNetworkLifecycleStateTerminated
	}
	return false
}

func DatabaseVmClusterNetworkSweepResponseFetchOperation(client *client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DatabaseClient().GetVmClusterNetwork(context.Background(), oci_database.GetVmClusterNetworkRequest{
		VmClusterNetworkId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getVmClusterNetworkImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("exadataInfrastructures/" + rs.Primary.Attributes["exadata_infrastructure_id"] + "/vmClusterNetworks/" + rs.Primary.Attributes["id"]), nil
	}
}
