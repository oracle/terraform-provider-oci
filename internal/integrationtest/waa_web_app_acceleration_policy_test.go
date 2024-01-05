// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	WaaWebAppAccelerationPolicyRequiredOnlyResource = WaaWebAppAccelerationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Required, acctest.Create, WaaWebAppAccelerationPolicyRepresentation)

	WaaWebAppAccelerationPolicyResourceConfig = WaaWebAppAccelerationPolicyResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Update, WaaWebAppAccelerationPolicyRepresentation)

	WaaWaaWebAppAccelerationPolicySingularDataSourceRepresentation = map[string]interface{}{
		"web_app_acceleration_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id}`},
	}

	WaaWaaWebAppAccelerationPolicyDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `WAAPolicy1`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: WaaWebAppAccelerationPolicyDataSourceFilterRepresentation}}
	WaaWebAppAccelerationPolicyDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id}`}},
	}

	WaaWebAppAccelerationPolicyRepresentation = map[string]interface{}{
		"compartment_id":              acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                acctest.Representation{RepType: acctest.Optional, Create: `WAAPolicy1`, Update: `displayName2`},
		"freeform_tags":               acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"response_caching_policy":     acctest.RepresentationGroup{RepType: acctest.Optional, Group: WaaWebAppAccelerationPolicyResponseCachingPolicyRepresentation},
		"response_compression_policy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: WaaWebAppAccelerationPolicyResponseCompressionPolicyRepresentation},
	}
	WaaWebAppAccelerationPolicyResponseCachingPolicyRepresentation = map[string]interface{}{
		"is_response_header_based_caching_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}
	WaaWebAppAccelerationPolicyResponseCompressionPolicyRepresentation = map[string]interface{}{
		"gzip_compression": acctest.RepresentationGroup{RepType: acctest.Optional, Group: WaaWebAppAccelerationPolicyResponseCompressionPolicyGzipCompressionRepresentation},
	}
	WaaWebAppAccelerationPolicyResponseCompressionPolicyGzipCompressionRepresentation = map[string]interface{}{
		"is_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
	}

	WaaWebAppAccelerationPolicyResourceDependencies = ""
)

// issue-routing-tag: waa/default
func TestWaaWebAppAccelerationPolicyResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaaWebAppAccelerationPolicyResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy"
	datasourceName := "data.oci_waa_web_app_acceleration_policies.test_web_app_acceleration_policies"
	singularDatasourceName := "data.oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WaaWebAppAccelerationPolicyResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Create, WaaWebAppAccelerationPolicyRepresentation), "waa", "webAppAccelerationPolicy", t)

	acctest.ResourceTest(t, testAccCheckWaaWebAppAccelerationPolicyDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Required, acctest.Create, WaaWebAppAccelerationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Create, WaaWebAppAccelerationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "WAAPolicy1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.0.is_response_header_based_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WaaWebAppAccelerationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(WaaWebAppAccelerationPolicyRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "WAAPolicy1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.0.is_response_header_based_caching_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.0.is_enabled", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Update, WaaWebAppAccelerationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_caching_policy.0.is_response_header_based_caching_enabled", "true"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "response_compression_policy.0.gzip_compression.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waa_web_app_acceleration_policies", "test_web_app_acceleration_policies", acctest.Optional, acctest.Update, WaaWaaWebAppAccelerationPolicyDataSourceRepresentation) +
				compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Optional, acctest.Update, WaaWebAppAccelerationPolicyRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),

				resource.TestCheckResourceAttr(datasourceName, "web_app_acceleration_policy_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "web_app_acceleration_policy_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Required, acctest.Create, WaaWaaWebAppAccelerationPolicySingularDataSourceRepresentation) +
				compartmentIdVariableStr + WaaWebAppAccelerationPolicyResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "web_app_acceleration_policy_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_caching_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_caching_policy.0.is_response_header_based_caching_enabled", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_compression_policy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_compression_policy.0.gzip_compression.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "response_compression_policy.0.gzip_compression.0.is_enabled", "true"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + WaaWebAppAccelerationPolicyRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaaWebAppAccelerationPolicyDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WaaClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waa_web_app_acceleration_policy" {
			noResourceFound = false
			request := oci_waa.GetWebAppAccelerationPolicyRequest{}

			tmp := rs.Primary.ID
			request.WebAppAccelerationPolicyId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waa")

			response, err := client.GetWebAppAccelerationPolicy(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waa.WebAppAccelerationPolicyLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("WaaWebAppAccelerationPolicy") {
		resource.AddTestSweepers("WaaWebAppAccelerationPolicy", &resource.Sweeper{
			Name:         "WaaWebAppAccelerationPolicy",
			Dependencies: acctest.DependencyGraph["webAppAccelerationPolicy"],
			F:            sweepWaaWebAppAccelerationPolicyResource,
		})
	}
}

func sweepWaaWebAppAccelerationPolicyResource(compartment string) error {
	waaClient := acctest.GetTestClients(&schema.ResourceData{}).WaaClient()
	webAppAccelerationPolicyIds, err := getWaaWebAppAccelerationPolicyIds(compartment)
	if err != nil {
		return err
	}
	for _, webAppAccelerationPolicyId := range webAppAccelerationPolicyIds {
		if ok := acctest.SweeperDefaultResourceId[webAppAccelerationPolicyId]; !ok {
			deleteWebAppAccelerationPolicyRequest := oci_waa.DeleteWebAppAccelerationPolicyRequest{}

			deleteWebAppAccelerationPolicyRequest.WebAppAccelerationPolicyId = &webAppAccelerationPolicyId

			deleteWebAppAccelerationPolicyRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waa")
			_, error := waaClient.DeleteWebAppAccelerationPolicy(context.Background(), deleteWebAppAccelerationPolicyRequest)
			if error != nil {
				fmt.Printf("Error deleting WebAppAccelerationPolicy %s %s, It is possible that the resource is already deleted. Please verify manually \n", webAppAccelerationPolicyId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &webAppAccelerationPolicyId, WaaWebAppAccelerationPolicySweepWaitCondition, time.Duration(3*time.Minute),
				WaaWebAppAccelerationPolicySweepResponseFetchOperation, "waa", true)
		}
	}
	return nil
}

func getWaaWebAppAccelerationPolicyIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WebAppAccelerationPolicyId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waaClient := acctest.GetTestClients(&schema.ResourceData{}).WaaClient()

	listWebAppAccelerationPoliciesRequest := oci_waa.ListWebAppAccelerationPoliciesRequest{}
	listWebAppAccelerationPoliciesRequest.CompartmentId = &compartmentId
	listWebAppAccelerationPoliciesRequest.LifecycleState = []oci_waa.WebAppAccelerationPolicyLifecycleStateEnum{oci_waa.WebAppAccelerationPolicyLifecycleStateActive}
	listWebAppAccelerationPoliciesResponse, err := waaClient.ListWebAppAccelerationPolicies(context.Background(), listWebAppAccelerationPoliciesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WebAppAccelerationPolicy list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, webAppAccelerationPolicy := range listWebAppAccelerationPoliciesResponse.Items {
		id := *webAppAccelerationPolicy.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WebAppAccelerationPolicyId", id)
	}
	return resourceIds, nil
}

func WaaWebAppAccelerationPolicySweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if webAppAccelerationPolicyResponse, ok := response.Response.(oci_waa.GetWebAppAccelerationPolicyResponse); ok {
		return webAppAccelerationPolicyResponse.LifecycleState != oci_waa.WebAppAccelerationPolicyLifecycleStateDeleted
	}
	return false
}

func WaaWebAppAccelerationPolicySweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WaaClient().GetWebAppAccelerationPolicy(context.Background(), oci_waa.GetWebAppAccelerationPolicyRequest{
		WebAppAccelerationPolicyId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
