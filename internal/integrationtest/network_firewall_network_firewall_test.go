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
	NetworkFirewallRequiredOnlyResource = NetworkFirewallResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Required, acctest.Create, networkFirewallRepresentation)

	NetworkFirewallResourceConfig = NetworkFirewallResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Optional, acctest.Update, networkFirewallRepresentation)

	networkFirewallSingularDataSourceRepresentation = map[string]interface{}{
		"network_firewall_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall.test_network_firewall.id}`},
	}

	networkFirewallDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MyFirewall`, Update: `displayName2`},
		"id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_firewall_network_firewall.test_network_firewall.id}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: networkFirewallDataSourceFilterRepresentation}}
	networkFirewallDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_network_firewall_network_firewall.test_network_firewall.id}`}},
	}

	networkFirewallRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"subnet_id":                  acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"availability_domain":        acctest.Representation{RepType: acctest.Optional, Create: `${data.oci_identity_availability_domains.test_availability_domains.availability_domains.0.name}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `MyFirewall`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"ipv4address":                acctest.Representation{RepType: acctest.Optional, Create: `10.0.0.3`},
	}

	NetworkFirewallResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, CoreSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, CoreVcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall_policy", "test_network_firewall_policy", acctest.Required, acctest.Create, networkFirewallPolicyRepresentation)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_network_firewall_network_firewall.test_network_firewall"
	datasourceName := "data.oci_network_firewall_network_firewalls.test_network_firewalls"
	singularDatasourceName := "data.oci_network_firewall_network_firewall.test_network_firewall"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NetworkFirewallResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Optional, acctest.Create, networkFirewallRepresentation), "networkfirewall", "networkFirewall", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NetworkFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap(
					"oci_network_firewall_network_firewall",
					"test_network_firewall",
					acctest.Optional, acctest.Create,
					networkFirewallRepresentation,
				),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyFirewall"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NetworkFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(networkFirewallRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "MyFirewall"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NetworkFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Optional, acctest.Update, networkFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "availability_domain"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "ipv4address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "subnet_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
					"oci_network_firewall_network_firewalls",
					"test_network_firewalls",
					acctest.Optional, acctest.Update,
					networkFirewallDataSourceRepresentation,
				) +
				compartmentIdVariableStr + NetworkFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Optional, acctest.Update, networkFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "network_firewall_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "network_firewall_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_network_firewall_network_firewall", "test_network_firewall", acctest.Required, acctest.Create, networkFirewallSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NetworkFirewallResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "availability_domain"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "ipv4address", "10.0.0.3"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + NetworkFirewallRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall" {
			noResourceFound = false
			request := oci_network_firewall.GetNetworkFirewallRequest{}

			tmp := rs.Primary.ID
			request.NetworkFirewallId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			response, err := client.GetNetworkFirewall(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewall") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewall", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewall",
			Dependencies: acctest.DependencyGraph["networkFirewall"],
			F:            sweepNetworkFirewallNetworkFirewallResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallIds, err := getNetworkFirewallIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallId := range networkFirewallIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallId]; !ok {
			deleteNetworkFirewallRequest := oci_network_firewall.DeleteNetworkFirewallRequest{}

			deleteNetworkFirewallRequest.NetworkFirewallId = &networkFirewallId

			deleteNetworkFirewallRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteNetworkFirewall(context.Background(), deleteNetworkFirewallRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewall %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &networkFirewallId, networkFirewallSweepWaitCondition, time.Duration(3*time.Minute),
				networkFirewallSweepResponseFetchOperation, "network_firewall", true)
		}
	}
	return nil
}

func getNetworkFirewallIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listNetworkFirewallsRequest := oci_network_firewall.ListNetworkFirewallsRequest{}
	listNetworkFirewallsRequest.CompartmentId = &compartmentId
	listNetworkFirewallsRequest.LifecycleState = oci_network_firewall.ListNetworkFirewallsLifecycleStateNeedsAttention
	listNetworkFirewallsResponse, err := networkFirewallClient.ListNetworkFirewalls(context.Background(), listNetworkFirewallsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NetworkFirewall list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, networkFirewall := range listNetworkFirewallsResponse.Items {
		id := *networkFirewall.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallId", id)
	}
	return resourceIds, nil
}

func networkFirewallSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if networkFirewallResponse, ok := response.Response.(oci_network_firewall.GetNetworkFirewallResponse); ok {
		return networkFirewallResponse.LifecycleState != oci_network_firewall.LifecycleStateDeleted
	}
	return false
}

func networkFirewallSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NetworkFirewallClient().GetNetworkFirewall(context.Background(), oci_network_firewall.GetNetworkFirewallRequest{
		NetworkFirewallId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
