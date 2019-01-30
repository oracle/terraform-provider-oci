// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

var (
	RuleSetResourceConfig = RuleSetResourceDependencies +
		generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update, ruleSetRepresentation)

	ruleSetSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_rule_set`},
	}

	ruleSetDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           RepresentationGroup{Required, ruleSetDataSourceFilterRepresentation}}
	ruleSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `name`},
		"values": Representation{repType: Required, create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
	}

	ruleSetRepresentation = map[string]interface{}{
		"items":            RepresentationGroup{Required, ruleSetItemsRepresentation},
		"load_balancer_id": Representation{repType: Required, create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             Representation{repType: Required, create: `example_rule_set`},
	}
	ruleSetItemsRepresentation = map[string]interface{}{
		"action": Representation{repType: Required, create: `ADD_HTTP_REQUEST_HEADER`, update: `EXTEND_HTTP_REQUEST_HEADER_VALUE`},
		"header": Representation{repType: Required, create: `example_header_name`, update: `example_header_name2`},
		"prefix": Representation{repType: Optional, create: `prefix`},
		"suffix": Representation{repType: Optional, create: `suffix`},
		"value":  Representation{repType: Required, create: `example_header_value`, update: ``},
	}

	RuleSetResourceDependencies      = LoadBalancerResourceConfig
	RuleSetResourceWithMultipleRules = `
resource "oci_load_balancer_rule_set" "test_rule_set" {
	#Required
	items {
		#Required
		action = "ADD_HTTP_REQUEST_HEADER"
		header = "example_header_name"
		value = "example_header_value"
	}
	items {
		#Required
		action = "ADD_HTTP_RESPONSE_HEADER"
		header = "example_header_name2"
		value = "example_header_value2"
	}
	items {
		#Required
		action = "EXTEND_HTTP_REQUEST_HEADER_VALUE"
		header = "example_header_name3"
		prefix = "prefix"
		suffix = "suffix"
	}
	items {
		#Required
		action = "EXTEND_HTTP_RESPONSE_HEADER_VALUE"
		header = "example_header_name4"
		prefix = "prefix"
		suffix = "suffix"
	}
	items {
		#Required
		action = "REMOVE_HTTP_REQUEST_HEADER"
		header = "example_header_name5"
	}
	items {
		#Required
		action = "REMOVE_HTTP_RESPONSE_HEADER"
		header = "example_header_name6"
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "example_rule_set"
}
`
)

func TestLoadBalancerRuleSetResource_basic(t *testing.T) {
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
			// verify create
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "ADD_HTTP_REQUEST_HEADER",
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

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
					RuleSetResourceWithMultipleRules,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "items.#", "6"),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "ADD_HTTP_REQUEST_HEADER",
						"header": "example_header_name",
						"value":  "example_header_value",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "ADD_HTTP_RESPONSE_HEADER",
						"header": "example_header_name2",
						"value":  "example_header_value2",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "EXTEND_HTTP_REQUEST_HEADER_VALUE",
						"header": "example_header_name3",
						"prefix": "prefix",
						"suffix": "suffix",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "EXTEND_HTTP_RESPONSE_HEADER_VALUE",
						"header": "example_header_name4",
						"prefix": "prefix",
						"suffix": "suffix",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "REMOVE_HTTP_REQUEST_HEADER",
						"header": "example_header_name5",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
						"action": "REMOVE_HTTP_RESPONSE_HEADER",
						"header": "example_header_name6",
					},
						[]string{}),
					resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

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
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_rule_sets", Optional, Update, ruleSetDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceDependencies +
					generateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Optional, Update, ruleSetRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

					resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "6"),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "ADD_HTTP_REQUEST_HEADER",
						"header": "example_header_name",
						"value":  "example_header_value",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "ADD_HTTP_RESPONSE_HEADER",
						"header": "example_header_name2",
						"value":  "example_header_value2",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "EXTEND_HTTP_REQUEST_HEADER_VALUE",
						"header": "example_header_name3",
						"prefix": "prefix",
						"suffix": "suffix",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "EXTEND_HTTP_RESPONSE_HEADER_VALUE",
						"header": "example_header_name4",
						"prefix": "prefix",
						"suffix": "suffix",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "REMOVE_HTTP_REQUEST_HEADER",
						"header": "example_header_name5",
					},
						[]string{}),
					CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
						"action": "REMOVE_HTTP_RESPONSE_HEADER",
						"header": "example_header_name6",
					},
						[]string{}),
					resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.name", "example_rule_set"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", Required, Create, ruleSetSingularDataSourceRepresentation) +
					compartmentIdVariableStr + RuleSetResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{},
						[]string{}),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_rule_set"),

					resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
					CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
						"action": "EXTEND_HTTP_REQUEST_HEADER_VALUE",
						"header": "example_header_name2",
						"prefix": "prefix",
						"suffix": "suffix",
						"value":  "",
					},
						[]string{}),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_rule_set"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + RuleSetResourceConfig,
			},
			// verify resource import
			{
				Config:            config,
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

func testAccCheckLoadBalancerRuleSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).loadBalancerClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_rule_set" {
			noResourceFound = false
			request := oci_load_balancer.GetRuleSetRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.RuleSetName = &value
			}

			_, err := client.GetRuleSet(context.Background(), request)

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
