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
	oci_datascience "github.com/oracle/oci-go-sdk/v46/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ModelRequiredOnlyResource = ModelResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Required, Create, modelRepresentation)

	ModelResourceConfig = ModelResourceDependencies +
		generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Update, modelRepresentation)

	modelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": Representation{repType: Required, create: `${oci_datascience_model.test_model.id}`},
	}

	modelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": Representation{repType: Required, create: `${var.compartment_id}`},
		"display_name":   Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"id":             Representation{repType: Optional, create: `${oci_datascience_model.test_model.id}`},
		"project_id":     Representation{repType: Optional, create: `${oci_datascience_project.test_project.id}`},
		"state":          Representation{repType: Optional, create: `ACTIVE`},
		"filter":         RepresentationGroup{Required, modelDataSourceFilterRepresentation}}
	modelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   Representation{repType: Required, create: `id`},
		"values": Representation{repType: Required, create: []string{`${oci_datascience_model.test_model.id}`}},
	}

	modelRepresentation = map[string]interface{}{
		"artifact_content_length":      Representation{repType: Required, create: `21002`},
		"model_artifact":               Representation{repType: Required, create: `datascience_model_resource.go`},
		"compartment_id":               Representation{repType: Required, create: `${var.compartment_id}`},
		"project_id":                   Representation{repType: Required, create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": Representation{repType: Optional, create: `attachment; filename=tfTestArtifact`},
		"custom_metadata_list":         RepresentationGroup{Optional, modelCustomMetadataListRepresentation},
		"defined_metadata_list":        RepresentationGroup{Optional, modelDefinedMetadataListRepresentation},
		"defined_tags":                 Representation{repType: Optional, create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  Representation{repType: Optional, create: `description`, update: `description2`},
		"display_name":                 Representation{repType: Optional, create: `displayName`, update: `displayName2`},
		"freeform_tags":                Representation{repType: Optional, create: map[string]string{"Department": "Finance"}, update: map[string]string{"Department": "Accounting"}},
		"input_schema":                 Representation{repType: Optional, create: "{}"},
		"output_schema":                Representation{repType: Optional, create: "{}"},
	}
	modelCustomMetadataListRepresentation = map[string]interface{}{
		"category":    Representation{repType: Optional, create: `Performance`, update: `Performance`},
		"description": Representation{repType: Optional, create: `description`, update: `description`},
		"key":         Representation{repType: Optional, create: `BaseModel1`, update: `BaseModel1`},
		"value":       Representation{repType: Optional, create: `xgb`, update: `xgb`},
	}
	modelDefinedMetadataListRepresentation = map[string]interface{}{
		"key":   Representation{repType: Optional, create: `UseCaseType`, update: `UseCaseType`},
		"value": Representation{repType: Optional, create: `ner`, update: `ner`},
	}

	ModelResourceDependencies = generateResourceFromRepresentationMap("oci_datascience_project", "test_project", Required, Create, projectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelResource_basic")
	defer httpreplay.SaveScenario()

	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := getEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model.test_model"
	datasourceName := "data.oci_datascience_models.test_models"
	singularDatasourceName := "data.oci_datascience_model.test_model"

	var resId, resId2 string
	// Save TF content to create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	saveConfigContent(config+compartmentIdVariableStr+ModelResourceDependencies+
		generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Create, modelRepresentation), "datascience", "model", t)

	ResourceTest(t, testAccCheckDatascienceModelDestroy, []resource.TestStep{
		// verify create
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Required, Create, modelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "21002"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),

				func(s *terraform.State) (err error) {
					resId, err = fromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next create
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies,
		},
		// verify create with optionals
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Create, modelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "21002"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ModelResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Create,
					representationCopyWithNewProperties(modelRepresentation, map[string]interface{}{
						"compartment_id": Representation{repType: Required, create: `${var.compartment_id_for_update}`},
					})),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "21002"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Update, modelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(resourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(resourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(resourceName, "output_schema", "{}"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
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
				generateDataSourceFromRepresentationMap("oci_datascience_models", "test_models", Optional, Update, modelDataSourceRepresentation) +
				compartmentIdVariableStr + ModelResourceDependencies +
				generateResourceFromRepresentationMap("oci_datascience_model", "test_model", Optional, Update, modelRepresentation),
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "models.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.created_by"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.defined_tags.%", "1"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.display_name", "displayName2"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.project_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.time_created"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				generateDataSourceFromRepresentationMap("oci_datascience_model", "test_model", Required, Create, modelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ModelResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_metadata_list.0.category", "Performance"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_metadata_list.0.description", "description"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_metadata_list.0.key", "BaseModel1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "custom_metadata_list.0.value", "xgb"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_metadata_list.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_metadata_list.0.key", "UseCaseType"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_metadata_list.0.value", "ner"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "defined_tags.%", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_schema", "{}"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_schema", "{}"),
				//resource.TestCheckResourceAttr(singularDatasourceName, "state", ACTIVE),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
			),
		},
		// verify resource import
		{Config: config,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"artifact_content_disposition",
				"artifact_content_md5",
				"artifact_last_modified",
				"artifact_content_length",
				"empty_model",
				"model_artifact",
				"model_id",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckDatascienceModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := testAccProvider.Meta().(*OracleClients).dataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model" {
			noResourceFound = false
			request := oci_datascience.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_datascience.ModelLifecycleStateDeleted): true,
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
	if !inSweeperExcludeList("DatascienceModel") {
		resource.AddTestSweepers("DatascienceModel", &resource.Sweeper{
			Name:         "DatascienceModel",
			Dependencies: DependencyGraph["model"],
			F:            sweepDatascienceModelResource,
		})
	}
}

func sweepDatascienceModelResource(compartment string) error {
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()
	modelIds, err := getModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_datascience.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = getRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			waitTillCondition(testAccProvider, &modelId, modelSweepWaitCondition, time.Duration(3*time.Minute),
				modelSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getModelIds(compartment string) ([]string, error) {
	ids := getResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := GetTestClients(&schema.ResourceData{}).dataScienceClient()

	listModelsRequest := oci_datascience.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_datascience.ListModelsLifecycleStateActive
	listModelsResponse, err := dataScienceClient.ListModels(context.Background(), listModelsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting Model list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, model := range listModelsResponse.Items {
		id := *model.Id
		resourceIds = append(resourceIds, id)
		addResourceIdToSweeperResourceIdMap(compartmentId, "ModelId", id)
	}
	return resourceIds, nil
}

func modelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_datascience.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_datascience.ModelLifecycleStateDeleted
	}
	return false
}

func modelSweepResponseFetchOperation(client *OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.dataScienceClient().GetModel(context.Background(), oci_datascience.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
