// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceAgentPluginDataSourceRepresentation = map[string]interface{}{
		"instanceagent_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_instance.test_instance.id}`},
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	InstanceAgentPluginResourceConfig = SubnetResourceConfig + utils.OciImageIdsVariable +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(acctest.RepresentationCopyWithRemovedProperties(instanceRepresentation, []string{"agent_config"}), map[string]interface{}{
			"agent_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: instanceAgentConfigRepresentation},
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_agent_plugins", "test_instance_agent_plugins", acctest.Required, acctest.Create, instanceAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceAgentPluginResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "instanceagent_id"),
			),
		},
	})
}
