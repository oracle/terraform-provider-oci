// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	createSecurityRuleResourceConfig = securityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_security_rule",
			"test_network_firewall_policy_security_rule",
			acctest.Required, acctest.Create,
			securityRuleCreateRepresentation,
		)

	securityRuleResourceConfig = securityRuleResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_security_rule",
			"test_network_firewall_policy_security_rule",
			acctest.Optional, acctest.Update,
			securityRuleRepresentation,
		)

	securityRuleSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	securityRuleDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	securityRuleCreateRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `INSPECT`, Update: `DROP`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: conditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `security_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"inspection":                 acctest.Representation{RepType: acctest.Required, Create: `INTRUSION_DETECTION`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: rulePositionRepresentation},
	}

	securityRuleRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `INSPECT`, Update: `DROP`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Optional, Group: conditionRepresentationNull},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `security_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Required, Group: rulePositionRepresentation},
	}

	conditionRepresentation = map[string]interface{}{
		"application":         acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group.name}`}, Update: nil},
		"destination_address": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}},
		"service":             acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list.name}`}},
		"source_address":      acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}},
		"url":                 acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list.name}`}},
	}

	conditionRepresentationNull = map[string]interface{}{
		"application":         acctest.Representation{RepType: acctest.Required, Create: nil},
		"destination_address": acctest.Representation{RepType: acctest.Required, Create: nil},
		"service":             acctest.Representation{RepType: acctest.Required, Create: nil},
		"source_address":      acctest.Representation{RepType: acctest.Required, Create: nil},
		"url":                 acctest.Representation{RepType: acctest.Required, Create: nil},
	}

	rulePositionRepresentation = map[string]interface{}{
		"after_rule":  acctest.Representation{RepType: acctest.Required, Create: nil},
		"before_rule": acctest.Representation{RepType: acctest.Optional, Create: nil},
	}

	securityRuleResourceDependencies = securityRuleComponentDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)

	securityRuleComponentDependencies = createApplicationResource + createServiceResource +
		createAddressListResource +
		createApplicationGroupResource +
		createServiceListResource +
		createURLListResource
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicySecurityRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicySecurityRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_security_rules.test_network_firewall_policy_security_rules"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_security_rule.test_network_firewall_policy_security_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+createSecurityRuleResourceConfig,
		"networkfirewall",
		"networkFirewallPolicySecurityRule", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicySecurityRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createSecurityRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "INSPECT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "security_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.application.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.service.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.url.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + securityRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "DROP"),
				resource.TestCheckResourceAttr(resourceName, "name", "security_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_security_rules",
					"test_network_firewall_policy_security_rules",
					acctest.Optional, acctest.Update,
					securityRuleDataSourceRepresentation,
				) +
				compartmentIdVariableStr + securityRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "security_rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "security_rule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_security_rule",
					"test_network_firewall_policy_security_rule",
					acctest.Required, acctest.Create,
					securityRuleSingularDataSourceRepresentation,
				) +
				compartmentIdVariableStr + securityRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "DROP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.application.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.destination_address.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.service.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.source_address.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.url.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "security_rule_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "position.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + securityRuleResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicySecurityRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_security_rule" {
			noResourceFound = false
			request := oci_network_firewall.GetSecurityRuleRequest{}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.SecurityRuleName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetSecurityRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicySecurityRule") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicySecurityRule", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicySecurityRule",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicySecurityRule"],
			F:            sweepNetworkFirewallNetworkFirewallPolicySecurityRuleResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicySecurityRuleResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicySecurityRuleIds, err := getNetworkFirewallNetworkFirewallPolicySecurityRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicySecurityRuleId := range networkFirewallPolicySecurityRuleIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicySecurityRuleId]; !ok {
			deleteSecurityRuleRequest := oci_network_firewall.DeleteSecurityRuleRequest{}

			deleteSecurityRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteSecurityRule(context.Background(), deleteSecurityRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicySecurityRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicySecurityRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicySecurityRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicySecurityRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listSecurityRulesRequest := oci_network_firewall.ListSecurityRulesRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicySecurityRule resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listSecurityRulesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listSecurityRulesResponse, err := networkFirewallClient.ListSecurityRules(context.Background(), listSecurityRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicySecurityRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicySecurityRule := range listSecurityRulesResponse.Items {
			id := *networkFirewallPolicySecurityRule.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicySecurityRuleId", id)
		}

	}
	return resourceIds, nil
}
