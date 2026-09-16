// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	OcvpGenerateVmwareBinaryDownloadInfoDataSourceRepresentation = map[string]interface{}{
		"sddc_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.sddc_id}`},
		"vmware_binary_file_name": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_ocvp_retrieve_vmware_binaries.test_retrieve_vmware_binaries.items[0].file_name}`},
	}
)

// issue-routing-tag: ocvp/default
func TestGenerateVmwareBinaryDownloadInfoDataSource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerateVmwareBinaryDownloadInfoDataSource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	sddcId := utils.GetEnvSettingWithBlankDefault("ocvp_sddc_id")
	if sddcId == "" {
		t.Skip("Skipping test because env var 'ocvp_sddc_id' is not set")
	}
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	sddcIdVariableStr := fmt.Sprintf("variable \"sddc_id\" { default = \"%s\" }\n", sddcId)

	retrieveVmwareBinariesConfig := acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_retrieve_vmware_binaries", "test_retrieve_vmware_binaries", acctest.Required, acctest.Create, OcvpRetrieveVmwareBinariesDataSourceRepresentation)
	datasourceName := "data.oci_ocvp_generate_vmware_binary_download_info.test_generate_vmware_binary_download_info"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		{
			Config: config + compartmentIdVariableStr + sddcIdVariableStr + retrieveVmwareBinariesConfig +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_generate_vmware_binary_download_info", "test_generate_vmware_binary_download_info", acctest.Required, acctest.Create, OcvpGenerateVmwareBinaryDownloadInfoDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "sddc_id", sddcId),
				resource.TestCheckResourceAttrSet(datasourceName, "vmware_binary_file_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "file_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "time_expires"),
				resource.TestCheckResourceAttrSet(datasourceName, "url"),
			),
		},
	})
}
