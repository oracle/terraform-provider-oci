// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	RuleSetResourceConfig = RuleSetResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Update, ruleSetRepresentation)

	ruleSetSingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_rule_set`},
	}

	ruleSetDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleSetDataSourceFilterRepresentation}}
	ruleSetDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_rule_set.test_rule_set.name}`}},
	}

	ruleSetRepresentation = map[string]interface{}{
		"items":            acctest.RepresentationGroup{RepType: acctest.Required, Group: ruleSetItemsRepresentation},
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":             acctest.Representation{RepType: acctest.Required, Create: `example_rule_set`},
	}
	ruleSetItemsRepresentation = map[string]interface{}{
		"action": acctest.Representation{RepType: acctest.Required, Create: `ADD_HTTP_REQUEST_HEADER`, Update: `EXTEND_HTTP_REQUEST_HEADER_VALUE`},
		"header": acctest.Representation{RepType: acctest.Required, Create: `example_header_name`, Update: `example_header_name2`},
		"prefix": acctest.Representation{RepType: acctest.Optional, Create: `prefix`},
		"suffix": acctest.Representation{RepType: acctest.Optional, Create: `suffix`},
		"value":  acctest.Representation{RepType: acctest.Required, Create: `example_header_value`, Update: ``},
	}

	RuleSetResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
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
	items {
		#Required
		action = "CONTROL_ACCESS_USING_HTTP_METHODS"
		allowed_methods = ["POST", "GET"]
		status_code = "400"
	}
	items {
		#Required
		action = "ALLOW"
		conditions {
			#Required
			attribute_name = "SOURCE_IP_ADDRESS"
			attribute_value = "129.0.0.0/8"
		}
		description = "description"
	}
	items {
		#Required
		action = "ALLOW"
		conditions {
			#Required
			attribute_name = "SOURCE_VCN_ID"
			attribute_value = "${oci_core_vcn.test_lb_vcn.id}"
		}
		conditions {
			#Required
			attribute_name = "SOURCE_VCN_IP_ADDRESS"
			attribute_value = "10.10.1.0/24"
		}
	}
	items {
		#Required
		action = "REDIRECT"
		conditions {
			#Required
			attribute_name = "PATH"
			attribute_value = "/example"
			operator = "SUFFIX_MATCH"
		}
		redirect_uri {
			protocol = "{protocol}"
			host = "in{host}"
			port = 8081
			path = "{path}/video"
			query = "?lang=en"
		}
		response_code = 302
	}
	items {
		#Required
		action = "HTTP_HEADER"
		are_invalid_characters_allowed = true
		http_large_header_size_in_kb = 8
	}
	load_balancer_id = "${oci_load_balancer_load_balancer.test_load_balancer.id}"
	name = "example_rule_set"
}
`
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerRuleSetResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerRuleSetResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_rule_set.test_rule_set"
	datasourceName := "data.oci_load_balancer_rule_sets.test_rule_sets"
	singularDatasourceName := "data.oci_load_balancer_rule_set.test_rule_set"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+RuleSetResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetRepresentation), "loadbalancer", "ruleSet", t)

	acctest.ResourceTest(t, testAccCheckLoadBalancerRuleSetDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "ADD_HTTP_REQUEST_HEADER",
					"header": "example_header_name",
					"value":  "example_header_value",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + RuleSetResourceDependencies +
				RuleSetResourceWithMultipleRules,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "items.#", "11"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "ADD_HTTP_REQUEST_HEADER",
					"header": "example_header_name",
					"value":  "example_header_value",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "ADD_HTTP_RESPONSE_HEADER",
					"header": "example_header_name2",
					"value":  "example_header_value2",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "EXTEND_HTTP_REQUEST_HEADER_VALUE",
					"header": "example_header_name3",
					"prefix": "prefix",
					"suffix": "suffix",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "EXTEND_HTTP_RESPONSE_HEADER_VALUE",
					"header": "example_header_name4",
					"prefix": "prefix",
					"suffix": "suffix",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "REMOVE_HTTP_REQUEST_HEADER",
					"header": "example_header_name5",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action": "REMOVE_HTTP_RESPONSE_HEADER",
					"header": "example_header_name6",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action":      "CONTROL_ACCESS_USING_HTTP_METHODS",
					"status_code": "400",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action":       "ALLOW",
					"conditions.#": "1",
					"description":  "description",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action":       "ALLOW",
					"conditions.#": "2",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action":         "REDIRECT",
					"conditions.#":   "1",
					"redirect_uri.#": "1",
					"response_code":  "302",
				},
					[]string{}),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "items", map[string]string{
					"action":                         "HTTP_HEADER",
					"are_invalid_characters_allowed": "true",
					"http_large_header_size_in_kb":   "8",
				},
					[]string{}),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_rule_set"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_sets", "test_rule_sets", acctest.Optional, acctest.Update, ruleSetDataSourceRepresentation) +
				compartmentIdVariableStr + RuleSetResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Optional, acctest.Update, ruleSetRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "rule_sets.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(datasourceName, "rule_sets.0.items", map[string]string{
					"action": "EXTEND_HTTP_REQUEST_HEADER_VALUE",
					"header": "example_header_name2",
					"prefix": "prefix",
					"suffix": "suffix",
					"value":  "",
				},
					[]string{}),
				resource.TestCheckResourceAttr(datasourceName, "rule_sets.0.name", "example_rule_set"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_rule_set", "test_rule_set", acctest.Required, acctest.Create, ruleSetSingularDataSourceRepresentation) +
				compartmentIdVariableStr + RuleSetResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{},
					[]string{}),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_rule_set"),

				resource.TestCheckResourceAttr(singularDatasourceName, "items.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(singularDatasourceName, "items", map[string]string{
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
			Config:            config + compartmentIdVariableStr + RuleSetResourceConfig,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"state",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckLoadBalancerRuleSetDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
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

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

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

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("LoadBalancerRuleSet") {
		resource.AddTestSweepers("LoadBalancerRuleSet", &resource.Sweeper{
			Name:         "LoadBalancerRuleSet",
			Dependencies: acctest.DependencyGraph["ruleSet"],
			F:            sweepLoadBalancerRuleSetResource,
		})
	}
}

func sweepLoadBalancerRuleSetResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	ruleSetIds, err := getRuleSetIds(compartment)
	if err != nil {
		return err
	}
	for _, ruleSetId := range ruleSetIds {
		if ok := acctest.SweeperDefaultResourceId[ruleSetId]; !ok {
			deleteRuleSetRequest := oci_load_balancer.DeleteRuleSetRequest{}

			deleteRuleSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteRuleSet(context.Background(), deleteRuleSetRequest)
			if error != nil {
				fmt.Printf("Error deleting RuleSet %s %s, It is possible that the resource is already deleted. Please verify manually \n", ruleSetId, error)
				continue
			}
		}
	}
	return nil
}

func getRuleSetIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "RuleSetId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

	listRuleSetsRequest := oci_load_balancer.ListRuleSetsRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for RuleSet resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listRuleSetsRequest.LoadBalancerId = &loadBalancerId

		listRuleSetsResponse, err := loadBalancerClient.ListRuleSets(context.Background(), listRuleSetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting RuleSet list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, ruleSet := range listRuleSetsResponse.Items {
			id := *ruleSet.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "RuleSetId", id)
		}

	}
	return resourceIds, nil
}
