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
	GoldenGateRecipeDataSourceRepresentation = map[string]interface{}{
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"display_name":   acctest.Representation{RepType: acctest.Optional, Create: `displayName`},
		"recipe_type":    acctest.Representation{RepType: acctest.Optional, Create: `ZERO_ETL`},
	}

	GoldenGateRecipeResourceConfig = ""
)

// issue-routing-tag: golden_gate/default
func TestGoldenGateRecipeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGateRecipeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_recipes.test_recipes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_recipes", "test_recipes", acctest.Required, acctest.Create, GoldenGateRecipeDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGateRecipeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "recipe_summary_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "recipe_summary_collection.0.items.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "recipe_summary_collection.0.items.0.recipe_type", "ZERO_ETL"),
			),
		},
	})
}
