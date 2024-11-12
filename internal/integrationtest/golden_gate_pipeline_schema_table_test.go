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
	GoldenGatePipelineSchemaTableDataSourceRepresentation = map[string]interface{}{
		"pipeline_id":        acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_pipeline.test_pipeline.id}`},
		"source_schema_name": acctest.Representation{RepType: acctest.Required, Create: `${var.source_schema}`},
		"target_schema_name": acctest.Representation{RepType: acctest.Required, Create: `${var.target_schema}`},
	}

	GoldenGatePipelineSchemaTableResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGatePipelineSchemaTableResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGatePipelineSchemaTableResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr("source_connection_id", t) +
		makeVariableStr("target_connection_id", t) +
		makeVariableStr("source_schema", t) +
		makeVariableStr("target_schema", t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	sourceSchema := utils.GetEnvSettingWithBlankDefault("source_schema")
	targetSchema := utils.GetEnvSettingWithBlankDefault("target_schema")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_pipeline_schema_tables.test_pipeline_schema_tables"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_pipeline_schema_tables", "test_pipeline_schema_tables", acctest.Required, acctest.Create, GoldenGatePipelineSchemaTableDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGatePipelineSchemaTableResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_id"),
				resource.TestCheckResourceAttr(datasourceName, "source_schema_name", sourceSchema),
				resource.TestCheckResourceAttr(datasourceName, "target_schema_name", targetSchema),

				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_schema_table_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_schema_table_collection.0.items.#", "2"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_schema_table_collection.0.items.0.source_table_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_schema_table_collection.0.items.0.target_table_name"),
			),
		},
	})
}
