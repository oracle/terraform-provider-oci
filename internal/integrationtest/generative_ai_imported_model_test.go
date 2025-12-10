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
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	GenerativeAiImportedModelRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Required, acctest.Create, GenerativeAiImportedModelRepresentation)

	GenerativeAiImportedModelResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Update, GenerativeAiImportedModelRepresentation)

	GenerativeAiImportedModelSingularDataSourceRepresentation = map[string]interface{}{
		"imported_model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_imported_model.test_imported_model.id}`},
	}

	GenerativeAiImportedModelsDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"capability":     acctest.Representation{RepType: acctest.Optional, Create: []string{`TEXT_TO_TEXT`}},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_imported_model.test_imported_model.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor":         acctest.Representation{RepType: acctest.Optional, Create: `vendor`, Update: `vendor2`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiImportedModelDataSourceFilterRepresentation},
	}

	GenerativeAiImportedModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_imported_model.test_imported_model.id}`}},
	}

	generativeAiImportedModelVersion  = fmt.Sprintf("version-%d", time.Now().UnixNano())
	generativeAiImportedModelVersion2 = fmt.Sprintf("version2-%d", time.Now().UnixNano())

	GenerativeAiImportedModelDataSourceRepresentation = map[string]interface{}{
		"source_type":  acctest.Representation{RepType: acctest.Required, Create: `HUGGING_FACE_MODEL`},
		"model_id":     acctest.Representation{RepType: acctest.Required, Create: `google/gemma-3-12b-it`},
		"access_token": acctest.Representation{RepType: acctest.Required, Create: `${var.hf_access_token}`},
	}

	GenerativeAiImportedModelRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"data_source":    acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiImportedModelDataSourceRepresentation},
		"capabilities":   acctest.Representation{RepType: acctest.Required, Create: []string{`TEXT_TO_TEXT`}},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"vendor":         acctest.Representation{RepType: acctest.Optional, Create: `vendor`, Update: `vendor2`},
		"version":        acctest.Representation{RepType: acctest.Optional, Create: generativeAiImportedModelVersion, Update: generativeAiImportedModelVersion2},
	}
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiImportedModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiImportedModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_imported_model.test_imported_model"

	datasourceName := "data.oci_generative_ai_imported_models.test_imported_models"
	singularDatasourceName := "data.oci_generative_ai_imported_model.test_imported_model"

	hfAccessToken := utils.GetEnvSettingWithBlankDefault("hf_access_token")
	hfAccessTokenVariableStr := fmt.Sprintf("variable \"hf_access_token\" { default = \"%s\" }\n", hfAccessToken)

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+hfAccessTokenVariableStr+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Create, GenerativeAiImportedModelRepresentation), "generativeai", "importedModel", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiImportedModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + hfAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Required, acctest.Create, GenerativeAiImportedModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.model_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + hfAccessTokenVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + hfAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Create, GenerativeAiImportedModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.source_type", "HUGGING_FACE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiImportedModelVersion),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + hfAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiImportedModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.source_type", "HUGGING_FACE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiImportedModelVersion),

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
			Config: config + compartmentIdVariableStr + hfAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Update, GenerativeAiImportedModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "capabilities.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "data_source.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.access_token"),
				resource.TestCheckResourceAttrSet(resourceName, "data_source.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "data_source.0.source_type", "HUGGING_FACE_MODEL"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor2"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiImportedModelVersion2),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_imported_models", "test_imported_models", acctest.Optional, acctest.Update, GenerativeAiImportedModelsDataSourceRepresentation) +
				compartmentIdVariableStr + hfAccessTokenVariableStr +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Optional, acctest.Update, GenerativeAiImportedModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "capability.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "vendor", "vendor2"),

				resource.TestCheckResourceAttr(datasourceName, "imported_model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "imported_model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_imported_model", "test_imported_model", acctest.Required, acctest.Create, GenerativeAiImportedModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + hfAccessTokenVariableStr + GenerativeAiImportedModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "imported_model_id"),
			),
		},
		// verify resource import
		{
			Config:            config + hfAccessTokenVariableStr + compartmentIdVariableStr + GenerativeAiImportedModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"capabilities",
				"compartment_id",
				"data_source",
				"description",
				"display_name",
				"freeform_tags",
				"id",
				"lifecycle_details",
				"previous_state",
				"state",
				"system_tags",
				"time_created",
				"time_updated",
				"vendor",
				"version",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGenerativeAiImportedModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_imported_model" {
			noResourceFound = false
			request := oci_generative_ai.GetImportedModelRequest{}

			tmp := rs.Primary.ID
			request.ImportedModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetImportedModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.ImportedModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiImportedModel") {
		resource.AddTestSweepers("GenerativeAiImportedModel", &resource.Sweeper{
			Name:         "GenerativeAiImportedModel",
			Dependencies: acctest.DependencyGraph["importedModel"],
			F:            sweepGenerativeAiImportedModelResource,
		})
	}
}

func sweepGenerativeAiImportedModelResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	importedModelIds, err := getGenerativeAiImportedModelIds(compartment)
	if err != nil {
		return err
	}
	for _, importedModelId := range importedModelIds {
		if ok := acctest.SweeperDefaultResourceId[importedModelId]; !ok {
			deleteImportedModelRequest := oci_generative_ai.DeleteImportedModelRequest{}

			deleteImportedModelRequest.ImportedModelId = &importedModelId

			deleteImportedModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteImportedModel(context.Background(), deleteImportedModelRequest)
			if error != nil {
				fmt.Printf("Error deleting ImportedModel %s %s, It is possible that the resource is already deleted. Please verify manually \n", importedModelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &importedModelId, GenerativeAiImportedModelSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiImportedModelSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiImportedModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ImportedModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listImportedModelsRequest := oci_generative_ai.ListImportedModelsRequest{}
	listImportedModelsRequest.CompartmentId = &compartmentId
	listImportedModelsRequest.LifecycleState = oci_generative_ai.ImportedModelLifecycleStateActive
	listImportedModelsResponse, err := generativeAiClient.ListImportedModels(context.Background(), listImportedModelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting ImportedModel list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, importedModel := range listImportedModelsResponse.Items {
		id := *importedModel.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ImportedModelId", id)
	}
	return resourceIds, nil
}

func GenerativeAiImportedModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if importedModelResponse, ok := response.Response.(oci_generative_ai.GetImportedModelResponse); ok {
		return importedModelResponse.LifecycleState != oci_generative_ai.ImportedModelLifecycleStateDeleted
	}
	return false
}

func GenerativeAiImportedModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetImportedModel(context.Background(), oci_generative_ai.GetImportedModelRequest{
		ImportedModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
