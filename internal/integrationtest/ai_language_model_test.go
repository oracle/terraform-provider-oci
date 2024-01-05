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
	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	"github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiLanguageModelRequiredOnlyResource = AiLanguageModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation)

	AiLanguageModelResourceConfig = AiLanguageModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Update, AiLanguageModelRepresentation)

	AiLanguageAiLanguageModelSingularDataSourceRepresentation = map[string]interface{}{
		"id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_model.test_model.id}`},
	}

	AiLanguageAiLanguageModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		// "model_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_model.test_model.id}`},
		"project_id": acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_language_project.test_project.id}`},
		"state":      acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"filter":     acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelDataSourceFilterRepresentation}}
	AiLanguageModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_ai_language_model.test_model.id}`}},
	}

	AiLanguageModelRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"model_details":    acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelModelDetailsRepresentation},
		"project_id":       acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_project.test_project.id}`},
		"training_dataset": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelTrainingDatasetRepresentation},
		// "defined_tags":     acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}},
		"test_strategy": acctest.RepresentationGroup{RepType: acctest.Optional, Group: AiLanguageModelTestStrategyRepresentation},
	}
	AiLanguageModelModelDetailsRepresentation = map[string]interface{}{
		"model_type": acctest.Representation{RepType: acctest.Required, Create: `NAMED_ENTITY_RECOGNITION`},
		// "classification_mode": acctest.RepresentationGroup{RepType: acctest.Optional, Group: AiLanguageModelModelDetailsClassificationModeRepresentation},
		"version":       acctest.Representation{RepType: acctest.Optional, Create: `V1.0`},
		"language_code": acctest.Representation{RepType: acctest.Required, Create: `en`},
	}
	AiLanguageModelTrainingDatasetRepresentation = map[string]interface{}{
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		// "dataset_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
		"location_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelTrainingDatasetLocationDetailsRepresentation},
	}
	AiLanguageModelTestStrategyRepresentation = map[string]interface{}{
		"strategy_type":      acctest.Representation{RepType: acctest.Required, Create: `TEST_AND_VALIDATION_DATASET`},
		"testing_dataset":    acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelTestStrategyTestingDatasetRepresentation},
		"validation_dataset": acctest.RepresentationGroup{RepType: acctest.Optional, Group: AiLanguageModelTestStrategyValidationDatasetRepresentation},
	}
	// AiLanguageModelModelDetailsClassificationModeRepresentation = map[string]interface{}{
	// 	"classification_mode": acctest.Representation{RepType: acctest.Optional, Create: `MULTI_CLASS`},
	// }
	AiLanguageModelTrainingDatasetLocationDetailsRepresentation = map[string]interface{}{
		"bucket":        acctest.Representation{RepType: acctest.Required, Create: `TERSI-Test`},
		"location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_LIST`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `idngwwc5ajp5`},
		"object_names":  acctest.Representation{RepType: acctest.Required, Create: []string{`test.jsonl`}},
	}
	AiLanguageModelTestStrategyTestingDatasetRepresentation = map[string]interface{}{
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		// "dataset_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
		"location_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelTestStrategyTestingDatasetLocationDetailsRepresentation},
	}
	AiLanguageModelTestStrategyValidationDatasetRepresentation = map[string]interface{}{
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		// "dataset_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_data_labeling_service_dataset.test_dataset.id}`},
		"location_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiLanguageModelTestStrategyValidationDatasetLocationDetailsRepresentation},
	}
	AiLanguageModelTestStrategyTestingDatasetLocationDetailsRepresentation = map[string]interface{}{
		"bucket":        acctest.Representation{RepType: acctest.Required, Create: `TERSI-Test`},
		"location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_LIST`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `idngwwc5ajp5`},
		"object_names":  acctest.Representation{RepType: acctest.Required, Create: []string{`test.jsonl`}},
	}
	AiLanguageModelTestStrategyValidationDatasetLocationDetailsRepresentation = map[string]interface{}{
		"bucket":        acctest.Representation{RepType: acctest.Required, Create: `TERSI-Test`},
		"location_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_LIST`},
		"namespace":     acctest.Representation{RepType: acctest.Required, Create: `idngwwc5ajp5`},
		"object_names":  acctest.Representation{RepType: acctest.Required, Create: []string{`test.jsonl`}},
	}

	AiLanguageModelResourceDependencies =
	// acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation) +
	acctest.GenerateResourceFromRepresentationMap("oci_ai_language_project", "test_project", acctest.Required, acctest.Create, AiLanguageProjectRepresentation)
	// +
	// acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetRepresentation) +
	// DefinedTagsDependencies + acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentationDataset) +
	// acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Required, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: ai_language/default
func TestAiLanguageModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiLanguageModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_ai_language_model.test_model"
	datasourceName := "data.oci_ai_language_models.test_models"
	singularDatasourceName := "data.oci_ai_language_model.test_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiLanguageModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Create, AiLanguageModelRepresentation), "ailanguage", "model", t)

	acctest.ResourceTest(t, testAccCheckAiLanguageModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "model_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.model_type", "NAMED_ENTITY_RECOGNITION"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiLanguageModelResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiLanguageModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Create, AiLanguageModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_details.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.0.classification_mode", "MULTI_CLASS"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.language_code", "en"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.model_type", "NAMED_ENTITY_RECOGNITION"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.version", "V1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.strategy_type", "TEST_AND_VALIDATION_DATASET"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.testing_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.validation_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "training_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.object_names.#", "1"),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + AiLanguageModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(AiLanguageModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_details.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.0.classification_mode", "MULTI_CLASS"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.language_code", "en"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.model_type", "NAMED_ENTITY_RECOGNITION"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.version", "V1.0"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.strategy_type", "TEST_AND_VALIDATION_DATASET"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.testing_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.validation_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "training_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.object_names.#", "1"),

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
			Config: config + compartmentIdVariableStr + AiLanguageModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Update, AiLanguageModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "model_details.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.#", "1"),
				// resource.TestCheckResourceAttr(resourceName, "model_details.0.classification_mode.0.classification_mode", "MULTI_CLASS"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.language_code", "en"),
				resource.TestCheckResourceAttr(resourceName, "model_details.0.model_type", "NAMED_ENTITY_RECOGNITION"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.strategy_type", "TEST_AND_VALIDATION_DATASET"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.testing_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.testing_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "test_strategy.0.validation_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "test_strategy.0.validation_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.#", "1"),
				// resource.TestCheckResourceAttrSet(resourceName, "training_dataset.0.dataset_id"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(resourceName, "training_dataset.0.location_details.0.object_names.#", "1"),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_models", "test_models", acctest.Optional, acctest.Update, AiLanguageAiLanguageModelDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Optional, acctest.Update, AiLanguageModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
				// resource.TestCheckResourceAttrSet(datasourceName, "model_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "project_id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttr(datasourceName, "model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageAiLanguageModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "evaluation_results.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_details.#", "1"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "model_details.0.classification_mode.#", "1"),
				// resource.TestCheckResourceAttr(singularDatasourceName, "model_details.0.classification_mode.0.classification_mode", "MULTI_CLASS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_details.0.language_code", "en"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_details.0.model_type", "NAMED_ENTITY_RECOGNITION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_details.0.version", "V1.0"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.strategy_type", "TEST_AND_VALIDATION_DATASET"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.testing_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "test_strategy.0.validation_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.location_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.location_details.0.bucket", "TERSI-Test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.location_details.0.location_type", "OBJECT_LIST"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.location_details.0.namespace", "idngwwc5ajp5"),
				resource.TestCheckResourceAttr(singularDatasourceName, "training_dataset.0.location_details.0.object_names.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "version"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiLanguageModelRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckAiLanguageModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).AiServiceLanguageClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_ai_language_model" {
			noResourceFound = false
			request := oci_ai_language.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_ai_language.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("AiLanguageModel") {
		resource.AddTestSweepers("AiLanguageModel", &resource.Sweeper{
			Name:         "AiLanguageModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepAiLanguageModelResource,
		})
	}
}

func sweepAiLanguageModelResource(compartment string) error {
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()
	modelIds, err := getAiLanguageModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_ai_language.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "ai_language")
			_, error := aiServiceLanguageClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, AiLanguageModelSweepWaitCondition, time.Duration(3*time.Minute),
				AiLanguageModelSweepResponseFetchOperation, "ai_language", true)
		}
	}
	return nil
}

func getAiLanguageModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	aiServiceLanguageClient := acctest.GetTestClients(&schema.ResourceData{}).AiServiceLanguageClient()

	listModelsRequest := oci_ai_language.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_ai_language.ModelLifecycleStateActive
	listModelsResponse, err := aiServiceLanguageClient.ListModels(context.Background(), listModelsRequest)

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

func AiLanguageModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_ai_language.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_ai_language.ModelLifecycleStateDeleted
	}
	return false
}

func AiLanguageModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.AiServiceLanguageClient().GetModel(context.Background(), oci_ai_language.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
