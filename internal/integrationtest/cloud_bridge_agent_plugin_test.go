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
	CloudBridgeCloudBridgeAgentPluginSingularDataSourceRepresentation = map[string]interface{}{
		"agent_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.agentId}`},
		"plugin_name": acctest.Representation{RepType: acctest.Required, Create: `gomon`},
	}
	CloudBridgeAgentPluginResourceDependencies = ""
)

// issue-routing-tag: cloud_bridge/default
func TestCloudBridgeAgentPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCloudBridgeAgentPluginResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	agentId := utils.GetEnvSettingWithBlankDefault("agentId")
	agentIdVariableStr := fmt.Sprintf("variable \"agentId\" { default = \"%s\" }\n", agentId)

	variableStr := compartmentIdVariableStr + agentIdVariableStr

	singularDatasourceName := "data.oci_cloud_bridge_agent_plugin.test_agent_plugin"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variableStr+CloudBridgeAgentPluginResourceDependencies, "cloudbridge", "agentPlugin", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_cloud_bridge_agent_plugin", "test_agent_plugin", acctest.Required, acctest.Create, CloudBridgeCloudBridgeAgentPluginSingularDataSourceRepresentation) +
				variableStr + CloudBridgeAgentPluginResourceDependencies,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "agent_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "plugin_name", "gomon"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "plugin_version"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
	})
}
