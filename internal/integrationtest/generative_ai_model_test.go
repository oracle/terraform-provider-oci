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
	GenerativeAiModelRequiredOnlyResource = GenerativeAiModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelTfewRepresentation)

	GenerativeAiModelResourceConfig = GenerativeAiModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelTfewRepresentation)

	GenerativeAiModelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_model.test_model.id}`},
	}

	GenerativeAiModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"capability":     acctest.Representation{RepType: acctest.Optional, Create: []string{`FINE_TUNE`}},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: Name1, Update: Name2},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_model.test_model.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor":         acctest.Representation{RepType: acctest.Optional, Create: `vendor`}, // Only base model vendor can be updated, not the case here
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelDataSourceFilterRepresentation}}
	GenerativeAiModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_model.test_model.id}`}},
	}

	GenerativeAiModelLoraRepresentation = map[string]interface{}{
		"base_model_id":     acctest.Representation{RepType: acctest.Required, Create: `${local.llama_base_model_id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fine_tune_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelLoraFineTuneDetailsRepresentation},
		"description":       acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":      acctest.Representation{RepType: acctest.Optional, Create: loraName1, Update: loraName2},
		"freeform_tags":     acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"vendor":            acctest.Representation{RepType: acctest.Optional, Create: `meta`}, // Only base model vendor can be updated, not the case here
		"version":           acctest.Representation{RepType: acctest.Optional, Create: Version1, Update: Version2},
	}
	GenerativeAiModelLoraFineTuneDetailsRepresentation = map[string]interface{}{
		"dedicated_ai_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster_large_v2.id}`},
		"training_dataset":        acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsTrainingDatasetRepresentation},
		"training_config":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelLoraFineTuneDetailsTrainingConfigRepresentation},
	}
	GenerativeAiModelLoraFineTuneDetailsTrainingConfigRepresentation = map[string]interface{}{
		"training_config_type":                acctest.Representation{RepType: acctest.Required, Create: `LORA_TRAINING_CONFIG`},
		"early_stopping_patience":             acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"early_stopping_threshold":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"learning_rate":                       acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"log_model_metrics_interval_in_steps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"total_training_epochs":               acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"training_batch_size":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lora_r":                              acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lora_alpha":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"lora_dropout":                        acctest.Representation{RepType: acctest.Optional, Create: `1.0`},
	}

	GenerativeAiModelTfewRepresentation = map[string]interface{}{
		"base_model_id":     acctest.Representation{RepType: acctest.Required, Create: `${local.cohere_base_model_id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fine_tune_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsRepresentation},
		//"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: Name1, Update: Name2},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"vendor":        acctest.Representation{RepType: acctest.Optional, Create: `vendor`}, // Only base model vendor can be updated, not the case here
		"version":       acctest.Representation{RepType: acctest.Optional, Create: Version1, Update: Version2},
	}
	GenerativeAiModelFineTuneDetailsRepresentation = map[string]interface{}{
		"dedicated_ai_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`},
		"training_dataset":        acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsTrainingDatasetRepresentation},
		"training_config":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsTrainingConfigRepresentation},
	}
	GenerativeAiModelFineTuneDetailsTrainingDatasetRepresentation = map[string]interface{}{
		"bucket":       acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.fine_tune_bucket.name}`},
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"namespace":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":       acctest.Representation{RepType: acctest.Required, Create: FineTuneDataObjectName},
	}
	GenerativeAiModelFineTuneDetailsTrainingConfigRepresentation = map[string]interface{}{
		"training_config_type":                acctest.Representation{RepType: acctest.Required, Create: `TFEW_TRAINING_CONFIG`},
		"early_stopping_patience":             acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"early_stopping_threshold":            acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"learning_rate":                       acctest.Representation{RepType: acctest.Optional, Create: `1`},
		"log_model_metrics_interval_in_steps": acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"total_training_epochs":               acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"training_batch_size":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	FineTuneDataObjectStorageBucketRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":           acctest.Representation{RepType: acctest.Required, Create: FineTuneDataBucketName},
		"namespace":      acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
	}

	FineTuneDataObjectStorageObjectRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.fine_tune_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: FineTuneDataObjectName},
		"content":   acctest.Representation{RepType: acctest.Optional, Create: Prompts},
	}

	GenerativeAiLlamaModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster_large_v2", acctest.Required, acctest.Create, GenerativeAiLoraFineTuningDedicatedAiClusterRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "fine_tune_bucket", acctest.Required, acctest.Create, FineTuneDataObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "fine_tune_data", acctest.Optional, acctest.Create, FineTuneDataObjectStorageObjectRepresentation) +
		llamaBaseModelDependencies

	llamaBaseModelDependencies = `
