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
	"github.com/oracle/oci-go-sdk/v56/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	ModelRequiredOnlyResource = ModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, modelRepresentation)

	ModelResourceConfig = ModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, modelRepresentation)

	modelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_model.test_model.id}`},
	}

	modelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_model.test_model.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: modelDataSourceFilterRepresentation}}
	modelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_datascience_model.test_model.id}`}},
	}

	modelRepresentation = map[string]interface{}{
		"artifact_content_length":      acctest.Representation{RepType: acctest.Required, Create: `21002`},
		"model_artifact":               acctest.Representation{RepType: acctest.Required, Create: `datascience_model_resource.go`},
		"compartment_id":               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=tfTestArtifact`},
		"custom_metadata_list":         acctest.RepresentationGroup{RepType: acctest.Optional, Group: modelCustomMetadataListRepresentation},
		"defined_metadata_list":        acctest.RepresentationGroup{RepType: acctest.Optional, Group: modelDefinedMetadataListRepresentation},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"input_schema":                 acctest.Representation{RepType: acctest.Optional, Create: "{}"},
		"output_schema":                acctest.Representation{RepType: acctest.Optional, Create: "{}"},
	}
	modelCustomMetadataListRepresentation = map[string]interface{}{
		"category":    acctest.Representation{RepType: acctest.Optional, Create: `Performance`, Update: `Performance`},
		"description": acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description`},
		"key":         acctest.Representation{RepType: acctest.Optional, Create: `BaseModel1`, Update: `BaseModel1`},
		"value":       acctest.Representation{RepType: acctest.Optional, Create: `xgb`, Update: `xgb`},
	}
	modelDefinedMetadataListRepresentation = map[string]interface{}{
		"key":   acctest.Representation{RepType: acctest.Optional, Create: `UseCaseType`, Update: `UseCaseType`},
		"value": acctest.Representation{RepType: acctest.Optional, Create: `ner`, Update: `ner`},
	}

	ModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, projectRepresentation) +
		DefinedTagsDependencies
)

// issue-routing-tag: datascience/default
func TestDatascienceModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_model.test_model"
	datasourceName := "data.oci_datascience_models.test_models"
	singularDatasourceName := "data.oci_datascience_model.test_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+ModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelRepresentation), "datascience", "model", t)

	acctest.ResourceTest(t, testAccCheckDatascienceModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, modelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "artifact_content_length", "21002"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_content_md5"),
				resource.TestCheckResourceAttrSet(resourceName, "artifact_last_modified"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create, modelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + ModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(modelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
			Config: config + compartmentIdVariableStr + ModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, modelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_models", "test_models", acctest.Optional, acctest.Update, modelDataSourceRepresentation) +
				compartmentIdVariableStr + ModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Optional, acctest.Update, modelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "models.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "models.0.compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(datasourceName, "models.0.created_by"),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_model", "test_model", acctest.Required, acctest.Create, modelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + ModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DataScienceClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_datascience_model" {
			noResourceFound = false
			request := oci_datascience.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

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
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DatascienceModel") {
		resource.AddTestSweepers("DatascienceModel", &resource.Sweeper{
			Name:         "DatascienceModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepDatascienceModelResource,
		})
	}
}

func sweepDatascienceModelResource(compartment string) error {
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()
	modelIds, err := getModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_datascience.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")
			_, error := dataScienceClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, modelSweepWaitCondition, time.Duration(3*time.Minute),
				modelSweepResponseFetchOperation, "datascience", true)
		}
	}
	return nil
}

func getModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataScienceClient := acctest.GetTestClients(&schema.ResourceData{}).DataScienceClient()

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
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "ModelId", id)
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

func modelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataScienceClient().GetModel(context.Background(), oci_datascience.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
