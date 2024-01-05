// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
)

var (
	DataflowDataflowRunLogSingularDataSourceRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `${data.oci_dataflow_run_logs.test_run_logs.run_logs.0.name}`},
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	DataflowDataflowRunLogDataSourceRepresentation = map[string]interface{}{
		"run_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_dataflow_invoke_run.test_invoke_run.id}`},
	}

	DataflowRunLogResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_dataflow_application", "test_application", acctest.Required, acctest.Create, DataflowApplicationRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_dataflow_invoke_run", "test_invoke_run", acctest.Required, acctest.Create, acctest.RepresentationCopyWithNewProperties(DataflowInvokeRunRepresentation, map[string]interface{}{
			"asynchronous": acctest.Representation{RepType: acctest.Required, Create: `false`},
		}))
)

// issue-routing-tag: dataflow/default
func TestDataflowRunLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataflowRunLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	fileUri := utils.GetEnvSettingWithBlankDefault("dataflow_file_uri")
	fileUriVariableStr := fmt.Sprintf("variable \"dataflow_file_uri\" { default = \"%s\" }\n", fileUri)

	datasourceName := "data.oci_dataflow_run_logs.test_run_logs"
	singularDatasourceName := "data.oci_dataflow_run_log.test_run_log"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_run_logs", "test_run_logs", acctest.Required, acctest.Create, DataflowDataflowRunLogDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + DataflowRunLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_run_logs", "test_run_logs", acctest.Required, acctest.Create, DataflowDataflowRunLogDataSourceRepresentation) +
				acctest.GenerateDataSourceFromRepresentationMap("oci_dataflow_run_log", "test_run_log", acctest.Required, acctest.Create, DataflowDataflowRunLogSingularDataSourceRepresentation) +
				compartmentIdVariableStr + fileUriVariableStr + DataflowRunLogResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "run_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "content_type"),
			),
		},
	})
}
