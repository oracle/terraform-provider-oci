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
	AiLanguageAiLanguageModelTypeSingularDataSourceRepresentation = map[string]interface{}{
		"model_type": acctest.Representation{RepType: acctest.Required, Create: `NAMED_ENTITY_RECOGNITION`},
	}

	AiLanguageModelTypeResourceConfig = ""
)

// issue-routing-tag: ai_language/default
func TestAiLanguageModelTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiLanguageModelTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_ai_language_model_type.test_model_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_language_model_type", "test_model_type", acctest.Required, acctest.Create, AiLanguageAiLanguageModelTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiLanguageModelTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "NAMED_ENTITY_RECOGNITION"),

				resource.TestCheckResourceAttr(singularDatasourceName, "versions.#", "3"),
			),
		},
	})
}
