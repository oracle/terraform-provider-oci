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
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	oci_ai_vision "github.com/oracle/oci-go-sdk/v65/aivision"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	AiVisionModelRequiredOnlyResource = AiVisionModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Required, acctest.Create, visionModelRepresentation)

	AiVisionModelResourceConfig = AiVisionModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Update, visionModelRepresentation)

	AiVisionmodelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_model.test_model.id}`},
	}

	AiVisionmodelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_vision_model.test_model.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_vision_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: visionModelDataSourceFilterRepresentation}}
	visionModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_vision_model.test_model.id}`}},
	}

	visionModelRepresentation = map[string]interface{}{
		"compartment_id":                 acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_type":                     acctest.Representation{RepType: acctest.Required, Create: `IMAGE_CLASSIFICATION`},
		"project_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_vision_project.test_project.id}`},
		"training_dataset":               acctest.RepresentationGroup{RepType: acctest.Required, Group: visionModelTrainingDatasetRepresentation},
		"description":                    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"is_quick_mode":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_training_duration_in_hours": acctest.Representation{RepType: acctest.Required, Create: `0.01`},
		"model_version":                  acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`},
	}
	visionModelTrainingDatasetRepresentation = map[string]interface{}{
		"dataset_type":   acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"bucket":         acctest.Representation{RepType: acctest.Required, Create: `golden_dataset`},
		"namespace_name": acctest.Representation{RepType: acctest.Required, Create: `axhheqi2ofpb`},
		"object":         acctest.Representation{RepType: acctest.Required, Create: `a_hymenoptera_v3.json`},
	}

	AiVisionModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_project", "test_project", acctest.Required, acctest.Create, visionvisionProjectRepresentation)
)

// issue-routing-tag: ai_vision/default
func TestAiVisionModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiVisionModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_vision_model.test_model"
	datasourceName := "data.oci_ai_vision_models.test_models"
	singularDatasourceName := "data.oci_ai_vision_model.test_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiVisionModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Create, visionModelRepresentation), "aivision", "model", t)

	acctest.ResourceTest(t, testAccCheckAiVisionModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiVisionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Required, acctest.Create, visionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_type", "IMAGE_CLASSIFICATION"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiVisionModelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiVisionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Create, visionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_duration_in_hours", "0.01"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "IMAGE_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "golden_dataset"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace_name", "axhheqi2ofpb"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "a_hymenoptera_v3.json"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiVisionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(visionModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_duration_in_hours", "0.01"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "IMAGE_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "golden_dataset"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace_name", "axhheqi2ofpb"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "a_hymenoptera_v3.json"),

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
			Config: config + compartmentIdVariableStr + AiVisionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Update, visionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_duration_in_hours", "0.01"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "IMAGE_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "golden_dataset"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace_name", "axhheqi2ofpb"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "a_hymenoptera_v3.json"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_models", "test_models", acctest.Optional, acctest.Update, AiVisionmodelDataSourceRepresentation) +
				compartmentIdVariableStr + AiVisionModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Optional, acctest.Update, visionModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_vision_model", "test_model", acctest.Required, acctest.Create, AiVisionmodelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiVisionModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "average_precision"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "confidence_threshold"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_training_duration_in_hours", "0.01"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "metrics"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "IMAGE_CLASSIFICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "precision"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "recall"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "test_image_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "total_image_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trained_duration_in_hours"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.bucket", "golden_dataset"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.namespace_name", "axhheqi2ofpb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.object", "a_hymenoptera_v3.json"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiVisionModelRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiVisionModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceVisionClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_vision_model" {
			noResourceFound = false
			request := oci_ai_vision.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_vision.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiVisionModel") {
		resource.AddTestSweepers("AiVisionModel", &resource.Sweeper{
			Name:         "AiVisionModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepAiVisionModelResource,
		})
	}
}

func sweepAiVisionModelResource(compartment string) error {
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()
	modelIds, err := getAiVisionModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_ai_vision.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_vision")
			_, error := aiServiceVisionClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, AiVisionmodelsSweepWaitCondition, time.Duration(3*time.Minute),
				AiVisionmodelsSweepResponseFetchOperation, "ai_vision", true)
		}
	}
	return nil
}

func getAiVisionModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceVisionClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceVisionClient()

	listModelsRequest := oci_ai_vision.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_ai_vision.ModelLifecycleStateActive
	listModelsResponse, err := aiServiceVisionClient.ListModels(context.Background(), listModelsRequest)

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

func AiVisionmodelsSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_ai_vision.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_ai_vision.ModelLifecycleStateDeleted
	}
	return false
}

func AiVisionmodelsSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceVisionClient().GetModel(context.Background(), oci_ai_vision.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
