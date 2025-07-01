// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	ApmTracesLogSingularDataSourceRepresentation = map[string]interface{}{
		"apm_domain_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.apm_domain_id}`},
		"log_key":                  acctest.Representation{RepType: acctest.Required, Create: `${var.log_key}`},
		"time_log_ended_less_than": acctest.Representation{RepType: acctest.Required, Create: `${var.time_log_ended_less_than}`},
		"time_log_started_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Required, Create: `${var.time_log_started_greater_than_or_equal_to}`},
	}

	//ApmTracesLogResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_apm_apm_domain", "test_apm_domain", acctest.Required, acctest.Create, apmDomainRepresentation)
)

// issue-routing-tag: apm_traces/default
func TestApmTracesLogResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestApmTracesLogResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	apmDomainId := utils.GetEnvSettingWithBlankDefault("apm_domain_id")
	logKey := utils.GetEnvSettingWithBlankDefault("log_key")
	timeLogStartedGreaterThanOrEqualTo := utils.GetEnvSettingWithBlankDefault("time_log_started_greater_than_or_equal_to")
	timeLogEndedLessThan := utils.GetEnvSettingWithBlankDefault("time_log_ended_less_than")

	if apmDomainId == "" || logKey == "" {
		t.Skip("Set apm_domain_id, logKey to run this test")
	}

	apmDomainIdVariableStr := fmt.Sprintf("variable \"apm_domain_id\" { default = \"%s\" }\n", apmDomainId)
	logKeyVariableStr := fmt.Sprintf("variable \"log_key\" { default = \"%s\" }\n", logKey)
	timeLogStartedGreaterThanOrEqualToVariableStr := fmt.Sprintf("variable \"time_log_started_greater_than_or_equal_to\" { default = \"%s\" }\n", timeLogStartedGreaterThanOrEqualTo)
	timeLogEndedLessThanVariableStr := fmt.Sprintf("variable \"time_log_ended_less_than\" { default = \"%s\" }\n", timeLogEndedLessThan)

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_apm_traces_log.test_log"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config + apmDomainIdVariableStr + logKeyVariableStr + timeLogStartedGreaterThanOrEqualToVariableStr + timeLogEndedLessThanVariableStr +
				acctest.GenerateDataSourceFromRepresentationMap("oci_apm_traces_log", "test_log", acctest.Required, acctest.Create, ApmTracesLogSingularDataSourceRepresentation) +
				compartmentIdVariableStr,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "apm_domain_id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "log_key", logKey),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_log_ended_less_than"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_log_started_greater_than_or_equal_to"),

				resource.TestCheckResourceAttr(singularDatasourceName, "attributes.#", "7"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "body"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "event_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "overflow_attributes"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity_number"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "severity_text"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "span_key"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_observed"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "timestamp"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_flags"),
				//resource.TestCheckResourceAttrSet(singularDatasourceName, "trace_key"),
			),
		},
	})
}
