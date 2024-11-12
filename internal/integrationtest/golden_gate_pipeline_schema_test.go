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
	GoldenGatePipelineSchemaDataSourceRepresentation = map[string]interface{}{
		"pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_pipeline.test_pipeline.id}`},
	}

	GoldenGatePipelineSchemaResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGatePipelineSchemaResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGatePipelineSchemaResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr("source_connection_id", t) +
		makeVariableStr("target_connection_id", t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_pipeline_schemas.test_pipeline_schemas"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_pipeline_schemas", "test_pipeline_schemas", acctest.Required, acctest.Create, GoldenGatePipelineSchemaDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGatePipelineSchemaResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_schema_collection.#"),
				//resource.TestCheckResourceAttr(datasourceName, "pipeline_schema_collection.0.items.#", "1"),
			),
		},
	})
}
