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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	LoadBalancerRoutingPolicyResourceConfig = LoadBalancerRoutingPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Optional, acctest.Update, loadBalancerRoutingPolicyRepresentation)

	loadBalancerRoutingPolicySingularDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id":    acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"routing_policy_name": acctest.Representation{RepType: acctest.Required, Create: `example_routing_rules`},
	}

	loadBalancerRoutingPolicyDataSourceRepresentation = map[string]interface{}{
		"load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: loadBalancerRoutingPolicyDataSourceFilterRepresentation}}
	loadBalancerRoutingPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `name`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy.name}`}},
	}

	loadBalancerRoutingPolicyRepresentation = map[string]interface{}{
		"condition_language_version": acctest.Representation{RepType: acctest.Required, Create: `V1`},
		"load_balancer_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `example_routing_rules`},
		"rules":                      acctest.RepresentationGroup{RepType: acctest.Required, Group: loadBalancerRoutingPolicyRulesRepresentation},
	}
	loadBalancerRoutingPolicyRulesRepresentation = map[string]interface{}{
		"actions":   acctest.RepresentationGroup{RepType: acctest.Required, Group: loadBalancerRoutingPolicyRulesActionsRepresentation},
		"condition": acctest.Representation{RepType: acctest.Required, Create: `all(http.request.url.path eq (i ''))`},
		"name":      acctest.Representation{RepType: acctest.Required, Create: `example_routing_rules`, Update: `name2`},
	}
	loadBalancerRoutingPolicyRulesActionsRepresentation = map[string]interface{}{
		"name":             acctest.Representation{RepType: acctest.Required, Create: `FORWARD_TO_BACKENDSET`, Update: `FORWARD_TO_BACKENDSET`},
		"backend_set_name": acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_backend_set.test_backend_set.name}`},
	}

	LoadBalancerRoutingPolicyResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_backend_set", "test_backend_set", acctest.Required, acctest.Create, backendSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies
)

// issue-routing-tag: load_balancer/default
func TestLoadBalancerLoadBalancerRoutingPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestLoadBalancerLoadBalancerRoutingPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy"
	datasourceName := "data.oci_load_balancer_load_balancer_routing_policies.test_load_balancer_routing_policies"
	singularDatasourceName := "data.oci_load_balancer_load_balancer_routing_policy.test_load_balancer_routing_policy"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckLoadBalancerLoadBalancerRoutingPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + LoadBalancerRoutingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Required, acctest.Create, loadBalancerRoutingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition_language_version", "V1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_routing_rules"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.actions.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.actions.0.name", "FORWARD_TO_BACKENDSET"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.condition", "all(http.request.url.path eq (i ''))"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.name", "example_routing_rules"),

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
			Config: config + compartmentIdVariableStr + LoadBalancerRoutingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Optional, acctest.Update, loadBalancerRoutingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "condition_language_version", "V1"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_routing_rules"),
				resource.TestCheckResourceAttr(resourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.actions.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "rules.0.actions.0.backend_set_name"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.actions.0.name", "FORWARD_TO_BACKENDSET"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.condition", "all(http.request.url.path eq (i ''))"),
				resource.TestCheckResourceAttr(resourceName, "rules.0.name", "name2"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policies", "test_load_balancer_routing_policies", acctest.Optional, acctest.Update, loadBalancerRoutingPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerRoutingPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Optional, acctest.Update, loadBalancerRoutingPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "load_balancer_id"),

				resource.TestCheckResourceAttr(datasourceName, "routing_policies.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.condition_language_version", "V1"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.name", "example_routing_rules"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.rules.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.rules.0.actions.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "routing_policies.0.rules.0.actions.0.backend_set_name"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.rules.0.actions.0.name", "FORWARD_TO_BACKENDSET"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.rules.0.condition", "all(http.request.url.path eq (i ''))"),
				resource.TestCheckResourceAttr(datasourceName, "routing_policies.0.rules.0.name", "name2"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_load_balancer_load_balancer_routing_policy", "test_load_balancer_routing_policy", acctest.Required, acctest.Create, loadBalancerRoutingPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + LoadBalancerRoutingPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "routing_policy_name"),

				resource.TestCheckResourceAttr(singularDatasourceName, "condition_language_version", "V1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "example_routing_rules"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.actions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.actions.0.name", "FORWARD_TO_BACKENDSET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.condition", "all(http.request.url.path eq (i ''))"),
				resource.TestCheckResourceAttr(singularDatasourceName, "rules.0.name", "name2"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + LoadBalancerRoutingPolicyResourceConfig,
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
	})
}

func testAccCheckLoadBalancerLoadBalancerRoutingPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).LoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_load_balancer_load_balancer_routing_policy" {
			noResourceFound = false
			request := oci_load_balancer.GetRoutingPolicyRequest{}

			if value, ok := rs.Primary.Attributes["load_balancer_id"]; ok {
				request.LoadBalancerId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.RoutingPolicyName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")

			_, err := client.GetRoutingPolicy(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("LoadBalancerLoadBalancerRoutingPolicy") {
		resource.AddTestSweepers("LoadBalancerLoadBalancerRoutingPolicy", &resource.Sweeper{
			Name:         "LoadBalancerLoadBalancerRoutingPolicy",
			Dependencies: acctest.DependencyGraph["loadBalancerRoutingPolicy"],
			F:            sweepLoadBalancerLoadBalancerRoutingPolicyResource,
		})
	}
}

func sweepLoadBalancerLoadBalancerRoutingPolicyResource(compartment string) error {
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()
	loadBalancerRoutingPolicyIds, err := getLoadBalancerRoutingPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, loadBalancerRoutingPolicyId := range loadBalancerRoutingPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[loadBalancerRoutingPolicyId]; !ok {
			deleteRoutingPolicyRequest := oci_load_balancer.DeleteRoutingPolicyRequest{}

			deleteRoutingPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "load_balancer")
			_, error := loadBalancerClient.DeleteRoutingPolicy(context.Background(), deleteRoutingPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting LoadBalancerRoutingPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", loadBalancerRoutingPolicyId, error)
				continue
			}
		}
	}
	return nil
}

func getLoadBalancerRoutingPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "LoadBalancerRoutingPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	loadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).LoadBalancerClient()

	listRoutingPoliciesRequest := oci_load_balancer.ListRoutingPoliciesRequest{}

	loadBalancerIds, error := getLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting loadBalancerId required for LoadBalancerRoutingPolicy resource requests \n")
	}
	for _, loadBalancerId := range loadBalancerIds {
		listRoutingPoliciesRequest.LoadBalancerId = &loadBalancerId

		listRoutingPoliciesResponse, err := loadBalancerClient.ListRoutingPolicies(context.Background(), listRoutingPoliciesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting LoadBalancerRoutingPolicy list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, loadBalancerRoutingPolicy := range listRoutingPoliciesResponse.Items {
			id := *loadBalancerRoutingPolicy.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "LoadBalancerRoutingPolicyId", id)
		}

	}
	return resourceIds, nil
}
