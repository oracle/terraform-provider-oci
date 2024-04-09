// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_network_firewall "github.com/oracle/oci-go-sdk/v65/networkfirewall"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	createFirewallPolicyRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRequiredOnlyRepresentation,
	)

	NetworkFirewallPolicyResourceConfig = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Optional, acctest.Update,
		networkFirewallPolicyRepresentation,
	)

	networkFirewallPolicySingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	networkFirewallPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
	}

	networkFirewallPolicyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	networkFirewallPolicyRequiredOnlyRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`, Update: `${var.compartment_id}`},
	}
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_network_firewall_network_firewall_policy.test_network_firewall_policy"
	datasourceName := "data.oci_network_firewall_network_firewall_policies.test_network_firewall_policies"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy.test_network_firewall_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy",
			"test_network_firewall_policy",
			acctest.Optional, acctest.Create,
			networkFirewallPolicyRepresentation,
		), "network_firewall", "networkFirewallPolicy", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyDestroy, []resource.TestStep{
		// verify Create - step 0
		{
			Config: config + compartmentIdVariableStr + createFirewallPolicyRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "display_name"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create - step 1
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals - step 2
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy",
					"test_network_firewall_policy",
					acctest.Optional, acctest.Update,
					networkFirewallPolicyRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step) - step 3
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy",
					"test_network_firewall_policy",
					acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkFirewallPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters - step 4
		{
			Config: config + compartmentIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy",
					"test_network_firewall_policy",
					acctest.Optional, acctest.Update,
					networkFirewallPolicyRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify datasource - step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policies",
					"test_network_firewall_policies",
					acctest.Optional, acctest.Update,
					networkFirewallPolicyDataSourceRepresentation,
				) +
				compartmentIdVariableStr + NetworkFirewallPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "network_firewall_policy_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_firewall_policy_summary_collection.0.items.#", "2"),
			),
		},
		// verify singular datasource -step 6
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkFirewallPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "attached_network_firewall_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkFirewallPolicyResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy" {
			noResourceFound = false
			request := oci_network_firewall.GetNetworkFirewallPolicyRequest{}

			tmp := rs.Primary.ID
			request.NetworkFirewallPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			response, err := client.GetNetworkFirewallPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_network_firewall.LifecycleStateDeleted): true,
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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicy") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicy", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicy",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicy"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyIds, err := getNetworkFirewallPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyId]; !ok {
			deleteNetworkFirewallPolicyRequest := oci_network_firewall.DeleteNetworkFirewallPolicyRequest{}

			deleteNetworkFirewallPolicyRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

			deleteNetworkFirewallPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteNetworkFirewallPolicy(context.Background(), deleteNetworkFirewallPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkFirewallPolicyId, networkFirewallPolicySweepWaitCondition, time.Duration(3*time.Minute),
				networkFirewallPolicySweepResponseFetchOperation, "network_firewall", true)
		}
	}
	return nil
}

func getNetworkFirewallPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listNetworkFirewallPoliciesRequest := oci_network_firewall.ListNetworkFirewallPoliciesRequest{}
	listNetworkFirewallPoliciesRequest.CompartmentId = &compartmentId
	listNetworkFirewallPoliciesRequest.LifecycleState = oci_network_firewall.ListNetworkFirewallPoliciesLifecycleStateNeedsAttention
	listNetworkFirewallPoliciesResponse, err := networkFirewallClient.ListNetworkFirewallPolicies(context.Background(), listNetworkFirewallPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkFirewallPolicy := range listNetworkFirewallPoliciesResponse.Items {
		id := *networkFirewallPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyId", id)
	}
	return resourceIds, nil
}

func networkFirewallPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkFirewallPolicyResponse, ok := response.Response.(oci_network_firewall.GetNetworkFirewallPolicyResponse); ok {
		return networkFirewallPolicyResponse.LifecycleState != oci_network_firewall.LifecycleStateDeleted
	}
	return false
}

func networkFirewallPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NetworkFirewallClient().GetNetworkFirewallPolicy(context.Background(), oci_network_firewall.GetNetworkFirewallPolicyRequest{
		NetworkFirewallPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
