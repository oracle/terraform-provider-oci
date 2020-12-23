// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
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
	instanceAvailablePluginDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"os_name":        Representation{repType: Required, create: `Oracle Linux`},
		"os_version":     Representation{repType: Required, create: `7.8`},
	}

	InstanceAvailablePluginResourceConfig = ""
)

func TestComputeinstanceagentInstanceAvailablePluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeinstanceagentInstanceAvailablePluginResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_computeinstanceagent_instance_available_plugins.test_instance_available_plugins"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_available_plugins", "test_instance_available_plugins", Required, Create, instanceAvailablePluginDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceAvailablePluginResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "os_name", "Oracle Linux"),
					resource.TestCheckResourceAttr(datasourceName, "os_version", "7.8"),

					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.is_enabled_by_default"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.is_supported"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.0.summary"),

					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.1.is_enabled_by_default"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.2.is_enabled_by_default"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.3.is_enabled_by_default"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.4.is_enabled_by_default"),
					resource.TestCheckResourceAttrSet(datasourceName, "available_plugins.5.is_enabled_by_default"),
				),
			},
		},
	})
}
