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
	AiLanguageAiLanguageModelEvaluationResultDataSourceRepresentation = map[string]interface{}{
		"model_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_ai_language_model.test_model.id}`},
	}

	AiLanguageModelEvaluationResultResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_ai_language_model", "test_model", acctest.Required, acctest.Create, AiLanguageModelRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_ai_language_project", "test_project", acctest.Required, acctest.Create, AiLanguageProjectRepresentation)
	// +
	// acctest.GenerateResourceFromRepresentationMap("oci_data_labeling_service_dataset", "test_dataset", acctest.Required, acctest.Create, datasetRepresentation) +
	// acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Required, acctest.Create, bucketRepresentation) +
	// GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", Required, Create, namespaceSingularDataSourceRepresentation)
)

// issue-routing-tag: ai_language/default
func TestAiLanguageModelEvaluationResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiLanguageModelEvaluationResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_ai_language_model_evaluation_results.test_model_evaluation_results"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_model_evaluation_results", "test_model_evaluation_results", acctest.Required, acctest.Create, AiLanguageAiLanguageModelEvaluationResultDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageModelEvaluationResultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				// resource.TestCheckResourceAttrSet(datasourceName, "model_id"),

				resource.TestCheckResourceAttr(datasourceName, "evaluation_result_collection.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "evaluation_result_collection.0.items.#"),
			),
		},
	})
}
