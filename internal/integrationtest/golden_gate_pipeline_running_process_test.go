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
	GoldenGatePipelineRunningProcessDataSourceRepresentation = map[string]interface{}{
		"pipeline_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_golden_gate_pipeline.test_pipeline.id}`},
	}

	GoldenGatePipelineRunningProcessResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_golden_gate_pipeline", "test_pipeline", acctest.Required, acctest.Create, GoldenGatePipelineRepresentation)
)

// issue-routing-tag: golden_gate/default
func TestGoldenGatePipelineRunningProcessResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestGoldenGatePipelineRunningProcessResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig() +
		makeVariableStr("source_connection_id", t) +
		makeVariableStr("target_connection_id", t)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_golden_gate_pipeline_running_processes.test_pipeline_running_processes"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_golden_gate_pipeline_running_processes", "test_pipeline_running_processes", acctest.Required, acctest.Create, GoldenGatePipelineRunningProcessDataSourceRepresentation) +
				compartmentIdVariableStr + GoldenGatePipelineRunningProcessResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "pipeline_running_process_collection.#"),
				resource.TestCheckResourceAttr(datasourceName, "pipeline_running_process_collection.0.items.#", "0"),
			),
		},
	})
}
