// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	AiDocumentProcessorJobRequiredOnlyResource = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Required, acctest.Create, AiDocumentProcessorJobRepresentation)

	AiDocumentProcessorJobResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Update, AiDocumentProcessorJobRepresentation)

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
		"source_type": acctest.Representation{RepType: acctest.Required, Create: `OBJECT_STORAGE_LOCATIONS`},
		//"data":             acctest.Representation{RepType: acctest.Optional, Create: `data`},
		"object_locations": acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobInputLocationObjectLocationsRepresentation},
	}
	AiDocumentProcessorJobOutputLocationRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `tf_test_bucket`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `axgexwaxnm7k`},
		"prefix":    acctest.Representation{RepType: acctest.Required, Create: `response`},
	}
	AiDocumentProcessorJobProcessorConfigRepresentation = map[string]interface{}{
		"features":       acctest.RepresentationGroup{RepType: acctest.Required, Group: AiDocumentProcessorJobProcessorConfigFeaturesRepresentation},
		"processor_type": acctest.Representation{RepType: acctest.Required, Create: `GENERAL`},
	}
	AiDocumentProcessorJobInputLocationObjectLocationsRepresentation = map[string]interface{}{
		"bucket":    acctest.Representation{RepType: acctest.Required, Create: `tf_test_bucket`},
		"namespace": acctest.Representation{RepType: acctest.Required, Create: `axgexwaxnm7k`},
		"object":    acctest.Representation{RepType: acctest.Required, Create: `amazon_inv.pdf`},
	}
	AiDocumentProcessorJobProcessorConfigFeaturesRepresentation = map[string]interface{}{
		"feature_type":            acctest.Representation{RepType: acctest.Required, Create: `DOCUMENT_CLASSIFICATION`},
		"generate_searchable_pdf": acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"max_results":             acctest.Representation{RepType: acctest.Optional, Create: `10`},
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
	acctest.SaveConfigContent(config+compartmentIdVariableStr+acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Create, AiDocumentProcessorJobRepresentation), "aidocument", "processorJob", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Required, acctest.Create, AiDocumentProcessorJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.object", "amazon_inv.pdf"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.prefix", "response"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.feature_type", "DOCUMENT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.processor_type", "GENERAL"),
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + acctest.GenerateResourceFromRepresentationMap("oci_ai_document_processor_job", "test_processor_job", acctest.Optional, acctest.Create, AiDocumentProcessorJobRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "input_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.object_locations.0.object", "amazon_inv.pdf"),
				resource.TestCheckResourceAttr(resourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(resourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(resourceName, "output_location.0.prefix", "response"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.feature_type", "DOCUMENT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.generate_searchable_pdf", "false"),
				resource.TestCheckResourceAttr(resourceName, "processor_config.0.features.0.max_results", "10"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.object_locations.0.object", "amazon_inv.pdf"),
				resource.TestCheckResourceAttr(singularDatasourceName, "input_location.0.source_type", "OBJECT_STORAGE_LOCATIONS"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.bucket", "tf_test_bucket"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.namespace", "axgexwaxnm7k"),
				resource.TestCheckResourceAttr(singularDatasourceName, "output_location.0.prefix", "response"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "percent_complete"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.feature_type", "DOCUMENT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.generate_searchable_pdf", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.features.0.max_results", "10"),
				resource.TestCheckResourceAttr(singularDatasourceName, "processor_config.0.is_zip_output_enabled", "false"),
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
