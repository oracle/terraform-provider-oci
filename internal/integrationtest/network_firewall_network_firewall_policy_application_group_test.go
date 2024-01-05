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
	createApplicationGroupResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_application_group",
		"test_network_firewall_policy_application_group",
		acctest.Required, acctest.Create,
		applicationGroupRepresentation,
	)
	createApplicationGroupResourceConfig = applicationGroupResourceDependencies + createApplicationGroupResource

	applicationGroupResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_application_group",
		"test_network_firewall_policy_application_group",
		acctest.Optional, acctest.Update,
		applicationGroupRepresentation,
	)

	applicationGroupResourceConfig = applicationGroupResourceDependencies + applicationGroupResource

	applicationGroupSingularDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	applicationGroupListDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	emptyApps = make([]string, 0)

	applicationGroupRepresentation = map[string]interface{}{
		"apps":                       acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application.name}`}, Update: emptyApps},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `application_group_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	applicationGroupResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required,
		acctest.Create,
		networkFirewallPolicyRepresentation,
	) + acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_application",
		"test_network_firewall_policy_application",
		acctest.Required,
		acctest.Create,
		applicationRepresentationWithOptionals,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyApplicationGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyApplicationGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_application_groups.test_network_firewall_policy_application_groups"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_application_group.test_network_firewall_policy_application_group"

	var resId, resId2 string
	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+createApplicationGroupResourceConfig, "networkfirewall", "networkFirewallPolicyApplicationGroup", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyApplicationGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + createApplicationGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "apps.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_group_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "total_apps", "1"),

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
			Config: config + compartmentIdVariableStr + applicationGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "apps.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_group_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "total_apps", "0"),

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
					"oci_network_firewall_network_firewall_policy_application_groups",
					"test_network_firewall_policy_application_groups",
					acctest.Optional, acctest.Update,
					applicationGroupListDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				applicationGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "application_group_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "application_group_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_application_group",
					"test_network_firewall_policy_application_group",
					acctest.Required, acctest.Create,
					applicationGroupSingularDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				applicationGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "apps.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "application_group_1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "total_apps", "0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + applicationGroupResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyApplicationGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_application_group" {
			noResourceFound = false
			request := oci_network_firewall.GetApplicationGroupRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.ApplicationGroupName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetApplicationGroup(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyApplicationGroup") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyApplicationGroup", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyApplicationGroup",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyApplicationGroup"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyApplicationGroupResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyApplicationGroupResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyApplicationGroupIds, err := getNetworkFirewallNetworkFirewallPolicyApplicationGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyApplicationGroupId := range networkFirewallPolicyApplicationGroupIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyApplicationGroupId]; !ok {
			deleteApplicationGroupRequest := oci_network_firewall.DeleteApplicationGroupRequest{}

			deleteApplicationGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteApplicationGroup(context.Background(), deleteApplicationGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyApplicationGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyApplicationGroupId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyApplicationGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyApplicationGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listApplicationGroupsRequest := oci_network_firewall.ListApplicationGroupsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyApplicationGroup resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listApplicationGroupsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listApplicationGroupsResponse, err := networkFirewallClient.ListApplicationGroups(context.Background(), listApplicationGroupsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyApplicationGroup list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyApplicationGroup := range listApplicationGroupsResponse.Items {
			id := *networkFirewallPolicyApplicationGroup.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyApplicationGroupId", id)
		}

	}
	return resourceIds, nil
}
