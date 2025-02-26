// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiDocumentProcessorJobDependencies = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_project", "test_project", acctest.Required, acctest.Create, AiDocumentProjectRepresentation) + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_model", "test_model", acctest.Required, acctest.Create, AiDocumentModelRepresentation2) +
		DefinedTagsDependencies

	AiDocumentProcessorJobRequiredOnlyResource = AiDocumentProcessorJobDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Required, acctest.Create, AiDocumentProcessorJobRepresentation)

	AiDocumentProcessorJobResourceConfig = AiDocumentProcessorJobDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Update, AiDocumentProcessorJobRepresentation)

	AiDocumentAiDocumentProcessorJobSingularDataSourceRepresentation = map[string]interface{}{
		"processor_job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_document_processor_job.test_processor_job.id}`},
	}

	AiDocumentProcessorJobRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"input_location":   acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobInputLocationRepresentation},
		"output_location":  acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobOutputLocationRepresentation},
		"processor_config": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobProcessorConfigRepresentation},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
	}
	AiDocumentProcessorJobInputLocationRepresentation = map[string]interface{}{
		"page_range":       acctest.Representation{RepType: acctest.Optional, Create: []string{}},
		"source_type":      acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_LOCATIONS`},
		"data":             acctest.Representation{RepType: acctest.Optional, Create: ""},
		"object_locations": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobInputLocationObjectLocationsRepresentation},
	}
	AiDocumentProcessorJobOutputLocationRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `canary_test`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `axylfvgphoea`},
		"prefix":    acctest.Representation{RepType: acctest.Required, Create: `test`},
	}
	AiDocumentProcessorJobProcessorConfigRepresentation = map[string]interface{}{
		"processor_type":        acctest.Representation{RepType: acctest.Required, Create: `GENERAL`},
		"document_type":         acctest.Representation{RepType: acctest.Optional, Create: `RECEIPT`},
		"features":              acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobProcessorConfigFeaturesRepresentation},
		"is_zip_output_enabled": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"language":              acctest.Representation{RepType: acctest.Optional, Create: `ENG`},
		"model_id":              acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
		"normalization_fields":  acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobProcessorConfigNormalizationFieldsRepresentation},
	}
	AiDocumentProcessorJobInputLocationObjectLocationsRepresentation = map[string]interface{}{
		"bucket":     acctest.Representation{RepType: acctest.Required, Create: `canary_test`},
		"namespace":  acctest.Representation{RepType: acctest.Required, Create: `axylfvgphoea`},
		"object":     acctest.Representation{RepType: acctest.Required, Create: `key_value_receipt.png`},
		"page_range": acctest.Representation{RepType: acctest.Optional, Create: []string{`1`}},
	}
	AiDocumentProcessorJobProcessorConfigFeaturesRepresentation = map[string]interface{}{
		"feature_type":             acctest.Representation{RepType: acctest.Required, Create: `KEY_VALUE_EXTRACTION`},
		"generate_searchable_pdf":  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_results":              acctest.Representation{RepType: acctest.Optional, Create: `0`},
		"model_id":                 acctest.Representation{RepType: acctest.Optional, Create: `${oci_ai_document_model.test_model.id}`},
		"selection_mark_detection": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"tenancy_id":               acctest.Representation{RepType: acctest.Optional, Create: `${var.compartment_id}`},
	}
	AiDocumentProcessorJobProcessorConfigNormalizationFieldsRepresentation = map[string]interface{}{
		"map": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobProcessorConfigNormalizationFieldsMapRepresentation},
	}
	AiDocumentProcessorJobProcessorConfigNormalizationFieldsMapRepresentation = map[string]interface{}{
		"normalization_type": acctest.Representation{RepType: acctest.Required, Create: `normalization_type_sample_val`},
	}
)

// issue-routing-tag: ai_document/default
func TestAiDocumentProcessorJobResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiDocumentProcessorJobResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_ai_document_processor_job.test_processor_job"

	singularDatasourceName := "data.oci_ai_document_processor_job.test_processor_job"

	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AiDocumentProcessorJobDependencies+acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Create, AiDocumentProcessorJobRepresentation), "aidocument", "processorJob", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AiDocumentProcessorJobDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Required, acctest.Create, AiDocumentProcessorJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.object", "key_value_receipt.png"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.prefix", "test"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.normalization_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.normalization_fields.0.map.0.normalization_type", "normalization_type_sample_val"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.feature_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.processor_type", "GENERAL"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AiDocumentProcessorJobDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + AiDocumentProcessorJobDependencies + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Create, AiDocumentProcessorJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.page_range.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.page_range.#", "0"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.object", "key_value_receipt.png"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.prefix", "test"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.feature_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.generate_searchable_pdf", "false"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.max_results", "0"),
				resource.TestCheckResourceAttrSet(resourceName, "processor_config.0.features.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.selection_mark_detection", "false"),
				resource.TestCheckResourceAttrSet(resourceName, "processor_config.0.features.0.tenancy_id"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.is_zip_output_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.language", "ENG"),
				resource.TestCheckResourceAttrSet(resourceName, "processor_config.0.model_id"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.normalization_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.normalization_fields.0.map.0.normalization_type", "normalization_type_sample_val"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.is_zip_output_enabled", "false"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.processor_type", "GENERAL"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_accepted"),
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Required, acctest.Create, AiDocumentAiDocumentProcessorJobSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiDocumentProcessorJobResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "processor_job_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.page_range.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.page_range.#", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.object", "key_value_receipt.png"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.bucket", "canary_test"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.namespace", "axylfvgphoea"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.prefix", "test"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "percent_complete"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.feature_type", "KEY_VALUE_EXTRACTION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.generate_searchable_pdf", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.max_results", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.selection_mark_detection", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.is_zip_output_enabled", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.language", "ENG"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.normalization_fields.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.normalization_fields.0.map.0.normalization_type", "normalization_type_sample_val"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.processor_type", "GENERAL"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_accepted"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_finished"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
			),
		},
		// verify resource import
		{
			Config:                  config + AiDocumentProcessorJobRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
