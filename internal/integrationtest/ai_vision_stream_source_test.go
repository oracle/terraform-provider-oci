// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiVisionStreamSourceRequiredOnlyResource = AiVisionStreamSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Required, acctest.Create, AiVisionStreamSourceRepresentation)

	AiVisionStreamSourceResourceConfig = AiVisionStreamSourceResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Update, AiVisionStreamSourceRepresentation)

	AiVisionStreamSourceSingularDataSourceRepresentation = map[string]interface{}{
		"stream_source_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_stream_source.test_stream_source.id}`},
	}

	AiVisionStreamSourceDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_vision_stream_source.test_stream_source.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamSourceDataSourceFilterRepresentation}}
	AiVisionStreamSourceDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_vision_stream_source.test_stream_source.id}`}},
	}

	AiVisionStreamSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"stream_source_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamSourceStreamSourceDetailsRepresentation},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":         acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}
	AiVisionStreamSourceStreamSourceDetailsRepresentation = map[string]interface{}{
		"camera_url":                    acctest.Representation{RepType: acctest.Required, Create: `rtsp://64.1.1.23`, Update: `rtsp://64.1.1.23`},
		"source_type":                   acctest.Representation{RepType: acctest.Required, Create: `RTSP`},
		"stream_network_access_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamSourceStreamSourceDetailsStreamNetworkAccessDetailsRepresentation},
	}
	AiVisionStreamSourceStreamSourceDetailsStreamNetworkAccessDetailsRepresentation = map[string]interface{}{
		"private_endpoint_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_vision_private_endpoint.test_vision_private_endpoint.id}`},
		"stream_access_type":  acctest.Representation{RepType: acctest.Required, Create: `PRIVATE`},
	}

	AiVisionStreamSourceResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_vision_private_endpoint", "test_vision_private_endpoint", acctest.Required, acctest.Create, AiVisionVisionPrivateEndpointRepresentation)
)

// issue-routing-tag: ai_vision/default
func TestAiVisionStreamSourceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiVisionStreamSourceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_vision_stream_source.test_stream_source"
	datasourceName := "data.oci_ai_vision_stream_sources.test_stream_sources"
	singularDatasourceName := "data.oci_ai_vision_stream_source.test_stream_source"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiVisionStreamSourceResourceDependencies+subnetIdVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Create, AiVisionStreamSourceRepresentation), "aivision", "streamSource", t)

	acctest.ResourceTest(t, testAccCheckAiVisionStreamSourceDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Required, acctest.Create, AiVisionStreamSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.camera_url", "rtsp://64.1.1.23"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.source_type", "RTSP"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_details.0.stream_network_access_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.0.stream_access_type", "PRIVATE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Create, AiVisionStreamSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.camera_url", "rtsp://64.1.1.23"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.source_type", "RTSP"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_details.0.stream_network_access_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.0.stream_access_type", "PRIVATE"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiVisionStreamSourceRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.camera_url", "rtsp://64.1.1.23"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.source_type", "RTSP"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_details.0.stream_network_access_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.0.stream_access_type", "PRIVATE"),
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
			Config: config + compartmentIdVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Update, AiVisionStreamSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.camera_url", "rtsp://64.1.1.23"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.source_type", "RTSP"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_details.0.stream_network_access_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "stream_source_details.0.stream_network_access_details.0.stream_access_type", "PRIVATE"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_sources", "test_stream_sources", acctest.Optional, acctest.Update, AiVisionStreamSourceDataSourceRepresentation) +
				compartmentIdVariableStr + AiVisionStreamSourceResourceDependencies + subnetIdVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Optional, acctest.Update, AiVisionStreamSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "stream_source_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_source_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Required, acctest.Create, AiVisionStreamSourceSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamSourceResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_source_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_details.0.camera_url", "rtsp://64.1.1.23"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_details.0.source_type", "RTSP"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_details.0.stream_network_access_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_source_details.0.stream_network_access_details.0.stream_access_type", "PRIVATE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiVisionStreamSourceRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiVisionStreamSourceDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceVisionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_vision_stream_source" {
			noResourceFound = false
			request := oci_ai_vision.GetStreamSourceRequest{}

			tmp := rs.Primary.ID
			request.StreamSourceId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")

			response, err := client.GetStreamSource(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_vision.StreamSourceLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiVisionStreamSource") {
		resource.AddTestSweepers("AiVisionStreamSource", &resource.Sweeper{
			Name:         "AiVisionStreamSource",
			Dependencies: acctest.DependencyGraph["streamSource"],
			F:            sweepAiVisionStreamSourceResource,
		})
	}
}

func sweepAiVisionStreamSourceResource(compartment string) error {
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()
	streamSourceIds, err := getAiVisionStreamSourceIds(compartment)
	if err != nil {
		return err
	}
	for _, streamSourceId := range streamSourceIds {
		if ok := acctest.SweeperDefaultResourceId[streamSourceId]; !ok {
			deleteStreamSourceRequest := oci_ai_vision.DeleteStreamSourceRequest{}

			deleteStreamSourceRequest.StreamSourceId = &streamSourceId

			deleteStreamSourceRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")
			_, error := aiServiceVisionClient.DeleteStreamSource(context.Background(), deleteStreamSourceRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamSource %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamSourceId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamSourceId, AiVisionStreamSourceSweepWaitCondition, time.Duration(3*time.Minute),
				AiVisionStreamSourceSweepResponseFetchOperation, "ai_vision", true)
		}
	}
	return nil
}

func getAiVisionStreamSourceIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamSourceId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()

	listStreamSourcesRequest := oci_ai_vision.ListStreamSourcesRequest{}
	listStreamSourcesRequest.CompartmentId = &compartmentId
	listStreamSourcesRequest.LifecycleState = oci_ai_vision.StreamSourceLifecycleStateActive
	listStreamSourcesResponse, err := aiServiceVisionClient.ListStreamSources(context.Background(), listStreamSourcesRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting StreamSource list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, streamSource := range listStreamSourcesResponse.Items {
		id := *streamSource.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamSourceId", id)
	}
	return resourceIds, nil
}

func AiVisionStreamSourceSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamSourceResponse, ok := response.Response.(oci_ai_vision.GetStreamSourceResponse); ok {
		return streamSourceResponse.LifecycleState != oci_ai_vision.StreamSourceLifecycleStateDeleted
	}
	return false
}

func AiVisionStreamSourceSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceVisionClient().GetStreamSource(context.Background(), oci_ai_vision.GetStreamSourceRequest{
		StreamSourceId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
