// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	pringProbeStartTime                                             = time.Now()
	HealthChecksHealthChecksPingProbeResultDataSourceRepresentation = map[string]interface{}{
		"probe_configuration_id":              acctest.Representation{RepType: acctest.Required, Create: `${oci_health_checks_ping_monitor.test_ping_monitor.id}`},
		"start_time_greater_than_or_equal_to": acctest.Representation{RepType: acctest.Optional, Create: strconv.FormatInt(pringProbeStartTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"start_time_less_than_or_equal_to":    acctest.Representation{RepType: acctest.Optional, Create: strconv.FormatInt(pringProbeStartTime.Add(30*time.Minute).UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"target":                              acctest.Representation{RepType: acctest.Optional, Create: `www.oracle.com`},
	}

	HealthChecksPingProbeResultResourceConfig = acctest.GenerateResourceFromRepresentationMap("oci_health_checks_ping_monitor", "test_ping_monitor", acctest.Required, acctest.Create, HealthChecksPingMonitorRepresentation)
)

// issue-routing-tag: health_checks/default
func TestHealthChecksPingProbeResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksPingProbeResultResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_ping_probe_results.test_ping_probe_results"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config + compartmentIdVariableStr + HealthChecksPingProbeResultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				func(s *terraform.State) (err error) {
					if httpreplay.ShouldRetryImmediately() {
						time.Sleep(10 * time.Millisecond)
					} else {
						time.Sleep(5 * time.Minute)
					}
					return nil
				},
			),
		},
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_health_checks_ping_probe_results", "test_ping_probe_results", acctest.Optional, acctest.Create, HealthChecksHealthChecksPingProbeResultDataSourceRepresentation) +
				compartmentIdVariableStr + HealthChecksPingProbeResultResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(datasourceName, "probe_configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_time_greater_than_or_equal_to"),
				resource.TestCheckResourceAttrSet(datasourceName, "start_time_less_than_or_equal_to"),
				resource.TestCheckResourceAttr(datasourceName, "target", "www.oracle.com"),

				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.#"),
				resource.TestCheckResourceAttr(datasourceName, "ping_probe_results.0.connection.#", "1"),
				resource.TestCheckResourceAttr(datasourceName, "ping_probe_results.0.dns.#", "1"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.domain_lookup_end"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.domain_lookup_start"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.error_category"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.icmp_code"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.is_healthy"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.is_timed_out"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.key"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.latency_in_ms"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.probe_configuration_id"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.protocol"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.start_time"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.target"),
				resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.vantage_point_name"),
			),
		},
	})
}
