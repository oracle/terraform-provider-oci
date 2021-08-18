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
	supportedVmwareSoftwareVersionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
	}

	SupportedVmwareSoftwareVersionResourceConfig = ""
)

// issue-routing-tag: ocvp/default
func TestOcvpSupportedVmwareSoftwareVersionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOcvpSupportedVmwareSoftwareVersionResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ocvp_supported_vmware_software_versions.test_supported_vmware_software_versions"

	saveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_ocvp_supported_vmware_software_versions", "test_supported_vmware_software_versions", Required, Create, supportedVmwareSoftwareVersionDataSourceRepresentation) +
				compartmentIdVariableStr + SupportedVmwareSoftwareVersionResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "items.#"),
			),
		},
	})
}
