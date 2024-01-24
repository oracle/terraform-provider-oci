// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseDatabaseVmClusterRecommendedNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Required, Create: `testVmClusterNw`},
		"exadata_infrastructure_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"networks":                   []acctest.RepresentationGroup{{RepType: acctest.Required, Group: DatabaseVmClusterRecommendedNetworkNetworksSingularDataSourceRepresentation}, {RepType: acctest.Required, Group: DatabaseVmClusterRecommendedNetworkbackupNetworksRepresentation}},
		"db_servers":                 acctest.Representation{RepType: acctest.Optional, Create: []string{`${data.oci_database_db_servers.test_db_servers.db_servers.0.id}`, `${data.oci_database_db_servers.test_db_servers.db_servers.1.id}`}},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.10`}},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ntp":                        acctest.Representation{RepType: acctest.Optional, Create: []string{`192.168.10.20`}},
		"scan_listener_port_tcp":     acctest.Representation{RepType: acctest.Optional, Create: `1521`},
		"scan_listener_port_tcp_ssl": acctest.Representation{RepType: acctest.Optional, Create: `2484`},
	}
	DatabaseVmClusterRecommendedNetworkNetworksSingularDataSourceRepresentation = map[string]interface{}{
		"cidr":         acctest.Representation{RepType: acctest.Required, Create: `192.168.19.2/16`},
		"domain":       acctest.Representation{RepType: acctest.Required, Create: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.168.20.1`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `CLIENT`},
		"prefix":       acctest.Representation{RepType: acctest.Required, Create: `myprefix1`},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `10`},
	}
	DatabaseVmClusterRecommendedNetworkbackupNetworksRepresentation = map[string]interface{}{
		"cidr":         acctest.Representation{RepType: acctest.Required, Create: `192.169.19.2/16`},
		"domain":       acctest.Representation{RepType: acctest.Required, Create: `oracle.com`},
		"gateway":      acctest.Representation{RepType: acctest.Required, Create: `192.169.20.1`},
		"netmask":      acctest.Representation{RepType: acctest.Required, Create: `255.255.0.0`},
		"network_type": acctest.Representation{RepType: acctest.Required, Create: `BACKUP`},
		"prefix":       acctest.Representation{RepType: acctest.Required, Create: `myprefix1`},
		"vlan_id":      acctest.Representation{RepType: acctest.Required, Create: `11`},
	}

	DatabaseVmClusterRecommendedNetworkDataSourceDependencies = ExadataInfrastructureResourceActivateDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Optional, acctest.Update, acctest.RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
			"maintenance_window": acctest.RepresentationGroup{RepType: acctest.Optional, Group: exadataInfrastructureMaintenanceWindowRepresentationComplete},
		})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_database_db_servers", "test_db_servers", acctest.Required, acctest.Create, DatabaseDatabaseDbServerDataSourceRepresentation)
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterRecommendedNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterRecommendedNetworkResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_vm_cluster_recommended_network.test_vm_cluster_recommended_network"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + DatabaseVmClusterRecommendedNetworkDataSourceDependencies +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_recommended_network", "test_vm_cluster_recommended_network", acctest.Optional, acctest.Create, DatabaseDatabaseVmClusterRecommendedNetworkSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "db_servers.#", "2"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dns.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ntp.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "scans.#", "1"),
			),
		},
	})
}
