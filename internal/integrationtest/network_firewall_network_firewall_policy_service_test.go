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
	createServiceResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_service",
		"test_network_firewall_policy_service",
		acctest.Required, acctest.Create,
		serviceRepresentation,
	)

	createServiceResourceConfig = serviceResourceDependencies + createServiceResource

	createServiceResourceConfigOptional = serviceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_service",
			"test_network_firewall_policy_service",
			acctest.Required, acctest.Create,
			serviceRepresentation)

	serviceResourceConfig = serviceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_service",
			"test_network_firewall_policy_service",
			acctest.Optional, acctest.Update,
			serviceRepresentation)

	singularServiceDataSourceRep = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	listServicesDataSourceRep = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	serviceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `service_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"port_ranges":                acctest.RepresentationGroup{RepType: acctest.Required, Group: servicePortRangesRepresentation},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `TCP_SERVICE`, Update: `TCP_SERVICE`},
	}

	servicePortRangesRepresentation = map[string]interface{}{
		"minimum_port": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"maximum_port": acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
	}

	serviceResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required, acctest.Create,
		networkFirewallPolicyRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyServiceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyServiceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_services.test_network_firewall_policy_services"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_service.test_network_firewall_policy_service"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+serviceResourceConfig, "networkfirewall", "networkFirewallPolicyService", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyServiceDestroy, []resource.TestStep{
		// verify Create - step 0
		{
			Config: config + compartmentIdVariableStr + createServiceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "service_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.0.minimum_port", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create - step 1
		{
			Config: config + compartmentIdVariableStr + serviceResourceDependencies,
		},

		// verify Create with optionals - step 2
		{
			Config: config + compartmentIdVariableStr + createServiceResourceConfigOptional,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "service_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.0.maximum_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.0.minimum_port", "10"),
				resource.TestCheckResourceAttr(resourceName, "type", "TCP_SERVICE"),

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

		// verify updates to updatable parameters - step 3
		{
			Config: config + compartmentIdVariableStr + serviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "name", "service_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.0.maximum_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "port_ranges.0.minimum_port", "11"),
				resource.TestCheckResourceAttr(resourceName, "type", "TCP_SERVICE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource - step 4
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_services",
					"test_network_firewall_policy_services",
					acctest.Optional, acctest.Update,
					listServicesDataSourceRep) +
				compartmentIdVariableStr + serviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "service_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "service_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource - step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_service",
					"test_network_firewall_policy_service",
					acctest.Required, acctest.Create,
					singularServiceDataSourceRep) +
				compartmentIdVariableStr + serviceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "name", "service_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port_ranges.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port_ranges.0.maximum_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "port_ranges.0.minimum_port", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "TCP_SERVICE"),
			),
		},
		// verify resource import
		{
			Config:                  config + createServiceResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyServiceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_service" {
			noResourceFound = false
			request := oci_network_firewall.GetServiceRequest{}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.ServiceName = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetService(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyService") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyService", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyService",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyService"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyServiceResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyServiceResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyServiceIds, err := getNetworkFirewallNetworkFirewallPolicyServiceIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyServiceId := range networkFirewallPolicyServiceIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyServiceId]; !ok {
			deleteServiceRequest := oci_network_firewall.DeleteServiceRequest{}

			deleteServiceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteService(context.Background(), deleteServiceRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyService %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyServiceId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyServiceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyServiceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listServicesRequest := oci_network_firewall.ListServicesRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyService resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listServicesRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listServicesResponse, err := networkFirewallClient.ListServices(context.Background(), listServicesRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyService list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyService := range listServicesResponse.Items {
			id := *networkFirewallPolicyService.Name
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyServiceId", id)
		}

	}
	return resourceIds, nil
}
