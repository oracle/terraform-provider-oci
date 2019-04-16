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
	startTime                               = time.Now()
	httpProbeResultDataSourceRepresentation = map[string]interface{}{
		"probe_configuration_id":              Representation{repType: Required, create: `${oci_health_checks_http_monitor.test_http_monitor.id}`},
		"start_time_greater_than_or_equal_to": Representation{repType: Optional, create: strconv.FormatInt(startTime.UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"start_time_less_than_or_equal_to":    Representation{repType: Optional, create: strconv.FormatInt(startTime.Add(5*time.Minute).UnixNano()/(int64(time.Millisecond)/int64(time.Nanosecond)), 10)},
		"target":                              Representation{repType: Optional, create: `www.oracle.com`},
	}

	HttpProbeResultResourceConfig = HttpMonitorRequiredOnlyResource
)

func TestHealthChecksHttpProbeResultResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestHealthChecksHttpProbeResultResource_basic")
	defer httpreplay.SaveScenario()

	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_health_checks_http_probe_results.test_http_probe_results"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + compartmentIdVariableStr + HttpProbeResultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					func(s *terraform.State) (err error) {
						if httpreplay.ShouldRetryImmediately() {
							time.Sleep(10 * time.Millisecond)
						} else {
							time.Sleep(2 * time.Minute)
						}
						return nil
					},
				),
			},
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_health_checks_http_probe_results", "test_http_probe_results", Optional, Create, httpProbeResultDataSourceRepresentation) +
					compartmentIdVariableStr + HttpProbeResultResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "probe_configuration_id"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time_greater_than_or_equal_to"),
					resource.TestCheckResourceAttrSet(datasourceName, "start_time_less_than_or_equal_to"),
					resource.TestCheckResourceAttr(datasourceName, "target", "www.oracle.com"),

					resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "http_probe_results.0.probe_configuration_id"),
				),
			},
		},
	})
}
