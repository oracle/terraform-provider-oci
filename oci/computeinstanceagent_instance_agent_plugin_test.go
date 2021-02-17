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
	instanceAgentPluginDataSourceRepresentation = map[string]interface{}{
		"instanceagent_id": Representation{repType: Required, create: `${oci_core_instance.test_instance.id}`},
		"compartment_id":   Representation{repType: Required, create: `${var.compartment_id}`},
	}

	InstanceAgentPluginResourceConfig = SubnetResourceConfig + OciImageIdsVariable +
		generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_instance", "test_instance", Required, Create, representationCopyWithNewProperties(representationCopyWithRemovedProperties(instanceRepresentation, []string{"agent_config"}), map[string]interface{}{
			"agent_config": RepresentationGroup{Required, instanceAgentConfigRepresentation},
		}))
)

func TestComputeinstanceagentInstanceAgentPluginResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestComputeinstanceagentInstanceAgentPluginResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_computeinstanceagent_instance_agent_plugins.test_instance_agent_plugins"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_computeinstanceagent_instance_agent_plugins", "test_instance_agent_plugins", Required, Create, instanceAgentPluginDataSourceRepresentation) +
					compartmentIdVariableStr + InstanceAgentPluginResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "instanceagent_id"),
				),
			},
		},
	})
}