locals {
  llama_filtered_models = [
	for item in data.oci_generative_ai_models.llama_base_models.model_collection[0].items : item
	  if (
        (item.version == "1.0.0")
		&& contains(item.capabilities, "FINE_TUNE")
		&& (item.display_name == "meta.llama-3-70b-instruct")
	  )
	]

 llama_base_model_id = local.llama_filtered_models[0].id
}
data "oci_generative_ai_models" "llama_base_models" {
  compartment_id = var.compartment_id 
}
`

	GenerativeAiModelResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiFineTuningDedicatedAiClusterRepresentation) +
		// Avoid loop - acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelTfewRepresentation) +
		// Cannot test in home region due to GPU - DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "fine_tune_bucket", acctest.Required, acctest.Create, FineTuneDataObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "fine_tune_data", acctest.Optional, acctest.Create, FineTuneDataObjectStorageObjectRepresentation) +
		baseModelDependencies

	baseModelDependencies = `
locals {

  filtered_models = [
	for item in data.oci_generative_ai_models.base_models.model_collection[0].items : item
	  if (
		(item.version == "14.2")
		&& contains(item.capabilities, "FINE_TUNE")
		&& (item.display_name == "cohere.command-light")
	  )
	]

  cohere_base_model_id = local.filtered_models[0].id
}

data "oci_generative_ai_models" "base_models" {
  compartment_id = var.compartment_id
  display_name = "cohere.command-light"
}
`
	FineTuneDataBucketName = "fineTuneData"
	FineTuneDataObjectName = "uhc_data.jsonl"
	// At least 32 prompts are needed to pass API validation
	Prompts = `{\"prompt\": \"1\", \"completion\": \"one\"}\n{\"prompt\": \"2\", \"completion\": \"two\"}\n{\"prompt\": \"3\", \"completion\": \"three\"}\n{\"prompt\": \"4\", \"completion\": \"four\"}\n{\"prompt\": \"5\", \"completion\": \"five\"}\n{\"prompt\": \"6\", \"completion\": \"six\"}\n{\"prompt\": \"7\", \"completion\": \"seven\"}\n{\"prompt\": \"8\", \"completion\": \"eight\"}\n{\"prompt\": \"9\", \"completion\": \"nine\"}\n{\"prompt\": \"10\", \"completion\": \"ten\"}\n{\"prompt\": \"11\", \"completion\": \"eleven\"}\n{\"prompt\": \"12\", \"completion\": \"twelve\"}\n{\"prompt\": \"13\", \"completion\": \"thirteen\"}\n{\"prompt\": \"14\", \"completion\": \"fourteen\"}\n{\"prompt\": \"15\", \"completion\": \"fifteen\"}\n{\"prompt\": \"16\", \"completion\": \"sixteen\"}\n{\"prompt\": \"17\", \"completion\": \"seventeen\"}\n{\"prompt\": \"18\", \"completion\": \"eighteen\"}\n{\"prompt\": \"19\", \"completion\": \"nineteen\"}\n{\"prompt\": \"20\", \"completion\": \"twenty\"}\n{\"prompt\": \"21\", \"completion\": \"twenty-one\"}\n{\"prompt\": \"22\", \"completion\": \"twenty-two\"}\n{\"prompt\": \"23\", \"completion\": \"twenty-three\"}\n{\"prompt\": \"24\", \"completion\": \"twenty-four\"}\n{\"prompt\": \"25\", \"completion\": \"twenty-five\"}\n{\"prompt\": \"26\", \"completion\": \"twenty-six\"}\n{\"prompt\": \"27\", \"completion\": \"twenty-seven\"}\n{\"prompt\": \"28\", \"completion\": \"twenty-eight\"}\n{\"prompt\": \"29\", \"completion\": \"twenty-nine\"}\n{\"prompt\": \"30\", \"completion\": \"thirty\"}\n{\"prompt\": \"31\", \"completion\": \"thirty-one\"}\n{\"prompt\": \"32\", \"completion\": \"thirty-two\"}`

	// (Name + version) need to be unique
	StrFromTime = time.Now().UTC().Format("20060102150405")
	Name1       = "OldName" + StrFromTime
	Version1    = "v" + StrFromTime + ".1"
	Name2       = "NewName" + StrFromTime
	Version2    = "v" + StrFromTime + ".2"
	loraName1   = "LoraModel" + StrFromTime
	loraName2   = "LoraModelUpdated" + StrFromTime
)

