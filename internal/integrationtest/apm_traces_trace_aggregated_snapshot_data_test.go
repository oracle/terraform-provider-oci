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
	ApmTracesApmTracestraceAggregatedSnapshotDataSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"trace_key":     acctest.Representation{RepType: acctest.Required, Create: `${var.trace_key}`},
	}

	//TraceAggregatedSnapshotDataResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_traces/default
func TestApmTracesTraceAggregatedSnapshotDataResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesTraceAggregatedSnapshotDataResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	//This is a manual test. It requires apm_domain_id and trace_key as environment variables.
	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")
	traceKey := utils.GetEnvSettingWithBlankDefault("trace_key")

	if apmDomainId == "" || traceKey == "" {
		t.Skip("Set apm_domain_id, trace_key to run this test")
	}

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	traceKeyVariableStr := fmt.Sprintf("variable \"trace_key\" { default = \"%s\" }\n", traceKey)

	singularDatasourceName := "data.oci_apm_traces_trace_aggregated_snapshot_data.test_trace_aggregated_snapshot_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + traceKeyVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_trace_aggregated_snapshot_data", "test_trace_aggregated_snapshot_data", acctest.Required, acctest.Create, ApmTracesApmTracestraceAggregatedSnapshotDataSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_key", traceKey),

				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "2"),
			),
		},
	})
}
