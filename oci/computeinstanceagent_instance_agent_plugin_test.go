// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	instanceAgentPluginDataSourceRepresentation = map[string]interface{}{
		"instanceagent_id": Representation{RepType: Required, Create: `${oci_core_instance.test_instance.id}`},
		"compartment_id":   Representation{RepType: Required, Create: `${var.compartment_id}`},
	}

	InstanceAgentPluginResourceConfig = SubnetResourceConfig + OciImageIdsVariable +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, RepresentationCopyWithNewProperties(RepresentationCopyWithRemovedProperties(instanceRepresentation, []string{"agent_config"}), map[string]interface{}{
			"agent_config": RepresentationGroup{Required, instanceAgentConfigRepresentation},
		}))
)

// issue-routing-tag: computeinstanceagent/default
func TestComputeinstanceagentInstanceAgentPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeinstanceagentInstanceAgentPluginResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_computeinstanceagent_instance_agent_plugins.test_instance_agent_plugins"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_agent_plugins", "test_instance_agent_plugins", Required, Create, instanceAgentPluginDataSourceRepresentation) +
				compartmentIdVariableStr + InstanceAgentPluginResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "instanceagent_id"),
			),
		},
	})
}
