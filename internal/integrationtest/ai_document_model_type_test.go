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
	AiDocumentModelTypeSingularDataSourceRepresentation = map[string]interface{}{
		"model_type":     acctest.Representation{RepType: acctest.Required, Create: `PRE_TRAINED_DOCUMENT_CLASSIFICATION`},
		"model_sub_type": acctest.Representation{RepType: acctest.Required, Create: `QR_BAR_CODE`},
	}

	AiDocumentModelTypeResourceConfig = ""
)

// issue-routing-tag: ai_document/default
func TestAiDocumentModelTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestAiDocumentModelTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_ai_document_model_type.test_model_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_ai_document_model_type", "test_model_type", acctest.Required, acctest.Create, AiDocumentModelTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AiDocumentModelTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(singularDatasourceName, "model_type", "PRE_TRAINED_DOCUMENT_CLASSIFICATION"),
				resource.TestCheckResourceAttr(singularDatasourceName, "model_sub_type", "QR_BAR_CODE"),
				resource.TestCheckResourceAttr(singularDatasourceName, "versions.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "capabilities", "{\"V1.0\":{\"capability\":{\"description\":{\"details\":[\"document classification model\"]},\"supportedLanguages\":{\"details\":[\"en\"]},\"userDisplayedVersion\":{\"details\":[\"V1.0\"]}}}}"),
			),
		},
	})
}
