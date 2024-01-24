// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	RuleSetControlAccessRequiredOnlyRepresentation = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Required, acctest.Create, ruleSetControlAccessRepresentation)

	ruleSetControlAccessSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_control_access_rule_set`},
	}

	ruleSetControlAccessDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleSetControlAccessDataSourceFilterRepresentation}}
	ruleSetControlAccessDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_rule_set.test_control_access_rule_set.name}`}},
	}

	ruleSetControlAccessRepresentation = map[string]interface{}{
		"items":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleSetItemsControlAccessRepresentation},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_control_access_rule_set`},
	}

	ruleSetItemsControlAccessRepresentation = map[string]interface{}{
		"action":          acctest.Representation{RepType: acctest.Required, Create: `CONTROL_ACCESS_USING_HTTP_METHODS`},
		"allowed_methods": acctest.Representation{RepType: acctest.Required, Create: []string{`GET`}, Update: []string{`GET`, `POST`}},
		"status_code":     acctest.Representation{RepType: acctest.Optional, Create: `405`, Update: `400`},
	}

	ruleSetItemsAnotherControlAccessRepresentation = map[string]interface{}{
		"action":          acctest.Representation{RepType: acctest.Required, Create: `CONTROL_ACCESS_USING_HTTP_METHODS`},
		"allowed_methods": acctest.Representation{RepType: acctest.Required, Create: []string{`GET`}, Update: []string{`GET`, `POST`, `PUT`}},
		"status_code":     acctest.Representation{RepType: acctest.Optional, Create: `405`, Update: `400`},
	}
)

// issue-routing-tag: load_balancer/default
func TestResourceLoadBalancerRuleSetResource_controlAccess_test(t *testing.T) {
	httpreplay.SetScenario("TestResourceLoadBalancerRuleSetResource_controlAccess_test")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_control_access_rule_set"

	var resId, resId2 string
	datasourceName := "data.oci_load_balancer_rule_sets.test_control_access_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_control_access_rule_set"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.PreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckLoadBalancerRuleSetDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Required, acctest.Create, ruleSetControlAccessRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "405",
						"allowed_methods.#": "1",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Optional, acctest.Update, ruleSetControlAccessRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "2",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
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
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Optional, acctest.Update,
						acctest.RepresentationCopyWithNewProperties(ruleSetControlAccessRepresentation, map[string]interface{}{
							"items": acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleSetItemsAnotherControlAccessRepresentation},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "3",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_control_access_rule_set"),

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
					acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_control_access_rule_sets", acctest.Optional, acctest.Update, ruleSetControlAccessDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Optional, acctest.Update, ruleSetControlAccessRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
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
					acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Required, acctest.Create, ruleSetControlAccessSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_control_access_rule_set", acctest.Optional, acctest.Update, ruleSetControlAccessRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_control_access_rule_set"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action":            "CONTROL_ACCESS_USING_HTTP_METHODS",
						"status_code":       "400",
						"allowed_methods.#": "2",
					},
						[]string{}),
				),
			},
			// verify resource import
			{
				Config:            config + RuleSetControlAccessRequiredOnlyRepresentation,
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
