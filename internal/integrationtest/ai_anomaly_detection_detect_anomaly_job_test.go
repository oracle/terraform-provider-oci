// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v65/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiAnomalyDetectionDetectAnomalyJobRequiredOnlyResource = AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Required, acctest.Create, AiAnomalyDetectionDetectAnomalyJobRepresentation)

	AiAnomalyDetectionDetectAnomalyJobResourceConfig = AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Update, AiAnomalyDetectionDetectAnomalyJobRepresentation)

	AiAnomalyDetectionDetectAnomalyJobSingularDataSourceRepresentation = map[string]interface{}{
		"detect_anomaly_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job.id}`},
	}

	AiAnomalyDetectionDetectAnomalyJobDataSourceRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"detect_anomaly_job_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job.id}`},
		"display_name":          acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"model_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_model.test_model.id}`},
		"project_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
		"filter":                acctest.RepresentationGroup{RepType: acctest.Required, Group: AiAnomalyDetectionDetectAnomalyJobDataSourceFilterRepresentation}}
	AiAnomalyDetectionDetectAnomalyJobDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job.id}`}},
	}

	AiAnomalyDetectionDetectAnomalyJobRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"input_details":  acctest.RepresentationGroup{RepType: acctest.Required, Group: AiAnomalyDetectionDetectAnomalyJobObjectListInputDetailsRepresentation},
		"model_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_model.test_model.id}`},
		"output_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiAnomalyDetectionDetectAnomalyJobOutputDetailsRepresentation},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"sensitivity":    acctest.Representation{RepType: acctest.Optional, Create: `0.5`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDetectAnomalyJobResultsChangesRep},
	}
	AiAnomalyDetectionDetectAnomalyJobObjectListInputDetailsRepresentation = map[string]interface{}{
		"input_type":       acctest.Representation{RepType: acctest.Required, Create: `OBJECT_LIST`},
		"object_locations": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiAnomalyDetectionDetectAnomalyJobInputDetailsObjectLocationsRepresentation},
	}
	AiAnomalyDetectionDetectAnomalyJobOutputDetailsRepresentation = map[string]interface{}{
		"bucket":      acctest.Representation{RepType: acctest.Required, Create: `mset-idp-test-datasets`},
		"namespace":   acctest.Representation{RepType: acctest.Required, Create: `id5zda5six9a`},
		"output_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"prefix":      acctest.Representation{RepType: acctest.Optional, Create: `ravi-async-go-output`},
	}
	AiAnomalyDetectionDetectAnomalyJobInputDetailsObjectLocationsRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `mset-idp-test-datasets`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `id5zda5six9a`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: `latest_test_data.csv`},
	}

	ignoreDetectAnomalyJobResultsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`output_details[0].prefix`, `defined_tags`, `system_tags`}},
	}

	AiAnomalyDetectionDetectAnomalyJobResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Create, aiAnomalyDetectionModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Required, acctest.Create, aiAnomalyDetectionDataAssetRepresentation)
)

