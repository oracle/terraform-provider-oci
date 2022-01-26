// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_waf "github.com/oracle/oci-go-sdk/v56/waf"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	WebAppFirewallRequiredOnlyResource = WebAppFirewallResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Required, acctest.Create, webAppFirewallRepresentation)

	WebAppFirewallResourceConfig = WebAppFirewallResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Update, webAppFirewallRepresentation)

	webAppFirewallSingularDataSourceRepresentation = map[string]interface{}{
		"web_app_firewall_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waf_web_app_firewall.test_web_app_firewall.id}`},
	}

	webAppFirewallDataSourceRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":                         acctest.Representation{RepType: acctest.Optional, Create: `${oci_waf_web_app_firewall.test_web_app_firewall.id}`},
		"state":                      acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"web_app_firewall_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
		"filter":                     acctest.RepresentationGroup{RepType: acctest.Required, Group: webAppFirewallDataSourceFilterRepresentation}}
	webAppFirewallDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waf_web_app_firewall.test_web_app_firewall.id}`}},
	}

	webAppFirewallRepresentation = map[string]interface{}{
		"backend_type":               acctest.Representation{RepType: acctest.Required, Create: `LOAD_BALANCER`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"load_balancer_id":           acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"web_app_firewall_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waf_web_app_firewall_policy.test_web_app_firewall_policy.id}`},
		//"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}}, //, Update: map[string]string{"Department": "Accounting"} but prevents from updating tags with policyID in body
	}

	WebAppFirewallResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall_policy", "test_web_app_firewall_policy", acctest.Required, acctest.Create, webAppFirewallPolicyRepresentation)
)

// issue-routing-tag: waf/default
func TestWafWebAppFirewallResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWafWebAppFirewallResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waf_web_app_firewall.test_web_app_firewall"
	datasourceName := "data.oci_waf_web_app_firewalls.test_web_app_firewalls"
	singularDatasourceName := "data.oci_waf_web_app_firewall.test_web_app_firewall"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WebAppFirewallResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Create, webAppFirewallRepresentation), "waf", "webAppFirewall", t)

	acctest.ResourceTest(t, testAccCheckWafWebAppFirewallDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Required, acctest.Create, webAppFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_firewall_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Create, webAppFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_firewall_policy_id"),

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

		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WebAppFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(webAppFirewallRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_firewall_policy_id"),

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
			Config: config + compartmentIdVariableStr + WebAppFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Update, webAppFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_firewall_policy_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewalls", "test_web_app_firewalls", acctest.Optional, acctest.Update, webAppFirewallDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Optional, acctest.Update, webAppFirewallRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "web_app_firewall_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "web_app_firewall_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "web_app_firewall_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waf_web_app_firewall", "test_web_app_firewall", acctest.Required, acctest.Create, webAppFirewallSingularDataSourceRepresentation) +
				compartmentIdVariableStr + WebAppFirewallResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "web_app_firewall_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + WebAppFirewallResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWafWebAppFirewallDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WafClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waf_web_app_firewall" {
			noResourceFound = false
			request := oci_waf.GetWebAppFirewallRequest{}

			tmp := rs.Primary.ID
			request.WebAppFirewallId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")

			response, err := client.GetWebAppFirewall(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waf.WebAppFirewallLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.GetLifecycleState())]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.GetLifecycleState())
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
	if !acctest.InSweeperExcludeList("WafWebAppFirewall") {
		resource.AddTestSweepers("WafWebAppFirewall", &resource.Sweeper{
			Name:         "WafWebAppFirewall",
			Dependencies: acctest.DependencyGraph["webAppFirewall"],
			F:            sweepWafWebAppFirewallResource,
		})
	}
}

func sweepWafWebAppFirewallResource(compartment string) error {
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()
	webAppFirewallIds, err := getWebAppFirewallIds(compartment)
	if err != nil {
		return err
	}
	for _, webAppFirewallId := range webAppFirewallIds {
		if ok := acctest.SweeperDefaultResourceId[webAppFirewallId]; !ok {
			deleteWebAppFirewallRequest := oci_waf.DeleteWebAppFirewallRequest{}

			deleteWebAppFirewallRequest.WebAppFirewallId = &webAppFirewallId

			deleteWebAppFirewallRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waf")
			_, error := wafClient.DeleteWebAppFirewall(context.Background(), deleteWebAppFirewallRequest)
			if error != nil {
				fmt.Printf("Error deleting WebAppFirewall %s %s, It is possible that the resource is already deleted. Please verify manually \n", webAppFirewallId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &webAppFirewallId, webAppFirewallSweepWaitCondition, time.Duration(3*time.Minute),
				webAppFirewallSweepResponseFetchOperation, "waf", true)
		}
	}
	return nil
}

func getWebAppFirewallIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WebAppFirewallId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	wafClient := acctest.GetTestClients(&schema.ResourceData{}).WafClient()

	listWebAppFirewallsRequest := oci_waf.ListWebAppFirewallsRequest{}
	listWebAppFirewallsRequest.CompartmentId = &compartmentId
	listWebAppFirewallsRequest.LifecycleState = []oci_waf.WebAppFirewallLifecycleStateEnum{oci_waf.WebAppFirewallLifecycleStateActive}
	listWebAppFirewallsResponse, err := wafClient.ListWebAppFirewalls(context.Background(), listWebAppFirewallsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WebAppFirewall list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, webAppFirewall := range listWebAppFirewallsResponse.Items {
		id := *webAppFirewall.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WebAppFirewallId", id)
	}
	return resourceIds, nil
}

func webAppFirewallSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if webAppFirewallResponse, ok := response.Response.(oci_waf.GetWebAppFirewallResponse); ok {
		return webAppFirewallResponse.GetLifecycleState() != oci_waf.WebAppFirewallLifecycleStateDeleted
	}
	return false
}

func webAppFirewallSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WafClient().GetWebAppFirewall(context.Background(), oci_waf.GetWebAppFirewallRequest{
		WebAppFirewallId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
