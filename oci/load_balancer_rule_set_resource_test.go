// Copyright (c) 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	allowRuleSetRepresentation = map[string]interface{}{
		"items":            RepresentationGroup{Required, allowRuleItemsRepresentation},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_rule_set`},
	}
	allowRuleSetRepresentationWithTwoItems = map[string]interface{}{
		"items":            []RepresentationGroup{{Required, allowRuleItemsRepresentation}, {Required, allowRuleItemsRepresentationWithTwoConditions}},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_rule_set`},
	}

	httpResponseRuleItemsRepresentation = map[string]interface{}{
		"action": Representation{repType: Required, create: `ADD_HTTP_RESPONSE_HEADER`},
		"header": Representation{repType: Optional, create: `example_header_name`},
		"value":  Representation{repType: Optional, create: `example_header_value`},
	}
	allowRuleItemsRepresentation = map[string]interface{}{
		"action":      Representation{repType: Required, create: `ALLOW`},
		"description": Representation{repType: Optional, create: `description`},
		"conditions":  []RepresentationGroup{{Optional, itemsConditionsRepresentationSourceIPCondition}},
	}
	allowRuleItemsRepresentationWithTwoConditions = map[string]interface{}{
		"action":      Representation{repType: Required, create: `ALLOW`},
		"description": Representation{repType: Optional, create: `description`},
		"conditions":  []RepresentationGroup{{Optional, itemsConditionsRepresentationSourceVCNID}, {Optional, itemsConditionsRepresentationSourceVCNIP}},
	}

	itemsConditionsRepresentationSourceIPCondition = map[string]interface{}{
		"attribute_name":  Representation{repType: Required, create: `SOURCE_IP_ADDRESS`, update: `SOURCE_VCN_ID`},
		"attribute_value": Representation{repType: Required, create: `129.0.0.0/8`, update: `${oci_core_vcn.test_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNID = map[string]interface{}{
		"attribute_name":  Representation{repType: Required, create: `SOURCE_VCN_ID`},
		"attribute_value": Representation{repType: Required, create: `${oci_core_vcn.test_vcn.id}`},
	}
	itemsConditionsRepresentationSourceVCNIP = map[string]interface{}{
		"attribute_name":  Representation{repType: Required, create: `SOURCE_VCN_IP_ADDRESS`},
		"attribute_value": Representation{repType: Required, create: `10.10.1.0/24`},
	}
)

func TestLoadBalancerRuleSetResource_allowAction(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerRuleSetResource_allowAction")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_rule_set"
	datasourceName := "data.oci_load_balancer_rule_sets.test_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_rule_set"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerRuleSetDestroy,
		Steps: []resource.TestStep{
			// create with ADD_HTTP_RESPONSE_HEADER item
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update,
						getUpdatedRepresentationCopy("items", RepresentationGroup{Required, httpResponseRuleItemsRepresentation}, allowRuleSetRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify update to 1 item with 1 condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify update of the condition
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update, allowRuleSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update,
						getUpdatedRepresentationCopy("items", RepresentationGroup{Required, allowRuleItemsRepresentationWithTwoConditions}, allowRuleSetRepresentation)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":       "ALLOW",
						"conditions.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},
			// verify update to two items and 3 conditions total
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: resource.ComposeAggregateTestCheckFunc(
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
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_rule_sets", Optional, Update, ruleSetDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
			},
			// verify resource import
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Create, allowRuleSetRepresentationWithTwoItems),
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
