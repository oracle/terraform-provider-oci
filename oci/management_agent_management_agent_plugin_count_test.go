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
	managementAgentPluginCountSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{RepType: Required, Create: `${var.compartment_id}`},
		"group_by":       Representation{RepType: Required, Create: `pluginName`},
	}

	ManagementAgentPluginCountResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentPluginCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentPluginCountResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_plugin_count.test_management_agent_plugin_count"

	SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugin_count", "test_management_agent_plugin_count", Required, Create, managementAgentPluginCountSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagementAgentPluginCountResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "group_by", "pluginName"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				),
			},
		},
	})
}
