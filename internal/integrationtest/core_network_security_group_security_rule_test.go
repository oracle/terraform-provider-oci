// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	CoreNetworkSecurityGroupSecurityRuleRequiredOnlyResource = CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Required, acctest.Create, CoreNetworkSecurityGroupSecurityRuleRepresentation)

	NetworkSecurityGroupSecurityRuleResourceConfig = CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create, CoreNetworkSecurityGroupSecurityRuleRepresentation)

	CoreCoreNetworkSecurityGroupSecurityRuleDataSourceRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Optional, Create: `INGRESS`},
	}
	CoreNetworkSecurityGroupSecurityRuleRepresentation = map[string]interface{}{
		"network_security_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_network_security_group.test_network_security_group.id}`},
		"direction":                 acctest.Representation{RepType: acctest.Required, Create: `EGRESS`},
		"protocol":                  acctest.Representation{RepType: acctest.Required, Create: `6`},
		"description":               acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `updated description`},
	}

	CoreNetworkSecurityGroupSecurityEgressSecurityRulesRepresentation = map[string]interface{}{
		"direction":        acctest.Representation{RepType: acctest.Required, Create: `EGRESS`},
		"destination":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.0/16`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"destination_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `6`},
		"stateless":        acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRulesTcpOptionsRepresentation},
	}
	CoreNetworkSecurityGroupSecurityIngressSecurityRulesRepresentation = map[string]interface{}{
		"direction":   acctest.Representation{RepType: acctest.Required, Create: `INGRESS`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Optional, Create: `10.0.1.0/24`, Update: `${lookup(data.oci_core_services.test_services.services[0], "cidr_block")}`},
		"source_type": acctest.Representation{RepType: acctest.Optional, Create: `CIDR_BLOCK`, Update: `SERVICE_CIDR_BLOCK`},
		"stateless":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRulesTcpOptionsRepresentation},
	}
	CoreNetworkSecurityGroupSecurityRulesIcmpOptionsRepresentation = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Optional, Create: `4`, Update: `0`},
	}
	CoreNetworkSecurityGroupSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRulesTcpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRulesTcpOptionsSourcePortRangeRepresentation},
	}
	CoreNetworkSecurityGroupSecurityRulesUdpOptionsRepresentation = map[string]interface{}{
		"destination_port_range": acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRulesUdpOptionsDestinationPortRangeRepresentation},
		"source_port_range":      acctest.RepresentationGroup{RepType: acctest.Optional, Group: CoreNetworkSecurityGroupSecurityRuleSecurityRulesUdpOptionsDestinationPortRangeRepresentation},
	}

	CoreNetworkSecurityGroupSecurityRulesTcpOptionsSourcePortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	CoreNetworkSecurityGroupSecurityRuleSecurityRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	CoreNetworkSecurityGroupSecurityRulesTcpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}
	CoreNetworkSecurityGroupSecurityRulesUdpOptionsDestinationPortRangeRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `1521`, Update: `1522`},
	}

	CoreNetworkSecurityGroupSecurityRuleResourceDependencies = ObjectStorageCoreService +
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", acctest.Required, acctest.Create, CoreNetworkSecurityGroupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation)
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+CoreNetworkSecurityGroupSecurityRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityEgressSecurityRulesRepresentation)), "core", "networkSecurityGroupSecurityRule", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityEgressSecurityRulesRepresentation)),
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
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityEgressSecurityRulesRepresentation)),
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
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityIngressSecurityRulesRepresentation)),
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
			Config: config + compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityIngressSecurityRulesRepresentation)),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_core_network_security_group_security_rules", "test_network_security_group_security_rules", acctest.Optional, acctest.Update, CoreCoreNetworkSecurityGroupSecurityRuleDataSourceRepresentation) +
				compartmentIdVariableStr + CoreNetworkSecurityGroupSecurityRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_core_network_security_group_security_rule", "test_network_security_group_security_rule", acctest.Optional, acctest.Update,
					acctest.RepresentationCopyWithNewProperties(CoreNetworkSecurityGroupSecurityRuleRepresentation, CoreNetworkSecurityGroupSecurityIngressSecurityRulesRepresentation)),
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
			Config:                  config + CoreNetworkSecurityGroupSecurityRuleRequiredOnlyResource,
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
