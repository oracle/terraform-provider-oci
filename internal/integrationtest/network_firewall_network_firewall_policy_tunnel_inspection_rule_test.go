// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	tunnelInspectionRuleResourceConfig = tunnelInspectionRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_tunnel_inspection_rule",
			"test_network_firewall_policy_tunnel_inspection_rule",
			acctest.Optional,
			acctest.Update,
			tunnelInspectionRuleRepresentation)

	tunnelInspectionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id":  acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"tunnel_inspection_rule_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_tunnel_inspection_rule.test_network_firewall_policy_tunnel_inspection_rule.name}`},
	}

	tunnelInspectionRuleDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	tunnelInspectionRuleRepresentation = map[string]interface{}{
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: tunnelInspectionRuleConditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `tunnel_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"protocol":                   acctest.Representation{RepType: acctest.Required, Create: `VXLAN`},
		"action":                     acctest.Representation{RepType: acctest.Optional, Create: `INSPECT`, Update: `INSPECT_AND_CAPTURE_LOG`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: tunnelInspectionRulePositionRepresentation},
		"profile":                    acctest.RepresentationGroup{RepType: acctest.Optional, Group: tunnelInspectionRuleProfileRepresentation},
	}
	tunnelInspectionRuleConditionRepresentation = map[string]interface{}{
		"destination_address": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}, Update: []string{}},
		"source_address":      acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}, Update: []string{}},
	}
	tunnelInspectionRulePositionRepresentation = map[string]interface{}{
		"after_rule":  acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"before_rule": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
	}
	tunnelInspectionRuleProfileRepresentation = map[string]interface{}{
		"must_return_traffic_to_source": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	tunnelInspectionRuleResourceDependencies = createAddressListResource + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule.test_network_firewall_policy_tunnel_inspection_rule"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_tunnel_inspection_rules.test_network_firewall_policy_tunnel_inspection_rules"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_tunnel_inspection_rule.test_network_firewall_policy_tunnel_inspection_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+tunnelInspectionRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Optional, acctest.Create, tunnelInspectionRuleRepresentation), "networkfirewall", "networkFirewallPolicyTunnelInspectionRule", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDestroy, []resource.TestStep{
		// verify Create - step 0
		{
			Config: config + compartmentIdVariableStr + tunnelInspectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Required, acctest.Create, tunnelInspectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "tunnel_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "VXLAN"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create - step 1
		{
			Config: config + compartmentIdVariableStr + tunnelInspectionRuleResourceDependencies,
		},
		// verify Create with optionals - step 2
		{
			Config: config + compartmentIdVariableStr + tunnelInspectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Optional, acctest.Create, tunnelInspectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "INSPECT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "tunnel_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "profile.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "profile.0.must_return_traffic_to_source", "false"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "VXLAN"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters - - step 3
		{
			Config: config + compartmentIdVariableStr + tunnelInspectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Optional, acctest.Update, tunnelInspectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "INSPECT_AND_CAPTURE_LOG"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "name", "tunnel_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "profile.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "profile.0.must_return_traffic_to_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "VXLAN"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource - step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rules", "test_network_firewall_policy_tunnel_inspection_rules", acctest.Optional, acctest.Update, tunnelInspectionRuleDataSourceRepresentation) +
				compartmentIdVariableStr + tunnelInspectionRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Optional, acctest.Update, tunnelInspectionRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "tunnel_inspection_rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "tunnel_inspection_rule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource - step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_tunnel_inspection_rule", "test_network_firewall_policy_tunnel_inspection_rule", acctest.Required, acctest.Create, tunnelInspectionRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + tunnelInspectionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "tunnel_inspection_rule_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "INSPECT_AND_CAPTURE_LOG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "tunnel_rule_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "position.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "priority_order"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "profile.0.must_return_traffic_to_source", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "VXLAN"),
			),
		},
		// verify resource import - step 6
		{
			Config:                  config + tunnelInspectionRuleResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_tunnel_inspection_rule" {
			noResourceFound = false
			request := oci_network_firewall.GetTunnelInspectionRuleRequest{}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.TunnelInspectionRuleName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetTunnelInspectionRule(context.Background(), request)

			if err == nil {
				return fmt.Errorf("resource still exists")
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyTunnelInspectionRule") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyTunnelInspectionRule", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyTunnelInspectionRule",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyTunnelInspectionRule"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyTunnelInspectionRuleIds, err := getNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyTunnelInspectionRuleId := range networkFirewallPolicyTunnelInspectionRuleIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyTunnelInspectionRuleId]; !ok {
			deleteTunnelInspectionRuleRequest := oci_network_firewall.DeleteTunnelInspectionRuleRequest{}

			deleteTunnelInspectionRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteTunnelInspectionRule(context.Background(), deleteTunnelInspectionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyTunnelInspectionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyTunnelInspectionRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyTunnelInspectionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listTunnelInspectionRulesRequest := oci_network_firewall.ListTunnelInspectionRulesRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyTunnelInspectionRule resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listTunnelInspectionRulesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listTunnelInspectionRulesResponse, err := networkFirewallClient.ListTunnelInspectionRules(context.Background(), listTunnelInspectionRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyTunnelInspectionRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyTunnelInspectionRule := range listTunnelInspectionRulesResponse.Items {
			id := *networkFirewallPolicyTunnelInspectionRule.GetName()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyTunnelInspectionRuleId", id)
		}
	}
	return resourceIds, nil
}
