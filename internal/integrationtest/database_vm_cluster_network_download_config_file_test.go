// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation = map[string]interface{}{
		"exadata_infrastructure_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_exadata_infrastructure.test_exadata_infrastructure.id}`},
		"vm_cluster_network_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_database_vm_cluster_network.test_vm_cluster_network.id}`},
		"base64_encode_content":     acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	VmClusterNetworkDownloadConfigFileResourceConfig = VmClusterNetworkValidatedResourceConfig
)

// issue-routing-tag: database/ExaCC
func TestDatabaseVmClusterNetworkDownloadConfigFileResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseVmClusterNetworkDownloadConfigFileResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_database_vm_cluster_network_download_config_file.test_vm_cluster_network_download_config_file"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_network_download_config_file", "test_vm_cluster_network_download_config_file", acctest.Required, acctest.Create, vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterNetworkDownloadConfigFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "false"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},

		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_database_vm_cluster_network_download_config_file", "test_vm_cluster_network_download_config_file", acctest.Optional, acctest.Create, vmClusterNetworkDownloadConfigFileSingularDataSourceRepresentation) +
				compartmentIdVariableStr + VmClusterNetworkDownloadConfigFileResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "exadata_infrastructure_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vm_cluster_network_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "base64_encode_content", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
			),
		},
	})
}
