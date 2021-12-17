// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	NetworkSecurityGroupSecurityRuleRequiredOnlyResource = NetworkSecurityGroupSecurityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Required, acctest.Create, networkSecurityGroupSecurityRuleRepresentation)

	NetworkSecurityGroupSecurityRuleResourceConfig = NetworkSecurityGroupSecurityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create, networkSecurityGroupSecurityRuleRepresentation)

	networkSecurityGroupSecurityRuleDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Optional, Create: `INGRESS`},
	}
	networkSecurityGroupSecurityRuleRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Required, Create: `EGRESS`},
		"protocol":                  acctest.Representation{RepType: acctest.Required, Create: `6`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `updated description`},
	}

	egressSecurityRulesRepresentation = map[string]interface{}{
		"direction":        acctest.Representation{RepType: acctest.Required, Create: `EGRESS`},
		"destination":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.0/16`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `6`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesTcpOptionsRepresentation},
	}
	ingressSecurityRulesRepresentation = map[string]interface{}{
		"direction":   acctest.Representation{RepType: acctest.Required, Create: `INGRESS`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesTcpOptionsRepresentation},
	}
	securityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `0`},
	}
	securityRulesTcpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesTcpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityRulesUdpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesUdpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: securityRulesUdpOptionsSourcePortRangeRepresentation},
	}

	securityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	securityRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}

	NetworkSecurityGroupSecurityRuleResourceDependencies = ObjectStorageCoreService +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, networkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreNetworkSecurityGroupSecurityRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupSecurityRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_network_security_group_security_rule.test_network_security_group_security_rule"
	datasourceName := "data.oci_core_network_security_group_security_rules.test_network_security_group_security_rules"

	var resId, resId2, compositeId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NetworkSecurityGroupSecurityRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)), "core", "networkSecurityGroupSecurityRule", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "direction", "EGRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_security_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttrSet(resourceName, "destination"),
				resource.TestCheckResourceAttr(resourceName, "destination_type", "CIDR_BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "stateless"),
				resource.TestCheckResourceAttr(resourceName, "tcp_options.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					networkSecurityGroupId, _ := acctest.FromInstanceState(s, resourceName, "network_security_group_id")
					compositeId = "networkSecurityGroups/" + networkSecurityGroupId + "/securityRules/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "direction", "EGRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_security_group_id"),

				resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
				resource.TestCheckResourceAttrSet(resourceName, "destination"),
				resource.TestCheckResourceAttr(resourceName, "destination_type", "SERVICE_CIDR_BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "stateless"),
				resource.TestCheckResourceAttr(resourceName, "tcp_options.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updatedr")
					}
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "direction", "INGRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_security_group_id"),
				resource.TestCheckResourceAttrSet(resourceName, "description"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "CIDR_BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "stateless"),
				resource.TestCheckResourceAttr(resourceName, "tcp_options.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "direction", "INGRESS"),
				resource.TestCheckResourceAttrSet(resourceName, "network_security_group_id"),

				resource.TestCheckResourceAttr(resourceName, "description", "updated description"),
				resource.TestCheckResourceAttrSet(resourceName, "source"),
				resource.TestCheckResourceAttr(resourceName, "source_type", "SERVICE_CIDR_BLOCK"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "protocol"),
				resource.TestCheckResourceAttrSet(resourceName, "stateless"),
				resource.TestCheckResourceAttr(resourceName, "tcp_options.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updatedr")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_network_security_group_security_rules", "test_network_security_group_security_rules", acctest.Optional, acctest.Update, networkSecurityGroupSecurityRuleDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "direction"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_security_group_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.description"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.direction"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.is_valid"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.protocol"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.stateless"),
				resource.TestCheckResourceAttr(datasourceName, "security_rules.0.tcp_options.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "security_rules.0.time_created"),
			),
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateIdFunc:       getNetworkSecurityGroupSecurityRuleImportId(resourceName),
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func getNetworkSecurityGroupSecurityRuleImportId(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("not found: %s", resourceName)
		}
		return fmt.Sprintf("networkSecurityGroups/" + rs.Primary.Attributes["network_security_group_id"] + "/securityRules/" + rs.Primary.Attributes["id"]), nil
	}
}
