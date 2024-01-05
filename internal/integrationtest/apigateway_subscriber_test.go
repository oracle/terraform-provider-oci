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
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApigatewaySubscriberRequiredOnlyResource = ApigatewaySubscriberResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Required, acctest.Create, ApigatewaySubscriberRepresentation)

	ApigatewaySubscriberResourceConfig = ApigatewaySubscriberResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Update, ApigatewaySubscriberRepresentation)

	ApigatewaySubscriberSingularDataSourceRepresentation = map[string]interface{}{
		"subscriber_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_apigateway_subscriber.test_subscriber.id}`},
	}

	ApigatewaySubscriberDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		// "usage_plan_id":  acctest.Representation{RepType: acctest.Optional, Create: `${oci_apigateway_usage_plan.test_usage_plan.id}`},
		"filter": acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewaySubscriberDataSourceFilterRepresentation}}
	ApigatewaySubscriberDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_subscriber.test_subscriber.id}`}},
	}

	ApigatewaySubscriberRepresentation = map[string]interface{}{
		"clients":        acctest.RepresentationGroup{RepType: acctest.Required, Group: ApigatewaySubscriberClientsRepresentation},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"usage_plans":    acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_apigateway_usage_plan.test_usage_plan.id}`}, Update: []string{`${oci_apigateway_usage_plan.test_usage_plan2.id}`}},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreChangesSubscriberRepresentation},
	}
	ignoreChangesSubscriberRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}
	ApigatewaySubscriberClientsRepresentation = map[string]interface{}{
		"name":  acctest.Representation{RepType: acctest.Required, Create: `name`, Update: `name2`},
		"token": acctest.Representation{RepType: acctest.Required, Create: `token`, Update: `token2`},
	}

	// ApigatewaySubscriberResourceDependencies = DefinedTagsDependencies
	ApigatewaySubscriberResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan", acctest.Required, acctest.Create, ApigatewayUsagePlanRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_usage_plan", "test_usage_plan2", acctest.Required, acctest.Create, ApigatewayUsagePlanRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: apigateway/default
func TestApigatewaySubscriberResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApigatewaySubscriberResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_apigateway_subscriber.test_subscriber"
	datasourceName := "data.oci_apigateway_subscribers.test_subscribers"
	singularDatasourceName := "data.oci_apigateway_subscriber.test_subscriber"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ApigatewaySubscriberResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Create, ApigatewaySubscriberRepresentation), "apigateway", "subscriber", t)

	acctest.ResourceTest(t, testAccCheckApigatewaySubscriberDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ApigatewaySubscriberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Required, acctest.Create, ApigatewaySubscriberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "clients.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.token", "token"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "usage_plans.#", "1"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ApigatewaySubscriberResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ApigatewaySubscriberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Create, ApigatewaySubscriberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "clients.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.token", "token"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "usage_plans.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ApigatewaySubscriberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(ApigatewaySubscriberRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "clients.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.name", "name"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.token", "token"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "usage_plans.#", "1"),

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
			Config: config + compartmentIdVariableStr + ApigatewaySubscriberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Update, ApigatewaySubscriberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "clients.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.name", "name2"),
				resource.TestCheckResourceAttr(resourceName, "clients.0.token", "token2"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "usage_plans.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_subscribers", "test_subscribers", acctest.Optional, acctest.Update, ApigatewaySubscriberDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewaySubscriberResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Optional, acctest.Update, ApigatewaySubscriberRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				// resource.TestCheckResourceAttrSet(datasourceName, "usage_plan_id"),

				resource.TestCheckResourceAttr(datasourceName, "subscriber_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscriber_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apigateway_subscriber", "test_subscriber", acctest.Required, acctest.Create, ApigatewaySubscriberSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ApigatewaySubscriberResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscriber_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "clients.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "clients.0.name", "name2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "clients.0.token", "token2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "usage_plans.#", "1"),
			),
		},
		// verify resource import
		{
			Config:            config + ApigatewaySubscriberRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"lifecycle_details",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckApigatewaySubscriberDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).SubscribersClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_apigateway_subscriber" {
			noResourceFound = false
			request := oci_apigateway.GetSubscriberRequest{}

			tmp := rs.Primary.ID
			request.SubscriberId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")

			response, err := client.GetSubscriber(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_apigateway.SubscriberLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("ApigatewaySubscriber") {
		resource.AddTestSweepers("ApigatewaySubscriber", &resource.Sweeper{
			Name:         "ApigatewaySubscriber",
			Dependencies: acctest.DependencyGraph["subscriber"],
			F:            sweepApigatewaySubscriberResource,
		})
	}
}

func sweepApigatewaySubscriberResource(compartment string) error {
	subscribersClient := acctest.GetTestClients(&schema.ResourceData{}).SubscribersClient()
	subscriberIds, err := getSubscriberIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriberId := range subscriberIds {
		if ok := acctest.SweeperDefaultResourceId[subscriberId]; !ok {
			deleteSubscriberRequest := oci_apigateway.DeleteSubscriberRequest{}

			deleteSubscriberRequest.SubscriberId = &subscriberId

			deleteSubscriberRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "apigateway")
			_, error := subscribersClient.DeleteSubscriber(context.Background(), deleteSubscriberRequest)
			if error != nil {
				fmt.Printf("Error deleting Subscriber %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriberId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subscriberId, subscriberSweepWaitCondition, time.Duration(3*time.Minute),
				subscriberSweepResponseFetchOperation, "apigateway", true)
		}
	}
	return nil
}

func getSubscriberIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriberId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	subscribersClient := acctest.GetTestClients(&schema.ResourceData{}).SubscribersClient()

	listSubscribersRequest := oci_apigateway.ListSubscribersRequest{}
	listSubscribersRequest.CompartmentId = &compartmentId
	listSubscribersRequest.LifecycleState = oci_apigateway.SubscriberLifecycleStateActive
	listSubscribersResponse, err := subscribersClient.ListSubscribers(context.Background(), listSubscribersRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subscriber list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subscriber := range listSubscribersResponse.Items {
		id := *subscriber.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriberId", id)
	}
	return resourceIds, nil
}

func subscriberSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subscriberResponse, ok := response.Response.(oci_apigateway.GetSubscriberResponse); ok {
		return subscriberResponse.LifecycleState != oci_apigateway.SubscriberLifecycleStateDeleted
	}
	return false
}

func subscriberSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.SubscribersClient().GetSubscriber(context.Background(), oci_apigateway.GetSubscriberRequest{
		SubscriberId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
