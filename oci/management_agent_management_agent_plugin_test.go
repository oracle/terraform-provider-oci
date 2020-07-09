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
	managementAgentPluginDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
	}

	ManagementAgentPluginResourceConfig = ""
)

func TestManagementAgentManagementAgentPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentPluginResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_management_agent_management_agent_plugins.test_management_agent_plugins"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugins", "test_management_agent_plugins", Required, Create, managementAgentPluginDataSourceRepresentation) +
					compartmentIdVariableStr + ManagementAgentPluginResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.display_name"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.is_console_deployable"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.name"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.supported_platform_types.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "management_agent_plugins.0.version"),
				),
			},
		},
	})
}
