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
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_ons "github.com/oracle/oci-go-sdk/v58/ons"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	SubscriptionRequiredOnlyResource = SubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Required, acctest.Create, subscriptionRepresentation)

	SubscriptionResourceConfig = SubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Update, subscriptionRepresentation)

	subscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_subscription.test_subscription.id}`},
	}

	subscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"topic_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: subscriptionDataSourceFilterRepresentation}}
	subscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ons_subscription.test_subscription.id}`}},
	}

	subscriptionRepresentation = map[string]interface{}{
		"compartment_id":  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"endpoint":        acctest.Representation{RepType: acctest.Required, Create: `john.smith@example.com`},
		"protocol":        acctest.Representation{RepType: acctest.Required, Create: `EMAIL`},
		"topic_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"defined_tags":    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delivery_policy": acctest.Representation{RepType: acctest.Optional, Update: `{\"backoffRetryPolicy\":{\"maxRetryDuration\":7000000,\"policyType\":\"EXPONENTIAL\"}}`},
	}

	SubscriptionResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(10, utils.CharsetWithoutDigits, "tsubscription"))
)

// issue-routing-tag: ons/default
func TestOnsSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnsSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ons_subscription.test_subscription"
	datasourceName := "data.oci_ons_subscriptions.test_subscriptions"
	singularDatasourceName := "data.oci_ons_subscription.test_subscription"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+SubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Create, subscriptionRepresentation), "ons", "subscription", t)

	acctest.ResourceTest(t, testAccCheckOnsSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Required, acctest.Create, subscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
				resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Create, subscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + SubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(subscriptionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

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
			Config: config + compartmentIdVariableStr + SubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Update, subscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "protocol", "EMAIL"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ons_subscriptions", "test_subscriptions", acctest.Optional, acctest.Update, subscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + SubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Optional, acctest.Update, subscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "topic_id"),

				resource.TestCheckResourceAttr(datasourceName, "subscriptions.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.created_time"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.delivery_policy.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.etag"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.id"),
				resource.TestCheckResourceAttr(datasourceName, "subscriptions.0.protocol", "EMAIL"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "subscriptions.0.topic_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ons_subscription", "test_subscription", acctest.Required, acctest.Create, subscriptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + SubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "subscription_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_time"),
				resource.TestCheckResourceAttr(singularDatasourceName, "endpoint", "john.smith@example.com"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "etag"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "protocol", "EMAIL"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + SubscriptionResourceConfig,
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

func testAccCheckOnsSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NotificationDataPlaneClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ons_subscription" {
			noResourceFound = false
			request := oci_ons.GetSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.SubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ons")

			response, err := client.GetSubscription(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ons.SubscriptionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("OnsSubscription") {
		resource.AddTestSweepers("OnsSubscription", &resource.Sweeper{
			Name:         "OnsSubscription",
			Dependencies: acctest.DependencyGraph["subscription"],
			F:            sweepOnsSubscriptionResource,
		})
	}
}

func sweepOnsSubscriptionResource(compartment string) error {
	notificationDataPlaneClient := acctest.GetTestClients(&schema.ResourceData{}).NotificationDataPlaneClient()
	subscriptionIds, err := getSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, subscriptionId := range subscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[subscriptionId]; !ok {
			deleteSubscriptionRequest := oci_ons.DeleteSubscriptionRequest{}

			deleteSubscriptionRequest.SubscriptionId = &subscriptionId

			deleteSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ons")
			_, error := notificationDataPlaneClient.DeleteSubscription(context.Background(), deleteSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting Subscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", subscriptionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &subscriptionId, subscriptionSweepWaitCondition, time.Duration(3*time.Minute),
				subscriptionSweepResponseFetchOperation, "ons", true)
		}
	}
	return nil
}

func getSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	notificationDataPlaneClient := acctest.GetTestClients(&schema.ResourceData{}).NotificationDataPlaneClient()

	listSubscriptionsRequest := oci_ons.ListSubscriptionsRequest{}
	listSubscriptionsRequest.CompartmentId = &compartmentId
	listSubscriptionsResponse, err := notificationDataPlaneClient.ListSubscriptions(context.Background(), listSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Subscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, subscription := range listSubscriptionsResponse.Items {
		id := *subscription.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SubscriptionId", id)
	}
	return resourceIds, nil
}

func subscriptionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if subscriptionResponse, ok := response.Response.(oci_ons.GetSubscriptionResponse); ok {
		return subscriptionResponse.LifecycleState != oci_ons.SubscriptionLifecycleStateDeleted
	}
	return false
}

func subscriptionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NotificationDataPlaneClient().GetSubscription(context.Background(), oci_ons.GetSubscriptionRequest{
		SubscriptionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
