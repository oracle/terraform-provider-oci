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
	vmClusterRecommendedNetworkSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":              Representation{repType: Required, create: `testVmClusterNw`},
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"networks":                  []RepresentationGroup{{Required, vmClusterRecommendedNetworkClientNetworksRepresentation}, {Required, vmClusterRecommendedNetworkbackupNetworksRepresentation}},
		"defined_tags":              Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"dns":                       Representation{repType: Optional, create: []string{`192.168.10.10`}},
		"freeform_tags":             Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"ntp":                       Representation{repType: Optional, create: []string{`192.168.10.20`}},
	}
	vmClusterRecommendedNetworkClientNetworksRepresentation = map[string]interface{}{
		"cidr":         Representation{repType: Required, create: `192.168.19.2/16`},
		"domain":       Representation{repType: Required, create: `oracle.com`},
		"gateway":      Representation{repType: Required, create: `192.168.20.1`},
		"netmask":      Representation{repType: Required, create: `255.255.0.0`},
		"network_type": Representation{repType: Required, create: `CLIENT`},
		"prefix":       Representation{repType: Required, create: `myprefix1`},
		"vlan_id":      Representation{repType: Required, create: `10`},
	}
	vmClusterRecommendedNetworkbackupNetworksRepresentation = map[string]interface{}{
		"cidr":         Representation{repType: Required, create: `192.169.19.2/16`},
		"domain":       Representation{repType: Required, create: `oracle.com`},
		"gateway":      Representation{repType: Required, create: `192.169.20.1`},
		"netmask":      Representation{repType: Required, create: `255.255.0.0`},
		"network_type": Representation{repType: Required, create: `BACKUP`},
		"prefix":       Representation{repType: Required, create: `myprefix1`},
		"vlan_id":      Representation{repType: Required, create: `11`},
	}

	VmClusterRecommendedNetworkDataSourceDependencies = ExadataInfrastructureResourceActivateDependencies +
		generateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, representationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{
			"maintenance_window": RepresentationGroup{Optional, exadataInfrastructureMaintenanceWindowRepresentationComplete},
		}))
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterRecommendedNetworkResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterRecommendedNetworkResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_vm_cluster_recommended_network.test_vm_cluster_recommended_network"

	saveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config + VmClusterRecommendedNetworkDataSourceDependencies +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_recommended_network", "test_vm_cluster_recommended_network", Required, Create, vmClusterRecommendedNetworkSingularDataSourceRepresentation) +
					compartmentIdVariableStr,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "testVmClusterNw"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "networks.#", "2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "scans.#", "1"),
				),
			},
		},
	})
}
