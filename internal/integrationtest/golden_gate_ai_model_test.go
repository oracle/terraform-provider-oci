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
	GoldenGateAiModelDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"provider_type":  acctest.Representation{RepType: acctest.Required, Create: `OPEN_AI`},
	}

	GoldenGateAiModelResourceConfig = ""
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateAiModelResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateAiModelResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_ai_models.test_ai_models"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_ai_models", "test_ai_models", acctest.Required, acctest.Create, GoldenGateAiModelDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGateAiModelResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "provider_type", "OPEN_AI"),

				resource.TestCheckResourceAttrSet(datasourceName, "ai_model_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "ai_model_collection.0.items.#", "2"),
			),
		},
	})
}
