// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/oracle/oci-go-sdk/common"
	oci_ons "github.com/oracle/oci-go-sdk/ons"
)

var (
	topicName                             = `t` + randomString(10, charset)
	NotificationTopicRequiredOnlyResource = NotificationTopicResourceDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation)

	NotificationTopicResourceConfig = NotificationTopicResourceDependencies +
		generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Optional, Update, notificationTopicRepresentation)

	notificationTopicSingularDataSourceRepresentation = map[string]interface{}{
		"topic_id": Representation{repType: Required, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
	}

	notificationTopicDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_ons_notification_topic.test_notification_topic.id}`},
		"name":           Representation{repType: Optional, create: topicName},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, notificationTopicDataSourceFilterRepresentation}}
	notificationTopicDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `topic_id`},
		"values": Representation{repType: Required, create: []string{`${oci_ons_notification_topic.test_notification_topic.id}`}},
	}

	notificationTopicRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"name":           Representation{repType: Required, create: topicName},
		"defined_tags":   Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    Representation{repType: Optional, create: `Channel for admin messages`, update: `description2`},
		"freeform_tags":  Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
	}

	NotificationTopicResourceDependencies = DefinedTagsDependencies
)

func getTopicRepresentationCopyWithRandomName() map[string]interface{} {
	return representationCopyWithNewProperties(notificationTopicRepresentation, map[string]interface{}{
		"name": Representation{repType: Required, create: randomString(10, charset)},
	})
}

func TestOnsNotificationTopicResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ons_notification_topic.test_notification_topic"
	datasourceName := "data.oci_ons_notification_topics.test_notification_topics"
	singularDatasourceName := "data.oci_ons_notification_topic.test_notification_topic"

	var resId, resId2 string

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckOnsNotificationTopicDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + NotificationTopicResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", topicName),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + NotificationTopicResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + NotificationTopicResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Optional, Create, notificationTopicRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "api_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "Channel for admin messages"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", topicName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + NotificationTopicResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Optional, Update, notificationTopicRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "api_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "name", topicName),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),
					resource.TestCheckResourceAttrSet(resourceName, "topic_id"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
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
					generateDataSourceFromRepresentationMap("oci_ons_notification_topics", "test_notification_topics", Optional, Update, notificationTopicDataSourceRepresentation) +
					compartmentIdVariableStr + NotificationTopicResourceDependencies +
					generateResourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Optional, Update, notificationTopicRepresentation),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "name", topicName),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "notification_topics.#", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "notification_topics.0.api_endpoint"),
					resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "notification_topics.0.defined_tags.%", "1"),
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
					generateDataSourceFromRepresentationMap("oci_ons_notification_topic", "test_notification_topic", Required, Create, notificationTopicSingularDataSourceRepresentation) +
					compartmentIdVariableStr + NotificationTopicResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttrSet(singularDatasourceName, "api_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", topicName),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "topic_id"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + NotificationTopicResourceConfig,
			},
			// verify resource import
			{
				Config:                  config,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ResourceName:            resourceName,
			},
		},
	})
}

func testAccCheckOnsNotificationTopicDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).notificationControlPlaneClient
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ons_notification_topic" {
			noResourceFound = false
			request := oci_ons.GetTopicRequest{}

			if value, ok := rs.Primary.Attributes["topic_id"]; ok {
				request.TopicId = &value
			}

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ons")

			_, err := client.GetTopic(context.Background(), request)

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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	resource.AddTestSweepers("OnsTopic", &resource.Sweeper{
		Name:         "OnsTopic",
		Dependencies: DependencyGraph["topic"],
		F:            sweepOnsTopicResource,
	})
}

func sweepOnsTopicResource(compartment string) error {
	notificationControlPlaneClient := GetTestClients(&schema.ResourceData{}).notificationControlPlaneClient
	topicIds, err := getTopicIds(compartment)
	if err != nil {
		return err
	}
	for _, topicId := range topicIds {
		if ok := SweeperDefaultResourceId[topicId]; !ok {
			deleteTopicRequest := oci_ons.DeleteTopicRequest{}

			deleteTopicRequest.TopicId = &topicId

			deleteTopicRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "ons")
			_, error := notificationControlPlaneClient.DeleteTopic(context.Background(), deleteTopicRequest)
			if error != nil {
				fmt.Printf("Error deleting Topic %s %s, It is possible that the resource is already deleted. Please verify manually \n", topicId, error)
				continue
			}
			waitTillCondition(testAccProvider, &topicId, topicSweepWaitCondition, time.Duration(3*time.Minute),
				topicSweepResponseFetchOperation, "ons", true)
		}
	}
	return nil
}

func getTopicIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "TopicId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	notificationControlPlaneClient := GetTestClients(&schema.ResourceData{}).notificationControlPlaneClient

	listTopicsRequest := oci_ons.ListTopicsRequest{}
	listTopicsRequest.CompartmentId = &compartmentId
	listTopicsResponse, err := notificationControlPlaneClient.ListTopics(context.Background(), listTopicsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Topic list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, topic := range listTopicsResponse.Items {
		if topic.LifecycleState != oci_ons.NotificationTopicSummaryLifecycleStateDeleting {
			id := *topic.TopicId
			resourceIds = append(resourceIds, id)
			addResourceIdToSweeperResourceIdMap(compartmentId, "TopicId", id)
		}
	}
	return resourceIds, nil
}

func topicSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond defined mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if topicResponse, ok := response.Response.(oci_ons.GetTopicResponse); ok {
		return topicResponse.LifecycleState == oci_ons.NotificationTopicLifecycleStateDeleting
	}
	return false
}

func topicSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.notificationControlPlaneClient.GetTopic(context.Background(), oci_ons.GetTopicRequest{
		TopicId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
