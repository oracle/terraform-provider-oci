// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var (
	metricDataEndTimeStr   string
	metricDataStartTimeStr string

	metricDataDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            Representation{repType: Required, create: `${var.compartment_id}`},
		"namespace":                 Representation{repType: Required, create: `oci_vcn`},
		"query":                     Representation{repType: Required, create: `VnicToNetworkPackets[4m].max()`},
		"compartment_id_in_subtree": Representation{repType: Optional, create: `false`},
		"end_time":                  Representation{repType: Optional, create: metricDataEndTimeStr},
		"resolution":                Representation{repType: Optional, create: `2m`},
		"start_time":                Representation{repType: Optional, create: metricDataStartTimeStr},
	}

	MetricDataResourceConfig = AvailabilityDomainConfig
)

func generateMetricDataRepresentationWithCurrentTimeInputs() map[string]interface{} {
	endTime := time.Now()
	startTime := endTime.Add(-2 * time.Hour)
	metricDataEndTimeStr = endTime.Format(time.RFC3339Nano)
	metricDataStartTimeStr = startTime.Format(time.RFC3339Nano)
	return representationCopyWithNewProperties(metricDataDataSourceRepresentation, map[string]interface{}{
		"end_time":   Representation{repType: Optional, create: metricDataEndTimeStr},
		"start_time": Representation{repType: Optional, create: metricDataStartTimeStr},
	})
}

func TestMonitoringMetricDataResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId := getEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_monitoring_metric_data.test_metric_data"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config +
					generateDataSourceFromRepresentationMap("oci_monitoring_metric_data", "test_metric_data", Optional, Create, generateMetricDataRepresentationWithCurrentTimeInputs()) +
					compartmentIdVariableStr + MetricDataResourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "false"),
					resource.TestCheckResourceAttr(datasourceName, "end_time", metricDataEndTimeStr),
					resource.TestCheckResourceAttr(datasourceName, "namespace", "oci_vcn"),
					resource.TestCheckResourceAttr(datasourceName, "query", "VnicToNetworkPackets[4m].max()"),
					resource.TestCheckResourceAttr(datasourceName, "resolution", "2m"),
					resource.TestCheckResourceAttr(datasourceName, "start_time", metricDataStartTimeStr),

					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.0.aggregated_datapoints.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.0.aggregated_datapoints.0.timestamp"),
					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.0.aggregated_datapoints.0.value"),
					resource.TestCheckResourceAttr(datasourceName, "metric_data.0.compartment_id", compartmentId),
					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.0.name"),
					resource.TestCheckResourceAttr(datasourceName, "metric_data.0.namespace", "oci_vcn"),
				),
			},
		},
	})
}