// issue-routing-tag: ai_anomaly_detection/default
func TestAiAnomalyDetectionDetectAnomalyJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionDetectAnomalyJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job"
	datasourceName := "data.oci_ai_anomaly_detection_detect_anomaly_jobs.test_detect_anomaly_jobs"
	singularDatasourceName := "data.oci_ai_anomaly_detection_detect_anomaly_job.test_detect_anomaly_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionDetectAnomalyJobResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Create, AiAnomalyDetectionDetectAnomalyJobRepresentation), "aianomalydetection", "detectAnomalyJob", t)

	acctest.ResourceTest(t, testAccCheckAiAnomalyDetectionDetectAnomalyJobDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Required, acctest.Create, AiAnomalyDetectionDetectAnomalyJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "input_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.input_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.object", "latest_test_data.csv"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "output_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.output_type", "OBJECT_STORAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies,
		},
		// verify Create with optionals, trigger test
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Create, AiAnomalyDetectionDetectAnomalyJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.input_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.object", "latest_test_data.csv"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "output_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.output_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "output_details.0.prefix"),
				resource.TestCheckResourceAttr(resourceName, "sensitivity", `0.5`),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiAnomalyDetectionDetectAnomalyJobRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.input_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.object", "latest_test_data.csv"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "output_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.output_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "output_details.0.prefix"),
				resource.TestCheckResourceAttr(resourceName, "sensitivity", `0.5`),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Update, AiAnomalyDetectionDetectAnomalyJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.input_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "input_details.0.object_locations.0.object", "latest_test_data.csv"),
				resource.TestCheckResourceAttrSet(resourceName, "model_id"),
				resource.TestCheckResourceAttr(resourceName, "output_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(resourceName, "output_details.0.output_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "output_details.0.prefix"),
				resource.TestCheckResourceAttr(resourceName, "sensitivity", `0.5`),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_jobs", "test_detect_anomaly_jobs", acctest.Optional, acctest.Update, AiAnomalyDetectionDetectAnomalyJobDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Optional, acctest.Update, AiAnomalyDetectionDetectAnomalyJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "detect_anomaly_job_id"),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "detect_anomaly_job_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "detect_anomaly_job_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_detect_anomaly_job", "test_detect_anomaly_job", acctest.Required, acctest.Create, AiAnomalyDetectionDetectAnomalyJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionDetectAnomalyJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "detect_anomaly_job_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.0.input_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.0.object_locations.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.0.object_locations.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_details.0.object_locations.0.object", "latest_test_data.csv"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_details.0.bucket", "mset-idp-test-datasets"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_details.0.namespace", "id5zda5six9a"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_details.0.output_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "output_details.0.prefix"),
				resource.TestCheckResourceAttr(singularDatasourceName, "sensitivity", `0.5`),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiAnomalyDetectionDetectAnomalyJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiAnomalyDetectionDetectAnomalyJobDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_detect_anomaly_job" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetDetectAnomalyJobRequest{}

			tmp := rs.Primary.ID
			request.DetectAnomalyJobId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")

			response, err := client.GetDetectAnomalyJob(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_anomaly_detection.DetectAnomalyJobLifecycleStateCanceled): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					return fmt.Errorf("resource still exists")
				}
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
	if !acctest.InSweeperExcludeList("AiAnomalyDetectionDetectAnomalyJob") {
		resource.AddTestSweepers("AiAnomalyDetectionDetectAnomalyJob", &resource.Sweeper{
			Name:         "AiAnomalyDetectionDetectAnomalyJob",
			Dependencies: acctest.DependencyGraph["detectAnomalyJob"],
			F:            sweepAiAnomalyDetectionDetectAnomalyJobResource,
		})
	}
}

func sweepAiAnomalyDetectionDetectAnomalyJobResource(compartment string) error {
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()
	detectAnomalyJobIds, err := getAiAnomalyDetectionDetectAnomalyJobIds(compartment)
	if err != nil {
		return err
	}
	for _, detectAnomalyJobId := range detectAnomalyJobIds {
		if ok := acctest.SweeperDefaultResourceId[detectAnomalyJobId]; !ok {
			deleteDetectAnomalyJobRequest := oci_ai_anomaly_detection.DeleteDetectAnomalyJobRequest{}

			deleteDetectAnomalyJobRequest.DetectAnomalyJobId = &detectAnomalyJobId

			deleteDetectAnomalyJobRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteDetectAnomalyJob(context.Background(), deleteDetectAnomalyJobRequest)
			if error != nil {
				fmt.Printf("Error deleting DetectAnomalyJob %s %s, It is possible that the resource is already deleted. Please verify manually \n", detectAnomalyJobId, error)
				continue
			}
		}
	}
	return nil
}

func getAiAnomalyDetectionDetectAnomalyJobIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "DetectAnomalyJobId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()

	listDetectAnomalyJobsRequest := oci_ai_anomaly_detection.ListDetectAnomalyJobsRequest{}
	listDetectAnomalyJobsRequest.CompartmentId = &compartmentId
	listDetectAnomalyJobsResponse, err := anomalyDetectionClient.ListDetectAnomalyJobs(context.Background(), listDetectAnomalyJobsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting DetectAnomalyJob list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, detectAnomalyJob := range listDetectAnomalyJobsResponse.Items {
		id := *detectAnomalyJob.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "DetectAnomalyJobId", id)
	}
	return resourceIds, nil
}
