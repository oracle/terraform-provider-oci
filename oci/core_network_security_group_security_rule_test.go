// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NetworkSecurityGroupSecurityRuleRequiredOnlyResource = NetworkSecurityGroupSecurityRuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Required, Create, networkSecurityGroupSecurityRuleRepresentation)

	NetworkSecurityGroupSecurityRuleResourceConfig = NetworkSecurityGroupSecurityRuleResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create, networkSecurityGroupSecurityRuleRepresentation)

	networkSecurityGroupSecurityRuleDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 Representation{RepType: Optional, Create: `INGRESS`},
	}
	networkSecurityGroupSecurityRuleRepresentation = map[string]interface{}{
		"network_security_group_id": Representation{RepType: Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 Representation{RepType: Required, Create: `EGRESS`},
		"protocol":                  Representation{RepType: Required, Create: `6`},
		"description":               Representation{RepType: Optional, Create: `description`, Update: `updated description`},
	}

	egressSecurityRulesRepresentation = map[string]interface{}{
		"direction":        Representation{RepType: Required, Create: `EGRESS`},
		"destination":      Representation{RepType: Optional, Create: `10.0.0.0/16`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"protocol":         Representation{RepType: Required, Create: `6`},
		"stateless":        Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options":      RepresentationGroup{Optional, securityRulesTcpOptionsRepresentation},
	}
	ingressSecurityRulesRepresentation = map[string]interface{}{
		"direction":   Representation{RepType: Required, Create: `INGRESS`},
		"protocol":    Representation{RepType: Required, Create: `6`},
		"source":      Representation{RepType: Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": Representation{RepType: Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   Representation{RepType: Optional, Create: `false`, Update: `true`},
		"tcp_options": RepresentationGroup{Optional, securityRulesTcpOptionsRepresentation},
	}
	securityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": Representation{RepType: Required, Create: `3`},
		"code": Representation{RepType: Optional, Create: `4`, Update: `0`},
	}
	securityRulesTcpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": RepresentationGroup{Optional, securityRulesTcpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      RepresentationGroup{Optional, securityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	securityRulesUdpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": RepresentationGroup{Optional, securityRulesUdpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      RepresentationGroup{Optional, securityRulesUdpOptionsSourcePortRangeRepresentation},
	}

	securityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityRulesUdpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}
	securityRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": Representation{RepType: Required, Create: `1521`, Update: `1522`},
		"min": Representation{RepType: Required, Create: `1521`, Update: `1522`},
	}

	NetworkSecurityGroupSecurityRuleResourceDependencies = ObjectStorageCoreService +
		GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, VcnRepresentation)
)

// issue-routing-tag: core/virtualNetwork
func TestCoreNetworkSecurityGroupSecurityRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestCoreNetworkSecurityGroupSecurityRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_core_network_security_group_security_rule.test_network_security_group_security_rule"
	datasourceName := "data.oci_core_network_security_group_security_rules.test_network_security_group_security_rules"

	var resId, resId2, compositeId string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	SaveConfigContent(config+compartmentIdVariableStr+NetworkSecurityGroupSecurityRuleResourceDependencies+
		GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create,
			RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)), "core", "networkSecurityGroupSecurityRule", t)

	ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					networkSecurityGroupId, _ := FromInstanceState(s, resourceName, "network_security_group_id")
					compositeId = "networkSecurityGroups/" + networkSecurityGroupId + "/securityRules/" + resId
					if isEnableExportCompartment, _ := strconv.ParseBool(GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := TestExportCompartmentWithResourceName(&compositeId, &compartmentId, resourceName); errExport != nil {
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
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Update,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, egressSecurityRulesRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Create,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId, err = FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Update,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
					resId2, err = FromInstanceState(s, resourceName, "id")
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
				GenerateDataSourceFromRepresentationMap("oci_core_network_security_group_security_rules", "test_network_security_group_security_rules", Optional, Update, networkSecurityGroupSecurityRuleDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkSecurityGroupSecurityRuleResourceDependencies +
				GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", Optional, Update,
					RepresentationCopyWithNewProperties(networkSecurityGroupSecurityRuleRepresentation, ingressSecurityRulesRepresentation)),
			Check: ComposeAggregateTestCheckFuncWrapper(
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
