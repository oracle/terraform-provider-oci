// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"fmt"
	"testing"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
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
		"resource_group":            Representation{repType: Optional, create: `resourceGroup`},
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
	httpreplay.SetScenario("TestMonitoringMetricDataResource_basic")
	defer httpreplay.SaveScenario()

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
					resource.TestCheckResourceAttr(datasourceName, "resource_group", "resourceGroup"),
					resource.TestCheckResourceAttr(datasourceName, "start_time", metricDataStartTimeStr),

					resource.TestCheckResourceAttrSet(datasourceName, "metric_data.#"),
				),
			},
		},
	})
}
