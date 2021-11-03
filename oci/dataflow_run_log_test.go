// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	runLogSingularDataSourceRepresentation = map[string]interface{}{
		"name":   Representation{RepType: Required, Create: `${data.oci_dataflow_run_logs.test_run_logs.run_logs.0.name}`},
		"run_id": Representation{RepType: Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	runLogDataSourceRepresentation = map[string]interface{}{
		"run_id": Representation{RepType: Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	RunLogResourceConfig = GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", Required, Create, dataFlowApplicationRepresentation) +
		GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", Required, Create, RepresentationCopyWithNewProperties(invokeRunRepresentation, map[string]interface{}{
			"asynchronous": Representation{RepType: Required, Create: `false`},
		}))
)

// issue-routing-tag: dataflow/default
func TestDataflowRunLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowRunLogResource_basic")
	defer httpreplay.SaveScenario()

	config := ProviderTestConfig()

	compartmentId := GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	fileUri := GetEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)

	datasourceName := "data.oci_dataflow_run_logs.test_run_logs"
	singularDatasourceName := "data.oci_dataflow_run_log.test_run_log"

	SaveConfigContent("", "", "", t)

	ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_dataflow_run_logs", "test_run_logs", Required, Create, runLogDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + RunLogResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "run_id"),

				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.name"),
				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.run_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.size_in_bytes"),
				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.source"),
				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "run_logs.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				GenerateDataSourceFromRepresentationMap("oci_dataflow_run_logs", "test_run_logs", Required, Create, runLogDataSourceRepresentation) +
				GenerateDataSourceFromRepresentationMap("oci_dataflow_run_log", "test_run_log", Required, Create, runLogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + RunLogResourceConfig,
			Check: ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
			),
		},
	})
}
