// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	allowRuleSetRepresentation = map[string]interface{}{
		"items":            acctest.RepresentationGroup{RepType: acctest.Required, Group: allowRuleItemsRepresentation},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_rule_set`},
	}
	allowRuleSetRepresentationWithTwoItems = map[string]interface{}{
		"items":            []acctest.RepresentationGroup{{RepType: acctest.Required, Group: allowRuleItemsRepresentation}, {RepType: acctest.Required, Group: allowRuleItemsRepresentationWithTwoConditions}},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_rule_set`},
	}

	httpResponseRuleItemsRepresentation = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `ADD_HTTP_RESPONSE_HEADER`},
		"header": acctest.Representation{RepType: acctest.Optional, Create: `example_header_name`},
		"value":  acctest.Representation{RepType: acctest.Optional, Create: `example_header_value`},
	}
	allowRuleItemsRepresentation = map[string]interface{}{
		"action":      acctest.Representation{RepType: acctest.Required, Create: `ALLOW`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"conditions":  []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: itemsConditionsRepresentationSourceIPCondition}},
	}
	allowRuleItemsRepresentationWithTwoConditions = map[string]interface{}{
		"action":      acctest.Representation{RepType: acctest.Required, Create: `ALLOW`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`},
		"conditions":  []acctest.RepresentationGroup{{RepType: acctest.Optional, Group: itemsConditionsRepresentationSourceVCNID}, {RepType: acctest.Optional, Group: itemsConditionsRepresentationSourceVCNIP}},
	}

	itemsConditionsRepresentationSourceIPCondition = map[string]interface{}{
		"attribute_name":  acctest.Representation{RepType: acctest.Required, Create: `SOURCE_IP_ADDRESS`, Update: `SOURCE_VCN_ID`},
		"attribute_value": acctest.Representation{RepType: acctest.Required, Create: `129.0.0.0/8`, Update: `${oci_core_vcn.test_lb_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNID = map[string]interface{}{
		"attribute_name":  acctest.Representation{RepType: acctest.Required, Create: `SOURCE_VCN_ID`},
		"attribute_value": acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_lb_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNIP = map[string]interface{}{
		"attribute_name":  acctest.Representation{RepType: acctest.Required, Create: `SOURCE_VCN_IP_ADDRESS`},
		"attribute_value": acctest.Representation{RepType: acctest.Required, Create: `10.10.1.0/24`},
	}
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerRuleSetResource_allowAction(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerRuleSetResource_allowAction")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_rule_set"
	datasourceName := "data.oci_load_balancer_rule_sets.test_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_rule_set"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerRuleSetDestroy,
		Steps: []resource.TestStep{
			// Create with ADD_HTTP_RESPONSE_HEADER item
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("items", acctest.RepresentationGroup{RepType: acctest.Required, Group: httpResponseRuleItemsRepresentation}, allowRuleSetRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "ADD_HTTP_RESPONSE_HEADER",
						"header": "example_header_name",
						"value":  "example_header_value",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify Update to 1 item with 1 condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Create, allowRuleSetRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify Update of the condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Update, allowRuleSetRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify new added condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Update,
						acctest.GetUpdatedRepresentationCopy("items", acctest.RepresentationGroup{RepType: acctest.Required, Group: allowRuleItemsRepresentationWithTwoConditions}, allowRuleSetRepresentation)),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify Update to two items and 3 conditions total
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Create, allowRuleSetRepresentationWithTwoItems),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_rule_sets", acctest.Optional, acctest.Update, ruleSetDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Create, allowRuleSetRepresentationWithTwoItems),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
						"description":  "description",
					},
						[]string{}),
					acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
						"description":  "description",
					},
						[]string{}),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.name", "example_rule_set"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Create, allowRuleSetRepresentationWithTwoItems),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
						"description":  "description",
					},
						[]string{}),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
						"description":  "description",
					},
						[]string{}),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_rule_set"),
				),
			},
			// verify resource import
			{
				Config:            config + RuleSetRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"state",
				},
				ResourceName: resourceName,
			},
		},
	})
}
