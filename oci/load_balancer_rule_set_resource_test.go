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
	allowRuleSetRepresentation = map[string]interface{}{
		"items":            RepresentationGroup{Required, allowRuleItemsRepresentation},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{RepType: Required, Create: `example_rule_set`},
	}
	allowRuleSetRepresentationWithTwoItems = map[string]interface{}{
		"items":            []RepresentationGroup{{Required, allowRuleItemsRepresentation}, {Required, allowRuleItemsRepresentationWithTwoConditions}},
		"load_balancer_id": Representation{RepType: Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{RepType: Required, Create: `example_rule_set`},
	}

	httpResponseRuleItemsRepresentation = map[string]interface{}{
		"action": Representation{RepType: Required, Create: `ADD_HTTP_RESPONSE_HEADER`},
		"header": Representation{RepType: Optional, Create: `example_header_name`},
		"value":  Representation{RepType: Optional, Create: `example_header_value`},
	}
	allowRuleItemsRepresentation = map[string]interface{}{
		"action":      Representation{RepType: Required, Create: `ALLOW`},
		"description": Representation{RepType: Optional, Create: `description`},
		"conditions":  []RepresentationGroup{{Optional, itemsConditionsRepresentationSourceIPCondition}},
	}
	allowRuleItemsRepresentationWithTwoConditions = map[string]interface{}{
		"action":      Representation{RepType: Required, Create: `ALLOW`},
		"description": Representation{RepType: Optional, Create: `description`},
		"conditions":  []RepresentationGroup{{Optional, itemsConditionsRepresentationSourceVCNID}, {Optional, itemsConditionsRepresentationSourceVCNIP}},
	}

	itemsConditionsRepresentationSourceIPCondition = map[string]interface{}{
		"attribute_name":  Representation{RepType: Required, Create: `SOURCE_IP_ADDRESS`, Update: `SOURCE_VCN_ID`},
		"attribute_value": Representation{RepType: Required, Create: `129.0.0.0/8`, Update: `${oci_core_vcn.test_lb_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNID = map[string]interface{}{
		"attribute_name":  Representation{RepType: Required, Create: `SOURCE_VCN_ID`},
		"attribute_value": Representation{RepType: Required, Create: `${oci_core_vcn.test_lb_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNIP = map[string]interface{}{
		"attribute_name":  Representation{RepType: Required, Create: `SOURCE_VCN_IP_ADDRESS`},
		"attribute_value": Representation{RepType: Required, Create: `10.10.1.0/24`},
	}
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerRuleSetResource_allowAction(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerRuleSetResource_allowAction")
	defer httpreplay.SaveScenario()

	provider := TestAccProvider
	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_rule_set"
	datasourceName := "data.oci_load_balancer_rule_sets.test_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_rule_set"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { PreCheck() },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerRuleSetDestroy,
		Steps: []resource.TestStep{
			// Create with ADD_HTTP_RESPONSE_HEADER item
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update,
						GetUpdatedRepresentationCopy("items", RepresentationGroup{Required, httpResponseRuleItemsRepresentation}, allowRuleSetRepresentation)),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "ADD_HTTP_RESPONSE_HEADER",
						"header": "example_header_name",
						"value":  "example_header_value",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId, err = FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify Update to 1 item with 1 condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update, allowRuleSetRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update,
						GetUpdatedRepresentationCopy("items", RepresentationGroup{Required, allowRuleItemsRepresentationWithTwoConditions}, allowRuleSetRepresentation)),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = FromInstanceState(s, resourceName, "id")
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
					GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_rule_sets", Optional, Update, ruleSetDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "2"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
						"description":  "description",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
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
					GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "2"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
						"description":  "description",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
						"description":  "description",
					},
						[]string{}),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_rule_set"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
			},
			// verify resource import
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
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
