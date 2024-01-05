// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	tf_provider "github.com/oracle/terraform-provider-oci/internal/provider"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_ons "github.com/oracle/oci-go-sdk/v65/ons"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	topicNameRequiredOnly                 = `t` + utils.RandomString(10, utils.Charset)
	topicName                             = `t` + utils.RandomString(10, utils.Charset)
	NotificationTopicRequiredOnlyResource = OnsNotificationTopicResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(OnsNotificationTopicRepresentation, map[string]interface{}{
			"name": acctest.Representation{RepType: acctest.Required, Create: topicNameRequiredOnly},
		}))

	OnsNotificationTopicResourceConfig = OnsNotificationTopicResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Update, OnsNotificationTopicRepresentation)

	OnsOnsNotificationTopicSingularDataSourceRepresentation = map[string]interface{}{
		"topic_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	OnsOnsNotificationTopicDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"name":           acctest.Representation{RepType: acctest.Optional, Create: topicName},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: OnsNotificationTopicDataSourceFilterRepresentation}}
	OnsNotificationTopicDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `topic_id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ons_notification_topic.test_notification_topic.id}`}},
	}

	OnsNotificationTopicRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: topicName},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `Channel for admin messages`, Update: `description2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	OnsNotificationTopicResourceDependencies = DefinedTagsDependencies
)

func getTopicRepresentationCopyWithRandomNameOrHttpReplayValue(length int, charset string, httpReplayValue string) map[string]interface{} {
	return acctest.RepresentationCopyWithNewProperties(OnsNotificationTopicRepresentation, map[string]interface{}{
		"name": acctest.Representation{RepType: acctest.Required, Create: utils.RandomStringOrHttpReplayValue(length, utils.Charset, httpReplayValue)},
	})
}

// issue-routing-tag: ons/default
func TestOnsNotificationTopicResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestOnsNotificationTopicResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ons_notification_topic.test_notification_topic"
	datasourceName := "data.oci_ons_notification_topics.test_notification_topics"
	singularDatasourceName := "data.oci_ons_notification_topic.test_notification_topic"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+OnsNotificationTopicResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Create, OnsNotificationTopicRepresentation), "ons", "notificationTopic", t)

	acctest.ResourceTest(t, testAccCheckOnsNotificationTopicDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotificationTopicRequiredOnlyResource,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "name", topicNameRequiredOnly),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + OnsNotificationTopicResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + OnsNotificationTopicResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Create, OnsNotificationTopicRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "api_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "Channel for admin messages"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", topicName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + OnsNotificationTopicResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(OnsNotificationTopicRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "api_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "Channel for admin messages"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", topicName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
			Config: config + compartmentIdVariableStr + OnsNotificationTopicResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Update, OnsNotificationTopicRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "api_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "name", topicName),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ons_notification_topics", "test_notification_topics", acctest.Optional, acctest.Update, OnsOnsNotificationTopicDataSourceRepresentation) +
				compartmentIdVariableStr + OnsNotificationTopicResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Optional, acctest.Update, OnsNotificationTopicRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "name", topicName),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "notification_topics.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "notification_topics.0.api_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.description", "description2"),
				resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.name", topicName),
				resource.TestCheckResourceAttrSet(datasourceName, "notification_topics.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "notification_topics.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "notification_topics.0.topic_id"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", acctest.Required, acctest.Create, OnsOnsNotificationTopicSingularDataSourceRepresentation) +
				compartmentIdVariableStr + OnsNotificationTopicResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttrSet(singularDatasourceName, "api_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "name", topicName),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "topic_id"),
			),
		},
		// verify resource import
		{
			Config:                  config + NotificationTopicRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckOnsNotificationTopicDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).NotificationControlPlaneClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ons_notification_topic" {
			noResourceFound = false
			request := oci_ons.GetTopicRequest{}

			if value, ok := rs.Primary.Attributes["topic_id"]; ok {
				request.TopicId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ons")

			response, err := client.GetTopic(context.Background(), request)

			if tf_provider.AvoidWaitingForDeleteTarget && response.LifecycleState == oci_ons.NotificationTopicLifecycleStateDeleting {
				return nil
			}

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
	if !acctest.InSweeperExcludeList("OnsNotificationTopic") {
		resource.AddTestSweepers("OnsNotificationTopic", &resource.Sweeper{
			Name:         "OnsNotificationTopic",
			Dependencies: acctest.DependencyGraph["notificationTopic"],
			F:            sweepOnsNotificationTopicResource,
		})
	}
}

func sweepOnsNotificationTopicResource(compartment string) error {
	notificationControlPlaneClient := acctest.GetTestClients(&schema.ResourceData{}).NotificationControlPlaneClient()
	notificationTopicIds, err := getOnsNotificationTopicIds(compartment)
	if err != nil {
		return err
	}
	for _, notificationTopicId := range notificationTopicIds {
		if ok := acctest.SweeperDefaultResourceId[notificationTopicId]; !ok {
			deleteTopicRequest := oci_ons.DeleteTopicRequest{}

			deleteTopicRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ons")
			_, error := notificationControlPlaneClient.DeleteTopic(context.Background(), deleteTopicRequest)
			if error != nil {
				fmt.Printf("Error deleting NotificationTopic %s %s, It is possible that the resource is already deleted. Please verify manually \n", notificationTopicId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &notificationTopicId, topicSweepWaitCondition, time.Duration(3*time.Minute),
				topicSweepResponseFetchOperation, "ons", true)
		}
	}
	return nil
}

func getOnsNotificationTopicIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "NotificationTopicId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	notificationControlPlaneClient := acctest.GetTestClients(&schema.ResourceData{}).NotificationControlPlaneClient()

	listTopicsRequest := oci_ons.ListTopicsRequest{}
	listTopicsRequest.CompartmentId = &compartmentId
	listTopicsResponse, err := notificationControlPlaneClient.ListTopics(context.Background(), listTopicsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting NotificationTopic list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, notificationTopic := range listTopicsResponse.Items {
		if notificationTopic.LifecycleState != oci_ons.NotificationTopicSummaryLifecycleStateDeleting {
			id := *notificationTopic.TopicId
			resourceIds = append(resourceIds, id)
			acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "NotificationTopicId", id)
		}
	}
	return resourceIds, nil
}

func topicSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond defined mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if topicResponse, ok := response.Response.(oci_ons.GetTopicResponse); ok {
		return topicResponse.LifecycleState != oci_ons.NotificationTopicLifecycleStateDeleting
	}
	return false
}

func topicSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.NotificationControlPlaneClient().GetTopic(context.Background(), oci_ons.GetTopicRequest{
		TopicId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
