// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	QueueConsumerGroupRequiredOnlyResource = QueueConsumerGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Required, acctest.Create, QueueConsumerGroupRepresentation)

	QueueConsumerGroupResourceConfig = QueueConsumerGroupResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Optional, acctest.Update, QueueConsumerGroupRepresentation)

	QueueConsumerGroupSingularDataSourceRepresentation = map[string]interface{}{
		"consumer_group_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_queue_consumer_group.test_consumer_group.id}`},
	}

	QueueConsumerGroupDataSourceRepresentation = map[string]interface{}{
		"display_name": acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":           acctest.Representation{RepType: acctest.Optional, Create: `${oci_queue_consumer_group.test_consumer_group.id}`},
		"queue_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_queue_queue.test_queue.id}`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":       acctest.RepresentationGroup{RepType: acctest.Required, Group: QueueConsumerGroupDataSourceFilterRepresentation}}
	QueueConsumerGroupDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_queue_consumer_group.test_consumer_group.id}`}},
	}

	QueueConsumerGroupRepresentation = map[string]interface{}{
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"queue_id":                         acctest.Representation{RepType: acctest.Required, Create: `${oci_queue_queue.test_queue.id}`},
		"consumer_group_filter":            acctest.Representation{RepType: acctest.Optional, Create: `:consumerGroupFilter`, Update: `:consumerGroupFilter2`},
		"dead_letter_queue_delivery_count": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_enabled":                       acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsRepresentation},
	}

	QueueConsumerGroupResourceDependencies = DefinedTagsDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Create, QueueQueueRepresentation)
)

// issue-routing-tag: queue/default
func TestQueueConsumerGroupResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestQueueConsumerGroupResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	customEncryptionKeyId := utils.GetEnvSettingWithBlankDefault("custom_encryption_key_id")
	customEncryptionKeyIdVariableStr := fmt.Sprintf("variable \"custom_encryption_key_id\" { default = \"%s\" }\n", customEncryptionKeyId)

	resourceName := "oci_queue_consumer_group.test_consumer_group"
	datasourceName := "data.oci_queue_consumer_groups.test_consumer_groups"
	singularDatasourceName := "data.oci_queue_consumer_group.test_consumer_group"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+QueueConsumerGroupResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Optional, acctest.Create, QueueConsumerGroupRepresentation), "queue", "consumerGroup", t)

	acctest.ResourceTest(t, testAccCheckQueueConsumerGroupDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + QueueConsumerGroupResourceDependencies + customEncryptionKeyIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Required, acctest.Create, QueueConsumerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "queue_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueConsumerGroupResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueConsumerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Optional, acctest.Create, QueueConsumerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "consumer_group_filter", ":consumerGroupFilter"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "queue_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueConsumerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Optional, acctest.Update, QueueConsumerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "consumer_group_filter", ":consumerGroupFilter2"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_enabled", "true"),
				resource.TestCheckResourceAttrSet(resourceName, "queue_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_queue_consumer_groups", "test_consumer_groups", acctest.Optional, acctest.Update, QueueConsumerGroupDataSourceRepresentation) +
				compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueConsumerGroupResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Optional, acctest.Update, QueueConsumerGroupRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "queue_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "consumer_group_collection.#", "1"),
				// There seems to be a bug in the list call . Will update below once updated.
				//resource.TestCheckResourceAttr(datasourceName, "consumer_group_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_queue_consumer_group", "test_consumer_group", acctest.Required, acctest.Create, QueueConsumerGroupSingularDataSourceRepresentation) +
				compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueConsumerGroupResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "consumer_group_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "consumer_group_filter", ":consumerGroupFilter2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "dead_letter_queue_delivery_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:            config + QueueConsumerGroupRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"is_enabled",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckQueueConsumerGroupDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).QueueAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_queue_consumer_group" {
			noResourceFound = false
			request := oci_queue.GetConsumerGroupRequest{}

			tmp := rs.Primary.ID
			request.ConsumerGroupId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "queue")

			response, err := client.GetConsumerGroup(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_queue.ConsumerGroupLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("QueueConsumerGroup") {
		resource.AddTestSweepers("QueueConsumerGroup", &resource.Sweeper{
			Name:         "QueueConsumerGroup",
			Dependencies: acctest.DependencyGraph["consumerGroup"],
			F:            sweepQueueConsumerGroupResource,
		})
	}
}

func sweepQueueConsumerGroupResource(compartment string) error {
	queueAdminClient := acctest.GetTestClients(&schema.ResourceData{}).QueueAdminClient()
	consumerGroupIds, err := getQueueConsumerGroupIds(compartment)
	if err != nil {
		return err
	}
	for _, consumerGroupId := range consumerGroupIds {
		if ok := acctest.SweeperDefaultResourceId[consumerGroupId]; !ok {
			deleteConsumerGroupRequest := oci_queue.DeleteConsumerGroupRequest{}

			deleteConsumerGroupRequest.ConsumerGroupId = &consumerGroupId

			deleteConsumerGroupRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "queue")
			_, error := queueAdminClient.DeleteConsumerGroup(context.Background(), deleteConsumerGroupRequest)
			if error != nil {
				fmt.Printf("Error deleting ConsumerGroup %s %s, It is possible that the resource is already deleted. Please verify manually \n", consumerGroupId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &consumerGroupId, QueueConsumerGroupSweepWaitCondition, time.Duration(3*time.Minute),
				QueueConsumerGroupSweepResponseFetchOperation, "queue", true)
		}
	}
	return nil
}

func getQueueConsumerGroupIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ConsumerGroupId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	queueAdminClient := acctest.GetTestClients(&schema.ResourceData{}).QueueAdminClient()

	listConsumerGroupsRequest := oci_queue.ListConsumerGroupsRequest{}
	//listConsumerGroupsRequest.CompartmentId = &compartmentId
	listConsumerGroupsRequest.LifecycleState = oci_queue.ConsumerGroupLifecycleStateActive
	listConsumerGroupsResponse, err := queueAdminClient.ListConsumerGroups(context.Background(), listConsumerGroupsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ConsumerGroup list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, consumerGroup := range listConsumerGroupsResponse.Items {
		id := *consumerGroup.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ConsumerGroupId", id)
	}
	return resourceIds, nil
}

func QueueConsumerGroupSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if consumerGroupResponse, ok := response.Response.(oci_queue.GetConsumerGroupResponse); ok {
		return consumerGroupResponse.LifecycleState != oci_queue.ConsumerGroupLifecycleStateDeleted
	}
	return false
}

func QueueConsumerGroupSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.QueueAdminClient().GetConsumerGroup(context.Background(), oci_queue.GetConsumerGroupRequest{
		ConsumerGroupId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
