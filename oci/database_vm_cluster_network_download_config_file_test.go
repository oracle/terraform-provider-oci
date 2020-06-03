// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": Representation{repType: Required, create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     Representation{repType: Required, create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"base64_encode_content":     Representation{repType: Optional, create: `true`},
	}

	VmClusterNetworkDownloadConfigFileResourceConfig = VmClusterNetworkValidatedResourceConfig
)

func TestDatabaseVmClusterNetworkDownloadConfigFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterNetworkDownloadConfigFileResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_vm_cluster_network_download_config_file.test_vm_cluster_network_download_config_file"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_network_download_config_file", "test_vm_cluster_network_download_config_file", Required, Create, vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterNetworkDownloadConfigFileResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},

			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_database_vm_cluster_network_download_config_file", "test_vm_cluster_network_download_config_file", Optional, Create, vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation) +
					compartmentIdVariableStr + VmClusterNetworkDownloadConfigFileResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				),
			},
		},
	})
}
