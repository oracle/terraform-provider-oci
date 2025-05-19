// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
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
	natRuleRequiredOnlyResource = natRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Required, acctest.Create, natRuleCreateRepresentation)

	natRuleResourceConfig = natRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Optional, acctest.Update, natRuleRepresentation)

	natRuleSingularDataSourceRepresentation = map[string]interface{}{
		"nat_rule_name":              acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	natRuleDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"nat_rule_priority_order":    acctest.Representation{RepType: acctest.Optional, Create: `1`},
	}

	natRuleRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `DIPP_SRC_NAT`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: natRuleConditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `nat_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `NATV4`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: natRulePositionRepresentation},
	}
	natRuleCreateRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `DIPP_SRC_NAT`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: natRuleConditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `nat_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `NATV4`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: natRulePositionRepresentation},
	}
	natRuleConditionRepresentation = map[string]interface{}{
		"destination_address": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}},
		"source_address":      acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}},
		"service":             acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service.name}`},
	}
	natRulePositionRepresentation = map[string]interface{}{
		"after_rule":  acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"before_rule": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
	}

	natRuleResourceDependencies = createAddressListResource + createServiceResource + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required,
		acctest.Create,
		networkFirewallPolicyRepresentation)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyNatRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyNatRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_nat_rules.test_network_firewall_policy_nat_rules"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_nat_rule.test_network_firewall_policy_nat_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+natRuleResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Optional, acctest.Create, natRuleCreateRepresentation), "networkfirewall", "networkFirewallPolicyNatRule", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyNatRuleDestroy, []resource.TestStep{
		// verify Create step 0
		{
			Config: config + compartmentIdVariableStr + natRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Required, acctest.Create, natRuleCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "DIPP_SRC_NAT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "nat_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "NATV4"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create Step 1
		{
			Config: config + compartmentIdVariableStr + natRuleResourceDependencies,
		},
		// verify Create with optionals Step 2
		{
			Config: config + compartmentIdVariableStr + natRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Optional, acctest.Create, natRuleCreateRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "DIPP_SRC_NAT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.0", "address_list_1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.0", "address_list_1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "name", "nat_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "NATV4"),

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

		// verify updates to updatable parameters step 3
		{
			Config: config + compartmentIdVariableStr + natRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Optional, acctest.Update, natRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "DIPP_SRC_NAT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "name", "nat_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "type", "NATV4"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource Step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rules", "test_network_firewall_policy_nat_rules", acctest.Optional, acctest.Update, natRuleDataSourceRepresentation) +
				compartmentIdVariableStr + natRuleResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Optional, acctest.Update, natRuleRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "nat_rule_priority_order", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "nat_rule_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "nat_rule_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_nat_rule", "test_network_firewall_policy_nat_rule", acctest.Required, acctest.Create, natRuleSingularDataSourceRepresentation) +
				compartmentIdVariableStr + natRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "nat_rule_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "action", "DIPP_SRC_NAT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "nat_rule_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "position.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "priority_order"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "NATV4"),
			),
		},
		// verify resource import Step 6
		{
			Config:                  config + natRuleRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyNatRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_nat_rule" {
			noResourceFound = false
			request := oci_network_firewall.GetNatRuleRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.NatRuleName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetNatRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyNatRule") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyNatRule", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyNatRule",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyNatRule"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyNatRuleResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyNatRuleResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyNatRuleIds, err := getNetworkFirewallNetworkFirewallPolicyNatRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyNatRuleId := range networkFirewallPolicyNatRuleIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyNatRuleId]; !ok {
			deleteNatRuleRequest := oci_network_firewall.DeleteNatRuleRequest{}

			deleteNatRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteNatRule(context.Background(), deleteNatRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyNatRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyNatRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyNatRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyNatRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listNatRulesRequest := oci_network_firewall.ListNatRulesRequest{}
	// listNatRulesRequest.CompartmentId = &compartmentId

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyNatRule resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listNatRulesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listNatRulesResponse, err := networkFirewallClient.ListNatRules(context.Background(), listNatRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyNatRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyNatRule := range listNatRulesResponse.Items {
			id := *networkFirewallPolicyNatRule.GetName()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyNatRuleId", id)
		}

	}
	return resourceIds, nil
}
