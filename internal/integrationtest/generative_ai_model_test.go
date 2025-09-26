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
	GenerativeAiModelRequiredOnlyResource = GenerativeAiModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelRepresentation)

	GenerativeAiModelResourceConfig = GenerativeAiModelResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelRepresentation)

	GenerativeAiModelSingularDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_model.test_model.id}`},
	}

	GenerativeAiModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"capability":     acctest.Representation{RepType: acctest.Optional, Create: []string{`FINE_TUNE`}},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"id":             acctest.Representation{RepType: acctest.Optional, Create: `${oci_generative_ai_model.test_model.id}`},
		"state":          acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"vendor":         acctest.Representation{RepType: acctest.Optional, Create: `vendor`},
		"filter":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelDataSourceFilterRepresentation},
	}

	GenerativeAiModelDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_generative_ai_model.test_model.id}`}},
	}

	generativeAiVersion  = fmt.Sprintf("version-%d", time.Now().UnixNano())
	generativeAiVersion2 = fmt.Sprintf("version2-%d", time.Now().UnixNano())

	GenerativeAiModelRepresentation = map[string]interface{}{
		"base_model_id":     acctest.Representation{RepType: acctest.Required, Create: `${local.base_model_id}`},
		"compartment_id":    acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"fine_tune_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsRepresentation},
		//"defined_tags":      acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":   acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":  acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags": acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"vendor":        acctest.Representation{RepType: acctest.Optional, Create: `vendor`},
		"version":       acctest.Representation{RepType: acctest.Optional, Create: generativeAiVersion, Update: generativeAiVersion2},
	}
	GenerativeAiModelFineTuneDetailsRepresentation = map[string]interface{}{
		"dedicated_ai_cluster_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id}`},
		"training_dataset":        acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsTrainingDatasetRepresentation},
		"training_config":         acctest.RepresentationGroup{RepType: acctest.Required, Group: GenerativeAiModelFineTuneDetailsTrainingConfigRepresentation},
	}
	GenerativeAiModelFineTuneDetailsTrainingDatasetRepresentation = map[string]interface{}{
		"bucket":       acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"dataset_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE`},
		"namespace":    acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":       acctest.Representation{RepType: acctest.Required, Create: `object.jsonl`},
	}
	GenerativeAiModelFineTuneDetailsTrainingConfigRepresentation = map[string]interface{}{
		"training_config_type":                acctest.Representation{RepType: acctest.Required, Create: `TFEW_TRAINING_CONFIG`},
		"early_stopping_patience":             acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"early_stopping_threshold":            acctest.Representation{RepType: acctest.Optional, Create: `0.005`},
		"learning_rate":                       acctest.Representation{RepType: acctest.Optional, Create: `0.001`},
		"log_model_metrics_interval_in_steps": acctest.Representation{RepType: acctest.Optional, Create: `100`},
		"total_training_epochs":               acctest.Representation{RepType: acctest.Optional, Create: `2`},
		"training_batch_size":                 acctest.Representation{RepType: acctest.Optional, Create: `32`},
	}

	BaseModelDependencies = `
							# Pulls available base models dynamically from your compartment
							data "oci_generative_ai_models" "base_models" {
							  compartment_id = var.compartment_id
							}
							
							locals {
							 
							  filtered_base_models = [
								for item in data.oci_generative_ai_models.base_models.model_collection[0].items : item
								if (
									contains(item.capabilities, "FINE_TUNE")
									&& (item.display_name == "cohere.command-light")									
								  )	
							  ]
							
							  # Choose the first matching model
							  base_model_id = (
								length(local.filtered_base_models) > 0 ?
								local.filtered_base_models[0].id :
								""
							  )
							}
							`

	Prompts = `{\"prompt\": \"1\", \"completion\": \"one\"}\n{\"prompt\": \"2\", \"completion\": \"two\"}\n{\"prompt\": \"3\", \"completion\": \"three\"}\n{\"prompt\": \"4\", \"completion\": \"four\"}\n{\"prompt\": \"5\", \"completion\": \"five\"}\n{\"prompt\": \"6\", \"completion\": \"six\"}\n{\"prompt\": \"7\", \"completion\": \"seven\"}\n{\"prompt\": \"8\", \"completion\": \"eight\"}\n{\"prompt\": \"9\", \"completion\": \"nine\"}\n{\"prompt\": \"10\", \"completion\": \"ten\"}\n{\"prompt\": \"11\", \"completion\": \"eleven\"}\n{\"prompt\": \"12\", \"completion\": \"twelve\"}\n{\"prompt\": \"13\", \"completion\": \"thirteen\"}\n{\"prompt\": \"14\", \"completion\": \"fourteen\"}\n{\"prompt\": \"15\", \"completion\": \"fifteen\"}\n{\"prompt\": \"16\", \"completion\": \"sixteen\"}\n{\"prompt\": \"17\", \"completion\": \"seventeen\"}\n{\"prompt\": \"18\", \"completion\": \"eighteen\"}\n{\"prompt\": \"19\", \"completion\": \"nineteen\"}\n{\"prompt\": \"20\", \"completion\": \"twenty\"}\n{\"prompt\": \"21\", \"completion\": \"twenty-one\"}\n{\"prompt\": \"22\", \"completion\": \"twenty-two\"}\n{\"prompt\": \"23\", \"completion\": \"twenty-three\"}\n{\"prompt\": \"24\", \"completion\": \"twenty-four\"}\n{\"prompt\": \"25\", \"completion\": \"twenty-five\"}\n{\"prompt\": \"26\", \"completion\": \"twenty-six\"}\n{\"prompt\": \"27\", \"completion\": \"twenty-seven\"}\n{\"prompt\": \"28\", \"completion\": \"twenty-eight\"}\n{\"prompt\": \"29\", \"completion\": \"twenty-nine\"}\n{\"prompt\": \"30\", \"completion\": \"thirty\"}\n{\"prompt\": \"31\", \"completion\": \"thirty-one\"}\n{\"prompt\": \"32\", \"completion\": \"thirty-two\"}`

	FineTuneDataObjectStorageBucketRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"name":                  acctest.Representation{RepType: acctest.Required, Create: FineTuneDataBucketName, Update: FineTuneDataBucketName2},
		"namespace":             acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"access_type":           acctest.Representation{RepType: acctest.Optional, Create: `NoPublicAccess`, Update: `ObjectRead`},
		"auto_tiering":          acctest.Representation{RepType: acctest.Optional, Create: `Disabled`, Update: `InfrequentAccess`},
		"defined_tags":          acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"kms_key_id":            acctest.Representation{RepType: acctest.Optional, Create: `${lookup(data.oci_kms_keys.test_keys_dependency.keys[0], "id")}`},
		"metadata":              acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"content-type": "text/plain"}, Update: map[string]string{"content-type": "text/xml"}},
		"object_events_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"storage_tier":          acctest.Representation{RepType: acctest.Optional, Create: `Standard`},
		"versioning":            acctest.Representation{RepType: acctest.Optional, Create: `Enabled`, Update: `Disabled`},
	}

	FineTuneDataObjectStorageObjectRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `${oci_objectstorage_bucket.test_bucket.name}`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: FineTuneDataObjectName},
		"content":   acctest.Representation{RepType: acctest.Optional, Create: Prompts},
	}

	GenerativeAINamespaceSingularDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}

	FineTuneDataBucketName  = utils.RandomStringOrHttpReplayValue(32, utils.Charset, "bucket")
	FineTuneDataBucketName2 = FineTuneDataBucketName + "2"
	FineTuneDataObjectName  = "object.jsonl"

	GenerativeAiFineTuningDedicatedAiClusterRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"type":           acctest.Representation{RepType: acctest.Required, Create: `FINE_TUNING`},
		"unit_count":     acctest.Representation{RepType: acctest.Required, Create: `2`},
		"unit_shape":     acctest.Representation{RepType: acctest.Required, Create: `SMALL_COHERE`},
		"defined_tags":   acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":    acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":  acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
	}

	GenerativeAiModelResourceDependencies = BaseModelDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_dedicated_ai_cluster", "test_dedicated_ai_cluster", acctest.Required, acctest.Create, GenerativeAiFineTuningDedicatedAiClusterRepresentation) +
		DefinedTagsDependencies +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, GenerativeAINamespaceSingularDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, FineTuneDataObjectStorageBucketRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_object", "fine_tune_data", acctest.Optional, acctest.Create, FineTuneDataObjectStorageObjectRepresentation)
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
		acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create, GenerativeAiModelRepresentation), "generativeai", "model", t)

	acctest.ResourceTest(t, testAccCheckGenerativeAiModelDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Required, acctest.Create, GenerativeAiModelRepresentation),
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
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + GenerativeAiModelResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Create, GenerativeAiModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "0.005"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "0.001"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "100"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "32"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", "object.jsonl"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiVersion),

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
					acctest.RepresentationCopyWithNewProperties(GenerativeAiModelRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "0.005"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "0.001"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "100"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "32"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", "object.jsonl"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiVersion),

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
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "base_model_id"),
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.dedicated_ai_cluster_id"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_patience", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.early_stopping_threshold", "0.005"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.learning_rate", "0.001"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.log_model_metrics_interval_in_steps", "100"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.total_training_epochs", "2"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_batch_size", "32"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_config.0.training_config_type", "TFEW_TRAINING_CONFIG"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.bucket", FineTuneDataBucketName),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.dataset_type", "OBJECT_STORAGE"),
				resource.TestCheckResourceAttrSet(resourceName, "fine_tune_details.0.training_dataset.0.namespace"),
				resource.TestCheckResourceAttr(resourceName, "fine_tune_details.0.training_dataset.0.object", "object.jsonl"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttr(resourceName, "vendor", "vendor"),
				resource.TestCheckResourceAttr(resourceName, "version", generativeAiVersion2),

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
				acctest.GenerateResourceFromRepresentationMap("oci_generative_ai_model", "test_model", acctest.Optional, acctest.Update, GenerativeAiModelRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "capability.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "display_name", "displayName2"),
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
			),
		},
		// verify resource import
		{
			Config:            config + GenerativeAiModelRequiredOnlyResource,
			ImportState:       true,
			ImportStateVerify: true,
			ImportStateVerifyIgnore: []string{
				"base_model_id",
				"capabilities",
				"compartment_id",
				"defined_tags",
				"description",
				"display_name",
				"fine_tune_details",
				"freeform_tags",
				"id",
				"is_long_term_supported",
				"lifecycle_details",
				"model_metrics",
				"previous_state",
				"state",
				"system_tags",
				"time_created",
				"time_deprecated",
				"time_updated",
				"type",
				"vendor",
				"version",
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
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
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
