// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v46/common"
	oci_streaming "github.com/oracle/oci-go-sdk/v46/streaming"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	StreamRequiredOnlyResource = StreamResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation)

	StreamResourceConfig = StreamResourceDependencies +
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Update, streamRepresentation)

	streamSingularDataSourceRepresentation = map[string]interface{}{
		"stream_id": Representation{repType: Required, create: `${oci_streaming_stream.test_stream.id}`},
	}

	streamDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"id":             Representation{repType: Optional, create: `${oci_streaming_stream.test_stream.id}`},
		"name":           Representation{repType: Optional, create: `mynewstream`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, streamDataSourceFilterRepresentation}}
	streamDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_streaming_stream.test_stream.id}`}},
	}

	streamRepresentation = map[string]interface{}{
		"compartment_id":     Representation{repType: Required, create: `${var.compartment_id}`},
		"name":               Representation{repType: Required, create: `mynewstream`},
		"partitions":         Representation{repType: Required, create: `1`},
		"defined_tags":       Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"freeform_tags":      Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"retention_in_hours": Representation{repType: Optional, create: `24`},
	}

	StreamResourceDependencies = DefinedTagsDependencies
)

// issue-routing-tag: streaming/default
func TestStreamingStreamResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestStreamingStreamResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_streaming_stream.test_stream"
	datasourceName := "data.oci_streaming_streams.test_streams"
	singularDatasourceName := "data.oci_streaming_stream.test_stream"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+StreamResourceDependencies+
		generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Create, streamRepresentation), "streaming", "stream", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckStreamingStreamDestroy,
		Steps: []resource.TestStep{
			// verify create
			{
				Config: config + compartmentIdVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(resourceName, "partitions", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// Verify that stream's compartment_id can be removed and stream_pool_id can be used
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream_pool", "test_stream_pool", Required, Create, representationCopyWithNewProperties(streamPoolRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})) +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Create, streampoolidRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(resourceName, "partitions", "1"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},

			// delete before next create
			{
				Config: config + compartmentIdVariableStr + StreamResourceDependencies,
			},
			// verify create with optionals
			{
				Config: config + compartmentIdVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Create, streamRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(resourceName, "partitions", "1"),
					resource.TestCheckResourceAttr(resourceName, "retention_in_hours", "24"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = fromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(getEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := testExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Create,
						representationCopyWithNewProperties(streamRepresentation, map[string]interface{}{
							"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
						})),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(resourceName, "partitions", "1"),
					resource.TestCheckResourceAttr(resourceName, "retention_in_hours", "24"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = fromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Update, streamRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "messages_endpoint"),
					resource.TestCheckResourceAttr(resourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(resourceName, "partitions", "1"),
					resource.TestCheckResourceAttr(resourceName, "retention_in_hours", "24"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

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
					generateDataSourceFromRepresentationMap("oci_streaming_streams", "test_streams", Optional, Update, streamDataSourceRepresentation) +
					compartmentIdVariableStr + StreamResourceDependencies +
					generateResourceFromRepresentationMap("oci_streaming_stream", "test_stream", Optional, Update, streamRepresentation),
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "id"),
					resource.TestCheckResourceAttr(datasourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

					resource.TestCheckResourceAttr(datasourceName, "streams.#", "1"),
					resource.TestCheckResourceAttr(datasourceName, "streams.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "streams.0.defined_tags.%", "1"),
					resource.TestCheckResourceAttr(datasourceName, "streams.0.freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "streams.0.id"),
					resource.TestCheckResourceAttrSet(datasourceName, "streams.0.messages_endpoint"),
					resource.TestCheckResourceAttr(datasourceName, "streams.0.name", "mynewstream"),
					resource.TestCheckResourceAttr(datasourceName, "streams.0.partitions", "1"),
					resource.TestCheckResourceAttrSet(datasourceName, "streams.0.state"),
					resource.TestCheckResourceAttrSet(datasourceName, "streams.0.time_created"),
				),
			},
			// verify singular datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_streaming_stream", "test_stream", Required, Create, streamSingularDataSourceRepresentation) +
					compartmentIdVariableStr + StreamResourceConfig,
				Check: ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "messages_endpoint"),
					resource.TestCheckResourceAttr(singularDatasourceName, "name", "mynewstream"),
					resource.TestCheckResourceAttr(singularDatasourceName, "partitions", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "retention_in_hours", "24"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// remove singular datasource from previous step so that it doesn't conflict with import tests
			{
				Config: config + compartmentIdVariableStr + StreamResourceConfig,
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

func testAccCheckStreamingStreamDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).streamAdminClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_streaming_stream" {
			noResourceFound = false
			request := oci_streaming.GetStreamRequest{}

			tmp := rs.Primary.ID
			request.StreamId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "streaming")

			response, err := client.GetStream(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_streaming.StreamLifecycleStateDeleted): true,
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
	if DependencyGraph == nil {
		initDependencyGraph()
	}
	if !inSweeperExcludeList("StreamingStream") {
		resource.AddTestSweepers("StreamingStream", &resource.Sweeper{
			Name:         "StreamingStream",
			Dependencies: DependencyGraph["stream"],
			F:            sweepStreamingStreamResource,
		})
	}
}

func sweepStreamingStreamResource(compartment string) error {
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()
	streamIds, err := getStreamIds(compartment)
	if err != nil {
		return err
	}
	for _, streamId := range streamIds {
		if ok := SweeperDefaultResourceId[streamId]; !ok {
			deleteStreamRequest := oci_streaming.DeleteStreamRequest{}

			deleteStreamRequest.StreamId = &streamId

			deleteStreamRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "streaming")
			_, error := streamAdminClient.DeleteStream(context.Background(), deleteStreamRequest)
			if error != nil {
				fmt.Printf("Error deleting Stream %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamId, error)
				continue
			}
			waitTillCondition(testAccProvider, &streamId, streamSweepWaitCondition, time.Duration(3*time.Minute),
				streamSweepResponseFetchOperation, "streaming", true)
		}
	}
	return nil
}

func getStreamIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "StreamId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	streamAdminClient := GetTestClients(&schema.ResourceData{}).streamAdminClient()

	listStreamsRequest := oci_streaming.ListStreamsRequest{}
	listStreamsRequest.CompartmentId = &compartmentId
	listStreamsRequest.LifecycleState = oci_streaming.StreamLifecycleStateActive
	listStreamsResponse, err := streamAdminClient.ListStreams(context.Background(), listStreamsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Stream list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, stream := range listStreamsResponse.Items {
		id := *stream.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "StreamId", id)
	}
	return resourceIds, nil
}

func streamSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamResponse, ok := response.Response.(oci_streaming.GetStreamResponse); ok {
		return streamResponse.LifecycleState != oci_streaming.StreamLifecycleStateDeleted
	}
	return false
}

func streamSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.streamAdminClient().GetStream(context.Background(), oci_streaming.GetStreamRequest{
		StreamId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
