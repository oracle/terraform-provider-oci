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
	createURLListResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_url_list",
		"test_network_firewall_policy_url_list",
		acctest.Required, acctest.Create,
		urlListRepresentation,
	)

	createUrlListResourceConfig = urlListResourceDependencies + createURLListResource

	updatedUrlListResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_url_list",
		"test_network_firewall_policy_url_list",
		acctest.Optional, acctest.Update,
		urlListRepresentation2,
	)

	updatedUrlListResourceConfig = urlListResourceDependencies + updatedUrlListResource

	urlListSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	urlListDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	urlListRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `url_list_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"urls":                       []acctest.RepresentationGroup{{RepType: acctest.Required, Group: urlsRepresentation1}},
	}

	urlListRepresentation2 = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `url_list_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"urls":                       []acctest.RepresentationGroup{{RepType: acctest.Required, Group: urlsRepresentation1}, {RepType: acctest.Required, Group: urlsRepresentation2}},
	}
	urlsRepresentation1 = map[string]interface{}{
		"pattern": acctest.Representation{RepType: acctest.Required, Create: `www.url1.com`},
		"type":    acctest.Representation{RepType: acctest.Required, Create: `SIMPLE`},
	}

	urlsRepresentation2 = map[string]interface{}{
		"pattern": acctest.Representation{RepType: acctest.Required, Create: `www.url2.com`},
		"type":    acctest.Representation{RepType: acctest.Required, Update: `SIMPLE`},
	}

	urlListResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyUrlListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyUrlListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_url_lists.test_network_firewall_policy_url_lists"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_url_list.test_network_firewall_policy_url_list"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+createUrlListResourceConfig, "networkfirewall", "networkFirewallPolicyUrlList", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyUrlListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createUrlListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "url_list_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "urls.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "urls.0.pattern", "www.url1.com"),
				resource.TestCheckResourceAttr(resourceName, "urls.0.type", "SIMPLE"),

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
			Config: config + compartmentIdVariableStr + updatedUrlListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "url_list_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "total_urls"),
				resource.TestCheckResourceAttr(resourceName, "urls.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "urls.0.pattern", "www.url1.com"),
				resource.TestCheckResourceAttr(resourceName, "urls.1.pattern", "www.url2.com"),
				resource.TestCheckResourceAttr(resourceName, "urls.0.type", "SIMPLE"),

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
					"oci_network_firewall_network_firewall_policy_url_lists",
					"test_network_firewall_policy_url_lists",
					acctest.Optional, acctest.Update,
					urlListDataSourceRepresentation,
				) +
				compartmentIdVariableStr + updatedUrlListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "url_list_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "url_list_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_url_list",
					"test_network_firewall_policy_url_list",
					acctest.Required, acctest.Create,
					urlListSingularDataSourceRepresentation,
				) +
				compartmentIdVariableStr + updatedUrlListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "name", "url_list_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_urls"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urls.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urls.0.pattern", "www.url1.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urls.1.pattern", "www.url2.com"),
				resource.TestCheckResourceAttr(singularDatasourceName, "urls.0.type", "SIMPLE"),
			),
		},
		// verify resource import
		{
			Config:                  config + createUrlListResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyUrlListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_url_list" {
			noResourceFound = false
			request := oci_network_firewall.GetUrlListRequest{}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.UrlListName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetUrlList(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyUrlList") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyUrlList", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyUrlList",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyUrlList"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyUrlListResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyUrlListResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyUrlListIds, err := getNetworkFirewallNetworkFirewallPolicyUrlListIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyUrlListId := range networkFirewallPolicyUrlListIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyUrlListId]; !ok {
			deleteUrlListRequest := oci_network_firewall.DeleteUrlListRequest{}

			deleteUrlListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteUrlList(context.Background(), deleteUrlListRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyUrlList %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyUrlListId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyUrlListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyUrlListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listUrlListsRequest := oci_network_firewall.ListUrlListsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyUrlList resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listUrlListsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listUrlListsResponse, err := networkFirewallClient.ListUrlLists(context.Background(), listUrlListsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyUrlList list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyUrlList := range listUrlListsResponse.Items {
			id := *networkFirewallPolicyUrlList.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyUrlListId", id)
		}

	}
	return resourceIds, nil
}
