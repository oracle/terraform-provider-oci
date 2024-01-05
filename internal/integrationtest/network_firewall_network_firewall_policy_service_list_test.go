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
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	createServiceListResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_service_list",
		"test_network_firewall_policy_service_list",
		acctest.Required, acctest.Create,
		serviceListRepresentation,
	)
	createServiceListResourceConfig = serviceListResourceDependencies + createServiceListResource

	serviceListResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_service_list",
		"test_network_firewall_policy_service_list",
		acctest.Optional, acctest.Update,
		serviceListRepresentation,
	)

	serviceListResourceConfig = serviceListResourceDependencies + serviceListResource

	serviceListSingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list.name}`},
	}

	serviceListDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	serviceListRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `service_list_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"services":                   acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service.name}`}, Update: []string{}},
	}

	serviceListResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	) + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_service",
		"test_network_firewall_policy_service",
		acctest.Required, acctest.Create,
		serviceRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyServiceListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyServiceListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_service_lists.test_network_firewall_policy_service_lists"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_service_list.test_network_firewall_policy_service_list"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+serviceListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy_service_list", "test_network_firewall_policy_service_list", acctest.Required, acctest.Create, serviceListRepresentation), "networkfirewall", "networkFirewallPolicyServiceList", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyServiceListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createServiceListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "service_list_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + serviceListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "service_list_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "services.#", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "total_services"),

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
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_service_lists",
					"test_network_firewall_policy_service_lists",
					acctest.Optional, acctest.Update,
					serviceListDataSourceRepresentation,
				) +
				compartmentIdVariableStr + serviceListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "service_list_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_list_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_service_list",
					"test_network_firewall_policy_service_list",
					acctest.Required, acctest.Create,
					serviceListSingularDataSourceRepresentation,
				) +
				compartmentIdVariableStr + serviceListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "name", "service_list_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "services.#", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_services"),
			),
		},
		// verify resource import
		{
			Config:                  config + createServiceListResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyServiceListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_service_list" {
			noResourceFound = false
			request := oci_network_firewall.GetServiceListRequest{}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.ServiceListName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetServiceList(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyServiceList") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyServiceList", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyServiceList",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyServiceList"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyServiceListResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyServiceListResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyServiceListIds, err := getNetworkFirewallNetworkFirewallPolicyServiceListIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyServiceListId := range networkFirewallPolicyServiceListIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyServiceListId]; !ok {
			deleteServiceListRequest := oci_network_firewall.DeleteServiceListRequest{}

			deleteServiceListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteServiceList(context.Background(), deleteServiceListRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyServiceList %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyServiceListId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyServiceListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyServiceListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listServiceListsRequest := oci_network_firewall.ListServiceListsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyServiceList resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listServiceListsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listServiceListsResponse, err := networkFirewallClient.ListServiceLists(context.Background(), listServiceListsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyServiceList list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyServiceList := range listServiceListsResponse.Items {
			id := *networkFirewallPolicyServiceList.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyServiceListId", id)
		}

	}
	return resourceIds, nil
}
