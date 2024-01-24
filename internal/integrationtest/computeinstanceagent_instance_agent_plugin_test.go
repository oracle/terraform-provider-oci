// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	ComputeinstanceagentComputeinstanceagentInstanceAgentPluginDataSourceRepresentation = map[string]interface{}{
		"instanceagent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	ComputeinstanceagentInstanceAgentPluginResourceConfig = CoreSubnetResourceConfig + utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(CoreInstanceRepresentation, []string{"agent_config"}), map[string]interface{}{
			"agent_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: CoreInstanceAgentConfigRepresentation},
		}))
)

// issue-routing-tag: computeinstanceagent/default
func TestComputeinstanceagentInstanceAgentPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeinstanceagentInstanceAgentPluginResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_computeinstanceagent_instance_agent_plugins.test_instance_agent_plugins"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_agent_plugins", "test_instance_agent_plugins", acctest.Required, acctest.Create, ComputeinstanceagentComputeinstanceagentInstanceAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + ComputeinstanceagentInstanceAgentPluginResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "instanceagent_id"),
			),
		},
	})
}
