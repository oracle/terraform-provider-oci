// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
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
	OcvpOcvpSupportedVmwareSoftwareVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ocvp_supported_vmware_software_versions", "test_supported_vmware_software_versions", acctest.Required, acctest.Create, OcvpOcvpSupportedVmwareSoftwareVersionDataSourceRepresentation) +
				compartmentIdVariableStr + OcvpSupportedVmwareSoftwareVersionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
			),
		},
	})
}
