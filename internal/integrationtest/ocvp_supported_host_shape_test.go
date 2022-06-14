// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"terraform-provider-oci/httpreplay"
	"terraform-provider-oci/internal/acctest"

	"terraform-provider-oci/internal/utils"
)

var (
	supportedHostShapeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: `BM.DenseIO.E4.128`},
		"sddc_type":      acctest.Representation{RepType: acctest.Optional, Create: `PRODUCTION`},
	}

	SupportedHostShapeResourceConfig = ""
)

// issue-routing-tag: ocvp/default
func TestOcvpSupportedHostShapeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSupportedHostShapeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ocvp_supported_host_shapes.test_supported_host_shapes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource with required params
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_host_shapes", "test_supported_host_shapes", acctest.Required, acctest.Create, supportedHostShapeDataSourceRepresentation) +
				compartmentIdVariableStr + SupportedHostShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttr(datasourceName, "items.#", "2"),
			),
		},
		// verify datasource with optional params
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_host_shapes", "test_supported_host_shapes", acctest.Optional, acctest.Create, supportedHostShapeDataSourceRepresentation) +
				compartmentIdVariableStr + SupportedHostShapeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
				resource.TestCheckResourceAttr(datasourceName, "items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.default_ocpu_count", "64"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.description", "2.55 GHz AMD EPYC™ 7J13 (Milan) processor."),
				resource.TestCheckResourceAttr(datasourceName, "items.0.is_support_shielded_instances", "false"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.name", "BM.DenseIO.E4.128"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.supported_operations.#", "2"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.shape_family", "AMD"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.supported_ocpu_count.#"),
				resource.TestCheckResourceAttr(datasourceName, "items.0.supported_sddc_types.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "items.0.supported_vmware_software_versions.0"),
			),
		},
	})
}
