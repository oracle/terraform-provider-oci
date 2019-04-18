// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	pringProbeStartTime                     = time.Now()
	pingProbeResultDataSourceRepresentation = map[string]interface{}{
		"probe_configuration_id":              Representation{repType: Required, create: `${oci_health_checks_ping_monitor.test_ping_monitor.id}`},
		"start_time_greater_than_or_equal_to": Representation{repType: Optional, create: strconv.FormatInt(pringProbeStartTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"start_time_less_than_or_equal_to":    Representation{repType: Optional, create: strconv.FormatInt(pringProbeStartTime.Add(5*time.Minute).UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"target":                              Representation{repType: Optional, create: `www.oracle.com`},
	}

	PingProbeResultResourceConfig = PingMonitorRequiredOnlyResource
)

func TestHealthChecksPingProbeResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksPingProbeResultResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_ping_probe_results.test_ping_probe_results"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + compartmentIdVariableStr + PingProbeResultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
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
					generateDataSourceFromRepresentationMap("oci_health_checks_ping_probe_results", "test_ping_probe_results", Optional, Create, pingProbeResultDataSourceRepresentation) +
					compartmentIdVariableStr + PingProbeResultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "probe_configuration_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time_greater_than_or_equal_to"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time_less_than_or_equal_to"),
					resource.TestCheckResourceAttr(datasourceName, "target", "www.oracle.com"),

					resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "ping_probe_results.0.probe_configuration_id"),
				),
			},
		},
	})
}
