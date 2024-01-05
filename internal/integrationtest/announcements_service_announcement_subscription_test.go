// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_announcements_service "github.com/oracle/oci-go-sdk/v65/announcementsservice"
	"github.com/oracle/oci-go-sdk/v65/common"
)

var (
	AnnouncementsServiceAnnouncementSubscriptionRequiredOnlyResource = AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation)

	AnnouncementsServiceAnnouncementSubscriptionResourceConfig = AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Update, AnnouncementsServiceAnnouncementSubscriptionRepresentation)

	AnnouncementsServiceAnnouncementsServiceAnnouncementSubscriptionSingularDataSourceRepresentation = map[string]interface{}{
		"announcement_subscription_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_announcements_service_announcement_subscription.test_announcement_subscription.id}`},
	}

	AnnouncementsServiceAnnouncementsServiceAnnouncementSubscriptionDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_announcements_service_announcement_subscription.test_announcement_subscription.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AnnouncementsServiceAnnouncementSubscriptionDataSourceFilterRepresentation}}
	AnnouncementsServiceAnnouncementSubscriptionDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_announcements_service_announcement_subscription.test_announcement_subscription.id}`}},
	}

	AnnouncementsServiceAnnouncementSubscriptionRepresentation = map[string]interface{}{
		"compartment_id":      acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":        acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"ons_topic_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"defined_tags":        acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":         acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"filter_groups":       acctest.RepresentationGroup{RepType: acctest.Optional, Group: AnnouncementsServiceAnnouncementSubscriptionFilterGroupsRepresentation},
		"freeform_tags":       acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":           acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsPlanChangesRepresentation},
		"preferred_language":  acctest.Representation{RepType: acctest.Optional, Create: `en-US`, Update: `sh-RS`},
		"preferred_time_zone": acctest.Representation{RepType: acctest.Optional, Create: `America/Mexico_City`, Update: `America/Los_Angeles`},
	}
	AnnouncementsServiceAnnouncementSubscriptionFilterGroupsRepresentation = map[string]interface{}{
		"filters": acctest.RepresentationGroup{RepType: acctest.Required, Group: AnnouncementsServiceAnnouncementSubscriptionFilterGroupsFiltersRepresentation},
	}
	AnnouncementsServiceAnnouncementSubscriptionFilterGroupsFiltersRepresentation = map[string]interface{}{
		"type":  acctest.Representation{RepType: acctest.Required, Create: `SERVICE`},
		"value": acctest.Representation{RepType: acctest.Required, Create: `Oracle Fusion Applications`},
	}

	ignoreDefinedTagsPlanChangesRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	AnnouncementsServiceAnnouncementSubscriptionResourceDependencies = DefinedTagsDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsNotificationTopicRepresentation)
)

// issue-routing-tag: announcements_service/default
func TestAnnouncementsServiceAnnouncementSubscriptionResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAnnouncementsServiceAnnouncementSubscriptionResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_announcements_service_announcement_subscription.test_announcement_subscription"
	datasourceName := "data.oci_announcements_service_announcement_subscriptions.test_announcement_subscriptions"
	singularDatasourceName := "data.oci_announcements_service_announcement_subscription.test_announcement_subscription"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnnouncementsServiceAnnouncementSubscriptionResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation), "announcementsservice", "announcementSubscription", t)

	acctest.ResourceTest(t, testAccCheckAnnouncementsServiceAnnouncementSubscriptionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Create, AnnouncementsServiceAnnouncementSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.type", "SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.value", "Oracle Fusion Applications"),
				resource.TestCheckResourceAttrSet(resourceName, "filter_groups.0.name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "preferred_time_zone", "America/Mexico_City"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AnnouncementsServiceAnnouncementSubscriptionRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.type", "SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.value", "Oracle Fusion Applications"),
				resource.TestCheckResourceAttrSet(resourceName, "filter_groups.0.name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "en-US"),
				resource.TestCheckResourceAttr(resourceName, "preferred_time_zone", "America/Mexico_City"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
			Config: config + compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Update, AnnouncementsServiceAnnouncementSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.type", "SERVICE"),
				resource.TestCheckResourceAttr(resourceName, "filter_groups.0.filters.0.value", "Oracle Fusion Applications"),
				resource.TestCheckResourceAttrSet(resourceName, "filter_groups.0.name"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "ons_topic_id"),
				resource.TestCheckResourceAttr(resourceName, "preferred_language", "sh-RS"),
				resource.TestCheckResourceAttr(resourceName, "preferred_time_zone", "America/Los_Angeles"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_announcements_service_announcement_subscriptions", "test_announcement_subscriptions", acctest.Optional, acctest.Update, AnnouncementsServiceAnnouncementsServiceAnnouncementSubscriptionDataSourceRepresentation) +
				compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Optional, acctest.Update, AnnouncementsServiceAnnouncementSubscriptionRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "announcement_subscription_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "announcement_subscription_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_announcements_service_announcement_subscription", "test_announcement_subscription", acctest.Required, acctest.Create, AnnouncementsServiceAnnouncementsServiceAnnouncementSubscriptionSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AnnouncementsServiceAnnouncementSubscriptionResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "announcement_subscription_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_groups.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_groups.0.filters.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_groups.0.filters.0.type", "SERVICE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "filter_groups.0.filters.0.value", "Oracle Fusion Applications"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "filter_groups.0.name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "preferred_language", "sh-RS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "preferred_time_zone", "America/Los_Angeles"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AnnouncementsServiceAnnouncementSubscriptionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAnnouncementsServiceAnnouncementSubscriptionDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnnouncementSubscriptionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_announcements_service_announcement_subscription" {
			noResourceFound = false
			request := oci_announcements_service.GetAnnouncementSubscriptionRequest{}

			tmp := rs.Primary.ID
			request.AnnouncementSubscriptionId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "announcements_service")

			response, err := client.GetAnnouncementSubscription(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_announcements_service.AnnouncementSubscriptionLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AnnouncementsServiceAnnouncementSubscription") {
		resource.AddTestSweepers("AnnouncementsServiceAnnouncementSubscription", &resource.Sweeper{
			Name:         "AnnouncementsServiceAnnouncementSubscription",
			Dependencies: acctest.DependencyGraph["announcementSubscription"],
			F:            sweepAnnouncementsServiceAnnouncementSubscriptionResource,
		})
	}
}

func sweepAnnouncementsServiceAnnouncementSubscriptionResource(compartment string) error {
	announcementSubscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).AnnouncementSubscriptionClient()
	announcementSubscriptionIds, err := getAnnouncementsServiceAnnouncementSubscriptionIds(compartment)
	if err != nil {
		return err
	}
	for _, announcementSubscriptionId := range announcementSubscriptionIds {
		if ok := acctest.SweeperDefaultResourceId[announcementSubscriptionId]; !ok {
			deleteAnnouncementSubscriptionRequest := oci_announcements_service.DeleteAnnouncementSubscriptionRequest{}

			deleteAnnouncementSubscriptionRequest.AnnouncementSubscriptionId = &announcementSubscriptionId

			deleteAnnouncementSubscriptionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "announcements_service")
			_, error := announcementSubscriptionClient.DeleteAnnouncementSubscription(context.Background(), deleteAnnouncementSubscriptionRequest)
			if error != nil {
				fmt.Printf("Error deleting AnnouncementSubscription %s %s, It is possible that the resource is already deleted. Please verify manually \n", announcementSubscriptionId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &announcementSubscriptionId, AnnouncementsServiceAnnouncementSubscriptionSweepWaitCondition, time.Duration(3*time.Minute),
				AnnouncementsServiceAnnouncementSubscriptionSweepResponseFetchOperation, "announcements_service", true)
		}
	}
	return nil
}

func getAnnouncementsServiceAnnouncementSubscriptionIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "AnnouncementSubscriptionId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	announcementSubscriptionClient := acctest.GetTestClients(&schema.ResourceData{}).AnnouncementSubscriptionClient()

	listAnnouncementSubscriptionsRequest := oci_announcements_service.ListAnnouncementSubscriptionsRequest{}
	listAnnouncementSubscriptionsRequest.CompartmentId = &compartmentId
	listAnnouncementSubscriptionsRequest.LifecycleState = oci_announcements_service.AnnouncementSubscriptionLifecycleStateActive
	listAnnouncementSubscriptionsResponse, err := announcementSubscriptionClient.ListAnnouncementSubscriptions(context.Background(), listAnnouncementSubscriptionsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting AnnouncementSubscription list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, announcementSubscription := range listAnnouncementSubscriptionsResponse.Items {
		id := *announcementSubscription.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "AnnouncementSubscriptionId", id)
	}
	return resourceIds, nil
}

func AnnouncementsServiceAnnouncementSubscriptionSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if announcementSubscriptionResponse, ok := response.Response.(oci_announcements_service.GetAnnouncementSubscriptionResponse); ok {
		return announcementSubscriptionResponse.LifecycleState != oci_announcements_service.AnnouncementSubscriptionLifecycleStateDeleted
	}
	return false
}

func AnnouncementsServiceAnnouncementSubscriptionSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnnouncementSubscriptionClient().GetAnnouncementSubscription(context.Background(), oci_announcements_service.GetAnnouncementSubscriptionRequest{
		AnnouncementSubscriptionId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
