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
	ApmTracesApmTracestraceSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"trace_key":     acctest.Representation{RepType: acctest.Required, Create: `${var.trace_key}`},
	}
)

// issue-routing-tag: apm_traces/default
func TestApmTracesTraceResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesTraceResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	//This is a manual test. It requires apm_domain_id and trace_key as environment variables.
	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")
	traceKey := utils.GetEnvSettingWithBlankDefault("trace_key")

	if apmDomainId == "" || traceKey == "" {
		t.Skip("Set apm_domain_id, trace_key to run this test")
	}

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	traceKeyVariableStr := fmt.Sprintf("variable \"trace_key\" { default = \"%s\" }\n", traceKey)

	singularDatasourceName := "data.oci_apm_traces_trace.test_trace"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + compartmentIdVariableStr + traceKeyVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_trace", "test_trace", acctest.Required, acctest.Create, ApmTracesApmTracestraceSingularDataSourceRepresentation),

			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_key", traceKey),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "error_span_count"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "is_fault"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "root_span_duration_in_ms"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "root_span_operation_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "root_span_service_name"),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_summaries.#", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "span_count"),
				resource.TestCheckResourceAttr(singularDatasourceName, "span_summary.#", "1"),
				// The number of spans would be different for different trace
				// the test with traceId 82876b361520c7dd55d4b0a2d4062fb0 returns 15
				resource.TestCheckResourceAttr(singularDatasourceName, "spans.#", "15"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_earliest_span_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_latest_span_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_root_span_ended"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_root_span_started"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_duration_in_ms"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_error_code"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_error_type"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_status"),
			),
		},
	})
}
