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
	createApplicationResource = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy_application",
		"test_network_firewall_policy_application",
		acctest.Required,
		acctest.Create,
		applicationRepresentationRequiredOnly,
	)
	createApplicationRequiredOnlyResourceConfig = applicationResourceDependencies + createApplicationResource

	createApplicationResourceConfig = applicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_application",
			"test_network_firewall_policy_application",
			acctest.Required,
			acctest.Create,
			applicationRepresentationWithOptionals,
		)

	singularApplicationDataSourceRepresentation = map[string]interface{}{
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application.name}`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	listApplicationDataSourceRepresentation = map[string]interface{}{
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
	}

	applicationRepresentationWithOptionals = map[string]interface{}{
		"icmp_type":                  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `application_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"icmp_code":                  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `ICMP`, Update: `ICMP`},
	}

	applicationRepresentationRequiredOnly = map[string]interface{}{
		"icmp_type":                  acctest.Representation{RepType: acctest.Required, Create: `10`, Update: `11`},
		"name":                       acctest.Representation{RepType: acctest.Required, Create: `application_1`},
		"network_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_network_firewall_network_firewall_policy.test_network_firewall_policy.id}`},
		"type":                       acctest.Representation{RepType: acctest.Required, Create: `ICMP`, Update: `ICMP`},
	}

	applicationResourceConfig = applicationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_application",
			"test_network_firewall_policy_application",
			acctest.Optional,
			acctest.Update,
			applicationRepresentationWithOptionals,
		)

	applicationResourceDependencies = acctest.GenerateResourceFromRepresentationMap(
		"oci_network_firewall_network_firewall_policy",
		"test_network_firewall_policy",
		acctest.Required,
		acctest.Create,
		networkFirewallPolicyRepresentation,
	)
)

// issue-routing-tag: network_firewall/default
func TestNetworkFirewallNetworkFirewallPolicyApplicationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestNetworkFirewallNetworkFirewallPolicyApplicationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application"
	datasourceName := "data.oci_network_firewall_network_firewall_policy_applications.test_network_firewall_policy_applications"
	singularDatasourceName := "data.oci_network_firewall_network_firewall_policy_application.test_network_firewall_policy_application"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+applicationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap(
			"oci_network_firewall_network_firewall_policy_application",
			"test_network_firewall_policy_application",
			acctest.Optional,
			acctest.Create,
			applicationRepresentationWithOptionals,
		),
		"networkfirewall", "networkFirewallPolicyApplication", t)

	acctest.ResourceTest(t, testAccCheckNetworkFirewallNetworkFirewallPolicyApplicationDestroy, []resource.TestStep{
		// verify create with required only - step 0
		{
			Config: config + compartmentIdVariableStr + createApplicationRequiredOnlyResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_type", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_1"),
				resource.TestCheckResourceAttr(resourceName, "type", "ICMP"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create - step 1
		{
			Config: config + compartmentIdVariableStr + applicationResourceDependencies,
		},

		// verify Create with optionals - step 2
		{
			Config: config + compartmentIdVariableStr + createApplicationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_code", "10"),
				resource.TestCheckResourceAttr(resourceName, "icmp_type", "10"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ICMP"),

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
			Config: config + compartmentIdVariableStr + applicationResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "icmp_code", "11"),
				resource.TestCheckResourceAttr(resourceName, "icmp_type", "11"),
				resource.TestCheckResourceAttr(resourceName, "name", "application_1"),
				resource.TestCheckResourceAttrSet(resourceName, "network_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(resourceName, "type", "ICMP"),

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
					"oci_network_firewall_network_firewall_policy_applications",
					"test_network_firewall_policy_applications",
					acctest.Optional, acctest.Update,
					listApplicationDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				applicationResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "application_summary_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "application_summary_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource - step 5
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap(
					"oci_network_firewall_network_firewall_policy_application",
					"test_network_firewall_policy_application",
					acctest.Required, acctest.Create,
					singularApplicationDataSourceRepresentation,
				) +
				compartmentIdVariableStr +
				applicationResourceConfig,

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "network_firewall_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "icmp_code", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "icmp_type", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", "application_1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "parent_resource_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "type", "ICMP"),
			),
		},
		// verify resource import
		{
			Config:                  config + createApplicationRequiredOnlyResourceConfig,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckNetworkFirewallNetworkFirewallPolicyApplicationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NetworkFirewallClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_network_firewall_network_firewall_policy_application" {
			noResourceFound = false
			request := oci_network_firewall.GetApplicationRequest{}

			if value, ok := rs.Primary.Attributes["name"]; ok {
				request.ApplicationName = &value
			}

			if value, ok := rs.Primary.Attributes["network_firewall_policy_id"]; ok {
				request.NetworkFirewallPolicyId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")

			_, err := client.GetApplication(context.Background(), request)

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
	if !acctest.InSweeperExcludeList("NetworkFirewallNetworkFirewallPolicyApplication") {
		resource.AddTestSweepers("NetworkFirewallNetworkFirewallPolicyApplication", &resource.Sweeper{
			Name:         "NetworkFirewallNetworkFirewallPolicyApplication",
			Dependencies: acctest.DependencyGraph["networkFirewallPolicyApplication"],
			F:            sweepNetworkFirewallNetworkFirewallPolicyApplicationResource,
		})
	}
}

func sweepNetworkFirewallNetworkFirewallPolicyApplicationResource(compartment string) error {
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()
	networkFirewallPolicyApplicationIds, err := getNetworkFirewallNetworkFirewallPolicyApplicationIds(compartment)
	if err != nil {
		return err
	}
	for _, networkFirewallPolicyApplicationId := range networkFirewallPolicyApplicationIds {
		if ok := acctest.SweeperDefaultResourceId[networkFirewallPolicyApplicationId]; !ok {
			deleteApplicationRequest := oci_network_firewall.DeleteApplicationRequest{}

			deleteApplicationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "network_firewall")
			_, error := networkFirewallClient.DeleteApplication(context.Background(), deleteApplicationRequest)
			if error != nil {
				fmt.Printf("Error deleting NetworkFirewallPolicyApplication %s %s, It is possible that the resource is already deleted. Please verify manually \n", networkFirewallPolicyApplicationId, error)
				continue
			}
		}
	}
	return nil
}

func getNetworkFirewallNetworkFirewallPolicyApplicationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NetworkFirewallPolicyApplicationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	networkFirewallClient := acctest.GetTestClients(&schema.ResourceData{}).NetworkFirewallClient()

	listApplicationsRequest := oci_network_firewall.ListApplicationsRequest{}

	networkFirewallPolicyIds, error := getNetworkFirewallPolicyIds(compartment)
	if error != nil {
		return resourceIds, fmt.Errorf("Error getting networkFirewallPolicyId required for NetworkFirewallPolicyApplication resource requests \n")
	}
	for _, networkFirewallPolicyId := range networkFirewallPolicyIds {
		listApplicationsRequest.NetworkFirewallPolicyId = &networkFirewallPolicyId

		listApplicationsResponse, err := networkFirewallClient.ListApplications(context.Background(), listApplicationsRequest)

		if err != nil {
			return resourceIds, fmt.Errorf("Error getting NetworkFirewallPolicyApplication list for compartment id : %s , %s \n", compartmentId, err)
		}
		for _, networkFirewallPolicyApplication := range listApplicationsResponse.Items {
			id := *networkFirewallPolicyApplication.GetName()
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NetworkFirewallPolicyApplicationId", id)
		}

	}
	return resourceIds, nil
}
