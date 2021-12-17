// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	networkSecurityGroupSecurityRuleResourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Required, Create: `EGRESS`},
		"protocol":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `updated description`},
		"destination":               acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.0/24`},
	}

	networkSecurityGroupIngressSecurityRuleResourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Required, Create: `INGRESS`},
		"protocol":                  acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":                    acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"icmp_options":              acctest.RepresentationGroup{RepType: acctest.Optional, Group: nsgSecurityRulesIcmpOptionsRepresentation},
		"source_type":               acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":                 acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	networkSecurityGroupIngressSecurityRuleUDPResourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Required, Create: `INGRESS`},
		"protocol":                  acctest.Representation{RepType: acctest.Required, Create: `17`},
		"source":                    acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type":               acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":                 acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"udp_options":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesUdpOptionsRepresentation},
	}

	nsgSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
	}
)

// issue-routing-tag: core/virtualNetwork
func TestAccResourceCoreNetworkSecurityGroupSecurityRule_scenarios(t *testing.T) {
	httpreplay.SetScenario("TestAccResourceCoreNetworkSecurityGroupSecurityRule_multipleRules")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_network_security_group_security_rule.test_network_security_group_security_rule"

	var resId1, resId2 [10]string

	acctest.ResourceTest(t, nil, []resource.TestStep{

		//verify Create 10 rules
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleResourceRepresentation, map[string]interface{}{
						"count": acctest.Representation{RepType: acctest.Optional, Create: `10`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					for i := 0; i < 10; i++ {
						resId, err := acctest.FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "id")
						if resId == "" {
							return err
						}
						resId1[i] = resId
					}
					return nil
				},
			),
		},
		//verify Update 10 rules
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleResourceRepresentation, map[string]interface{}{
						"count": acctest.Representation{RepType: acctest.Optional, Create: `10`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					for i := 0; i < 10; i++ {

						resId, err := acctest.FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "id")
						if resId == "" {
							return err
						}
						resId2[i] = resId

						if resId1[i] != resId2[i] {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						description, err := acctest.FromInstanceState(s, fmt.Sprintf("%s.%d", resourceName, i), "description")
						if description == "" {
							return err
						}
						if description != "updated description" {
							return fmt.Errorf("%s: Attribute 'description' expected \"updated description\", got %s", fmt.Sprintf("%s.%d", resourceName, i), description)
						}
					}
					return nil
				},
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		// Create rule without specifying `code` in icmp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create, networkSecurityGroupIngressSecurityRuleResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "icmp_options.0.code", "-1"),
			),
		},
		// Update rule without specifying code in icmp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update, networkSecurityGroupIngressSecurityRuleResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_options.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "icmp_options.0.code", "-1"),
			),
		},
		// delete
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		// Create rule without specifying `code` in udp options
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create, networkSecurityGroupIngressSecurityRuleUDPResourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "udp_options.#", "1"),
			),
		},
	})
}
