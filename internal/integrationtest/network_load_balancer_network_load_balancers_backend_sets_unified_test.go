// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRequiredOnlyResource = NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_backend_sets_unified", "test_network_load_balancers_backend_sets_unified", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation)

	NetworkLoadBalancersBackendSetsUnifiedResourceConfig = NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_backend_sets_unified", "test_network_load_balancers_backend_sets_unified", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation)

	NetworkLoadBalancersBackendSetsUnifiedSingularDataSourceRepresentation = map[string]interface{}{
		"backend_set_name":         acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancers_backend_sets_unified.test_network_load_balancers_backend_sets_unified.name}`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation = map[string]interface{}{
		"health_checker":           acctest.RepresentationGroup{RepType: acctest.Required, Group: NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedHealthCheckerRepresentation},
		"name":                     acctest.Representation{RepType: acctest.Required, Create: `example_backend_set`},
		"network_load_balancer_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
		"policy":                   acctest.Representation{RepType: acctest.Required, Create: `FIVE_TUPLE`, Update: `THREE_TUPLE`},
		"backends":                 acctest.RepresentationGroup{RepType: acctest.Optional, Group: NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedBackendsRepresentation},
		"ip_version":               acctest.Representation{RepType: acctest.Optional, Create: `IPV4`},
		"is_preserve_source":       acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedHealthCheckerRepresentation = map[string]interface{}{
		"protocol":           acctest.Representation{RepType: acctest.Required, Create: `TCP`, Update: `TCP`},
		"interval_in_millis": acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
		"port":               acctest.Representation{RepType: acctest.Optional, Create: `80`, Update: `8080`},
		"request_data":       acctest.Representation{RepType: acctest.Optional, Create: `SGVsbG9Xb3JsZA==`, Update: `QnllV29ybGQ=`},
		"response_data":      acctest.Representation{RepType: acctest.Optional, Create: `SGVsbG9Xb3JsZA==`, Update: `QnllV29ybGQ=`},
		"retries":            acctest.Representation{RepType: acctest.Optional, Create: `3`, Update: `5`},
		"timeout_in_millis":  acctest.Representation{RepType: acctest.Optional, Create: `10000`, Update: `30000`},
	}

	NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedBackendsRepresentation = map[string]interface{}{
		"port":       acctest.Representation{RepType: acctest.Required, Create: `10`},
		"ip_address": acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.3`},
		"is_backup":  acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_drain":   acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"is_offline": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"name":       acctest.Representation{RepType: acctest.Optional, Create: `example_backend`},
		"weight":     acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
	}

	NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancerRepresentation)
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_load_balancer_network_load_balancers_backend_sets_unified.test_network_load_balancers_backend_sets_unified"

	var resId, resId2 string

	acctest.ResourceTest(t, testAccCheckNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_backend_sets_unified", "test_network_load_balancers_backend_sets_unified", acctest.Required, acctest.Create, NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_backend_sets_unified", "test_network_load_balancers_backend_sets_unified", acctest.Optional, acctest.Create, NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "backends", map[string]string{
					"ip_address": "10.0.0.3",
					"is_backup":  "false",
					"is_drain":   "false",
					"is_offline": "false",
					"name":       "example_backend",
					"port":       "10",
					"weight":     "10",
				}, []string{}),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "80"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.request_data", "SGVsbG9Xb3JsZA=="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_body_regex", ""),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_data", "SGVsbG9Xb3JsZA=="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "3"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "10000"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "false"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "FIVE_TUPLE"),

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
			Config: config + compartmentIdVariableStr + NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers_backend_sets_unified", "test_network_load_balancers_backend_sets_unified", acctest.Optional, acctest.Update, NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backends.#", "1"),
				acctest.CheckResourceSetContainsElementWithProperties(resourceName, "backends", map[string]string{
					"ip_address": "10.0.0.3",
					"is_backup":  "true",
					"is_drain":   "true",
					"is_offline": "true",
					"name":       "example_backend",
					"port":       "10",
					"weight":     "11",
				},
					[]string{}),
				resource.TestCheckResourceAttr(resourceName, "health_checker.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.interval_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.port", "8080"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.protocol", "TCP"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.request_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.response_data", "QnllV29ybGQ="),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.retries", "5"),
				resource.TestCheckResourceAttr(resourceName, "health_checker.0.timeout_in_millis", "30000"),
				resource.TestCheckResourceAttr(resourceName, "ip_version", "IPV4"),
				resource.TestCheckResourceAttr(resourceName, "is_preserve_source", "true"),
				resource.TestCheckResourceAttr(resourceName, "name", "example_backend_set"),
				resource.TestCheckResourceAttrSet(resourceName, "network_load_balancer_id"),
				resource.TestCheckResourceAttr(resourceName, "policy", "THREE_TUPLE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkLoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_load_balancer_network_load_balancers_backend_sets_unified" {
			noResourceFound = false
			request := oci_network_load_balancer.GetBackendSetRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.BackendSetName = &value
			}

			if value, ok := rs.Primary.Attributes["network_load_balancer_id"]; ok {
				request.NetworkLoadBalancerId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")

			_, err := client.GetBackendSet(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified") {
		resource.AddTestSweepers("NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified", &resource.Sweeper{
			Name:         "NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified",
			Dependencies: acctest.DependencyGraph["networkLoadBalancersBackendSetsUnified"],
			F:            sweepNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource,
		})
	}
}

func sweepNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource(compartment string) error {
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()
	networkLoadBalancersBackendSetsUnifiedIds, err := getNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedIds(compartment)
	if err != nil {
		return err
	}
	for _, networkLoadBalancersBackendSetsUnifiedId := range networkLoadBalancersBackendSetsUnifiedIds {
		if ok := acctest.SweeperDefaultResourceId[networkLoadBalancersBackendSetsUnifiedId]; !ok {
			deleteBackendSetRequest := oci_network_load_balancer.DeleteBackendSetRequest{}

			deleteBackendSetRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_load_balancer")
			_, error := networkLoadBalancerClient.DeleteBackendSet(context.Background(), deleteBackendSetRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkLoadBalancersBackendSetsUnified %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkLoadBalancersBackendSetsUnifiedId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkLoadBalancersBackendSetsUnifiedId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkLoadBalancerClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkLoadBalancerClient()

	listBackendSetsRequest := oci_network_load_balancer.ListBackendSetsRequest{}

	networkLoadBalancerIds, error := getNetworkLoadBalancerNetworkLoadBalancerIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkLoadBalancerId required for NetworkLoadBalancersBackendSetsUnified resource requests \n")
	}
	for _, networkLoadBalancerId := range networkLoadBalancerIds {
		listBackendSetsRequest.NetworkLoadBalancerId = &networkLoadBalancerId

		listBackendSetsResponse, err := networkLoadBalancerClient.ListBackendSets(context.Background(), listBackendSetsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkLoadBalancersBackendSetsUnified list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkLoadBalancersBackendSetsUnified := range listBackendSetsResponse.Items {
			id := *networkLoadBalancersBackendSetsUnified.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkLoadBalancersBackendSetsUnifiedId", id)
		}

	}
	return resourceIds, nil
}
