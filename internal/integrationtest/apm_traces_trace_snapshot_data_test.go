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
	ApmTracesApmTracestraceSnapshotDataSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"trace_key":     acctest.Representation{RepType: acctest.Required, Create: `${var.trace_key}`},
		"is_summarized": acctest.Representation{RepType: acctest.Optional, Create: `${var.isSummarized}`},
		"snapshot_time": acctest.Representation{RepType: acctest.Optional, Create: `${var.snapshotTime}`},
		"thread_id":     acctest.Representation{RepType: acctest.Optional, Create: `${var.thread.id}`},
	}
)

// issue-routing-tag: apm_traces/default
func TestApmTracesTraceSnapshotDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesTraceSnapshotDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//This is a manual test. It requires apm_domain_id and trace_key as environment variables.
	// Optional environment variables are thread_id, is_summarized and snapshot_time.
	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")
	traceKey := utils.GetEnvSettingWithBlankDefault("trace_key")
	threadId := utils.GetEnvSettingWithBlankDefault("thread_id")
	isSummarized := utils.GetEnvSettingWithBlankDefault("is_summarized")
	snapshotTime := utils.GetEnvSettingWithBlankDefault("snapshot_time")

	if apmDomainId == "" || traceKey == "" {
		t.Skip("Set apm_domain_id and trace_key to run this test")
	}

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	traceKeyVariableStr := fmt.Sprintf("variable \"trace_key\" { default = \"%s\" }\n", traceKey)
	threadIdVariableStr := fmt.Sprintf("variable \"thread_id\" { default = \"%s\" }\n", threadId)
	isSummarizedStr := fmt.Sprintf("variable \"is_summarized\" { default = \"%s\" }\n", isSummarized)
	snapshotTimeStr := fmt.Sprintf("variable \"snapshot_time\" { default = \"%s\" }\n", snapshotTime)

	singularDatasourceName := "data.oci_apm_traces_trace_snapshot_data.test_trace_snapshot_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + traceKeyVariableStr + compartmentIdVariableStr + threadIdVariableStr + isSummarizedStr + snapshotTimeStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_trace_snapshot_data", "test_trace_snapshot_data", acctest.Required, acctest.Create, ApmTracesApmTracestraceSnapshotDataSingularDataSourceRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_key", traceKey),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_started"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_snapshot_details.#", "2"),
			),
		},
	})
}
