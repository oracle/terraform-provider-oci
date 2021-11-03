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
	hostInsightRequiredRepresentation = map[string]interface{}{
		"compartment_id":      Representation{RepType: Required, Create: `${var.compartment_id}`},
		"entity_source":       Representation{RepType: Required, Create: `MACS_MANAGED_EXTERNAL_HOST`},
		"management_agent_id": Representation{RepType: Required, Create: `${var.managed_agent_id}`},
		"status":              Representation{RepType: Required, Create: `DISABLED`},
		"defined_tags":        Representation{RepType: Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`},
		"freeform_tags":       Representation{RepType: Optional, Create: map[string]string{"bar-key": "value"}},
		"lifecycle":           RepresentationGroup{Required, ignoreHIChangesRepresentation},
	}

	ignoreHIChangesRepresentation = map[string]interface{}{
		"ignore_changes": Representation{RepType: Required, Create: []string{`defined_tags`}},
	}
)

// This test basically covers the required field
// issue-routing-tag: opsi/controlPlane
func TestResourceOpsiHostInsight(t *testing.T) {
	httpreplay.SetScenario("TestResourceOpsiHostInsight")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	managementAgentId := GetEnvSettingWithBlankDefault("managed_agent_id")
	if managementAgentId == "" {
		t.Skip("Manual install agent and set managed_agent_id to run this test")
	}
	managementAgentIdVariableStr := fmt.Sprintf("variable \"managed_agent_id\" { default = \"%s\" }\n", managementAgentId)

	resourceName := "oci_opsi_host_insight.test_host_insight"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+managementAgentIdVariableStr+HostInsightResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", Required, Create, hostInsightRequiredRepresentation), "opsi", "hostInsight", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOpsiHostInsightDestroy,
		Steps: []resource.TestStep{
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + managementAgentIdVariableStr + HostInsightResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_opsi_host_insight", "test_host_insight", Required, Create, hostInsightRequiredRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "status", "DISABLED"),

					func(s *terraform.State) (err error) {
						_, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}
