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
	AiVisionStreamJobRequiredOnlyResource = AiVisionStreamJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Required, acctest.Create, AiVisionStreamJobRepresentation)

	AiVisionStreamJobResourceConfig = AiVisionStreamJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Update, AiVisionStreamJobRepresentation)

	AiVisionStreamJobSingularDataSourceRepresentation = map[string]interface{}{
		"stream_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_stream_job.test_stream_job.id}`},
	}

	AiVisionStreamJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_vision_stream_job.test_stream_job.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamJobDataSourceFilterRepresentation}}
	AiVisionStreamJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_vision_stream_job.test_stream_job.id}`}},
	}

	AiVisionStreamJobRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"features":               acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamJobFeaturesRepresentation},
		"stream_output_location": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiVisionStreamJobStreamOutputLocationRepresentation},
		"stream_source_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_stream_source.test_stream_source.id}`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags2": "freeformTags2"}},
	}
	AiVisionStreamJobFeaturesRepresentation = map[string]interface{}{
		"feature_type": acctest.Representation{RepType: acctest.Required, Create: `FACE_DETECTION`, Update: `FACE_DETECTION`},
	}
	AiVisionStreamJobStreamOutputLocationRepresentation = map[string]interface{}{
		"bucket":               acctest.Representation{RepType: acctest.Required, Create: `Test`, Update: `Test`},
		"namespace":            acctest.Representation{RepType: acctest.Required, Create: `axfelw9p2fyr`, Update: `axfelw9p2fyr`},
		"output_location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`, Update: `OBJECT_STORAGE`},
		"prefix":               acctest.Representation{RepType: acctest.Required, Create: `prefix`, Update: `prefix2`},
		"obo_token":            acctest.Representation{RepType: acctest.Optional, Create: `oboToken`, Update: `oboToken2`},
	}
	AiVisionStreamJobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_vision_private_endpoint", "test_vision_private_endpoint", acctest.Required, acctest.Create, AiVisionVisionPrivateEndpointRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_source", "test_stream_source", acctest.Required, acctest.Create, AiVisionStreamSourceRepresentation)
)

// issue-routing-tag: ai_vision/default
func TestAiVisionStreamJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiVisionStreamJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	subnetId := utils.GetEnvSettingWithBlankDefault("subnet_id")
	subnetIdVariableStr := fmt.Sprintf("variable \"subnet_id\" { default = \"%s\" }\n", subnetId)

	resourceName := "oci_ai_vision_stream_job.test_stream_job"
	datasourceName := "data.oci_ai_vision_stream_jobs.test_stream_jobs"
	singularDatasourceName := "data.oci_ai_vision_stream_job.test_stream_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+subnetIdVariableStr+AiVisionStreamJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Create, AiVisionStreamJobRepresentation), "aivision", "streamJob", t)

	acctest.ResourceTest(t, testAccCheckAiVisionStreamJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Required, acctest.Create, AiVisionStreamJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "features.0.feature_type", "FACE_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "features.0.tracking_types.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.bucket", "Test"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.namespace", "axfelw9p2fyr"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.output_location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.prefix", "prefix"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Create, AiVisionStreamJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "features.0.feature_type", "FACE_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.bucket", "Test"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.namespace", "axfelw9p2fyr"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.obo_token", "oboToken"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.output_location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.prefix", "prefix"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "false")); isEnableExportCompartment {
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + compartmentIdUVariableStr + AiVisionStreamJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiVisionStreamJobRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "features.0.feature_type", "FACE_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.bucket", "Test"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.namespace", "axfelw9p2fyr"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.obo_token", "oboToken"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.output_location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.prefix", "prefix"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_id"),
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
			Config: config + compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Update, AiVisionStreamJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "features.0.feature_type", "FACE_DETECTION"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.bucket", "Test"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.namespace", "axfelw9p2fyr"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.obo_token", "oboToken2"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.output_location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "stream_output_location.0.prefix", "prefix2"),
				resource.TestCheckResourceAttrSet(resourceName, "stream_source_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_jobs", "test_stream_jobs", acctest.Optional, acctest.Update, AiVisionStreamJobDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Optional, acctest.Update, AiVisionStreamJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "stream_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "stream_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_stream_job", "test_stream_job", acctest.Required, acctest.Create, AiVisionStreamJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + subnetIdVariableStr + AiVisionStreamJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "stream_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "features.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "features.0.feature_type", "FACE_DETECTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "INACTIVE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.0.bucket", "Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.0.namespace", "axfelw9p2fyr"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.0.obo_token", "oboToken2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.0.output_location_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "stream_output_location.0.prefix", "prefix2"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiVisionStreamJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiVisionStreamJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceVisionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_vision_stream_job" {
			noResourceFound = false
			request := oci_ai_vision.GetStreamJobRequest{}

			tmp := rs.Primary.ID
			request.StreamJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")

			response, err := client.GetStreamJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_vision.StreamJobLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiVisionStreamJob") {
		resource.AddTestSweepers("AiVisionStreamJob", &resource.Sweeper{
			Name:         "AiVisionStreamJob",
			Dependencies: acctest.DependencyGraph["streamJob"],
			F:            sweepAiVisionStreamJobResource,
		})
	}
}

func sweepAiVisionStreamJobResource(compartment string) error {
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()
	streamJobIds, err := getAiVisionStreamJobIds(compartment)
	if err != nil {
		return err
	}
	for _, streamJobId := range streamJobIds {
		if ok := acctest.SweeperDefaultResourceId[streamJobId]; !ok {
			deleteStreamJobRequest := oci_ai_vision.DeleteStreamJobRequest{}

			deleteStreamJobRequest.StreamJobId = &streamJobId

			deleteStreamJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")
			_, error := aiServiceVisionClient.DeleteStreamJob(context.Background(), deleteStreamJobRequest)
			if error != nil {
				fmt.Printf("Error deleting StreamJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", streamJobId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &streamJobId, AiVisionStreamJobSweepWaitCondition, time.Duration(3*time.Minute),
				AiVisionStreamJobSweepResponseFetchOperation, "ai_vision", true)
		}
	}
	return nil
}

func getAiVisionStreamJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "StreamJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()

	listStreamJobsRequest := oci_ai_vision.ListStreamJobsRequest{}
	listStreamJobsRequest.CompartmentId = &compartmentId
	listStreamJobsRequest.LifecycleState = oci_ai_vision.StreamJobLifecycleStateNeedsAttention
	listStreamJobsResponse, err := aiServiceVisionClient.ListStreamJobs(context.Background(), listStreamJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting StreamJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, streamJob := range listStreamJobsResponse.Items {
		id := *streamJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "StreamJobId", id)
	}
	return resourceIds, nil
}

func AiVisionStreamJobSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if streamJobResponse, ok := response.Response.(oci_ai_vision.GetStreamJobResponse); ok {
		return streamJobResponse.LifecycleState != oci_ai_vision.StreamJobLifecycleStateDeleted
	}
	return false
}

func AiVisionStreamJobSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceVisionClient().GetStreamJob(context.Background(), oci_ai_vision.GetStreamJobRequest{
		StreamJobId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
