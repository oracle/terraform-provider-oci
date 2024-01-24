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
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	QueueQueueRequiredOnlyResource = QueueQueueResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation)

	QueueQueueResourceConfig = QueueQueueResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Update, QueueQueueRepresentation)

	QueueQueueQueueSingularDataSourceRepresentation = map[string]interface{}{
		"queue_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_queue_queue.test_queue.id}`},
	}

	ignoreDefinedTagsRepresentation = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`defined_tags`}},
	}

	QueueQueueQueueDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_queue_queue.test_queue.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: QueueQueueDataSourceFilterRepresentation}}
	QueueQueueDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_queue_queue.test_queue.id}`}},
	}

	QueueQueueRepresentation = map[string]interface{}{
		"compartment_id":                   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":                     acctest.Representation{RepType: acctest.Required, Create: `displayName`, Update: `displayName2`},
		"channel_consumption_limit":        acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"custom_encryption_key_id":         acctest.Representation{RepType: acctest.Optional, Create: `${var.custom_encryption_key_id}`},
		"dead_letter_queue_delivery_count": acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"defined_tags":                     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":                    acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"retention_in_seconds":             acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"timeout_in_seconds":               acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"visibility_in_seconds":            acctest.Representation{RepType: acctest.Optional, Create: `10`, Update: `11`},
		"lifecycle":                        acctest.RepresentationGroup{RepType: acctest.Optional, Group: ignoreDefinedTagsRepresentation},
	}

	QueueQueueResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: queue/default
func TestQueueQueueResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestQueueQueueResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	customEncryptionKeyId := utils.GetEnvSettingWithBlankDefault("custom_encryption_key_id")
	customEncryptionKeyIdVariableStr := fmt.Sprintf("variable \"custom_encryption_key_id\" { default = \"%s\" }\n", customEncryptionKeyId)

	resourceName := "oci_queue_queue.test_queue"
	datasourceName := "data.oci_queue_queues.test_queues"
	singularDatasourceName := "data.oci_queue_queue.test_queue"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+QueueQueueResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Create, QueueQueueRepresentation), "queue", "queue", t)

	acctest.ResourceTest(t, testAccCheckQueueQueueDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + QueueQueueResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Create, QueueQueueRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_consumption_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key_id"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "visibility_in_seconds", "10"),

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
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + compartmentIdUVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(QueueQueueRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_consumption_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key_id"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "visibility_in_seconds", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// Verify purge queue operation. This should trigger purge operation.
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + compartmentIdUVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(QueueQueueRepresentation, map[string]interface{}{
						"purge_queue": acctest.Representation{RepType: acctest.Required, Create: `true`},
						"purge_type":  acctest.Representation{RepType: acctest.Required, Create: `normal`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key_id"),
				resource.TestCheckResourceAttrSet(resourceName, "purge_queue"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "10"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "10"),
				resource.TestCheckResourceAttr(resourceName, "visibility_in_seconds", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters and we are not setting  the purge queue related
		// parameters. So it should not trigger purge queue.
		{
			Config: config + compartmentIdVariableStr + customEncryptionKeyIdVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Update, QueueQueueRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "channel_consumption_limit", "11"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "custom_encryption_key_id"),
				resource.TestCheckResourceAttr(resourceName, "dead_letter_queue_delivery_count", "11"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(resourceName, "retention_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttr(resourceName, "timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(resourceName, "visibility_in_seconds", "11"),

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
			Config: config + customEncryptionKeyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_queue_queues", "test_queues", acctest.Optional, acctest.Update, QueueQueueQueueDataSourceRepresentation) +
				compartmentIdVariableStr + QueueQueueResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Optional, acctest.Update, QueueQueueRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "queue_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "queue_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config + customEncryptionKeyIdVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_queue_queue", "test_queue", acctest.Required, acctest.Create, QueueQueueQueueSingularDataSourceRepresentation) +
				compartmentIdVariableStr + QueueQueueResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "queue_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "channel_consumption_limit", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "dead_letter_queue_delivery_count", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "messages_endpoint"),
				resource.TestCheckResourceAttr(singularDatasourceName, "retention_in_seconds", "10"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "timeout_in_seconds", "11"),
				resource.TestCheckResourceAttr(singularDatasourceName, "visibility_in_seconds", "11"),
			),
		},
		// verify resource import
		{
			Config:            config + QueueQueueRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"purge_queue",
				"purge_type",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckQueueQueueDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).QueueAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_queue_queue" {
			noResourceFound = false
			request := oci_queue.GetQueueRequest{}

			tmp := rs.Primary.ID
			request.QueueId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "queue")

			response, err := client.GetQueue(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_queue.QueueLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("QueueQueue") {
		resource.AddTestSweepers("QueueQueue", &resource.Sweeper{
			Name:         "QueueQueue",
			Dependencies: acctest.DependencyGraph["queue"],
			F:            sweepQueueQueueResource,
		})
	}
}

func sweepQueueQueueResource(compartment string) error {
	queueAdminClient := acctest.GetTestClients(&schema.ResourceData{}).QueueAdminClient()
	queueIds, err := getQueueQueueIds(compartment)
	if err != nil {
		return err
	}
	for _, queueId := range queueIds {
		if ok := acctest.SweeperDefaultResourceId[queueId]; !ok {
			deleteQueueRequest := oci_queue.DeleteQueueRequest{}

			deleteQueueRequest.QueueId = &queueId

			deleteQueueRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "queue")
			_, error := queueAdminClient.DeleteQueue(context.Background(), deleteQueueRequest)
			if error != nil {
				fmt.Printf("Error deleting Queue %s %s, It is possible that the resource is already deleted. Please verify manually \n", queueId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &queueId, QueueQueueSweepWaitCondition, time.Duration(3*time.Minute),
				QueueQueueSweepResponseFetchOperation, "queue", true)
		}
	}
	return nil
}

func getQueueQueueIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "QueueId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	queueAdminClient := acctest.GetTestClients(&schema.ResourceData{}).QueueAdminClient()

	listQueuesRequest := oci_queue.ListQueuesRequest{}
	listQueuesRequest.CompartmentId = &compartmentId
	listQueuesRequest.LifecycleState = oci_queue.QueueLifecycleStateActive
	listQueuesResponse, err := queueAdminClient.ListQueues(context.Background(), listQueuesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Queue list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, queue := range listQueuesResponse.Items {
		id := *queue.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "QueueId", id)
	}
	return resourceIds, nil
}

func QueueQueueSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if queueResponse, ok := response.Response.(oci_queue.GetQueueResponse); ok {
		return queueResponse.LifecycleState != oci_queue.QueueLifecycleStateDeleted
	}
	return false
}

func QueueQueueSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.QueueAdminClient().GetQueue(context.Background(), oci_queue.GetQueueRequest{
		QueueId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
