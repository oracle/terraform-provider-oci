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
	oci_ai_document "github.com/oracle/oci-go-sdk/v65/aidocument"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiDocumentModelRequiredOnlyResource = AiDocumentModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentModelRepresentation)

	AiDocumentModelResourceConfig = AiDocumentModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Update, AiDocumentModelRepresentation)

	AiDocumentAiDocumentModelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_model.test_model.id}`},
	}

	AiDocumentAiDocumentModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
		"project_id":     acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_project.test_project.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentModelDataSourceFilterRepresentation}}
	AiDocumentModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_document_model.test_model.id}`}},
	}

	AiDocumentModelRepresentation = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_type":                 acctest.Representation{RepType: acctest.Required, Create: `KEY_VALUE_EXTRACTION`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_project.test_project.id}`},
		"training_dataset":           acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentModelTrainingDatasetRepresentation},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags": "freeformTags2"}},
		"is_quick_mode":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_training_time_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `0.5`},
		"model_version":              acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAiDocumentDefinedTagsChangesRepresentation},
	}

	AiDocumentModelRepresentation2 = map[string]interface{}{
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_type":                 acctest.Representation{RepType: acctest.Required, Create: `KEY_VALUE_EXTRACTION`},
		"project_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_project.test_project.id}`},
		"training_dataset":           acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentModelTrainingDatasetRepresentation2},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags": "freeformTags2"}},
		"is_quick_mode":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_training_time_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `0.5`},
		"model_version":              acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAiDocumentDefinedTagsChangesRepresentation},
	}

	AiDocumentComposeModelRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_type":     acctest.Representation{RepType: acctest.Required, Create: `KEY_VALUE_EXTRACTION`},
		"project_id":     acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_project.test_project.id}`},
		"component_models": []acctest.RepresentationGroup{
			{RepType: acctest.Optional, Group: AiDocumentModelComponentModelRepresentation1},
			{RepType: acctest.Optional, Group: AiDocumentModelComponentModelRepresentation2},
		},
		"defined_tags":               acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":               acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"freeformTags": "freeformTags"}, Update: map[string]string{"freeformTags": "freeformTags2"}},
		"is_quick_mode":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_training_time_in_hours": acctest.Representation{RepType: acctest.Optional, Create: `0.5`},
		"model_version":              acctest.Representation{RepType: acctest.Optional, Create: `modelVersion`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreAiDocumentDefinedTagsChangesRepresentation},
	}
	AiDocumentModelComponentModelRepresentation1 = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_model.test_model.id}`},
	}
	AiDocumentModelComponentModelRepresentation2 = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_model.test_model2.id}`},
	}

	AiDocumentModelTrainingDatasetRepresentation = map[string]interface{}{
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"bucket":       acctest.Representation{RepType: acctest.Required, Create: `tf_test_bucket`},
		"namespace":    acctest.Representation{RepType: acctest.Required, Create: `axgexwaxnm7k`},
		"object":       acctest.Representation{RepType: acctest.Required, Create: `tf_test_dataset_1680065500556.jsonl`},
	}

	AiDocumentModelTrainingDatasetRepresentation2 = map[string]interface{}{
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"bucket":       acctest.Representation{RepType: acctest.Required, Create: `tf_test_bucket`},
		"namespace":    acctest.Representation{RepType: acctest.Required, Create: `axgexwaxnm7k`},
		"object":       acctest.Representation{RepType: acctest.Required, Create: `tf_test_aadhar_1686719828190.jsonl`},
	}

	AiDocumentModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_project", "test_project", acctest.Required, acctest.Create, AiDocumentProjectRepresentation) +
		DefinedTagsDependencies

	AiDocumentComposeModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_project", "test_project", acctest.Required, acctest.Create, AiDocumentProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model2", acctest.Required, acctest.Create, AiDocumentModelRepresentation2) +
		DefinedTagsDependencies
)

// issue-routing-tag: ai_document/default
func TestAiDocumentModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiDocumentModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_document_model.test_model"
	datasourceName := "data.oci_ai_document_models.test_models"
	singularDatasourceName := "data.oci_ai_document_model.test_model"
	composeResourceName := "oci_ai_document_model.test_compose_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiDocumentModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Create, AiDocumentModelRepresentation), "aidocument", "model", t)

	acctest.ResourceTest(t, testAccCheckAiDocumentModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiDocumentModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "tf_test_dataset_1680065500556.jsonl"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiDocumentModelResourceDependencies,
		},

		// verify Create Compose Model
		{
			Config: config + compartmentIdVariableStr + AiDocumentComposeModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_compose_model", acctest.Optional, acctest.Create, AiDocumentComposeModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(composeResourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(composeResourceName, "description", "description"),
				resource.TestCheckResourceAttr(composeResourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(composeResourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(composeResourceName, "id"),
				resource.TestCheckResourceAttr(composeResourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(composeResourceName, "max_training_time_in_hours", "0.5"),
				resource.TestCheckResourceAttr(composeResourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(composeResourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(composeResourceName, "project_id"),
				resource.TestCheckResourceAttr(composeResourceName, "component_models.#", "2"),
				resource.TestCheckResourceAttrSet(composeResourceName, "component_models.0.model_id"),
				resource.TestCheckResourceAttrSet(composeResourceName, "state"),
				resource.TestCheckResourceAttrSet(composeResourceName, "time_created"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, composeResourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiDocumentModelResourceDependencies,
		},

		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiDocumentModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Create, AiDocumentModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_time_in_hours", "0.5"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "tf_test_dataset_1680065500556.jsonl"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiDocumentModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiDocumentModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_time_in_hours", "0.5"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "tf_test_dataset_1680065500556.jsonl"),

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
			Config: config + compartmentIdVariableStr + AiDocumentModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Update, AiDocumentModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "max_training_time_in_hours", "0.5"),
				resource.TestCheckResourceAttr(resourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.object", "tf_test_dataset_1680065500556.jsonl"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_document_models", "test_models", acctest.Optional, acctest.Update, AiDocumentAiDocumentModelDataSourceRepresentation) +
				compartmentIdVariableStr + AiDocumentModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Optional, acctest.Update, AiDocumentModelRepresentation),
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentAiDocumentModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiDocumentModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "component_models.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_composed_model"),
				resource.TestCheckResourceAttr(singularDatasourceName, "is_quick_mode", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "max_training_time_in_hours", "0.5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "metrics.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_version", "modelVersion"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trained_time_in_hours"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.object", "tf_test_dataset_1680065500556.jsonl"),
			),
		},
		// verify resource import
		{
			Config:            config + AiDocumentModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ResourceName:      resourceName,
		},
	})
}

func testAccCheckAiDocumentModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceDocumentClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_document_model" {
			noResourceFound = false
			request := oci_ai_document.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_document")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_document.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiDocumentModel") {
		resource.AddTestSweepers("AiDocumentModel", &resource.Sweeper{
			Name:         "AiDocumentModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepAiDocumentModelResource,
		})
	}
}

func sweepAiDocumentModelResource(compartment string) error {
	aiServiceDocumentClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceDocumentClient()
	modelIds, err := getAiDocumentModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_ai_document.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_document")
			_, error := aiServiceDocumentClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, AiDocumentModelSweepWaitCondition, time.Duration(3*time.Minute),
				AiDocumentModelSweepResponseFetchOperation, "ai_document", true)
		}
	}
	return nil
}

func getAiDocumentModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceDocumentClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceDocumentClient()

	listModelsRequest := oci_ai_document.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_ai_document.ModelLifecycleStateActive
	listModelsResponse, err := aiServiceDocumentClient.ListModels(context.Background(), listModelsRequest)

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

func AiDocumentModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_ai_document.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_ai_document.ModelLifecycleStateDeleted
	}
	return false
}

func AiDocumentModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceDocumentClient().GetModel(context.Background(), oci_ai_document.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
