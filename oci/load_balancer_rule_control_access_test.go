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
	ruleSetControlAccessSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_control_access_rule_set`},
	}

	ruleSetControlAccessDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, ruleSetControlAccessDataSourceFilterRepresentation}}
	ruleSetControlAccessDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_rule_set.test_control_access_rule_set.name}`}},
	}

	ruleSetControlAccessRepresentation = map[string]interface{}{
		"items":            RepresentationGroup{Required, ruleSetItemsControlAccessRepresentation},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_control_access_rule_set`},
	}

	ruleSetItemsControlAccessRepresentation = map[string]interface{}{
		"action":          Representation{repType: Required, create: `CONTROL_ACCESS_USING_HTTP_METHODS`},
		"allowed_methods": Representation{repType: Required, create: []string{`GET`}, update: []string{`GET`, `POST`}},
		"status_code":     Representation{repType: Optional, create: `405`, update: `400`},
	}

	ruleSetItemsAnotherControlAccessRepresentation = map[string]interface{}{
		"action":          Representation{repType: Required, create: `CONTROL_ACCESS_USING_HTTP_METHODS`},
		"allowed_methods": Representation{repType: Required, create: []string{`GET`}, update: []string{`GET`, `POST`, `PUT`}},
		"status_code":     Representation{repType: Optional, create: `405`, update: `400`},
	}
)

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerRuleSetResource_controlAccess_test(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerRuleSetResource_controlAccess_test")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_control_access_rule_set"

	var resId, resId2 string
	datasourceName := "data.oci_load_balancer_rule_sets.test_control_access_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_control_access_rule_set"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerRuleSetDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Required, Create, ruleSetControlAccessRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "405",
						"allowed_methods.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update, ruleSetControlAccessRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters allowed_methods only
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update,
						representationCopyWithNewProperties(ruleSetControlAccessRepresentation, map[string]interface{}{
							"items": RepresentationGroup{Required, ruleSetItemsAnotherControlAccessRepresentation},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "3",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_control_access_rule_sets", Optional, Update, ruleSetControlAccessDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update, ruleSetControlAccessRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "1"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.name", "example_control_access_rule_set"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Required, Create, ruleSetControlAccessSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update, ruleSetControlAccessRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_control_access_rule_set"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "2",
					},
						[]string{}),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update, ruleSetControlAccessRepresentation),
			},
			// verify resource import
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", Optional, Update, ruleSetControlAccessRepresentation),
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
