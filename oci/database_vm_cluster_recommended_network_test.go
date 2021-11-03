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
	vmClusterRecommendedNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             Representation{RepType: Required, Create: `${var.compartment_id}`},
		"display_name":               Representation{RepType: Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":  Representation{RepType: Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"networks":                   []RepresentationGroup{{Required, vmClusterRecommendedNetworkClientNetworksRepresentation}, {Required, vmClusterRecommendedNetworkbackupNetworksRepresentation}},
		"defined_tags":               Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                        Representation{RepType: Optional, Create: []string{`192.168.10.10`}},
		"freeform_tags":              Representation{RepType: Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                        Representation{RepType: Optional, Create: []string{`192.168.10.20`}},
		"scan_listener_port_tcp":     Representation{RepType: Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl": Representation{RepType: Optional, Create: `2484`},
	}
	vmClusterRecommendedNetworkClientNetworksRepresentation = map[string]interface{}{
		"cidr":         Representation{RepType: Required, Create: `192.168.19.2/16`},
		"domain":       Representation{RepType: Required, Create: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.168.20.1`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`},
		"network_type": Representation{RepType: Required, Create: `CLIENT`},
		"prefix":       Representation{RepType: Required, Create: `myprefix1`},
		"vlan_id":      Representation{RepType: Required, Create: `10`},
	}
	vmClusterRecommendedNetworkbackupNetworksRepresentation = map[string]interface{}{
		"cidr":         Representation{RepType: Required, Create: `192.169.19.2/16`},
		"domain":       Representation{RepType: Required, Create: `oracle.com`},
		"gateway":      Representation{RepType: Required, Create: `192.169.20.1`},
		"netmask":      Representation{RepType: Required, Create: `255.255.0.0`},
		"network_type": Representation{RepType: Required, Create: `BACKUP`},
		"prefix":       Representation{RepType: Required, Create: `myprefix1`},
		"vlan_id":      Representation{RepType: Required, Create: `11`},
	}

	VmClusterRecommendedNetworkDataSourceDependencies = ExadataInfrastructureResourceActivateDependencies +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
			"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
		}))
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterRecommendedNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterRecommendedNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_vm_cluster_recommended_network.test_vm_cluster_recommended_network"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + VmClusterRecommendedNetworkDataSourceDependencies +
				GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_recommended_network", "test_vm_cluster_recommended_network", Optional, Create, vmClusterRecommendedNetworkSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.cidr", "192.168.19.2/16"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.domain", "oracle.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.gateway", "192.168.20.1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.netmask", "255.255.0.0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.network_type", "CLIENT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "networks.0.prefix", "myprefix1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "networks.0.vlan_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_port_tcp", "1521"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scan_listener_port_tcp_ssl", "2484"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scans.#", "1"),
			),
		},
	})
}
