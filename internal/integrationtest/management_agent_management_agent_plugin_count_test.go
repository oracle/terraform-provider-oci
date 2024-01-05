// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	managementAgentPluginCountSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"group_by":       acctest.Representation{RepType: acctest.Required, Create: `pluginName`},
	}

	ManagementAgentPluginCountResourceConfig = ""
)

// issue-routing-tag: management_agent/default
func TestManagementAgentManagementAgentPluginCountResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestManagementAgentManagementAgentPluginCountResource_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_management_agent_management_agent_plugin_count.test_management_agent_plugin_count"

	acctest.SaveConfigContent("", "", "", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_management_agent_management_agent_plugin_count", "test_management_agent_plugin_count", acctest.Required, acctest.Create, managementAgentPluginCountSingularDataSourceRepresentation) +
					compartmentIdVariableStr + ManagementAgentPluginCountResourceConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "group_by", "pluginName"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				),
			},
		},
	})
}
