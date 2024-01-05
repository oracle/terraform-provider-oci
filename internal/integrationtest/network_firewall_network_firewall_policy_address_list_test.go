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
	createAddressListResourceConfig = addressListResourceDependencies + createAddressListResource
	createAddressListResource       = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_address_list",
		"test_network_firewall_policy_address_list",
		acctest.Required, acctest.Create,
		addressListRepresentationFQDN,
	)

	addressListSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	addressListsDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	addressListRepresentationFQDN = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `address_list_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `FQDN`, Update: `FQDN`},
		"addresses":                  acctest.Representation{RepType: acctest.Required, Create: []string{`www.google.com`, `www.facebook.com`}, Update: []string{`www.google.com`, `www.twitter.com`}},
	}

	addressListResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required,
		acctest.Create,
		networkFirewallPolicyRepresentation)

	addressListResourceConfig = addressListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_address_list",
			"test_network_firewall_policy_address_list",
			acctest.Optional,
			acctest.Update,
			addressListRepresentationFQDN,
		)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyAddressListResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyAddressListResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_address_lists.test_network_firewall_policy_address_lists"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_address_list.test_network_firewall_policy_address_list"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(
		config+compartmentIdVariableStr+addressListResourceDependencies+
			acctest.GenerateResourceFromRepresentationMap(
				"oci_network_firewall_network_firewall_policy_address_list",
				"test_network_firewall_policy_address_list",
				acctest.Required,
				acctest.Create,
				addressListRepresentationFQDN),
		"networkfirewall",
		"networkFirewallPolicyAddressList", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyAddressListDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createAddressListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "address_list_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "FQDN"),

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
			Config: config + compartmentIdVariableStr + addressListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "addresses.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "name", "address_list_1"),
				resource.TestCheckResourceAttr(resourceName, "type", "FQDN"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttrSet(resourceName, "total_addresses"),

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
					"oci_network_firewall_network_firewall_policy_address_lists",
					"test_network_firewall_policy_address_lists",
					acctest.Optional,
					acctest.Update,
					addressListsDataSourceRepresentation) +
				compartmentIdVariableStr +
				addressListResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(datasourceName, "address_list_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "address_list_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_address_list",
					"test_network_firewall_policy_address_list",
					acctest.Required,
					acctest.Create,
					addressListSingularDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				addressListResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "addresses.#", "2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "address_list_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_addresses"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "FQDN"),
			),
		},
		// verify resource import
		{
			Config:                  config + addressListResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyAddressListDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_address_list" {
			noResourceFound = false
			request := oci_network_firewall.GetAddressListRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.AddressListName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetAddressList(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyAddressList") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyAddressList", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyAddressList",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyAddressList"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyAddressListResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyAddressListResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyAddressListIds, err := getNetworkFirewallNetworkFirewallPolicyAddressListIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyAddressListId := range networkFirewallPolicyAddressListIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyAddressListId]; !ok {
			deleteAddressListRequest := oci_network_firewall.DeleteAddressListRequest{}

			deleteAddressListRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteAddressList(context.Background(), deleteAddressListRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyAddressList %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyAddressListId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyAddressListIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyAddressListId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listAddressListsRequest := oci_network_firewall.ListAddressListsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallNetworkFirewallPolicyAddressListIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyAddressList resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listAddressListsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listAddressListsResponse, err := networkFirewallClient.ListAddressLists(context.Background(), listAddressListsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyAddressList list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyAddressList := range listAddressListsResponse.Items {
			id := *networkFirewallPolicyAddressList.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyAddressListId", id)
		}

	}
	return resourceIds, nil
}
