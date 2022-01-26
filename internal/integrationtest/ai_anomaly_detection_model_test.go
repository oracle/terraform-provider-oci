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
	oci_ai_anomaly_detection "github.com/oracle/oci-go-sdk/v56/aianomalydetection"
	"github.com/oracle/oci-go-sdk/v56/common"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	AiAnomalyDetectionModelRequiredOnlyResource = AiAnomalyDetectionModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Required, acctest.Create, aiAnomalyDetectionModelRepresentation)

	AiAnomalyDetectionModelResourceConfig = AiAnomalyDetectionModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Update, aiAnomalyDetectionModelRepresentation)

	aiAnomalyDetectionModelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_model.test_model.id}`},
	}

	aiAnomalyDetectionModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_anomaly_detection_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: aiAnomalyDetectionModelDataSourceFilterRepresentation}}
	aiAnomalyDetectionModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_model.test_model.id}`}},
	}

	aiAnomalyDetectionModelRepresentation = map[string]interface{}{
		"compartment_id":         acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_training_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: modelModelTrainingDetailsRepresentation},
		"project_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_anomaly_detection_data_asset.test_data_asset.project_id}`},
		"defined_tags":           acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":            acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":           acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":          acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"lifecycle":              acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}
	modelModelTrainingDetailsRepresentation = map[string]interface{}{
		"data_asset_ids":    acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_anomaly_detection_data_asset.test_data_asset.id}`}},
		"target_fap":        acctest.Representation{RepType: acctest.Optional, Create: `0.01`},
		"training_fraction": acctest.Representation{RepType: acctest.Optional, Create: `0.7`},
	}

	ignoreModelTrainingResultsChangesRep = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{`model_training_details[0].target_fap`, `model_training_details[0].target_fap`, `model_training_results`}},
	}

	AiAnomalyDetectionModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_project", "test_project", acctest.Required, acctest.Create, aiAnomalyDetectionProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_data_asset", "test_data_asset", acctest.Required, acctest.Create, aiAnomalyDetectionDataAssetRepresentation) +
		DefinedTagsDependencies
)

func TestAiAnomalyDetectionModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiAnomalyDetectionModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_anomaly_detection_model.test_model"
	datasourceName := "data.oci_ai_anomaly_detection_models.test_models"
	singularDatasourceName := "data.oci_ai_anomaly_detection_model.test_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiAnomalyDetectionModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Create, aiAnomalyDetectionModelRepresentation), "aianomalydetection", "model", t)

	acctest.ResourceTest(t, testAccCheckAiAnomalyDetectionModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Required, acctest.Create, aiAnomalyDetectionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.0.data_asset_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionModelResourceDependencies,
		},
		//verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Create, aiAnomalyDetectionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.0.data_asset_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.target_fap"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.training_fraction"),
				resource.TestCheckResourceAttr(resourceName, "model_training_results.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiAnomalyDetectionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(aiAnomalyDetectionModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.0.data_asset_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.target_fap"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.training_fraction"),
				resource.TestCheckResourceAttr(resourceName, "model_training_results.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Update, aiAnomalyDetectionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_training_details.0.data_asset_ids.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.target_fap"),
				resource.TestCheckResourceAttrSet(resourceName, "model_training_details.0.training_fraction"),
				resource.TestCheckResourceAttr(resourceName, "model_training_results.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_models", "test_models", acctest.Optional, acctest.Update, aiAnomalyDetectionModelDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Optional, acctest.Update, aiAnomalyDetectionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_anomaly_detection_model", "test_model", acctest.Required, acctest.Create, aiAnomalyDetectionModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiAnomalyDetectionModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_training_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_training_details.0.data_asset_ids.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_training_details.0.target_fap"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_training_details.0.training_fraction"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_training_results.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AiAnomalyDetectionModelResourceConfig,
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

func testAccCheckAiAnomalyDetectionModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AnomalyDetectionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_anomaly_detection_model" {
			noResourceFound = false
			request := oci_ai_anomaly_detection.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_anomaly_detection.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiAnomalyDetectionModel") {
		resource.AddTestSweepers("AiAnomalyDetectionModel", &resource.Sweeper{
			Name:         "AiAnomalyDetectionModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepAiAnomalyDetectionModelResource,
		})
	}
}

func sweepAiAnomalyDetectionModelResource(compartment string) error {
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()
	modelIds, err := aiAnomalyDetectionGetModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_ai_anomaly_detection.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_anomaly_detection")
			_, error := anomalyDetectionClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, aiAnomalyDetectionModelSweepWaitCondition, time.Duration(3*time.Minute),
				aiAnomalyDetectionModelSweepResponseFetchOperation, "ai_anomaly_detection", true)
		}
	}
	return nil
}

func aiAnomalyDetectionGetModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	anomalyDetectionClient := acctest.GetTestClients(&schema.ResourceData{}).AnomalyDetectionClient()

	listModelsRequest := oci_ai_anomaly_detection.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_ai_anomaly_detection.ModelLifecycleStateActive
	listModelsResponse, err := anomalyDetectionClient.ListModels(context.Background(), listModelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Model list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, model := range listModelsResponse.Items {
		id := *model.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelId", id)
	}
	return resourceIds, nil
}

func aiAnomalyDetectionModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_ai_anomaly_detection.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_ai_anomaly_detection.ModelLifecycleStateDeleted
	}
	return false
}

func aiAnomalyDetectionModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AnomalyDetectionClient().GetModel(context.Background(), oci_ai_anomaly_detection.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
