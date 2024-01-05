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
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	WaaWebAppAccelerationRequiredOnlyResource = WaaWebAppAccelerationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Required, acctest.Create, WaaWebAppAccelerationRepresentation)

	WaaWebAppAccelerationResourceConfig = WaaWebAppAccelerationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Update, WaaWebAppAccelerationRepresentation)

	WaaWaaWebAppAccelerationSingularDataSourceRepresentation = map[string]interface{}{
		"web_app_acceleration_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waa_web_app_acceleration.test_web_app_acceleration.id}`},
	}

	WaaWaaWebAppAccelerationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `WAA1`, Update: `displayName2`},
		"id":                             acctest.Representation{RepType: acctest.Optional, Create: `${oci_waa_web_app_acceleration.test_web_app_acceleration.id}`},
		"state":                          acctest.Representation{RepType: acctest.Optional, Create: []string{`ACTIVE`}},
		"web_app_acceleration_policy_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id}`},
		"filter":                         acctest.RepresentationGroup{RepType: acctest.Required, Group: WaaWebAppAccelerationDataSourceFilterRepresentation}}
	WaaWebAppAccelerationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_waa_web_app_acceleration.test_web_app_acceleration.id}`}},
	}

	WaaWebAppAccelerationRepresentation = map[string]interface{}{
		"backend_type":                   acctest.Representation{RepType: acctest.Required, Create: `LOAD_BALANCER`},
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"load_balancer_id":               acctest.Representation{RepType: acctest.Required, Create: `${oci_load_balancer_load_balancer.test_load_balancer.id}`},
		"web_app_acceleration_policy_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_waa_web_app_acceleration_policy.test_web_app_acceleration_policy.id}`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `WAA1`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
	}

	WaaWebAppAccelerationResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_load_balancer_load_balancer", "test_load_balancer", acctest.Required, acctest.Create, loadBalancerRepresentation) +
		LoadBalancerSubnetDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration_policy", "test_web_app_acceleration_policy", acctest.Required, acctest.Create, WaaWebAppAccelerationPolicyRepresentation)
)

// issue-routing-tag: waa/default
func TestWaaWebAppAccelerationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestWaaWebAppAccelerationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_waa_web_app_acceleration.test_web_app_acceleration"
	datasourceName := "data.oci_waa_web_app_accelerations.test_web_app_accelerations"
	singularDatasourceName := "data.oci_waa_web_app_acceleration.test_web_app_acceleration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+WaaWebAppAccelerationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Create, WaaWebAppAccelerationRepresentation), "waa", "webAppAcceleration", t)

	acctest.ResourceTest(t, testAccCheckWaaWebAppAccelerationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Required, acctest.Create, WaaWebAppAccelerationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_acceleration_policy_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Create, WaaWebAppAccelerationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "WAA1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_acceleration_policy_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + WaaWebAppAccelerationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(WaaWebAppAccelerationRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "backend_type", "LOAD_BALANCER"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "WAA1"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "load_balancer_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "system_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "web_app_acceleration_policy_id"),

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
			Config: config + compartmentIdVariableStr + WaaWebAppAccelerationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Update, WaaWebAppAccelerationRepresentation),
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
				resource.TestCheckResourceAttrSet(resourceName, "web_app_acceleration_policy_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_waa_web_app_accelerations", "test_web_app_accelerations", acctest.Optional, acctest.Update, WaaWaaWebAppAccelerationDataSourceRepresentation) +
				compartmentIdVariableStr + WaaWebAppAccelerationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Optional, acctest.Update, WaaWebAppAccelerationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "web_app_acceleration_policy_id"),

				resource.TestCheckResourceAttr(datasourceName, "web_app_acceleration_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "web_app_acceleration_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_waa_web_app_acceleration", "test_web_app_acceleration", acctest.Required, acctest.Create, WaaWaaWebAppAccelerationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + WaaWebAppAccelerationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "web_app_acceleration_id"),

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
		// verify resource import
		{
			Config:                  config + WaaWebAppAccelerationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckWaaWebAppAccelerationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).WaaClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_waa_web_app_acceleration" {
			noResourceFound = false
			request := oci_waa.GetWebAppAccelerationRequest{}

			tmp := rs.Primary.ID
			request.WebAppAccelerationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waa")

			response, err := client.GetWebAppAcceleration(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_waa.WebAppAccelerationLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("WaaWebAppAcceleration") {
		resource.AddTestSweepers("WaaWebAppAcceleration", &resource.Sweeper{
			Name:         "WaaWebAppAcceleration",
			Dependencies: acctest.DependencyGraph["webAppAcceleration"],
			F:            sweepWaaWebAppAccelerationResource,
		})
	}
}

func sweepWaaWebAppAccelerationResource(compartment string) error {
	waaClient := acctest.GetTestClients(&schema.ResourceData{}).WaaClient()
	webAppAccelerationIds, err := getWaaWebAppAccelerationIds(compartment)
	if err != nil {
		return err
	}
	for _, webAppAccelerationId := range webAppAccelerationIds {
		if ok := acctest.SweeperDefaultResourceId[webAppAccelerationId]; !ok {
			deleteWebAppAccelerationRequest := oci_waa.DeleteWebAppAccelerationRequest{}

			deleteWebAppAccelerationRequest.WebAppAccelerationId = &webAppAccelerationId

			deleteWebAppAccelerationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "waa")
			_, error := waaClient.DeleteWebAppAcceleration(context.Background(), deleteWebAppAccelerationRequest)
			if error != nil {
				fmt.Printf("Error deleting WebAppAcceleration %s %s, It is possible that the resource is already deleted. Please verify manually \n", webAppAccelerationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &webAppAccelerationId, WaaWebAppAccelerationSweepWaitCondition, time.Duration(3*time.Minute),
				WaaWebAppAccelerationSweepResponseFetchOperation, "waa", true)
		}
	}
	return nil
}

func getWaaWebAppAccelerationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "WebAppAccelerationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	waaClient := acctest.GetTestClients(&schema.ResourceData{}).WaaClient()

	listWebAppAccelerationsRequest := oci_waa.ListWebAppAccelerationsRequest{}
	listWebAppAccelerationsRequest.CompartmentId = &compartmentId
	listWebAppAccelerationsRequest.LifecycleState = []oci_waa.WebAppAccelerationLifecycleStateEnum{oci_waa.WebAppAccelerationLifecycleStateActive}
	listWebAppAccelerationsResponse, err := waaClient.ListWebAppAccelerations(context.Background(), listWebAppAccelerationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting WebAppAcceleration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, webAppAcceleration := range listWebAppAccelerationsResponse.Items {
		id := *webAppAcceleration.GetId()
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "WebAppAccelerationId", id)
	}
	return resourceIds, nil
}

func WaaWebAppAccelerationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if webAppAccelerationResponse, ok := response.Response.(oci_waa.GetWebAppAccelerationResponse); ok {
		return webAppAccelerationResponse.GetLifecycleState() != oci_waa.WebAppAccelerationLifecycleStateDeleted
	}
	return false
}

func WaaWebAppAccelerationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.WaaClient().GetWebAppAcceleration(context.Background(), oci_waa.GetWebAppAccelerationRequest{
		WebAppAccelerationId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