// issue-routing-tag: generative_ai/default
func TestGenerativeAiModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGenerativeAiModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_generative_ai_model.test_model"
	datasourceName := "data.oci_generative_ai_models.test_models"
	singularDatasourceName := "data.oci_generative_ai_model.test_model"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+GenerativeAiModelResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create, GenerativeAiModelTfewRepresentation), "generativeai", "model", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelTfewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies,
		},
		// verify Create with optionals Lora
		{
			Config: config + compartmentIdVariableStr + GenerativeAiLlamaModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create, GenerativeAiModelLoraRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "capabilities"), - won't have it in state file
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", loraName1),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.lora_alpha", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.lora_dropout", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.lora_r", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "LORA_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "meta"),
				resource.TestCheckResourceAttr(resourceName, "version", Version1),

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

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiLlamaModelResourceDependencies,
		},
		// verify Create with optionals T-few
		{
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create, GenerativeAiModelTfewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "capabilities"), - won't have it in state file
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", Name1),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", Version1),

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
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(GenerativeAiModelTfewRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				// resource.TestCheckResourceAttrSet(resourceName, "capabilities"), - won't have it in state file
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", Name1),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", Version1),

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
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelTfewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				//resource.TestCheckResourceAttrSet(resourceName, "capabilities"), - won't have it in state file
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", Name2),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttrSet(resourceName, "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", Version2),

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
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_models", "test_models", acctest.Optional, acctest.Update, GenerativeAiModelDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelTfewRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "capability.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", Name2),
				resource.TestCheckResourceAttrSet(datasourceName, "id"),
				resource.TestCheckResourceAttr(datasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttr(datasourceName, "vendor", "vendor"),

				resource.TestCheckResourceAttr(datasourceName, "model_collection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "model_collection.0.items.#", "1"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelSingularDataSourceRepresentation) +
				compartmentIdVariableStr + GenerativeAiModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "model_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", Version2),
				resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", Name2),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.learning_rate", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.training_batch_size", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(singularDatasourceName, "fine_tune_details.0.training_dataset.0.object", FineTuneDataObjectName),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_metrics.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "vendor"),
				resource.TestCheckResourceAttr(singularDatasourceName, "version", Version2),
			),
		},
		// verify resource import
		{
			Config:            config + GenerativeAiModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"base_model_id",
				"description",
				"fine_tune_details",
				"freeform_tags",
				"is_long_term_supported",
				"lifecycle_details",
				"model_metrics",
				"previous_state",
				"system_tags",
				"time_created",
				"time_updated",
				"vendor",
			},
			ResourceName: resourceName,
		},
	})
}

func testAccCheckGenerativeAiModelDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).GenerativeAiClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_generative_ai_model" {
			noResourceFound = false
			request := oci_generative_ai.GetModelRequest{}

			tmp := rs.Primary.ID
			request.ModelId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")

			response, err := client.GetModel(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_generative_ai.ModelLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("GenerativeAiModel") {
		resource.AddTestSweepers("GenerativeAiModel", &resource.Sweeper{
			Name:         "GenerativeAiModel",
			Dependencies: acctest.DependencyGraph["model"],
			F:            sweepGenerativeAiModelResource,
		})
	}
}

func sweepGenerativeAiModelResource(compartment string) error {
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()
	modelIds, err := getGenerativeAiModelIds(compartment)
	if err != nil {
		return err
	}
	for _, modelId := range modelIds {
		if ok := acctest.SweeperDefaultResourceId[modelId]; !ok {
			deleteModelRequest := oci_generative_ai.DeleteModelRequest{}

			deleteModelRequest.ModelId = &modelId

			deleteModelRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "generative_ai")
			_, error := generativeAiClient.DeleteModel(context.Background(), deleteModelRequest)
			if error != nil {
				fmt.Printf("Error deleting Model %s %s, It is possible that the resource is already deleted. Please verify manually \n", modelId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &modelId, GenerativeAiModelSweepWaitCondition, time.Duration(3*time.Minute),
				GenerativeAiModelSweepResponseFetchOperation, "generative_ai", true)
		}
	}
	return nil
}

func getGenerativeAiModelIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "ModelId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	generativeAiClient := acctest.GetTestClients(&schema.ResourceData{}).GenerativeAiClient()

	listModelsRequest := oci_generative_ai.ListModelsRequest{}
	listModelsRequest.CompartmentId = &compartmentId
	listModelsRequest.LifecycleState = oci_generative_ai.ModelLifecycleStateActive
	listModelsResponse, err := generativeAiClient.ListModels(context.Background(), listModelsRequest)

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

func GenerativeAiModelSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is ACTIVE beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if modelResponse, ok := response.Response.(oci_generative_ai.GetModelResponse); ok {
		return modelResponse.LifecycleState != oci_generative_ai.ModelLifecycleStateDeleted
	}
	return false
}

func GenerativeAiModelSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.GenerativeAiClient().GetModel(context.Background(), oci_generative_ai.GetModelRequest{
		ModelId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
