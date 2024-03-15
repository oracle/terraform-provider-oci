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
	ApmTracesTraceAggregatedSnapshotDataSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id": acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"trace_key":     acctest.Representation{RepType: acctest.Required, Create: `${var.trace_key}`},
		"server_name":   acctest.Representation{RepType: acctest.Optional, Create: `${var.server_name}`},
		"service_name":  acctest.Representation{RepType: acctest.Optional, Create: `${var.service_name}`},
		"span_key":      acctest.Representation{RepType: acctest.Optional, Create: `${var.span_key}`},
		"span_name":     acctest.Representation{RepType: acctest.Optional, Create: `${var.span_name}`},
	}
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
	serverName := utils.GetEnvSettingWithBlankDefault("server_name")
	serviceName := utils.GetEnvSettingWithBlankDefault("service_name")
	spanKey := utils.GetEnvSettingWithBlankDefault("span_key")
	spanName := utils.GetEnvSettingWithBlankDefault("span_name")

	if apmDomainId == "" || traceKey == "" {
		t.Skip("Set apm_domain_id, trace_key to run this test")
	}

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	traceKeyVariableStr := fmt.Sprintf("variable \"trace_key\" { default = \"%s\" }\n", traceKey)
	serverNameVariableStr := fmt.Sprintf("variable \"server_name\" { default = \"%s\" }\n", serverName)
	serviceNameVariableStr := fmt.Sprintf("variable \"service_name\" { default = \"%s\" }\n", serviceName)
	spanKeyVariableStr := fmt.Sprintf("variable \"span_key\" { default = \"%s\" }\n", spanKey)
	spanNameVariableStr := fmt.Sprintf("variable \"span_name\" { default = \"%s\" }\n", spanName)

	singularDatasourceName := "data.oci_apm_traces_trace_aggregated_snapshot_data.test_trace_aggregated_snapshot_data"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + traceKeyVariableStr + serverNameVariableStr + serviceNameVariableStr + spanKeyVariableStr + spanNameVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_trace_aggregated_snapshot_data", "test_trace_aggregated_snapshot_data", acctest.Optional, acctest.Create, ApmTracesTraceAggregatedSnapshotDataSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "trace_key", traceKey),
				resource.TestCheckResourceAttr(singularDatasourceName, "span_key", spanKey),
				resource.TestCheckResourceAttr(singularDatasourceName, "service_name", serviceName),
				resource.TestCheckResourceAttr(singularDatasourceName, "span_name", spanName),
				resource.TestCheckResourceAttr(singularDatasourceName, "server_name", serverName),
				resource.TestCheckResourceAttr(singularDatasourceName, "details.#", "2"),
			),
		},
	})
}
