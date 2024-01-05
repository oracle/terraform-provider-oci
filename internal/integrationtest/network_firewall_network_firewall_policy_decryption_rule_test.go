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
	createDecryptionRuleResourceConfig = decryptionRuleResourceDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_decryption_rule",
		"test_network_firewall_policy_decryption_rule",
		acctest.Required, acctest.Create,
		decryptionRuleCreateRepresentation,
	)

	decryptionRuleResourceConfig = decryptionRuleResourceDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_decryption_rule",
		"test_network_firewall_policy_decryption_rule",
		acctest.Optional, acctest.Update,
		decryptionRuleRepresentation,
	)

	decryptionRuleSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_decryption_rule.test_network_firewall_policy_decryption_rule.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	decryptionRuleDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	decryptionRuleCreateRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `DECRYPT`, Update: `NO_DECRYPT`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: decryptionRuleConditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `decryption_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"decryption_profile":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_decryption_profile.test_network_firewall_policy_decryption_profile.name}`, Update: nil},
		"secret":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_mapped_secret.test_network_firewall_policy_mapped_secret.name}`, Update: nil},
	}

	decryptionRuleRepresentation = map[string]interface{}{
		"action":                     acctest.Representation{RepType: acctest.Required, Create: `DECRYPT`, Update: `NO_DECRYPT`},
		"condition":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: decryptionRuleConditionRepresentation},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `decryption_rule_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"position":                   acctest.RepresentationGroup{RepType: acctest.Optional, Group: decryptionRulePositionRepresentation},
	}

	decryptionRuleConditionRepresentation = map[string]interface{}{
		"destination_address": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}, Update: []string{}},
		"source_address":      acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`}, Update: []string{}},
	}

	decryptionRulePositionRepresentation = map[string]interface{}{
		"after_rule":  acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
		"before_rule": acctest.Representation{RepType: acctest.Optional, Create: nil, Update: nil},
	}

	decryptionRuleResourceDependencies = decryptionRuleComponentDependencies + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)

	decryptionRuleComponentDependencies = createAddressListResource + createMappedSecretResource + vaultSecretResource + createDecryptionProfileResource
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyDecryptionRuleResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyDecryptionRuleResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	vaultId := utils.GetEnvSettingWithBlankDefault("kms_vault_id")
	vaultIdVariableStr := fmt.Sprintf("variable \"vault_id\" { default = \"%s\" }\n", vaultId)

	keyId := utils.GetEnvSettingWithBlankDefault("kms_key_id")
	keyIdVariableStr := fmt.Sprintf("variable \"key_id\" { default = \"%s\" }\n", keyId)

	variablesConfig := compartmentIdVariableStr + vaultIdVariableStr + keyIdVariableStr

	resourceName := "oci_network_firewall_network_firewall_policy_decryption_rule.test_network_firewall_policy_decryption_rule"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_decryption_rules.test_network_firewall_policy_decryption_rules"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_decryption_rule.test_network_firewall_policy_decryption_rule"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+variablesConfig+createDecryptionRuleResourceConfig, "networkfirewall", "networkFirewallPolicyDecryptionRule", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyDecryptionRuleDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + variablesConfig + createDecryptionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "DECRYPT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "decryption_rule_1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.destination_address.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "condition.0.source_address.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "decryption_profile", "decryption_profile_1"),
				resource.TestCheckResourceAttr(resourceName, "secret", "mapped_secret_1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + variablesConfig + decryptionRuleResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "action", "NO_DECRYPT"),
				resource.TestCheckResourceAttr(resourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "decryption_rule_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "position.#", "1"),

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
			Config: config + variablesConfig + decryptionRuleResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_decryption_rules",
					"test_network_firewall_policy_decryption_rules",
					acctest.Optional, acctest.Update,
					decryptionRuleDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "decryption_rule_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "decryption_rule_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + variablesConfig + decryptionRuleResourceConfig +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_decryption_rule",
					"test_network_firewall_policy_decryption_rule",
					acctest.Required, acctest.Create,
					decryptionRuleSingularDataSourceRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "action", "NO_DECRYPT"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.destination_address.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "condition.0.source_address.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "decryption_rule_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "position.#", "1"),
			),
		},
		// verify resource import
		{
			Config:                  config + createDecryptionRuleResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyDecryptionRuleDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_decryption_rule" {
			noResourceFound = false
			request := oci_network_firewall.GetDecryptionRuleRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.DecryptionRuleName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetDecryptionRule(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyDecryptionRule") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyDecryptionRule", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyDecryptionRule",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyDecryptionRule"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyDecryptionRuleResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyDecryptionRuleResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyDecryptionRuleIds, err := getNetworkFirewallNetworkFirewallPolicyDecryptionRuleIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyDecryptionRuleId := range networkFirewallPolicyDecryptionRuleIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyDecryptionRuleId]; !ok {
			deleteDecryptionRuleRequest := oci_network_firewall.DeleteDecryptionRuleRequest{}

			deleteDecryptionRuleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteDecryptionRule(context.Background(), deleteDecryptionRuleRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyDecryptionRule %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyDecryptionRuleId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyDecryptionRuleIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyDecryptionRuleId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listDecryptionRulesRequest := oci_network_firewall.ListDecryptionRulesRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyDecryptionRule resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listDecryptionRulesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listDecryptionRulesResponse, err := networkFirewallClient.ListDecryptionRules(context.Background(), listDecryptionRulesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyDecryptionRule list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyDecryptionRule := range listDecryptionRulesResponse.Items {
			id := *networkFirewallPolicyDecryptionRule.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyDecryptionRuleId", id)
		}

	}
	return resourceIds, nil
}
