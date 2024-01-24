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
	OcvpSupportedVmwareSoftwareVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"host_shape_name": acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO2.52`},
		"version":         acctest.Representation{RepType: acctest.Optional, Create: `7.0 update 3`},
	}

	OcvpSupportedVmwareSoftwareVersionResourceConfig = ""
)

// issue-routing-tag: ocvp/default
func TestOcvpSupportedVmwareSoftwareVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSupportedVmwareSoftwareVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ocvp_supported_vmware_software_versions.test_supported_vmware_software_versions"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_vmware_software_versions", "test_supported_vmware_software_versions", acctest.Required, acctest.Create, OcvpSupportedVmwareSoftwareVersionDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSupportedVmwareSoftwareVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.supported_host_shape_names.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.description"),
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_vmware_software_versions", "test_supported_vmware_software_versions", acctest.Optional, acctest.Create, OcvpSupportedVmwareSoftwareVersionDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSupportedVmwareSoftwareVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "host_shape_name"),
				resource.TestCheckResourceAttr(datasourceName, "version", `7.0 update 3`),

				resource.TestCheckResourceAttr(datasourceName, "items.0.description", "7.0 update 3"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.version"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.supported_host_shape_names.0"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.esxi_software_versions.0.description"),
			),
		},
	})
}
