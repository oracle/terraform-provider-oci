// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v45/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v45/networkloadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	NetworkLoadBalancerRequiredOnlyResource = NetworkLoadBalancerResourceDependencies +
		generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerRepresentation)

	NetworkLoadBalancerResourceConfig = NetworkLoadBalancerResourceDependencies +
		generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Update, networkLoadBalancerRepresentation)

	networkLoadBalancerSingularDataSourceRepresentation = map[string]interface{}{
		"network_load_balancer_id": Representation{repType: Required, create: `${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`},
	}

	networkLoadBalancerDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"state":          Representation{repType: Optional, create: `Active`},
		"filter":         RepresentationGroup{Required, networkLoadBalancerDataSourceFilterRepresentation}}
	networkLoadBalancerDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_network_load_balancer_network_load_balancer.test_network_load_balancer.id}`}},
	}

	networkLoadBalancerRepresentation = map[string]interface{}{
		"compartment_id":                 Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":                   Representation{repType: Required, create: `displayName`, update: `displayName2`},
		"subnet_id":                      Representation{repType: Required, create: `${oci_core_subnet.test_subnet.id}`},
		"defined_tags":                   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"is_preserve_source_destination": Representation{repType: Optional, create: `false`, update: `true`},
		"is_private":                     Representation{repType: Optional, create: `false`},
		"network_security_group_ids":     Representation{repType: Optional, create: []string{`${oci_core_network_security_group.test_network_security_group.id}`}},
		"reserved_ips":                   RepresentationGroup{Optional, networkLoadBalancerReservedIpsRepresentation},
	}
	networkLoadBalancerReservedIpsRepresentation = map[string]interface{}{
		"id": Representation{repType: Optional, create: `${oci_core_public_ip.test_public_ip.id}`},
	}

	NetworkLoadBalancerReservedIpDependencies = generateResourceFromRepresentationMap("oci_core_public_ip", "test_public_ip", Required, Create, publicIpRepresentation)

	NetworkLoadBalancerResourceDependencies = generateResourceFromRepresentationMap("oci_core_network_security_group", "test_network_security_group", Required, Create, networkSecurityGroupRepresentation) +
		generateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", Required, Create, subnetRepresentation) +
		generateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", Required, Create, vcnRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: network_load_balancer/default
func TestNetworkLoadBalancerNetworkLoadBalancerResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkLoadBalancerNetworkLoadBalancerResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_network_load_balancer_network_load_balancer.test_network_load_balancer"
	datasourceName := "data.oci_network_load_balancer_network_load_balancers.test_network_load_balancers"
	singularDatasourceName := "data.oci_network_load_balancer_network_load_balancer.test_network_load_balancer"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckNetworkLoadBalancerNetworkLoadBalancerDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Create, networkLoadBalancerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.reserved_ip.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Create,
						representationCopyWithNewProperties(networkLoadBalancerRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "false"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.ip_address"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_addresses.0.reserved_ip.0.id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Update, networkLoadBalancerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "1"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},

			// verify removal of NSGs
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Update,
						representationCopyWithNewProperties(networkLoadBalancerRepresentation, map[string]interface{}{
							"network_security_group_ids": Representation{repType: Required, create: []string{}},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.0.is_public", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_preserve_source_destination", "true"),
					resource.TestCheckResourceAttr(resourceName, "is_private", "false"),
					resource.TestCheckResourceAttr(resourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttr(resourceName, "network_security_group_ids.#", "0"),

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
					generateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancers", "test_network_load_balancers", Optional, Update, networkLoadBalancerDataSourceRepresentation) +
					compartmentIdVariableStr + NetworkLoadBalancerResourceDependencies + NetworkLoadBalancerReservedIpDependencies +
					generateResourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Optional, Update, networkLoadBalancerRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(datasourceName, "state", "Active"),

					resource.TestCheckResourceAttr(datasourceName, "network_load_balancer_collection.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "network_load_balancer_collection.0.items.#", "1"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_network_load_balancer_network_load_balancer", "test_network_load_balancer", Required, Create, networkLoadBalancerSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NetworkLoadBalancerResourceConfig + NetworkLoadBalancerReservedIpDependencies,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "network_load_balancer_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "ip_addresses.0.is_public", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_preserve_source_destination", "true"),
					resource.TestCheckResourceAttr(singularDatasourceName, "is_private", "false"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
					resource.TestCheckResourceAttr(singularDatasourceName, "network_security_group_ids.#", "1"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NetworkLoadBalancerResourceConfig + NetworkLoadBalancerReservedIpDependencies,
			},
			// verify resource import
			{
				Config:            config,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"reserved_ips",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func testAccCheckNetworkLoadBalancerNetworkLoadBalancerDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).networkLoadBalancerClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_load_balancer_network_load_balancer" {
			noResourceFound = false
			request := oci_network_load_balancer.GetNetworkLoadBalancerRequest{}

			tmp := rs.Primary.ID
			request.NetworkLoadBalancerId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "network_load_balancer")

			response, err := client.GetNetworkLoadBalancer(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_network_load_balancer.LifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("NetworkLoadBalancerNetworkLoadBalancer") {
		resource.AddTestSweepers("NetworkLoadBalancerNetworkLoadBalancer", &resource.Sweeper{
			Name:         "NetworkLoadBalancerNetworkLoadBalancer",
			Dependencies: DependencyGraph["networkLoadBalancer"],
			F:            sweepNetworkLoadBalancerNetworkLoadBalancerResource,
		})
	}
}

func sweepNetworkLoadBalancerNetworkLoadBalancerResource(compartment string) error {
	networkLoadBalancerClient := GetTestClients(&schema.ResourceData{}).networkLoadBalancerClient()
	networkLoadBalancerIds, err := getNetworkLoadBalancerIds(compartment)
	if err != nil {
		return err
	}
	for _, networkLoadBalancerId := range networkLoadBalancerIds {
		if ok := SweeperDefaultResourceId[networkLoadBalancerId]; !ok {
			deleteNetworkLoadBalancerRequest := oci_network_load_balancer.DeleteNetworkLoadBalancerRequest{}

			deleteNetworkLoadBalancerRequest.NetworkLoadBalancerId = &networkLoadBalancerId

			deleteNetworkLoadBalancerRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "network_load_balancer")
			_, error := networkLoadBalancerClient.DeleteNetworkLoadBalancer(context.Background(), deleteNetworkLoadBalancerRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkLoadBalancer %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkLoadBalancerId, error)
				continue
			}
			waitTillCondition(testAccProvider, &networkLoadBalancerId, networkLoadBalancerSweepWaitCondition, time.Duration(3*time.Minute),
				networkLoadBalancerSweepResponseFetchOperation, "network_load_balancer", true)
		}
	}
	return nil
}

func getNetworkLoadBalancerIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "NetworkLoadBalancerId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkLoadBalancerClient := GetTestClients(&schema.ResourceData{}).networkLoadBalancerClient()

	listNetworkLoadBalancersRequest := oci_network_load_balancer.ListNetworkLoadBalancersRequest{}
	listNetworkLoadBalancersRequest.CompartmentId = &compartmentId
	listNetworkLoadBalancersRequest.LifecycleState = oci_network_load_balancer.ListNetworkLoadBalancersLifecycleStateActive
	listNetworkLoadBalancersResponse, err := networkLoadBalancerClient.ListNetworkLoadBalancers(context.Background(), listNetworkLoadBalancersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkLoadBalancer list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkLoadBalancer := range listNetworkLoadBalancersResponse.Items {
		id := *networkLoadBalancer.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "NetworkLoadBalancerId", id)
	}
	return resourceIds, nil
}

func networkLoadBalancerSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkLoadBalancerResponse, ok := response.Response.(oci_network_load_balancer.GetNetworkLoadBalancerResponse); ok {
		return networkLoadBalancerResponse.LifecycleState != oci_network_load_balancer.LifecycleStateDeleted
	}
	return false
}

func networkLoadBalancerSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.networkLoadBalancerClient().GetNetworkLoadBalancer(context.Background(), oci_network_load_balancer.GetNetworkLoadBalancerRequest{
		NetworkLoadBalancerId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
